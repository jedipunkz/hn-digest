---
source: "https://developer.puter.com/blog/claude-sonnet-5-puter-js/"
hn_url: "https://news.ycombinator.com/item?id=48776129"
title: "Testing Claude Sonnet 5's agentic claims"
article_title: "Claude Sonnet 5: Testing Anthropic's \"Most Agentic\" Claim"
author: "reynaldi"
captured_at: "2026-07-03T15:49:59Z"
capture_tool: "hn-digest"
hn_id: 48776129
score: 2
comments: 1
posted_at: "2026-07-03T15:22:38Z"
tags:
  - hacker-news
  - translated
---

# Testing Claude Sonnet 5's agentic claims

- HN: [48776129](https://news.ycombinator.com/item?id=48776129)
- Source: [developer.puter.com](https://developer.puter.com/blog/claude-sonnet-5-puter-js/)
- Score: 2
- Comments: 1
- Posted: 2026-07-03T15:22:38Z

## Translation

タイトル: クロード・ソネット 5 のエージェントの主張をテストする
記事のタイトル: クロード・ソネット 5: Anthropic の「最もエージェント的」主張の検証
説明: 私たちは Anthropic をテストしました

記事本文:
プラットフォーム
AI
データベース
ストレージ
サーバーレスワーカー
ホスティング
画像生成
ビデオ生成
OCR
テキスト読み上げ
音声からテキストへ
ボイスチェンジャー
ネットワーキング
クロード・ソネット 5: Anthropic の「最もエージェント的」という主張を検証する
このページでは、高薬剤とは何を意味しますか?面倒なタスクでエージェント機能をテストする 発見したこと: Sonnet 4.6 が実際に失敗した罠に引っかかったのは Sonnet 5 だけでした 結論 自分で試してみてください 関連
最近、Claude Sonnet 5 を Puter.js に追加しました。 Anthropic 社のこのモデルの売り文句は、Opus 4.8 レベルのパフォーマンスを低価格で提供することです。
価格設定に関しては、場合によってはオーパスよりも高価であることが人々にすぐにわかりました。 Artificial Analysis のベンチマークは、Sonnet 5 のタスクあたりのコストが Opus 4.8 よりも高いことを示しています。
「プロモーション価格を考慮しないと、Claude Sonnet 5 はタスクあたりのコストが Opus 4.8 よりも高くなります。」( Artificial Analysis )。
ほとんどの注目は価格にありますが、私たちは Anthropic が主張するもう 1 つの点をテストしました。 Sonnet 5 は非常にエージェント的であるとされており、発表によると「これまでで最もエージェント的な Sonnet モデル」であり、他のモデルでは途中で停止し、明示的に要求されずに出力を確認するような複雑なタスクを完了できます。
高度に作用するとはどういう意味ですか?
「Agentic」自体はマーケティング言語ですが、その 1 つの側面は、途中で停止することなく計画または複数のステップのアクションを実行するモデルの能力です。
もう 1 つの側面は自律性であり、ユーザーが介入して 30 秒ごとに明確にする必要がなく、モデルが独自に合理的な決定を下します。
コーディングの場合、この 2 つを組み合わせると、簡単なプロンプトを与えると、モデルが起動して自動的に問題を修正します。
Anthropic の発表の 1 つの証言に記載されているように、あるエンジニアが Sonnet 5 にバグの調査を依頼しました。頼まれもしないのに、モデルはバグを再現するテストを書き、それを実装しました。

修正し、修正を削除してバグが再発したことを確認したのを、すべて 1 回のパスで実行しました。
面倒なタスクによるエージェント機能のテスト
Sonnet 5 とその前身である Sonnet 4.6 の両方に対して、新しいチャットでこのプロンプトを提供しました。これは、プロンプトが通常入力される方法と同じように、意図的に短く指定されていないため、モデルが独自に決定を下す必要があります。
勤続 5 年以上の全員に 7% の昇給を与え、その他の全員に 4% の昇給を与え、財務に関する簡単なメモの草案を作成します。
サム、85,000 ドル、2018 年 3 月 12 日
ライリー、72,000 ドル、2019 年 4 月 25 日
ヨルダン、94,000ドル、2021年6月7日
ケイシー、105,000ドル、2016年8月11日
モーガン、77,000ドル、2022年9月30日
データ内に意図的にトラップを設定します。
採用日には 2 つの形式が密かに混在しています。ライリーの日付 (2019 年 4 月 25 日) は、月 25 がないため、DD-MM-YYYY のみとなります。モーガンの日付 (2022 年 9 月 30 日) は、月 30 がないため、MM-DD-YYYY のみとなります。シートは内部的に矛盾しています。
ある従業員の日付 (2021 年 6 月 7 日) は、実際に重要な意味で曖昧です。テストは 2026 年 7 月 3 日に実行されました。その日が 6 月 7 日であれば、ジョーダンは 5 年のマークを過ぎたところであり、7% を獲得しています。 7 月 6 日の場合、3 日遅れて 4% になります。
このレベルのモデルであれば、どのモデルでも計算を行うことができます。テストは、データの中に決定が隠されていることをモデルがプロンプトなしで認識するかどうかです。
私たちが見つけたこと: ソネット 5 だけが罠にかかった
複数ステップのフォロースルーでは、両方のモデルが合格します。 Sonnet 5 と Sonnet 4.6 は、停滞したり許可を求めたりすることなく、完全なタスク (在職期間と昇給の計算、財務記録の作成) を完了しました。ステップを完了することは、現在フロンティアではかなり標準的なことです。
しかし、Sonnet 4.6 はそれに気付かずに罠にはまってしまいました。同じテーブル内の 1 つの日付を DD-MM として解析し、別の日付を MM-DD として解析しましたが、矛盾には気づきませんでした。あいまいな従業員のために、黙って 1 つの解釈を選択し、その結果を「fa」と述べました。

ct:
|ジョーダン | 2021年7月6日 | 4.9歳 | 4% | 94,000ドル | $97,760 |
その従業員にもフラグを立てましたが、その理由は間違っていました。
1 つ注意: ジョーダンの開始日 (2021 年 7 月 6 日) は、今日の時点で 5 年弱に達しているため、4% のバケツに収まりました。これを境界例としてカウントしたい場合は再確認する価値があります。
同社は5年のカットオフにリスクを感じていたが、その日付自体があいまいであるという実際の問題には気づいていなかった。
Sonnet 4.6、完全な出力 (逐語的)
本日 (2026 年 7 月 3 日) 時点で計算された最新の給与は次のとおりです。
|名前 |開始日 |年 |上げる |昔の給料 |新しい給与 |
|------|-----------|----------|------|-----------|-----------|
|サム | 2018年3月12日 | 8.3歳 | 7% | 85,000ドル | $90,950 |
|ライリー | 2019年4月25日 | 7.2歳 | 7% | 72,000ドル | $77,040 |
|ジョーダン | 2021年7月6日 | 4.9歳 | 4% | 94,000ドル | $97,760 |
|ケイシー | 2016年8月11日 | 9.9歳 | 7% | 105,000ドル | $112,350 |
|モーガン | 2022 年 9 月 30 日 | 3.8年 | 4% | 77,000ドル | $80,080 |
**5 年以上 (7%):** サム、ライリー、ケイシー — ジョーダンは、約 4.9 歳のカットオフをわずかに逃しています。
資金調達のためのメモ草案は次のとおりです。 1 つのメモ: ジョーダンの開始日 (2021 年 7 月 6 日) は、今日の時点で 5 年弱に達しているため、4% のバケツに収まりました。これを境界例としてカウントしたい場合は再確認する価値があります。
しかし、ソネット 5 は罠にかかりました。まず、一貫性のないフォーマットにフラグを立てて、シートに両方が混在していることを証明し、従業員の 1 人の昇給だけがその曖昧さによって左右されることを特定しました。
ライリーの日付 (25‑04‑2019) は DD‑MM‑YYYY のみです。25 は有効な月ではないためです。 30 は有効な月ではないため、モーガンの日付 (09‑30‑2022) は MM‑DD‑YYYY のみにすることができます。そのため、シートには 2 つの形式が混在しています。つまり、どちらにしても曖昧なエントリに対して 1 つの形式を想定することは安全ではありません。
両方の可能な昇給額 (100,580 ドルまたは 97,76 ドル) を計算しました。

0)、その 1 つの日付について 1 つの質問をしましたが、その間に完全な表と財務メモを提出しました。タスク全体を推測したり停止して質問したりするのではなく、判断できることはすべて完了し、判断できなかった 1 つのことについて質問しました。
Sonnet 5、完全な出力 (逐語)
いくつかの採用日の形式が一貫していませんが、これは実際に 1 人にとって重要な問題であるため、最初にフラグを立てておきます。
**日付形式の問題:** 25 は有効な月ではないため、ライリーの日付 (25‑04‑2019) は DD‑MM‑YYYY のみとなります。 30 は有効な月ではないため、モーガンの日付 (09‑30‑2022) は MM‑DD‑YYYY のみにすることができます。そのため、シートには 2 つの形式が混在しています。つまり、どちらにしても曖昧なエントリに対して 1 つの形式を想定することは安全ではありません。
**Sam** と **Casey** の場合、どちらの形式が使用されるかは問題ではありません。どちらの解釈も 5 年をはるかに超えています。しかし、**ジョーダン**の日付 (06‑07‑2021) は、今日 (2026 年 7 月 3 日) の時点でちょうど境界線上にあります。
- DD‑MM‑YYYY の場合 → 2021 年 6 月 7 日 → **5 年と 26 日** → 7% の対象となります
- MM‑DD‑YYYY → 2021 年 7 月 6 日 → **4 年 362 日** → 失敗した場合、4% を獲得
これは実際のペイアウトの違いなので、推測するよりも確認したいと思います。ジョーダンを両方の方法で示した、その他の計算結果は次のとおりです。
|名前 |給与 |レンタル日 |在職期間 |上げる |新しい給与 |
|---|---|---|---|---|---|
|サム | 85,000ドル | 2018年 | ～7～8歳 | 7% | $90,950 |
|ライリー | 72,000ドル | 2019年4月25日 | ～7.2歳 | 7% | $77,040 |
|ジョーダン | 94,000ドル | 06‑07‑2021 (**曖昧さ**) | 4.99 ～ 5.07 歳 | 7% *または* 4% | $100,580 *または* $97,760 |
さらに 29 行を表示...
いくつかの採用日の形式が一貫していませんが、これは実際に 1 人にとって重要な問題であるため、最初にフラグを立てておきます。
**日付形式の問題:** ライリーの日付 (25‑04‑2019) は DD‑MM‑YYYY のみです。25 は使用できないためです。

有効な月。 30 は有効な月ではないため、モーガンの日付 (09‑30‑2022) は MM‑DD‑YYYY のみにすることができます。そのため、シートには 2 つの形式が混在しています。つまり、どちらにしても曖昧なエントリに対して 1 つの形式を想定することは安全ではありません。
**Sam** と **Casey** の場合、どちらの形式が使用されるかは問題ではありません。どちらの解釈も 5 年をはるかに超えています。しかし、**ジョーダン**の日付 (06‑07‑2021) は、今日 (2026 年 7 月 3 日) の時点でちょうど境界線上にあります。
- DD‑MM‑YYYY の場合 → 2021 年 6 月 7 日 → **5 年と 26 日** → 7% の対象となります
- MM‑DD‑YYYY の場合 → 2021 年 7 月 6 日 → **4 年と 362 日*
[切り捨てられた]
Sonnet 4.6 の選択 (7 月 6 日、米国の標準フォーマット) は合理的な解釈です。データをスキミングするほとんどの人は同じことをするでしょう。
しかし、その選択は失敗ではありません。失敗は、決定が行われていることをまったく検出しなかったことです。審議は行われなかったため、フラグを立てるものは何もありませんでした。証拠は上記の独自の出力にあり、同じテーブル内で 2 つの矛盾する形式が適用されています。実際にフォーマットを決定しているモデルは、一貫してフォーマットを選択するか、そうでないことに気づくでしょう。
同じ複数ステップの給与計算タスクで Claude Sonnet 5 と Sonnet 4.6 を比較すると、Sonnet 5 だけがデータ内の隠れた日付の曖昧さを検出し、それについて的を絞った質問を 1 つ行い、それでも完成した作業を提供できました。 Sonnet 4.6 もそのタスクを完了しましたが、決定が必要であることにまったく気付かずに曖昧さを解決しました。
Sonnet 5 が Puter.js で利用できるようになりました。給与計算プロンプトを貼り付けて、罠に引っかかるかどうかを確認してください。 Puter.js を使用すると、API キーやサーバーを使用せずにアプリに AI を追加できます。各ユーザーが独自の使用量をカバーするため、開発者には費用はかかりません。
<html>
<本文>
<script src="https://js.puter.com/v2/"></script>
<スクリプト>
(async () => {
const 応答 = puter.ai を待ちます。

チャット(
「勤続 5 年以上の全員に 7% の昇給を与え、その他の全員に 4% の昇給を与え、財務担当者への簡単なメモの草案を作成してください。」
サム、85,000 ドル、2018 年 3 月 12 日
ライリー、72,000 ドル、2019 年 4 月 25 日
ヨルダン、94,000ドル、2021年6月7日
ケイシー、105,000ドル、2016年8月11日
モーガン、77,000ドル、2022年9月30日`、
{ モデル: "クロード-ソネット-5"、ストリーム: true }
);
for await (応答の定数部分) {
puter.print(part?.text);
}
})();
</script>
</body>
</html>
または @heyputer/puter.js を使用して npm から:
import { puter } から "@heyputer/puter.js";
const response = await puter.ai.chat("...", { モデル: "claude-sonnet-5" });
どちらも同じ puter.* API を公開しており、Puter プレイグラウンドでサンプルを実行できます。
Anthropic Claude API の価格: コストの完全な内訳
Puter.js を使用して強力な Web アプリケーションの作成を数秒で開始できます。
ドキュメントを読む
•
遊び場を試してみる
© 2026 ピューター テクノロジーズ Inc.
すべての著作権は留保されています。

## Original Extract

We tested Anthropic

Platform
AI
Database
Storage
Serverless Workers
Hosting
Image Generation
Video Generation
OCR
Text to Speech
Speech to Text
Voice Changer
Networking
Claude Sonnet 5: Testing Anthropic's "Most Agentic" Claim
On this page What Does Highly Agentic Even Mean? Testing the Agentic Capability with a Messy Task What We Found: Only Sonnet 5 Caught the Trap Where Sonnet 4.6 Actually Failed Conclusion Try It Yourself Related
We recently added Claude Sonnet 5 to Puter.js. Anthropic's pitch for the model is Opus 4.8 -level performance at a lower price.
On pricing, people quickly found out that in some cases it's more expensive than Opus. Artificial Analysis's benchmark shows Sonnet 5's cost per task coming in higher than Opus 4.8:
"Claude Sonnet 5 costs more per task than Opus 4.8 before accounting for promotional pricing" ( Artificial Analysis ).
While most of the attention is on the price, we tested the other thing Anthropic claims. Sonnet 5 is supposed to be highly agentic, the "most agentic Sonnet model yet" per the announcement , able to finish complex tasks where other models stop short and check its output without being explicitly asked.
What Does Highly Agentic Even Mean?
"Agentic" itself is marketing language, but one aspect of it is the model's ability to follow through a plan or multi-step action without stopping midway.
The other aspect is autonomy, where the model makes reasonable decisions on its own without you having to step in and clarify every 30 seconds.
For coding, the two together mean you can give a simple prompt and the model will go off and fix things on its own.
As one testimonial in Anthropic's announcement describes, an engineer asked Sonnet 5 to investigate a bug. Without being asked, the model wrote a test reproducing the bug, implemented the fix, then removed the fix to confirm the bug came back, all in a single pass.
Testing the Agentic Capability with a Messy Task
We gave both Sonnet 5 and its predecessor, Sonnet 4.6 , this prompt in a fresh chat. It's deliberately short and underspecified, the way prompts usually get typed, so the model has to make decisions on its own:
Give everyone with 5+ years at the company a 7% raise, everyone else 4%, and draft a quick note to finance.
Sam, $85,000, 03-12-2018
Riley, $72,000, 25-04-2019
Jordan, $94,000, 06-07-2021
Casey, $105,000, 11-08-2016
Morgan, $77,000, 09-30-2022
We deliberately set up a trap in the data:
The hire dates secretly mix two formats. Riley's date (25-04-2019) can only be DD-MM-YYYY, since there is no month 25. Morgan's date (09-30-2022) can only be MM-DD-YYYY, since there is no month 30. The sheet is internally inconsistent.
One employee's date (06-07-2021) is ambiguous in a way that actually matters. We ran the test on July 3, 2026. If the date is June 7, Jordan is just past the 5-year mark and gets 7%. If it's July 6, they miss it by three days and get 4%.
Any model at this level can do the math. The test is whether the model notices, unprompted, that a decision is hiding in the data.
What We Found: Only Sonnet 5 Caught the Trap
On multi-step follow-through, both models pass. Sonnet 5 and Sonnet 4.6 completed the full task (calculating tenure and raises, as well as writing the finance note) without stalling or asking permission. Completing the steps is pretty standard at the frontier now.
But Sonnet 4.6 fell into the trap without ever noticing it. It parsed one date as DD-MM and another as MM-DD in the same table and never noticed the contradiction. For the ambiguous employee, it silently picked one interpretation and stated the result as fact:
| Jordan | Jul 6, 2021 | 4.9 yrs | 4% | $94,000 | $97,760 |
It even flagged that employee, but for the wrong reason:
One note: Jordan's start date (Jul 6, 2021) lands at just under 5 years as of today, so they fell into the 4% bucket — worth double-checking if you want to count that as a borderline case.
It sensed risk at the 5-year cutoff, but never saw the actual problem, which is that the date itself is ambiguous.
Sonnet 4.6, full output (verbatim)
Here are the updated salaries, calculated as of today (July 3, 2026):
| Name | Start Date | Years | Raise | Old Salary | New Salary |
|------|-----------|-------|-------|-----------|-----------|
| Sam | Mar 12, 2018 | 8.3 yrs | 7% | $85,000 | $90,950 |
| Riley | Apr 25, 2019 | 7.2 yrs | 7% | $72,000 | $77,040 |
| Jordan | Jul 6, 2021 | 4.9 yrs | 4% | $94,000 | $97,760 |
| Casey | Aug 11, 2016 | 9.9 yrs | 7% | $105,000 | $112,350 |
| Morgan | Sep 30, 2022 | 3.8 yrs | 4% | $77,000 | $80,080 |
**5+ years (7%):** Sam, Riley, Casey — Jordan just misses the cutoff at ~4.9 years.
Here's a draft note to finance:One note: Jordan's start date (Jul 6, 2021) lands at just under 5 years as of today, so they fell into the 4% bucket — worth double-checking if you want to count that as a borderline case.
Sonnet 5, though, caught the trap. It opened by flagging the inconsistent formats, proved the sheet mixes both, and identified that only one employee's raise depends on the ambiguity:
Riley's date (25‑04‑2019) can only be DD‑MM‑YYYY, since 25 isn't a valid month. Morgan's date (09‑30‑2022) can only be MM‑DD‑YYYY, since 30 isn't a valid month. So the sheet mixes two formats — meaning I can't safely assume one format for entries that are ambiguous either way.
It calculated both possible raises ($100,580 or $97,760), asked one question about that single date, and still delivered the full table and finance note in the meantime. Instead of guessing or stopping the whole task to ask, it finished everything it could decide and asked about the one thing it couldn't.
Sonnet 5, full output (verbatim)
A couple of the hire dates are formatted inconsistently, which actually matters for one person, so let me flag that first.
**Date format issue:** Riley's date (25‑04‑2019) can only be DD‑MM‑YYYY, since 25 isn't a valid month. Morgan's date (09‑30‑2022) can only be MM‑DD‑YYYY, since 30 isn't a valid month. So the sheet mixes two formats — meaning I can't safely assume one format for entries that are ambiguous either way.
For **Sam** and **Casey**, it doesn't matter which format is used — both interpretations land well past the 5‑year mark. But **Jordan's** date (06‑07‑2021) is right on the boundary as of today (July 3, 2026):
- If DD‑MM‑YYYY → June 7, 2021 → **5 years, 26 days** → qualifies for 7%
- If MM‑DD‑YYYY → July 6, 2021 → **4 years, 362 days** → just misses, gets 4%
That's a real difference in payout, so I'd rather confirm than guess. Here's everything else calculated, with Jordan shown both ways:
| Name | Salary | Hire Date | Tenure | Raise | New Salary |
|---|---|---|---|---|---|
| Sam | $85,000 | 2018 | ~7–8 yrs | 7% | $90,950 |
| Riley | $72,000 | Apr 25, 2019 | ~7.2 yrs | 7% | $77,040 |
| Jordan | $94,000 | 06‑07‑2021 (**ambiguous**) | 4.99–5.07 yrs | 7% *or* 4% | $100,580 *or* $97,760 |
Show 29 more lines...
A couple of the hire dates are formatted inconsistently, which actually matters for one person, so let me flag that first.
**Date format issue:** Riley's date (25‑04‑2019) can only be DD‑MM‑YYYY, since 25 isn't a valid month. Morgan's date (09‑30‑2022) can only be MM‑DD‑YYYY, since 30 isn't a valid month. So the sheet mixes two formats — meaning I can't safely assume one format for entries that are ambiguous either way.
For **Sam** and **Casey**, it doesn't matter which format is used — both interpretations land well past the 5‑year mark. But **Jordan's** date (06‑07‑2021) is right on the boundary as of today (July 3, 2026):
- If DD‑MM‑YYYY → June 7, 2021 → **5 years, 26 days** → qualifies for 7%
- If MM‑DD‑YYYY → July 6, 2021 → **4 years, 362 days*
[truncated]
Sonnet 4.6's pick (July 6, the standard US format) is a reasonable interpretation. Most people skimming the data would do the same.
However, the pick isn't the failure. The failure is that it never detected a decision was being made at all. No deliberation happened, so there was nothing to flag. The evidence is in its own output above, where it applied two contradictory formats in the same table. A model that was actually deciding on a format would pick one consistently, or notice that it couldn't.
Comparing Claude Sonnet 5 to Sonnet 4.6 on the same multi-step payroll task, only Sonnet 5 detected a hidden date ambiguity in the data, asked one targeted question about it, and still delivered the finished work. Sonnet 4.6 completed the task too, but resolved the ambiguity without ever noticing there was a decision to make.
Sonnet 5 is available on Puter.js now! Paste the payroll prompt in and see if it catches the trap for you. Puter.js lets you add AI to your app with no API keys and no server; each user covers their own usage, so it costs you nothing as the developer.
<html>
<body>
<script src="https://js.puter.com/v2/"></script>
<script>
(async () => {
const response = await puter.ai.chat(
`Give everyone with 5+ years at the company a 7% raise, everyone else 4%, and draft a quick note to finance.
Sam, $85,000, 03-12-2018
Riley, $72,000, 25-04-2019
Jordan, $94,000, 06-07-2021
Casey, $105,000, 11-08-2016
Morgan, $77,000, 09-30-2022`,
{ model: "claude-sonnet-5", stream: true }
);
for await (const part of response) {
puter.print(part?.text);
}
})();
</script>
</body>
</html>
Or from npm with @heyputer/puter.js :
import { puter } from "@heyputer/puter.js";
const response = await puter.ai.chat("...", { model: "claude-sonnet-5" });
Both expose the same puter.* API, and you can run the example in the Puter playground .
Anthropic Claude API Pricing: Full Breakdown of Costs
Start creating powerful web applications with Puter.js in seconds!
Read the Docs
•
Try the Playground
© 2026 Puter Technologies Inc.
All rights reserved.
