---
source: "https://twitter.com/lafalcemateo/status/2082250304330809738"
hn_url: "https://news.ycombinator.com/item?id=49103230"
title: "There is no way to know if an LLM API is manipulating you"
article_title: "Mateo Lafalce on X: \"what if they're rigging your API to manipulate you deliberately? Is there any way to know?\nwith a provider-hosted model: no https://t.co/SBoHEZhzXH\" / X"
author: "lafalce"
captured_at: "2026-07-29T21:49:27Z"
capture_tool: "hn-digest"
hn_id: 49103230
score: 1
comments: 0
posted_at: "2026-07-29T21:19:13Z"
tags:
  - hacker-news
  - translated
---

# There is no way to know if an LLM API is manipulating you

- HN: [49103230](https://news.ycombinator.com/item?id=49103230)
- Source: [twitter.com](https://twitter.com/lafalcemateo/status/2082250304330809738)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T21:19:13Z

## Translation

タイトル: LLM API がユーザーを操作しているかどうかを知る方法はありません
記事のタイトル: Mateo Lafalce on X: 「もし彼らがあなたを意図的に操作するために API を不正操作していたらどうしますか? それを知る方法はありますか?」
プロバイダーホスト型モデルの場合: いいえ https://t.co/SBoHEZhzXH" / X
説明: もし彼らがあなたを意図的に操作するためにあなたの API を不正に操作していたらどうなるでしょうか?知る方法はありますか？
プロバイダーホスト型モデルの場合: いいえ

記事本文:
X についてのマテオ・ラファルス: 「もし彼らがあなたを意図的に操作するために API を不正操作していたらどうしますか? それを知る方法はありますか?
プロバイダーホスト型モデルの場合: いいえ https://t.co/SBoHEZhzXH" / X Post
Mateo Lafalce @lafalcemateo 彼らがあなたを意図的に操作するためにあなたの API を操作しているとしたらどうしますか?知る方法はありますか？
プロバイダーホスト型モデルの場合: いいえ 11:42 PM · 2026 年 7 月 28 日 21 ビュー 1
Mateo Lafalce @lafalcemateo 22h API を使用すると、重要な 3 つのレイヤーが見えなくなります。
- 実際に応答しているモデルはどれですか
- 効果的なシステム プロンプトとは何ですか
- 重みが変更されたかどうか
あなたは完全にプロバイダーのなすがままです。 1 12 Mateo Lafalce @lafalcemateo 22h 自分の VPS で推論を実行すると、そのうち 2 つが返されます。モデルのバージョンとシステム プロンプトを検証できます。
しかし、ウェイトはアキレス腱です。これは一度書き込まれたファイルであり、二度と見ることはなく、ボックスの所有者が誰でもアクセスできます。 1 12 Mateo Lafalce @lafalcemateo 22h 理想的なプライバシーの世界では、モデルを独自のハードウェアで実行します。モデル、プロンプト、重みはすべて正当でハッシュ検証可能です。
それが実質的な天井です。 1 10 Mateo Lafalce @lafalcemateo 22h それでも安全というわけではありません。実行するソフトウェアやハードウェアにはバックドアが組み込まれている可能性があります。ファームウェア、ドライバー、マイクロコード、シリコン。
信頼は消えるのではなく、動きます。問題は決して「私は信頼できるか？」ということではありません。しかし、「誰が、どの層で？」 3
今すぐサインアップして、自分だけのタイムラインを手に入れましょう!
Google でサインアップ Apple でサインアップ アカウントを作成 サインアップすると、Cookie の使用を含むサービス利用規約とプライバシー ポリシーに同意したことになります。
© 2026 X Corp. 最新情報をお見逃しなく
X の人々が最初に知りました。

## Original Extract

what if they're rigging your API to manipulate you deliberately? Is there any way to know?
with a provider-hosted model: no

Mateo Lafalce on X: "what if they're rigging your API to manipulate you deliberately? Is there any way to know?
with a provider-hosted model: no https://t.co/SBoHEZhzXH" / X Post
Mateo Lafalce @lafalcemateo what if they're rigging your API to manipulate you deliberately? Is there any way to know?
with a provider-hosted model: no 11:42 PM · Jul 28, 2026 21 Views 1
Mateo Lafalce @lafalcemateo 22h Through an API you're blind to the three layers that matter:
- which model is actually answering
- what the effective system prompt is
- whether the weights were modified
You're entirely at the provider's mercy 1 12 Mateo Lafalce @lafalcemateo 22h Running inference on your own VPS gets two of them back: you can validate the model version and the system prompt.
But the weights are the Achilles' heel. It's a file written once that nobody ever looks at again, and whoever owns the box has access. 1 12 Mateo Lafalce @lafalcemateo 22h In the ideal privacy world you run the model on your own hardware: model, prompt and weights all legitimate and hash-verifiable.
That's the practical ceiling. 1 10 Mateo Lafalce @lafalcemateo 22h Even then you're not safe: the software and hardware you run can carry backdoors. Firmware, drivers, microcode, silicon.
Trust doesn't disappear, it moves. The question is never "do I trust?" but "whom, and at which layer?" 3
Sign up now to get your own personalized timeline!
Sign up with Google Sign up with Apple Create account By signing up, you agree to the Terms of Service and Privacy Policy , including Cookie Use.
© 2026 X Corp. Don't miss what's happening
People on X are the first to know.
