---
source: "https://twitter.com/_can1357/status/2087228354399265125"
hn_url: "https://news.ycombinator.com/item?id=49265135"
title: "OpenAI and Anthropic hidden CoT leaks when given deep_think tool."
article_title: "Can Bölük on X: \"guys you do know you can just disable thinking, and instead give it a \"deep_think\" tool, and it will call it with internal CoT reasoning format right?\ngl fixing that https://t.co/eWnPGbwxXs\" / X"
author: "himata4113"
captured_at: "2026-08-11T22:32:09Z"
capture_tool: "hn-digest"
hn_id: 49265135
score: 6
comments: 0
posted_at: "2026-08-11T22:06:14Z"
tags:
  - hacker-news
  - translated
---

# OpenAI and Anthropic hidden CoT leaks when given deep_think tool.

- HN: [49265135](https://news.ycombinator.com/item?id=49265135)
- Source: [twitter.com](https://twitter.com/_can1357/status/2087228354399265125)
- Score: 6
- Comments: 0
- Posted: 2026-08-11T22:06:14Z

## Translation

タイトル: deep_think ツールを使用すると、OpenAI と Anthropic の隠れた CoT がリークします。
記事のタイトル: X で Bölük はできますか: 「思考を無効にするだけで、代わりに「deep_think」ツールを与えると、内部の CoT 推論形式でそれを呼び出すことができることをご存知ですか?
gl はそれを修正します https://t.co/eWnPGbwxXs" / X
説明: 思考を無効にして代わりに「deep_think」ツールを与えると、内部 CoT 推論形式でそれを呼び出すことができることを知っていますよね?
glそれを修正します

記事本文:
X の Bölük: 「思考を無効にするだけで、代わりに "deep_think" ツールを与えると、内部の CoT 推論形式でそれを呼び出すことができることをご存知ですか?
glはそれを修正しています https://t.co/eWnPGbwxXs" / X Post
X の Bölük: 「思考を無効にするだけで、代わりに "deep_think" ツールを与えると、内部の CoT 推論形式でそれを呼び出すことができることをご存知ですか?
gl を修正してください https://t.co/eWnPGbwxXs"
Bölük @_can1357 さん、思考を無効にして代わりに「deep_think」ツールを与えると、内部の CoT 推論形式でそれを呼び出すことができることを知っていますか?
gl Alexander Panfilov @kotekjedi_ml を修正中 10 時間 ようやくそれについて話せます。
私たちは、あらゆるフロンティア AI 企業の API の脆弱性を利用して、フロンティア モデルの隠された推論を抽出する方法を発見しました。
クエリしたプロンプトのほとんどについて、推論トークンの数が請求された API 思考トークンと 1:1 で一致することを確認しました。 5:23 PM · 2026 年 8 月 11 日 255.1K 再生数 85 186 2.9K 1.6K
Can Bölük @_can1357 5 時間 これが PoC です pasta.can.ac/omegiligox.py
引数の形式については何も述べていないことに注意してください。gpt55+ の grug-talk 形式が得られます。
これ以外にも、ちょっと面白いことがある。 2 1 255 30K Can Bölük @_can1357 3 時間、適切な思考の流れを得ることができなかった ngl。思考レベルの切り替えも問題なく機能します。これは常にシステム プロンプトに数字が表示されます。 00:00 2 149 14K Can Bölük @_can1357 次のリリースでは 3 時間が延長されます。 8 152 11K
Can Bölük @_can1357 4 時間の寓話 ちなみに、人類には通常素晴らしい正規表現ゲームがあります、smh 2 134 10K
Can Bölük @_can1357 5 時間かかりましたが、そのための奇妙なドメインさえ取得できませんでした。クレイジー！！！ 3 100 10K
Kevin @kcosr 4 時間 これで完了です。間もなく、これらのモデルは、各ラボが提供するクラウドホスト型ハーネスでのみ使用できるようになります。 3 1 115 5.7K
何が起こっているかを確認してください

会話に参加してください
電話で続行 Apple で続行 Google で続行、またはユーザー名または電子メールでログイン 関係者

## Original Extract

guys you do know you can just disable thinking, and instead give it a "deep_think" tool, and it will call it with internal CoT reasoning format right?
gl fixing that

Can Bölük on X: "guys you do know you can just disable thinking, and instead give it a "deep_think" tool, and it will call it with internal CoT reasoning format right?
gl fixing that https://t.co/eWnPGbwxXs" / X Post
Can Bölük on X: "guys you do know you can just disable thinking, and instead give it a "deep_think" tool, and it will call it with internal CoT reasoning format right?
gl fixing that https://t.co/eWnPGbwxXs"
Can Bölük @_can1357 guys you do know you can just disable thinking, and instead give it a "deep_think" tool, and it will call it with internal CoT reasoning format right?
gl fixing that Alexander Panfilov @kotekjedi_ml 10h We can finally talk about it:
We found a way to extract hidden reasoning of frontier models using a vulnerability in the APIs of every frontier AI company.
We verified that our reasoning token count matches billed API thinking tokens 1:1 for most of the prompts we queried. 5:23 PM · Aug 11, 2026 255.1K Views 85 186 2.9K 1.6K
Can Bölük @_can1357 5h Here's a PoC pasta.can.ac/omegiligox.py
Note how I said nothing about the format of the argument and we get gpt55+'s grug-talk format.
Beyond this, which is kinda funny. 2 1 255 30K Can Bölük @_can1357 3h kinda missed having proper thinking streams ngl. switching thinking levels work fine too, it has always been a number in the system prompt! 00:00 2 149 14K Can Bölük @_can1357 3h will be up in next release, hf! 8 152 11K
Can Bölük @_can1357 4h fable btw, anthropic usually has great regex game, smh 2 134 10K
Can Bölük @_can1357 5h and I didn't even get a freaking domain for it! crazy!!! 3 100 10K
Kevin @kcosr 4h Now you've done it. Soon we will only be able to use these models in cloud-hosted harnesses provided by each lab. 3 1 115 5.7K
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
