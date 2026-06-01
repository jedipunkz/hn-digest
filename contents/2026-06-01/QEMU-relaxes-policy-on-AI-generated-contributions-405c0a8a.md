---
source: "https://lists.nongnu.org/archive/html/qemu-devel/2026-05/msg07614.html"
hn_url: "https://news.ycombinator.com/item?id=48345891"
title: "QEMU relaxes policy on AI-generated contributions"
article_title: "[PATCH] docs/devel: relax policy on AI-generated contributions"
author: "tambourine_man"
captured_at: "2026-06-01T00:29:50Z"
capture_tool: "hn-digest"
hn_id: 48345891
score: 3
comments: 0
posted_at: "2026-05-31T14:22:48Z"
tags:
  - hacker-news
  - translated
---

# QEMU relaxes policy on AI-generated contributions

- HN: [48345891](https://news.ycombinator.com/item?id=48345891)
- Source: [lists.nongnu.org](https://lists.nongnu.org/archive/html/qemu-devel/2026-05/msg07614.html)
- Score: 3
- Comments: 0
- Posted: 2026-05-31T14:22:48Z

## Translation

タイトル: QEMU、AI 生成の貢献に関するポリシーを緩和
記事のタイトル: [PATCH] docs/devel: AI 生成のコントリビューションに関するポリシーを緩和する

記事本文:
[PATCH] docs/devel: AI 生成のコントリビューションに関するポリシーを緩和
これまで、QEMU のコード来歴ポリシーはいかなる貢献も拒否していました
AI によって生成されたコンテンツが含まれるか、AI によって生成されたコンテンツから派生すると考えられます。全面禁止
LLM 出力を単独で使用できることはほとんどありませんでしたが、メンテナンスは簡単でしたが、
ツールが改良されるにつれて、絶対的な禁止はより困難になりました
正当化する。
この政策の動機となった懸念は変わっておらず、価値がある
正確に言うと、DCO は、提出者が法的権限を持っているかどうかに関するものです。
「創造的な表現」に関するものではなく、コードを貢献する権利。の
LLM 出力の著作権とライセンスのステータスはまだ未確定であるため、
質問はまだ残っています。変化したのはリスクのバランスです。
- AI支援コンテンツを受け入れるプロジェクトは深刻な事態に陥っていない
これまでに起きた法的トラブル、リスクの可能性を示唆
物質化は高くありません。
- Red Hat[1] などの他の組織は、リスクを次のように評価しています。
許容できる -- ただし、個々の開発者のコミュニティは許容できない
会社の法的支援があり、根拠のない紛争さえある
QEMU の作業から長期にわたって気が散ることになるでしょう。
AI 支援を許可するようにポリシーを改訂します。
著作権侵害は少なくとも元に戻すのが簡単であり、広がる可能性は低いです。
テスト、ドキュメント、機械的な変更、小さなバグ修正。コアコード
他のものが依存していて、一度では簡単に捨てられないもの
問題が事後になってから気づいた場合、事前の連絡なしに立ち入りを禁止したままにする
メンテナーからの同意。
これに関連しており、すでに驚くべき増加が見られます。
セキュリティ要件は、メンテナーの燃え尽き症候群の問題です。
コードの作成者からレビュー者に労力が移ります。 AIが下げる
パッチの作成コストはかかりますが、コストを削減することはできません。
理解と

1 つをレビューする。どちらかといえば、それはそれを引き起こします。
査読者は、提出者が推論を行ったとはもはや想定できません。
すべての行。上記の制限は、ボリュームを維持するのと同じように機能します。
持続可能な仕事を見直す。
さらに、記録のトレーラーとして「AI-used-for:」を紹介
[切り捨てられた]
Re: [PATCH] docs/devel: AI 生成のコントリビューションに関するポリシーの緩和、Michael S. Tsirkin、2026/05/28
Re: [PATCH] docs/devel: AI 生成のコントリビューションに関するポリシーの緩和、Paolo Bonzini、2026/05/28
Re: [PATCH] docs/devel: AI 生成のコントリビューションに関するポリシーの緩和、Michael S. Tsirkin、2026/05/28
Re: [PATCH] docs/devel: AI 生成のコントリビューションに関するポリシーの緩和、Paolo Bonzini、2026/05/28
Re: [PATCH] docs/devel: AI 生成のコントリビューションに関するポリシーの緩和、Michael S. Tsirkin、2026/05/28
Re: [PATCH] docs/devel: AI 生成のコントリビューションに関するポリシーの緩和、Philippe Mathieu-Daudé、2026/05/28
Re: [PATCH] docs/devel: AI 生成のコントリビューションに関するポリシーを緩和する、Alex Bennée 、2026/05/28
前の日付:
Re: 生成された AI とコードの出所について
次の日付:
Re: [PATCH] disas: sparc: Compare_opcodes() for ループの整数オーバーフローを修正
スレッドごとに前へ:
[お知らせ] QEMU 11.0.1 安定版がリリースされました
次にスレッド別に:
Re: [PATCH] docs/devel: AI 生成のコントリビューションに関するポリシーを緩和する

## Original Extract

[PATCH] docs/devel: relax policy on AI-generated contributions
Until now QEMU's code provenance policy declined any contribution
believed to include or derive from AI-generated content. A blanket ban
was easy to maintain while LLM output was rarely usable on its own, but
as the tools improved an absolute prohibition has become harder to
justify.
The concern that motivated the policy is unchanged, and it is worth
stating precisely: the DCO is about whether the submitter has the legal
right to contribute the code, not about "creative expression". The
copyright and license status of LLM output remains unsettled, so that
question is still open. What has shifted is the balance of risk:
- projects accepting AI-assisted content have not run into serious
legal trouble so far, which suggests the probability of the risk
materializing is not high;
- other organizations, such as Red Hat[1], have assessed the risk as
acceptable -- though a community of individual developers does not
have the legal backing of a company, and even an unfounded dispute
would be a long-lasting distraction from work on QEMU.
Revise the policy to permit AI assistance where the ramifications of
copyright violations are at least easy to revert and unlikely to spread:
tests, documentation, mechanical changes, and small bug fixes. Core code
that other things depend on, and that cannot simply be thrown away once
a problem is noticed long after the fact, stays off-limits without prior
agreement from a maintainer.
Related to this, and already visible in the incredible uptick in
security requirements, is the question of maintainer burnout and the
shift in effort from the author to the reviewer of the code. AI lowers
the cost of producing a patch but does nothing to lower the cost of
understanding and reviewing one; if anything it raises it, since a
reviewer can no longer assume that the submitter has reasoned through
every line. The limits above work just as much to keep the volume of
review work sustainable.
Furthermore, introduce "AI-used-for:" as a trailer to record wh
[truncated]
Re: [PATCH] docs/devel: relax policy on AI-generated contributions , Michael S. Tsirkin , 2026/05/28
Re: [PATCH] docs/devel: relax policy on AI-generated contributions , Paolo Bonzini , 2026/05/28
Re: [PATCH] docs/devel: relax policy on AI-generated contributions , Michael S. Tsirkin , 2026/05/28
Re: [PATCH] docs/devel: relax policy on AI-generated contributions , Paolo Bonzini , 2026/05/28
Re: [PATCH] docs/devel: relax policy on AI-generated contributions , Michael S. Tsirkin , 2026/05/28
Re: [PATCH] docs/devel: relax policy on AI-generated contributions , Philippe Mathieu-Daudé , 2026/05/28
Re: [PATCH] docs/devel: relax policy on AI-generated contributions , Alex Bennée , 2026/05/28
Prev by Date:
Re: on ai generated and code provenance
Next by Date:
Re: [PATCH] disas: sparc: Fix integer overflow in compare_opcodes() for loop
Previous by thread:
[ANNOUNCE] QEMU 11.0.1 Stable released
Next by thread:
Re: [PATCH] docs/devel: relax policy on AI-generated contributions
