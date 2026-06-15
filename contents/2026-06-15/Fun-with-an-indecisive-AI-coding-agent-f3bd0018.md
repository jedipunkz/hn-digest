---
source: "https://benhoyt.com/writings/indecisive-ai-agent/"
hn_url: "https://news.ycombinator.com/item?id=48542935"
title: "Fun with an indecisive AI coding agent"
article_title: "Fun with an indecisive AI coding agent"
author: "azhenley"
captured_at: "2026-06-15T16:42:46Z"
capture_tool: "hn-digest"
hn_id: 48542935
score: 2
comments: 0
posted_at: "2026-06-15T15:40:35Z"
tags:
  - hacker-news
  - translated
---

# Fun with an indecisive AI coding agent

- HN: [48542935](https://news.ycombinator.com/item?id=48542935)
- Source: [benhoyt.com](https://benhoyt.com/writings/indecisive-ai-agent/)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T15:40:35Z

## Translation

タイトル: 優柔不断なAIコーディングエージェントと楽しむ
説明: ここでは、一見優柔不断な AI コーディング エージェントである Claude Opus 4.6 を楽しんでいます。

記事本文:
優柔不断な AI コーディング エージェントを楽しむ
AI エージェントを使用して GoAWK の重大なバグを修正することで、ある程度の成功を収めてきました。しかし、時には非常に優柔不断になることもあります（つい先日、ミッチェル・ハシモト氏がツイートでこのことを指摘しました）。
私は次の質問をしました (私は GoLand IDE の GitHub Copilot プラグインで Claude の Opus 4.6 を使用しています)。
GoAWK は、Awk プログラム BEGIN { a["x"]=1; に対して、「x 1\n」の代わりに「0\n0\n」を出力します。 for (NR in a) print NR, a[NR] } 。なぜ？
数段落以内に、私よりも早く問題を診断してくれました。これらの特殊変数をネイティブの Go int として保存していることがわかり、文字列表現が失われます。
それからそれを修正しようとし始めました。数分の間に、私はますます面白くなり、バグを修正するための 7 つのユニークなオプションが表示され、それらの間で 25 回以上フリップフロップしました。
出てきたオプションは次のとおりです。
オプション A: 特殊変数の文字列表現を保持する
オプション B: 特殊変数を値型として保存する
オプション C: 特殊変数が文字列に設定されている場合に文字列のオーバーライドを保存する
オプション D: ForIn オペコードを具体的に修正する
オプション E: 元の値をサイドフィールドに保存
オプション F: lineNum (NR) と fileLineNum (FNR) を値に変更するだけです
オプション G: 値タイプの特別なオーバーライド マップを追加します
数秒ごとに、「実は、最も簡単な解決策は…」または「しかし、いいえ、本当の問題は…」と表示されます。私がキャンセルするまでそれは続きました。
名誉のために言っておきますが、同社は最も頻繁に (26 回中 11 回) オプション B を提案しており、それがほぼ確実に適切な修正です。今回のPRでやってみました。
参考までに、チャット出力のバージョンを以下に示します。別のオプションを提案しているときだけを示すために大幅に切り取られています。これらを次のように太字でマークしました: Option A 。
見ている

このバグ、問題は ForIn オペコード ハンドラーにあります…
しかし、待ってください。実際の ForIn 実装をもっと注意深く見てみましょう…
今ではバグがはっきりとわかります…
これでバグが確認されました。全体像は次のとおりです: …
根本的な原因は、NR、NF、FNR などの特殊変数がネイティブ Go int 型として格納され、本質的に文字列表現が失われることです。 …
修正は setSpecial と getSpecial で行う必要があります。これらの数値スペシャルの文字列表現を保存する必要があります。 … オプション A
ただし、より単純で対象を絞ったアプローチでは、問題は特に特殊な変数を使用した for-in に関するものです。しかし実際には、NR = “x” のような代入でも同じ問題が発生し、文字列が失われます。最善の解決策を考えさせてください。
最もクリーンな修正は、これらの特殊変数を値の型として保存することです…しかし、これは大規模なリファクタリングです。オプション B
より外科的な修正: 文字列を介して設定される数値特殊文字の「文字列オーバーライド」を保存します。しかし、それは危険です。オプション C
実は、これについては別のことを考えてみましょう。 …
正しい修正は、値を値型として保存し、内部で必要な場合にのみ int を抽出することです。 … オプション B
実際、現在のアーキテクチャを維持する最も簡単な修正は、特に for-in ループについてはアプローチを変更することです。オプション D
しかし、いいえ、本当の問題はもっと一般的なものです。数値特殊文字への文字列の代入では文字列が保持される必要があります。
最小限だが正しい修正について考えさせてください。最も的を絞ったアプローチ: 数値特殊変数のストレージを int から value に変更し、それらを読み取り/書き込みするいくつかの場所を更新します。 … オプション B
これは私がやりたい以上に大規模なリファクタリングです。もっと簡単な方法を考えてみましょう。 …
実際、最も実用的で最小限の修正は、これを ForIn レベルで処理することです。

… オプション D
実際、正しい修正は、これらの数値特殊値を値型として格納するように変更することだと思います。 … オプション B
これらすべてを変更するのは危険です。もっと簡単な方法を考えてみましょう。
実際、最も単純な正しい修正は次のとおりです。setSpecial で、数値スペシャルの場合は、元の値をサイド フィールドに保存し、一致したときに getSpecial がその値を返せるようにします。 … オプション E
わかりました、これについてもっと慎重に考えてみましょう。 …
これは複雑になってきています。もっと単純なアプローチをとってみましょう。lineNum と fileLineNum は、for-in のターゲットになる可能性が最も高いため、値型に変更するだけです。 … オプション F
待ってください。実際には、もっと簡単なアプローチがあります。 …
わかりました。実際的なアプローチをとらせてください。この特定のバグに対する最もクリーンな最小限の修正は、数値特殊変数 (lineNum、fileLineNum、matchLength、matchStart、argc、numFields) を値型として格納するように変更することです。 … オプション B
実際には、お待ちください。numFields は内部で int として頻繁に使用されます。別のアプローチを試してみましょう。 …
実際、最も単純で正しいアプローチは次のとおりです。数値特殊変数ごとに、int (内部使用用) と値 (AWK レベルの表現用) の両方を格納します。 … オプション E
しかし、それでも多くの変化があります。もっとシンプルに考えてみます…
特定のケースを処理する最も単純で正しい修正: setSpecial で、数値特殊の場合、値が文字列 (意味のある数値に変換できないか、typeStr である) の場合、値を直接保存することで文字列を保持します。 … オプション B
OK、最もクリーンな最小限の修正は、specialOverrides map[int]value フィールドを interp 構造体に追加することだと思います。 … オプション G
実際のところ、ちょっと待ってください。それは非常に複雑です。具体的なケースをもう一度見てみましょう。 …
先に進み、適切な修正をさせてください。 int の代わりに値の型を格納するように数値スペシャルを変更します。 … オプション

B
別の、もっと単純なアプローチを考えてみましょう。 SpecialValues マップを追加することで、setSpecial と getSpecial が値を適切に処理できるようにします… オプション G
実際、内部に手を加える必要のない最も簡単なアプローチは、ForIn ハンドラーを変更するだけです。 …しかし問題は、getSpecial が常に int から再構築することです…
OK、適切な修正をしましょう。 「特別な値を上書きする」マップ アプローチを使用します。 … オプション G
さて、実際にアプローチを再考させてください。 … これに一致する最もクリーンなアプローチ: 数値特殊文字を int から value に変更するだけです。 … オプション B
待ってください、実際には、さらに現実的なことを言わせてください。アプローチ: int フィールドは内部使用のために保持しますが、AWK レベルの値を追跡する並列値フィールドを追加します。 … オプション E
いいえ、それは二重簿記です。ここで少しだけ、関連するフィールドを値タイプに変更してみましょう。 … オプション B
うーん、実際のところ、一歩下がって最も簡単な正しい修正とは何かを考えてみましょう。最も最小限の修正: for-in の場合のみを処理します。 … オプション D
いや、待ってください。問題は、setSpecial が int に切り捨てられるため、str 値が失われることです。さて、これが私の計画です。 lineNum と fileLineNum (for-in ループの最も可能性の高い候補、NR と FNR) だけを int から value に変更します。 … オプション F
実際には、すべてを適切に実行させてください。私の計画は次のとおりです: … オプション B
ここでは、for-in (NR および FNR) で使用される可能性が最も高い変数である lineNum と fileLineNum だけに焦点を当てます。オプション F
実際、ユーザーのバグレポートは特に NR に関するものでした。すべての数値特殊項目を修正しますが、頻繁に使用される数値特殊記号については、より単純なアプローチを採用します。オプション B (多かれ少なかれ)
実際のところ、すべてのケースを解決できる最も単純なアプローチを採用しましょう。map[int]value フィールドを追加します。

ユーザー設定の特別なオーバーライドを保存します。 … オプション G
興味深いですね！すでにspecialOverridesフィールドがあり、それに対する呼び出しを削除します。これらは私が変更される前から存在していました。 …
今、それは非常に混乱し始めました。コードを編集して、specialOverrides マップを追加していましたが、そのことを忘れていました。この時点では、キャンセルしただけです。
「実は」は19個あります。これらは素晴らしいマシンですが、これほど優柔不断な友人を持ちたくないです。
より一般的には、大規模な言語モデルとエージェントがいつ停止すべきかをどのように認識するのかに興味があります。トークン プレディクターがトークンを予測し続けることは、私にとっては驚くことではありません…

## Original Extract

In which I have some fun with Claude Opus 4.6, a seemingly indecisive AI coding agent.

Fun with an indecisive AI coding agent
I’ve been having moderate success using an AI agent to help me fix non-trival bugs in GoAWK . But it’s sometimes very indecisive (Mitchell Hashimoto noted this in a tweet just the other day).
I asked it the following question (I’m using Claude’s Opus 4.6 with the GitHub Copilot plugin in the GoLand IDE):
GoAWK prints out "0\n0\n" instead of "x 1\n" for the Awk program BEGIN { a["x"]=1; for (NR in a) print NR, a[NR] } . Why?
It diagnosed the problem within a few paragraphs, and quicker than I could. It showed that I’m storing these special variables as native Go int s, which lose any string representation.
Then it started trying to fix it. Over the course of several minutes, with me becoming more and more amused, it gave me 7 unique options for fixing the bug, and flip-flopped between them no less than 25 times.
Here are the options it came up:
Option A: preserve string representation for special variables
Option B: store special variables as value type
Option C: store string override when special variables are set to string
Option D: fix ForIn opcode specifically
Option E: store original value in side field
Option F: just change lineNum (NR) and fileLineNum (FNR) to value
Option G: add special overrides map of value type
Every few seconds it would say “Actually, the simplest fix…” or “But no, the real issue is…”. It went on and on until I cancelled it.
To its credit, it proposed Option B most often (11 out of 26 times), and that’s almost certainly the proper fix. I’ve done it in this PR.
For your interest, I’m including a version of the chat output below. It’s heavily clipped to just show when it was suggesting another option. I’ve marked those in bold, like this: Option A .
Looking at this bug, the issue is in the ForIn opcode handler…
But wait, let me look more carefully at the actual ForIn implementation …
Now I can see the bug clearly …
This confirms the bug. Here’s the complete picture: …
The root cause is that special variables like NR, NF, FNR, etc. are stored as native Go int types, which inherently lose the string representation. …
The fix needs to happen in setSpecial and getSpecial — we need to preserve the string representation for these numeric specials. … Option A
However, a simpler and more targeted approach: the issue is specifically about for-in with special variables. But really the same problem would occur with any assignment like NR = “x” — it would lose the string. Let me think about the best fix.
The cleanest fix is to store these special variables as value types … but that’s a large refactor. Option B
A more surgical fix: store a “string override” for numeric specials when they’re set via string. But that’s hacky. Option C
Actually, let me think about this differently. …
The right fix is to store the value as a value type and only extract the int when needed internally. … Option B
Actually, the simplest fix that preserves the current architecture: for the for-in loop specifically, we could change the approach. Option D
But no, the real issue is more general — any string assignment to a numeric special should preserve the string.
Let me think about a minimal but correct fix. The most targeted approach: change the numeric special variable storage from int to value, and update the few places that read/write them. … Option B
This is a bigger refactor than I want to do. Let me think of a simpler approach. …
Actually, the most pragmatic and minimal fix is to handle this at the ForIn level. … Option D
Actually, I think the right fix IS to change these numeric specials to store as value type. … Option B
Changing all of these is risky. Let me think of an even simpler approach.
Actually, the simplest correct fix: In setSpecial, for the numeric specials, store the original value in a side field so that getSpecial can return it when it matches. … Option E
OK, let me think about this more carefully. …
This is getting complex. Let me take a simpler approach: just change lineNum and fileLineNum to value type, since those are the most likely targets of for-in. … Option F
Wait, there’s actually a much simpler approach. …
OK let me take the pragmatic approach. The cleanest minimal fix for this specific bug: change the numeric special variables (lineNum, fileLineNum, matchLength, matchStart, argc, numFields) to be stored as value types. … Option B
Actually wait, numFields is heavily used internally as an int. Let me try a different approach. …
Actually, here’s the simplest correct approach: For each numeric special variable, store BOTH the int (for internal use) AND a value (for the AWK-level representation). … Option E
But that’s still a lot of changes. Let me think even simpler…
The simplest correct fix that handles the specific case: in setSpecial, for numeric specials, if the value is a string (not convertible to a meaningful number or is a typeStr), preserve the string by storing a value directly. … Option B
OK, I think the cleanest minimal fix is: add a specialOverrides map[int]value field to the interp struct. … Option G
Actually wait - that’s overcomplicating it. Let me re-examine the specific case. …
Let me just go ahead and do the proper fix. I’ll change the numeric specials to store value types instead of int. … Option B
Let me take a different, much simpler approach. I’ll just make setSpecial and getSpecial properly handle values by adding a specialValues map … Option G
Actually, the simplest approach that doesn’t require touching internals: just change the ForIn handler. … But the problem is getSpecial always reconstructs from int…
OK, let me just do the proper fix. I’ll use a “special values override” map approach: … Option G
OK actually let me reconsider the approach. … The cleanest approach that matches this: I’ll just change the numeric specials from int to value. … Option B
Wait, actually, let me just be even more pragmatic. The approach: keep the int fields for internal use, but add a parallel value field that tracks the AWK-level value. … Option E
Actually no, that’s double-bookkeeping. Let me just bite the bullet and change the relevant fields to value type. … Option B
Hmm, actually, let me step back and think about what the simplest CORRECT fix is. The most minimal fix: only handle the for-in case. … Option D
No wait, the issue is setSpecial truncates to int, so the str value is lost regardless. OK here’s my plan. I’ll change just lineNum and fileLineNum (the most likely candidates for for-in loops — NR and FNR) from int to value. … Option F
Actually, let me just do ALL of them properly. Here’s my plan: … Option B
Let me focus on just lineNum and fileLineNum, which are the variables most likely to be used in for-in (NR and FNR). Option F
Actually, the user’s bug report is specifically about NR. Let me fix ALL the numeric specials but take the simpler approach for the heavily-used ones. Option B (more or less)
Actually, you know what, let me just take the simplest possible approach that fixes all cases: add a map[int]value field to store user-set special overrides. … Option G
Interesting! There’s already a specialOverrides field and delete calls on it. These existed before my changes. …
Now it started to get very confused. It had edited the code to add a specialOverrides map, but then forgot that it had done so. At this point, I just cancelled it.
There are 19 “Actually”s. These are amazing machines, but I’d hate to have a friend who was this indecisive!
More generally, I’d be interested to know how large language models and agents know when to stop. It’s not surprising to me that a token predictor keeps predicting tokens…
