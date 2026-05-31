---
source: "https://autolang.vercel.app/docs/philosophy-vision"
hn_url: "https://news.ycombinator.com/item?id=48336380"
title: "Show HN: A lightweight compiler for untrusted AI Agent scripts"
article_title: "Vision and Philosophy - Autolang | Autolang Docs"
author: "hoansdz"
captured_at: "2026-05-31T00:29:26Z"
capture_tool: "hn-digest"
hn_id: 48336380
score: 1
comments: 1
posted_at: "2026-05-30T14:09:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A lightweight compiler for untrusted AI Agent scripts

- HN: [48336380](https://news.ycombinator.com/item?id=48336380)
- Source: [autolang.vercel.app](https://autolang.vercel.app/docs/philosophy-vision)
- Score: 1
- Comments: 1
- Posted: 2026-05-30T14:09:27Z

## Translation

タイトル: Show HN: 信頼できない AI エージェント スクリプト用の軽量コンパイラ
記事のタイトル: ビジョンと哲学 - Autolang | Autolang ドキュメント
説明: Autolang のビジョンと中心となる哲学
HN テキスト: それを追求する価値があると思いますか?

記事本文:
ビジョンと哲学 - Autolang | Autolang ドキュメント Autolang
Autolang ドキュメント メニュー ドキュメントの概要
チャット デモを試す GitHub リポジトリ Autolang は、既存の機能を再利用しながら、AI が安全、迅速、低コストでコードを作成できるように設計されたスクリプト言語です。
AI エージェントの時代において、問題がますます明らかになってきています。それは、現在のツールの大部分が AI 向けではなく人間向けに設計されているということです。
これはパラドックスを生み出します。人間は、何でもできる柔軟で機能豊富な環境を必要とします。しかし、AI エージェントにはそれほどのパワーは必要ありません。 AI に Python、JavaScript、その他の汎用ランタイムなどのツールを直接使用させると、ファイルの読み取り、書き込み、削除、ネットワーク呼び出し、システムの操作、および意図しないアクションの実行など、あまりにも広範なエコシステムが AI に与えられることになります。この柔軟性は人間にとっては優れていますが、AI にとっては現実的なリスクとなります。
このリスクを軽減するために、開発者は通常、言語制限、権限のブロック、サンドボックス化、または仮想化による環境分離を適用します。ただし、これらの方法にはすべて共通の弱点があります。複雑で完全に制御するのが難しく、強力な分離にはシステム リソース、特に RAM、CPU、および運用オーバーヘッドの多大なコストがかかります。
Autolang はまさにこの問題を解決するために生まれました。
Autolang は、Python や JavaScript に代わるものではありません。これは、これらの言語の上または並行して配置されるオーケストレーション レイヤーであり、Python、JavaScript、C++、およびその他のエコシステムの既存の関数を、統合され、安全で、制御された実行環境にラップすることができます。
Autolang は、すでに機能している書き換えツールを必要としません。代わりに、npm、C++、またはその他のエコシステムからの既存の関数をラップし、Autolang バインディングを通じて AI に公開します。 AIコール

あなたが登録したものだけです。他には何もアクセスできません。
言い換えれば、Autolang は何でもできる言語になろうとしているわけではありません。 Autolang はマネージャーとして機能します。
2. AI エージェント システムの Autolang とは何ですか?
Autolang を使用すると、AI は制御範囲外の関数を自由に直接呼び出すことができなくなります。 AI が実行するすべてのアクションは、Autolang の静的コンパイラーとシステムによって確立されたルールを通過する必要があります。
これにより、非常に明確なモデルが作成されます。
AI は高レベルのロジックを記述します。
Autolang は、そのロジックを検証、制限、検証します。
実際の基礎となる実行関数は、Python、JavaScript、C++、または既存のライブラリから取得できます。
AI は Autolang によって定義された範囲内でのみ動作することが許可されます。
このアプローチにより、既存のソフトウェア エコシステムを再利用する機能を維持しながら、AI によって生成されたコードに関連するリスクが大幅に軽減されます。
ライブチャットボットエージェントのデモを試す
LLM エージェントが安全でリソースが制限されたサンドボックス内で Autolang コードを動的に作成して実行し、リクエストを解決する方法をご覧ください。
3. AI生成コードの特徴
Autolang は、AI によって記述されたコードの現実世界の特性に基づいて設計されました。 AI によって生成されたスニペットは通常、次の特徴を共有します。
それらは短く、通常は 100 行未満です。
極端で徹底的なパフォーマンスの最適化は必要ありません。
安全性の点で完全に信頼できるわけではありません。
無限ループ、不正なデータ型、null ポインター アクセス、不正な API 呼び出しなどの論理エラーが発生する傾向があります。
Autolang が解決に焦点を当てているのはまさにこの種類の問題です。 Autolang は、大規模で重いプログラムには最適化されていません。安全性、厳密な制御、総実行時間の短縮が求められる、頻繁に実行される小規模なコード スニペット向けに最適化されています。
これは、Autolang が行っている賭けも示しています。AI はモノラル全体を記述する必要がないということです。

ithic システムではありますが、構文が明確で、セマンティクスが厳密で、環境が安全で、外部関数を呼び出す機能が厳密に制御されている限り、ロジックの小さな断片を確実に記述することができます。これらはまさに、Autolang が提供するために構築された条件です。
4. パフォーマンスの考え方: 実行時間の合計時間
パフォーマンスについて議論するとき、Autolang はランタイムだけに焦点を当てているわけではありません。代わりに、Autolang は合計時間を優先します。
その理由は、AI スクリプトが、コンパイル時に大きなペナルティを必要とするほど高速に実行する必要がほとんどないためです。短いスクリプトを徹底的に最適化するためだけにコンパイルに 1 秒もかかる場合、全体的なユーザー エクスペリエンスは低下します。
したがって、Autolang は、コンパイル速度、コンパイラ インテリジェンス、および実行コストのバランスを取るように設計されています。目標は、コンパイルコストを追加する不必要な徹底的な最適化を行わずに、十分な最適化を行うことですが、Autolang はそのように設計されていません。
GitHub src/tests フォルダーの内部ベンチマークによると、Autolang は、第 12 世代 i5 プロセッサーと 16 GB RAM を搭載した Windows 11 上で約 900 行のコードを約 25 ミリ秒で実行します。単一のコンパイラー インスタンスの RAM 使用量は約 5.1MB (ピーク時約 7MB) です。コールドスタートはウォームシステムでは最大 18ms です。 Windows での初回実行時間は、OS レベルのファイル スキャンにより長くなる可能性があります。
5. 安全に対する考え方: 人間を信頼するように AI を信頼してはいけない
AI にコードを作成させる場合の最大の問題の 1 つは、AI が絶対的な安全性を保証できないことです。
コンテキスト外の関数呼び出し、
Autolang は、ユーザーが厳密な最大オペコード実行制限を設定できるようにすることで、この問題を軽減します。スクリプトがそのしきい値を超えると、ホスト プロセスをクラッシュさせることなく、スクリプトはただちに終了されます。
AI によって生成されたコードは計算集約的であることはほとんどないため、オペコードを制限することはエンドランを防ぐ非常に効果的な方法です。

実行中のコード、欠陥のあるループ ロジック、または無限サイクルに陥ったエージェント。
目的は、AI が常に正しいコードを書くと信じることではなく、たとえ AI が間違ったコードを書いたとしても、環境全体をダウンさせることができないシステムを設計することです。
6. データ型と信頼性の考え方
null 値が 1 つ誤ってアクセスされるだけで、プログラムがクラッシュするのに十分です。 AI によって生成されたコードでは、これは永続的なリスクとなります。
Autolang は、どのライブラリが lateinit 型および null 許容型の使用を許可されるかをライブラリごとに明示的に制御することで、この問題を解決します。これは次のことを意味します。
システム ライブラリは柔軟性を維持できます。
AI によって生成されたコードは、より厳しい制約に従う必要があります。
Null 許容値は、?? を使用して明示的に処理することが強制されます。オペレーター。
Null ポインターを介して VM をクラッシュすることは非常に困難になります。
また、Autolang は静的分析を利用して、AI に最初からタイプセーフなコードを強制的に作成させます。これは、JavaScript や Python のような高度に動的な言語とは大きく異なります。
7. 構文の考え方: 簡潔だが明確である
Autolang は、TypeScript の精神を引き継いだデータ構造を備えた Kotlin からインスピレーションを得た構文を使用します。理由は簡単です。Autolang は、完全に異質な構文の学習を AI モデルに強制するのではなく、AI モデルがこれらの人気のある言語からすでに持っている事前知識を活用したいと考えています。
ただし、すべての構文の決定は、使い慣れた言語を模倣するだけでなく、AI を支援するという目標を達成する必要があります。リストを例に挙げます。
宣言中 (例: val list = <Int>[] )、AI は型を明示的に宣言することを強制され、後でデータ型の不一致を防ぎます。
型が明確に推測できたら、[] だけで十分です。トークンが節約され、不必要な複雑さが軽減されます。
明確さが必要な場合は明示的にします。
十分な文脈がある場合は簡潔にします。
8. 構造化された制御ではなく、

無限の自由
JavaScript や Python などの言語は、非常に動的で柔軟性があります。人間にとって、これは利点です。 AI にとって、それは多くの場合責任となります。
Autolang は逆のアプローチを採用しています。自由度は低くなりますが、安全性は高くなります。 AI にすべてを自由に推測させる代わりに、Autolang は静的分析に依存して、データ型を強制し、アクセス許可を制限し、関数呼び出しの範囲を制御し、コードが実行前に厳密に有効であることを確認します。
AI システムでは、あいまいさは大きなコストとなります。 Autolang はそのコストを削減するために作成されました。
9. Autolang を使用する場合 (および使用しない場合)
Autolang は、すべてのシナリオで必須の選択ではありません。
LLM が実行時にコードを生成して実行するシステムを構築しており、LLM がアクセスできるものを制御する必要があります。
多数の AI エージェントを同時に実行しているため、ランタイムを完全に分離することによる RAM オーバーヘッドがコストの問題になりつつあります。
既存の JS または C++ 関数を書き換えることなく、制御された方法で AI に公開したいと考えています。
AI によって生成されたスクリプトは短く (約 100 行未満)、頻繁に実行されます。
大まかなしきい値: 5 つ以上のエージェントを同時に実行していて、メモリまたはコールドスタートの遅延が重要になる場合、トレードオフは意味を持ち始めます。
次の場合、Autolang はおそらく適切ではありません。
エージェントの数が少ないため、実行時の分離オーバーヘッドは問題になりません。
AI は長く複雑なプログラムを作成する必要があります。Autolang は大規模なアプリケーションではなく、短いスクリプト向けに最適化されています。
OS レベルのセキュリティ保証が必要です。Autolang は言語レベルのサンドボックスであり、プロセスの分離に代わるものではありません。
Python バインディングが必要ですが、まだ利用できません。
Wasm は、最新のサンドボックスにとって非常に効果的なソリューションです。ただし、Wasm を適切に使用するには、ビルド プロセス、関数の呼び出し、データの受け渡し、実行時エラーの処理を管理する必要があります。

NG、および環境オーケストレーション。
言い換えれば、Wasm は問題の一部を解決しますが、AI の調整層として自動的に機能するわけではありません。
Autolang はその役割を果たすことを目的としています。複数の構成ファイル、JSON、個別のランタイムの管理をユーザーに強いる代わりに、Autolang を使用すると、AI が Autolang 構文だけでその意図を表現できるようになり、調整と制御はシステムによって処理されます。
11. 運用哲学: AI が「何を」指示し、システムが「どのように」処理するか
AI を複数の異なるランタイム間で直接動作させる場合、環境、JSON ファイル、アクセス許可、および無数の安全メカニズムを管理するという負担がかかります。
Autolang はこのアーキテクチャを反転します。 AI があらゆる環境に直接触れるのではなく、すべてを事前定義された関数にラップします。 AI は、Autolang を通じてこれらの関数を呼び出すだけで済みます。システム全体は、複数のランタイムをやりくりする混乱ではなく、AI の意図を中心に展開します。
人間は、ネイティブ関数をラップし、インターフェイスを定義し、実行制限を設定し、安全に使用できるように AI に公開するだけで済みます。 Autolang は人間のワークフローを置き換えるために生まれた言語ではありません。人間と AI が連携するために最適化された中間層です。
Autolang は、何でもできるようにするという意味で AI を「強く」するために作られたわけではありません。これは、AI をより安全、安価、制御しやすくし、基本的に短期間のタスクの実行に適したものにするために作成されました。
Autolang は、あらゆる面で Python や JavaScript と競合するために作成されたわけではありません。これは、非常に具体的だが非常に重要な問題を解決するために生まれました。
より軽く、より安全で、消費リソースが少なく、制御が容易な方法で AI がソフトウェア ツールと対話できるようにする方法。
Autolang の中心となる哲学は次のとおりです。
既存のエコシステムを置き換えないでください

幹、
すべてを最初から書き直すことを強制しないでください。
ランタイムを不必要に肥大化させないでください。
偽りの柔軟性と引き換えに安全性を犠牲にすることは決してありません。
Autolang は、AI と実行環境の間の管理レイヤーです。
基礎となるシステムが実際の作業を実行します。
最新の AI エージェントと組み込みシステム向けに設計された、超低レイテンシーの厳密に型指定されたスクリプト言語です。ゼロトラスト メンブレン アーキテクチャで構築されています。
後援: ドキュメント紹介
© 2026 Autolang · MIT ライセンス

## Original Extract

The vision and core philosophy of Autolang

Do you think it's worth pursuing ?

Vision and Philosophy - Autolang | Autolang Docs Autolang
Autolang Docs Menu Docs Introduction
Try Chat Demo GitHub Repository Autolang is a scripting language designed for AI to write code safely, quickly, and at low cost while reusing existing functions.
In the era of AI Agents, an issue has become increasingly apparent: the vast majority of current tools are designed for humans, not for AI.
This creates a paradox. Humans require flexible, feature-rich environments where they can do anything. But AI Agents do not need all that power. When we let AI directly use tools like Python, JavaScript, or other general-purpose runtimes, we hand it an ecosystem that is far too broad: the ability to read, write, and delete files, make network calls, manipulate the system, and perform unintended actions. That flexibility is excellent for humans, but it represents a real risk for AI.
To mitigate this risk, developers typically apply language restrictions, permission blocking, sandboxing, or environment isolation via virtualization. However, these methods all share common weaknesses: they are complex, hard to fully control, and strong isolation comes at the heavy cost of system resources — specifically RAM, CPU, and operational overhead.
Autolang was born to solve this exact problem.
Autolang is not meant to replace Python or JavaScript. It is an orchestration layer that sits on top of or alongside those languages, allowing you to wrap existing functions from Python, JavaScript, C++, and other ecosystems into a unified, safe, and controlled execution environment.
Autolang does not require rewriting tools that already work. Instead, you wrap existing functions — from npm, C++, or any other ecosystem — and expose them to the AI through Autolang bindings. The AI calls only what you register. Nothing else is accessible.
In other words, Autolang doesn't try to be a language that can do everything. Autolang acts as a manager.
2. What is Autolang in an AI Agent System?
With Autolang, AI is no longer free to directly invoke functions outside its controlled scope. Every action the AI takes must pass through Autolang's static compiler and the rules established by the system.
This creates a very clear model:
The AI writes high-level logic.
Autolang verifies, restricts, and validates that logic.
The actual underlying execution functions can come from Python, JavaScript, C++, or existing libraries.
The AI is only permitted to operate within the scope defined by Autolang.
This approach significantly reduces the risks associated with AI-generated code, while preserving the ability to reuse existing software ecosystems.
Try the Live Chatbot Agent Demo
See how an LLM agent writes and runs Autolang code dynamically in a secure, resource-limited sandbox to solve your requests.
3. Characteristics of AI-Generated Code
Autolang was designed based on the real-world characteristics of code written by AI. Snippets generated by AI typically share the following traits:
They are short, usually under 100 lines.
They do not require extreme, deep performance optimization.
They cannot be fully trusted in terms of safety.
They are prone to logical errors such as infinite loops, incorrect data types, null pointer accesses, or wrong API calls.
That is the exact class of problems Autolang focuses on solving. Autolang is not optimized for massive, heavy programs. It is optimized for small, frequently-run code snippets that demand safety, strict control, and low total execution time.
This also informs the bet Autolang is making: AI doesn't need to write entire monolithic systems, but it can write small fragments of logic reliably — provided the syntax is clear, the semantics are strict, the environment is safe, and the ability to call external functions is tightly controlled. These are exactly the conditions Autolang is built to provide.
4. Performance Mindset: Total Time over Runtime
When discussing performance, Autolang doesn't solely focus on runtime. Instead, Autolang prioritizes total time:
The reason is that AI scripts rarely need to run so fast that it justifies a heavy compile-time penalty. If compiling takes a whole second just to deeply optimize a short script, the overall user experience degrades.
Therefore, Autolang is designed to strike a balance between compilation speed, compiler intelligence, and execution cost. The goal is to optimize just enough — without unnecessary deep optimizations that add compile cost but benefit long-running programs Autolang is not designed for.
According to internal benchmarks in the GitHub src/tests folder, Autolang executes roughly 900 lines of code in ~25ms on Windows 11 with a 12th-gen i5 processor and 16GB RAM. RAM usage for a single compiler instance is approximately 5.1MB (peak ~7MB). Cold start is ~18ms on a warm system; first-run times on Windows may be higher due to OS-level file scanning.
5. Safety Mindset: Don't Trust AI Like You Trust Humans
One of the biggest issues with letting AI write code is that AI cannot guarantee absolute safety.
out-of-context function calls,
Autolang mitigates this by allowing users to set a strict maximum opcode execution limit. If a script exceeds that threshold, it is immediately terminated — without crashing the host process.
Because AI-generated code is rarely compute-intensive, limiting opcodes is a highly effective way to prevent endlessly running code, faulty loop logic, or an agent getting stuck in an infinite cycle.
The objective isn't to trust that the AI will always write correct code, but to engineer a system where even if the AI writes bad code, it cannot bring down the entire environment.
6. Data Types and Reliability Mindset
A single misaccessed null value is enough to crash a program. With AI-generated code, this is a persistent risk.
Autolang solves this by providing explicit per-library control over which libraries are permitted to use lateinit and nullable types. This means:
System libraries can maintain flexibility.
AI-generated code must adhere to stricter constraints.
Nullable values are forced to be handled explicitly with the ?? operator.
Crashing the VM via null pointers becomes exceedingly difficult.
Autolang also utilizes static analysis to force the AI to write type-safe code from the start — a significant departure from highly dynamic languages like JavaScript or Python.
7. Syntax Mindset: Concise but Unambiguous
Autolang uses Kotlin-inspired syntax, with data structures that carry the spirit of TypeScript. The reasoning is straightforward: Autolang wants to leverage the prior knowledge AI models already have from these popular languages, rather than forcing them to learn an entirely alien syntax.
However, every syntax decision must serve the goal of assisting the AI — not just mimicking familiar languages. Take lists, for example:
During declaration (e.g., val list = <Int>[] ), the AI is forced to explicitly declare the type, preventing mismatched data types later on.
Once the type can be clearly inferred, [] is sufficient — saving tokens and reducing unnecessary complexity.
Explicit when clarity is needed.
Concise when there is enough context.
8. Structured Control, Not Infinite Freedom
Languages like JavaScript or Python are highly dynamic and flexible. For humans, this is an advantage. For AI, it is often a liability.
Autolang takes the opposite approach: less freedom, but more safety. Instead of letting the AI freely infer everything, Autolang relies on static analysis to enforce data types, limit access permissions, control the scope of function calls, and ensure the code is strictly valid before it ever runs.
In an AI system, ambiguity is a heavy cost. Autolang was created to eliminate that cost.
9. When to Use (and When Not to Use) Autolang
Autolang is not a mandatory choice for every scenario.
You are building a system where an LLM generates and executes code at runtime, and you need to control what it can access.
You are running many concurrent AI agents and the RAM overhead of full runtime isolation is becoming a cost problem.
You want to expose existing JS or C++ functions to AI in a controlled way, without rewriting them.
Your AI-generated scripts are short (under ~100 lines) and run frequently.
A rough threshold: if you are running 5+ concurrent agents and memory or cold-start latency matters, the tradeoff starts making sense.
Autolang is probably not the right fit if:
You only have a small number of agents and runtime isolation overhead is not a concern.
Your AI needs to write long, complex programs — Autolang is optimized for short scripts, not large applications.
You need OS-level security guarantees — Autolang is a language-level sandbox and does not replace process isolation.
You need Python bindings — they are not available yet.
Wasm is a highly effective solution for modern sandboxing. But to use Wasm properly, you still have to manage the build process, function invocations, data passing, runtime error handling, and environment orchestration.
In other words, Wasm solves part of the problem, but it does not automatically serve as a coordination layer for AI.
Autolang aims to fill that role: instead of forcing you to manage multiple config files, JSONs, and distinct runtimes, Autolang allows the AI to express its intent solely through Autolang syntax, while coordination and control are handled by the system.
11. Operational Philosophy: AI Dictates 'What', the System Handles 'How'
If you let the AI directly operate across multiple different runtimes, you are burdened with managing environments, JSON files, access permissions, and a myriad of safety mechanisms.
Autolang flips this architecture. Instead of letting the AI directly touch every environment, you wrap everything into pre-defined functions. The AI only needs to call those functions through Autolang. The entire system revolves around the AI's intent, rather than the chaos of juggling multiple runtimes.
Humans only need to: wrap native functions, define interfaces, establish execution limits, and expose them to the AI for safe usage. Autolang is not a language born to replace human workflows — it is an intermediary layer optimized for humans and AI working together.
Autolang was not created to make AI "stronger" in the sense of being able to do anything. It was created to make AI safer, cheaper, easier to control, and fundamentally better suited for short-lived task execution.
Autolang was not created to compete with Python or JavaScript on every front. It was born to solve a very specific, but highly critical problem:
how to allow AI to interact with software tools in a way that is lighter, safer, consumes fewer resources, and is easier to control.
The core philosophy of Autolang is:
do not replace the existing ecosystem,
do not force rewriting everything from scratch,
do not let runtimes bloat unnecessarily,
and never sacrifice safety in exchange for fake flexibility.
Autolang is the management layer between the AI and the execution environment.
The underlying system executes the real work.
An ultra-low latency, strictly-typed scripting language designed for modern AI Agents and Embedded Systems. Built with a zero-trust membrane architecture.
Sponsored by Documentation Introduction
© 2026 Autolang · MIT Licensed
