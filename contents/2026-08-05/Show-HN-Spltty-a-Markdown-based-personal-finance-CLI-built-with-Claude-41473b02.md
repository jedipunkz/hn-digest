---
source: "https://tdiniz.dev/thoughts/building-a-personal-finance-tracker-with-claude"
hn_url: "https://news.ycombinator.com/item?id=49187132"
title: "Show HN: Spltty – a Markdown-based personal finance CLI built with Claude"
article_title: "Building a personal finance tracker with Claude · Thiago Diniz"
author: "thiagovsdiniz"
captured_at: "2026-08-05T19:14:51Z"
capture_tool: "hn-digest"
hn_id: 49187132
score: 5
comments: 0
posted_at: "2026-08-05T18:43:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Spltty – a Markdown-based personal finance CLI built with Claude

- HN: [49187132](https://news.ycombinator.com/item?id=49187132)
- Source: [tdiniz.dev](https://tdiniz.dev/thoughts/building-a-personal-finance-tracker-with-claude)
- Score: 5
- Comments: 0
- Posted: 2026-08-05T18:43:10Z

## Translation

タイトル: Show HN: Spltty – Claude で構築されたマークダウンベースのパーソナルファイナンス CLI
記事のタイトル: クロードとチアゴ・ディニスによるパーソナル・ファイナンス・トラッカーの構築
説明: Markdown 台帳のフォルダーを、Claude がファジー入力を処理している間に計算を処理する決定論的な CLI にどのように変換したか。
HN テキスト: 私は、プレーンな Markdown ファイルを使用して共有経費、カスタム分割、決済を追跡するための Ruby CLI である Spltty を構築しました。この記事では、Claude が管理するフォルダーから CLI にどのように進化したかについて説明しています。

記事本文:
Claude とパーソナル ファイナンス トラッカーを構築する · Thiago Diniz tdiniz の取り組み キャリアに関する考察 ← 考察に戻る
クロードと一緒に個人財務トラッカーを構築する
妻と私は、ほとんどのカップルがやっているようにお金を共有します。つまり、非常に貧弱な文書を使います。彼女のもの、私のもの、そして私たちのものもありますが、必ずしも 50/50 ではありません。ある時点では、結婚式、アパートの改築、モントリオールへの 6 週間の旅行がすべて同時に行われていました。記録は、PDF 明細書、クレジット カード請求書のスクリーンショット、Pix の領収書 (Pix はブラジルの即時決済システムです) の写真として届き、あるいは私が「ドレスのフィッティングに 700 レアル支払いました」と大声で言うこともありました。
一緒に住む前は、私たちが追跡しようとしていたのは旅行だけでした。私たちは割り勘アプリを使用しましたが、通常は旅行ごとに異なるアプリでした。どれも引っかかりませんでした。
簿記は私の仕事で、いつも事後処理をしていました。記憶と領収書の束から1週間分のタクシーと夕食を再構築し、数日後に出費が発生したときに記録できるように設計されたアプリにすべてを入力しました。一緒に住むことに決めたとき、リスクが高まったので、プロセス全体をスプレッドシートに移行しました。
スプレッドシートも同じ理由で腐ります。退屈な部分は数学ではありません。クレジット カードの明細から 40 行を転記し、それぞれが誰に属するかを判断するのです。
そこで私は可能な限り愚かなシステムからやり直しました。アプリやデータベースを持たない、記録システムとしてのプレーンテキストであるledger-cliに触発されて、Markdownファイルのフォルダーと、クロードにそれらの入力方法を指示する単一のテキストファイルを作成しました。
そのシステムも壊れるまで数週間続いた。転写の問題は解決されました。数学の問題はそうではありませんでした。 LLM に数百行の合計を求めることは、小学校の算数に数千のトークンを費やすことを意味します

ic を使用して、生成されるすべての数値を毎週チェックします。お金の計算は決定的で再現可能である必要があります。
そこで、計算を把握するために CLI を作成しました。 10 週間後、Markdown ファイルのフォルダーには独自のツールと 86 個のテストが含まれていました。
お急ぎですか？試してみましょう: https://github.com/thiagodiniz/spltty
プロジェクトがどのように進化したかについてもっと知りたい場合は、一度に 1 つのイテレーションを確認してください。読み続けてください
最初の反復: プレーン ファイルと 1 つの命令ファイル
全体的な手順を含む CLAUDE.md と、経費といくつかのメモを保持する 1 つのファイルを追加しました。
最初の経費表は次のようになります。
|日付 |タイトル |価値 (R$) |責任者 |出典 |
|-----------|----------|-----------:|-------------|----------|
それでおしまい。私がカード請求書のスクリーンショットをドロップすると、クロードは 38 行を抽出し、ファイルに書き込む前に私と一緒に各行を確認しました。
当時、私はアプリを構築していませんでした。データの形状とそれを埋めるためのルールを文書化していました。その後に登場したものはすべて、その形状を改良したものでした。
2 回目の反復: リアルマネーがスキーマを破壊した
2 週間後、スキーマは、使用することによってのみ発見できる方法で崩れ始めました。
「責任者」は二つの仕事をしていました。口論でカミラが言ったとき、それは彼女がその代金を払ったという意味ですか、それとも借りがあるという意味ですか？これらは異なる事実であり、両方がなければ和解は不可能です。したがって、それは 2 つの列になりました: お金を前払いした人を表す Paid By と、コストを負担する人を表す Responsible です。
台帳は 1 つでは不十分でした。改築、結婚式、レンタルユニット、旅行は「毎月の出費」ではありませんでした。これらは独自の寿命と分割ルールを持つプロジェクトでした。したがって、アカウントはプロジェクトごとに 1 つの台帳を取得しました。
通貨換算はすべてを複雑にします。この旅行では物事が面白くなりました。私たちは過ごしていました

カナダと米国では、取引はカナダドルと米ドルで到着しましたが、元の金額も一緒に保存された状態でブラジルレアルで保管する必要がありました。ブラジルのクレジット カードで使用される為替レートも購入日のレートと異なる場合があるため、各支払いがどこでどのように行われたかを追跡する必要がありました。
3 回目: 話すだけでは不十分な瞬間
1 か月ほど後、台帳には数百行が含まれていたので、私は明白な質問をしました。誰が誰に借りがあるのですか?
クロードはすべてを合計することができますが、LLM が数百行にわたって演算を行う場合は確認する必要があります。毎週再確認するのは、自分で行うよりも悪いです。
そこで、Claude の助けを借りて、 scripts/totals.py を作成しました。これは、表のヘッダーから台帳を解析し、関連する分割を適用して、各人の合計を出力できるスクリプトです。
その後、この変更で実際にどれくらい節約できたのか知りたいと思いました。私は、現在存在する台帳に対して両方のアプローチを実行しました。つまり、3 つの台帳にわたる 457 行が、1 つの結合された決済に解決されました。同じモデル、同じマシン、同じ質問。
CLI パスは、約 200 個の入力トークン (コマンドによって生成される 738 文字の概要) を送信しました。 LLM のみのパスは、台帳テーブル自体から 67,709 個のトークンすべてを送信しました。最後の行の範囲は 3 回の別々の実行から得られます。
そして、ここに私が予想していなかった部分があります。モデルはそれを正しく理解していました。
3 つの LLM のみの実行はすべて、CLI とセントが一致しました: R$17,930.52。
つまり、これはモデルが算術が苦手だという話ではありません。それは、およそ 73 ～ 82 倍の出力トークンと 25 倍の実測時間により、毎週、永遠に何かと照合しなければならない数値が得られるということです。
「過去 3 回は正しかった」というようなことは、決済システムを構築できる性質のものではありません。
今は AI にあいまいな作業を処理させます。つまり、

ぼやけた領収書や、MOMENTOS BUFFET E RECE がおそらく結婚式のケータリング業者であることを特定し、コーディングに正確な作業を与えます。
7 月 5 日に、スクリプトを実際のツール splttty (Ruby CLI) に置き換えました。 Python が間違っていたからではなく、Ruby が私のお気に入りの言語だからです。
問題は合計を計算することではなくなりました。それは、行がまだ手書きで (より正確には、Markdown を編集するクロードによって) 行われており、手書きの行が漂流しているということでした。日付が不一致になり、列の位置がずれて、単なるデータ入力の作業でトークンが焼き付けられます。
add は行を書き込みます。列の配置、カードのデフォルト、請求日のスタンプはすべて機械的なものであるため、現在はコード内に組み込まれています。 CLAUDE.md は、台帳テーブルを手動で編集しないことを明示的に示しています。
グループは文字列解析を置き換えました。どちらももはや魔法の言葉ではありません。これは、グローバルでは {Thiago: 50, Camila: 50} という名前付きグループですが、結婚式台帳では独自の Both グループを 65/35 で定義し、グローバル定義をオーバーライドします。 3 番目の参加者の追加は、パーサーの変更ではなく、構成の変更になりました。
コマンド spltty totals に移動されたスクリプトの合計を計算します。
Settle は、債務者を Paid By 、債権者を Responsible として、支払いを実際の台帳行として記録します。その支払いは、将来のすべての計算において負債を相殺します。
設定はメモ ファイルに保存されます。各台帳の YAML ヘッダーには、独自の設定とグループが含まれています。これらは config.json と双方向で同期され、最後に変更されたファイルが優先されます。どちら側でも編集できます。
テスト。そのうち86個あります。これにより、AI がツールを変更し続けても安全になります。
最後に、ループが閉じられました。取り込みワークフロー自体が /ingest スキルになりました。 「これはスクリーンショットです。それに対処してください」という手順は、私が持っていないものではなく、文書化された手順になりました。

o セッションごとに再説明する。
10 週間分のレシート、一番下に 1 つの数字。
試してみましょう: 最初に学ぶ価値のあるコマンド
# どのような台帳が存在しますか?これにより、設定とメモのヘッダーも調整されます。
分割リスト
# 経費を追加します。タイトルは位置的なものです。それ以外はすべてフラグです。
spltty add "Leite 1L" -l CASA -v 14.20 -p Thiago -mCash -r 両方
# カードに名前を一度付けると、支払者、責任者、請求日が自動的に入力されます。
splttty メソッドは「Mastercard 9759」を追加します -s mc-9759 -p Thiago -r Both -b last
spltty add "Extra" -l CASA -v 66.69 -m mc-9759 -y
# 分割がどのように機能するかを元帳ごとまたはグローバルに定義します。
spltty グループは両方を追加します -s Thiago:65,Camila:35 -l CASAMENTO
splttty グループは両方を追加します -s Thiago:50,Camila:50 --global
# 台帳ごとまたは台帳全体で、誰が誰に借りているのかを確認します。
分割された合計
splttty 合計 -- 結合
# この時点から借金を相殺するために支払いを記録します。
分割決済 CASA
-y フラグを指定すると、確認プロンプトがスキップされます。これは、Claude がすべての必須フィールドを収集した後に渡すものです。これを使用しないと、対話型の確認が表示されるため、最初のいくつかのエントリで使用する価値があります。
試してみて、どうなるか教えてください
splttty は単独で動作するようになりました。好みのインストール方法を選択してください。
醸造インストール thiagodiniz/spltty/spltty
gem インストール splttty
カール -fsSL https://raw.githubusercontent.com/thiagodiniz/spltty/main/install.sh |しー
次に、空のフォルダーを指定します。
mkdir ~/finances && cd ~/finances
分割インストール
このシステムは、私が当初望んでいたものであり、検査、編集、永久保存できるプレーンテキスト ファイルのフォルダーです。クロードはあいまいな入力を処理しますが、CLI は構造を強制し、計算を所有します。
その分業がプロジェクト全体になった。

## Original Extract

How I turned a folder of Markdown ledgers into a deterministic CLI that handles the math while Claude handles the fuzzy input.

I built Spltty, a Ruby CLI for tracking shared expenses, custom splits, and settlements using plain Markdown files. The article explains how it evolved from a Claude-managed folder into a CLI

Building a personal finance tracker with Claude · Thiago Diniz tdiniz Initiatives Career Thoughts ← Back to Thoughts
Building a personal finance tracker with Claude
My wife and I share money the way most couples do: with very poor documentation. Some things are hers, some are mine, and some are ours—but not necessarily 50/50. At one point, a wedding, an apartment renovation, and a six-week trip to Montreal were all happening at once. Records arrived as PDF statements, screenshots of credit card bills, photos of Pix receipts—Pix is Brazil’s instant payment system—or me saying out loud, “I paid R$700 for the dress fitting.”
Before we lived together, the only thing we tried to track was travel. We used split-the-bill apps, usually a different one for every trip. None of them stuck.
Bookkeeping was my job, and I always did it after the fact: reconstructing a week of taxis and dinners from memory and a pile of receipts, then entering everything days later into an app designed for expenses to be recorded as they happened. When we decided to move in together, the stakes got higher, so I moved the whole process to spreadsheets.
Spreadsheets rot too, for the same reason. The boring part isn’t the math—it’s transcribing forty rows from a credit card statement and deciding who each one belongs to.
So I started over with the dumbest possible system. Inspired by ledger-cli —plain text as the system of record, with no app and no database—I created a folder of Markdown files and a single text file telling Claude how to fill them in.
That system lasted a few weeks before it broke too. The transcription problem was solved; the math problem wasn’t. Asking an LLM to add up hundreds of rows means spending thousands of tokens on grade-school arithmetic and then checking every number it produces, every week. Money math needs to be deterministic and repeatable.
So I wrote a CLI to own the math. Ten weeks later, that folder of Markdown files has its own tool—and 86 tests.
In a rush? Go ahead and test it out: https://github.com/thiagodiniz/spltty
If you are more curious about how the project evolved, one iteration at a time. keep reading
First iteration: plain files and one instruction file
I added a CLAUDE.md containing the overall instructions, plus one file holding the expenses and a few notes.
The first expense table looked like this:
| Date | Title | Value (R$) | Responsible | Source |
|------------|-------|-----------:|-------------|--------|
That’s it. I dropped in a screenshot of a card bill, and Claude extracted 38 rows, confirming each one with me before writing it to the file.
At the time, I wasn’t building an app. I was documenting the shape of the data and the rules for filling it. Everything that came after was a refinement of that shape.
Second iteration: real money broke the schema
Two weeks in, the schema started to fall apart in ways I could only discover by using it.
“Responsible” was doing two jobs. When a row said Camila , did that mean she paid for it or that she owed it? Those are different facts, and settlements are impossible without both. So it became two columns: Paid By , for the person who fronted the money, and Responsible , for the person who bears the cost.
One ledger wasn’t enough. The renovation, the wedding, the rental unit, and the trip weren’t “monthly expenses.” They were projects with their own lifespans and split rules. So accounts/ got one ledger per project.
Currency conversion complicated everything. The trip was where things got interesting. We were spending in Canada and the US, so transactions arrived in CAD and USD but had to be stored in BRL, with the original amount preserved alongside them. The exchange rate used by a Brazilian credit card can also differ from the rate on the day of the purchase, so we had to track where and how each payment was made.
Third iteration: the moment talking wasn’t enough
After a month or so, the ledgers contained a few hundred rows, and I asked the obvious question: who owes whom?
Claude could add everything up, but an LLM doing arithmetic over hundreds of rows is something you have to check . Rechecking it every week is worse than doing it yourself.
So, with Claude’s help, I wrote scripts/totals.py : a script that could parse any ledger from its table header, apply the relevant splits, and print the totals for each person.
Later, I became curious about how much that change had actually saved. I ran both approaches against the ledgers as they exist today: 457 rows across three ledgers, resolved into one combined settlement. Same model, same machine, same question.
The CLI path sent roughly 200 input tokens: a 738-character summary produced by the command. The LLM-only path sent all 67,709 tokens from the ledger tables themselves. The ranges in the final row come from three separate runs.
And here’s the part I didn’t expect: the model got it right.
All three LLM-only runs matched the CLI to the cent: R$17,930.52.
So this isn’t a story about the model being bad at arithmetic. It’s that roughly 73–82 times more output tokens and 25 times more wall-clock time still buy me a number I have to verify against something, every week, forever.
“It was right the last three times” is not a property you can build a settlement system on.
Now I let AI handle the fuzzy work—reading a blurry receipt or identifying that MOMENTOS BUFFET E RECE is probably the wedding caterer—and give the exact work to code.
On July 5, I replaced the script with a real tool: spltty , a Ruby CLI. Not because Python was wrong, but because Ruby is my favourite language.
The problem was no longer computing the totals. It was that rows were still being written by hand—or, more accurately, by Claude editing Markdown—and hand-written rows drift. Dates become inconsistent, columns become misaligned, and tokens get burned on work that is really just data entry.
add writes the row. Column alignment, card defaults, and bill-date stamping are all mechanical, so they now live in code. CLAUDE.md explicitly says never to edit a ledger table by hand.
groups replaced string parsing. Both is no longer a magic word. It is a named group: {Thiago: 50, Camila: 50} globally, while the wedding ledger defines its own Both group at 65/35, which overrides the global definition. Adding a third participant is now a configuration change, not a parser change.
totals the script moved in to the commmand spltty totals .
settle records a payment as an actual ledger row, with the debtor as Paid By and the creditor as Responsible . That payment then offsets the debt in every future calculation.
Configuration lives in the notes files. Each ledger’s YAML header contains its own settings and groups. These are kept in two-way sync with config.json , with the most recently modified file taking precedence. I can edit either side.
Tests. There are 86 of them. This is what makes it safe to let an AI continue changing the tool.
Finally, the loop closed: the ingestion workflow itself became an /ingest skill. “Here’s a screenshot, deal with it” is now a documented procedure instead of something I have to re-explain in every session.
Ten weeks of receipts, one number at the bottom.
Try it: the commands worth learning first
# What ledgers exist? This also reconciles config with the notes headers.
spltty list
# Add an expense. The title is positional; everything else is a flag.
spltty add "Leite 1L" -l CASA -v 14.20 -p Thiago -m cash -r Both
# Name a card once, and Paid By, Responsible, and bill date are filled automatically.
spltty methods add "Mastercard 9759" -s mc-9759 -p Thiago -r Both -b last
spltty add "Extra" -l CASA -v 66.69 -m mc-9759 -y
# Define how a split works, either per ledger or globally.
spltty groups add Both -s Thiago:65,Camila:35 -l CASAMENTO
spltty groups add Both -s Thiago:50,Camila:50 --global
# See who owes whom, either per ledger or across all of them.
spltty totals
spltty totals --combined
# Record the payment so it offsets the debt from this point onward.
spltty settle CASA
The -y flag skips the confirmation prompt. That’s what Claude passes after it has collected every required field. Without it, you get an interactive confirmation, which is worth using for your first few entries.
Give it a try, and tell me how it goes
spltty now lives on its own. Pick whichever installation method you prefer:
brew install thiagodiniz/spltty/spltty
gem install spltty
curl -fsSL https://raw.githubusercontent.com/thiagodiniz/spltty/main/install.sh | sh
Then point it at an empty folder:
mkdir ~/finances && cd ~/finances
spltty install
The system is still what I wanted in the beginning: a folder of plain-text files that I can inspect, edit, and keep forever. Claude handles the ambiguous input, while the CLI enforces the structure and owns the math.
That division of labour turned out to be the whole project.
