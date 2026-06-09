---
source: "https://www.bejarano.io/terraform-autoapply/"
hn_url: "https://news.ycombinator.com/item?id=48458516"
title: "Safe Terraform auto-apply with conftest"
article_title: "Safe Terraform auto-apply with conftest"
author: "ricardbejarano"
captured_at: "2026-06-09T10:18:57Z"
capture_tool: "hn-digest"
hn_id: 48458516
score: 2
comments: 0
posted_at: "2026-06-09T09:01:56Z"
tags:
  - hacker-news
  - translated
---

# Safe Terraform auto-apply with conftest

- HN: [48458516](https://news.ycombinator.com/item?id=48458516)
- Source: [www.bejarano.io](https://www.bejarano.io/terraform-autoapply/)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T09:01:56Z

## Translation

タイトル: conftest を使用した安全な Terraform 自動適用
説明: Terraform 計画の人間によるレビューは拡張性がなく、AI レビューは本番環境に十分な決定性がありません。ここでは、conftest を使用して、明示的でテスト可能なポリシーを満たすプランを自動適用する方法を説明します。

記事本文:
安全なTerraform
conftest で自動適用
2026 年 6 月 6 日発行、Ricard Bejarano
儀式はご存知でしょう。変更が行われ、Terraform が計画され、誰かがそれをレビューし、
承認され、適用されます。速度が十分に低い場合、これは機能します。の
査読者は奇妙な間違いを見つけ、全員がよく眠ります。
ある時点を超えると、レビュー担当者がボトルネックになります。計画の山
立ち上がると、エンジニアが急いで通り抜けるか放置するかのどちらかになり、あなたは負け始めます
速度かレビューの質のいずれかです。多くの場合、両方です。
私たちがすぐに考えているのは、レビューを AI に委任することです。
計画のレビューを AI で補完することもできますが、最も興味深いのは
この分野で私が見つけた解決策は、Overmind↗ —あなたです
実稼働インフラストラクチャではなく、計画レビューを完全に委任することはできません。
それは非決定的です。同じ計画が今日は成功しても、明日は失敗する可能性があります。
多くの場合、人間による監査/コンプライアンスの要件に違反します。
明確な責任を持って承認する。そして批判的に
フィードバック ループから責任が取り除かれ、誰も所有者がなくなります。
何かが壊れたとき、それはまさに望ましくない決断です。
3 番目のオプションは、Terraform プランをプログラムで評価することです。
コードとしてのポリシーを使用して決定的に。それが私たちがやっていることであり、
コンテスト↗ 。
conftest↗ は、上に構築されたコードとしてのポリシー ツールです。
ポリシーエージェント↗ を開きます。ポリシーを書き込むのは、
Rego↗ 、食べさせてください
JSON データ。データがポリシーを満たしているかどうかがわかります。
重要な洞察は、Terraform が計画を JSON としてエクスポートできることです。
terraform plan -out=plan.tfplan
terraform show -json plan.tfplan > plan.json
この JSON ファイルには、Terraform が行う予定のすべてのリソース変更が含まれています。
何が作成、更新、削除されているか、およびそれぞれの前後の値
属性。これは人間のレビュー担当者が検討する情報と同じです。

ある
conftest などのポリシー エンジンが評価できる構造化フォーマット:
conftest テスト計画.json
計画がポリシーを満たしていれば、合格します。そうでない場合は、次のエラーが発生して失敗します。
明確な理由。決定は監査可能、テスト可能、再現可能です。
これは、すべての変更がノーオペであるプランのみを許可する Rego ポリシーです。
リソースの作成、またはデータ ソースの読み取り。更新または削除はポリシーに失敗します。
パッケージメイン
rego.v1 をインポートする
safe_actions := {"操作なし"、"作成"、"読み取り"}
拒否にメッセージが含まれている場合、{
input.resource_changes 内の一部の resource_change
resource_change.change.actions 内のアクション
safe_actions のアクションではありません
msg := sprintf(
"リソース %q にはアクション %q がありますが、これは安全セット %v にありません",
[resource_change.address、アクション、safe_actions]、
)
}
このポリシーは、JSON 形式のすべての resource_changes エントリを反復します。
テラフォーミング計画。それぞれについて、そのすべてのアクションが
安全なアクションが設定されています。アクションがそのセットの範囲外にある場合 (更新または
delete )、ポリシーは問題のあるリソースとアクションを含む拒否を発行します。
それでおしまい。このポリシーが合格した場合、計画は新しいリソースを作成するだけです。
データ ソースを読み取るか、何もしないので、自動適用しても安全です。もしそうなら
失敗するとパイプラインが停止し、人間がレビューします。
注: 使用する Terraform プロバイダーによっては、
新しいリソースの作成は完全に無害ではない可能性があります。ここでのポイントは、あなたが
組織の「安全なポリシー」の定義に合わせて独自のポリシーを作成します。
「自動適用」プランとは、以下で説明するように意味します。
CI/CD の統合は簡単です。 Terraform プランを作成したら、プランをエクスポートします
JSON に変換し、 conftest を実行し、結果に応じて分岐します。
terraform plan -out=plan.tfplan
terraform show -json plan.tfplan > plan.json
if conftest テスト計画.json;それから
terraform 適用 plan.tfplan
それ以外の場合
# 人間のアプリのゲート

ローバル
フィ
これがうまく機能するのは、決定境界が明確であるためです。あなたは
計画が「安全そうに見える」かどうかを誰か（または何か）に判断してもらうのではありません。あなたは
定義、テスト、バージョン管理した一連のルールを満たしているかどうかを確認する
インフラストラクチャ コードと一緒に。
上記の例は意図的に最小限にしています。データ ソースの作成のみを許可します。
読み取り、操作なし。実際には、より充実したポリシーと JSON が必要になります。
Terraform プランでは、次のようなさまざまな作業が可能です。
リソースの種類。すべてのリソースが同じリスクを伴うわけではありません。自動で申請してもらえるかも知れません
CloudWatch アラームに変更されますが、常に RDS インスタンスまたは IAM ポリシーをゲートオンします。
各 resource_changes エントリの type フィールドからは次のことがわかります。
safe_resource_types := {"aws_cloudwatch_metric_alarm"}
拒否にメッセージが含まれている場合、{
input.resource_changes 内の一部の resource_change
safe_resource_types の resource_change.type ではありません
resource_change.change.actions 内のアクション
{"no-op"、"read"} にないアクション
msg := sprintf("リソース %q のタイプ %q は自動適用セーフ セットに含まれていません", [resource_change.address, resource_change.type])
}
リソースフィールド。場合によっては、リソース タイプだけでは不十分な場合があります。
特定の属性にのみ影響を与える変更を自動適用します。の変更オブジェクト
JSON プランでは、個々のフィールドを比較できます。このポリシーは、次のような更新を拒否します。
タグを超えてフィールドを変更します。
拒否にメッセージが含まれている場合、{
input.resource_changes 内の一部の resource_change
resource_change.change.actions 内のアクション
アクション == "更新"
変更されたキー := {キー |
object.keys(resource_change.change.after) 内のいくつかのキー
resource_change.change.before[キー] != resource_change.change.after[キー]
}
Changed_keys != {"タグ", "タグすべて"}
msg := sprintf("リソース %q はタグ以外のフィールドを変更します: %v", [resource_change.address,Changed_keys])
}
爆発範囲。あ

2 つのリソースに関わるプランは、
実際の変更とゲートを含むリソースをカウントできます。
数値が指定されたしきい値を超えています:
max_auto_apply_changes := 10
拒否にメッセージが含まれている場合、{
変更されました := {リソース変更アドレス |
input.resource_changes 内の一部の resource_change
resource_change.change.actions 内のアクション
{"no-op"、"read"} にないアクション
}
カウント(変更) > max_auto_apply_changes
msg := sprintf("計画は %d 個のリソースに影響します。これは %d の自動適用制限を超えています", [count(changed), max_auto_apply_changes])
}
環境。ステージングでは自動適用され、本番ではゲートされます。あなたのリソースがあれば、
環境がタグ付けされている場合は、計画からそれを読み取ることができます。このポリシー
環境タグがそうでないリソースへの重要な変更を拒否します。
ステージング:
拒否にメッセージが含まれている場合、{
input.resource_changes 内の一部の resource_change
resource_change.change.actions 内のアクション
{"no-op"、"read"} にないアクション
resource_change.change.after.tags.Environment != "ステージング"
msg := sprintf("リソース %q はステージングされていないため、人間によるレビューが必要です", [resource_change.address])
}
これらのルールが構成されます。これらを同じポリシー ファイル内で組み合わせて、conftest を実行できます。
それらすべてを評価します。プランが自動適用されるには、すべてのルールに合格する必要があります。
たった 1 回の拒否でも、ポリシーは失敗するのに十分です。ポリシーはあなたの成長とともに成長します
自信があり、コードなので、自分と同じようにバージョン管理してテストできます。
他のコードでも。
AIを導入するとこのような仕組みがますます重要になります
エージェントを SDLC に派遣し、エージェントに変更を提案して実行させます。
ライブインフラストラクチャ。計画の安全性を証明する決定的な方法がなければ、
信頼性か速度、またはその両方を妥協する必要があります。
探していたものは見つかりましたか?
そうでない場合はお知らせください。
Copyright © 2026 リカール・ベジ

梅

## Original Extract

Human review of Terraform plans doesn’t scale, and AI review isn’t deterministic enough for production. Here’s how we use conftest to auto-apply plans that meet an explicit, testable policy.

Safe Terraform
auto-apply with conftest
Published Jun 6, 2026 by Ricard Bejarano
You know the ritual: a change is made, Terraform plans, someone reviews it,
approves it, and it gets applied. At low enough velocity, this works. The
reviewer catches the odd mistakes, and everyone sleeps well.
Past a certain point, the reviewer becomes the bottleneck . Plans pile
up, engineers either rush through them or let them sit, and you start losing
either velocity or review quality. Often both.
Our immediate next thought is to delegate review to AI .
And while you can complement your plan review with AI—the most interesting
solution I’ve found in this space is Overmind↗ —you
cannot fully delegate plan review to it, not for production infrastructure:
it’s non-deterministic : the same plan may pass today and fail tomorrow;
it often violates audit/compliance requirements that mandate human
sign-off with clear accountability; and critically
it removes responsibility from the feedback loop , no one owns the
decision, which is exactly what you don’t want when something breaks.
There’s a third option: evaluating Terraform plans programmatically and
deterministically using policy-as-code . That’s what we do, with
conftest↗ .
conftest↗ is a policy-as-code tool built on
Open Policy Agent↗ . You write policies in
Rego↗ , feed it
JSON data, and it tells you whether your data satisfies your policy.
The key insight is that Terraform can export its plan as JSON :
terraform plan -out=plan.tfplan
terraform show -json plan.tfplan > plan.json
That JSON file contains every resource change Terraform intends to make:
what’s being created, updated, deleted, and the before/after values of each
attribute. It’s the same information a human reviewer would look at, in a
structured format a policy engine—like conftest—can evaluate:
conftest test plan.json
If the plan satisfies your policy, it passes. If it doesn’t, it fails with an
explicit reason. The decision is auditable, testable, and reproducible.
Here’s a Rego policy that only allows plans where every change is a no-op, a
resource create, or a data source read. Any update or delete fails the policy:
package main
import rego.v1
safe_actions := {"no-op", "create", "read"}
deny contains msg if {
some resource_change in input.resource_changes
some action in resource_change.change.actions
not action in safe_actions
msg := sprintf(
"resource %q has action %q, which is not in the safe set %v",
[resource_change.address, action, safe_actions],
)
}
This policy iterates over every resource_changes entry in the JSON-formatted
Terraform plan. For each one, it checks whether all of its actions are in the
safe_actions set. If any action falls outside that set (an update or a
delete ), the policy emits a denial with the offending resource and action.
That’s it. If this policy passes, the plan only creates new resources,
reads data sources, or does nothing , so it’s safe to auto-apply. If it
fails, the pipeline stops and a human reviews.
Note: depending on what Terraform providers you use,
new resource creation may not be completely harmless. Point here is that you
create your own policy to suit your organization’s definition of what a “safe to
auto-apply” plan means, as we will see below.
The CI/CD integration is straightforward. After Terraform plans, export the plan
to JSON, run conftest , and branch on the result:
terraform plan -out=plan.tfplan
terraform show -json plan.tfplan > plan.json
if conftest test plan.json; then
terraform apply plan.tfplan
else
# gate on human approval
fi
What makes this work well is that the decision boundary is explicit . You’re
not asking someone (or something) to judge whether a plan “looks safe”. You’re
checking whether it satisfies a set of rules you defined, tested, and versioned
alongside your infrastructure code.
The example above is deliberately minimal: it only allows creates, data source
reads, and no-ops. In practice, you’ll want a richer policy, and the JSON
Terraform plan gives you plenty to work with:
Resource types. Not all resources carry the same risk. You might auto-apply
changes to CloudWatch alarms, but always gate on RDS instances or IAM policies.
The type field on each resource_changes entry gives you this:
safe_resource_types := {"aws_cloudwatch_metric_alarm"}
deny contains msg if {
some resource_change in input.resource_changes
not resource_change.type in safe_resource_types
some action in resource_change.change.actions
action not in {"no-op", "read"}
msg := sprintf("resource %q has type %q, which is not in the auto-apply safe set", [resource_change.address, resource_change.type])
}
Resource fields. Sometimes the resource type isn’t enough—you want to
auto-apply changes that only touch certain attributes. The change object in
the JSON plan let you diff individual fields. This policy denies any update that
modifies fields beyond tags :
deny contains msg if {
some resource_change in input.resource_changes
some action in resource_change.change.actions
action == "update"
changed_keys := {key |
some key in object.keys(resource_change.change.after)
resource_change.change.before[key] != resource_change.change.after[key]
}
changed_keys != {"tags", "tags_all"}
msg := sprintf("resource %q changes fields other than tags: %v", [resource_change.address, changed_keys])
}
Blast radius. A plan that touches 2 resources is different from one that
touches 200. You can count the resources with actual changes and gate when the
number exceeds a given threshold:
max_auto_apply_changes := 10
deny contains msg if {
changed := {resource_change.address |
some resource_change in input.resource_changes
some action in resource_change.change.actions
action not in {"no-op", "read"}
}
count(changed) > max_auto_apply_changes
msg := sprintf("plan affects %d resources, which exceeds the auto-apply limit of %d", [count(changed), max_auto_apply_changes])
}
Environment. Auto-apply in staging, gate in production. If your resources
are tagged with their environment, you can read that from the plan. This policy
denies any non-trivial change to a resource whose Environment tag is not
staging :
deny contains msg if {
some resource_change in input.resource_changes
some action in resource_change.change.actions
action not in {"no-op", "read"}
resource_change.change.after.tags.Environment != "staging"
msg := sprintf("resource %q is not in staging, requires human review", [resource_change.address])
}
These rules compose. You can combine them in the same policy file, and conftest
will evaluate all of them. A plan must pass every rule to auto-apply, and
any single denial is enough to fail the policy. The policy grows with your
confidence, and because it’s code, you can version it and test it like you do
with any other code.
A mechanism like this becomes ever more important as you introduce AI
agents to your SDLC , and let them propose and execute changes to your
live infrastructure. Without a deterministic way of attesting plan safety, you
either compromise on confidence, velocity, or both .
Did you find what you were looking for?
Let me know if you didn't.
Copyright © 2026 Ricard Bejarano
