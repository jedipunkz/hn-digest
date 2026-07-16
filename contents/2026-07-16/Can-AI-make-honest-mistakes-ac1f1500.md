---
source: "https://twitter.com/thsottiaux/status/2077630111499882637"
hn_url: "https://news.ycombinator.com/item?id=48941561"
title: "Can AI make honest mistakes?"
article_title: "Tibo on X: \"On file deletions. We’ve investigated a handful of reports where GPT-5.6 unexpectedly deleted files.\nWhat we have found is that this most commonly occurs when:\n- Full access mode is enabled and codex is run without sandboxing protections, including without auto review being\" / X"
author: "throwaway2027"
captured_at: "2026-07-16T23:50:20Z"
capture_tool: "hn-digest"
hn_id: 48941561
score: 3
comments: 1
posted_at: "2026-07-16T23:19:36Z"
tags:
  - hacker-news
  - translated
---

# Can AI make honest mistakes?

- HN: [48941561](https://news.ycombinator.com/item?id=48941561)
- Source: [twitter.com](https://twitter.com/thsottiaux/status/2077630111499882637)
- Score: 3
- Comments: 1
- Posted: 2026-07-16T23:19:36Z

## Translation

タイトル: AI は正直な間違いを犯すことができますか?
記事のタイトル: Tibo on X: 「ファイルの削除について。私たちは、GPT-5.6 がファイルを予期せず削除したといういくつかの報告を調査しました。
私たちが発見したのは、これは次の場合に最も一般的に発生するということです。
- フル アクセス モードが有効になっており、コーデックスはサンドボックス保護なしで実行されます (自動レビューなしなど)。" / X
説明: ファイルの削除について。私たちは、GPT-5.6 によってファイルが予期せず削除されたといういくつかのレポートを調査しました。
私たちが発見したのは、これは次の場合に最も一般的に発生するということです。
- フル アクセス モードが有効になっており、コーデックスはサンドボックス保護なし (自動レビューなしなど) で実行されます。

記事本文:
@thsottiaux ファイルの削除について。私たちは、GPT-5.6 によってファイルが予期せず削除されたといういくつかのレポートを調査しました。
私たちが発見したのは、これは次の場合に最も一般的に発生するということです。
- フル アクセス モードが有効で、自動レビューが有効になっていない場合を含め、サンドボックス保護なしでコーデックスが実行されます。
- モデルは、$ HOME 環境変数をオーバーライドして一時ディレクトリを定義しようとします。
- モデルは正直な間違いを犯し、代わりに誤って $ HOME を削除しました。
もちろん、これは、ユーザーがサンドボックスの保護手段なしで、またはこの種の高リスクのアクションをチェックして拒否する自動レビューを使用せずにフルアクセス モードでモデルを操作する場合でも、システムがどのように動作するかを望んでいるわけではありません。
私たちは、開発者メッセージを更新し、より多くのユーザーをより安全な許可モードに誘導し、追加のハーネス保護手段を追加するなど、このリスクを軽減するための措置を講じています。このようなことが起こることは非常にまれですが、詳細とリスクをさらに最小限に抑えるために私たちが行っていることについては、数日以内に詳細な事後分析を共有する予定です。 5:43 AM · 2026 年 7 月 16 日 979.5K 回視聴 6 4 6 646 3 7 0 370 8 。 2K 8.2K 1 。 5,000 1.5,000 646 件の返信を読む X は初めてですか?
今すぐサインアップして、自分だけのタイムラインを手に入れましょう!
Google でサインアップ Apple でサインアップ アカウントを作成 サインアップすると、Cookie の使用を含むサービス利用規約とプライバシー ポリシーに同意したことになります。
© 2026 X Corp. 何が起こっているかを見逃さないでください X の人々が最初に知ります。ログイン サインアップ 今トレンド

## Original Extract

On file deletions. We’ve investigated a handful of reports where GPT-5.6 unexpectedly deleted files.
What we have found is that this most commonly occurs when:
- Full access mode is enabled and codex is run without sandboxing protections, including without auto review being

@thsottiaux On file deletions. We’ve investigated a handful of reports where GPT-5.6 unexpectedly deleted files.
What we have found is that this most commonly occurs when:
- Full access mode is enabled and codex is run without sandboxing protections, including without auto review being enabled
- The model attempts to override the $ HOME env var to define a temporary directory.
- The model makes an honest mistake and mistakenly deletes $ HOME instead.
This is of course not how we want the system to behave, even when a user operates the model in full-access mode without the safeguards of our sandbox or without using auto review which checks for these kinds of high risk actions and rejects them.
We are taking steps to mitigate this risk including by updating the developer message, guiding more users towards safer permission modes, and adding additional harness safeguards. Even though this happens extremely rarely, we’ll share a detailed post-mortem in the coming days that goes into more details and what we are doing to minimize risks further. 5:43 AM · Jul 16, 2026 979.5K Views 6 4 6 646 3 7 0 370 8 . 2 K 8.2K 1 . 5 K 1.5K Read 646 replies New to X?
Sign up now to get your own personalized timeline!
Sign up with Google Sign up with Apple Create account By signing up, you agree to the Terms of Service and Privacy Policy , including Cookie Use.
© 2026 X Corp. Don't miss what's happening People on X are the first to know. Log in Sign up Trending now
