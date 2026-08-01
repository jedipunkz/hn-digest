---
source: "https://weeraman.com/the-prototype-isnt-the-product/"
hn_url: "https://news.ycombinator.com/item?id=49132130"
title: "AI doesn't generate working products, that's still your job"
article_title: "The Prototype Isn't the Product | Anuradha Weeraman"
author: "smckk"
captured_at: "2026-08-01T08:03:45Z"
capture_tool: "hn-digest"
hn_id: 49132130
score: 2
comments: 1
posted_at: "2026-08-01T07:52:10Z"
tags:
  - hacker-news
  - translated
---

# AI doesn't generate working products, that's still your job

- HN: [49132130](https://news.ycombinator.com/item?id=49132130)
- Source: [weeraman.com](https://weeraman.com/the-prototype-isnt-the-product/)
- Score: 2
- Comments: 1
- Posted: 2026-08-01T07:52:10Z

## Translation

タイトル: AI は実用的な製品を生成しません、それは依然としてあなたの仕事です
記事タイトル: プロトタイプは製品ではない |アヌラーダ・ウィーラマン
説明: AI により、最初の実用バージョンへの道のりが劇的に加速されました。最初の実用バージョンと実稼働グレードのものとの間の距離は縮まりませんでした。

記事本文:
プロトタイプは製品ではありません |アヌラーダ・ウィーラマン
アヌラーダ・ウィーラマン
技術者・建築家・創設者
← すべての投稿 2026 年 3 月 9 日 プロトタイプは製品ではありません
ソフトウェアの構築において、これほどアクセスしやすいとは感じたことはありませんでした。アイデアを平易な英語で説明すると、数分以内に実際に動作するプロトタイプが画面に表示されます。 UIがあります。データベースに接続します。それはあなたが想像したことを実現します。コード行を一度も書いたことがない人にとって、その瞬間は魔法のように感じられます。何年もコンパイラとスタック トレースと格闘してきた人にとって、これは本当に驚くべきことです。
プロトタイプはラップトップ上で実行されます。負荷がかかると壊れます。エラー処理はありません。 API トークンが漏洩している可能性があることがわかりました。データ モデルはデモでは理にかなっていましたが、2 人目のユーザーを追加した瞬間に崩れてしまいます。認証は仮定とともに行われます。すべてが安全かどうかというしつこい不安があります。それを導入したいと思っているのに、突然「これは機能する」と「これは準備ができている」の間にある溝を見つめることになります。
プロトタイプを作成するのは決して難しいことではありませんでした
ソフトウェア エンジニアは、常に何かを迅速に実行することができました。時間がかかったのは他のすべてでした。つまり、大規模に耐えられるシステムの設計、ユーザーが遭遇するはずがなかったが必然的に発生するケースへの対処、いつ問題が発生するかを知るための可観測性の組み込み、3 年後に後悔することの少ないデータ アーキテクチャに関する慎重な決定を下すことです。それは何も変わっていません。 AI により、最初の実用バージョンへの道のりが劇的に加速されました。最初の実用バージョンと実稼働グレードのものとの間の距離は縮まりませんでした。
混乱が生じるのは、旅の初期部分のフィードバック ループが非常に高速になり、非常にやりがいのあるものになったためです。あなたは尋ね、あなたは受け取り、あなたは結果を見る

ts.このサイクルは本当に刺激的で、残りのソフトウェア開発も同様に圧縮する必要があるという印象を与えます。そうではありません。ソフトウェア構築の難しい問題は、構文の記述が主なものではありませんでした。何を構築するか、どのように構築するか、何を延期するか、いつノーと言うかなど、判断に関するものでした。その判断が、バイブコーダーを彫刻家、そして職人に変えるのです。これは、プロトタイプと実稼働グレードのシステムを区別するものでもあります。
コンピューターサイエンスの学習に対する誤った事例
予想通り、AI が生成したコードのアクセシビリティは、コンピューター サイエンスを学ぶことにまだ意味があるのかと疑問を抱く業界への新規参入者の波を引き起こしました。アプリケーションを動作させる方法を説明できるのであれば、アルゴリズム、データ構造、オペレーティング システム、理論の研究に何年も費やす必要はありません。
コンピューター サイエンス教育の価値は、コードを作成する能力だけにあるわけではありません。それは、システムがどのように動作するか、どのように失敗するか、そしてその理由についてのメンタル モデルを開発することでした。このモデルにより、AI が生成したコードを見て、AI が作成したクエリによって 5,000 万行のテーブルでフル テーブル スキャンが発生することを認識できるようになります。これにより、提案されたキャッシュ戦略が同時負荷の下で競合状態を引き起こすことがわかります。これは、提案されたアーキテクチャがあなたが説明した問題を解決しますが、次の問題を大幅に困難にすることを示しています。
その基礎がなければ、モデルの判断に完全に依存することになります。そしてモデルには判断力がありません。彼らはパターン マッチングを備えており、ユーザーの意図と一致すると思われるコードを生成することに熱心です。彼らは、見た目が正しく、慣例に従い、本番環境では失敗するコードを自信を持って生成します。原因が分からない場合は、診断に何日もかかります。

あなたが探しているところで。
理解と成果の間のギャップが崩壊した今は、コンピューター サイエンスを学ぶのに歴史上最適な時期であることは間違いありません。分散システムがどのように機能するかを真に理解した学生は、10 年前に比べてわずかな時間で分散システムを構築できるようになりました。
何が変わり、何が変わらないのか
コードを機械的に書くことしかできず、要件を一行ずつ実装に変換するエンジニアの需要は、本当に減少しています。仕事のその部分は自動化されています。
起こっていることは、生産性分布の下限の圧縮と、上位の人々の上限の拡大です。最新の AI ツールを使用する経験豊富なエンジニアは、5 年前には想像もできなかったペースで作業を進めることができます。難しい問題がなくなったからではなく、多大な時間と注意力を費やしていた機械的な作業が大幅に処理されたからです。実際に専門知識が必要な作業に費やす時間は 1 日により多くなります。
取り残されるエンジニアはAIスキルのないエンジニアではありません。彼らは、理解の代わりに AI を使用し、論理的に理解できないシステムをバイブコーディングして、壊れたものを修正できず、成長したものを拡張できず、保守する必要がある人に自分が構築したものを説明できないことに気づく人たちです。
異なるレベルでの運用方法を学ぶ
必要な変化は、新しいツールセットを採用することではありません。それは、基本に根ざしたものを維持しながら、より高い抽象レベルで動作することです。この組み合わせは本当に強力で、本当に珍しいものです。
今後数年間で他のエンジニアを飛び越えるエンジニアは、AI を深い知識に代わるものとしてではなく、それを強化するものとして扱う人たちです。彼らは自分たちが何を求めているかを理解しています

製作するモデル。彼らは、ジュニア エンジニアのプル リクエストに適用するのと同じ批判的な目で、生成されたコードをレビューします。これらは、機能の説明だけでなく、建築上の考え方を会話に持ち込んでいます。彼らは、モデルが示唆することをいつ押し返すべきかを知っています。
これは、古いスキルセットを新しいコンテキストに適用し、劇的に活用できるようにしたものです。
プロトタイプは簡単な部分です。今の違いは、誰もがそれをはっきりと認識できることです。プロトタイプの後に来るものは依然として困難であり、依然として実際のエンジニアリング上の判断が必要であり、信頼性の高いソフトウェアを出荷するビルダーとデモを出荷するビルダーを依然として分離しています。
基礎を学びましょう。次に、新しいツールを学びます。その順番で。
前へ 次へ 複雑な問題: AI が生成したコードベースが静かに劣化している理由 Anuradha Weeraman 創設者、CTO、Debian 開発者。拡張可能なインテリジェントな製品とシステムを構築します。

## Original Extract

AI has dramatically accelerated the path to a first working version. It has not shortened the distance between a first working version and something production-grade.

The Prototype Isn't the Product | Anuradha Weeraman
Anuradha Weeraman
Technologist · Architect · Founder
← All posts March 9, 2026 The Prototype Isn't the Product
Building software has never felt this accessible. You describe an idea in plain English, and within minutes, a working prototype appears on your screen. It has a UI. It connects to a database. It does the thing you imagined. For someone who has never written a line of code, that moment feels like magic. For someone who has spent years wrestling with compilers and stack traces, it's genuinely astonishing.
The prototype runs on your laptop. It breaks under load. It has no error handling. You find out that it may be leaking your API tokens. The data model made sense for the demo but falls apart the moment you add a second user. The authentication is held together with assumptions. There's a nagging worry whether everything is secure. You want to deploy it, and suddenly you're staring at a chasm between "this works" and "this is ready."
Getting to a prototype was never the hard part
Software engineers have always been able to get something running quickly. What took time was everything else: designing systems that hold up at scale, handling the cases users weren't supposed to encounter but inevitably do, building in observability so you know when things break, making deliberate decisions about data architecture that you'll regret less three years from now. None of that has changed. AI has dramatically accelerated the path to a first working version. It has not shortened the distance between a first working version and something production-grade.
The confusion arises because the feedback loop for the early part of the journey has become so fast and so rewarding. You ask, you receive, you see results. That cycle is genuinely exciting, and it creates the impression that the rest of software development must be similarly compressed. It isn't. The hard problems of building software were never primarily about writing syntax. They were about judgment: what to build, how to structure it, what to defer, when to say no. That judgement is what turns a vibe-coder into a sculptor, and an artisan. It is also what differentiates a prototype from a production-grade system.
The misguided case against learning computer science
Predictably, the accessibility of AI-generated code has sparked a wave of new entrants to the industry who are questioning whether it still makes sense to learn computer science. If you can describe your way to a working application, why spend years studying algorithms, data structures, operating systems, and theory?
The value of a computer science education was never purely in the ability to produce code. It was in developing a mental model of how systems behave, how they fail, and why. That model is what allows you to look at AI-generated code and recognize that the query it wrote will cause a full table scan on a table with fifty million rows. It's what allows you to see that the caching strategy it proposed will create a race condition under concurrent load. It's what tells you that the architecture it suggested solves the problem you described but will make the next problem significantly harder.
Without that foundation, you are entirely dependent on the model's judgment. And models don't have judgment. They have pattern matching, with an eagerness to produce code that it believes matches your intent. They will confidently generate code that looks right, follows convention, and fails in production in ways that take days to diagnose if you don't know what you're looking for.
Now is arguably the best time in history to learn computer science, because the gap between understanding and output has collapsed. A student who genuinely grasps how a distributed system works can now build one in a fraction of the time it would have taken a decade ago.
What changes, and what doesn't
The demand for engineers who can only write code mechanically, and who translates requirements into implementations line by line, are genuinely declining. That part of the job is being automated.
What's happening is a compression of the lower end of the productivity distribution and an expansion of the ceiling for those at the top. An experienced engineer using modern AI tools can move at a pace that would have been unimaginable five years ago. Not because the hard problems have disappeared, but because the mechanical work that consumed so much time and attention is largely handled. More hours in the day for the work that actually requires expertise.
The engineers who will be left behind are not those who lack AI skills. They're the ones who use AI as a substitute for understanding, who vibe-code their way through systems they can't reason about, and then find themselves unable to fix what breaks, unable to scale what grows, unable to explain what they built to anyone who needs to maintain it.
Learning to operate at a different level
The shift required isn't about adopting a new set of tools. It's about operating at a higher level of abstraction while keeping your roots in the fundamentals. That combination is genuinely powerful and genuinely rare.
The engineers who will leapfrog their peers in the next few years are the ones who treat AI as a force multiplier on deep knowledge rather than a replacement for it. They understand what they're asking the model to produce. They review generated code with the same critical eye they'd apply to a junior engineer's pull request. They bring architectural thinking to the conversation, not just feature descriptions. They know when to push back on what the model suggests.
It's the old skill set, applied to a new context, with dramatically higher leverage.
The prototype is the easy part. The difference now is that everyone can see that clearly. What comes after the prototype is still hard, still requires real engineering judgment, and still separates the builders who ship reliable software from the ones who ship demos.
Learn the fundamentals. Then learn the new tools. In that order.
Previous Because We Can Next The Compounding Problem: Why Your AI-Generated Codebase Is Quietly Deteriorating Anuradha Weeraman Founder, CTO and Debian Developer. Building intelligent products and systems that scale.
