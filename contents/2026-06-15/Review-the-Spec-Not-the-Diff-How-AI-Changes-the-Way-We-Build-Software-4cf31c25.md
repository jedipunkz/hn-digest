---
source: "https://mkdshare.dev/d/review-the-spec-not-the-diff-how-ai-changes-the-way-we-NTdScTRg"
hn_url: "https://news.ycombinator.com/item?id=48540453"
title: "Review the Spec, Not the Diff: How AI Changes the Way We Build Software"
article_title: "Review the Spec, Not the Diff: How AI Changes the Way We Build Software — mkdshare.DEV"
author: "PhilipA"
captured_at: "2026-06-15T14:15:36Z"
capture_tool: "hn-digest"
hn_id: 48540453
score: 2
comments: 0
posted_at: "2026-06-15T12:46:49Z"
tags:
  - hacker-news
  - translated
---

# Review the Spec, Not the Diff: How AI Changes the Way We Build Software

- HN: [48540453](https://news.ycombinator.com/item?id=48540453)
- Source: [mkdshare.dev](https://mkdshare.dev/d/review-the-spec-not-the-diff-how-ai-changes-the-way-we-NTdScTRg)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T12:46:49Z

## Translation

タイトル: 違いではなく仕様を確認してください: AI がソフトウェアの構築方法をどのように変えるか
記事のタイトル: 違いではなく仕様を確認する: AI がソフトウェアの構築方法をどのように変えるか — mkdshare.DEV
説明: アクセス制御とインライン コメントを備えたクリーンで読みやすいドキュメントとしてマークダウンを公開します。手動、API 経由、または Claude、Cursor、Codex などの MCP クライアントから動作します。

記事本文:
mkdshare.DEV
テーマ#切り替え"
class="flex items-center gap-2rounded-lg px-3 py-1.5 text-sm font-medium text-muted transition-colors period-150 hover:bg-well hover:text-ink"
aria-label="テーマを選択">
システム
テーマ#選択"
データテーマ値パラメータ = "ライト"
class="flex w-full items-center justify-between px-3 py-2 text-left text-sm text-bodytransition-colors period-150 hover:bg-well">
ライト
✓
テーマ#選択"
data-theme-value-param="ダーク"
class="flex w-full items-center justify-between px-3 py-2 text-left text-sm text-bodytransition-colors period-150 hover:bg-well">
暗い
✓
Googleでサインイン
公共
フィリップ・アンダーセン著
メニュー#トグル"
class="flex items-center justify-center w-9 h-9 border border-edgerounded-lg hover:bg-well text-muted transition-colors duration-150">
アクション
メニュー#プリントドキュメント"
class="flex items-center gap-2.5 w-full px-4 py-2 text-sm text-body hover:bg-well transition-colors period-150 text-left">
PDFをダウンロード
クリップボード#コピー"
class="flex items-center gap-2.5 w-full px-4 py-2 text-sm text-body hover:bg-well transition-colors period-150 text-left">
マークダウンのコピー
違いではなく仕様を確認する: AI がソフトウェアの構築方法をどう変えるか
私たちのソフトウェア構築方法に何かが静かに崩れ、ほとんどのチームがまだそれに適応できていません。何十年もの間、コードを正直に保つための安全策は、最後にレビューすることでした。つまり、マージする前に人が差分を読むことでした。 AI によってその安全策が改善されたわけでも、悪化したわけでもありません。そのため、現在コードを作成している規模では実用的ではありませんでした。これに代わるものはライフサイクル全体で異なる形状であり、デフォルトで登場する前に理解しておく価値があります。
仕様は徹底的に見直し、コードは軽く見直してください。
角が削られているような気がします。それはその逆です。作業のほとんどを人間ではなく機械が行うと、厳密さはこのようになります。

ピング。
「spec」という単語が多すぎるため、先に進む前に 1 つ説明します。この記事で重要なのは技術的な仕様、つまり提案されたアプローチと変更がどのように構築されるかです。それは製品仕様ではなく、何を構築するのか、そしてなぜ構築するのかを決定する概要です。製品要件はこのステップへの入力であり、主題ではありません。レビュー対象のドキュメントはエンジニアが満足させるための計画であり、その技術計画はコードが存在する前によく読まれるものです。
コードの作成は、これまで仕事の中で時間とコストがかかる部分でした。レビューは、時間のかかる作業中に人々が犯した間違いを見つけるために存在していたので、マージの直前の最後に置かれていました。コードが人間が読める速度とほぼ同じ速度で到着したため、機能しました。プロデューサーとレビュアーは同じペースで走りました。
AI はそのバランスを壊します。コードは、誰もが 1 行ずつ読むよりもはるかに速く生成されるようになりました。書くのは安くなり、量は増えたので、すべてを読むのは安くならず、むしろ高くなりました。査読者はもはや制作者に追いつくことができず、いくら規律を尽くしても桁違いに拡大する差を埋めることはできません。
コストがかかる場所
BEFORE — 人間がコードを書く
─────────────────
仕様 ░░░
████████████のビルドが遅い: 書き込みがボトルネックになっています
██████████のレビューは実現可能です: コードは読み取り速度で到着します
現在 — AI がコードの大部分を作成
─────────────────
スペック ██████████ ここに前払いで投資してください
░░░ 速くて安いものを構築する
████████████████をレビューする

より多くのコードが、より早く到着します
誰でも読めるのでレビューします
すべてのラインを手作業で行うと、次のコストがかかります
以前は、それ以下ではありませんでした。
これは人々が読み間違えている部分です。手動によるコードレビューの費用は安くなりませんでした。コストも高くなり、すべての行を手で読もうとするのはもはや適切ではありません。そのため、人間の少量の注意が結果を変えるところに取り組みが移ります。コードが存在する前に、何を構築するか、どのように構築するかについての決定が行われます。
仕様に引っかかる間違ったアプローチは、編集に 5 分かかります。生成されたコードの山がすでに存在する後に同じアプローチを捕らえると、一日の手直しが必要になるか、レビュー キューが絶望的であるためまったく捕らえられません。プロセスの最前部にレバレッジが移動しました。これは、少しレビューするだけで多くのことがカバーされる最後のポイントであるためです。
ライフサイクルの形
このロジックに従うと、ライフサイクルが自動的に再編成されます。重心はスペックとそのレビューになります。その後のすべては、早期の決断が報われるように存在します。
┌───────────────────────────┐
│ 新たなライフサイクル (7 つのステップ) │
━━━━━━━━━━━━━━━━━━━━━━━━┘
1. 仕様 短い技術仕様: 問題と
│ なぜ 1 ～ 2 行で、提案されたものを
│ アプローチ — 変化をどのように構築するか。
│ 製品概要ではなく、エンジニアリング計画。
▼
2. スペックレビュー ◀══════ ★メインクオリティゲート★
│ 事前にピアが技術仕様をレビューします
│任意のコード。アプローチに合わせて、間違って殺す
│

アイデアを安価に変更できるうちに。
▼
3. 承認された仕様に基づいてビルドを実装します。
│ タイピングのほとんどを AI が行います。人間
│ はすべての行を所有し、理解しています。
▼
4. 受け入れチェック (試せる UX がある場合のみ)
│ ··· 任意 誰かが機能を試して確認する
│ それは仕様が意図したとおりに動作します。小切手
│ コードではなく結果について。
▼
5. SHIP 自動レビューはすべてに対して実行されます。
│ 人間によるレビューは一か八かの場合のみ
│変化↓
│ ┌─────────────────┐
│ │ 一か八かの賭けですか？ → 人間のレビュー │
│ │ マージ前に必須 │
│ ━━━━━━━━━━━━━━┘
▼
6. 検証 ステージングで動作することを確認してから、ステージングで動作することを確認します。
│制作、発表前。
▼
7. 発表する 検証されたら、人々に知らせます。
ステップは、その下の動きほど重要ではありません。人間による重いゲートは、プロセスの終わりから始まり、コードからインテントに移行します。
無謀ではなく安全な理由は、レビューが消えないからです。それぞれの種類の査読者が実際に得意とすることによって、それは 2 つに分かれます。
機械がコードをレビューします。静的分析、セキュリティ スキャン、依存関係チェック、およびテストは、例外なくすべての変更に対して毎回実行されます。これは、AI が生成するボリュームに応じて拡張できる唯一の種類のレビューであり、仕様レビューでは決して検出できない実装レベルの欠陥 (インジェクション、アクセス チェックの欠落、安全でない依存関係、壊れたエッジ ケース) を検出する層です。それはあらゆるものの下にある床です。
人間は意図とコードをレビューするのは、間違いが高くつく場合のみです。現在、人間の注意力は希少で高価なリソースなので、爆発範囲が広い場所にそれを費やし、

自動化と仕様レビューで残りをカバーします。
すべての変更 → 自動レビュー
(静的解析、セキュリティスキャン、テスト)
│ 常に、例外なく
▼
┌───────────────┐
│ これは一か八かの賭けですか │
│変化しますか？ │
━───────┬───────┘
はい ┌──┴──┐ いいえ
▼ ▼
┌─────────┐ ┌─────────────┐
│ + ヒューマンレビュー │ │ 個別に発送 │
│ マージ前 │ │ (自動レビュー │
│ │ │のみ）│
━───┬───┘ ━━━━━┬───┘
━─────┬───────┘
▼
検証する
一か八かの勝負は、特徴のサイズではなく、爆発の半径についてです。
• お金の動き
• 認証と認可
• 個人データまたは機密データ
• 他のすべてが依存する共有サービス
重要なのは、1 つのルールをすべてのコードに適用するのではなく、リスクによってプロセスを重み付けすることです。影響範囲が小さい変更は、金銭の移動やデータ漏洩の可能性がある変更と同じ税金を支払うべきではありません。ほとんどの変更は一か八かの賭けではないため、ほとんどの変更は自動化のみで出荷されます。
機械が機械的なチェックを所有すると、残る人間のレビューはスタイルに関するものではありません。それは自動化では提供できない判断に関するものです。
┌─────────────────────┐
│ 人間によるレビューとは何のためにあるのか │
━━━━━━━━━━━━━━━━━━━━━━┤
│ ✓

これは正しい問題を解決しているのでしょうか? │
│ ✓ アプローチは健全ですか? │
│ ✓ ユーザーにとって正しく動作しますか? │
│ ✓ アーキテクチャと意図は維持されていますか? │
━━━━━━━━━━━━━━━━━━━━━━┤
│ ✗ 非: 中括弧のスタイル、名前付け、行ごと │
│ メソッドの批評 (自動化がこれを実現します) │
━━━━━━━━━━━━━━━━┘
これらはツールでは答えられない質問であり、モデル自体の作業についての答えをモデルが信頼できるものではありません。その下にある機械層はすべて自動化に属します。
メカニカル レイヤーには、オートメーションよりも優れた家、つまりモデル自体の命令があります。命名、構造、ハウス スタイルなどの規則は、エージェントの常駐命令である CLAUDE.md または同等の命令に属しているため、AI は、間違った命令を作成したり、リンターに依存してフラグを立てたりするのではなく、最初の草稿にそれらを書き込みます。
これは、フィードバックのカテゴリー全体を再構成したものです。繰り返し現れる細かい問題は、差分の問題ではなく、指示のギャップです。指示を一度修正すれば、プロジェクトの存続期間中、変更のたびに追われることなく、この種の間違いはまったく発生しなくなります。最も安価な欠陥は、モデルが決して書き込まない欠陥です。
変わらないもの: 所有権
人間による一行ごとのレビューなしで出荷できるプロセスは、説明責任が明確な場合にのみ機能します。 3 つのことが真実である必要があり、それらは AI が関与しているかどうかに関係なく同じです。
仕様は短くしてください。実行時間が長い仕様は正しく読み取られず、目的が果たせません。問題、アプローチ、結果を述べる

来て、立ち止まってください。
共有する前に自分の仕様を読んでください。 AI が起草したかどうかにかかわらず、自分で読んでいないものを渡さないでください。あなたがそれを所有しています。
配布するコードを理解します。モデルが書いたとしても、それを理解する必要があります。 「AI が書いた」ということは、本番環境では答えになりません。
これらが防御しているのは、AI 支援作業の本当の失敗モードです。これは遅いコードではなく、所有されていないコードです。仕様書は誰も読まず、コードは誰も理解せず、機械が生成したために出荷されたものです。自主性と信頼が高まるのは、責任が伴うからです。
┌─────────────┬─────────────┐
│ 以前 │ 今 │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤
│ 人間がコードをレビューします │ 人間が同意します │
│ 最後 │ 前からのアプローチ │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤
│ 多くの場合、合意が得られない │ 事前に調整されたアプローチ │
│ 最初のアプローチ │ 任意のコードが存在する │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤
│ │ 一度独立して出荷するときの手直し │
│ アプローチが間違っていた │ 仕様は承認されている │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤
│ 一部の変更を出荷 │ 自動化によるすべてのレビュー │
│ まったくレビューなし │ コード。人間

レビューのみ │
│ │ 一か八かの変化 │
━━━━━━━━━━━━━━━━━━━━━━━━┘
旧モデルには静かな故障がありました。最後にある重いゲートは、納品のプレッシャーで人々がスキップするゲートだったので、チームはレビューと未レビューのコードの摩擦を同時に受けることがよくありました。新しいモデルは、1 つの信頼性の低い遅延ゲートを 2 つの信頼できるゲートと交換します。つまり、すべての自動化と、リスクが高い場合の人間の判断です。
単一チームのプロセスから一歩下がって、AI がコーディングを行う場合はどこでも 3 つのシフトが発生します。
ボトルネックは、コードの作成から、何を書くかの決定に移りました。コードが安価で大量にある場合、すべてを手で読むことは不可能になるため、機械が行レベルのレビューを担当し、人間は意図を正しく理解することに集中します。
コントロールは均等に適用されるのではなく、リスクによって重み付けされます。仕様といくつかのリスクの高い領域に重点を置き、あらゆる部分を自動化して、残りの部分には軽く触れます。
予防は検出に勝ります。品質を仕様やモデルの指示に早く組み込むほど、下流で同じ間違いを何度も発見するために支払う費用が少なくなります。
これはハードルを下げるものではありません。現在作業が行われている場所にバーを移動します。 「仕様を厳しく見直し、コードを軽く見直す」というのは基準の緩和ではありません。これは、それらをプロセスの最前部に再配置することであり、そこでは、少量の人間の判断によって、次の 1,000 行が書く価値があるかどうかが決定されます。
自由にサインインしてこのドキュメントにコメントし、アクセス制御を使用して独自のマークダウンを公開します。
Googleでサインイン
© 2026 mkdshare.DEV

## Original Extract

Publish markdown as clean, readable documents with access control and inline comments. Works manually, through the API, or from MCP clients like Claude, Cursor, and Codex.

mkdshare.DEV
theme#toggle"
class="flex items-center gap-2 rounded-lg px-3 py-1.5 text-sm font-medium text-muted transition-colors duration-150 hover:bg-well hover:text-ink"
aria-label="Choose theme">
System
theme#select"
data-theme-value-param="light"
class="flex w-full items-center justify-between px-3 py-2 text-left text-sm text-body transition-colors duration-150 hover:bg-well">
Light
✓
theme#select"
data-theme-value-param="dark"
class="flex w-full items-center justify-between px-3 py-2 text-left text-sm text-body transition-colors duration-150 hover:bg-well">
Dark
✓
Sign in with Google
Public
by Philip Andersen
menu#toggle"
class="flex items-center justify-center w-9 h-9 border border-edge rounded-lg hover:bg-well text-muted transition-colors duration-150">
Actions
menu#printDocument"
class="flex items-center gap-2.5 w-full px-4 py-2 text-sm text-body hover:bg-well transition-colors duration-150 text-left">
Download PDF
clipboard#copy"
class="flex items-center gap-2.5 w-full px-4 py-2 text-sm text-body hover:bg-well transition-colors duration-150 text-left">
Copy Markdown
Review the Spec, Not the Diff: How AI Changes the Way We Build Software
Something quietly broke in how we build software, and most teams have not adjusted to it yet. For decades the safeguard that kept code honest was review at the end: a person read the diff before it merged. AI did not make that safeguard better or worse. It made it impractical at the scale we now produce code. What is replacing it is a different shape for the whole lifecycle, and it is worth understanding before it arrives by default.
Review the spec heavily, the code lightly.
It sounds like a corner being cut. It is the opposite. It is what rigor looks like once a machine, not a person, is doing most of the typing.
One clarification before going further, because the word spec is overloaded. The spec this piece cares about is the technical one: the proposed approach and how the change will be built. It is not the product spec — the brief that decides what to build and why. Product requirements are the input to this step, not its subject. The document under review is the engineer's plan for satisfying them, and that technical plan is what gets read heavily before any code exists.
Writing code used to be the slow and expensive part of the job. Review existed to catch the mistakes people made while doing that slow work, so it sat at the end, right before merge. It worked because code arrived at roughly the speed a person could read it. The producer and the reviewer ran at the same pace.
AI breaks that balance. Code now gets produced far faster than anyone can read it line by line. Writing got cheap and the volume went up, so reading all of it got more expensive, not less. The reviewer can no longer keep up with the producer, and no amount of discipline closes a gap that is growing by orders of magnitude.
WHERE THE COST LIVES
BEFORE — humans write the code
───────────────────────────────────────────────
spec ░░░
build ████████████ slow: writing is the bottleneck
review ██████████ feasible: code arrives at reading speed
NOW — AI writes most of the code
───────────────────────────────────────────────
spec ██████████ invest here, up front
build ░░░ fast and cheap
review ████████████████ more code, arriving faster than
anyone can read it, so reviewing
every line by hand costs more than
before, not less.
This is the part people misread. Manual code review did not get cheaper. It got more expensive, and trying to read every line by hand is the thing that no longer fits. So the effort moves to where a small amount of human attention still changes the outcome: the decision about what to build and how to structure it, made before any code exists.
A wrong approach caught in a spec is a five-minute edit. The same approach caught after a pile of generated code already exists is a day of rework, or it never gets caught at all because the review queue is hopeless. The leverage moved to the front of the process because that is the last point where reviewing a little still covers a lot.
The shape the lifecycle is taking
When you follow that logic through, the lifecycle reorganizes itself. The center of gravity becomes the spec and its review. Everything after exists to make that early decision pay off.
┌─────────────────────────────────────────────────────────────┐
│ THE EMERGING LIFECYCLE (7 steps) │
└─────────────────────────────────────────────────────────────┘
1. SPEC A short TECHNICAL spec: the problem and
│ why in a line or two, then the proposed
│ approach — how the change will be built.
│ The engineering plan, not the product brief.
▼
2. SPEC REVIEW ◀══════ ★ THE MAIN QUALITY GATE ★
│ A peer reviews the technical spec BEFORE
│ any code. Align on the approach, kill wrong
│ ideas while they're cheap to change.
▼
3. BUILD Implement against the approved spec.
│ AI does most of the typing. A human
│ owns and understands every line.
▼
4. ACCEPTANCE CHECK (only if there's UX to try)
│ ··· optional Someone tries the feature and confirms
│ it does what the spec intended. A check
│ on outcome, not on code.
▼
5. SHIP Automated review runs on everything.
│ Human review only for high-stakes
│ changes ↓
│ ┌─────────────────────────────────┐
│ │ HIGH-STAKES? → human review │
│ │ required before merge │
│ └─────────────────────────────────┘
▼
6. VALIDATE Confirm it works in staging, then in
│ production, before announcing it.
▼
7. ANNOUNCE Tell people once it is validated.
The steps matter less than the move underneath them: the heavy human gate shifts from the end of the process to the beginning, and from the code to the intent.
The reason this is safe and not reckless is that review does not disappear. It splits in two, by what each kind of reviewer is actually good at.
Machines review the code. Static analysis, security scanning, dependency checks, and tests run on every change, every time, with no exceptions. This is the only kind of review that scales with the volume AI produces, and it is the layer that catches the implementation-level defects a spec review can never see: injection, a missing access check, an unsafe dependency, a broken edge case. It is the floor under everything.
Humans review the intent, and the code only where a mistake is expensive. Human attention is the scarce, costly resource now, so you spend it where the blast radius is high, and let automation plus the spec review cover the rest.
Every change → AUTOMATED REVIEW
(static analysis, security scan, tests)
│ always, no exceptions
▼
┌───────────────────────┐
│ Is this a HIGH-STAKES │
│ change? │
└───────────┬───────────┘
yes ┌──┴──┐ no
▼ ▼
┌─────────────────┐ ┌──────────────────────┐
│ + HUMAN REVIEW │ │ Ship independently │
│ before merge │ │ (automated review │
│ │ │ only) │
└────────┬────────┘ └───────────┬──────────┘
└──────────┬────────────┘
▼
VALIDATE
High-stakes is about blast radius, not feature size:
• Money movement
• Authentication and authorization
• Personal or sensitive data
• The shared services everything else depends on
The point is to weight the process by risk instead of applying one rule to all code. A change with a small blast radius should not pay the same tax as one that can move money or leak data. Most changes are not high-stakes, so most changes ship on automation alone.
Once machines own the mechanical checks, the human review that remains is not about style. It is about judgment that automation cannot supply.
┌────────────────────────────────────────────────┐
│ WHAT A HUMAN REVIEW IS FOR │
├────────────────────────────────────────────────┤
│ ✓ Is this solving the RIGHT problem? │
│ ✓ Is the approach SOUND? │
│ ✓ Will it behave correctly FOR THE USER? │
│ ✓ Does the architecture & intent hold up? │
├────────────────────────────────────────────────┤
│ ✗ NOT: brace style, naming, line-by-line │
│ method critique (automation does this) │
└────────────────────────────────────────────────┘
These are questions a tool cannot answer and a model cannot be trusted to answer about its own work. Everything below them, the mechanical layer, belongs to automation.
The mechanical layer has an even better home than automation: the model's own instructions. Conventions like naming, structure, and house style belong in the agent's standing instructions, the CLAUDE.md or its equivalent, so the AI writes to them in the first draft instead of producing them wrong and relying on a linter to flag them.
This reframes a whole category of feedback. A nitpick that keeps reappearing is not a problem with the diff, it is a gap in the instructions. Fix the instructions once and that class of mistake stops being generated at all, instead of being caught on every change for the life of the project. The cheapest defect is the one the model never writes.
What does not change: ownership
A process that lets people ship without line-by-line human review only works if accountability is explicit. Three things have to stay true, and they are the same whether or not an AI is involved.
Keep the spec short. A spec that runs long does not get read properly, which defeats the purpose. State the problem, the approach, and the outcome, and stop.
Read your own spec before you share it. Do not pass along something you have not read yourself, AI-drafted or not. You own it.
Understand the code you ship. If a model wrote it, you still have to understand it. "The AI wrote it" is not an answer in production.
What these guard against is the real failure mode of AI-assisted work, which is not slow code but unowned code: specs nobody read and code nobody understands, shipped because a machine produced it. Autonomy and trust go up only because accountability goes up with them.
┌──────────────────────────┬──────────────────────────┐
│ BEFORE │ NOW │
├──────────────────────────┼──────────────────────────┤
│ Humans review CODE at │ Humans agree on the │
│ the end │ APPROACH up front │
├──────────────────────────┼──────────────────────────┤
│ Often no agreement on │ Approach aligned before │
│ the approach first │ any code exists │
├──────────────────────────┼──────────────────────────┤
│ Rework when the │ Ship independently once │
│ approach was wrong │ the spec is approved │
├──────────────────────────┼──────────────────────────┤
│ Some changes shipped │ Automation reviews all │
│ with no review at all │ code; humans review only │
│ │ the high-stakes changes │
└──────────────────────────┴──────────────────────────┘
The old model had a quiet failure. The heavy gate at the end was the one people skipped under delivery pressure, so teams often got the friction of review and unreviewed code at the same time. The new model trades one unreliable late gate for two dependable ones: automation on everything, and human judgment where the stakes are high.
Step back from any single team's process, and three shifts show up wherever AI does the coding.
The bottleneck moved from writing code to deciding what to write. When code is cheap and plentiful, reading all of it by hand stops being possible, so machines take the line-level review and humans concentrate on getting the intent right.
Controls get weighted by risk instead of applied evenly. Heavy attention on the spec and on the few high-stakes areas, automation everywhere, and a light touch on the rest.
Prevention beats detection. The earlier you push quality, into the spec, into the model's instructions, the less you pay to catch the same mistake over and over downstream.
None of this lowers the bar. It moves the bar to where the work now happens. "Review the spec heavily, the code lightly" is not a relaxation of standards. It is a relocation of them, to the front of the process, where a small amount of human judgment still decides whether the next thousand lines were worth writing.
Free to sign in and comment on this document — and publish your own markdown with access control.
Sign in with Google
© 2026 mkdshare.DEV
