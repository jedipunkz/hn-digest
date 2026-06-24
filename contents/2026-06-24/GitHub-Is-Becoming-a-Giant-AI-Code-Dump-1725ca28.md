---
source: "https://maref.cc/en/blog/vibe-coding-crisis/"
hn_url: "https://news.ycombinator.com/item?id=48656807"
title: "GitHub Is Becoming a Giant AI Code Dump"
article_title: "MAREF — GitHub Is Becoming a Giant AI Code Dump"
author: "Athena-maref"
captured_at: "2026-06-24T09:03:25Z"
capture_tool: "hn-digest"
hn_id: 48656807
score: 21
comments: 22
posted_at: "2026-06-24T08:21:59Z"
tags:
  - hacker-news
  - translated
---

# GitHub Is Becoming a Giant AI Code Dump

- HN: [48656807](https://news.ycombinator.com/item?id=48656807)
- Source: [maref.cc](https://maref.cc/en/blog/vibe-coding-crisis/)
- Score: 21
- Comments: 22
- Posted: 2026-06-24T08:21:59Z

## Translation

タイトル: GitHub は巨大な AI コードダンプになりつつある
記事のタイトル: MAREF — GitHub は巨大な AI コード ダンプになりつつある
説明: 6 億 3,000 万のリポジトリ、新しいコードの半分は AI によって作成されましたが、信頼性は 77% から 60% に低下しました。 CMUは600万人の偽スターを発見した。 CodeRabbit: AI コードには人間が書いたコードよりも 1.7 倍深刻な問題があります。開発者の 63% は、AI コードを修正するのは、ゼロから記述するよりも時間がかかると述べています。そして最もワイルドな部分 — AI ユーザーは 19% 遅いですが、
[切り捨てられた]

記事本文:
マレフ
機能ドキュメント 仕組み 無料コミュニティ統合ブログ FAQ GitHub について 中文 機能ドキュメント 仕組み 無料コミュニティ統合ブログ FAQ GitHub について 中文 2026-06-13 GitHub は巨大な AI コード ダンプになりつつある
GitHub には 6 億 3,000 万のリポジトリがあります。すべての新しいコードのほぼ半分は AI によって書かれています。
生産性が飛躍的に向上したように思えますよね?
しかし、同じレポートには別の数字もあります。AI コードに対する開発者の信頼は 77% から 60% に低下しました。利用する人が増えています。それを信じる人は少なくなります。
あなたが目にするプロジェクトの半分は偽物です
カーネギーメロン大学が GitHub 上で 600 万人の偽スターを発見しました。セキュリティ会社 Socket は、詐欺に直接関係する 370,000 件の「フィックススター」を発見しました。何千人ものスターが参加するあのバイブコーディングプロジェクト？おそらく半分は捏造だと思います。
それは怖い部分ではありません。
CodeRabbit は 470 の PR をスキャンしました。 AI が書いたコードには、人間が書いたコードよりも 1.7 倍重大な問題がありました。 AI コードの 45% には、OWASP トップ 10 の脆弱性が含まれています。開発者の 63% は、AI コードを修正するのは、ゼロから記述するよりも時間がかかると述べています。
しかし、ここで最も衝撃的な発見があります。
遅くなっているけど、何もわかっていない
メーターはランダム化対照試験を実施した。結果: AI ユーザーは実際には 19% 遅くなりましたが、彼らは自分たちが 20% 速いと思っていました。
参加者にデータを見せても、やはり自分たちの方が速いと主張しました。
それが本当の恐怖なのです。だんだん遅くなってきて、何もわかりません。
オープンソースはゴミ PR に埋もれている
主要なプロジェクトがどのようなことを行っているかを見てみましょう。
curl — 7 年間のバグ報奨金プログラムを終了しました。 AI が生成したバグ レポート: 本物は 5% のみでした。残りはノイズでした。
Ghost — AI が提出したコードを完全に禁止しました。
Tailscale — さらに進んで、すべての外部 PR を閉じました。 AIかどうかはもう関係ありません。
GitHub 自体 — メンテナが 1 つの CL を実行できるように「PR キル スイッチ」を構築

外部送信を無効にします。
オープンソースはハッカーによって破壊されていません。バイブコーダーが AI で生成したゴミ PR によって溺死しつつあります。
この一連の流れの中で最も皮肉な部分は次のとおりです。
AI にコードを書くように依頼する → AI が GitHub を参照して検索する → 他の AI が書いたコードを見つける → ガベージイン、ガベージアウト。何が出ると思いますか？
GitHub は以前はコード リポジトリでした。それは巨大な AI コード ダンプと化しつつあります。
問題はAIではありません。それはガバナンスです。
AI がコードを書くことが問題なのではありません。問題は、AI の動作を誰も監査していないことです。
バイブコーディングの中核となる前提は、「AI が生成するものはすべて受け入れる、レビューしない、変更しない」です。これは TODO アプリで機能します。製品コードの場合、時限爆弾を仕掛けることになります。
しかし、人間に AI コードを 1 行ずつレビューしてもらうのも現実的ではありません。開発者の 63% が、修正のほうが執筆よりも時間がかかると回答している場合、レビューのほうがさらに時間がかかります。
本当の解決策は「AIの使用をやめる」とか「すべてを手動で見直す」ということではありません。これは、AI コードと本番環境の間の自動ガバナンスです。
すべてのツール呼び出しが監査され、すべてのファイル変更、すべての API 呼び出しが暗号署名され、記録されます。
危険な行為はブロックされます — データベースを削除しますか?生産を変更しますか？最初に 4 レベルのデシジョン ツリーを通過します。 97% は自動解決され、3% は人間によるレビューにエスカレーションされました。
セキュリティ ポリシーは進化します。あらゆる誤検知やバイパスの試みがガバナンス エンジンにフィードバックされます。
正式な検証 — 「これが安全であることを願っています」ではありません。数学的に証明可能な安全性への収束。
ガバナンスはスピードバンプではありません。そうすることで安全に速く走ることができるのです。
管理されていない AI コードは高速化されません。ただクラッシュが早くなるだけです。
AI の投稿を禁止するプロジェクト (curl、Ghost、Tailscale) は反 AI ではありません。彼らは同じことを言っています：「AIの支援を歓迎します。AIなしのゴミは受け入れません」

品質管理。」
MAREF はその品質管理層です。あなたとAIの間ではありません。ガベージ コードと運用環境の間。
コードをバイブすることもできます。 AI にコードを作成させることもできます。しかし、データベースが削除される前に、誰かがそれを停止します。
遅くなってきていますか？いいえ、実際に出荷するものをついに確認しました。
📊 出典: GitHub Octoverse Report、カーネギーメロンの偽スター調査、Socket セキュリティレポート、CodeRabbit 470 PR 分析、Meter RCT トライアル、curl/BT/Ghost/Tailscale 公式発表。 MAREF は、オープンソースのエージェント ガバナンス オペレーティング システムです。 5 分以内に始めましょう。
MAREF — マルチエージェント再帰的進化フレームワーク

## Original Extract

630M repos, half of new code is AI-written, but trust dropped from 77% to 60%. CMU found 6M fake stars. CodeRabbit: AI code has 1.7x more serious issues than human-written. 63% of developers say fixing AI code takes longer than writing from scratch. And the wildest part — AI users are 19% slower but
[truncated]

MAREF
Features Docs How It Works Free Community Integrations Blog FAQ About GitHub 中文 Features Docs How It Works Free Community Integrations Blog FAQ About GitHub 中文 2026-06-13 GitHub Is Becoming a Giant AI Code Dump
GitHub has 630 million repos. Nearly half of all new code is written by AI.
Sounds like a productivity explosion, right?
But there's another number in that same report: developer trust in AI code dropped from 77% to 60% . More people are using it. Fewer people believe in it.
Half the projects you see are fake
Carnegie Mellon University found 6 million fake stars on GitHub. Security firm Socket uncovered 370,000 "fix stars" directly tied to scams. That vibe coding project with thousands of stars? Probably half are fabricated.
That's not even the scary part.
CodeRabbit scanned 470 PRs. AI-written code had 1.7x more critical issues than human-written code. 45% of AI code ships with OWASP Top 10 vulnerabilities. 63% of developers say fixing AI code takes longer than writing it from scratch.
But here's the most mind-bending finding.
You're getting slower, but you have no idea
Meter ran a randomized controlled trial. The result: AI users were actually 19% slower , but they thought they were 20% faster .
When they showed the participants the data, they still insisted they were faster.
That's the real horror. You're getting slower, and you have no idea.
Open source is drowning in garbage PRs
Look at what major projects are doing:
curl — shut down its 7-year bug bounty program. AI-generated bug reports: only 5% were real. The rest was noise.
Ghost — outright banned AI-submitted code.
Tailscale — went further: closed all external PRs. AI or not, doesn't matter anymore.
GitHub itself — building a "PR kill switch" so maintainers can one-click disable external submissions.
Open source isn't being destroyed by hackers. It's being drowned to death by garbage AI-generated PRs from vibe coders.
Here's the most ironic part of the whole chain:
You ask AI to write code → AI searches GitHub for references → It finds code written by other AIs → Garbage in, garbage out. What do you think comes out?
GitHub used to be a code repository. It's turning into a giant AI code dump.
The problem isn't AI. It's governance.
AI writing code isn't the problem. The problem is: nobody is auditing what AI does.
The core premise of vibe coding is "accept everything AI generates — don't review, don't modify." That works for a TODO app. For production code, it's planting time bombs.
But asking humans to line-by-line review AI code isn't realistic either. If 63% of devs say fixing is slower than writing, reviewing is even slower.
The real solution isn't "stop using AI" or "review everything manually." It's automated governance between AI code and production :
Every tool call is audited — every file change, every API call, cryptographically signed and recorded.
Dangerous actions are blocked — delete database? modify production? First pass through a 4-level decision tree. 97% auto-resolved, 3% escalated to human review.
Security policies evolve — every false positive, every bypass attempt feeds back into the governance engine.
Formal verification — not "we hope this is safe." Mathematically provable convergence toward safety.
Governance isn't a speed bump. It's what lets you go fast safely.
Un-governed AI code doesn't go faster. It just crashes faster.
The projects banning AI submissions — curl, Ghost, Tailscale — aren't anti-AI. They're saying the same thing: "We welcome AI assistance. We won't accept garbage without quality control."
MAREF is that quality control layer. Not between you and AI. Between garbage code and your production environment.
You can still vibe code. You can still have AI write your code. But before it deletes your database, someone will stop it.
You're getting slower? No. You're finally seeing what you're actually shipping.
📊 Sources: GitHub Octoverse Report, Carnegie Mellon fake star study, Socket security report, CodeRabbit 470 PR analysis, Meter RCT trial, curl/BT/Ghost/Tailscale official announcements. MAREF is an open-source agent governance operating system. Get started in 5 minutes .
MAREF — Multi-Agent Recursive Evolution Framework
