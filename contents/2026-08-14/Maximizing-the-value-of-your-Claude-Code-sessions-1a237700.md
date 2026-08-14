---
source: "https://claude.com/blog/maximizing-the-value-of-your-claude-code-sessions"
hn_url: "https://news.ycombinator.com/item?id=49300800"
title: "Maximizing the value of your Claude Code sessions"
article_title: "Maximizing the value of your Claude Code sessions | Claude by Anthropic"
author: "twapi"
captured_at: "2026-08-14T16:40:56Z"
capture_tool: "hn-digest"
hn_id: 49300800
score: 3
comments: 0
posted_at: "2026-08-14T16:15:21Z"
tags:
  - hacker-news
  - translated
---

# Maximizing the value of your Claude Code sessions

- HN: [49300800](https://news.ycombinator.com/item?id=49300800)
- Source: [claude.com](https://claude.com/blog/maximizing-the-value-of-your-claude-code-sessions)
- Score: 3
- Comments: 0
- Posted: 2026-08-14T16:15:21Z

## Translation

タイトル: クロード コード セッションの価値を最大化する
記事のタイトル: クロード コード セッションの価値を最大化する |クロード by Anthropic
説明: すべてのトークンから最大限の価値を引き出す効率的なセッションを実行する方法に関する実践的なヒント。

記事本文:
クロードコードセッションの価値を最大化する |クロード by Anthropic
クロード製品のご紹介 クロード
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者に連絡する 営業担当者に連絡する 営業担当者に問い合わせる
クロードを試してみる クロードを試してみる クロードを試してみる
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
クロードコードセッションの価値を最大化する
クロードコードセッションの価値を最大化する
すべてのトークンから最大の価値を得る効率的なセッションを実行する方法。
カテゴリ クロード コード エンタープライズ AI
共有 リンクをコピー https://claude.com/blog/maximizing-the-value-of-your-claude-code-sessions
タスク間で /clear を実行します。これにより、以前の無関係なコンテキストがモデルに送り返されることがなくなり、トークンの使用量が削減されます。
開始する前に、モデルと作業レベルを設定します。会話中にどちらかを変更すると、プロンプト キャッシュが破壊され、トークン コストが増加する可能性があります。
ファイルに名前を付ける代わりに @ メンションを付けます。ファイルはメッセージに直接添付されるため、Read 呼び出しや、クロードがファイルを探しに行かなければならない場合は検索の手間が省けます。
ノイズの多いコマンドに Quiet フラグを追加するか、サブエージェントで実行します。コマンド出力はファイルと同様に会話に追加され、セッションの残りの間そこに残ります。
新しいセッションで /context を 1 回実行します。ロードされているもの ( CLAUDE.md 、MCP ツール定義) が表示されるので、不要なものを削除できます。
キーボードから休憩する前に /compact を実行してください。プロンプトのキャッシュは 1 時間で期限切れになり、会話を要約するのは非常に手間がかかります

キャッシュされている間は安くなります。
つい最近まで、コードを作成するツールは定額 (または無料) でした。その日の午後に 1 つのテストを修正しても、50 のテストを修正しても、エディターのコストは同じなので、個々のタスクにはそれ自体の価格がありませんでした。
Claude Code のようなエージェント コーディング ツールを使用すると、それが可能になります。同じ完了したタスクでも、使用方法に応じて費用が異なる場合があります。
1 回のセッションで、クロードはテストとその対象となるファイルを読み取り、編集を行い、数回のターンで完了します。別の例では、最初にリポジトリ内を grep し、同じ 2 つのファイルに到達する途中で 12 個のファイルを読み取り、それらのターンごとに、今朝以降に読み取られた他のすべての内容も会話に引きずり込みます。
同じ修正ですが、費やしたトークンの数が異なり、モデルはその間ずっと、必要のない 10 個のファイルについても考慮する必要がありました。
トークンを効率的に使用するということは、全体として使用するトークンの数が少なくなるということではありません。それは、実際に使用しているものが、実際に要求したものに対応していることを確認することを意味します。
それでは、何がトークンの価格を決定するのか、次にセッションが送信するトークンの数を決定するものは何なのか、そしてその過程でそれがセッションの実行方法に何を意味するのかを見てみましょう。
トークンの価格を決めるもの
トークンごとに料金が請求されますが、実際に支払っているのは推論です。つまり、GPU (または TPU、またはモデルがたまたま実行されているもの) がトークン上でモデルを実行するのにかかる時間です。
トークンにかかる時間を決定するのは 3 つの要素です。実行しているモデル、入力トークン (入ってくる) か出力トークン (出てくる) か、キャッシュされたかどうかです。
モデルが大きくなると、入力トークンと出力トークンの両方でより多くの作業が行われます。どのモデルがどの種類の作業に価値があるかは、それ自体がトピックであり、これについては「クロード モデルと労力レベルの選択」で説明しました。

n クロード・コード。
この投稿で知っておく必要があるのは、これから説明する他のすべてのことにはモデルの価格が乗算されるということだけです。問題が本当に難しいか曖昧な場合はより大きなモデルを使用し、作業が日常的な場合はより小さなモデルを使用します。
前へ 前へ 0 / 5 次へ 次へ Claude Code Desktop を入手
IRM https://claude.ai/install.ps1 | iex コマンドをクリップボードにコピーする またはドキュメントを読む クロード コードを試す クロード コードを試す クロード コードを試す 開発者ドキュメント 開発者ドキュメント 開発者ドキュメント 電子書籍
リクエストは 2 つのフェーズで GPU を通過し、それぞれに異なるコストがかかります。
まず、プリフィル中に、モデルはリクエストとコンテキスト、つまりシステム プロンプト、 CLAUDE.md 、メッセージ、およびそれ以降に会話に追加されたすべてのもの (クロードが読み取ったファイルと実行したコマンドの出力) を読み取ります。これらが入力トークンです。
次に、デコード中に、出力トークン (その思考、ツール呼び出し、および表示されるテキスト) を書き込みます。これは一度に 1 つのトークンごとに行われます。 200 トークンの応答は、モデルを 200 回ずつ実行することを意味します。トークンごとに、デコードにより GPU が長時間ビジー状態に保たれるため、出力の価格は入力の約 5 倍となります。
セッション内の出力トークンの多くは思考トークンであり、モデルがターンごとにどれだけの思考を行うかは、努力レベルによって制御されます。モデルと同様に、/effort で選択したレベルは、次のセッションでもデフォルトとして残ります。
ヒント: 新しいセッションで /model と /effort を 1 回実行して、実際の状態を確認します。どちらもあなたが最後に選んだものを覚えており、その決定が慎重であることを望んでいます。
ヒント: セッションが単調な作業になることがすでにわかっている場合、MAX_THINKING_TOKENS=0 クロードは、その 1 つのセッションについて思考をオフにします (Fable 5 を除く)。これは /effort low より下のステップです。
リクエストがちょうど sa で始まる場合

me トークンをサーバーが見たばかりのリクエストとして取得すると、その共有開始の状態は同じになるため、サーバーはそれを前回から維持し、その後に来るもののみを事前入力できます。これはプロンプト キャッシュと呼ばれます。
サーバーは状態を計算するのではなくロードするため、キャッシュからの読み取りには入力価格の 0.1 倍のコストがかかります。トークンをキャッシュに書き込むには、サーバーがその後も状態を保持する必要があるため、通常の入力よりも若干コストがかかり、最大 2 倍になります。ただし、書き込みはトークンごとに 1 回発生し、その後は 0.1x 読み取りがすべてのターンで発生します。
Claude Code はリクエストごとにプロンプ​​ト キャッシュを管理します。オンにするものは何もありません。ただし、壊れる可能性があるため、このようなコストの高騰を回避する方法を知ることが重要です。
「utils.test.ts の失敗したテストを修正する」と入力するとします。 Claude Code が送信するものは次のとおりです。
Claude Code は、システム プロンプト (ツール定義を含む)、 CLAUDE.md 、およびメッセージから最初のリクエストを組み立て、それを送信します (トークンの入力)。キャッシュにはまだ何も入っていないため、すべてが事前に入力されてキャッシュに書き込まれます。
モデルはまだ見ていないテストを修正できないため、少し考えて、utils.test.ts (出力トークン) の Read 呼び出しで応答します。 Claude Code はファイルを読み取り、会話に追加し、全体を再度送信します (入力トークン)。今回は、リクエスト 1 からのすべてが 10 分の 1 の価格でキャッシュから読み戻されます。フル価格で事前に入力されているのは、Read 呼び出しとファイルだけです。
ここで、モデルはテスト対象のファイル (出力) を必要とします。別の読み取り、別の追加、そしてすべてが再び実行されます。キャッシュからのリクエスト 1 と 2、2 番目のファイルはフルプライス (入力) です。
モデルは編集 (出力) で応答します。クロード コードはそれを適用し、結果を追加して、すべてを再度送信します。同じ話: 編集

とその結果は新しく、その前にあるものはすべてキャッシュ読み取り (入力) です。
モデルは npm test (出力) を実行します。クロード コードはテスト出力を追加し、テスト出力を唯一の新しい部分 (入力) として、すべてを再送信します。
テストに合格すると、モデルは短い概要 (出力) で応答します。ツール呼び出しがないということは追加するものが何もなく、リクエスト 6 もないことを意味するため、これで完了です。
これは 1 つの小さな修正に対して 5 つのリクエストであり、それぞれのリクエストにはその時点までの会話全体が含まれていました。典型的なターンは偏っています。数万枚のトークンが入ってきて、数百枚が出てきます。ただし、そのターンで新しいもののみが正規価格で事前に入力されます。
これがターンごとの請求額全体です。履歴のキャッシュ読み取り、新しいものすべての入力価格、および応答の出力価格です。
これはサブスクリプションにも当てはまります。これらの価格は直接表示されませんが、同じリクエストによって制限が引き下げられます。
キャッシュはリクエストの最初から一致する必要があり、リクエストは常に同じ順序で送信されます。ツール定義、システム プロンプト、会話 (先頭に CLAUDE.md) が続きます。
そのプレフィックスの何かが変更されると、その後ろにあるものはすべて再度事前に入力されます。会話の最後にツールの結果が追加されるのは、背後に何もないため、理想的なケースです。キャッシュを破棄するものは、リクエストをさらに前方に変更するもの、またはキャッシュのキーとなる内容を変更するものです。
/model : すべてのモデルには独自のキャッシュがあるため、次のターンでは会話全体が定価で再度事前入力されます。 (これには、プラン モードに入る、またはプラン モードから抜けるたびにモデルを切り替える opusplan が含まれます。)
/effort: エフォート レベルもキャッシュのキー設定の一部であるため、同じ話になります。 /model と /effort の両方で、m に切り替えるときに確認を求めるのはこのためです。

会話の無駄。
高速モード : これもキーの一部であり、再プレフィルは高速モードの価格で行われるため、オンにする場合は最初にオンにしてください。 (キャッシュに関しては、再度オフにするのは無料です。)
/compact : 会話は短いものに置き換えられるため、会話内の何も一致しなくなります (会話の前にあるシステム プロンプトは残ります)。昔の会話がキャッシュに残っている限り、要約を書くこと自体はコストがかからないので、長い休憩の後よりも長い休憩の前の方がはるかにコストがかかりません。
時間: 毎ターン、時計はリセットされますが、キャッシュはサブスクリプションの場合は 1 時間、API キーの場合は 5 分後に期限切れになります (ENABLE_PROMPT_CACHING_1H=1 で 1 時間になります)。それ以降に戻ってくると、次のターンで会話全体が再びプレフィルされます。古いセッションを再開すると、ほとんどの場合、これも行われます。通常、その時点でキャッシュは失われており、とにかく起動時にシステム プロンプトが再構築されます。
これは、決してモデルや努力を切り替えるべきではないという意味ではありません。これは、セッションの開始時または /clear の直後など、それを実行するのに費用がかからない瞬間と、長い会話の途中など、費用がかかる瞬間があることを意味します。
ヒント: 最後の数ターンが保持したくない場所に移動した場合は、 /compact を実行する代わりに、 /rewind をその直前に戻します。巻き戻しはそれらのターンを末尾からカットするだけなので、それ以前のすべてはキャッシュされたままであり、コストはかかりません。圧縮すると会話全体が書き換えられるため、常に何らかのコストがかかります。
セッションが送信するトークンの数を決定するもの
ここで知っておくべき主な点は、一度だけ送信されるものは何もないということです。クロードが読み取ったファイルや実行したコマンドの出力など、会話で終了するすべての内容は、セッションの残りの間、その後のターンごとに再送信されます。
キャッシュされているため、各再送信は安価ですが、安いだけでは意味がなく、モードのコンテキスト内でスペースを占有しています。

私も毎ターン考えなければなりません。
これは、実際にはセッションのコスト モデル全体です。つまり、コンテキスト内で終了するトークンの数、そこに留まるターン数、および同時に実行するコンテキストの数です。
コンテキスト内の内容の一部は、ツール定義、システム プロンプト、 CLAUDE.md 、および起動時に読み込まれるその他のものなど、何かを入力する前に表示されます。
ヒント: 何かを入力する前に、新しいセッションで /context を実行して、その内容を確認してください。 CLAUDE.md を特定の命令に保ち、ワークフロー固有の命令をスキルに移動します。スキルは、使用時にのみ読み込まれます。このセッションに不要な MCP サーバーがある場合は、 /mcp を使用してオフにします。
セッション中に追加されるその他のほとんどすべてはツールの結果です。つまり、Claude が読み取るファイルと、Claude が実行するコマンドの出力です。
クロードがどれだけ本を読むかは、ほとんどの場合、どれだけ自分で理解する必要があるかによって決まります。 「テストが失敗している」という場合は、まずどのテストを特定する必要があります。1 つまたは 2 つの grep を実行し、どれが関連しているかを確認するためにいくつかのファイルを開いて、それらの結果はすべて、役に立たなくなった後もずっとコンテキスト内に残ります。
「utils.test.ts の失敗したテストを修正する」では検索がスキップされ、読み取り呼び出しが 1 回かかります。

[切り捨てられた]

## Original Extract

Practical tips for how to run efficient sessions that get the most value from every token.

Maximizing the value of your Claude Code sessions | Claude by Anthropic
Meet Claude Products Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Maximizing the value of your Claude Code sessions
Maximizing the value of your Claude Code sessions
How to run efficient sessions that get the most value from every token.
Category Claude Code Enterprise AI
Share Copy link https://claude.com/blog/maximizing-the-value-of-your-claude-code-sessions
Run /clear between tasks. This prevents prior irrelevant context from being sent back to the model, which can reduce token usage.
Set your model and effort level before you start. Changing either one mid-conversation can bust your prompt cache, which can increase token cost.
@-mention files instead of naming them. The file gets attached to your message directly, which saves a Read call, or a search if Claude has to go find it.
Add quiet flags to noisy commands, or run them in a subagent. Command output is added to the conversation just like a file, and stays there for the rest of the session.
Run /context once in a fresh session. It shows what's loaded ( CLAUDE.md , MCP tool definitions), so you can cut out anything unnecessary.
/compact before you take a break from your keyboard. The prompt cache expires after an hour, and summarizing a conversation is much cheaper while it's still cached.
Until pretty recently, the tools you wrote code with were a flat fee (or free). Your editor cost the same whether you fixed one test or fifty that afternoon, so an individual task didn't really have a price of its own.
With agentic coding tools like Claude Code, it does. The same completed task can also cost different amounts depending on how you use it.
In one session, Claude reads the test and the file it covers, makes the edit, and is done in a handful of turns. In another, it greps around the repo first, reads a dozen files on its way to the same two, and every one of those turns also drags along everything else that's been read into the conversation since this morning.
It's the same fix, but you spent a different number of tokens on it, and the whole time the model was also having to think about ten files it didn't need.
Being efficient with tokens doesn't mean using fewer of them overall. It means making sure the ones you do use go towards the thing you actually asked for.
So let's look at what decides the price of a token, then what decides how many of them a session sends, and along the way, what that means for how you run a session.
What decides the price of a token
You're billed per token, but what you're actually paying for is inference: the time it takes a GPU (or a TPU, or whatever the model happens to be running on) to run the model over your tokens.
Three things decide how much of that time a token takes: which model you're running, whether it's an input token (going in) or an output token (coming out), and whether it was cached.
A bigger model does more work on both input and output tokens. Which model is worth it for which kind of work is a topic on its own, and we covered it in Choosing a Claude model and effort level in Claude Code .
For this post, all you need to know is that everything else we're about to cover gets multiplied by the model's price: use a larger model when the problem is genuinely hard or ambiguous, and a smaller one when the work is routine.
Prev Prev 0 / 5 Next Next Get Claude Code Desktop
irm https://claude.ai/install.ps1 | iex Copy command to clipboard Or read the documentation Try Claude Code Try Claude Code Try Claude Code Developer docs Developer docs Developer docs eBook
A request goes through the GPU in two phases, and they cost different amounts.
First, during prefill, the model reads your request and context: the system prompt, your CLAUDE.md , your message, and everything that's been added to the conversation since (the files Claude has read and the output of the commands it ran). Those are your input tokens.
Then, during decode, it writes output tokens: its thinking, the tool calls it makes, and the text you see. This happens one token at a time; a 200-token response is 200 runs of the model, one after the other. Per token, decode keeps the GPU busy for a lot longer, which is why output is priced at roughly 5x input.
A lot of the output tokens in a session are thinking tokens, and how much thinking the model does per turn is what the effort level controls. Like the model, the level you pick with /effort sticks around as your default for the next session too.
Tip: run /model and /effort once in a fresh session to see what you're actually on. Both remember whatever you picked last time, and you want that decision to be deliberate.
Tip: if you already know a session is going to be grunt work, MAX_THINKING_TOKENS=0 claude turns thinking off for that one session (except on Fable 5), which is the step below /effort low.
If a request starts with exactly the same tokens as a request the server just saw, the state for that shared beginning comes out the same, so the server can keep it around from last time and only prefill whatever comes after it. This is called prompt caching.
Reading from the cache costs 0.1x the input price, because the server loads the state instead of computing it. Writing tokens into the cache costs a bit more than normal input, up to 2x, since the server also has to hold on to the state afterwards. But the write happens once per token, and the 0.1x reads happen on every turn after it.
Claude Code manages the prompt cache on every request, there's nothing to turn on. However you can break it, so it's important to know how to avoid these cost spikes.
Say we type "fix the failing test in utils.test.ts ". Here's what Claude Code sends for it:
Claude Code assembles the first request out of the system prompt (tool definitions included), your CLAUDE.md , and your message, and sends it off (input tokens). Nothing is in the cache yet, so all of it gets prefilled and written into the cache.
The model can't fix a test it hasn't seen, so it thinks for a moment and responds with a Read call for utils.test.ts (output tokens). Claude Code reads the file, appends it to the conversation, and sends the whole thing again (input tokens). This time everything from request 1 is read back out of the cache at a tenth of the price, and the only thing prefilled at full price is what's new: the Read call and the file.
Now the model wants the file under test (output). Another Read, another append, and everything goes out again: requests 1 and 2 from the cache, the second file at full price (input).
The model responds with an Edit (output). Claude Code applies it, appends the result, and sends everything again. Same story: the Edit and its result are new, everything in front of them is a cache read (input).
The model runs npm test (output). Claude Code appends the test output and sends everything again, with the test output as the only new part (input).
The tests pass, and the model responds with a short summary (output). No tool call means nothing to append and no request 6, so we're done.
That's five requests for one small fix, and every one of them contained the entire conversation up to that point. A typical turn is lopsided: tens of thousands of tokens going in, a few hundred coming out. But only what's new in that turn gets prefilled at full price.
That's the whole per-turn bill: cache reads on the history, full input price on whatever's new, and the output price on the response.
This applies on a subscription too. You don't see these prices directly, but the same requests are what draw down your limits.
The cache has to match from the very start of the request forward, and requests always go out in the same order: tool definitions, then the system prompt, then the conversation (with CLAUDE.md at the front of it).
If anything in that prefix changes, everything behind it gets prefilled again. A tool result appended to the end of the conversation is the ideal case, since nothing is behind it. What throws the cache away is anything that changes the request further towards the front, or changes what the cache is keyed on:
/model : every model has its own cache, so on the next turn the entire conversation gets prefilled again at full price. (This includes opusplan, which switches models every time you go in or out of plan mode.)
/effort: the effort level is part of what the cache is keyed on too, so it's the same story. It's why both /model and /effort ask you to confirm when you switch in the middle of a conversation.
Fast mode : also part of the key, and the re-prefill happens at fast mode prices, so if you're going to turn it on, turn it on at the start. (Turning it off again is free, cache-wise.)
/compact : the conversation gets replaced with a shorter one, so nothing in it matches anymore (the system prompt in front of it survives). Writing the summary itself is cheap as long as the old conversation is still in the cache, so it's a lot cheaper before a long break than after one.
Time: every turn resets the clock, but the cache expires after an hour on a subscription or five minutes on an API key ( ENABLE_PROMPT_CACHING_1H=1 makes it an hour). Come back later than that, and the next turn prefills the whole conversation again. Resuming an old session almost always does too: the cache is usually gone by then, and the system prompt gets rebuilt at launch anyway.
None of this means you should never switch models or effort. It means there are cheap moments to do it, the start of a session or right after a /clear , and expensive ones, the middle of a long conversation.
Tip: if the last few turns went somewhere you don't want to keep, /rewind to just before them instead of running /compact . Rewinding only cuts those turns off the end, so everything before them is still cached and it costs nothing. Compacting rewrites the whole conversation, so it always costs something.
What decides how many tokens a session sends
The main thing to know here is that nothing gets sent just once. Everything that ends up in the conversation, a file Claude read or the output of a command it ran, gets sent again on every turn after it, for the rest of the session.
It's cached, so each of those re-sends is cheap, but cheap isn't nothing, and it's taking up room in the context the model has to think around on every turn too.
That's really the whole cost model of a session: how many tokens end up in the context, how many turns they stay there, and how many contexts you're running at the same time.
Part of what's in the context is there before you type anything: the tool definitions, the system prompt, CLAUDE.md , and whatever else gets loaded at startup.
Tip : run /context in a fresh session to see what's in there before you've typed anything. Keep CLAUDE.md to specific instructions and move workflow-specific ones into skills, which only get loaded when they're used. If there's an MCP server you don't need in this session, turn it off with /mcp .
Nearly everything else that gets added during the session is tool results: the files Claude reads, and the output of the commands it runs.
How much Claude reads mostly comes down to how much it has to figure out on its own. If you say "the tests are failing", it first has to find out which tests: a grep or two, a few files opened to see which one is relevant, and all of those results stay in the context long after they've stopped being useful.
"Fix the failing test in utils.test.ts " skips the searching and costs one Read call for the

[truncated]
