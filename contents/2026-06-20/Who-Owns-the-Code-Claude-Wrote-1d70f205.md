---
source: "https://www.oreilly.com/radar/who-owns-the-code-claude-wrote/"
hn_url: "https://news.ycombinator.com/item?id=48607274"
title: "Who Owns the Code Claude Wrote?"
article_title: "Who Owns the Code Claude Wrote? – O’Reilly"
author: "Garbage"
captured_at: "2026-06-20T10:05:44Z"
capture_tool: "hn-digest"
hn_id: 48607274
score: 2
comments: 0
posted_at: "2026-06-20T07:58:49Z"
tags:
  - hacker-news
  - translated
---

# Who Owns the Code Claude Wrote?

- HN: [48607274](https://news.ycombinator.com/item?id=48607274)
- Source: [www.oreilly.com](https://www.oreilly.com/radar/who-owns-the-code-claude-wrote/)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T07:58:49Z

## Translation

タイトル: クロードが書いたコードの所有者は誰ですか?
記事のタイトル: クロードが書いたコードの所有者は誰ですか? – オライリー
説明: AI 生成コードの著作権についてビルダー向けに説明。

記事本文:
メインコンテンツにスキップ
エンタープライズ向け
スキルを調べる クラウド コンピューティング Microsoft Azure
データ エンジニアリング データ ウェアハウス
ソフトウェア アーキテクチャ オブジェクト指向
ペネトレーションテスト/エシカルハッキング
ソフトスキル プロフェッショナルコミュニケーション
書籍、コース、イベントなどを検索します
計画
クラウド コンピューティング クラウド コンピューティング
データエンジニアリング データエンジニアリング
プログラミング言語 プログラミング言語
ソフトウェア アーキテクチャ ソフトウェア アーキテクチャ
ペネトレーションテスト/エシカルハッキング
書籍、コース、イベントなどを検索します
ダークモードを切り替える
AI と機械学習
ビジネス
データ
イノベーション
研究
セキュリティ オライリー学習プラットフォームをお試しください
オライリーの学習プラットフォームを使用すると、スキルを磨き、先を行くためのリソースとガイダンスが得られます。最大 14 日間無料でお試しいただけます。
オライリー プラットフォームのライブ オンライン イベントに参加して、テクノロジーを形作る専門家から学びましょう。
レーダートレンドのニュースレターを入手する
当社のプライバシーポリシーをお読みください。
O’Reilly Radar Trends to Watch ニュースレターをご購読いただきありがとうございます。
クロードが書いたコードの所有者は誰ですか?
AI 生成コードの著作権についてビルダー向けに説明。
Sena Evren 著 2026 年 6 月 15 日 • 16 分で読めます
以下の記事は元々、Sena Evren の Legal Layer ニュースレターに掲載されたもので、著者の許可を得てここに再投稿されています。
TL; DR
Claude Code、Cursor、Codex などのエージェント コーディング ツールは、著作権のないコード、雇用主が所有するコード、または目に見えないオープン ソース ライセンスによって汚染されている可能性のあるコードを生成します。この法律には、決着した法律もあれば、活発に争われている法律もあり、この記事ではどちらがどっちであるかは明らかです。 AI 支援コードを配布していて、これについて何も考えたことがない場合は、この記事が最適です。
今週コードを出荷した場合、その一部はおそらく AI によって書かれたものでしょう。そのコードを誰が法的に所有するかという問題は、

ほとんどの開発者が想定しているよりも問題は解決しており、その答えはコードの良さとは関係のない 3 つのことに依存します。
人間が著作権を確立するのに十分な創造的な決定を下したかどうか
雇用契約がすでに雇用主に割り当てられているかどうか
モデルが GPL ライセンスのトレーニング データから取得され、コードベースを密かに汚染したかどうか
2026 年 3 月 31 日、Anthropic は、定期的なソフトウェア アップデートで、構成ファイルが欠落していたため、クロード コードの 512,000 行のソース コードを誤って公開してしまいました。日の出前に、コードベースが GitHub 全体にミラーリングされました。朝食前に、開発者が AI ツールを使用してすべてを Python で書き直したところ、「クローコード」リポジトリは 1 日で GitHub のスターを 100,000 個獲得し、これは史上最速の記録となりました。その後、DMCA の削除が起こり、誰も明確な答えを持っていなかった次のような質問が起こりました。
もし、Anthropic 自身の主任エンジニアが認めたように、Claude Code が主に Claude 自身によって書かれたものであるとしたら、Anthropic はそれを所有しているのでしょうか?著作権法で保護されない可能性のあるコードに対して DMCA 削除を発行できますか?
この事件により、AI が生成したコードの所有権に関するあらゆる未解決の疑問が 1 つのニュース サイクルに圧縮されました。同じ疑問がコードベースにも当てはまります。
誰も教えてくれなかった著作権ルール
法的基準は、簡単に言うと次のとおりです。著作権は人間によって作成された作品のみを保護します。
米国著作権局はこれを一貫して認めており、DC巡回裁判所はセイラー事件でこれを支持した。最高裁判所は2026年3月にセイラー氏の上訴の審理を拒否したが、下級裁判所の推論を支持することも、全国的に問題を解決することもなかった。証明の否認は、裁判所が事件を審理しないことを選択したことを意味するものであり、それ以上のことではありません。これが意味するのは、DC 巡回裁判所の判決は有効であり、著作権局の立場は無傷であり、いかなる権利も持たないということです。

法廷は依然として別の方向に進んでいる。意味のある人間の著作権を持たずに主に AI によって生成された作品は、現在の原則では著作権保護の対象とはならず、最終的に決着がつかなくても、その立場は安定しています。
セイラーが実際に決定したことに対する 2 つの重要な制限。
この事件には、人間の関与がまったくなく描かれた絵画が関係していた。セイラー氏はAIシステムを単独著者として挙げており、人間の創造的貢献については一切主張していない。この判決は、人間が関与したAI支援作業というより難しい問題には直接言及していないが、その関与の程度については議論がある。
セイラーは視覚芸術に携わっていました。 AI コーディング ツールからのコード出力に特に人間の著作者責任の原則を適用した裁判所はまだありません。論理は当てはまりますが、直接の前例はまだ存在していません。
あなたにとっての意味: Claude Code または Cursor が生成し、意味のある変更を加えずに受け入れたコードは、誰にも著作権で保護されない可能性があります。コードは名前以外のすべてがパブリック ドメインにあるため、競合他社がそれをコピーした場合、法的手段がなくなる可能性があります。
コードが保護されているかどうかを決定するフレーズは「意味のある人間の著作権」ですが、著作権局はそれをパーセンテージや編集数で定量化することを意図的に拒否しています。裁判所が求めているのは、人間が真に創造的な決定を下したという証拠であるためです。
特定の設計に合わせて出力を再構成する
モデルの目的を指定するだけでは十分ではありません。作品をどのように構築するかを指示することが重要です。
エージェント ワークフローでは、この区別を確立するのは思っているよりも困難です。典型的なクロード コード セッションを考えてみましょう。
「API のレート制限モジュールを構築する」という 1 行のプロンプトを作成します。
Claude Code はアプローチを計画し、5 つのファイルを生成し、それを繰り返します。

3つのバージョンがあります。
出力を確認し、テストを実行して、マージします。
そのシーケンスにおけるあなたの貢献は、あなたのアーキテクチャ上の意図であり、最終的な承認となります。それが法廷において意味のある人間の著作物に該当するかどうかは未解決の問題であり、最終的な裁判所の判決はまだ出ていない。
正直な答えは、実質的にリダイレクトしたモジュールについてはおそらく「はい」、そのまま受け入れたコードについてはおそらく「いいえ」、そしてその間のすべてについては不明です。
現在、その中間点での訴訟が活発に行われている。アレン対パールマッター事件では、アーティストのジェイソン・アレンが、600 を超える詳細なプロンプトとその後の Photoshop での編集を使用して作成した作品の著作権局の登録拒否に異議を唱えています。著作権局はPhotoshopの編集が人間によって作成されたものであることを認めたが、それでもAIが生成した基本要素の登録は拒否した。この訴訟はまだ判決が下されておらず、どのような判決が下されるにせよ、人間の関与がどの程度であれば十分であるかについての判決に最も近いものとなるだろう。
部分的保護に関する既存の最も近い先例は、グラフィックノベル「Zarya of the Dawn」で、著作権局は人間が書いたテキストの登録を認めたが、ミッドジャーニーが生成した画像については登録を拒否した。この決定により、開発者がすぐに使用できる実用的な原則が確立されます。AI 支援コードベースの人間が作成した要素は、生成されたコード自体が保護できなくても、個別に保護できる可能性があります。アーキテクチャ ドキュメント、コミット メッセージに記録された設計上の決定、ADR、意図的なリダイレクトを示すプロンプト ログなどは、たとえ生成されたコードがそうでなくても、人間が作成した表現として保護できる可能性があります。できることを守るには、実際に何をしたかを文書化することから始まります。
あなたの雇用主がおそらくすでに所有しているもの
あなたのタラかどうかを考える前に

e は著作権で保護されていますが、より差し迫った疑問があります。たとえ著作物であったとしても、それは実際にあなたのものなのでしょうか?
雇用契約には、職場で構築したものはすべて雇用主に属するとほぼ確実に記載されています。この原則は著作権法において「雇用労働原則」と呼ばれています。この規定では、従業員が雇用の範囲内で作成したコードはすべて雇用主が所有し、そのコードが手書きで書かれたものであるか、クロード コードによって生成されたものであるか、またはその組み合わせであるかに関係なく、雇用主が法的作成者として扱われます。勤務時間中、作業プロジェクト、作業マシン上で AI コーディング ツールを使用しても、結果の所有者が変わることはありません。
ほとんどの雇用契約は、原則の規定をさらに超えています。 「知的財産」、「IP 譲渡」、または「成果物」というセクションを探してください。契約書を開いて条件を検索し、そのセクションを読みます。次のいずれかに該当する条項は、ほぼ確実に AI 支援コードに適用されます。
「会社の設備やリソースを使用して作成された成果物」
「雇用期間中に行われた発明または開発」
「企業がライセンスを取得したツールを利用して作成されたソフトウェア」
3番目は注目すべきものです。雇用主がチームに対して Claude Code、Cursor、または Copilot のライセンスを付与しており、あなたがそれらと同じツールを使用してサイド プロジェクトを構築する場合、たとえ自分の時間でプロジェクトを構築したとしても、広範な IP 譲渡条項により雇用主がそのプロジェクトに対する権利を取得できる可能性があります。
サンフランシスコの上級開発者は、今年初めにまさにこの状況を説明しました。彼は、仕事のプロジェクトや、夜や週末に作成した個人的なフィットネス追跡アプリに Claude Code を使用していました。彼の会社は知的財産ポリシーを更新し、個人用アプリを含む彼が AI 支援を利用して構築したすべてのものを主張し、クロードが所有していたものであると主張した。

IDE で作業ファイルを開くためのアクセス権がなければ、AI 出力はすべて企業 IP の派生作品でした。
これは、これがどこまで拡張できるかを示す最も明らかな例です。彼の会社の主張は、AI ツールが彼の会社のコードベースの「コンテキストを認識する」という 1 つのフレーズに基づいていました。この議論は法的には成立しません。IDE でのコンテキストの可視性によって、AI の出力が近くで開かれていたファイルの派生作品にはならず、クロードが見ることができるものとクロードが生成するものとの間の関係はコピーではなく確率的なパターン補完であるためです。しかし、この議論は雇用主が何を主張し始めているかを示している。その条項が十分に広範であれば、AI が実際に何をしたかに関係なく、その条項は表面的な妥当性を持ちます。
実際的なルール: 副業で何かを構築している場合は、自分で支払った個人アカウント、個人用マシン、ツールを使用してください。雇用主のライセンスを取得したツールをそのワークフローから完全に排除してください。
オープンソースの汚染問題
AI が生成したコードを所有している場合でも、目に見えないオープンソース ライセンスによってコードがすでに汚染されている可能性があります。
AI コーディング ツールは、GPL、LGPL、その他のコピーレフト ライセンスに基づいてライセンスされたコードを含む、大量のパブリック コードでトレーニングされます。コピーレフト ライセンスには、コードとともに適用される特定の義務が伴います。
GPL ライセンスのコードから派生したソフトウェアを配布する場合は、同じライセンスの下で独自のソース コードをリリースする必要があります。
これは、組み込んだコードが GPL ライセンスであることを知らなかった場合でも当てはまります。
「知らなかった」ということはコピーレフト違反に対する弁護にはなりません。
AI ツールがトレーニング データから GPL ライセンス コードのかなりの部分をそのまま再現し、そのコードをソースを公開せずに商用製品として出荷すると、元のコードにまったく触れずにコピーレフト違反が作成された可能性があります。

ポジトリ。侵害の法的基準は、機能的な類似性や類似性ではなく、実質的な逐語的な複製であり、この区別が重要です。GPL コードのように機能するコードを生成する AI ツールは、GPL コードを一語一語再現する AI ツールとは異なります。リスクは文字通りその範囲の端にあり、問題は、スキャンを実行しない限り、コードベースがどちら側にあるのかを知る方法がないことです。
シャルデ コミュニティの紛争により、2026 年初頭にこのことが具体化されました。これは提起された訴訟ではなく、法的な解決なしに問題を提起したオープンソース コミュニティ内の公的紛争でした。開発者は、Claude を使用して Python 文字エンコード ライブラリであるchardet を書き換え、MIT ライセンスの下で再リリースし、AI の書き換えは元の LGPL ライセンスを必要としない「クリーン ルーム」実装であると主張しました。
コミュニティが争った法的問題: クロードが LGPL ライセンスのコードベースでトレーニングされ、その出力がそのコードの実質的な部分をそのまま再現した場合、その出力はライセンスフリーとして扱うことができますか?シャルデ紛争はきれいに解決されず、この特定の問題に関して最終的な判決を下した裁判所はありません。解決されたのは、GPL コードをそのままコピーすることは、それがどのように作成されたかに関係なく、ライセンスに違反するということです。

[切り捨てられた]

## Original Extract

AI-generated code copyright explained for builders.

Skip to main content
For Enterprise
Explore Skills Cloud Computing Microsoft Azure
Data Engineering Data Warehouse
Software Architecture Object-Oriented
Penetration Testing / Ethical Hacking
Soft Skills Professional Communication
Search for books, courses, events, and more
Plans
Cloud Computing Cloud Computing
Data Engineering Data Engineering
Programming Languages Programming Languages
Software Architecture Software Architecture
Penetration Testing / Ethical Hacking
Search for books, courses, events, and more
Toggle dark mode
AI & ML
Business
Data
Innovation
Research
Security Try the O’Reilly learning platform
With the O’Reilly learning platform, you get the resources and guidance to keep your skills sharp and stay ahead. Try it free for up to 14 days.
Join a live online event on the O’Reilly platform to learn from the experts shaping tech.
Get the Radar Trends newsletter
Please read our privacy policy .
Thank you for subscribing to the O’Reilly Radar Trends to Watch newsletter.
Who Owns the Code Claude Wrote?
AI-generated code copyright explained for builders.
By Sena Evren June 15, 2026 • 16 minute read
The following article originally appeared on Sena Evren’s Legal Layer newsletter and is being reposted here with the author’s permission.
TL; DR
Agentic coding tools like Claude Code, Cursor, and Codex generate code that may be uncopyrightable, owned by your employer, or contaminated by open source licenses you cannot see. Some of this is settled law, some is actively contested, and this piece is clear about which is which. If you are shipping AI-assisted code and have not thought about any of this, this piece is for you.
If you shipped code this week, some of it was probably written by an AI. The question of who legally owns that code is less settled than most developers assume, and the answer depends on three things that have nothing to do with how good the code is:
Whether a human made enough creative decisions to establish copyright
Whether your employment contract already assigned it to your employer
Whether the model pulled from GPL-licensed training data and quietly contaminated your codebase
On March 31, 2026, Anthropic accidentally published 512,000 lines of Claude Code’s source code in a routine software update through a missing configuration file. Before sunrise, the codebase was mirrored across GitHub. Before breakfast, a developer had used an AI tool to rewrite the entire thing in Python, and the “claw-code” repository hit 100,000 GitHub stars in a single day, the fastest in history. Then came the DMCA takedowns, and then came the question nobody had a clean answer to:
If Claude Code was, by Anthropic’s own lead engineer’s admission, predominantly written by Claude itself, does Anthropic even own it? Can you issue a DMCA takedown for code that copyright law may not protect?
That incident compressed every open question about AI-generated code ownership into a single news cycle. The same questions apply to your codebase.
The copyright rule nobody told you
Here is the legal baseline, in plain terms: Copyright only protects work created by a human .
The US Copyright Office has confirmed this consistently, and the DC Circuit upheld it in the Thaler case. When the Supreme Court declined to hear the Thaler appeal in March 2026, it did not endorse the lower court’s reasoning or settle the question nationally. Cert denial means the court chose not to hear the case, nothing more. What it does mean is that the DC Circuit’s ruling stands, the Copyright Office’s position is intact, and no court has yet gone the other way. Works predominantly generated by AI without meaningful human authorship are not eligible for copyright protection under current doctrine, and that position is stable even if it is not finally settled.
Two important limits on what Thaler actually decided.
The case involved a painting created with zero human involvement at all. Thaler listed the AI system as sole author and made no claim of any human creative contribution. The ruling does not directly address the harder question of AI-assisted work where a human was involved but the degree of that involvement is disputed.
Thaler involved visual art. No court has yet applied the human authorship doctrine specifically to code output from an AI coding tool. The logic applies, but the direct precedent does not exist yet.
What it means for you : Code that Claude Code or Cursor generated and you accepted without meaningful modification may not be copyrightable by anyone. If a competitor copies it, you may have no legal recourse, because the code sits in the public domain in everything but name.
The phrase that determines whether your code is protected is “ meaningful human authorship ,” and the Copyright Office has deliberately refused to quantify it with a percentage or a number of edits, because what courts look for is evidence that a human made genuine creative decisions:
Restructuring the output to fit a specific design
Specifying an objective to the model is not enough. Directing how the work is constructed is what counts.
In an agentic workflow, this distinction is harder to establish than it sounds. Consider a typical Claude Code session:
You write a one-line prompt: “build a rate limiting module for the API.”
Claude Code plans the approach, generates five files, and iterates through three versions.
You review the output, run the tests, and merge.
Your contribution in that sequence is your architectural intent and your final approval. Whether that constitutes meaningful human authorship in a courtroom is an unresolved question with no definitive court ruling yet.
The honest answer is: probably yes for modules you substantially redirected, probably no for code you accepted verbatim, and unclear for everything in between.
The middle ground is actively being litigated right now. In Allen v. Perlmutter, artist Jason Allen is challenging the Copyright Office’s denial of registration for a work he created using more than 600 detailed prompts and subsequent editing in Photoshop. The Copyright Office acknowledged the Photoshop edits as human-authored but still denied registration for the AI-generated underlying elements. That case has not been decided yet, and whatever it decides will be the closest thing to a ruling on how much human involvement is enough.
The closest existing precedent on partial protection is Zarya of the Dawn , a graphic novel where the Copyright Office granted registration for the human-authored text but denied it for the Midjourney-generated images. That decision establishes a practical principle developers can use right now: The human-authored elements of an AI-assisted codebase may be separately protectable even if the generated code itself is not. Your architecture documents, your design decisions recorded in commit messages, your ADRs, your prompt logs showing deliberate redirection, these may be protectable as human-authored expression even if the code they produced is not. Protecting what you can starts with documenting what you actually did.
What your employer probably already owns
Before you think about whether your code is copyrightable, there is a more immediate question: Even if it is, is it actually yours?
Your employment contract almost certainly says that anything you build at work belongs to your employer. That principle has a name in copyright law: the work-for-hire doctrine. Under it, any code created by an employee within the scope of their employment is owned by the employer, who is treated as the legal author, regardless of whether the code was written by hand, generated by Claude Code, or some combination. Using an AI coding tool during work hours, on a work project, on a work machine, does not change who owns the result.
Most employment contracts go further than the doctrine’s defaults. Look for a section in yours called “Intellectual Property,” “IP Assignment,” or “Work Product.” Open the contract, search for those terms, and read that section. A clause that says any of the following almost certainly covers your AI-assisted code:
“Any work product created using company equipment or resources”
“Any invention or development made during the term of employment”
“Any software created with the assistance of company-licensed tools”
The third one is the one to watch. If your employer licenses Claude Code, Cursor, or Copilot for the team, and you use those same tools to build a side project, a broad IP assignment clause may give the employer a claim over that project, even if you built it on your own time.
A senior developer in San Francisco described exactly this situation earlier this year. He had used Claude Code for work projects and for a personal fitness tracking app built on evenings and weekends. His company updated its IP policy and claimed everything he had built with AI assistance, including the personal app, arguing that because Claude had access to open work files in the IDE, any AI output was a derivative work of company IP.
This is the clearest example of how far this can stretch. His company’s claim rested on one phrase: The AI tools were “context-aware” of his company’s codebase. The argument does not hold up legally, because context visibility in an IDE does not make AI output a derivative work of files that were open nearby, and the connection between what Claude can see and what it generates is probabilistic pattern completion, not copying. But the argument illustrates what employers are starting to claim. If the clause is broad enough, it has surface validity regardless of what the AI actually did.
The practical rule : If you are building something on the side, use a personal account, a personal machine, and tools you pay for yourself. Keep your employer’s licensed tools out of that workflow entirely.
The open source contamination problem
Even if you own your AI-generated code, you may have already contaminated it with an open source license you cannot see.
AI coding tools are trained on massive amounts of public code, including code licensed under the GPL, LGPL, and other copyleft licenses. Copyleft licenses carry a specific obligation that travels with the code :
If you distribute software that is a derivative of GPL-licensed code, you must release your own source code under the same license.
This applies even if you did not know the code you incorporated was GPL-licensed.
“I did not know” is not a defense to a copyleft violation.
When an AI tool reproduces a substantial verbatim portion of GPL-licensed code from its training data, and you ship that code in a commercial product without releasing source, you may have created a copyleft violation without ever touching the original repository. The legal standard for infringement is substantial verbatim reproduction, not functional similarity or resemblance, and this distinction matters: an AI tool generating code that works like GPL code is different from an AI tool that reproduces GPL code word for word. The risk sits at the verbatim end of that spectrum, and the problem is that you have no way to know which side of the line your codebase is on without running a scan.
The chardet community dispute made this concrete in early 2026. This was not a filed lawsuit but a public dispute within the open source community that raised the question without resolving it legally. A developer used Claude to rewrite chardet, a Python character encoding library, and rereleased it under an MIT license, arguing that the AI rewrite was a “clean room” implementation free of the original LGPL license.
The legal question the community fought over : If Claude was trained on the LGPL-licensed codebase and its output reproduces substantial verbatim portions of that code, can the output be treated as license-free? The chardet dispute did not resolve cleanly and no court has issued a definitive ruling on this specific question. What is settled is that verbatim copying of GPL code violates the license regardless of how it was prod

[truncated]
