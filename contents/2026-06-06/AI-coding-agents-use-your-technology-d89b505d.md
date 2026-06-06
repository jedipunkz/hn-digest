---
source: "https://developer.microsoft.com/blog/how-ai-coding-agents-actually-use-your-technology"
hn_url: "https://news.ycombinator.com/item?id=48422295"
title: "AI coding agents use your technology"
article_title: "How AI coding agents actually use your technology - Microsoft for Developers"
author: "waldekm"
captured_at: "2026-06-06T09:46:10Z"
capture_tool: "hn-digest"
hn_id: 48422295
score: 2
comments: 0
posted_at: "2026-06-06T07:17:49Z"
tags:
  - hacker-news
  - translated
---

# AI coding agents use your technology

- HN: [48422295](https://news.ycombinator.com/item?id=48422295)
- Source: [developer.microsoft.com](https://developer.microsoft.com/blog/how-ai-coding-agents-actually-use-your-technology)
- Score: 2
- Comments: 0
- Posted: 2026-06-06T07:17:49Z

## Translation

タイトル: AI コーディング エージェントがあなたのテクノロジーを使用します
記事のタイトル: AI コーディング エージェントが実際にテクノロジーを使用する方法 - Microsoft for Developers
説明: SDK、CLI、API を出荷し、開発者がそれを使用します。現在では、AI コーディング エージェントもこれを使用していますが、人間とは異なる使い方をしています。ほとんどの場合

記事本文:
メインコンテンツにスキップ
マイクロソフト
開発者ブログ
開発者ブログ
開発者ブログ
ホーム
開発者
開発者向けのマイクロソフト
Microsoft Entra ID 開発者
プラットフォームとツール
ビジュアルスタジオ
AI コーディング エージェントが実際にテクノロジーをどのように使用するか
AI コーディング エージェントが実際にテクノロジーをどのように使用するか
SDK、CLI、API を配布すると、開発者がそれを使用します。現在では、AI コーディング エージェントもこれを使用していますが、人間とは異なる使い方をしています。ほとんどの場合、「開発者がプロ​​ンプトを入力する」と「エージェントがテクノロジーを使用してコードを生成する」の間に実際に何が起こっているのかわかりません。エージェントはあなたの文書を読んでいますか? MCP サーバーを呼び出していますか?両方を無視して記憶から推測しているのでしょうか？
前回の記事 では、AX スタック (モデル、ハーネス、エージェント拡張機能) について紹介しました。何が修正され、何が影響を与えることができるかについて話し合いました。今回は、エージェントがテクノロジーに遭遇したときに実際に何が起こるかを段階的に追跡してみましょう。仕組みを理解するまでは、壊れているものを修正することはできないからです。
開発者が「何か作って」と言ったらどうなるか
開発者はコーディング エージェントを開き、「Contoso Identity を使用した認証を備えた REST API を構築してください」というプロンプトを入力します。次に何が起こるのかを説明します。
ステップ 1: ハーネスがコンテキストを組み立てる
モデルに何かがヒットする前に、ハーネス (コパイロット、クロード コード、カーソル) がコンテキスト ウィンドウを組み立てます。 VS Code チームは最近、コンテキストのアセンブリ、ツールの公開、エージェント ループをカバーする、ハーネスの仕組みに関する詳細な記事を公開しました。ハーネスは次のようにまとめられます。
システム プロンプト (ハーネス固有。表示または変更することはできません)
環境の詳細: 開発者の OS、作業ディレクトリへのフルパス
ハーネスが関連すると考えるワークスペース ファイル
インストールされている拡張機能 (MCP サーバー、スキル) からのツールの説明

s、カスタムエージェント)
任意の命令ファイル (.github/copilot-instructions.md、AGENTS.md など)
ハーネスはそれぞれ異なるため、これは単なる例です。それにもかかわらず、コーディングに使用される一般的な LLM のコンテキスト ウィンドウ サイズを考慮すると、そのようなセットアップがどのようにして利用可能なトークンをすぐに埋めてしまうのかがわかります。何がカットを行うかを決定するのはハーネスです。開発者が 20 個の拡張機能をインストールしている場合、ハーネスはツールの説明を要約したり、一部を完全に削除したり、推定された関連性によってそれらをランク付けしたりする可能性があります。拡張機能の説明は、モデルが認識する前にスペースをめぐって競合します。ハーネスの長さ制限 (各ハーネスが独自に設定) を超える場合は、どれほど関連性があるとしても完全に無視されます。また、OS やディレクトリ パスなど、普段考えたこともない詳細がモデルの決定に影響します。プラットフォーム固有のコードを生成し、さまざまなツールチェーンを想定し、ここで確認された内容に基づいてさまざまなデフォルト構成を選択することもできます。
ステップ 2: モデルが部屋を読み取ります
モデルはこの組み立てられたコンテキストを受け取り、人間がやらないことを行います。つまり、すべてを一度に読み取ります。システム プロンプト、ツールの説明、ワークスペース コンテキスト、開発者のプロンプト。利用可能なものと開発者が要求したもののメンタル モデルを構築します。
ここでトレーニング データが重要になります。モデルが事前トレーニング中にテクノロジーを認識している場合、モデルはすでに意見を持っています。 API は、API パターン、SDK 規約、一般的なエラー メッセージを知っています (または知っていると考えています)。あなたのテクノロジーを見たことがなければ、何も持たず、助けを求めるか、同様のテクノロジーに基づいて推測することになります。いずれにせよ、この時点でのモデルの仕事は、最初に何をすべきかを決定することです。コーディングを開始するのに十分な情報があるのか​​、それともツールを呼び出す必要があるのか​​。
それは変わります

この決定は、モデルにエンコードされた動作とハーネスが上に追加する命令の組み合わせであることがわかります。エージェントの中には、すぐにツールを呼び出す傾向がある人もいますが、自分の知識に頼って、ユーザーが要求した場合にのみツールを使用するエージェントもいます。エージェントの中には、まずインターネットで最新情報を検索する傾向がある人もいますが、効率を優先し、十分に知っていると感じたらタスクに取り組み始める人もいます。そのため、たとえ優れた拡張機能を出荷したとしても、あるエージェントはそれを積極的に呼び出す一方で、別のエージェントは開発者が明示的に要求しない限り、決してその拡張機能に触れない可能性があります。
ステップ 3: ツールの選択 (または選択しない)
モデルは、さらに情報が必要であると判断した場合、利用可能なツールとスキルを調べます。ここで、MCP サーバーのツールの説明とスキルの定義が重要になります。モデルは各説明を読んで、これが開発者の質問に役立つかどうかを判断します。この決定は、キーワード検索ではなく、意味上の一致に基づいていることに注意してください。モデルは意図を解釈しています。開発者が「認証」と述べ、ツールが「アイデンティティ プロバイダーの設定を構成する」と説明されている場合、モデルはそのギャップを埋める必要があります。
また、説明が意図と完全に一致している場合でも、モデルはそれをスキップする可能性があります。タスクが十分に単純そうに見える場合、またはモデルがすでに答えを知っていると確信している場合、モデルはわざわざツールを呼び出すことはありません。それはそれが持っているものと一緒に行きます。これは、モデルにテクノロジーのトレーニング データが含まれているものの、それが古い場合、つまり信頼性が高く、情報が古い場合に特に問題になります。
モデルがツールやスキルを選択しない場合、モデルはすでに知っているものをすべて使用して続行し、事前トレーニング データを使用します。トレーニングの打ち切りが何ヶ月前であっても、モデルが学習した SDK のどのバージョンも使用します。 AP の場合

それ以来変更しましたが、生成されたコードは間違っており、開発者もエージェントもそれを知りませんでした。
たとえば、モデルがあなたのツールを選択したとします。次に、それを正しく呼び出す必要があります。ツールのスキーマと開発者の意図に基づいてパラメータを構築します。しかし、モデルは開発者の意図を間違ったパラメータにマッピングしたり、スキーマと一致しない値を作成したりする可能性があります。
複数のツールが意図に一致する場合、モデルはいずれか 1 つまたは両方を呼び出す可能性があります。そして、それは関連性をめぐる競争だけではありません。より多くの情報を提供できる、異なるサーバーからの 2 つの MCP ツールがあるとします。 1 つは直接呼び出すことができ、もう 1 つはサルーティングを使用します。つまり、実際に答えを保持しているサブツールに関する情報を返す親ツールです。これは、LLM に公開されるツールの数を減らし、トークンの使用を最適化するための大規模な MCP サーバーで一般的な手法です。最初のターンで一方のツールが実際の情報を返し、もう一方のツールがサブツールのルーティング命令を返した場合、モデルは十分な知識を持っていると判断し、サブツールを呼び出さず、ルーティングが完全に失敗する可能性があります。このため、ツールを単独でテストするだけでは十分ではなく、開発者が実際に使用する組み合わせでテストする必要があります。
MCP サーバーは呼び出しを受信して​​処理し、コンテンツを返します。何を返すかは非常に重要です。返されるコンテンツが多すぎると、モデルがその一部を無視したり、混乱したりします。戻り値が少なすぎると、モデルは仮定でギャップを埋めてしまいます。コンテンツをわかりにくい形式でフォーマットすると、モデルはそれを誤って解析します。間違った情報を提供すると、エージェントが本来のタスクから完全に脱線してしまう可能性があります。拡張機能により、エージェントが開発者が要求したものとは異なるフレームワーク バージョンにプロジェクトをアップグレードしたり、タスクの途中で別のプログラミング言語に切り替えたりすることが確認されています。そしてそれはすべて行きます

コンテキスト ウィンドウに戻り、トークンを消費します。 200 個で済むドキュメントのトークンを 3,000 個返した場合、他の関連するコンテキストをウィンドウの外に押し出したことになります。それがドラッグです。
ステップ 5: モデルが応答を処理する
モデルはツールの応答を受け取り、それをコンテキスト内の他のすべてのものと統合します。これで、開発者の意図、ワークスペースのコンテキスト、ツールの出力が認識され、コードを生成するか、別のツールを呼び出すか、開発者に質問するかを決定します。ツールが明確で具体的なコンテンツ (コード サンプル、スキーマ、ステップバイステップの指示) を返した場合、モデルを続行できます。ツールが膨大な参照ドキュメントを返した場合、モデルは関連するものを抽出する必要があり、その抽出では重要な詳細が見逃される可能性があります。
どうすればわかりますか？そうは思わないでしょう。ツールが呼び出され、コンテンツが返され、外から見るとすべてが正常に見えます。しかし、モデルは間違った段落に引っかかり、内部専用のエンドポイントを使用するコードを生成してしまいました。これは最悪の種類の品質障害であり、誰かがコードを実行するまではわかりません。
次に、モデルはコードを生成します。ここで、上流のすべてがうまくいくか、破綻するかのどちらかになります。検出が機能した場合、モデルは拡張機能を見つけました。選択が機能した場合は、タスクに適切なツールが呼び出されています。ツールの応答が良好であれば、モデルには正確な最新の情報が含まれており、生成されたコードでは適切な SDK バージョン、適切なパターン、適切な認証フローが使用されています。
いずれかのステップが失敗した場合、モデルはトレーニング データに戻ります。古い API バージョンを使用したり、競合する SDK を選択したり、存在しないエンドポイントを作成したりする可能性があります。開発者は、動作しているように見えるコードが実行時に失敗するのを見て、エージェントのせいにします。
多くのエージェントは言語サーバー (LSP) と統合されており、実際に報告された問題に対応します。

最後まで待つのではなく、コードが生成されるときにコードを調整します。 VS Code では、エージェントは、アクティブな拡張機能がタイプ エラー、lint 違反、非推奨の警告などの診断を表面化できる [問題] パネルも監視します。あなたのテクノロジーが (VS Code 拡張機能、言語サーバー、またはリンターを通じて) そのフィードバックに貢献している場合、生成直後だけでなく、生成中にエージェントの出力に影響を与えることになります。それはあなたがコントロールするもう一つの表面です。
エージェントは最初の世代で停止しません。ハーネスがサポートしている場合 (そしてほとんどのハーネスがサポートしている場合)、エージェントはコードを構築し、テストを実行し、エラーを観察し、再試行します。ここでもテクノロジー面が重要です。 CLI が明確なエラー メッセージを生成した場合、エージェントは迅速かつ効率的に自己修正できます。ビルド ツールまたはテスト ランナーが特定のコードと提案を含む有用な出力を返した場合、エージェントは問題を修正します。エラーがわかりにくい場合 (「エラー: 操作が失敗しました」)、エージェントは盲目的に飛行しており、諦めるまで 10 ターンにわたって間違った方向に繰り返す可能性があります。
これは、AX の見落としがちな部分です。エラー メッセージはもはや人間の開発者だけのものではありません。これらはエージェント向けのものであり、エージェントには頼るべき直感がありません。彼らはあなたのエラーメッセージを文字通りに受け取ります。
コンテキストアセンブリ → モデルが何を認識できるかを決定します
モデルの解釈 → モデルが利用できると考えるものを決定します
ツールの選択 → モデルが拡張機能を使用するかどうかを決定します
ツールの呼び出し → モデルがどのような情報を取得するかを決定します
応答処理 → モデルがその情報をどのように使用するかを決定します
コード生成 → 開発者が実際に何を受け取るかを決定します
反復 → エージェントが自己修正できるかどうかを決定する
ステップ 1 での失敗はすべてに波及します。ハーネスがツールの説明を落とした場合

この場合、ステップ 3 ～ 7 は実行されません。あなたのツールは存在しないかもしれません。ステップ 3 が失敗した場合 (モデルはツールを認識しますが、タスクに接続しません)、同じ結果になります。ステップ 4 は成功したが、応答がわかりにくい場合は、ステップ 6 で壊れたコードが生成されます。
これが、「私のツールが呼ばれたかどうか」だけを測定する理由です。ほとんど何も教えてくれません。応答品質が悪い場合でも、ツールを正しく呼び出すことができ、ドラッグが発生する可能性があります。ツールはインストールされているにもかかわらず呼び出すことができません。これは、検出が壊れていることを意味します。そしてそれらの両極端の間でしょうか？各ステップが次のステップを静かに低下させる可能性があるのは、さまざまな方法があります。
AX の障害のほとんどは目に見えません。それらは上流で静かに発生し、目にすることはありません。出力だけを見て診断することはできません。それぞれに異なる修正があるため、どのステップが失敗したかを知る必要があります。次の記事では、AX の作業によって各ステップで揚力または抗力が生じているかどうかを測定する方法について説明します。
Waldek Mastykarz は主任クラウド擁護者です。
最初にディスカッションを開始してください。
AX スタック: 何が修正され、どこで勝てるのか
Microsoft Learn を基盤にしてエージェント開発者ツールを改善する
教育者のトレーニングと能力開発
学生と保護者向けの特典
AI マーケットプレイス アプリのサポート
プライバシーに関する選択
消費者の健康に関するプライバシー
サイトマップ

## Original Extract

You ship an SDK, a CLI, an API, and developers use it. Now AI coding agents use it too, except they use it differently than humans do. Most of the time

Skip to main content
Microsoft
Dev Blogs
Dev Blogs
Dev Blogs
Home
Developer
Microsoft for Developers
Microsoft Entra Identity Developer
Platform and Tools
Visual Studio
How AI coding agents actually use your technology
How AI coding agents actually use your technology
You ship an SDK, a CLI, an API, and developers use it. Now AI coding agents use it too, except they use it differently than humans do. Most of the time you have no idea what’s actually happening between “developer types a prompt” and “agent generates code with your technology.” Is the agent reading your docs? Is it calling your MCP server? Is it ignoring both and guessing from memory?
In the previous article , we introduced the AX stack: model, harness, and agent extensions. We talked about what’s fixed and what you can influence. This time, let’s trace through what actually happens, step by step, when an agent encounters your technology. Because until you see the mechanics, you can’t fix what’s breaking.
What happens when a developer says “build me something”
A developer opens their coding agent, types a prompt: “Build me a REST API with authentication using Contoso Identity.” Here’s what happens next.
Step 1: The harness assembles context
Before anything hits the model, the harness (Copilot, Claude Code, Cursor) assembles the context window. The VS Code team recently published a deep dive into how their harness works , covering context assembly, tool exposure, and the agent loop. The harness pulls together:
the system prompt (harness-specific, you can’t see or change it)
environment details: the developer’s OS, the full path to the working directory
workspace files the harness thinks are relevant
tool descriptions from installed extensions (MCP servers, skills, custom agents)
any instruction files (.github/copilot-instructions.md, AGENTS.md, etc.)
This is just an example, because every harness is different. Nonetheless, if you consider the context window size of any of the popular LLMs used for coding, you can start to see how such a setup quickly fills up the available tokens. The harness decides what makes the cut. If the developer has 20 extensions installed, the harness might summarize tool descriptions, drop some entirely, or rank them by estimated relevance. Your extension’s description is competing for space before the model even sees it. If it exceeds the harness’s length limit (each harness sets its own), it gets ignored entirely, no matter how relevant it is. And details you’d never think about, like the OS, or the directory path, influence the model’s decisions. It’ll generate platform-specific code, assume different toolchains, even pick different default configurations based on what it sees here.
Step 2: The model reads the room
The model receives this assembled context and does something humans don’t: it reads everything at once. The system prompt, the tool descriptions, the workspace context, the developer’s prompt. It builds a mental model of what’s available and what the developer asked for.
Here’s where training data matters. If the model has seen your technology during pre-training, it already has opinions. It knows (or thinks it knows) your API patterns, your SDK conventions, your common error messages. If it hasn’t seen your technology, it has nothing, and it’ll either ask for help or guess based on similar technologies. Either way, the model’s job at this point is to decide what to do first: does it have enough information to start coding, or does it need to call a tool?
It turns out, that this decision is a combination of the behavior encoded in the model and the instructions the harness adds on top. Some agents are more inclined to call tools straight away, while others rely on their own knowledge and only use tools when the user asks them to. Some agents tend to search for the latest information on the internet first, while others prioritize efficiency and start working on the task if they feel they know enough. So even if you ship a great extension, one agent might call it proactively while another never touches it unless the developer explicitly asks.
Step 3: Tool selection (or not)
If the model decides it needs more information, it looks at the available tools and skills. This is where your MCP server’s tool descriptions and skill definitions matter. The model reads each description and decides: does this help with what the developer asked? Notice, that this decision is based on semantic matching, not keyword search. The model is interpreting intent. If the developer said “authentication” and your tool is described as “configure identity provider settings,” the model has to bridge that gap.
And even when your description matches the intent perfectly, the model might still skip it. If the task looks simple enough, or if the model feels confident it already knows the answer, it won’t bother calling your tool. It’ll just go with what it has. This is especially painful when the model has some training data for your technology but it’s outdated: high confidence, stale information.
If the model doesn’t select any tool or skill, it proceeds with whatever it already knows: it’ll use pre-training data. It’ll use whatever version of your SDK the model learned from, however many months ago that training cut-off was. If your API changed since then, the generated code is wrong and neither the developer nor the agent knows it.
Say, the model selected your tool. Now it needs to call it correctly. It constructs the parameters based on the tool’s schema and the developer’s intent. But the model might map the developer’s intent to the wrong parameters, or invent values that don’t match the schema.
If multiple tools match the intent, the model might call either one or both. And it’s not just about competition for relevance. Say there are two MCP tools from different servers that can provide more information. One is invokable directly, the other uses subrouting: a parent tool that returns information about subtools that actually hold the answer. This is a common technique in larger MCP servers to lower the number of tools exposed to the LLM and optimize token usage. If in the first turn one tool returns the actual information while the other returns routing instructions for the subtool, the model might decide it knows enough and never invoke the subtool, failing the routing entirely. This is why testing tools in isolation isn’t enough, and you need to test them in the combinations developers actually use.
Your MCP server receives the call, processes it, and returns content. What you return matters a lot . Return too much content and the model ignores parts of it or gets confused. Return too little, and the model fills gaps with assumptions. Format the content in a confusing way, and the model misparses it. Provide wrong information and you can derail the agent from the original task entirely. We’ve seen extensions cause agents to upgrade a project to a different framework version than what the developer asked for, or switch to a different programming language mid-task. And all of it goes back into the context window, consuming tokens. If you return 3,000 tokens of documentation when 200 would do, you just pushed other relevant context out of the window. That’s drag.
Step 5: The model processes the response
The model takes your tool’s response and integrates it with everything else in context. It now has the developer’s intent, workspace context, and your tool’s output, and it decides whether to generate code, call another tool, or ask the developer a question. If your tool returned clear, specific content (a code sample, a schema, step-by-step instructions), the model can proceed. If your tool returned a wall of reference documentation, the model has to extract what’s relevant, and its extraction might miss the critical detail.
How would you know? You wouldn’t. The tool got called, it returned content, everything looks fine from the outside. But the model latched onto the wrong paragraph and generated code that uses an internal-only endpoint. That’s the worst kind of quality failure: invisible until someone runs the code.
Next, the model generates code. This is where everything upstream either pays off or falls apart. If discovery worked, the model found your extension. If selection worked, it called the right tool for the task. If the tool response was good, the model has accurate, current information, and the generated code uses the right SDK version, the right patterns, the right authentication flow.
If any step failed, the model falls back to its training data. It might use an older API version, pick a competing SDK, or invent an endpoint that doesn’t exist. The developer sees working-looking code that fails at runtime, and blames the agent.
Many agents integrate with language servers (LSPs) and respond to problems reported in real time, adjusting code as it’s generated rather than waiting until the end. In VS Code, the agent also monitors the Problems panel where any active extension can surface diagnostics: type errors, lint violations, deprecation warnings. If your technology contributes to that feedback (through a VS Code extension, a language server, or a linter), you’re influencing the agent’s output during generation, not just after. That’s another surface you control.
The agent doesn’t stop at first generation. If the harness supports it (and most do), the agent builds the code, runs tests, observes errors, and tries again. Your technology surface matters here too. If your CLI produces clear error messages, the agent can quickly and efficiently self-correct. If your build tooling or test runner returns helpful output with specific codes and suggestions, the agent fixes the problem. If your errors are cryptic (“Error: operation failed”), the agent is flying blind and might iterate in the wrong direction for 10 turns before giving up.
This is a part of AX that people overlook. Your error messages aren’t just for human developers anymore. They’re for agents, and agents have no intuition to fall back on. They take your error message literally.
Context assembly → determines what the model can see
Model interpretation → determines what the model thinks is available
Tool selection → determines whether the model uses your extension
Tool invocation → determines what information the model gets
Response processing → determines how the model uses that information
Code generation → determines what the developer actually receives
Iteration → determines whether the agent can self-correct
A failure at step 1 cascades through everything. If the harness drops your tool description, steps 3-7 never happen. Your tool might as well not exist. If step 3 fails (the model sees your tool but doesn’t connect it to the task), it’s the same result. If step 4 succeeds but the response is confusing, step 6 produces broken code.
This is why measuring only “was my tool called?” tells you almost nothing. Your tool can be called correctly and still produce drag if the response quality is poor. Your tool can never be called despite being installed, which means your discovery is broken . And between those two extremes? A dozen ways each step can silently degrade the next.
Most AX failures are invisible. They happen upstream, silently, and you never see them. You can’t diagnose them by looking at outputs alone: you need to know which step failed, because each one has a different fix. In the next article, we’ll cover how to measure whether your AX work is creating lift or drag at each step.
Waldek Mastykarz is a Principal Cloud Advocate.
Be the first to start the discussion.
The AX stack: what’s fixed, where you can win
Improve your agentic developer tools by grounding in Microsoft Learn
Educator training and development
Deals for students and parents
Support for AI marketplace apps
Your Privacy Choices
Consumer Health Privacy
Sitemap
