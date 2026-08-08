---
source: "https://curtispoe.org/paad/"
hn_url: "https://news.ycombinator.com/item?id=49225778"
title: "How to write production-quality code with AI"
article_title: "PAAD — Defense-in-Depth for AI-Assisted Development · Interactive Course"
author: "Ovid"
captured_at: "2026-08-08T21:20:09Z"
capture_tool: "hn-digest"
hn_id: 49225778
score: 3
comments: 1
posted_at: "2026-08-08T20:50:42Z"
tags:
  - hacker-news
  - translated
---

# How to write production-quality code with AI

- HN: [49225778](https://news.ycombinator.com/item?id=49225778)
- Source: [curtispoe.org](https://curtispoe.org/paad/)
- Score: 3
- Comments: 1
- Posted: 2026-08-08T20:50:42Z

## Translation

タイトル: AI を使用して本番品質のコードを記述する方法
記事タイトル: PAAD — AI支援開発のための多層防御・インタラクティブコース

記事本文:
著者からのメモ
この文書は何なのか、そして何ではないのか
PAAD は、AI コーディング アシスタントのための一連のスキルであり、AI 主導のエンジニアリングではなく、エンジニアリング主導の AI という 1 つのアイデアに基づいて構築されています。エンジニアは決定を下し、変更する時間がまだある間にそれぞれを確認します。違いが現れるのは技術的負債です。PAAD のレビュー ゲートにより、仕様駆動型開発がより厳密になるため、負債の蓄積がより緩やかになります。また、エージェント アーキテクチャにより、経験豊富な開発者は後で蓄積する負債を管理する方法が提供されます。このステップは、AI 主導のエンジニアリングには含まれていません。 PAAD を実行するために必要なすべてがこのコースに含まれています。各スキルの機能、そのように設計されている理由、その背後にある証拠などです。これは文書でもあり、文書は特定の種類の教育ツールの 1 つです。この最初のセクションでは、この文書が優れている点と、この文書がインストラクターによるトレーニングの代わりにならない理由について説明します。
ドキュメントがうまく機能していることを、このドキュメントは完全に実行しようとしています。これは参照可能であり、検索可能で、引用可能で、チームにループを導入する前の夜に再読可能です。自分のペースとタイムゾーンで動作し、配布に費用はかからず、安定しています。設計のディスカッションや意思決定記録にリンクできます。
ただし、どんなに注意深く書かれた文書でもできないこともあります。これらは媒体の特性であり、材料のギャップではありません。 PAAD は実践するものであり、学ぶべき一連の事実ではありません。エージェントの計画に反論する必要があることを読むのは、途中で誰かが修正してくれるライブセッションでそれを行うのとは異なるスキルです。ここでの例は必然的に一般的なものであるか、私自身の研究から抜粋されたものですが、これらのいずれかを採用する価値があるかどうかを決定するケースは、制約の下でコードベースに存在します。エージェントの場合

奇妙なことに、リアルタイムで一緒に奇妙なことを乗り越えてきたし、これからもそうなるだろうが、これは価値が高く、散文では再現できない。そして、文書は「しかし、私たちの状況はどうですか」という答えを返すことも、それが理解されたかどうかを確認することもできません。「私はそれを読んで、私はそれを理解しました」ということは内部からは見えません。これがインストラクターの存在理由の大部分を占めています。
さらに 2 つの制限は、個人ではなく組織的なものです。方法論はチームの財産です。5 人が別々に読むと 5 つの個人的な解釈が生まれ、チームが一緒に資料を読み進めると 1 つの合意された規範が生まれます。そして、導入が停滞するときは、期待の不一致、所有権の不明確さ、上級エンジニアがまだ納得していないなど、組織上の理由で停滞することがほとんどです。これらは正当な障害であり、文書では見えず、ライブディスカッションで対処できます。そしてドキュメントは古くなります。これらのモジュールで名前が付けられているツールは毎月変わります。このページは、それが書かれた日の時点で正確です。今日の会話は正確です。オプションの付録「チームにこれを教える」では、これについてさらに詳しく説明しています。内容を理解した後もチームが何をしなければならないか、およびその各部分を誰が所有する必要があるかについて説明します。
これらすべての結果の 1 つは、行動する価値があります。ドキュメントはあなたの声を聞くことができません。そのため、モジュールがあなたを混乱させた場合、スキルが説明どおりに動作しなかった場合、またはこれを読んでから実行するまでのどこかで行き詰まった場合、読者ができることは私に伝えることが最も有益であり、それがこのコースを改善する方法です。ドキュメントでは提供できない上記の項目については、ライブ トレーニングが存在します。独自のコードでの練習、自分の状況に関する質問、複数の個人的な読み物ではなく 1 つの共有規範を使用し、それをアレンジしたり、フィードバックを送信したりすることは、両方とも同じ場所、つまり LinkedIn から始まります。もしそうなら

あなたはトレーニングを企画する人ではなく、エンジニアとしてこれを読んでいますが、それはあなたの問題ではありません。以下に続く内容はすべて現状のままで完了しており、役立つと思われる場合は同僚に簡単に転送できます。
AI は問題に振りかける魔法の妖精の粉ではありません。これはツールであり、その逆ではなく、エンジニアリング上の判断によって動かされる場合に最も効果的に機能します。 PAAD は、その方向を向いた配置を維持するためのツールキットです。以下の 2 つのセクションにある「PAAD とは」セクションで完全に説明されています。
AI コーディング ツールを広範囲に使用しているものの、コードを慎重にレビューしていない開発者は、何か不快なことを発見します。初日は魔法のように感じます。最初の 1 か月は生産的です。しかし、時間が経つにつれて、コードベースが認識できない場所になっていることに気づきます。エラーを処理する 3 つの異なる方法、2 か所にハードコーディングされた構成値、何もアサートせずに合格するテスト、および要求されたことは技術的に実行するが意図したことは実行しない機能。
劇的に問題が起こることはありませんでした。それが問題のすべてだ。 1 回の劇的な変更でコードの品質が低下することはほとんどありません。それは、それぞれ単独では合理的であると思われる一連の小さな決定をすり抜けます。アシスタントは毎回、彼らの要求に応えましたが、それらの合理的な答えをすべて合計すると、彼らが決して受け入れることを選択しなかった負債を抱えたコードベースになります。
ここで問題が発生したのは、実際にはコードの品質の問題ではありません。誰が決定を下したのかが問題です。開発者はアシスタントに何をすべきかを伝え、アシスタントが担当してくれることを望みました。アシスタントは言われたとおりに実行し、それはいつもそうなのですが、開発者はコードが書かれて初めて何を決定したかを知りました。
それが AI 主導のエンジニアリングです。アシスタントがペースを設定し、エンジニアは下流で自分たちが経験したことのない決定を下します。

で。 PAAD — AI コーディング アシスタントのためのオープンソース プロジェクトで、詳細は次のセクションで説明します — はその逆を提案しており、その README にはそれがキャッチフレーズとして記載されています。
AI主導のエンジニアリングではなく、エンジニアリング主導のAI。 PAAD の README
同じ速度で反対方向。 AI を変更できるうちに AI が何をしようとしているのかを見て、決定を下すのはあなたです。具体的には、作業の各段階で何か異なることが起こることを意味します。右側の列の名前は PAAD の 4 つのコア スキルです。名前で呼び出すもので、それぞれに後で独自のモジュールが付いています。
右側の列を下に読むと、このコースの残りの部分の形状がわかります。各モジュールは 1 行をカバーします。
運転席に留まることがポイントであり、コストでもあります。 PAAD は自動操縦ではなく、可視性と制御を提供します。アシスタントに考えてもらいたいのであれば、これは間違ったツールキットです。ツールキットに含まれるすべてのスキルは、マージではなく人間の決定で終わります。 README には、その結果が明確に記載されています。報告を無視することは、報告をまったく行わないよりも悪いということです。
AI コーディング アシスタントはそのプロセスを大幅に圧縮できますが、弱い要件に確実に挑戦したり、計画からの逸脱を検出したり、長期的なコード品質を保護したりすることはできません。 PAAD の README
PAAD (「パッド」と発音) は、Curtis "Ovid" Poe によって作成された、GitHub 上にある MIT ライセンスの AI エージェント スキルのオープンソース システムです。これは AI 支援開発ツールに代わるものではありません。AI 支援開発ツールを補完するものであり、その README では、AI 支援開発ツールと併用し続けることができるものの例として、コーディング エージェント用の別個のオープンソース ワークフロー ツールキットである Superpowers を挙げています。モジュール 3 では、スーパーパワーと他のツールキットを紹介します。 4 つの文字は 4 つの一般的な障害モードの名前を示しており、各障害モードにはそれに対処するスキルがあります。
スキルはマークダウン命令です。

ファイル (従来は SKILL.md) であり、コーディング アシスタントがオンデマンドでロードしてこれに従います。それはコードでもモデルでもありません。これは文書化された手順です。ユーザーがこれを要求した場合、その方法は次のとおりであり、何を確認するかは次のとおりで、報告の方法は次のとおりです。最新のアシスタントのほとんどは、何らかのバージョンの形式をサポートしています。
PAAD はそのうちのおよそ 12 つです。このコース全体を通じて、小文字で書かれたスキルはスキル自体 (プッシュバック、アラインメント) を意味し、先頭にスラッシュが付いた同じ単語 /pushback はアシスタントに入力することを意味します。わかりやすい英語で質問することもできます。モジュール 11 では、それらのインストールと、それらを呼び出す 3 つの方法について説明します。それまでは、名前を指示ではなくラベルとして扱ってください。
アイデアを明確にするアナロジー
この方法論を紹介する記事では、注目に値する歴史的な類似点が描かれています。 1950 年代に Fortran が導入されたとき、コンパイラーは手書きのアセンブラーほど効率的なアセンブラーを生成できないため、IBM は開発者が反対することを知っていました。そのため、マーケティングでは、生成されたコードがほぼ同じ速度であることに注目しながら、開発者の時間を節約することに重点を置きました。
コンパイラーが人間が作成したアセンブリよりも高速に実行されるアセンブリを日常的に作成するようになるまで、数十年かかりました。現在、コンパイラを上回るパフォーマンスのアセンブリ コードを作成できる人間はほとんどいません。これが私たちの AI の現場です。カーティス・ポー「クロード・ソネットがオーパスを上回るパフォーマンスを鑑賞」
この主張は、AI があなたよりも優れたコードを書くということではありません。主張は、規律 (PAAD の D) を使えば、自分が書いたコードとほぼ同じ品質のコードを生成し、かかる時間のほんのわずかな時間で生成できるということです。このアナロジーは、コードの実行速度ではなく、コードが作成される速度についてのものです。そして、コンパイラと同様に、品質の差は縮まると予想されます。
それは良いたとえです、そして他の良いアナと同様に

当然のことながら、賞賛するのではなくチェックする必要があります。付録「証拠はどの程度優れていますか?」まさにそれを行います。
創設記事のタイトルは「クロード・ソネットがオーパスを上回るパフォーマンスを観察する」です。この結果はまさに、PAAD プロセスを実行する小型で安価なモデル (Sonnet) の方が、PAAD プロセスを使用せずに実行するより大型でより高性能なモデル (Opus) よりも優れた結果をもたらしたということです。方法論はモデルの能力に勝ります。
また、これは単一チームの現場レポートであり、制御条件なしで受け売りで中継されたものであり、Sonnet と Opus のコントラストは設計されたものではなく偶然であり、IDE のバグにより一時的にチームが Opus から締め出されたために生じました。記事自身の評決は正しいもので、「ベンチマークではなく、1つのチームからの『実際の状況』レポート」である。それを念頭に置いてください。付録「証拠はどの程度優れていますか?」独立した研究に反することになります。
4 つの文字を順番に処理し、それらを機能ごとにまとめてループにまとめます。 2 つのモジュールは、その引数を進める以外のことを行います。モジュール 3 は、PAAD 固有ではなく前提条件です。これは、仕様駆動開発について説明しています。PAAD のプラクティスは、あなたがすでに従っていることを前提としています。今日 spec-kit、Kiro、または OpenSpec を実行し、その理由がわかっている場合は、スキップしてください。モジュール 4 は、出力をどれほど注意深く確認しても、モデルが認識できないものを尋ねるために停止します。その過程で知識チェックが行われます。これらは採点されません。レビュー プロセスについて読むことと、レビュー プロセスを実行できることは別のスキルであるため、存在します。最後に最終評価があります。
評価の後は、オプションの付録があります。そのうちの 1 つは、「証拠はどの程度優れているか?」というもので、議論全体を独立した研究文献に照らしてテストします。コースではそれらに依存するものは何もなく、評価の質問もそれらに依存しません。この方法を採用してください。

必要に応じて信頼学を参照し、事件のどの部分が実際に十分な証拠があるかを知りたい場合に参照してください。もう 1 つの「寓話のレビューで見つかったこと」では、モジュール 2 のケーススタディの背後にある調査結果が印刷されているため、信頼に頼らずに自分で調査結果を判断できます。もう 1 つの「なぜ小さい仕様なのか」では、これらのゲートはすべて、大きな成果物よりも小さな成果物に対してはるかにうまく機能すると主張しています。仕様を作成した後ではなく、最初の仕様を作成するときに読んでください。
PAAD の使用経験がないことを前提としています。 AI コーディング アシスタント (Claude Code、Cursor、Kiro、Antigravity) にある程度の知識があると役に立ちますが、アイデアはどのアシスタントにも伝わります。
PAAD は、軽量の「ただ構築するだけ」のワークフローよりも多くのトークンを使用します。プロジェクトはこれについて明確にしています。最適化は、トークンの消費量を最小限にするためではなく、より適切な意思決定と回避可能な間違いを減らすために最適化されます。
具体的には、コードを作成するだけでなく、仕様のレビュー、計画のレビュー、コードのレビューに料金を支払うことになるため、1 つの機能に対してより多くのトークンのコストがかかります。これは、間違ったものを出荷して再構築するよりも安価であることに賭けます。通常、何年にもわたって保守する運用ソフトウェアについては、この賭けが正しいでしょう。廃棄する予定のプロトタイプ、スパイク (1 つの質問に答える短期間の使い捨て実験)、または

[切り捨てられた]

## Original Extract

A note from the author
What this document is — and what it is not
PAAD is a set of skills for AI coding assistants, built on one idea: engineering-driven AI, not AI-driven engineering. The engineer makes the decisions, and sees each one while there is still time to change it. Technical debt is where the difference appears: PAAD's review gates make specification-driven development more rigorous, so debt accumulates more slowly, and agentic-architecture later gives an experienced developer a way to manage the debt that does accumulate — a step AI-led engineering does not include. Everything you need to run PAAD is in this course: what each skill does, why it is designed the way it is, and the evidence behind it. It is also a document, and a document is one particular kind of teaching tool. This first section describes what this document does well, and why it cannot replace training with an instructor.
What a document does well, this one tries to do in full. It is a reference: searchable, quotable, re-readable the night before you introduce the loop to your team. It works at your pace and in your timezone, it costs nothing to distribute, and it is stable — you can link to it in a design discussion or a decision record.
Some things, though, no document can do, however carefully written. They are properties of the medium, not gaps in the material. PAAD is something you practise, not a set of facts to learn. Reading that you should argue with an agent's plan is a different skill from doing it in a live session, with someone there to correct you halfway through. The examples here are necessarily generic or drawn from my own work, while the cases that decide whether any of this is worth adopting live in your codebase, under your constraints. When an agent behaves strangely — and it will — working through the strangeness together in real time is high-value and unrepeatable in prose. And a document cannot answer "but what about our situation," nor check whether it has been understood: "I read it, and I understand it" is invisible from the inside, which is a large part of why instructors exist.
Two further limits are organisational rather than individual. A methodology is a team property — five people reading separately produce five private interpretations, where a team working through the material together produces one agreed norm. And when adoption stalls, it usually stalls for organisational reasons: mismatched expectations, unclear ownership, a senior engineer who is not yet convinced. Those are legitimate obstacles, they are invisible to a document, and they can be addressed in a live discussion. And a document goes out of date: the tools named in these modules change monthly. This page is accurate as of the day it was written; a conversation is accurate today. The optional appendix “Teaching this to your team” covers this in more detail: what a team still has to do once the material is understood, and who has to own each piece of it.
One consequence of all this is worth acting on. A document cannot hear you — so if a module confused you, a skill did not behave as described, or you got stuck somewhere between reading this and running it, telling me is the most useful thing a reader can do, and it is how this course improves. Live training exists for the pieces named above that a document cannot provide — practice on your own code, questions about your situation, one shared norm instead of several private readings — and arranging that, or sending that feedback, both start in the same place: LinkedIn . If you are reading this as an engineer rather than the person who arranges training, none of that is your problem: everything that follows is complete as it stands, and easy to forward to a colleague if it seems useful.
AI is not magic pixie dust you sprinkle on a problem. It is a tool, and it works best when your engineering judgement drives it rather than the other way around. PAAD is a toolkit for keeping the arrangement pointing in that direction. The section "What PAAD is", two sections below, explains it in full.
Developers who use AI coding tools extensively – but don't carefully review their code – discover something unpleasant. The first day feels like magic. The first month is productive. But over time, they notice that the codebase has become a place they no longer recognise. Three different ways of handling errors, a config value hard-coded in two places, tests that pass without asserting anything, and a feature that technically does what they asked but not what they meant.
Nothing went dramatically wrong. That is the whole problem. Code quality rarely degrades in one dramatic change; it slips through a series of small decisions that each seemed reasonable in isolation. The assistant did what they asked, every time, and the sum of all those reasonable answers is a codebase carrying debt they never chose to accept.
What went wrong there is not really a code-quality problem. It is a question of who was making the decisions. The developer told the assistant what to do and hoped it would take charge. The assistant did as it was told — it always does — and the developer found out what it had decided only once the code was written.
That is AI-driven engineering : the assistant sets the pace, and the engineers are downstream of decisions they never saw being made. PAAD — an open-source project for AI coding assistants, explained in full in the next section — proposes the inverse, and its README puts it as a tagline:
Engineering-driven AI, not AI-driven engineering. PAAD README
Same speed, opposite direction. You see what the AI is about to do while you can still change it, and you make the decisions. Concretely, that means something different happens at each stage of the work. The names in the right-hand column are PAAD's four core skills — things you invoke by name, each with a module of its own later:
Read the right-hand column downward and you have the shape of the rest of this course. Each module covers one row.
Staying in the driver's seat is the point, and it is also the cost. PAAD gives you visibility and control, not autopilot. If what you want is for the assistant to think for you, this is the wrong toolkit — every skill in it ends with a human decision rather than a merge. The README states the consequence plainly: a report you ignore is worse than no report at all.
AI coding assistants can compress that process dramatically, but they do not reliably challenge weak requirements, detect drift from the plan, or protect long-term code quality on their own. PAAD README
PAAD (pronounced "pad") is an open-source system of AI agent skills — MIT-licensed, on GitHub , created by Curtis "Ovid" Poe. It does not replace your AI-assisted development tools — it complements them, and its README names Superpowers , a separate open-source workflow toolkit for coding agents, as an example of something you can keep using alongside it. Module 3 introduces Superpowers with the other toolkits. The four letters name four common failure modes, and each failure mode has skills that address it.
A skill is a markdown instruction file — conventionally SKILL.md — that your coding assistant loads on demand and follows. It is not code, and it is not a model. It is a written procedure: when the user asks for this, here is how to do it, here is what to check, here is how to report back. Most modern assistants support some version of the format.
PAAD is roughly a dozen of them. Throughout this course, a skill written in lowercase means the skill itself — pushback , alignment — and the same word with a leading slash, /pushback , means typing it to your assistant. You can also just ask in plain English. Module 11 covers installing them and the three ways to invoke them ; until then, treat the names as labels rather than instructions.
The analogy that makes the idea clear
The article introducing the methodology draws a historical parallel worth noting. When Fortran was introduced in the 1950s, IBM knew developers would object, because a compiler could not produce assembler as efficient as hand-written assembler. So the marketing focused on saving developer time , while noting the generated code was almost as fast.
It was decades before compilers routinely wrote assembly that ran faster than human-written assembly. Today, very few humans can write assembly code that outperforms a compiler. This is where we are with AI. Curtis Poe, "Watching Claude Sonnet Outperform Opus"
The claim is not that AI writes better code than you. The claim is that with discipline — the D in PAAD — it produces code that is almost as good as what you would have written, and produces it in a fraction of the time it would have taken you. The analogy is about how quickly the code gets written, not how quickly it runs. And as with compilers, the quality gap is expected to narrow.
It is a good analogy, and like all good analogies it should be checked rather than admired. The appendix “How good is the evidence?” does exactly that.
The founding article is titled "Watching Claude Sonnet Outperform Opus." The finding is exactly that: a smaller, cheaper model (Sonnet) running the PAAD process produced better outcomes than a larger, more capable model (Opus) running without it. Methodology beat model capability.
It is also a single team's field report, relayed second-hand, with no control condition — and the Sonnet/Opus contrast was accidental rather than designed, arising because an IDE bug temporarily locked the team out of Opus. The article's own verdict is the right one: "a 'boots on the ground' report from one team, not a benchmark." Keep that in mind. The appendix “How good is the evidence?” sets it against the independent research.
You will work through the four letters in order, then assemble them into the per-feature loop that ties them together. Two modules do something other than advance that argument. Module 3 is prerequisite rather than PAAD-specific: it explains spec-driven development , the practice PAAD assumes you are already following — skip it if you run spec-kit, Kiro, or OpenSpec today and know why. Module 4 stops to ask what the model cannot see no matter how carefully you review its output. Along the way there are knowledge checks — they are not graded, they exist because reading about a review process and being able to run one are different skills. There is a final assessment at the end.
After the assessment there are optional appendices . One of them, “How good is the evidence?”, tests the whole argument against the independent research literature. Nothing in the course depends on any of them and no assessment question draws on them — take the methodology on trust if you prefer, and go there when you want to know which parts of the case are actually well evidenced. Another, “What the Fable review found”, prints the findings behind the case study in Module 2, so you can judge them yourself instead of taking a count on trust. Another, “Why small specs”, argues that all of these gates work far better on small artefacts than large ones — read it when you are writing your first spec rather than after.
No prior experience with PAAD is assumed. Some familiarity with an AI coding assistant — Claude Code, Cursor, Kiro, Antigravity — will help, but the ideas transfer to any of them.
PAAD uses more tokens than a lightweight "just build it" workflow, and the project is explicit about this: it optimises for better decisions and fewer avoidable mistakes, not for minimum token consumption.
Concretely: a single feature will cost you more tokens because you are paying to review the spec, review the plan, and review the code — on top of writing it. The bet is that this is cheaper than shipping the wrong thing and rebuilding it. That bet is usually right for production software you will maintain for years. It is a bad bet for a prototype you plan to throw away, a spike (a short throwaway experiment that answers one question), or a

[truncated]
