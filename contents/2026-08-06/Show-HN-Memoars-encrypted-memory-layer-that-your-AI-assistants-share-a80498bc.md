---
source: "https://memoars.com/"
hn_url: "https://news.ycombinator.com/item?id=49195751"
title: "Show HN: Memoars - encrypted memory layer that your AI assistants share"
article_title: ""
author: "shonster88"
captured_at: "2026-08-06T12:51:17Z"
capture_tool: "hn-digest"
hn_id: 49195751
score: 1
comments: 1
posted_at: "2026-08-06T12:29:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Memoars - encrypted memory layer that your AI assistants share

- HN: [49195751](https://news.ycombinator.com/item?id=49195751)
- Source: [memoars.com](https://memoars.com/)
- Score: 1
- Comments: 1
- Posted: 2026-08-06T12:29:24Z

## Translation

タイトル: Show HN: Memoars - AI アシスタントが共有する暗号化されたメモリ層
HN テキスト: AI エージェント間でのコンテキストの共有、知識の共有、記憶の共有に少し苦労しました (毎日 2 つまたは 3 つを使用しています)。それぞれが独自のメモリを持っており、メモリを共有していないか、メモリのキュレーションを行って改善するのが少し面倒です (特に、異なるモデルから時々実行される API AI 呼び出しがある場合) Memoars は、それを解決する試みです。ユーザーに属する 1 つのメモリ (ストレージ ベンダー ロックなし) で、アシスタントは MCP を介して読み書きできます。仕組み: - メモリの内容はマシン上で暗号化されます (XChaCha20-Poly1305、
キーは Argon2id で派生し、ストレージに直接書き込まれます。
独自の - R2、S3、MinIO、Supabase など) - コーディネーターはメタデータ プレーンを処理します: シーケンス番号、
バージョン、許可、競合解決。決して受け取らない
ワークスペースコンテンツキーがあるため、メモリコンテンツを読み取ることができません。それはあります
運用メタデータを参照 - 組織、ワークスペース、ID、バージョン、
使用法 - すべての変更は、追加専用のハッシュチェーンされたログに記録されます。
書き込み時に比較と交換を行うため、2 つのクライアントが黙ってデータを破壊することはありません
相互に確認すると、メモリがどのようにして現在の状態に到達したかを確認できます。 - 権限は組織→ワークスペース→アイデンティティであり、ワークスペースごとにあります
助成金。各ワークスペースには独自のパスフレーズがあるため、分離は
API だけでなく暗号化によっても強制されます。実際の場所: 現在は招待制です。正直に言うと、これは
今日の午後にインストールできる製品ではなく、招待リストを使用してください (出荷前にそれが意味があり、問題を解決できるかどうかを確認したいからです)。クライアントはオープンソース化されており、ホスト型コーディネーターが開きます
その後すぐに。詳細については、メールまたは ping を送信してください。すべての問い合わせに返信しますので、お気軽にお問い合わせください。

見てください！

## Original Extract

I struggled a bit with context sharing, knowledge sharing, memories sharing between AI agents (I use two or three on a daily basis). Each of them has its own memory, they dont share it or its a bit cumbersome to do memory curation and improve it (especially if there are some API AI calls that run occasionally from different models) Memoars is an attempt to solve it - one memory that belongs to you (no storage vendor lock) that any assistant can read and write through MCP. How it works: - Memory content is encrypted on your machine (XChaCha20-Poly1305,
key derived with Argon2id) and written directly to storage you
own - R2, S3, MinIO, Supabase etc) - A coordinator handles the metadata plane: sequence numbers,
versions, grants, conflict resolution. It never receives the
workspace content key, so it can't read memory content. It does
see operational metadata - org, workspace, identity, version,
usage - Every change lands in an append-only, hash-chained log with
compare-and-swap on writes, so two clients can't silently clobber
each other and you can see how a memory got to its current state. - Permissions are orgs → workspaces → identities, with per-workspace
grants. Each workspace has its own passphrase, so isolation is
enforced by encryption as well as by the API. Where it actually is: It's invite-only right now, and I want to be honest that this is a
invite list rather than a product you can go install this afternoon (as I want to make sure it makes sense and that it solves a problem for you before its shipped). The client is being open-sourced and the hosted coordinator opens
shortly after. Drop me a mail or ping me for more info - as I will reply to all inquiries Tnx for taking a look!

