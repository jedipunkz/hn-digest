---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48527480"
title: "I built a bookkeeping app for UK sole traders as a new developer using AI"
article_title: ""
author: "JamesQP"
captured_at: "2026-06-14T15:37:41Z"
capture_tool: "hn-digest"
hn_id: 48527480
score: 1
comments: 1
posted_at: "2026-06-14T14:24:30Z"
tags:
  - hacker-news
  - translated
---

# I built a bookkeeping app for UK sole traders as a new developer using AI

- HN: [48527480](https://news.ycombinator.com/item?id=48527480)
- Score: 1
- Comments: 1
- Posted: 2026-06-14T14:24:30Z

## Translation

タイトル: 新しい開発者として AI を使用して英国の個人事業主向けの簿記アプリを構築しました
HN テキスト: 約 1 か月前、私は QuarterPerfect の構築を開始しました。これは英国の個人事業主と家主を対象とした簿記ツールで、今年発効する MTD ITSA の変更に基づいて特別に設計されました (最初の四半期締め切り: 8 月 7 日)。まず私の背景を少し説明します。私は上級開発者ではありません。私の以前のプロジェクトは、XLSX のシフト勤務表を年次休暇をサポートするカレンダーに変換する小さな勤務表アプリでした。主に私と職場の数人の同僚のために構築されました。実際のユーザーに関するものや、支払いやコンプライアンス要件に関するものは何もありません。 QuarterPerfect はそれからの大幅な進歩です。 QuarterPerfect は、ユーザーに銀行取引明細書のインポート (CSV および PDF)、レシート照合による取引分類、定期的な取引を自動分類するための販売者ルール、および読み取り専用の会計士共有リンクを提供するため、クライアントは追跡することなくクリーンな記録を引き渡すことができます。テスト中のモバイルアプリではウェブファーストです。技術的な面は難しかったですが、学ぶことができました。さらに困難だったのは、私に実際の経験のない領域で構築することでした。HMRC/MTD の領域には、さまざまな収入タイプの扱い、「四半期ごとの更新」の実際の意味と年間利益、会計士のニーズと HMRC のニーズなど、多くのニュアンスがあります。それを間違えることはUXの問題ではなく、責任の問題です。 2 番目に大変だったのは、AI 支援を効果的に活用する方法を学ぶことでした。全体を通してクロードとパープレクシティを多用しました。初期の私は、アウトプットを無批判に受け入れすぎて、不安定な基盤の上に構築していました。最も役に立った変化は、コードジェネレーターとしてではなく、ペアのプログラマーのように扱うことでした。決定は依然として私が行う必要があり、私はそれを守るために自分が出荷しているものを十分に理解する必要があります。正直に言うと、本当に

良いワークフローは、Perplexity に私のアイデアを与え、Claude に詳細なプロンプトを生成させ、その後結果を Perplexity にポストするというものでした。多くのコピーアンドペーストが行われましたが、結果は正確でした。また、プロジェクト管理構造についても両者に同意してもらい、物事を整理整頓するのに本当に役立ちました。まだテスト中であり、一般ユーザーはまだいません。この製品はエンドツーエンドで機能しますが、コンプライアンス面でより自信が持てるまでは、見知らぬ人に製品を勧めるのは気が進みません。また、アプリを iOS と Google Play ストアにネイティブにしたいと考えています。後者の場合、アプリを 14 日間何らかの形で使用するテスターが少なくとも 12 名必要です。ソーシャル メディアで友人に協力を求めた後でも、まだ最低人数に達していません。正直に言うとかなり気が滅入っています。膨大な専門知識を持たずに、規制されたスペースやコンプライアンスが重視されるスペースを構築した人は他にいますか?たとえば、「これは機能します」と「これは誰かの納税記録を信頼するのに十分正しい」との間のギャップにどのように対処しましたか。 AI 支援開発ワークフローが実際に機能する便利な Web サイトやアプリの作成に本当に役立ったと感じた人はいますか? あなたの経験はどうですか?この投稿からユーザーを募集するつもりはありません。とにかく、この製品を適切に評価するにはサインアップが必要です。もっと知りたいのは、誰かが同じような立場に立ったことがあるかどうか、そしてその境遇を乗り越えるために何をしたかということだ。ありがとう、ジェームズ

## Original Extract

About a month ago I started building QuarterPerfect — a bookkeeping tool aimed at UK sole traders and landlords, specifically designed around the MTD ITSA changes coming into effect this year (first quarterly deadline: 7 August). A bit of background on me firstly. I'm not a senior developer. My previous project was a small roster app that converted an XLSX shift rota into a calendar with annual leave support — mostly built for me and a few colleagues at work. Nothing with real users, nothing with payments or compliance requirements. QuarterPerfect is a significant step up from that. QuarterPerfect offers users bank statement import (CSV and PDF), transaction categorization with receipt matching, merchant rules for auto-categorizing recurring transactions, and a read-only accountant share link so clients can hand off clean records without chasing. It's web-first with a mobile app in testing. The technical side was challenging but learnable. The harder thing was building in a domain I have no real background in. The HMRC/MTD space has a lot of nuance, treatment of different income types, what "quarterly updates" actually mean vs the annual return, what accountants need vs what HMRC needs. Getting that wrong isn't a UX problem, it's a liability problem. The second hard thing was learning to work effectively with AI assistance. I used Claude and Perplexity heavily throughout. Early on I'd accept output too uncritically and build on shaky foundations. The shift that helped most was treating it less like a code generator and more like a pair programmer, the decisions still have to be mine, and I have to understand what I'm shipping well enough to defend it. Honestly though, one really good work flow was to give Perplexity my ideas and have it generate detailed prompts for Claude, I would then post results back to Perplexity. A LOT of copy and pasting but results were spot on. I also had them both agree on a project management structure, which REALLY helped keep things tidy. It's still in testing, with no public users yet. The product works end-to-end but I'm not comfortable pushing it to strangers until I'm more confident in the compliance side. Also I want the app to be native to iOS and Google Play store, the latter requires me to have a minimum of 12 testers using the app in some capacity for 14 days, I haven't reached the minimum amount yet, even after asking for help from friends on social media, quite deflating to be honest! Has anyone else built in a regulated or compliance-heavy space without vast expertise? How did you handle the gap between "this works" and "this is correct enough to trust with someone's tax records" for example. Has anyone else found the AI-assisted development workflow genuinely helped them create a working/useful website/app, what has been your experience? I'm not looking to recruit users from this post, the product requires signup to evaluate properly anyway. More curious whether anyone's been in a similar position and what they did to get it over the line. Thanks, James

