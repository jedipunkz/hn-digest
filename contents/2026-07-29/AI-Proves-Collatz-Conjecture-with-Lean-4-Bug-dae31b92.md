---
source: "https://twitter.com/gro_tsen/status/2082483878480977959"
hn_url: "https://news.ycombinator.com/item?id=49101465"
title: "AI \"Proves\" Collatz Conjecture with Lean 4 Bug"
article_title: "Gro-Tsen on X: \"So, someone came up with an AI-generated formal proof, in Lean, of a solution to the Collatz problem, and it turned out that the “proof” was merely exploiting a bug in the Lean kernel (allowing you to prove anything).\" / X"
author: "pfdietz"
captured_at: "2026-07-29T18:57:25Z"
capture_tool: "hn-digest"
hn_id: 49101465
score: 3
comments: 0
posted_at: "2026-07-29T18:50:47Z"
tags:
  - hacker-news
  - translated
---

# AI "Proves" Collatz Conjecture with Lean 4 Bug

- HN: [49101465](https://news.ycombinator.com/item?id=49101465)
- Source: [twitter.com](https://twitter.com/gro_tsen/status/2082483878480977959)
- Score: 3
- Comments: 0
- Posted: 2026-07-29T18:50:47Z

## Translation

タイトル: AI がリーン 4 バグでコラッツ予想を「証明」
記事のタイトル: Gro-Tsen on X: 「誰かが、AI が生成したコラッツ問題の解決策の正式な証明をリーンで考え出しました。そして、その「証明」はリーン カーネルのバグを利用しているだけであることが判明しました (何でも証明できるようになります)。」 /X
説明: そこで、誰かが、Collatz 問題の解決策について、AI が生成した正式な証明を Lean で考え出しました。そして、その「証明」は単に Lean カーネルのバグを悪用しているだけであることが判明しました (何でも証明できるようになります)。

記事本文:
Gro-Tsen @gro_tsen そこで、誰かが AI が生成した Collatz 問題の解決策の正式な証明を Lean で考え出しましたが、その「証明」は単に Lean カーネルのバグを悪用しているだけであることが判明しました (何でも証明できるようになります)。 3:10 PM · 2026 年 7 月 29 日 55.8K 再生数 46 72 1.3K 155
Gro-Tsen @gro_tsen 3h 詳細と説明は次のとおりです: infosec.exchange abadidea (@ [email protected] ) さて、これまでに起こった最も AI の新しい候補者がいます 1) 7 月 25 日: 誰かが LLM をいじり、実際に定理で検証するコラッツ予想の証明を投稿しました... 3 3 71 6.1K Gro-Tsen @gro_tsen 3h 誰かが、リーンで形式的証明を要求することが、数学的証明における AI のスロップと幻覚に対する最終的な解決策であると主張する場合、このことを念頭に置いてください。 2 6 141 6.1K Gro-Tsen @gro_tsen 3h (もちろん、公式に証明された定理が非公式に理解されている定理と一致しない可能性があるという事実のような、より頻繁に起こる問題は他にもあります。または、より基本的には、証明が人間には読めず役に立たなくなる可能性があるということです。) 5 2 60 5.8K Gro-Tsen @gro_tsen 1h * これによって上記のいずれかが変わるわけではありませんが、私は知らされました。証拠を投稿した人は、これがリーンカーネルの健全性に関するバグであり、コラッツの問題の解決策として真剣に受け止められることを意図したものではないことを実際に認識していました。ジェイソン・ルート @JasonRute 2h リーン健全性バグを実証する楽しい方法としてコラッツを使っているだけだと思います @gro_tsen定理証明コミュニティにはこれに関する歴史があります。作者はそれが健全性に関するバグであり、意図的なものであることを知っていました。 1 33 4.6K
Antoine Ducros @antoineducros 3h しかし、このような事件は、リーンチェックされた多くの証拠に疑惑を投げかけませんか?検出されにくい形でそれがすでに起こっていないことをどうやって確信できるでしょうか? 2

54 3.8K
今すぐサインアップして、自分だけのタイムラインを手に入れましょう!
Google でサインアップ Apple でサインアップ アカウントを作成 サインアップすると、Cookie の使用を含むサービス利用規約とプライバシー ポリシーに同意したことになります。

## Original Extract

So, someone came up with an AI-generated formal proof, in Lean, of a solution to the Collatz problem, and it turned out that the “proof” was merely exploiting a bug in the Lean kernel (allowing you to prove anything).

Gro-Tsen @gro_tsen So, someone came up with an AI-generated formal proof, in Lean, of a solution to the Collatz problem, and it turned out that the “proof” was merely exploiting a bug in the Lean kernel (allowing you to prove anything). 3:10 PM · Jul 29, 2026 55.8K Views 46 72 1.3K 155
Gro-Tsen @gro_tsen 3h More details and explanations at: infosec.exchange abadidea (@ [email protected] ) Okay, we have a new contender for Most AI Thing to Ever Happen 1) July 25th: someone messes around with an LLM and posts a proof of the Collatz conjecture that does, in fact, verify in the theorem... 3 3 71 6.1K Gro-Tsen @gro_tsen 3h Keep this in mind when someone claims that requiring formal proofs in Lean is the end-all solution to AI slop and hallucination in mathematical proofs. 2 6 141 6.1K Gro-Tsen @gro_tsen 3h (There are other — more frequent — problems, of course, like the fact that the formally proven theorem might not match the informally understood one. Or, more basically, that the proof might be unreadable for humans, making it useless.) 5 2 60 5.8K Gro-Tsen @gro_tsen 1h * Not that this changes any of the above, but I am informed that the person posting the proof was actually aware that this was a Lean kernel soundness bug, and it was not intended to be taken seriously as a solution to Collatz's problem. Jason Rute @JasonRute 2h Replying to @gro_tsen I think they just use collatz as a fun way to demonstrate the lean soundness bug. There is a history of this in the theorem proving community. The author knew it was a soundness bug and it was intentional. 1 33 4.6K
Antoine Ducros @antoineducros 3h But doesn’t such an incident shed suspicion on plenty of Lean-checked proofs? How can one be sure that it did not already happen in a less detectable way? 2 54 3.8K
Sign up now to get your own personalized timeline!
Sign up with Google Sign up with Apple Create account By signing up, you agree to the Terms of Service and Privacy Policy , including Cookie Use.
