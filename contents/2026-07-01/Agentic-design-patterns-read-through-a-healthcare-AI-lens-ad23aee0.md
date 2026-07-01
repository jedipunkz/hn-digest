---
source: "https://jenniferjiangkells.com/thoughts/agentic-patterns-healthcare-lens/"
hn_url: "https://news.ycombinator.com/item?id=48752601"
title: "Agentic design patterns, read through a healthcare AI lens"
article_title: "Agentic design patterns, read through a healthcare AI lens"
author: "adjks"
captured_at: "2026-07-01T20:33:16Z"
capture_tool: "hn-digest"
hn_id: 48752601
score: 1
comments: 0
posted_at: "2026-07-01T20:21:48Z"
tags:
  - hacker-news
  - translated
---

# Agentic design patterns, read through a healthcare AI lens

- HN: [48752601](https://news.ycombinator.com/item?id=48752601)
- Source: [jenniferjiangkells.com](https://jenniferjiangkells.com/thoughts/agentic-patterns-healthcare-lens/)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T20:21:48Z

## Translation

タイトル: エージェントのデザイン パターン、ヘルスケア AI のレンズを通して読み解く
説明: 一般的なエージェント エンジニアリング パターンに再度慣れるために、効果的な AI エージェントの構築に関する Anthropic のガイドを読みました。全体的にシンプルで分かりやすかったです...

記事本文:
エージェントのデザイン パターン、ヘルスケア AI のレンズを通して読み取る |ジェニファー・ジャン・ケルズ
ジェニファー・ジャン・ケルズ
について
執筆
思い
講演とプレス
エージェントのデザイン パターン、ヘルスケア AI のレンズを通して読み取る
一般的なエージェント エンジニアリング パターンに再度慣れるために、効果的な AI エージェントの構築に関する Anthropic のガイドを読みました。全体的にシンプルで簡潔で、楽しく読めました。最もエレガントな技術文書は、最も単純なアドバイスを提供していることが多いと思います。
私は、ガイド内の各パターンを取り上げて、それに対するヘルスケアの使用例を見つけるという、機械的なことを計画していました。予想外だったのは、練習の途中で気分が切り替わってしまったことです。 3 番目のパターンのあたりで、パターンは興味深い部分ではなくなり、別の質問が静かに引き継がれました。医療におけるどの問題が実際に検証可能ですか?それは引っ張る価値のあるスレッドであることが判明し、この投稿はここで終わります。しかし、私がそれを見つけた順序は要点の半分なので、実際に起こった方法を見てみましょう。
エージェント システムの最初の原則についての私の要点: シンプル > 複雑、透明性 > 抽象化。エージェントは、環境と対話するために解放される、ツール、メモリ、および取得機能を備えた単なる LLM であることを覚えておくことが重要です。この技術には、設計を過度に複雑にすることなくユースケースに合わせて調整するという制約があり、ツールを明確に文書化して説明する忍耐力があります。 LLM をジュニア開発者と考えるのは、私にとってとても魅力的です。
ここで、各パターンがヘルスケア AI のレンズを通して着地できると私は考えます。 (ヘルスケア AI を現実世界に導入した経験のある者として、これによってインフラストラクチャとプライバシーに関する山ほどの疑問が生じることは承知していますが、ここでは純粋にグリーンフィールドを想像しています。)
音声から臨床文書を生成します。 T

相談内容を文字に起こし、ステップを連鎖させて生のテキストを SOAP のような固定構造に従うメモに整形します。各ステップで 1 つのジョブ (文字起こし → 構造 → 検証) が実行されます。
臨床試験の基準をわかりやすい言葉に翻訳します。分厚い医療専門用語を、読みやすい一般人向けの要約に段階的に翻訳し、チェーン内の各リンクが複雑さを取り除きます。
医療 Q&A トリアージ。ロジスティックな質問 (「予定はいつですか?」) をデータ取得ツールに送信し、一般的な健康に関する質問を LLM に送信し、臨床的に複雑な質問や曖昧な質問を人間に送信します。ここで、ルーティングは安全機構です。
次のいくつかのパターンは、ユースケースを見つけるのに苦労しました。それらはすべて、余分なステップを含むルーティングのように見え、ガイド自身の例は、進むにつれてより抽象的になっていきました。おそらく重要なのは、各ワークフローに 1 つのユースケースを組み込むことではなく、要件に応じて成長する単純なものから複雑なものへの卒業とみなすことなのかもしれません。ヘルスケアのような一か八かの分野では、追加できるガードレールが常に 1 つあるため、これを特定するのは困難です。ここで本当の疑問が生じます。いつになったら十分なのでしょうか? eval の世界に入りましょう。 (これについては後で詳しく説明します。)
このガイドの例の多くはリーン コーディングに関連しています。 AI 支援の開発ツールが急増していること、そして今年の Code with Claude でダリオ・アモデイがそれをどのように組み立てたかを考えれば、驚くことではありません。コーディングは検証可能であるため最初の橋頭堡であり、次のフロンティアはより多くのドメインを検証可能にすることです。
これにより、私は FHIR を新たな観点から再考するようになりました。もしかしたら、それは必ずしも面倒なことではなく、目的を達成するための手段である必要もないのかもしれません。 FHIR は標準化された構造化された JSON です。それが検証可能にし、医療における特定の問題も検証可能にします。
そこで私は質問をリダイレクトしました。どのエージェントが

医療におけるユースケースは実際に検証可能ですか?
フリーテキストを FHIR リソースに変換します。何かが返される前に、生成された FHIR JSON のさまざまな側面 (構造、コード、エンコードされた値) をそれぞれ検証する並列サブ呼び出しを実行します。この構造により、具体的なチェック対象が得られます。
ソース全体の健康記録を集約します。オーケストレーターはワーカーをファンアウトし、各ワーカーが異なるシステムからレコードを取得して正規化し、それらを 1 つの一貫した画像に調整します。
私はこれについて考えるのにかなりの時間を費やしました。エバリュエーターは、次の場合に最も役立ちます。
a) 明確な品質バーがあり、それに対して反復処理を行うと、出力が目に見えて改善されます。
b) フィードバックは、LLM が人間の監督なしで独自に与えることができるものです。
最もわかりやすい例は匿名化です。ジェネレーターは臨床ノートの編集されたバージョンを生成します。評価者は、残留 PHI (紛らわしい名前、MRN、本来あるべきではない場所にある日付) がないかスキャンし、漏洩したものはすべて返却します。ジェネレーターは再び編集し、パスがクリーンに戻るまでループが繰り返されます。バーはほぼ 2 値であり (PHI はまだ存在するか、はいかいいえか)、フィードバックは LLM が自動的に与えることができるもので、各ラウンドは前回よりも確実に優れています。明確に定義された成功基準があれば、閉じて一日中実行できるループです。
臨床コーディングも同様の形をしています。メモから ICD コードを生成し、それぞれがテキストでサポートされているか、十分に具体的であるかを評価し、修正します。ただし、フィードバックが臨床的判断を必要とする瞬間 (「これは安全ですか」、「これは正しい判断ですか」)、ループは人間に戻される必要があります。
これは、私が繰り返し考えていることを示しています。ヘルスケア AI は、すべてが均一に高リスクである一枚岩ではありません。スペクトルがあり、評価者が最適化する

r パターンは、リスクが低く、自己検証可能な目的で正確に機能します。
ここには 2 本のストランドがあり、同じスペクトルの反対側に位置します。
臨床上の意思決定を行うエージェントは、人間が常に関与しなければ実行できません。行動の場は無制限であり、賭け金は誰かの健康であり、「検証可能」は完全に自動化できるものではなくなります。
しかし、医療データを照合するエージェントは、特に FHIR のような標準化され構造化された形式の場合、完全にエージェント的になる可能性があると思います。これは、上記のオーケストレーターとワーカーの例をさらに一歩進めたもので、違いは自律性です。ワークフローは制限されています: ファンアウト、プル、正規化、合成、完了。完全なエージェントは、調整が無制限である場合に使用します。次に何をフェッチするかを決定し、ソース間の競合が見つかったときにそれを解決し、事前に決定されたステップ数なしでデータが一貫するまでループし続けます。
重要なのは、出力がまだ構造化されており、チェック可能であるということです。これが、ここでの自律性が安全である理由です。タスクが検証可能であればあるほど、エージェントはより多くの自律性を保持できます。
本当の制限は、それを実行するためのテスト環境です。実際の PHI に触れる前に、医療データの忠実性を十分に強調し、グリーンフィールドの方向で実験するのに十分な厳密なサンドボックスがあるでしょうか?それは想像するのは簡単ですが、稼ぐのは難しいです。セ・ラ・ヴィ！
それでは、私たちはどうなるでしょうか？医療を強制するパターンのチェックリストではなく、検証可能性を通じて獲得する自律性の勾配のようなものです。ヘルスケアの構造化された部分 (FHIR、データ調整) は、ユーザーが検証できる部分であり、ユーザーが検証できる部分はエージェントが所有できる部分です。臨床判断の最終段階が人間のままであるのは、エージェントが到達できないからではなく、私たちがまだ測定できないからです

そこは彼らを信頼するのに十分です。
そしてここで、先ほど止めた質問に一周して戻ります。いつになったら十分ですか?検証可能性は、その質問にまったく答えられるものにするものです。チェックできない領域では、「十分な品質」を測定することはできません。 eval の世界をチェックしてみましょう!
ソフトウェア、ヘルスケア、AI に関するメモ

## Original Extract

I read Anthropic’s guide on Building Effective AI Agents to re-familiarize myself with common agentic engineering patterns. The whole thing was simple, conci...

Agentic design patterns, read through a healthcare AI lens | Jennifer Jiang-Kells
Jennifer Jiang-Kells
About
Writing
Thoughts
Talks & press
Agentic design patterns, read through a healthcare AI lens
I read Anthropic’s guide on Building Effective AI Agents to re-familiarize myself with common agentic engineering patterns. The whole thing was simple, concise, and a pleasure to read. I find that the most elegant technical writing often offers the simplest advice.
I went in planning to do something mechanical: take each pattern in the guide and find a healthcare use case for it. What I didn’t expect was for the exercise to flip on me partway through. Somewhere around the third pattern, the patterns stopped being the interesting part, and a different question quietly took over: which problems in healthcare are actually verifiable? That turned out to be the thread worth pulling, and it’s where this post ends up. But the order I found it in is half the point, so let me walk through it the way it actually happened.
My takeaway for the first principles of agentic systems: simple > complex , transparency > abstraction . It’s important to remember that agents are just LLMs with tools, memory, and retrieval, set loose to interact with an environment. The craft is in the restraint of tailoring to a use case without overcomplicating the design, and patience to document and explain your tools clearly. It’s very endearing to me to think of LLMs as junior devs.
So here’s where I think each pattern could land through a healthcare AI lens. (As someone who’s deployed healthcare AI into the real world, I know this raises a mountain of infrastructure and privacy questions — but here I’m imagining purely greenfield.)
Generating clinical documents from speech. Transcribe a consultation, then chain steps to shape the raw text into a note that follows a fixed structure like SOAP — each step doing one job (transcribe → structure → validate).
Translating clinical-trial criteria into plain language. A staged translation from dense medical jargon to a readable, layperson-friendly summary, each link in the chain stripping away a layer of complexity.
Medical Q&A triage. Send logistical questions (“when’s my appointment?”) to a data-fetch tool, general health queries to an LLM, and anything clinically complex or ambiguous to a human. Here the routing is the safety mechanism.
The next few patterns I struggled to find use cases for — they all looked like routing with extra steps, and the guide’s own examples got more abstract as it went. Then it clicked: maybe the point isn’t to slot one use case into each workflow, but to see them as a graduation from simple → complex that grows with your requirements. That’s hard to pin down in a high-stakes domain like healthcare, where there’s always one more guardrail you could add. Which raises the real question: when is it good enough ? Enter the world of evals. (More on that later.)
Many of the guide’s examples also lean coding-related. Unsurprising, given the explosion of AI-assisted dev tools, and given how Dario Amodei framed it at this year’s Code with Claude : coding was the first beachhead because it was verifiable , and the next frontier is making more domains verifiable.
This made me reconsider FHIR in a new light. Maybe it doesn’t have to be a pain in the ass and a means to an end. FHIR is standardized, structured JSON. That makes it verifiable — and that makes certain problems in healthcare verifiable too.
So then I redirected the question: which agentic use cases in healthcare are actually verifiable ?
Converting free text to FHIR resources. Run parallel sub-calls that each validate a different facet of the generated FHIR JSON — structure, codes, encoded values — before anything is returned. The structure gives you something concrete to check against.
Aggregating health records across sources. An orchestrator fans out to workers that each pull and normalize records from a different system, then reconciles them into one coherent picture.
I spent a good chunk of time thinking about this one. The evaluator is most useful when:
a) there’s a clear quality bar and iterating against it measurably improves the output, and
b) the feedback is something an LLM can give on its own, without human supervision.
The cleanest example is de-identification. A generator produces a redacted version of a clinical note; an evaluator scans it for any residual PHI — a stray name, an MRN, a date sitting where it shouldn’t — and hands back whatever leaked; the generator redacts again, and the loop repeats until the pass comes back clean. The bar is almost binary (is there still PHI, yes or no?), the feedback is something an LLM can give itself, and each round is verifiably better than the last: a loop you can close and run all day, given clearly defined success criteria.
Clinical coding has the same shape: generate ICD codes from a note, evaluate whether each is supported by the text and specific enough, revise. The moment the feedback requires clinical judgment, though (“is this safe”, “is this the right call”?), the loop has to open back up to a human.
This points to something I keep circling back to: healthcare AI isn’t a monolith where everything is uniformly high-risk. There’s a spectrum, and the evaluator–optimizer pattern works precisely on the low-risk, self-verifiable end of it.
Two strands here, and they sit at opposite ends of that same spectrum.
Agents that make clinical decisions can never run without a human in the loop, full stop. The action space is open-ended, the stakes are someone’s health, and “verifiable” stops being something you can fully automate.
But agents that reconcile healthcare data can, I think, be fully agentic — especially when they are in a standardized, structured format like FHIR. This is a step beyond the orchestrator–workers example above, and the difference is autonomy. Workflows are bounded : fan out, pull, normalize, synthesize, done. A full agent is for when reconciliation is open-ended : it decides what to fetch next, resolves conflicts between sources as it finds them, and keeps looping against the data until it’s coherent, with no predetermined number of steps.
The key is that the output is still structured and checkable, which is why that autonomy is safe here: the more verifiable the task, the more autonomy an agent can hold.
The real limitation is the test environment to run it in. Do we have rigorous enough sandboxes to fully stress healthcare-data fidelity, to experiment in a greenfield direction, before any of this touches real PHI? That’s easy to imagine and hard to earn. C’est la vie !
So where does that leave us? Less of a checklist of patterns to force healthcare into, and more of a gradient of autonomy you earn through verifiability. The parts of healthcare that are structured — FHIR, data reconciliation — are the parts you can verify, and the parts you can verify are the parts an agent can own. The clinical-judgment end stays human, not because agents can’t reach it, but because we can’t yet measure them well enough to trust them there.
And now we loop full circle back to the question I parked earlier: when is it good enough ? Verifiability is what makes that question answerable at all: you can’t measure “good enough” in a domain you can’t check. Time to check out the world of evals!
Some notes on software, healthcare and AI
