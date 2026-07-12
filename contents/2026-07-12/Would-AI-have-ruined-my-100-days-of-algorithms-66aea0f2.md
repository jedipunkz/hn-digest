---
source: "https://danielsada.tech/blog/100-days-of-algorithms/"
hn_url: "https://news.ycombinator.com/item?id=48884624"
title: "Would AI have ruined my 100 days of algorithms?"
article_title: "Daniel Sada Caraveo – Would AI have ruined my 100 days of algorithms? – Software, Notes & Culture"
author: "dshacker"
captured_at: "2026-07-12T21:39:51Z"
capture_tool: "hn-digest"
hn_id: 48884624
score: 2
comments: 0
posted_at: "2026-07-12T20:47:09Z"
tags:
  - hacker-news
  - translated
---

# Would AI have ruined my 100 days of algorithms?

- HN: [48884624](https://news.ycombinator.com/item?id=48884624)
- Source: [danielsada.tech](https://danielsada.tech/blog/100-days-of-algorithms/)
- Score: 2
- Comments: 0
- Posted: 2026-07-12T20:47:09Z

## Translation

タイトル: AI は私の 100 日間のアルゴリズムを台無しにするでしょうか?
記事のタイトル: Daniel Sada Caraveo – AI は私の 100 日間のアルゴリズムを台無しにするでしょうか? – ソフトウェア、メモ、文化
説明: ソフトウェア開発エンジニア。何百万ものユーザーがいるサービスの同期テクノロジとストレージ ソリューションに取り組んでいます。私は人々がより生産的になれるよう支援します。

記事本文:
AI は私の 100 日間のアルゴリズムを台無しにしてしまうでしょうか?
8 年前、私は「100 日間のアルゴリズム」という挑戦を始めました。私の主な目標は、ロバート セジウィックが教えるプリンストン大学のアルゴリズム I と II のクラスを受講し、面白そうなアルゴリズム、またはそれらについて何かを学ぶことができそうなアルゴリズムを実装することで、CS のキャリアから学んだアルゴリズムを強化することでした。当時、私はreadmeに次のように書きました。
こんにちは！このアルゴリズムのほとんどについては以前にクラスで説明しましたが、自分で実装したわけではなく、どのような障害があるのか​​を確認しました。私は 1 日に 1 つか 2 つのアルゴリズムを実行するか、少なくとも 1 日に 1 つの問題を解決しようとしています。新しいアルゴリズムを交互に使用し、それらのアルゴリズムの問​​題に取り組みます。
このプロジェクトは面接の準備にも役立ちました。 Union Find が面接で人気の質問になるずっと前に、Sedgewick は、Union Find をいくつかの方法で実装する方法について見事なプレゼンテーションを行っており、アルゴリズムを教える方法についての黄金の例のように思えます。
8 年前、このコードのほとんどは LLM を使用せずに手作りされたことを考えると、欠陥はあるものの、私はそれを大切にしています。これらの問題の一部を解決するには、アルゴリズムを学習して使用するだけで、仕事後に時間がかかりました (1 時間の場合もありますが、それ以上の場合もあります)。大学の CS の授業では、実際に赤と黒の木のバランスを取る方法を学んだことはありませんでしたが、この授業でそれについて学ぶのは楽しいです。 (たとえ私の実装が少し壊れていたとしても、それは簡単なことではありません)
100 日のアルゴリズムを完了するのにわずか 8 年しかかかりませんでしたので、100 日チャレンジをもう一度行う気力があるかどうかはわかりません。
最終日には、GPT-5.6 Sol にプロジェクト全体をレビューしてもらうことにしました (下記を参照)。LLM があれば内容をもっと理解できるようになるのか、それともショートカットをするよう説得してくれるのかわかりません。

プロジェクト。それでも、私のコードに何らかのフィードバック メカニズムがあれば素晴らしいでしょう。
見つかったすべての問題を修正するのではなく、コードを当時のプログラマーの歴史的成果物として保存しています。ただし、実装したもの、検討しただけのもの、未完成のままのものを正確に区別できるように、README を更新します。
これは強力な学習ジャーナルですが、現時点では信頼できるアルゴリズム ライブラリではありません。文書化されたスイートでは、62 件の合格したテストが報告されていますが、ソース全体のカバレッジは約 74% で、いくつかの主要な実装が壊れているか未発見です。
明らかに動作しないファイル
- **最大フローはスタブです:** fordfulkerson.py:24-30 ;そのテストは True ( `test_flownetwork.py:17-20` ) のみをアサートします。
- **グラフ アルゴリズム:** BFS は深さ優先で動作し、最短以外のパス ( `bfs.py:26-32` ) を返すことができます。無向サイクル検出は永久に再帰します (cycle_detection.py:25)。有向サイクル検出は開始されず、逆の結果が報告されます ( `di_cycle_detection.py:19-28` )。
- **BST:** 不正なサイズ、天井トラバーサル、およびルートの削除 ( `bst.py:102-114,142-166` )。赤黒木サイズの簿記もありません。
- **MST/Kruskal:** mst.py:26 はインポート中に失敗します。 __len__ / __repr__ は何も返さず、Kruskal は存在しない g.v を参照します。
- **バイナリ I/O:** バイトおよび chr(bytes) の代わりに整数を書き込むとエラーが発生します。 isEmpty() は、ファイルが読み取り可能かどうか ( binarystreams.py ) のみを報告します。
- **明示的なスタブ:** BFS 順列 ( `permutations.py:9-10` ) およびフィボナッチ ( `fibonaccimatrixes.py:39-43` )。
- **その他の再現可能な欠陥:** [0,1,0] で選択ソートが失敗します。重複する文字列では 3 方向の基数ソートが失敗します。有向重み付きグラフ次数は、存在しない close() を呼び出します。
- **ML スクリプト:** 依存関係マニフェストがなく、廃止された TensorFlow に依存しています

1 API ( `tf.logging` 、 `tf.contrib` )。
**個人的な感想**
あなたの最も強力な仕事は、広範で一貫性のある毎日のログ、小さな構造への分解、および実装をテストする意欲です。 Union 検索、基本的な並べ替え、Prim、try、およびトラバーサルの基礎は、真の理解を示します。
ここでの主な教訓は、アルゴリズムの追加ではなく統合です。
1. README 内で未完了のトピックに正直にマークを付けます。
2. すべてのインライン テストを、検出されたスイートに移動します。
3. エッジケースを追加し、結果を Python リファレンス実装と比較します。
4. 修理
[切り捨てられた]
AI Lazyslop と個人の責任
レコードプレイヤーを持たずにレコードを持っている理由
私が何か新しいことを書いたときは、サインアップして詳細を聞いてください。
私の購読者リストに正常に参加しました。

## Original Extract

Software Development Engineer, I work on sync technologies and storage solutions for services that have millions of users. I empower people to be more productive.

Would AI have ruined my 100 days of algorithms?
Eight years ago, I started with a challenge called “ 100 Days of Algorithms ”. My main goal was to reinforce my algorithm learning from my CS career by taking Princeton’s Algorithms I and II classes taught by Robert Sedgewick, by implementing some algorithms that looked interesting or I could learn something about them. At the time, I wrote in the readme:
Hello! I’ve covered most of this algorithms in classes before, but I didn’t implement them myself, and see what roadblocks I’ve got. I’m trying to do an algorithm or two a day or at least try to solve a problem a day. I’ll be alternating between new algorithms and working in problems with those algorithms.
The project also helped me prep for interviews. Long before Union Find became a popular interview question, Sedgewick does a masterful presentation on how to implement Union Find in several ways that I find like the golden example on how to teach algorithms.
Given that 8 years ago, most of this code was handcrafted without the use of LLMs, I cherish it despite it’s flaws. It required time after work (sometimes 1 hour, but sometimes more!) of just learning and using algorithms to solve some of these problems. I never learned in college to see how to actually balance a red-black tree in my CS classes, but learning about it with this class is fun. (even if my implementation is a little bit broken, it’s not easy folks)
I’m not sure I have the drive to do 100-day challenge ever again, as it only took me 8 years to finish 100 days of algorithms.
For my final day, I decided to have GPT-5.6 Sol review my entire project, (see below) I’m not sure whether LLMs would have made me understand the material more or persuaded me to take shortcuts on the project. Some kind of feedback mechanism for my code would have been great nevertheless!
Rather than fixing every issue it found, I’m preserving the code as historical artifacts of the programmer I was then. I will, however, update the README so it accurately distinguishes what I implemented, what I only explored, and what remains unfinished.
This is a strong learning journal, but not currently a dependable algorithm library. The documented suite reports 62 passing tests, while whole-source coverage is roughly 74% and several major implementations are broken or undiscovered.
Files that plainly do not work
- **Max flow is a stub:** fordfulkerson.py:24-30 ; its test only asserts True ( `test_flownetwork.py:17-20` ).
- **Graph algorithms:** BFS behaves depth-first and can return non-shortest paths ( `bfs.py:26-32` ); undirected cycle detection recurses forever ( cycle_detection.py:25 ); directed cycle detection never starts and reports the opposite result ( `di_cycle_detection.py:19-28` ).
- **BST:** incorrect sizes, ceiling traversal, and root deletion ( `bst.py:102-114,142-166` ). Red-black-tree size bookkeeping is also absent.
- **MST/Kruskal:** mst.py:26 fails during import; __len__ / __repr__ return nothing, and Kruskal references nonexistent g.v .
- **Binary I/O:** writing integers instead of bytes and chr(bytes) raise errors; isEmpty() only reports whether the file is readable ( binarystreams.py ).
- **Explicit stubs:** BFS permutations ( `permutations.py:9-10` ) and Fibonacci ( `fibonaccimatrixes.py:39-43` ).
- **Other reproducible defects:** selection sort fails on [0,1,0] ; three-way radix sort fails on duplicate strings; directed weighted graph degree calls nonexistent close() .
- **ML script:** no dependency manifest and it relies on obsolete TensorFlow 1 APIs ( `tf.logging` , `tf.contrib` ).
**Personal reflection**
Your strongest work is the breadth, consistent daily log, decomposition into small structures, and willingness to test implementations. Union-find, basic sorting, Prim, tries, and traversal foundations show real understanding.
The main lesson is now consolidation rather than adding algorithms:
1. Mark unfinished topics honestly in the README.
2. Move every inline test into the discovered suite.
3. Add edge cases and compare results against Python reference implementations.
4. Repair
[truncated]
AI Lazyslop, and Personal Responsibility
Why I have vinyls with no vinyl player
When I write something new, sign up to hear more.
You have successfully joined my subscriber list.
