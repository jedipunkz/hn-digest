---
source: "https://twitter.com/kotekjedi_ml/status/2087147042888114428"
hn_url: "https://news.ycombinator.com/item?id=49258338"
title: "Stealing Reasoning Traces from Proprietary LLM APIs"
article_title: "Alexander Panfilov on X: \"We can finally talk about it:\nWe found a way to extract hidden reasoning of frontier models using a vulnerability in the APIs of every frontier AI company.\nWe verified that our reasoning token count matches billed API thinking tokens 1:1 for most of the prompts we queried.\n[truncated]"
author: "rzk"
captured_at: "2026-08-11T14:13:26Z"
capture_tool: "hn-digest"
hn_id: 49258338
score: 3
comments: 0
posted_at: "2026-08-11T13:50:33Z"
tags:
  - hacker-news
  - translated
---

# Stealing Reasoning Traces from Proprietary LLM APIs

- HN: [49258338](https://news.ycombinator.com/item?id=49258338)
- Source: [twitter.com](https://twitter.com/kotekjedi_ml/status/2087147042888114428)
- Score: 3
- Comments: 0
- Posted: 2026-08-11T13:50:33Z

## Translation

タイトル: 独自の LLM API から推論トレースを盗む
記事のタイトル: Xに関するアレクサンダー・パンフィロフ: 「ついにそれについて話すことができます:
私たちは、あらゆるフロンティア AI 企業の API の脆弱性を利用して、フロンティア モデルの隠された推論を抽出する方法を発見しました。
クエリしたプロンプトのほとんどについて、推論トークンの数が請求された API 思考トークンと 1:1 で一致することを確認しました。
[切り捨てられた]
説明: ついにそれについて話すことができます:
私たちは、あらゆるフロンティア AI 企業の API の脆弱性を利用して、フロンティア モデルの隠された推論を抽出する方法を発見しました。
クエリしたプロンプトのほとんどについて、推論トークンの数が請求された API 思考トークンと 1:1 で一致することを確認しました。

記事本文:
アレクサンダー・パンフィロフ、Xについて：「ようやくそれについて話すことができます。
私たちは、あらゆるフロンティア AI 企業の API の脆弱性を利用して、フロンティア モデルの隠された推論を抽出する方法を発見しました。
クエリしたプロンプトのほとんどについて、推論トークンの数が請求された API 思考トークンと 1:1 で一致することを確認しました。 https://t.co/S7wN8aP3X7" / X ポスト
Alexander Panfilov @kotekjedi_ml ようやくそれについて話すことができます。
私たちは、あらゆるフロンティア AI 企業の API の脆弱性を利用して、フロンティア モデルの隠された推論を抽出する方法を発見しました。
クエリしたプロンプトのほとんどについて、推論トークンの数が請求された API 思考トークンと 1:1 で一致することを確認しました。 12:00 PM · 2026 年 8 月 11 日 80.9K 再生数 42 208 1.5K 851
Alexander Panfilov @kotekjedi_ml 2 時間 背景: 5 月に、@ matthew_d_green は、暗号化された推論が元のコンテキストの外で再生される可能性があることを発見し、それを研究室に報告しました (blog.cryptographyengineering.com/2026/05/29/foo…)。
研究所は「サイドチャネルやリプレイにはセキュリティ上の影響はないと考えている」と述べた。
もっと見る 暗号化された推論について話しましょう blog.cryptographyengineering.com より 2 5 69 8.3K Alexander Panfilov @kotekjedi_ml 2h モデル間の移植性は、Haiku 4.5 が Opus 4.8 の考えを読み取ることができることを意味します。
まあ、Opus の考えを取り入れて、少し脱獄すれば、直接攻撃することなく、Haiku に Opus の生の推論をそのまま転写させることができます。
OpenAI と Gemini でも同じトリックが機能します。 さらに表示 1 4 105 7.8K Alexander Panfilov @kotekjedi_ml 2h ご想像のとおり、これは、暗号を解読することなく推論トレースを抽出することが長い間可能であった可能性があることを示唆しています。
逸話: Kimi-K3 推論に Opus 推論のいくつかのトークンを事前に入力すると、反応が Show mo に向けて変化することがわかりました。

re 2 6 113 7.9K Alexander Panfilov @kotekjedi_ml 2h さらに、暗号化された推論 BLOB を含むクロード コード/コーデックス セッションをオンラインで共有したことがある場合、それらが解読されて個人データが漏洩する可能性があります。
約 7,000 件のパブリック トレースの予備スキャンを行ったところ、62 個の一意の API キー、33 個の電子メール アドレス、33 個のパスワード、その他の機密情報が見つかりました。 詳細を表示 2 15 120 7.1K Alexander Panfilov @kotekjedi_ml 2h 論文では、不正使用アップリフト (添付の写真を参照)、ジェイルブレイク、目に見えないプロンプト インジェクションなどのさらなる脅威について説明しています。 1 3 42 5.8K Alexander Panfilov @kotekjedi_ml 2h しかし、私たちはまた、いくつかの実際の陰謀、報酬追求などの例を見てみる機会を得て、それを付録にダンプしました。
1) サマライザーの不貞
推論の要約では、元のトレースから重要な情報が省略されることがよくあります。
ここで、Opus 4.8 は、自分が次のことを知っていることに気づきます。 詳細を表示 1 8 77 7.3K Alexander Panfilov @kotekjedi_ml 2h 2) 判読できない推論
[切り捨てられた]
あんぱーれ @anpaure 2h クソ 1 33 1.8K
何が起こっているかを確認して会話に参加してください
電話で続行 Apple で続行 Google で続行、またはユーザー名または電子メールでログイン 関係者

## Original Extract

We can finally talk about it:
We found a way to extract hidden reasoning of frontier models using a vulnerability in the APIs of every frontier AI company.
We verified that our reasoning token count matches billed API thinking tokens 1:1 for most of the prompts we queried.

Alexander Panfilov on X: "We can finally talk about it:
We found a way to extract hidden reasoning of frontier models using a vulnerability in the APIs of every frontier AI company.
We verified that our reasoning token count matches billed API thinking tokens 1:1 for most of the prompts we queried. https://t.co/S7wN8aP3X7" / X Post
Alexander Panfilov @kotekjedi_ml We can finally talk about it:
We found a way to extract hidden reasoning of frontier models using a vulnerability in the APIs of every frontier AI company.
We verified that our reasoning token count matches billed API thinking tokens 1:1 for most of the prompts we queried. 12:00 PM · Aug 11, 2026 80.9K Views 42 208 1.5K 851
Alexander Panfilov @kotekjedi_ml 2h Some background: In May, @ matthew_d_green found that encrypted reasoning could be replayed outside its original context, and reported it to the labs ( blog.cryptographyengineering.com/2026/05/29/foo… ).
The labs said that "they don’t see any security implications in side channels or replays".
In our Show more Let’s talk about encrypted reasoning From blog.cryptographyengineering.com 2 5 69 8.3K Alexander Panfilov @kotekjedi_ml 2h Cross-model portability means Haiku 4.5 can read Opus 4.8’s thoughts.
Well, if you take Opus thought, do a bit of jailbreaking, you can make Haiku transcribe the Opus' raw reasoning verbatim, without ever attacking it directly.
The same trick works with OpenAI and Gemini Show more 1 4 105 7.8K Alexander Panfilov @kotekjedi_ml 2h As you might guess, this suggests that distilling reasoning traces may have been possible for a long time without ever breaking the cryptography.
An anecdote: we find that prefilling Kimi-K3 reasoning with a few tokens of Opus reasoning measurably shifts its response toward Show more 2 6 113 7.9K Alexander Panfilov @kotekjedi_ml 2h Further, if you ever shared online a Claude Code/Codex session with encrypted reasoning blobs, they can be decoded and leak your personal data.
We did a preliminary scan of ~7,000 public traces and found 62 unique API keys, 33 email addresses, 33 passwords, and other sensitive Show more 2 15 120 7.1K Alexander Panfilov @kotekjedi_ml 2h In the paper we discuss more threats like misuse uplift (see the pic attached), jailbreaking and invisible prompt injection. 1 3 42 5.8K Alexander Panfilov @kotekjedi_ml 2h But we also took a chance to have a look at some in-the-wild scheming, reward seeking, etc. examples, and dumped it in appendix.
1) Summarizer unfaithfulness
Reasoning summaries often omit important information from the original trace.
Here, Opus 4.8 realizes it knows the Show more 1 8 77 7.3K Alexander Panfilov @kotekjedi_ml 2h 2) Illegible reasoning
[truncated]
anpaure @anpaure 2h holy shit 1 33 1.8K
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
