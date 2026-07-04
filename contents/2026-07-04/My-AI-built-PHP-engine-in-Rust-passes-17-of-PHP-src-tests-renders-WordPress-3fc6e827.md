---
source: "https://ekinertac.com/blog/i-dont-know-rust-my-ai-is-rewriting-php-in-it/"
hn_url: "https://news.ycombinator.com/item?id=48789325"
title: "My AI-built PHP engine in Rust passes 17% of PHP-src tests, renders WordPress"
article_title: "I Don't Know Rust. My AI Is Rewriting PHP in It Anyway. - ekinertac"
author: "ekinertac"
captured_at: "2026-07-04T21:49:03Z"
capture_tool: "hn-digest"
hn_id: 48789325
score: 1
comments: 0
posted_at: "2026-07-04T21:35:30Z"
tags:
  - hacker-news
  - translated
---

# My AI-built PHP engine in Rust passes 17% of PHP-src tests, renders WordPress

- HN: [48789325](https://news.ycombinator.com/item?id=48789325)
- Source: [ekinertac.com](https://ekinertac.com/blog/i-dont-know-rust-my-ai-is-rewriting-php-in-it/)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T21:35:30Z

## Translation

タイトル: Rust で AI が構築した PHP エンジンが PHP-src テストの 17% に合格し、WordPress をレンダリング
記事タイトル：Rustが分かりません。私の AI はとにかく PHP を書き換えています。 - エキナータック
説明: できます

記事本文:
Rustは分かりません。私の AI はとにかく PHP を書き換えています。 - ekinertac > _ ekinertac プロジェクト
← エッセイ 2026年7月5日
プログラミング
錆びる
php
llm
エージェント
ファーゴ
クロード 私はRustを知りません。私の AI はとにかく PHP を書き換えています。
数日前の夜、私は端末が 26 KB の WordPress フロントページ、<title>Phargo Test Site</title>、ブロック ライブラリ CSS、「Hello world!」を印刷するのを見ました。 SQLite データベースから取得され、最後にクリーンな </html> が表示されます。 1 つの詳細を除いて、まったく目立った出力はありません。出力された PHP エンジンには、PHP の実際のソース コードが 0 行も含まれていません。これは、Rust で書かれた最初からのインタプリタです。
ここで、皆さんに理解していただきたいのは、私は Rust を知りません。私はレクサーを書いたことはありません。別のタブにある Wikipedia の記事を読まなければ、「ツリーウォーキング評価者」が何であるかを説明することはできません。パーティーで私を追い詰めて、PHP のガベージ コレクターがどのように機能するのか尋ねたら、私は電話を装うでしょう。
このエンジンは Phargo と呼ばれており、これに対する私の貢献は、大まかに言えば、エイミングです。 AIがコードを書きます。私はそれを標的に向け、海図を検討する中世の王のように返ってきた内容を読み（厳粛にうなずき、理解力ゼロ）、現代のソフトウェア開発において最も強力なフレーズを入力します。
実験: 構築システムとしての徹底した誠実さ
現在、誰もが、そしてその観葉植物が AI によって構築されたプロジェクトを抱えており、どのプロジェクトも「それはうまくいく!」という反証不可能な同じ主張をしています。誰に従って機能しますか？それを書いたのはAIですか？ 4テイク目に録音されたデモ？
したがって、実験全体は、Bun チームが現実世界のテスト スイートに対して JavaScript ランタイムを実行する様子を観察したことから借用した 1 つのアイデアに基づいています。それは、AI に宿題を採点させないことです。
PHP には独自のテスト スイートが同梱されており、約 22,000 の .phpt ファイルが書き込まれています。

30 年以上にわたる PHP 内部チーム。私が書いていないテスト。 AIが書かなかったテスト。 DateTime の夏時間計算から、float に対して var_dump() が正確に出力するものまで、言語の隅々までエンコードするテスト。そのスイートはオラクルです。スコアボードはすべてを実行し、実行ごとに合格率がリポジトリに自動生成されます。
この数字はお世辞にしたり、交渉したり、気分を良くしたりすることはできません。 bug40261.phpt がパスするか、パスしないかのどちらかです。
現在のスコア: 22,037 点中 3,844 点 — アップストリーム PHP テスト スイート全体の 17.4%。 17% と鼻で笑う前に、現実的な上限は約 40 ～ 45% です。スイートの残りの部分は、明示的に範囲外である C 拡張機能 (GD、curl、SOAP、intl、MySQL ドライバーなど) をテストするためです。実際の競技場では、登りは非常に現実的であり、ゼロから始まりました。
人間としての私のループは、恥ずかしいほど薄いです。
AI はコーパス全体に対して失敗ヒストグラムを実行し、実際に修正できる失敗したテストの最大のクラスターを見つけます。
約 22,000 のテスト スコアボードを実行します (約 7 分間のファン騒音)
数値が上がった場合: コミット、プッシュ、繰り返し
数値が下がった場合: 「うーん、下がった、もう一度見てください」という別のセリフを言うようになります。
それだけです。それが仕事です。私はピークの委任を達成しましたが、申し訳ないとさえ思っていません。
オラクルは賄賂を受け取ることはできません。しかし、ハーネスは私の顔に嘘をつきました。
早い段階で、合格率は間違っていると感じられるほど頭打ちになりました。テストのカテゴリ全体 (明らかに単純なテスト) は、予想される出力と同一に見える差分で失敗し続けました。私は、間違い探しパズルで 2 枚の同じ写真を見つめる男性のように、それらの差分を見つめましたが、何も見つかりませんでした。
キャリッジリターンという文字通り目に見えないものであるため、違いは目に見えませんでした。テストコーパスはチェックアウトされていました。

n CRLF 行末を持つ Windows とスコアボードは出力をバイトごとに比較しました。 PHP 独自のテスト ランナーは、比較する前に行末を正規化します。私たちのものはそうではありませんでした。つまり、ハーネスは行末だけでコーパス内の基本的にすべての複数行テストに静かに失敗しており、それが何週間も続いていました。
正規化コードは 1 行です。何百ものテストが即座に緑色に切り替わりました。
「自分の寸法を測ってください」という教訓がプロジェクトに反映されました。あなたのオラクルは、あなたをそれに接続する配管と同じくらい正直です。現在、run-tests.php とまったく同じ方法で正規化しています。それ以降、疑わしいプラトーが発生すると、最初に同じ疑問が生じます。エンジンが間違っているのか、それともスコアボードが嘘をついているのでしょうか?
PHP のテスト スイートは Readme を備えた地雷原である
他人の 22,000 ファイルのテスト コーパスを実行することについて誰も教えてくれないことがあります。それらのファイルの一部は爆弾です。悪意のあるものではなく、偶然のものです。不条理な構造を割り当てる古代のメモリ バグの回帰テスト、無限に拡張するジェネレータ テスト、PHP 独自の慎重にフェンスされた CI 内でのみ実行されることを目的としたテスト。
私はこれを、すべての偉大な発見が行われる方法で見つけました。つまり、開発マシンをハード再起動しました。 「プログラムがクラッシュした」ではありません。 「端末がフリーズした」ではありません。コンピューター全体が真っ暗になって再起動したのは、発電機のテストでエンジンがブレーキのないロケット弾のショッピングカートのように家中の RAM をすべて食い尽くしてしまったためです。
その余波でエンジンは偏執的なものになったが、正直なところ、かなり偏執的な状態になっている。
上限付きのグローバル アロケータ - テストがどれほど独創的であっても、エンジンは物理的に 6 GiB を超える割り当てを行うことができません。
ステップ制限があるため、無限ループはスペース ヒーターではなくエラーで終了します。
文字列サイズ、配列ノード、出力長、ジェネレータ拡張の上限
あ

ブレッドクラム ファイルを使用すると、スコアボードが現在のテスト名で更新されるため、何かがハングしたときにどのファイルを確認すればよいか正確にわかります。
これらはどれも魅力的な言語実装作業ではありません。すべては、「研究プロジェクト」と「コーヒーを淹れている間、無人で 22,000 の敵対的なファイルを安全に読み進めることができるもの」との違いです。
静かに眠っていた機能
私の好きなバグのジャンルは、コーパスが容赦なくバグを検出するもので、存在し、解析され、エラーなしで実行され、まったく何も行わない機能です。ポチョムキン内蔵。数か月間にわたって、スイートは次のようなことを明らかにしました。
clone — 適切に解析され、 NULL と評価されます。エンジン全体。不変の日付計算はクローンで作られているため、すべてのテストのすべての DateTimeImmutable は静かに永久に壊れていました。
unset($arr[$key]) — まったく何もしません。鍵はただ…そのまま残っていた
Trim($str, $charlist) — 初期から charlist 引数を無視し、とにかく空白をトリミングしました
$$variableVariables — 存在しませんでした
静的関数変数 - 存在しませんでした
spl_autoload_register() — 温かい笑顔でオートローダーを受け入れ、決して呼び出しませんでした
catch (\Throwable) — 何も一致しませんでした。これは、キャッチオールが持つ非常に面白いプロパティです。
これらはいずれもデモで生き残るだろう。それぞれは、Rust を読まない人 (私) によるコードレビューに耐えることができます。彼らの誰もコーパスから生き残っていませんでした。これが、この実験の全理論を 1 つの箇条書きリストにまとめたものです。私はコードを監査できないので、22,000 のテストが私に代わってコードを監査します。その徹底ぶりは、人間のレビュー担当者が昼食を過ぎても耐えられないほどです。
そしてWordPressページを提供しました
北極星は常に WordPress でした。これは PHP 互換性の最終ボスであり、2003 年以降のあらゆる PHP イディオムの堆積層を含むのに十分な古いコードベースです。
wp-load.php を b にする

ootstrap は、言語弁護士の熱狂的な夢のような一連のブロッカーを通過しました。 goto (はい、WordPress の HTML パーサーは goto を使用します)、 str_replace の参照 $count パラメータ、正規表現文字クラス内での \xNN のエスケープ、組み込み関数の半分を認識していない function_exists() です。 preg_split が wpdb::prepare 内の PREG_SPLIT_DELIM_CAPTURE フラグを無視していたため、インストーラーは自身のデータベースを破損しました。このバグは、私が曇りガラス越しに開胸手術を見守る男性のように自信を持って監督していた間に、AI によって診断、発見、修正できたバグの 4 層下にありました。
そしてある晩、wp_install() が完了します。管理者ユーザーが作成され、オプション テーブルが設定され、SQLite に 3 つの投稿が作成されました。フロントページには、実際のテーマ、実際の投稿、実際のパーマリンクが表示されます。
徹底的な正直さがすべてであるため、完全な開示を行います。
✅ 新規インストールが実行され、フロントページがデータベースからレンダリングされます
✅ /wp-admin/ もレンダリングされます — 実際のダッシュボードは問題なく表示され、正直言ってフロントページよりも驚きました
⚠️ REST API は未踏の領域です
⚠️ 現在、そのページでは実際の PHP よりも約 55 倍遅い (7.1 秒対 126 ミリ秒) — ただし、ピカピカの新しいバイトコード VM はすでに PHP 8.5 の 1 ～ 3 倍でマイクロベンチマークを実行しており、次はその数値で実行される予定です。
私は、AI が言語エンジンを作成できるかどうかを学ぶことを期待して参加しました。驚くべき答えは、これは本当の質問ではなかったということです。もちろん、レクサーを作成することもできます。これまでに発行されたすべてのレクサーを読み取ります。本当の疑問は常に「コードを検証できない人が、どうやってプロジェクトを誠実に保つことができるのか」ということでした。そして、その答えは時代遅れで、退屈で、美しいものであることが判明しました。テストは他の人が書いたもので、現実が動いたときにのみ動く数値であり、結果がフロリダかどうかにかかわらず、すべての結果が公開リポジトリにプッシュされました。

私たちを攻撃するかどうか。
Rustはまだ分かりません。エンジンは現在、約 24,000 行あります。スコアボードには一度に数十のテストが表示され、開発ログには戦争の話が収集され、どこかに、WASM にコンパイルされた Rust エンジン上のブラウザで WordPress が実行されるバージョンがあります。
数値の上昇を観察してください: github.com/ekinertac/Phargo 。あるいは、そんなことはせず、パーサーを書けない人がどこかでパーサーを出荷しているということを知って楽しむだけです。
ああ、そして当然のことですが、この投稿も LLM で下書きされ、私がそれらしく聞こえるまで編集しました。機械が書く、私が目指す。それが要点だ。

## Original Extract

I can

I Don't Know Rust. My AI Is Rewriting PHP in It Anyway. - ekinertac > _ ekinertac projects
← essays July 5, 2026
programming
rust
php
llm
agents
phargo
claude I Don’t Know Rust. My AI Is Rewriting PHP in It Anyway.
A few nights ago I watched my terminal print out a 26 KB WordPress front page — <title>Phargo Test Site</title> , the block-library CSS, “Hello world!” pulled from a SQLite database, and a clean </html> at the bottom. Completely unremarkable output, except for one detail: the PHP engine that served it contains zero lines of PHP’s actual source code . It’s a from-scratch interpreter written in Rust.
Here’s the part I need you to sit with: I don’t know Rust. I have never written a lexer. I could not explain to you what a “tree-walking evaluator” is without reading the Wikipedia article in another tab. If you cornered me at a party and asked how PHP’s garbage collector works, I would fake a phone call.
The engine is called Phargo , and my contribution to it is, roughly, aiming. An AI writes the code. I point it at a target, read what comes back like a medieval king reviewing naval charts — solemn nod, zero comprehension — and type the most powerful phrase in modern software development:
The Experiment: Radical Honesty as a Build System
Everyone and their houseplant has an AI-built project right now, and every single one comes with the same unfalsifiable claim: “it works!” Works according to whom? The AI that wrote it? The demo that was recorded on the fourth take?
So the whole experiment rests on one idea, borrowed from watching the Bun team drive their JavaScript runtime against real-world test suites: don’t let the AI grade its own homework.
PHP ships with its own test suite — about 22,000 .phpt files written by the PHP internals team over three decades. Tests I didn’t write. Tests the AI didn’t write. Tests that encode every cursed corner of the language, from DateTime daylight-saving math to what exactly var_dump() prints for a float. That suite is the oracle. The scoreboard runs all of it, and the pass rate is auto-generated into the repo after every run.
The number cannot be flattered, negotiated with, or prompted into a better mood. Either bug40261.phpt passes or it doesn’t.
Current score: 3,844 of 22,037 — 17.4% of the entire upstream PHP test suite. And before you snort-laugh at 17%: the realistic ceiling is around 40–45%, because the rest of the suite tests C extensions (GD, curl, SOAP, intl, MySQL drivers…) that are explicitly out of scope. Within the actual playing field, the climb is very real — it started at zero.
My loop as the human is almost embarrassingly thin:
The AI runs a failure histogram over the whole corpus to find the biggest cluster of failing tests it can actually fix
It runs the ~22,000-test scoreboard (about 7 minutes of fan noise)
If the number went up: commit, push, repeat
If the number went down: I get to say my other line, “hmm, that regressed, look again”
That’s it. That’s the job. I’ve achieved peak delegation and I’m not even sorry.
The Oracle Cannot Be Bribed. The Harness, However, Lied to My Face.
Early on, the pass rate plateaued in a way that felt wrong. Whole categories of tests — obviously simple ones — kept failing with diffs that looked identical to the expected output. I stared at those diffs like a man staring at two identical photos in a spot-the-difference puzzle, finding nothing.
The difference was invisible because it was literally invisible: carriage returns. The test corpus had been checked out on Windows with CRLF line endings, and our scoreboard compared output byte-for-byte. PHP’s own test runner normalizes line endings before comparing. Ours didn’t. Which means the harness was silently failing essentially every multi-line test in the corpus on line endings alone, and had been for weeks.
One line of normalization code. Hundreds of tests flipped to green instantly.
The lesson tattooed itself onto the project: measure your measurement. Your oracle is only as honest as the plumbing that connects you to it. We now normalize exactly the way run-tests.php does, and every suspicious plateau since has triggered the same question first — is the engine wrong, or is the scoreboard lying?
PHP’s Test Suite Is a Minefield With a Readme
Here’s something nobody tells you about running somebody else’s 22,000-file test corpus: some of those files are bombs. Not malicious ones — accidental ones. Regression tests for ancient memory bugs that allocate absurd structures, generator tests that expand into infinity, tests that were only ever meant to run inside PHP’s own carefully-fenced CI.
I found this out the way all great discoveries are made: my development machine hard-restarted . Not “the program crashed.” Not “the terminal froze.” The entire computer went black and rebooted, because a generator test convinced our engine to eat every byte of RAM in the house like a rocket-powered shopping cart with no brakes.
The aftermath turned the engine paranoid, and honestly it wears paranoia well:
A capped global allocator — the engine physically cannot allocate more than 6 GiB, no matter how creative the test gets
A step limit, so infinite loops die with an error instead of a space heater
Caps on string sizes, array nodes, output length, generator expansion
A breadcrumb file the scoreboard updates with the current test name, so when something hangs, we know exactly which file to glare at
None of this is glamorous language-implementation work. All of it is the difference between “research project” and “thing that can safely chew through 22,000 hostile files unattended while I make coffee.”
The Features That Were Quietly Lying
My favorite genre of bug — and the corpus finds them relentlessly — is the feature that exists, parses, runs without error, and does absolutely nothing. The Potemkin builtin. Over the months the suite has exposed, among others:
clone — parsed fine, evaluated to NULL . Engine-wide. Every DateTimeImmutable in every test had been silently broken forever, because immutable date math is made of clone
unset($arr[$key]) — a total no-op. The key simply… stayed
trim($str, $charlist) — ignored the charlist argument since the beginning of time and trimmed whitespace anyway
$$variableVariables — didn’t exist
static function variables — didn’t exist
spl_autoload_register() — accepted your autoloader with a warm smile and never called it
catch (\Throwable) — matched nothing, which is a very funny property for a catch-all to have
Each of these would survive a demo. Each would survive a code review by someone (me) who doesn’t read Rust. None of them survived the corpus. That’s the entire thesis of the experiment in one bullet list: I can’t audit the code, so the 22,000 tests audit it for me, with a thoroughness no human reviewer could sustain past lunch.
And Then It Served a WordPress Page
The north star was always WordPress — it’s the final boss of PHP compatibility, a codebase old enough to contain sedimentary layers of every PHP idiom since 2003.
Getting wp-load.php to even bootstrap burned through a chain of blockers that reads like a language-lawyer’s fever dream: goto (yes, WordPress’s HTML parser uses goto ), str_replace ’s by-reference $count parameter, \xNN escapes inside regex character classes, function_exists() being blind to half our builtins. The installer then corrupted its own database because preg_split was ignoring the PREG_SPLIT_DELIM_CAPTURE flag inside wpdb::prepare — a bug four layers below anything I could have diagnosed, found and fixed by the AI while I supervised with the confidence of a man watching open-heart surgery through frosted glass.
And then one evening: wp_install() completes. Admin user created, options table populated, three posts in SQLite. The front page renders — real theme, real posts, real permalinks.
Full disclosure, because radical honesty is the whole bit:
✅ Fresh install runs, front page renders from the database
✅ /wp-admin/ renders too — the actual dashboard, without any issues, which frankly surprised me more than the front page did
⚠️ The REST API is unexplored territory
⚠️ It’s currently ~55x slower than real PHP on that page (7.1 s vs 126 ms) — though the shiny new bytecode VM already runs micro-benchmarks at 1–3x of PHP 8.5 and is coming for that number next
I went in expecting to learn whether an AI could write a language engine. The surprising answer is that this was never really the question. Of course it can write a lexer — it’s read every lexer ever published. The real question was always: how does a person who can’t verify the code keep the project honest? And the answer turned out to be old-fashioned, boring, and beautiful: tests somebody else wrote, a number that only moves when reality moves, and every result pushed to a public repo whether it flatters us or not.
I still don’t know Rust. The engine is now ~24,000 lines of it. The scoreboard goes up a few dozen tests at a time, the dev log collects war stories, and somewhere out there is a version of this where WordPress runs in your browser on a Rust engine compiled to WASM.
Watch the number climb: github.com/ekinertac/Phargo . Or don’t, and just enjoy knowing that somewhere, a man who cannot write a parser is shipping one.
Oh, and obviously: this post was drafted with an LLM too, then edited by me until it sounded like me. The machine writes, I aim. It’s the whole point.
