---
source: "https://alejandromp.com/development/blog/replacing-bash-with-swift-in-an-ai-harness/"
hn_url: "https://news.ycombinator.com/item?id=48396260"
title: "Replacing Bash with Swift in an AI Harness"
article_title: "Replacing Bash with Swift in an AI Harness | Alejandro M. P."
author: "ianhxu"
captured_at: "2026-06-04T10:22:11Z"
capture_tool: "hn-digest"
hn_id: 48396260
score: 2
comments: 0
posted_at: "2026-06-04T09:38:38Z"
tags:
  - hacker-news
  - translated
---

# Replacing Bash with Swift in an AI Harness

- HN: [48396260](https://news.ycombinator.com/item?id=48396260)
- Source: [alejandromp.com](https://alejandromp.com/development/blog/replacing-bash-with-swift-in-an-ai-harness/)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T09:38:38Z

## Translation

タイトル: AI ハーネスで Bash を Swift に置き換える
記事のタイトル: AI ハーネスで Bash を Swift に置き換える |アレハンドロ M. P.
説明: 数週間前、Cocoanetics の『An Interpreter for Swift』を読みましたが、頭の中にぼんやり浮かんでいたものを、他の誰かがすでに明確に表現しているという素晴らしい感覚を感じました。

記事本文:
AI ハーネスで Bash を Swift に置き換える |アレハンドロ M. P.
アレハンドロ M. P.
開発
スウィフト
開発
スウィフト
もっと見る
シリーズ: エージェント ハーネスを構築する
エージェント ハーネスはもう構築しましたか?
AI はプロジェクトを覚えていませんが、Markdown は覚えています
複数のツールも必要なのでしょうか?
macOS で AI ハーネスをサンドボックス化する
AI ハーネスにスキルを教える
AI ハーネスで Bash を Swift に置き換える
AI ハーネスで Bash を Swift に置き換える
数週間前、Cocoanetics の『 An Interpreter for Swift 』を読みましたが、頭の中に漠然と浮かんでいたものを、他の誰かがすでに明確に表現してくれているという素晴らしい感覚を感じました。
はっきり言っておきますが、誰かが TypeScript がエージェント コーディングの未来であると主張するたびに、私はとても悲しくなります。そこで私は、スウィフトもそのレースに参加してほしいという願いを宇宙に送り続けました。今、ついに、欠けていた部分、つまり大好きな言語の通訳を実現することができました。 - コアネティクス
これには Swift が理想的な言語だと常々思っていたので、その言葉は私の中に響きました。そして、これらのシステムがどのように機能するかを学ぶためにすでに小さな Swift AI ハーネスを構築しているのであれば、ここはアイデアをテストするのに最適な場所であると感じました。そこで、この Swift インタープリターをハーネスで使用してみる必要があるとすぐに思いました。たとえプロジェクトがまだ小さく、主に学習用だったとしても、それでも役立つ可能性があると思いました。
前回の投稿で、私はすでに小さなファイル ツールボックスを bash に置き換えていましたが、主な目的は、これらのファイル ツールをシェル コマンドのラッパー以外のものであるかのように振る舞うのをやめたらどうなるかを確認することでした。
したがって、実験の次の部分を実行して、汎用ツールが bash ではなく Swift だったらどうなるか、という答えを得るのにすでに良い位置にありました。
モデルにシェルを書かせる代わりに、必要なトップレベルの Swift を何でも書かせます。

任務を達成するために。それはすでに私にはずっと素敵に聞こえます。
この投稿の要点は、自分で構築するということであり、重要なアイデアはすべてすでにここにあります。ただし、私の執筆をサポートしたい場合、または単に時間を節約したい場合は、ダウンロードできるようにプロジェクトをパッケージ化しました。
「コーヒーを買って」で私をサポートしてハーネスを手に入れましょう
このアイデアの最も明白なバージョンは非常にシンプルです。ハーネスはそのままにしておきます。 bash を Swift ツールに置き換えるだけです。モデルにトップレベルの Swift を記述させ、それを実行するためにコンパイラーを呼び出します。
それは間違いなく機能しますが、実際の Swift を実行すると、ツールチェーン全体が付属します。モジュールキャッシュ。コンパイラの動作。 Swift を実行するにはコンパイル ステップを実行する必要があることは誰もが知っていますが、特に毎回新しいコンパイラ プロセスを開始する必要がある場合、これは世界最速のことではありません。また、LLM がタスクを達成するために何度も往復する可能性があることを考えると、ここで得られる利益にはそれだけの価値があります。これは、「これは小さな組み込みランタイムです」ではなく、「これは言語ツールチェーンです」という通常の現実です。
これは、Bash で発生したセキュリティの問題も抱えていることを意味します。もちろんシートベルトは引き続き使用できますが、この機会を利用して改善できるかもしれません。
そして最後に、Swift インタプリタを使いたかっただけです。公式な解決策ではないにもかかわらず、これは非常に有望で有用に見え、長年にわたってもっと真剣に受け止められるべきもののように思えます。私はスクリプト作成においても Swift が最高の言語であると今でも思っていますが、パフォーマンスはそこにはありません。したがって、私の頭の中には常に 2 つの選択肢がありました。ツールチェーンのパフォーマンスを向上させて Jai と同じくらい速くするか、制限を受け入れてより高速なインタープリターを作成するかです。したがって、コアチームからのものではないにもかかわらず、私たちが今ここに持っているものを使って、2番目のアイデアを証明しましょう。
そこが SwiftScript の興味深いところです

。
SwiftScript は、ライブラリとして埋め込むことができるツリーウォーキング Swift インタープリターです。スイフトはありません。 swift -e へのシェルアウトはありません。 Swift を解析して解釈するだけです。
これで、Swift ツールは別のサブプロセスではなくなりました。ハーネス自体の一部である場合もあります。したがって、bash のものによく似た新しい ToolDefinition を作成するだけです。
static func swift() -> Self {
ツール定義(
名前：「スウィフト」、
説明: "実行エンジンとして SwiftScript を使用し、ワークスペースの制限に ShellKit を使用して、トップレベルの Swift コードをハーネス プロセス内で直接実行します。",
引数: ["コード"]、
run: { 引数、workspaceRoot in
コード = 引数["コード"] ?? 「」
let result = await SwiftScriptToolRunner.executeInProcess(
コード: コード、
場所: workspaceRoot
)
return if result.exitCode == "ExitStatus(0)",
let stdout = result.stdout.nonEmpty {
標準出力
} それ以外の場合は {
実行結果JSON(
stdout: 結果.stdout、
標準エラー出力: result.stderr、
終了コード: result.exitCode
)
}
}
)
}
インタプリタはどこで実行されますか?
これが私が理解しなければならない次の質問でした。私にはインタプリタを書いた経験がありますが、純粋なコードを実行するだけの場合は簡単です。数学的計算を必要とする構文ツリーを解釈するのは簡単ですが、I/O や「環境」と対話する必要があるその他の機能について話し始めると、話はさらに難しくなります。
つまり、埋め込み可能なインタプリタがあり、Swift の実行方法を知っていますが、その周囲の環境はどうなっているのでしょうか?印刷物はどこへ行くのでしょうか？現在のディレクトリは何ですか?ファイル API は何を許可するかをどのようにして知るのでしょうか?
ここで、同じ作者による別のライブラリである ShellKit が登場します。
SwiftScript のインタープリターは、デフォルトで出力およびランタイム コンテキストに Shell.current を使用します。ShellKit を使用すると、組み込み者は withCurrent { ... } を使用して現在の非同期タスク用のカスタム シェルをインストールできます。

したがって、実際の実行パスは次のようになります。
シェル = シェル(
stdout: stdoutSink、
stderr: stderrSink、
環境: .synthetic(作業ディレクトリ: workspaceRoot.path)、
スクリプト名: "<swift-tool>",
サンドボックス: 試してみますか? makeSandbox(...)
)
let インタープリター = Interpreter()
let status = try await shell.withCurrent {
awaitを試してくださいinterpreter.evalScript(コード、ファイル名: "<swift-tool>")
}
withCurrent 呼び出しが重要です。スクリプトの実行中、合成シェルを Shell.current としてバインドします。
したがって、解釈された Swift が ShellKit ブリッジを介してルーティングする処理を実行すると、次のようになります。
これは、読んだ後に実装の魔法が解ける素晴らしい瞬間の 1 つでした。通訳は何も不思議なことをしているわけではありません。私たちが提供するランタイムコンテキスト内で実行されているだけです。
この変更により、単に bash を実行するよりも優れた効果が得られていると感じています。 Swift が抽象的な意味で「より安全」だからではなく、実行モデルがより狭く、より制御されているからです。
bash の場合、ツールの自然な形状は「ここにシェルがあります。頑張ってください。楽しんでください」というものです。埋め込み SwiftScript の場合、自然な形は「ここにホストするインタープリターがあり、ここに公開することを選択したランタイム コンテキスト」になります。それは非常に異なる出発点です。
もちろん、これは bash シェル仮想化ツールを使用しても実現できますが、それが今日私たちがここにいる理由ではありません ^^。
ワンツールの実験でモデルに汎用言語を与えることを想定している場合、Swift が有力な候補のように感じられます。
ということは、サンドボックスはもう必要ないということですか?
良い点は、通訳がただ虚空に浮かんでいるわけではないということです。これは私たちが制御する仮想化シェル内で実行され、ShellKit はそのシェルにサンドボックスも提供します。
したがって、Swift ツールのシェルを構築するときに、次のようなポリシーをアタッチできます。
砂を返す

ボックス(
ドキュメントディレクトリ: workspaceRoot、
ダウンロードディレクトリ: workspaceRoot、
libraryDirectory: workspaceRoot.appendingPathComponent(".swift-script-library", isDirectory: true),
一時ディレクトリ: 一時ディレクトリ、
ホームディレクトリ: workspaceRoot、
承認: { の URL
ガード url.isFileURL else {
Guard let host = url.host, host.isEmpty == false else {
throw Sandbox.Denial(url: url, 理由: "非ファイル URL には承認するホストがありません")
}
戻る
}
解決済み = url.standardizedFileURL.resolveSymlinksInPath()
let path = 解決済み.path
if パス == gitDirectory.path || path.hasPrefix(gitDirectory.path + "/") {
throw Sandbox.Denial(url: url, 理由: "Swift ツールは .git にアクセスできません")
}
if isAllowed(パス: パス、下: workspaceRoot.path)
|| isAllowed(パス: パス、下: globalSkillsRoot.path)
|| isAllowed(パス: パス、下:temporaryDirectory.path) {
戻る
}
throw Sandbox.Denial(url: url, 理由: "ファイル URL が Swift ツール サンドボックスの外にあります")
}
)
これはすでに bash よりもはるかに優れた状況です。
通訳者は私たちが形作る世界の見方を理解します。ワークスペース、ユーザーレベルのスキルディレクトリ、一時ディレクトリを読み取ることができ、ブリッジ API を介してネットワークリクエストを行うこともできます。 .git には触れられません。ブリッジされたファイル API を介してマシンの残りの部分に侵入することはできません。
このため、私の最初の直感は、SwiftScript と ShellKit ですでに十分かもしれないということでした。実行モデルが bash から組み込みインタープリターに移行すると、全体がより制御されているように感じられます。デフォルトでは、任意のシェル パイプラインはありません。コンパイラのサブプロセスはありません。 「ツールをさらに生成して何が起こるかを確認する」という自然な方法はありません。
しかし、それでもインタプリタが厳しいセキュリティ境界になるわけではありません。
重要な違いは、ShellKit.Sandbox がインプロセスであることです。

ss ポリシー層。これは非常に便利ですが、OS によって強制されるサンドボックスとは異なります。また、これは依然として同じプロセスと同じランタイムの一部であり、私たちが制御しているため、ハーネスとコードの品質もセキュリティ攻撃にさらされます。
私が採用したもう 1 つのアプローチは、物事をシンプルにするために破棄されましたが、同じ組み込みインタープリターとシェル サンドボックスを使用しますが、bash で行ったのと同じように別のプロセスで生成することでした。そうすれば、その別個のプロセスを同じシートベルト システムにラップし、OS によって強化された追加の保護層をさらに得ることができます。
Sandbox-exec -p '<シートベルト プロファイル>' swiftagentharness swift-script-helper 'print(40 + 2)'
これは、ブラウザーが分離のためにマルチプロセスになり始めたときのことを思い出させます。この学習ハーネスでは、ShellKit とインタープリタのおかげでシンプルに保つだけで十分に思えますが、アウトプロセス システムが実際にどのように機能するかを見るのは依然として興味深いものです。
では、これは実際にはどのように見えるのでしょうか?
その時点で、残る唯一の興味深い質問は実践的な質問でした。この Swift ツールを使用してハーネスを実行すると、実際に希望どおりに動作しますか?
少なくともこのシリーズで重要なことに関しては、答えは「はい」でした。
LLM は Swift ツールを呼び出して、Swift を実行するだけです。
<ステップ1>
ツール呼び出し: TOOL_CALL {"名前":"swift","引数":{"コード":"print(40 + 2)"}}
ツール成功: 迅速 -> 42
<ステップ2>
アシスタント: Swift ツールは機能しました。最上位の Swift から 42 を出力しました。
Foundation を使用し、ワークスペースを検査し、ファイルを読み取ることもできます。
<ステップ1>
ツール呼び出し: TOOL_CALL {"名前":"swift","arguments":{"code":"import Foundation ..."}}
ツール成功: swift -> cwd=/Users/me/code/swiftagentharness
トップレベルカウント=11
top_level_sample=.agents,.build,.git,.gitignore,AGENTS.md,Package.resolved
パッケージ行数=34
エージェントライン

=6
<ステップ2>
アシスタント: Swift ツールは機能し、要求された結果を返しました。
したがって、LLM の観点から見ると、これはすでにターミナル ラッパーよりも実際の汎用言語ツールのように動作します。
そして、スクリプトがエスケープしようとすると、サンドボックスが起動します。
<ステップ1>
ツール呼び出し: TOOL_CALL {"name":"swift","arguments":{"code":"import Foundation ... \"hello\" を試してください。write(toFile: \"../escape.txt\"、アトミック: true、エンコーディング: .utf8)"}}
tool-success: swift -> {"stdout":"","stderr":"ファイル URL が Swift ツール サンドボックスの外にあります\n","exit_code":"interpreter-error"}
それはまさに私が望んでいた種類の動作です。モデルはファイルの読み書きができますが、公開することに決めたワールド内でのみ可能です。
もちろん、これは依然として単一の作成者によって構築されたインタプリタ ライブラリであり、巨大な公式ランタイムではないため、表面にはまだ完全に完成していない部分があります。それでいいのです。それは今でもその要点を非常によく証明しています。重要なことは、現在考えられるすべての Swift API をサポートしているということではありません。重要なことは、Swift を使用すると、ワンツールの実験がすでに現実的に感じられるということです。
bash を組み込み Swift に置き換えるのは単なるギミックではありません。ハーネスの形状が真に変わり、ワンツールでの実験がより楽しくなります。

[切り捨てられた]

## Original Extract

A few weeks ago I read An Interpreter for Swift, from Cocoanetics, and I had that nice feeling of somebody else having already articulated the thing that was vaguely floating in your head.

Replacing Bash with Swift in an AI Harness | Alejandro M. P.
Alejandro M. P.
Development
Swift
Development
Swift
More
Series: Build an Agent Harness
Have You Built an Agent Harness Yet?
AI doesn't remember your project, Markdown does
Do We Even Need Multiple Tools?
Sandboxing an AI Harness on macOS
Teaching Skills to an AI Harness
Replacing Bash with Swift in an AI Harness
Replacing Bash with Swift in an AI Harness
A few weeks ago I read An Interpreter for Swift , from Cocoanetics, and I had that nice feeling of somebody else having already articulated the thing that was vaguely floating in your head.
Let me just say it outright: I get very sad every time somebody insists that TypeScript is the future of agentic coding. So I kept sending my wish to the universe that Swift should be also in the race for that. Now – finally – I was able to manifest the missing piece: an interpreter for the language I love. - Cocoanetics
That resounded inside me as I’ve always thought that Swift was the ideal language for this. And if I am already building a tiny Swift AI harness to learn how these systems work, then this felt like the perfect place to test the idea. So I immediately knew I had to try using this Swift interpreter in the harness. Even if the project is still tiny and mostly for learning, I thought it could still be useful.
In the previous post I had already replaced the little file toolbox with bash , mostly to see what happened if I stopped pretending those file tools were anything other than wrappers around shell commands.
So we were already in a good position to perform the next part of the experiment and answer, what if the generic tool was not bash , but Swift ?
Instead of letting the model write shell, let it write whatever top-level Swift it needs to accomplish the task. That sounds much nicer to me already.
The whole point of this post is that you build it yourself, and all the important ideas are already here. But if you want to support my writing, or you just want to save time, I packaged the project for you to download.
Get the harness by supporting me on Buy me a coffee
The most obvious version of the idea is very simple. Keep the harness the same. Just replace bash with a swift tool. Let the model write top-level Swift, and call out to the compiler to execute it.
That would definitely work, but running real Swift brings the whole toolchain with it. Module caches. Compiler behavior. We all know that running Swift means going through a compilation step, and that is not the fastest thing in the world, especially if we need to start a new compiler process every time. And given LLMs might roundtrip many times to accomplish a task, any gains we get here are worth it. It is the usual reality of “this is a language toolchain”, not “this is a tiny embedded runtime”.
It also means we have the security issues we had with Bash, we can still use Seatbelt of course, but maybe we can take this chance to do better.
And finally, I just wanted to use the Swift Interpreter. Despite not being an official solution, it looks very promising and useful, and like something that should have been taken more seriously over the years. I still think Swift is the best language even for scripting, but performance is not there. So in my head there have always been two options: increase the performance of the toolchain and be as fast as Jai, or accept the limits and make a faster interpreter. So what we have here now, despite not coming from the core team, lets us prove the second idea.
That is where SwiftScript becomes interesting.
SwiftScript is a tree-walking Swift interpreter that you can embed as a library. No swiftc . No shelling out to swift -e . Just parse Swift and interpret it.
Now the swift tool is not another subprocess. It can just be part of the harness itself. So we just make a new ToolDefinition that looks very similar to the bash one.
static func swift() -> Self {
ToolDefinition(
name: "swift",
description: "Run top-level Swift code directly inside the harness process, using SwiftScript as the execution engine and ShellKit for workspace confinement.",
arguments: ["code"],
run: { arguments, workspaceRoot in
let code = arguments["code"] ?? ""
let result = await SwiftScriptToolRunner.executeInProcess(
code: code,
at: workspaceRoot
)
return if result.exitCode == "ExitStatus(0)",
let stdout = result.stdout.nonEmpty {
stdout
} else {
executionResultJSON(
stdout: result.stdout,
stderr: result.stderr,
exitCode: result.exitCode
)
}
}
)
}
Where does the interpreter run?
This was the next question I had to understand. I have experience writing interpreters and things are easy when you just need to exercise pure code. Interpreting a syntax tree that wants a mathematical computation is trivial, but things get trickier when you start talking about I/O and other features that need to interact with an “environment”.
So we have an interpreter that we can embed and it knows how to run Swift, but what is the environment around it? Where does print go? What is the current directory? How do file APIs know what to allow?
This is where ShellKit , another library from the same author, enters the picture.
SwiftScript’s interpreter defaults to using Shell.current for output and runtime context, and ShellKit lets an embedder install a custom Shell for the current async task with withCurrent { ... } .
So the actual execution path ends up looking like this:
let shell = Shell(
stdout: stdoutSink,
stderr: stderrSink,
environment: .synthetic(workingDirectory: workspaceRoot.path),
scriptName: "<swift-tool>",
sandbox: try? makeSandbox(...)
)
let interpreter = Interpreter()
let status = try await shell.withCurrent {
try await interpreter.evalScript(code, fileName: "<swift-tool>")
}
That withCurrent call is the key. It binds our synthetic shell as Shell.current for the duration of the script.
So when interpreted Swift does things that route through the ShellKit bridge, it sees:
That was one of those nice moments where the implementation becomes less magical after you read it. The interpreter is not doing anything mysterious. It is just running inside a runtime context we provide.
This change is already feeling better than just running bash . Not because Swift is “safer” in some abstract sense, but because the execution model is narrower and more controlled.
With bash , the natural shape of the tool is “here is a shell, good luck, have fun”. With the embedded SwiftScript, the natural shape is “here is an interpreter we host, and here is the runtime context we chose to expose”. That is a very different starting point.
And of course this could be accomplished too by using some bash shell virtualizer, but that’s not why we are here today ^^.
If the one-tool experiment is supposed to give the model a general-purpose language, Swift feels like a great candidate.
Does that mean we no longer need sandboxing?
The good part is that the interpreter is not just floating there in the void. It runs inside the virtualized shell we control, and ShellKit gives that shell a sandbox too.
So when we build the shell for the swift tool, we can attach a policy like this:
return Sandbox(
documentsDirectory: workspaceRoot,
downloadsDirectory: workspaceRoot,
libraryDirectory: workspaceRoot.appendingPathComponent(".swift-script-library", isDirectory: true),
temporaryDirectory: temporaryDirectory,
homeDirectory: workspaceRoot,
authorize: { url in
guard url.isFileURL else {
guard let host = url.host, host.isEmpty == false else {
throw Sandbox.Denial(url: url, reason: "non-file URL has no host to authorize")
}
return
}
let resolved = url.standardizedFileURL.resolvingSymlinksInPath()
let path = resolved.path
if path == gitDirectory.path || path.hasPrefix(gitDirectory.path + "/") {
throw Sandbox.Denial(url: url, reason: "the Swift tool cannot access .git")
}
if isAllowed(path: path, under: workspaceRoot.path)
|| isAllowed(path: path, under: globalSkillsRoot.path)
|| isAllowed(path: path, under: temporaryDirectory.path) {
return
}
throw Sandbox.Denial(url: url, reason: "file URL is outside the Swift tool sandbox")
}
)
That is already a much nicer situation than bash .
The interpreter gets a view of the world that we shape. It can read the workspace, the user-level skills directory, a temp directory, and it can also make network requests through the bridged APIs. It cannot touch .git . It cannot just wander off into the rest of the machine through the bridged file APIs.
So due to this, my first instinct was that SwiftScript plus ShellKit might already be enough. Once the execution model moves from bash to an embedded interpreter, the whole thing already feels much more controlled. There is no arbitrary shell pipeline by default. No compiler subprocess. No “just spawn more tools and see what happens” as the natural path.
But that still does not make the interpreter a hard security boundary.
The important distinction is that ShellKit.Sandbox is an in-process policy layer. It is very useful, but is different from an OS-enforced sandbox. It is also still part of the same process and the same runtime we are controlling so it also exposes our harness and the quality of our code to security attacks.
Another approach I took, but discarded to keep things simple, was to use the same embedded interpreter and shell sandbox, but spawning it in a separate process just like we did with bash. If we do that, then we can wrap that separate process in the same Seatbelt system and get yet an extra, OS-enforced layer, of protection.
sandbox-exec -p '<seatbelt profile>' swiftagentharness swift-script-helper 'print(40 + 2)'
This reminds me of when browsers started going multi-process for isolation. Even though for this learning harness keeping it simple seems enough thanks to ShellKit and the interpreter, it’s still interesting to see how an out-of-process system would play in practice.
So what does this look like in practice?
At that point, the only interesting question left was the practical one. If I now run the harness with this swift tool, does it actually behave the way I want?
The answer was yes, at least for the kind of things that matter in this series.
The LLM can call the swift tool and just execute Swift:
<Step 1>
tool-call: TOOL_CALL {"name":"swift","arguments":{"code":"print(40 + 2)"}}
tool-success: swift -> 42
<Step 2>
Assistant: The Swift tool worked. It printed 42 from top-level Swift.
It can also use Foundation , inspect the workspace, and read files:
<Step 1>
tool-call: TOOL_CALL {"name":"swift","arguments":{"code":"import Foundation ..."}}
tool-success: swift -> cwd=/Users/me/code/swiftagentharness
top_level_count=11
top_level_sample=.agents,.build,.git,.gitignore,AGENTS.md,Package.resolved
package_lines=34
agents_lines=6
<Step 2>
Assistant: The Swift tool worked and returned the requested result.
So from the LLM point of view, this already behaves much more like a real general-purpose language tool than a terminal wrapper.
And the sandbox does kick in when the script tries to escape:
<Step 1>
tool-call: TOOL_CALL {"name":"swift","arguments":{"code":"import Foundation ... try \"hello\".write(toFile: \"../escape.txt\", atomically: true, encoding: .utf8)"}}
tool-success: swift -> {"stdout":"","stderr":"file URL is outside the Swift tool sandbox\n","exit_code":"interpreter-error"}
Which is exactly the kind of behavior I wanted. The model can read and write files, but only inside the world we decided to expose.
Of course, because this is still an interpreter library built by a single author and not some giant official runtime, there are parts of the surface that are not fully complete yet. That is fine. It still proves the point very well. The important thing is not that it supports every possible Swift API today. The important thing is that the one-tool experiment already feels real with Swift.
Replacing bash with embedded Swift is not just a gimmick. It genuinely changes the shape of the harness, and makes the one-tool experiment feel much more delibe

[truncated]
