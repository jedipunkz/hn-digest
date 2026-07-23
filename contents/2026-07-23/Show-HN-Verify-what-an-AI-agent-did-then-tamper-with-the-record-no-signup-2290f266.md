---
source: "https://api.cinchor.com/proof/refund"
hn_url: "https://news.ycombinator.com/item?id=49022766"
title: "Show HN: Verify what an AI agent did, then tamper with the record (no signup)"
article_title: "Cinchor — Verify a real agent decision"
author: "foh_quarters"
captured_at: "2026-07-23T15:02:31Z"
capture_tool: "hn-digest"
hn_id: 49022766
score: 1
comments: 0
posted_at: "2026-07-23T15:01:05Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Verify what an AI agent did, then tamper with the record (no signup)

- HN: [49022766](https://news.ycombinator.com/item?id=49022766)
- Source: [api.cinchor.com](https://api.cinchor.com/proof/refund)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T15:01:05Z

## Translation

タイトル: HN を表示: AI エージェントが行ったことを検証し、記録を改ざんします (サインアップなし)
記事のタイトル: Cinchor — 実際のエージェントの決定を検証する
HN テキスト: 市場における自律型 AI の現状は、「何かをしてもらいたいが、それをやってくれると信じている」状態です。モデルを実行する操作ではすべてのログが保持され、ファーストパーティのログは本質的に最も強力ではありません。それは、「私は宿題をしました、私は宿題を採点しました。私の宿題は合格点を獲得しました。信頼してください。」です。 Cinchor は、自律エージェントのコンテキスト、指示、出力の検証可能な記録を共有するために運用で使用できる軽量ツールです。証明はアンカーされており、改ざん明示的な方法で取得可能です (私はこのための基盤としてカスタム L1 を構築し、現在すべてのバリデーターを操作しています。AI アカウンタビリティー層のトークン露出はゼロですが、バリデーターの独立性がロードマップにあります)。 HN コミュニティからの意見を求めています。重複している場合、すでに使用しているものとの重複はどこにありますか?

記事本文:
Cinchor · AI エージェントの責任記録
AI エージェントが下した実際の決定 - それを自分で検証してください。
サポート エージェントは、規定のポリシー上限に基づいてこの払い戻しを自動承認しました。それができた瞬間、
レコードのフィンガープリントは Omne チェーンに書き込まれ、追加専用であり、アプリケーションの範囲外にありました。
キーを貼り付けていないので、私たちの言葉をそのまま信じる必要はありません。以下の文字を変更してください。
チェーンには実際に起こったことがすでに保持されているため、証明は失敗します。
自己証明されたログではなく、第三者による証拠。決断がなされたとき、シンコールは、
その時点で、そのフィンガープリントは追加のみで Omne チェーンにコミットされました。その時に書かれたものなので、オペレーターもシンコールも戻って書き直すことはできません。
エージェントはそうしましたが、それでも合格しました。それが、取らなければならない内部監査ログとのまったくの違いです
信念: ここでの緑色の結果はチェーンに対してチェック可能であり、データベース自体が保証するものではありません。
完全な信頼モデル、譲歩を含む ·
コードで確認する ·
自分の記録を発行して検証する

## Original Extract

The current state of in-market autonomous AI is “want something done; trust that they do it”. The operation running the models hold all of the logs, and 1st party logs are inherently not the strongest. It’s: “I did my homework, I graded my homework.. my homework got a passing grade. Trust”. Cinchor is the lightweight tool that operations can use to share verifiable records of autonomous agents’ context, instructions and output. The proofs are anchored and retrievable in a tamper-evident manner (I built a custom L1 as a substrate for this and right now I operate all of the validators. There’s zero token-exposure for the AI accountability layer, while validator independence is on the roadmap). Looking for input from the HN community. Where is this redundant with something you already use, if so?

Cinchor · accountability records for AI agents
A real decision an AI agent made — verify it yourself.
A support agent auto-approved this refund under a stated policy ceiling. The instant it did,
the record's fingerprint was written to the Omne chain, append-only and outside the application's reach.
You didn't paste a key and you don't have to take our word for it: change any character below and the
proof fails, because the chain already holds what actually happened.
Third-party evidence, not a self-attested log. When the decision was made, Cinchor
committed its fingerprint to the Omne chain at that moment, append-only. Because it was written then, neither the operator nor Cinchor can go back, rewrite what
the agent did, and still pass. That's the whole difference from an internal audit log you'd have to take
on faith: a green result here is checkable against the chain, not our database vouching for itself.
The full trust model, concessions included ·
Verify it in code ·
Issue and verify your own record
