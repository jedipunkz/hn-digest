---
source: "https://bertrandmeyer.com/2026/07/21/questions-and-answers-on-ai-and-verification-a-follow-up-to-my-may-acm-tech-talk/"
hn_url: "https://news.ycombinator.com/item?id=49085108"
title: "Questions and answers on AI and verification: follow-up to my May ACM Tech Talk"
article_title: "Questions and answers on AI and verification: a follow-up to my May ACM Tech Talk - Bertrand Meyer's technology+ blog"
author: "rbanffy"
captured_at: "2026-07-28T15:12:18Z"
capture_tool: "hn-digest"
hn_id: 49085108
score: 1
comments: 0
posted_at: "2026-07-28T15:09:26Z"
tags:
  - hacker-news
  - translated
---

# Questions and answers on AI and verification: follow-up to my May ACM Tech Talk

- HN: [49085108](https://news.ycombinator.com/item?id=49085108)
- Source: [bertrandmeyer.com](https://bertrandmeyer.com/2026/07/21/questions-and-answers-on-ai-and-verification-a-follow-up-to-my-may-acm-tech-talk/)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T15:09:26Z

## Translation

タイトル: AI と検証に関する質疑応答: 5 月の ACM Tech Talk のフォローアップ
記事のタイトル: AI と検証に関する質問と回答: 5 月の ACM Tech Talk のフォローアップ - Bertrand Meyer の technology+ ブログ
説明: 5 月 7 日、私は「人工知能時代のソフトウェア検証」に関する ACM Tech 講演を行い、多くの参加者が集まりました。最後には答えきれないほどの質問がありましたが、私はすべてに答えると参加者に約束しました。その間に私は著書「AI for Smarties」を完成させなければなりませんでした
[切り捨てられた]

記事本文:
コンテンツにスキップ
Bertrand Meyer のテクノロジー+ブログ
ソフトウェアエンジニアリング、プログラミング方法論、言語、検証、一般的なテクノロジー、出版文化など
@Bertrand_Meyer をフォローしてください
メニュー
ホーム
AI と検証に関する質疑応答: 5 月の ACM Tech Talk のフォローアップ
5 月 7 日、私は「人工知能時代のソフトウェア検証」に関する ACM Tech 講演を行い、多くの参加者が集まりました。最後には答えきれないほどの質問がありましたが、私はすべてに答えると参加者に約束しました。その間に私は著書『AI for Smarties: Understanding Artificial Intelligence』を完成させる必要がありましたが、出版されたのでこの議論に戻ります。素晴らしい質問をありがとうございました!
主催者のヤン・ティマノフスキーは、会議のチャットに書き込まれて未回答のままになっていた質問のリストを私に提供してくれました。それらを以下に再掲します。今後 2 ～ 3 週間かけて、可能な限りの回答を提供するつもりですので、時々戻ってきてください。
質問には各質問者の名前と電子メールが含まれていました。参加者のプライバシーを保護するために、この情報は削除しました。
以下の「コメント」セクションを使用して、お気軽に質問やコメントを追加してください。早急にお答えできるよう最善を尽くします。
議論を続けることを楽しみにしています！
質問と（進行中の）いくつかの回答
[他の出席者のコメント]、「はい、その通りです。もし不純な起源から来たイノベーションを利用しなかったら、私たちは非常に貧しく、より病気になり、繁栄も減っていたでしょう。」、
[他の出席者の回答]、彼らは反対のようです。タイプ L は競争の場を平等にし、タイプ E は不平等を拡大します。、2026 年 5 月 7 日 12:12、
【他の参加者のコメント】 「スキル破壊」というのはレベリングタイプLのことだと思います。
本当にイノベーションなんてないよ」

スキルを「破壊」しますが、スキルのない人でも以前は高いスキル レベルが必要だった結果を達成できるようにすることで、スキルを無意味にするものもあります。」、2026 年 5 月 7 日 12:56、
18、「運用前に検証が完了したとみなされた後、セーフティ クリティカルなアプリケーション向けに AI によって作成されたソフトウェアが運用中に安全であることをどのように保証できますか?」、2026 年 5 月 7 日 12:16、
カテゴリー 人工知能 , コンピュータサイエンス , 契約による設計 , 教育 , エッフェル , 数学 , オブジェクトテクノロジー , プログラミング技術 , セミナー , ソフトウェアエンジニアリング
人工知能時代のソフトウェア検証に関する ACM ウェビナー後の Q&A
コメントを残す 返信をキャンセル
コメントを投稿するにはログインする必要があります。

## Original Extract

On May 7, I gave a well-attended ACM Tech talk on “Software Verification in the Age of Artificial Intelligence”. There were more questions at the end than I had the time to answer, but I promised the participants that I would answer everything. In the meantime I had to finish my book AI for Smarties
[truncated]

Skip to content
Bertrand Meyer's technology+ blog
Software engineering, programming methodology, languages, verification, general technology, publication culture, and more
Follow @Bertrand_Meyer
Menu
Home
Questions and answers on AI and verification: a follow-up to my May ACM Tech Talk
On May 7, I gave a well-attended ACM Tech talk on “ Software Verification in the Age of Artificial Intelligence ”. There were more questions at the end than I had the time to answer, but I promised the participants that I would answer everything. In the meantime I had to finish my book AI for Smarties: Understanding Artificial Intelligence , but now that it is out I am coming back to this discussion. Thanks for all the great questions!
Yan Timanovsky, the organizer, provided me with the list of questions that had been written on the meeting chat and left unanswered. They are reproduced below. My plan is to provide answers — to those for which I can — over the next two or three weeks, so please come back once in a while.
The questions contained the name and email of each questioner; I removed this information to preserve participants’ privacy.
Feel free to add questions, or comments, using the Comment section below. I will do my best to answer promptly.
Looking forward to continuing the discussion!
Questions and (in progress) some answers
[Other attendee comment],”Yes, absolutely. If we didn’t use innovations that came from impure origins, we’d be greatly impoverished, sicker, and would flourish less.”,
[Other attendee answer],They seem opposite. Type L levels the playing field while type E amplifies inequality.,5/7/2026 12:12,
[Comment by other attendeed] I think what you mean by “”skill-destroying”” is the levelling type-L.
No innovation really “”destroys”” a skill, but some make skills irrelevant by allowing unskilled people to achieve the results that previously required a high skill-level.”,5/7/2026 12:56,
18,”How can we ensure software written by AI for safety-critical applications remains safe during operations, after verication was deemed complete prior to operations?”,5/7/2026 12:16,
Categories Artificial Intelligence , Computer science , Design by Contract , Education , Eiffel , Mathematics , Object technology , Programming techniques , Seminar , Software engineering
Q&A after ACM Webinar on Software Verification in the Age of Artificial Intelligence
Leave a Comment Cancel reply
You must be logged in to post a comment.
