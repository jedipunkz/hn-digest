---
source: "https://devgeist.com/blog/ai-infra-mutation/"
hn_url: "https://news.ycombinator.com/item?id=48992621"
title: "AI Agents Should Propose Infrastructure, Not Mutate It"
article_title: "AI Agents Should Propose Infrastructure, Not Mutate It"
author: "VMyroslav"
captured_at: "2026-07-21T14:57:38Z"
capture_tool: "hn-digest"
hn_id: 48992621
score: 1
comments: 0
posted_at: "2026-07-21T14:12:16Z"
tags:
  - hacker-news
  - translated
---

# AI Agents Should Propose Infrastructure, Not Mutate It

- HN: [48992621](https://news.ycombinator.com/item?id=48992621)
- Source: [devgeist.com](https://devgeist.com/blog/ai-infra-mutation/)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T14:12:16Z

## Translation

タイトル: AI エージェントはインフラストラクチャを変更するのではなく、提案する必要がある
説明: エージェントがクラウド認証情報を保持することで、10 年にわたるコードとしてのインフラストラクチャが元に戻される理由と、意図と効果の間にポリシーチェックされたゲートを置きながらエージェントの速度を維持する設計が無効になる理由。

記事本文:
レンダリングします。 --> AI エージェントはインフラストラクチャを変更するのではなく、提案する必要がある devgeist ブログ · エッセイ · 自動 / ライト / ダークを選択するドロップダウンについて。
JS を使用しないオープン/クローズを無料で提供します。以下のスクリプト
ワイヤの選択、永続化、および終了。アイコン + テキスト ラベル。
バーの右端、ナビリンクを過ぎたところにあります。 --> 自動
AI エージェントはインフラストラクチャを変更するのではなく、提案する必要がある
Myroslav Vivcharyk · 2026 年 7 月 18 日 · 14 分で読む
エージェントがクラウド認証情報を保有することで、10 年にわたるコードとしてのインフラストラクチャが元に戻される理由と、意図と効果の間にポリシーでチェックされたゲートを置きながらエージェントの速度を維持する設計。
現在、ほとんどの主要なクラウド ベンダーは、インフラストラクチャを変更できる MCP サーバーを出荷しています。これは構築が簡単で、デモも良好です。
また、10 年間にわたってコードとしてのインフラストラクチャが提供するために構築されてきたガバナンス特性も捨て去られますが、現時点ではそれが最も重要です。
おそらくインシデントレポートをすでにご覧になっていると思いますので、要約としてリンクのみを記載します。
Replit のエージェントがプロダクションを削除しました
データベース
明示的なコードのフリーズ中。
VS Code に同梱されている Amazon Q 拡張機能の侵害されたリリース
注射された
プロンプト
ローカル ファイル システムとクラウド リソースを消去するようにエージェントに指示します。
AWS CLI。ペイロードは実行されなかったと報告されています。エラーではなく構文エラーです。
コントロール、止めました。
1 つ目は、エージェントが自らの仕事を不適切に行っている場合です。 2 つ目は、攻撃者から与えられたジョブをエージェントが実行する場合です。
間違いは両方とも同じです。つまり、資格情報を使用した確率的な実行によって破壊的な権限に到達できるということです。
「エージェント XYZ、データベースを作成してください」は重要なデモ ケースではありません。それは
音声起動コンソール: 人間が意図を形成し、モデルがそれを
コマンド。何年も前にAlexaを使って構築できたかもしれません。
さらに興味深い質問は、2 日目の運用がどのようなものになるかです。

のように。
開発者はエージェントに新機能を求めます。エージェントが仕事を分割する
サブエージェント間で実行され、そのうちの 1 つがサービスにデータベースが必要であると判断します。
ユーザー、ACL、および特定の保持期間を持つトピック。または: 「クエリが遅い」 —
そして 3 つの推論ステップの後、エージェントはインスタンスが次のとおりであると結論付けます。
サイズが小さく、一段上のものになります。
これらのチェーンでは、インフラストラクチャの突然変異は推論の副作用です。
これは緊急事態です。トリガーとなったリクエストからいくつかのレイヤーが削除されました。
算術演算が壊れます。エージェントは生産コストを削減します
変化なので、どのエージェントでも突然変異量が増加します
組織が購入するスループット。レビュー能力は（まだ？）成長しません。
それでも人間であり、人間のリズムで。私たちが依存したガバナンス モデル
これら 2 つのレートが一致すると仮定します。変異量がレビューを上回る場合
容量が大きくなると、レビューがボトルネックになるか、現実的でなくなるかのどちらかです。
現在の共通チェーンは開発者のマシン上で実行されます。
エージェントのホストがベンダーの MCP サーバーをローカル プロセスとして起動します
そしてツールリストを取得します: create_database 、 update_acl 、
delete_topic — それぞれの引数に JSON スキーマが含まれます。
モデルはタスクの途中でツールを選択します。データベースが必要であると判断し、データベースを見つけます。
カタログ内の create_database に、引数を入力します。
どのような文脈であっても。
MCP サーバーは、呼び出しをクラウド API に変換します。
開発者の環境から継承した認証情報。 API
実行します。リソースは現在存在しています。
結果はツール結果としてモデルに返されます。
どの ID が変更を行ったかに注意してください: 開発者のものです。エージェントには何もありません
独自のプリンシパルであるため、監査証跡には人間が示され、IAM ルールは示されません。
エージェントの呼び出しと開発者の呼び出しを区別することもできます。ホストが求めるのは、
通話ごとの確認 — 人間が「」をクリックするまで

常に許可してください。」
AWS、Google、Azure は OAuth と IAM の背後でリモート MCP エンドポイントをホストするようになりました。エージェントは
独自の ID を持ち、すべての呼び出しが許可されます。このプリンシパルは
プロトコル、このリソースに触れてもよいか — すべてが監査に表示されます
ログ。リクエストレベルのガバナンスが存在します。
クラウド API を呼び出すエージェントは、エンジニアと同じようにインフラストラクチャを変更します。
2014 年にコンソールをクリックして実行しました。
このモデルは規模を維持できなかったため廃止されました。
コードとしてのインフラストラクチャは、宣言的でバージョン管理されたチェック可能な機能を導入しました。
意図と効果の間のアーティファクト。それらのプロパティは、
発信者はエージェントになっているため、負けることは許容されます。
権限を分離します。インフラを提案する権限
変化を起こすには、それを判断する権限と実行する権限が必要です。
さまざまなアクターに属しており、エージェントは最初のアクターのみを保持します。
これは機械に適用される構造設計です。
このように権力が分割されるのは、単一の権限が存在しないという古い理由によるものです。
俳優は常に正しいと信頼できるため、保証が適用されます
俳優の性格ではなく、構造にあります。
代替案では新しい機能は導入されません。同じガバナンス IaC です
すでに提供されており、高速かつ同時実行可能な呼び出し元用に再構築されました。
エージェントは API 呼び出しではなく、インテントを生成します。 「Postgres インスタンスが必要です
サービスXの場合。」
サーバー側レイヤーは、インテントを宣言的なアーティファクトに変換します。
Terraform/OpenTofu HCL — 組織の既存の HCL に基づく
不動産および承認されたモジュール。
成果物は現在の状態に対して計画され、計画はチェックされます
組織のポリシー (OPA) に反する。
拒否は交渉です。ポリシー違反がドラフトに戻ってくる
構造化された機械可読なヒントとしてステップを実行します。アーティファクトが再生される
合格するか、明らかにスタックするまで再送信されます。どれ

h
違反は自己修正可能であり、人間が必要とする違反は、
ポリシーのメタデータ。
提案する俳優と応募する俳優は別です。の
製図者はアーティファクトのみを作成できます。通過後はゲートのみ
チェックして、プランを実行して適用できます。
意図、生成されたアーティファクト、違反など、すべてのステップが監査されます。
修正、結果 — 追加専用のログ。
派生したインテントはトランスポート経由で到着します。
目的: サービス支払い - 偵察には実稼働 Postgres インスタンスが必要です
資産スナップショットとモジュールから作業する製図者
カタログ、最初のドラフトを発行します。
モジュール "payments_recon_db" {
ソース = "registry.internal/platform/postgres"
バージョン = "3.2.1"
サービス = "支払い-偵察"
環境 = 「生産」
インスタンスクラス = "小さい"
削除保護 = false
}
門はそれを否定しますが、その否定は次のように構造化されています。
[
{
"ポリシー" : "prod_deletion_protection" ,
"message" : "運用環境では deletion_protection が true である必要があります" ,
"self_fixable" : true 、
"fix_hint" : "deletion_protection = true を設定"
}、
{
"ポリシー" : "prod_instance_class" ,
"message" : "instance_class \" small \" は運用環境では許可されません" ,
"self_fixable" : true 、
"fix_hint" : "中、大のいずれかを使用します"
}
】
どちらの違反も self_fixable: true — ポリシーで宣言されています。
メタデータなので、ループは
アーティファクト: deletion_protection = true 、instance_class = "medium" 、
再提出してください。合格すると、計画が作成およびチェックされ、実行が適用されます。
ここで、同じインテントが別のルールをトリップすると仮定します。
{
"ポリシー" : "データ常駐" ,
"message" : "新しい運用データ ストアの常駐分類はありません" ,
"self_fixable" : false 、
"info_needed" : "サービス所有者からのデータ常駐分類"
}
self_fixable: false ループを終了します: ジョブは次のようにパークされます。
need_input とエスカレーションします。いくら回復してもダメ

これは×です、
なぜなら、欠落している事実はマシンのコンテキストにないからです。
どの違反が収束し、どの違反がエスカレートするかは、いつ決定されたか
方針が書かれていました。
写真で注意すべき点がいくつかあります。
2 つのトラストゾーン。起草者は何もないサンドボックスから提案します。
資格情報。企画・申請段階のみ 通過点の向こう側
チェック — 雲に触れることができます。
ドアが 1 つ。ゲートは複数の突然変異パスの中の 1 つではなく、側面もありません
1 つのエージェントが CLI を保持し、別のエージェントが独自の MCP サーバーを使用するチャネル。
組織レベルのポリシーにより、すべてのプリンシパルへの状態変更呼び出しが拒否されます
適用段階を除く: エージェント、開発者、CI も同様です。
プロバイダーは運用知識をエンコードします。ほとんどの Terraform プロバイダー
API の癖、つまり再試行セマンティクス、依存関係の順序付け、
最終的な整合性、冪等性、部分的な障害のクリーンアップ。エージェント
クラウド API を呼び出すと、実行時にすべてが再検出されます。 HCL を発行するエージェントはそれを継承します。
モジュールは生成を語彙に変えます。プラットフォームチームは、
承認されたモジュール、その構成要素、および起草者の仕事
「型指定されたパラメータを使用して承認されたモジュールをインスタンス化する」に絞り込みます。
影響範囲が小さく、ポリシーだけでは表現できない好みの一部をコード化する場所です。
既存の不動産が文脈であり、それが複雑になります。ほとんど
組織にはすでに HCL リポジトリがあり、何が存在するかを説明する状態: 名前付け
慣例、既存のネットワーク、使用パターンなど。
そして、レイヤーは読み取ったものをフィードします。レイヤーが適用するすべてのミューテーションが返されます。
不動産の中で、それを求めた意図に結びついています。インフラストラクチャがループを通過するほど、より豊かになります
製図者の世界が広がります。直接 API エージェントは同じリポジトリを読み取ることができます。
しかし、書き込みがそこに到達することはありません。それが引き起こすすべての突然変異は外側に落ちます
宣言された不動産、つまりそれが読み取るコンテキスト

それぞれ少しずつ間違っています
時間。 1 つのループがそのコンテキストを複雑にします。もう一方はそれを侵食します。
起草者は状態を見ることはありません。 HCL 宣言のみが表示されます。
写真のポリシー ゲートは 1 つのボックスですが、2 つの異なる質問に答えるために 2 つの異なるチェックを実行します。
最初のチェックは、何かが実行される前に、生成された HCL 自体に対して実行されます。
Terraform は実行エンジンです: init ダウンロード プロバイダー
プラグイン (任意のバイナリ) とプランがそれらを実行します。データソースの取得
計画中にネットワーク経由で実行され、外部データ ソースが実行されます。
ローカルコマンドの設定名は何でも構いません。悪いアーティファクトはそうではありません
ダメージを与えるには到達する必要があります。計画は十分です。とても静的です
構成ポリシーは、許可されたプロバイダーとそのバージョンを固定します。
モジュールとプロバイダーを承認されたソースに制限し、エスケープを禁止します
ハッチ。アーティファクトが安全に実行できるかどうかを確認します。
次に、アーティファクトは初期化とプランを取得します。
読み取り専用の資格情報を使用し、プラン自体に対して 2 番目のチェックが実行されます。
結果として生じるリソース、削除と置換、コスト、リージョン、
暗号化、クォータ、および変更がすでに設定されている資産にどのように関連するか
存在します。このチェックでは、「この変更が準拠しているかどうか」という質問に答えます。
これは、計画する宣言的な成果物がある場合にのみ存在できます。
合格した正確な保存された計画が適用されます。
保存された計画は、それが計算された状態のバージョンにバインドされているため、
別の適用が最初に実行されると、古い計画は拒否され、再生成されます。
移動した州に対しては決して適用されません。
すべてのエージェント変更には 3 つの層があります。タスクの目的: 「機能 X を実装する」 —
人間が言ったもの。派生インフラストラクチャの意図: 「機能 X にはトピックが必要です」
およびデータベース ユーザー」 - エージェントによって推測されます。具体的な突然変異：これ
これらのパラメータを含むモジュール。ゲートの検証

は
ポリシーに反する第 3 層。 2番目がそうであったことを証明することはできません。
最初の解釈を正しく行うと、インフラストラクチャ レベルの決定は行われませんでした。
チェックするために記載されています。
これが正確な制限です。ゲートは意図の忠実性ではなく、コンプライアンスを保証します。
計画はすべてのルールを満たしていても、間違った計画である可能性があります。
これらのどれもが HCL を神聖なものにするものではありません。 Pulumi または別の後継者が代替となる可能性があります。
重要なのは、宣言型 IaC エコシステムでは、
業界はすでに運用知識を蓄えており、ガード層は
その残高を再構築する代わりに費やします。
この設計は、望ましい状態の突然変異を制御します。異常なサービスを再起動してフェイルオーバーを強制し、
一時的な認証情報のローテーションとプッシュはアクションです
計画サイクルを通して演劇が行われます。同じ不変式が必要です。
サーバー側のチェックは誰もバイパスしません。
これらのアイデアのプロトタイプを作成しているときに、もちろん、私はそうではないことに気づきました。
一人で。 AARM — 自律アクション ランタイム
Management 、クラウド セキュリティ アライアンス プロジェクト
エージェントアクションのランタイムガバナンスを定義します: 実行前のインターセプト、
ID にバインドされ、ポリシーに対してチェックされます。この記事の前提が当てはまるなら
あなたにとって、この仕様は読む価値があります。
1. 「これは MCP 内の単なる CI ではありませんか? エージェントに次のようなプル リクエストを開いてもらいます
他の全員 - パイプラインは通常どおり実行されます。」
はい、それが最も重要な点です。
ビルドします

[切り捨てられた]

## Original Extract

Why agents holding cloud credentials undo a decade of infrastructure-as-code — and a design that keeps agent speed while putting a policy-checked gate between intent and effect.

renders. --> AI Agents Should Propose Infrastructure, Not Mutate It devgeist Blog · Essays · About dropdown to pick auto / light / dark.
gives a no-JS open/close for free; the script below
wires selection, persistence, and closing. Icons + text labels.
Sits at the far right of the bar, past the nav links. --> Auto
AI Agents Should Propose Infrastructure, Not Mutate It
Myroslav Vivcharyk · July 18, 2026 · 14 min read
Why agents holding cloud credentials undo a decade of infrastructure-as-code — and a design that keeps agent speed while putting a policy-checked gate between intent and effect.
Most major cloud vendors now ship an MCP server that can alter your infrastructure. This is easy to build and it demos well.
It also throws away the governance properties that a decade of infrastructure-as-code was built to provide, at the moment they matter most.
You have probably seen the incident reports already, so links only, as a recap:
Replit’s agent deleted a production
database
during an explicit code freeze.
A compromised release of the Amazon Q extension for VS Code shipped with
an injected
prompt
telling the agent to wipe local filesystems and cloud resources through
the AWS CLI. The payload reportedly never ran — a syntax error, not a
control, stopped it.
The first is an agent doing its own job badly. The second is an agent doing a job an attacker gave it.
The mistake is the same in both: destructive authority reachable from probabilistic execution, with credentials.
“Agent XYZ, create me a database” is not the demo case that matters. That’s a
voice-activated console: a human forms the intent, a model transforms it into the
command. You could have built it with Alexa many years ago.
The more interesting question is what day-2 operation looks like.
A developer asks an agent for a new feature. The agent splits the work
across sub-agents, and one of them decides the service needs a database
user, ACLs, and a topic with specific retention. Or: “my queries are slow” —
and three reasoning steps later, an agent concludes the instance is
undersized and bumps it a tier.
In these chains, the infrastructure mutation is a side effect of reasoning .
It’s emergent — several layers removed from the request that triggered it.
The arithmetic breaks . Agents reduce the cost of producing
a change, so mutation volume grows with whatever agent
throughput an organization buys. Review capacity doesn’t grow with it (yet?): it is
still humans, at human cadence. The governance models we relied on
assume those two rates are matched. When mutation volume outruns review
capacity, the review either becomes the bottleneck or stops being real.
The common chain today runs on a developer’s machine:
The agent’s host launches the vendor’s MCP server as a local process
and gets the tools list: create_database , update_acl ,
delete_topic — each with a JSON schema for its arguments.
The model picks a tool mid-task: it decides it needs a database, finds
create_database in the catalog, and fills in the arguments from
whatever context it has.
The MCP server translates the call to the cloud API — using the
credentials it inherited from the developer’s environment. The API
executes. The resource now exists.
The result returns to the model as a tool result.
Note which identity did the mutating: the developer’s. The agent has no
principal of its own, so the audit trail shows a human, and no IAM rule
can even tell the agent’s calls from the developer’s. The host asks for
per-call confirmation — until the human clicks “always allow.”
AWS, Google and Azure now host remote MCP endpoints behind OAuth and IAM: the agent gets
its own identity, every call is authorized — may this principal use the
protocol, may it touch this resource — and everything appears in the audit
log. Request-level governance exists.
An agent calling cloud APIs mutates infrastructure the way an engineer
clicking through a console did in 2014.
That model was retired because it didn’t survive scale.
Infrastructure-as-code introduced a declarative, versioned, checkable
artifact between intent and effect. Those properties did not become
acceptable to lose because the caller is now an agent.
Separate the powers. The authority to propose an infrastructure
change, the authority to judge it, and the authority to execute it must
belong to different actors — and the agent holds only the first.
This is constitutional design, applied to machines.
Power is split this way for the old reason: no single
actor can be trusted to be right every time, so the guarantee is placed
in the structure, not in the actor’s character.
The alternative introduces no new capability. It’s the same governance IaC
already gave us, rebuilt for a caller that is fast and concurrent:
The agent produces intent, not API calls. “I need a Postgres instance
for service X.”
A server-side layer turns the intent into a declarative artifact —
Terraform/OpenTofu HCL — based on the organization’s existing
estate and approved modules.
The artifact is planned against current state, and the plan is checked
against organizational policies (OPA) .
Denial is a negotiation. Policy violations come back to the drafting
step as structured, machine-readable hints. The artifact is regenerated
and resubmitted until it passes or gets detectably stuck. Which
violations are self-fixable and which need a human is declared in the
policy’s metadata.
The actor that proposes and the actor that applies are separate. The
drafter can only produce an artifact. Only the gate, after a passing
check, can run the plan and apply it.
Every step is audited — intent, generated artifact, violations,
fixes, outcome — append-only log.
The derived intent arrives over the transport:
intent: service payments-recon needs a production Postgres instance
The drafter, working from the estate snapshot and the module
catalog, emits a first draft:
module "payments_recon_db" {
source = "registry.internal/platform/postgres"
version = "3.2.1"
service = "payments-recon"
environment = "production"
instance_class = "small"
deletion_protection = false
}
The gate denies it — but the denial is structured:
[
{
"policy" : "prod_deletion_protection" ,
"message" : "deletion_protection must be true in production" ,
"self_fixable" : true ,
"fix_hint" : "set deletion_protection = true"
},
{
"policy" : "prod_instance_class" ,
"message" : "instance_class \" small \" is not allowed in production" ,
"self_fixable" : true ,
"fix_hint" : "use one of: medium, large"
}
]
Both violations carry self_fixable: true — declared in the policy’s
metadata, so the loop regenerates the
artifact: deletion_protection = true , instance_class = "medium" ,
resubmit. It passes, the plan is produced and checked, apply runs.
Now suppose the same intent trips a different rule:
{
"policy" : "data_residency" ,
"message" : "no residency classification for a new production data store" ,
"self_fixable" : false ,
"info_needed" : "data residency classification from the service owner"
}
self_fixable: false ends the loop: the job parks as
needs_input and escalates. No amount of regeneration fixes this one,
because the missing fact isn’t in the machine’s context.
Which violations converge and which escalate was decided when
the policy was written.
A few things to notice in the picture.
Two trust zones. The drafter proposes from a sandbox with no
credentials. Only the plan/apply stage — on the far side of a passing
check — can touch the cloud.
One door. The gate is not one mutation path among several — no side
channel where one agent holds a CLI and another uses its own MCP server.
Organization-level policy denies state-changing calls to every principal
except the apply stage: agents, developers, and CI alike.
Providers encode the operational knowledge. Most Terraform providers
are years of handling API quirks: retry semantics, dependency ordering,
eventual consistency, idempotency, partial-failure cleanup. An agent
calling cloud APIs rediscovers all of it at runtime. An agent emitting HCL inherits it.
Modules turn generation into vocabulary. A platform team has
approved modules — its building blocks — and the drafter’s job
narrows to “instantiate an approved module with typed parameters.”
Smaller blast radius, and a place to encode some of your taste that policy alone can’t express.
The existing estate is the context — and it compounds. Most
organizations already have HCL repos and state describing what exists: naming
conventions, existing networks, patterns in use, etc.
And the layer feeds what it reads: every mutation it applies lands back
in the estate, tied to the intent that asked for it. The more infrastructure goes through the loop, the richer
the drafter’s world gets. A direct-API agent can read the same repo —
but its writes never land there. Every mutation it makes falls outside
the declared estate, so the context it read is a little more wrong each
time. One loop compounds its context; the other erodes it.
The drafter never sees state. It sees only the HCL declarations.
The policy gate in the picture is one box, but it runs two different checks to answer two different questions.
The first check runs on the generated HCL itself, before anything executes.
Terraform is an execution engine: init downloads provider
plugins — arbitrary binaries — and plan runs them; data sources fetch
over the network during planning, and an external data source executes
whatever local command the configuration names. A bad artifact doesn’t
need to reach apply to do damage — plan is enough. So static
configuration policy pins the allowed providers and their versions,
restricts modules and providers to approved sources, and bans the escape
hatches. We check the artifact is even safe to run.
The artifact then gets an init and a plan under
read-only credentials, and the second check runs on the plan itself: the
resources that would result, the deletions and replacements, cost, regions,
encryption, quotas, and how the change relates to the estate that already
exists. This check answers “does this change comply” — the one
that can only exist because there is a declarative artifact to plan.
The exact saved plan that passed gets applied.
A saved plan is bound to the state version it was computed against, so if
another apply lands first, the stale plan is rejected and regenerated —
never applied against a state that has moved.
Every agentic change has three layers. Task intent: “implement feature X” —
stated by a human. Derived infrastructure intent: “feature X needs a topic
and a database user” — inferred by an agent. Concrete mutation: this
module with these parameters. The gate validates the
third layer against policy. It cannot certify that the second was the
right reading of the first — no infrastructure-level decision was ever
stated for it to check against.
That is the precise limitation: the gate guarantees compliance, not intent-fidelity.
A plan can pass every rule and still be the wrong plan.
None of this makes HCL sacred. Pulumi or another successor could substitute.
The point is that the declarative-IaC ecosystem is where the
industry already banked its operational knowledge, and the guard layer
spends that balance instead of rebuilding it.
This design governs desired-state mutations. Restarting an unhealthy service, forcing a failover,
rotating an ephemeral credential are actions, and pushing them
through a plan cycle would be theater. They need the same invariant — a
server-side check nobody bypasses.
While prototyping these ideas, I discovered — of course — that I was not
alone. AARM — Autonomous Action Runtime
Management , a Cloud Security Alliance project that
defines runtime governance of agent actions: pre-execution interception,
bound to identity, checked against policy. If this article’s premise holds
for you, the spec is worth reading.
1. “Isn’t this just CI inside MCP? Have the agent open a pull request like
everyone else — the pipeline runs as usual.”
Yes — and that’s most of the point.
Build t

[truncated]
