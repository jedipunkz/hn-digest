---
source: "https://www.al3rez.com/building-zehn"
hn_url: "https://news.ycombinator.com/item?id=48545494"
title: "Zehn, a fuzzy finder for every prompt I've sent an AI agent"
article_title: "zehn, a fuzzy finder for every prompt I've sent an AI agent — al3rez"
author: "speckx"
captured_at: "2026-06-15T19:24:50Z"
capture_tool: "hn-digest"
hn_id: 48545494
score: 1
comments: 0
posted_at: "2026-06-15T18:54:09Z"
tags:
  - hacker-news
  - translated
---

# Zehn, a fuzzy finder for every prompt I've sent an AI agent

- HN: [48545494](https://news.ycombinator.com/item?id=48545494)
- Source: [www.al3rez.com](https://www.al3rez.com/building-zehn)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T18:54:09Z

## Translation

タイトル: Zehn、AI エージェントに送信したすべてのプロンプトに対するファジー ファインダー
記事のタイトル: zehn、AI エージェントに送信したすべてのプロンプトに対するファジー ファインダー — al3rez
説明: 4 つのコーディング エージェントを使用していますが、それらのエージェント全体で古いプロンプトが見つかりませんでした。そこで私は、彼らの履歴をすべて検索し、セッションに戻すツールを作成しました。

記事本文:
zehn、AI エージェントに送信したすべてのプロンプトのあいまいファインダー
ある日はclaudeを使用し、次の日はcodexを使用し、その後はpiまたはopencodeを使用します。それらはすべて、私が入力した内容の履歴を保持します。問題は、1 週間後、頼んだことをやり返したいと思ったときに、どのエージェントに言ったのか思い出せないことです。あるいは、自分がどのプロジェクトに参加していたのか。そのため、4 つの異なる履歴ファイルを 4 つの異なる形式で手作業で調べてみましたが、たいていは諦めていました。
それで私はツェンを作りました。 4 つの履歴すべてを一度に読み取り、これまでに送信したすべてのプロンプトをあいまい検索できます。いくつかの文字を入力し、目的の文字を見つけて Enter キーを押すと、そのセッションを所有するエージェントの正確なセッションに戻ります。
名前はذهن、「心」です。これは 1 つの小さな Zig バイナリです。
問題は、フォーマットについて誰も同意していないことです
これは簡単なことのように思えます。いくつかの履歴ファイルを読み取り、リストを表示します。厄介な点は、エージェントごとに履歴の保存方法が異なるため、4 つすべてを学習する必要があることです。
クロードは、history.jsonl をホーム ディレクトリに保存します。 1 行に 1 つの JSON オブジェクトが表示され、プロンプトが表示されます。 codex には独自のhistory.jsonlがありますが、プロンプトはテキストの下にあり、プロジェクトは添付されていません。 pi はセッションごとにフォルダーを書き込みます。それぞれが .jsonl ファイルで、最初の行がセッション メタデータで、後の行がメッセージです。また、opencode はファイルをまったく使用しません。 SQLiteデータベースです。
したがって、zehn にはエージェントごとに 1 つずつ、計 4 つの小さなリーダーがあり、それらはすべて同じものを生成します。つまり、プロンプト テキスト、プロジェクト、セッション ID、およびタイムスタンプ (存在する場合) を含むレコードです。すべてが記録されると、ツールの残りの部分はそれがどこから来たのかを気にしません。
そのうちのいくつかは小さな変化球を投げた。メッセージのコンテンツはプレーンな文字列である場合もあれば、型指定されたブロックの配列である場合もあるため、両方を処理する必要がありました。 opencode は sh を意味しました

sqlite3 コマンドを使い始めましたが、私はそれが好きではありませんでしたが、Zig で SQLite リーダーをゼロから書くのは、私が死にたいと思っていた丘ではありませんでした。 sqlite3 がインストールされていない場合、zehn はオープンコードをスキップしてその理由を説明します。
それから重複排除を行います。同じプロンプトを 5 回入力した場合は、最新のエントリが 1 つ必要になります。
fzf のような感じにしたかったのです。それが私がすでに持っているマッスルメモリーだからです。文字を入力すると、最も一致する文字が上位にランク付けされ、一致した文字が強調表示されます。
したがって、これは部分文字列検索ではなく、実際のファジーマッチャーです。これは fzf と同じ方法で一致をスコアします。つまり、単語の先頭での一致、連続する文字の一致、キャメルケース境界のボーナスです。スコアリングは小さな動的プログラミングの調整として実行され、異常に長い行には貪欲なフォールバックがあるため、何を投げても高速なままです。私はテストでその隣に総当たりスコアラーを作成しました。これは純粋に、何千ものランダムな入力に対して高速バージョンが明らかだが遅いバージョンと一致することを確認するためでした。
そのテストでは、私が認めたい以上に多くのバグが見つかりました。
実際に魔法のように感じる部分
プロンプトを見つけるだけで半分は完了です。残りの半分はセッションに戻ることであり、それが実際にツールを使用する部分です。
各エージェントには独自の再開方法があります。クロードは claude --resume <id> を望んでいます。 codex は codex 履歴書 <id> を必要としています。 pi と opencode にはそれぞれ独自のものがあります。したがって、プロンプトを選択すると、zehn はそれがどのエージェントに属しているかを認識し、適切なコマンドを作成して、それを生成します。また、セッションが最初に実行されたプロジェクト ディレクトリに cd が書き込まれます。これは通常、エージェントが期待していることだからです。そのディレクトリがなくなった場合は、現在の場所に戻って通知します。
検索し、Enter キーを押し、エージェントに再度入力します。セッション ID をコピー＆ペーストすることはできません

ログファイル。
正直に言うと、私が好きだからという理由もあります。しかし、このようなツールには本当の理由があります。
これはランタイムのない 1 つの静的バイナリです。インストーラーがそれをビルドし、 ~/.local/bin にドロップするだけです。これですべてが完了します。同時にインストールするものはなく、最新の状態に保つための言語ランタイムもありません。
もう 1 つの理由はアロケータです。 Zig ではアロケータを渡すため、背後で何も割り当てられません。私はプログラム全体に 1 つの領域を与えました。つまり、これらすべての履歴を読み取っている間は自由に割り当て、何も解放せず、プロセスが終了するときにすべてをドロップするようにしました。起動して 1 つのジョブを実行して終了するツールとしては、まさに正しい形状です。約 1,300 のセッションを約 0.2 秒で読み取り、解析します。速度を上げるために費やした時間のほとんどは、必要のない作業を行っていないだけでした。
Zig は 1.0 より前であり、標準ライブラリがあなたの下に移動します。 I/O API が最近変更され、問題が発生しました。完全に理解したい小さなツールについては、その取引に応じます。
bash <(カール -L https://al3rez.com/zehn)
コードは github.com/al3rez/zehn にあります。

## Original Extract

I use four coding agents and could never find an old prompt across them. So I built one tool that searches all of their histories and drops me back into the session.

zehn, a fuzzy finder for every prompt I've sent an AI agent
I use claude one day, codex the next, then pi or opencode after that. They all keep a history of what I’ve typed. The problem is that a week later, when I want to get back to something I asked for, I can’t remember which agent I said it to. Or which project I was in. So I’d go digging through four different history files by hand, in four different formats, and usually give up.
So I built zehn. It reads all four histories at once and lets me fuzzy-search every prompt I’ve ever sent. I type a few letters, find the one I want, hit enter, and it puts me back in that exact session, in the agent that owns it.
The name is ذهن, “the mind.” It’s one small Zig binary.
the problem is that nobody agrees on a format
This sounds simple. Read some history files, show a list. The annoying part is that every agent stores its history differently, and I had to learn all four.
claude keeps a history.jsonl in your home directory. One JSON object per line, with the prompt under display . codex has its own history.jsonl , but the prompt is under text and there’s no project attached. pi writes a folder per session, each one a .jsonl file where the first line is session metadata and later lines are messages. And opencode doesn’t use files at all. It’s a SQLite database.
So zehn has four little readers, one per agent, and they all produce the same thing: a record with the prompt text, the project, the session id, and a timestamp if there is one. Once everything is a record, the rest of the tool doesn’t care where it came from.
A couple of these threw small curveballs. A message’s content is sometimes a plain string and sometimes an array of typed blocks, so I had to handle both. opencode meant shelling out to the sqlite3 command, which I didn’t love, but writing a SQLite reader from scratch in Zig was not the hill I wanted to die on. If sqlite3 isn’t installed, zehn just skips opencode and tells you why.
Then I dedupe. If you’ve typed the same prompt five times, you want one entry, the most recent one.
I wanted it to feel like fzf, because that’s the muscle memory I already have. Type letters, get the best matches ranked at the top, with the matched characters highlighted.
So it’s a real fuzzy matcher, not a substring search. It scores matches the way fzf does: bonuses for matching at the start of a word, for matching consecutive characters, for camelCase boundaries. The scoring runs as a small dynamic-programming alignment, and there’s a greedy fallback for pathologically long lines so it stays fast no matter what you throw at it. I wrote a brute-force scorer next to it in the tests, purely to check the fast version agrees with the obvious-but-slow version on thousands of random inputs.
That test caught more bugs than I’d like to admit.
the part that actually feels like magic
Finding the prompt is half of it. The other half is getting back into the session, and that’s the bit that makes me actually use the tool.
Each agent has its own way to resume. claude wants claude --resume <id> . codex wants codex resume <id> . pi and opencode each have their own. So when you pick a prompt, zehn knows which agent it belongs to, builds the right command, and spawns it for you. It also cd s into the project directory the session ran in first, because that’s usually what the agent expects. If that directory is gone, it falls back to where you are and tells you.
You search, you hit enter, you’re typing to the agent again. No copy-pasting session ids out of a log file.
Honestly, partly because I like it. But there are real reasons for a tool like this.
It’s one static binary with no runtime. The installer builds it and drops it in ~/.local/bin , and that’s the whole story. Nothing to install alongside it, no language runtime to keep current.
The other reason is the allocator. In Zig you pass the allocator in, so nothing allocates behind your back. I gave the whole program a single arena: allocate freely while reading all those histories, never free anything, and let it all drop when the process exits. For a tool that starts, does one job, and quits, that’s exactly the right shape. It reads and parses around 1,300 sessions in about 0.2 seconds, and most of the time I spent on speed was just not doing work I didn’t need to do.
Zig is pre-1.0 and the standard library moves under you. The I/O API changed recently and broke things. For a small tool I want to understand completely, I’ll take that trade.
bash <( curl -L https://al3rez.com/zehn)
The code is at github.com/al3rez/zehn .
