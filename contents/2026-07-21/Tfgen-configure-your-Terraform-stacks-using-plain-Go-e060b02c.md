---
source: "https://ably.com/blog/cdktf-migration-tfgen"
hn_url: "https://news.ycombinator.com/item?id=48990838"
title: "Tfgen: configure your Terraform stacks using plain Go"
article_title: ""
author: "surminus"
captured_at: "2026-07-21T12:21:07Z"
capture_tool: "hn-digest"
hn_id: 48990838
score: 1
comments: 0
posted_at: "2026-07-21T11:29:03Z"
tags:
  - hacker-news
  - translated
---

# Tfgen: configure your Terraform stacks using plain Go

- HN: [48990838](https://news.ycombinator.com/item?id=48990838)
- Source: [ably.com](https://ably.com/blog/cdktf-migration-tfgen)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T11:29:03Z

## Translation

タイトル: Tfgen: プレーン Go を使用して Terraform スタックを構成する
説明: CDKTF は廃止されつつあります。私たちは、オープンソースの Go 代替ツールである tfgen を構築し、AI を使用して 2 週間で 150 の Terraform ワークスペースを移行しました。

記事本文:
お問い合わせ ログイン 無料で始める このページで ← ブログ Ably Engineering tfgen の紹介: プレーン Go を使用して Terraform スタックを構成する
HashiCorp のサンセット CDKTF は代替品が見当たりません。ここでは、独自の Go 代替手段である tfgen を構築し、AI を使用して 150 の実稼働 Terraform ワークスペースを数か月ではなく 2 週間で移行した方法を説明します。
最近まで、私たちは Terraform 用の HashiCorp の CDK を広範囲に使用していました。その後、彼らは開発の終了を発表しました。新しいエコシステムへの移行に何か月も費やしたくはありませんでしたが、代わりのエコシステムが必要でした。
Ably のインフラストラクチャ チームでは Go が好きです。もっとシンプルなものを作る機会はないだろうかと考えました。
Ably でのインフラストラクチャのプロビジョニング
可能な場合は、Infrastructure as Code を使用して、コア Pub/Sub プラットフォーム用のすべてのインフラストラクチャをプロビジョニングします。最初は Terraform を使用し、次に OpenTofu に切り替えます (ただし、tofu を使用している場合でも、簡単にするために、全体で「Terraform」と呼びます)。当社のインフラストラクチャに対するすべての運用上の変更はピアレビューを経る必要があり、Scalr を使用してデプロイされます。
これには、アプリケーションのワークロードとは別のアカウントで、非本番環境と本番環境の両方の複数のリージョンにプロビジョニングされた一連の Amazon VPC であるネットワーキング スタックが含まれます。
各 VPC はセキュリティ グループ ルールや NAT ゲートウェイなどのリソースとともに他のすべての VPC とピアリングする必要があるため、HCL でこれを完全に構成するのはすぐに困難になってきました。
この文脈で、「CDKTF」としても知られる「 terraform-cdk 」を採用しました。このツールは HCL の制限から抜け出すための自然な方法であり、エンジニアは Terraform 互換の JSON を生成するコードをさまざまな言語で書くことができました。 TypeScriptを使用してこれを採用しました。
私たちはこのツールを使用して大きな成功を収めました。

これは、多くの繰り返しロジックを含む機能の一部にこれを採用しましたが、別のワークスペース、別のリージョン、別のアカウントにリソースをプロビジョニングする必要がありました。 CDKTF を使用するとこれが非常に簡単になり、人に優しい構成ファイルを使用して新しいリソースをプロビジョニングしました。
生成された JSON はリポジトリにコミットされ、標準プロセスを使用してデプロイされました。誰が構成ファイルにどのような変更を加えたかを確認でき、それがより複雑な JSON 出力で明確に表現されていることがわかります。
このメカニズムを使用して、約 150 の本番ワークスペースを維持しました。
残念ながら、2025 年 12 月 10 日に、HashiCorp は CDKTF を廃止することを決定しました。
このため、私たちは少し窮地に陥りました。はい、ツールは正常に動作していましたが、新機能はなく、未解決のバグは修正されません。
私たちは、代替となる可能性のあるものについていくつかの調査を行いました。私たちは、非常にうまく機能する既存のプロセスを維持したいと考えていました。また、すべての重要な生産リソースに影響を与える長期にわたる移行にはコミットしたくありませんでした。これにより、リスクが増大し、比較的価値の低い多くのエンジニアリング作業が必要となります。
他のさまざまなツール ( Terragrunt 、 Pulumi 、 Jsonnet 、 Cue ) を検討しました。
Terragrunt は広く使用されており、多くのプロジェクトを大規模に管理できるため、私たちにとってはうまく機能しますが、このモデルに適合できるようにするには、ワークフローを変更する必要があり、これがおそらく移行の大部分を占めていたでしょう。
一元化された構成ファイルを引き続き使用したいと考えた場合、それらを維持するか、別のものに移行するには、とにかくカスタムのものを作成する必要がありました。基本的に、それは膨大な作業量だったでしょう。
Pulumi も同じ問題に悩まされていました。これは素晴らしいツールですが、完全に新しいエコシステムへの移行を意味していました。
テンプレート言語

Jsonnet と Cue の量はおそらく必要なものに近かったのですが、すでにリポジトリにコミットしたものとバイト同一の出力を確実に得るのは大変な作業だったでしょう。
本質的に、これらすべてのツールは同じ問題、つまり、私たちが引き受ける準備ができていない膨大な量の作業を抱えていました。最終的には、Go を使用して自分たちで何かを書くことに落ち着きました。
私たちはすでに主にインフラストラクチャの管理に Go を使用しているため、ホームスパンの代替品として Go を選択するのは自然なことでした。
既存の構成を使用し、同一の出力を既存のパスに書き込むドロップイン置換を作成するだけでなく、CDKTF で以前に遭遇した制約も緩和したいと考えていました。
CDKTF では、強力な型を生成した特定のプロバイダーのバージョンを「取得」する必要がありました。これにより、LSP を使用した TypeScript の作成が少し簡単になりましたが、これらの定義をフェッチするためのビルド ステップも必要になりました。実際には、より強力なタイプが追加の構築ステップのコストに見合う価値があるとは思いませんでした。 Terraform ツール自体によって検証されるコードを生成する、より柔軟なものが必要でした。
私のアイデアは、Terraform の概念に一致する、各コンポーネントが型である単純なパブリック インターフェイスを作成することでした。
// Resource は、Terraform 構成内のリソース ブロックを表します。
type リソース構造体 {
Type string // リソースタイプ、例: 「aws_s3_バケット」
名前文字列 // ローカル名、スタック内のタイプごとに一意
Config Config // リソースの引数
}
ここでの Config タイプにより、柔軟性が高まります。異なるプロバイダーごとに大量の型を生成し、バージョン間で型がどのように変化するかという問題に対処するのではなく、ユーザーは必要な値を設定する責任を負います。
Terraform 互換の JSON が生成されると、ユーザーは val を使用できるようになります。

idate を使用して、生成された型の要件をバイパスして、それが正しいことを確認します。
これは、HCL を書く経験を与えながら、Go を使用して自然に自分自身を表現できることを意味します。
CDKTF の言語を模倣して、各コンポーネントが「スタック」に追加されます。
スタック := tfgen.NewStack( "例" )
バケット := &tfgen.Resource{
入力: "aws_s3_bucket" 、
名前: "例" 、
構成: tfgen.Config{
"バケット" : "私の例のバケット" ,
}、
}
stack.AddResource(バケット)
その後、このスタックを任意の場所に書き込むことができます。
f, _ := os.Create( "main.tf.json" )
stack.Write(f)
結果の JSON は Terraform と完全に互換性があり、tofu apply で適用できます。
{
「リソース」: {
"aws_s3_bucket" : {
「例」: {
"バケット" : "私の例のバケット"
}
}
}
}
AIを活用した移行
2026 年のほとんどのエンジニアと同様に、私も AI を使用して多くのコードを作成しています。
何かをゼロから設計する場合、AI ツールには当たり外れがあることがわかります。実際、このパッケージを作成するためにいくつかのアイデアを AI から反映させましたが、ひどい結果でした。
これが、私がパブリック インターフェイスを事前に自分で作成した理由です。私はパッケージをどのように使用したいかを正確に知っていたため、結果を制御することができました。
AI で重要なのは、AI が機能できる強力なサンプルを用意することです。これにより、AI が行う必要がある設計上の決定の数が減ります。
バイト同一の出力を必要とする既存のパスに書き込んでいたため、AI ツール (この場合はクロード コード) にも明確なフィードバック ループがあることを意味しました。
強く型付けされたパブリック インターフェイス、人間が判読できるコードがどのようなものであるかを示す明確な例、そして私たちが求めていた正確な出力 (git diff を使用して簡単に確認できる) があれば、AI が短期間で移行を完了してくれると確信していました。
私は、それぞれに別のエージェントを使用して、移行を段階的に実行するように依頼しました。
研究: 既存の CDKTF コードを読み取り、

既存の出力、および構成ファイルの形式。これにより、これらの特定のワークスペースにどのロジックをコピーする必要があるか、およびそれらをどこに記述する必要があるかがまとめられました。
計画: このエージェントは、前の出力を使用して、次のような特定のフィードバック ループと各ステップでのテストを含む、実装のための明確で順序付けられた計画を作成しました。
go build、go test、go vet などの検証ツールを実行する
JSON の生成後に git diff を使用して、変更がゼロであること、または影響のないもののみであることを確認します (CDKTF は削除しても問題ないメタデータ フィールドを書き込みました)
実装: 最終エージェントは計画を実行し、検証も実行しました。
いくつかの異なるワークスペースのセットがあり、それぞれを順番に作業しました。次のセットに進む前に、本番環境の一連のワークスペースを完全に移行しました。これにより、生成される変更に対する信頼性が高まりました。
以前は手書きで書かなければならなかったとしたら、何か月もかかっていたと思います。実際の結果の値がツールの変更を除いて変更がゼロだったことを考えると、ほぼ確実にバックログの最後に置かれていたでしょう。
幸いなことに、AI を使用することで、実際にすべてのワークスペースの移行を数週間で完了することができました。それだけでなく、私は他の日常業務と並行してこの作業を行っていました。なぜなら、私が割り当てられた緊急の仕事をしている間、バックグラウンドで AI を動作させることができたからです。
このプロジェクトの成功には少々驚きましたが、AI ツールを使用するメリットがはっきりと分かりました。
CDKTF に代わる明確なツールが存在しないことを考えると、コミュニティの他の人々がこのツールを便利だと思うかもしれないと思い、この理由から、私たちは tfgen をオープンソースにすることにしました。
これは github.com/ively/tfgen で見つけることができます。
コア パッケージ全体はわずか約 600 行の Go であり、必要なものがすべて含まれています。

完全な Terraform 構成とプロバイダー ヘルパーを生成します。たとえば、Ably Terraform プロバイダー用に小さなリソース セットを提供しました。
stack := tfgen.NewStack( "ively" )
ブロック := &tfgen.TerraformBlock{必須バージョン: ">= 1.9" }
block.AddRequiredProvider( "ively" , ively.ProviderVersion())
stack.SetTerraformBlock(ブロック)
stack.AddProvider(ively.Provider())
app := ively.App( "main" , ively.AppOptions{TLSOnly: true })
stack.AddResource(アプリ)
var buf bytes.Buffer
エラーの場合:= stack.Write(&buf);エラー != nil {
パニック（えー）
}
fmt.Print(buf.String())
生成された型がないため、リソース間の参照は Ref メソッドを使用して行われ、生成された JSON で Terraform 式 (${ively_app.main.id} など) としてレンダリングされます。
stack.AddResource(ively.Namespace( "ai-transport" , ively.NamespaceOptions{
AppID: app.Ref( "id" ),
MutableMessages: true 、
}))
また、Ably では複数のスタックを記述するという概念も組み込まれており、それが Ably での使用方法です。
stack1 := tfgen.NewStack( "スタック1" )
stack2 := tfgen.NewStack( "スタック2" )
mystacks := tfgen.Stacks{スタック1, スタック2}
mystacks.Write( "my/stacks/ディレクトリ" )
よくある質問
CDKTF とは何ですか?なぜ日没になったのですか? CDKTF (terraform-cdk) を使用すると、エンジニアは TypeScript、Python、Go などの汎用言語でインフラストラクチャ コードを記述し、Terraform 互換の JSON にコンパイルできます。 HashiCorp は 2025 年 12 月に、新機能やバグ修正を行わずにツールの開発を中止すると発表しました。
tfgen は Terraform だけでなく OpenTofu とも互換性がありますか?はい。 tfgen は標準の Terraform 互換 JSON を生成するため、構成の記述方法を変更することなく、Terraform と OpenTofu の両方で動作します。
tfgen は CDKTF とどう違うのですか? tfgen はプロバイダー固有の型を生成しません。フェッチしてメンテナンスする代わりに

厳密に型指定されたプロバイダー定義では、構成値を直接設定し、Terraform 独自のツールを使用して出力を検証します。これにより、必要な追加のビルド ステップ CDKTF が削除されます。
Ably 以外のプロバイダーに tfgen を使用できますか?はい。コア パッケージはプロバイダーに依存しません。 Ably には、実例として Ably Terraform プロバイダー用の小さなヘルパー セットが含まれていますが、tfgen の Resource および Config タイプはどの Terraform プロバイダーでも機能します。
すべての Terraform リソース タイプに対して Go コードを記述する必要がありますか?いいえ。tfgen はリソースごとに生成された型を必要としないため、プロバイダー固有のインストールや同期を維持するものは何もありません。リソースは、汎用リソース タイプとその構成マップを使用して表現します。
便利だと思われ、機能のリクエストや質問がある場合は、ご連絡いただくか、GitHub の問題を提起してください。
信頼性の高いリアルタイムを数分で構築します。
Ably は、私たちが記事にしているのと同じパターンを本番環境で大規模に実行します。無料枠には毎月 600 万のメッセージが含まれており、カードは必要ありません。
Ably チームからの新しい投稿 (毎月)。
Ably でのインフラストラクチャのプロビジョニング
Twitter / X LinkedIn Y ハッカー ニュース リンクをコピー 続きを読む
Ably に直接接続する方法 (そしておそらくそうすべきではない理由) – パート 2
VPC ピアリングとトランス

[切り捨てられた]

## Original Extract

CDKTF is sunsetting. We built tfgen, an open source Go alternative, and used AI to migrate 150 Terraform workspaces in two weeks.

Contact us Login Start free On this page ← Blog Ably engineering Introducing tfgen: configure your Terraform stacks using plain Go
HashiCorp sunset CDKTF with no replacement in sight. Here's how we built our own Go alternative, tfgen, and used AI to migrate 150 production Terraform workspaces in two weeks instead of months.
Until recently, we extensively used HashiCorp's CDK for Terraform. Then they announced the end of its development. We didn't want to spend months migrating to a new ecosystem, but we needed a replacement.
In the Infrastructure Team at Ably, we like Go. I wondered if there was an opportunity for something simpler.
Provisioning infrastructure at Ably
We provision all our infrastructure for the core Pub/Sub platform where possible using Infrastructure as Code , initially using Terraform and then making the switch to OpenTofu (but for simplicity I will refer to "Terraform" throughout, even though we use tofu). Every production change to our infrastructure has to go through peer review and is deployed using Scalr .
This includes our networking stack, which is a set of Amazon VPCs provisioned across multiple regions for both our non-production and production environments, in accounts separate from our application workloads.
Since each VPC must be peered with every other VPC along with resources such as security group rules and NAT gateways, configuring this entirely in HCL became unwieldy very quickly.
It was in this context that we adopted the " terraform-cdk ", also known as "CDKTF". This tool was a natural way to get out of HCL's limitations, where engineers were able to write code in a variety of languages which produced Terraform-compatible JSON. We adopted this using TypeScript.
We had a lot of success using this tool, and because of this adopted it for some of our functionality that had a lot of repeated logic, but required us to provision resources in separate workspaces, in separate regions, and in separate accounts. Using CDKTF made this extremely simple, and we provisioned new resources using human-friendly configuration files.
The resulting generated JSON was committed to our repository, and deployed using our standard processes. We could see who made what changes to a configuration file, and see it clearly expressed in the more complex JSON output.
We maintained roughly 150 production workspaces using this mechanism.
Unfortunately, on December 10, 2025, HashiCorp decided to sunset CDKTF .
This put us in a bit of a bind: yes, the tool was working fine, but there will be no new features, and any open bugs will never be fixed.
We did some research on a potential replacement: we wanted to keep the process we have that works really well for us, and we didn't want to commit to a lengthy migration that touches all of our critical production resources, which increases risk and takes a lot of engineering effort with relatively low value.
We considered various other tools: Terragrunt , Pulumi , Jsonnet and Cue .
Terragrunt is widely used and allows you to manage many projects at scale, which would work for us, but to be able to fit into this model, we would have needed to change our workflows, which probably would have been the bulk of the migration.
Given we wanted to continue using our centralised configuration files, we would have needed to write something custom anyway to keep them, or migrate to something else. Basically, it would have been a huge amount of work.
Pulumi suffered the same problem: it's a great tool, but would have meant a migration to a completely new ecosystem.
The templating languages of Jsonnet and Cue were probably closer to what we needed, but it would have been a great deal of work to ensure we had byte-identical output to what we already had committed to the repository.
Essentially, all these tools suffered the same issue: a huge amount of work that we were not prepared to undertake. In the end we settled on writing something ourselves, using Go.
We already primarily use Go for managing our infrastructure, so it was the natural choice for a homespun replacement.
As well as writing a drop-in replacement, which used our existing configuration and wrote identical output to existing paths, I also wanted to ease the constraints that we had previously bumped into with CDKTF.
CDKTF required us to "get" specific provider versions, which had generated strong types. This made writing TypeScript a little easier with LSPs, but also involved a build step to fetch these definitions. In practice, I didn't think the stronger types were worth the cost of the extra build step. I wanted something more flexible, that produced code that was validated by Terraform tooling itself.
My idea was to produce a simple public interface that matched Terraform concepts, where each component was a type:
// Resource represents a resource block in a Terraform configuration.
type Resource struct {
Type string // resource type, e.g. "aws_s3_bucket"
Name string // local name, unique per type within the stack
Config Config // the resource's arguments
}
The Config type here is what makes it flexible. Rather than having to generate a large amount of types for each different provider and handle the problems of how they change between versions, instead the user is responsible for setting the required values.
When the Terraform-compatible JSON is generated, then the user can use validate to ensure it's correct, bypassing the requirement for generated types.
This means that we can express ourselves naturally using Go, while giving us the experience of writing HCL.
Mimicking CDKTF's language, each component is then appended to a "stack":
stack := tfgen.NewStack( "example" )
bucket := &tfgen.Resource{
Type: "aws_s3_bucket" ,
Name: "example" ,
Config: tfgen.Config{
"bucket" : "my-example-bucket" ,
},
}
stack.AddResource(bucket)
You can then write this stack to wherever you want:
f, _ := os.Create( "main.tf.json" )
stack.Write(f)
The resulting JSON is fully Terraform compatible and can be applied with tofu apply:
{
"resource" : {
"aws_s3_bucket" : {
"example" : {
"bucket" : "my-example-bucket"
}
}
}
}
Migrating using AI
Like most engineers in 2026, I am producing a lot of my code using AI.
I can find AI tooling hit and miss when designing something from scratch. In fact, I did bounce some ideas off AI for writing this package and it did a terrible job.
This is why I wrote the public interface myself upfront: I knew exactly how I would like to use the package, which put me in control of the outcome.
The key with AI is to have strong examples that it can work from, which reduces the number of design decisions it has to make.
Since we were writing to an existing path requiring byte-identical output, it meant that my AI tool (in this case, Claude Code) also had a clear feedback loop.
With a strongly typed public interface, a clear example of what I want the human-readable code to look like, and the exact output we were after, which could easily be checked using git diff, I was confident that AI would make short work of the migration.
I asked it to perform the migrations in stages, using a separate agent for each:
Research: read the existing CDKTF code, the existing output, and the format of the configuration files. This produced a summary of what logic we needed to copy for these particular workspaces and where they needed to be written.
Planning: this agent used the previous output to create a clear, ordered plan for implementation, including specific feedback loops and tests at each step, such as:
Running go build, go test and validation tools such as go vet
Using git diff after generating the JSON to ensure zero changes, or only non-impacting ones (CDKTF wrote a metadata field that was fine to drop)
Implementation: the final agent worked through the plan, and also performed validation.
We had several differing sets of workspaces, and I worked through each in turn. I completely migrated a set of workspaces in production before moving on to the next set, which increased confidence in the changes it was producing.
I think in the past if I had to write it by hand, it would have taken me months. It almost certainly would have been put to the bottom of our backlog, given the actual resulting value was zero changes, other than a change of tooling.
Fortunately, using AI meant that I actually completed the migration of all our workspaces in a couple of weeks. Not only that, I did it alongside my other day-to-day work, because I could let AI work in the background while I did the more urgent work I was assigned.
I was slightly bowled over by the success of this project, and it clearly proved to me the merits of using AI tools.
Given the lack of a clear replacement to CDKTF, I wondered if others in the community might find this tool useful, and for this reason we have decided to open source tfgen.
You can find it at github.com/ably/tfgen .
The entire core package is only around 600 lines of Go, and contains everything needed to generate complete Terraform configurations, as well as provider helpers. For example, we've provided a small set of resources for the Ably Terraform provider :
stack := tfgen.NewStack( "ably" )
block := &tfgen.TerraformBlock{RequiredVersion: ">= 1.9" }
block.AddRequiredProvider( "ably" , ably.ProviderVersion())
stack.SetTerraformBlock(block)
stack.AddProvider(ably.Provider())
app := ably.App( "main" , ably.AppOptions{TLSOnly: true })
stack.AddResource(app)
var buf bytes.Buffer
if err := stack.Write(&buf); err != nil {
panic (err)
}
fmt.Print(buf.String())
Since there are no generated types, references between resources are made using the Ref method, which renders as a Terraform expression (such as ${ably_app.main.id}) in the generated JSON:
stack.AddResource(ably.Namespace( "ai-transport" , ably.NamespaceOptions{
AppID: app.Ref( "id" ),
MutableMessages: true ,
}))
It also has the concept of writing multiple stacks built-in, since that's how we use it at Ably:
stack1 := tfgen.NewStack( "stack1" )
stack2 := tfgen.NewStack( "stack2" )
mystacks := tfgen.Stacks{stack1, stack2}
mystacks.Write( "my/stacks/directory" )
Frequently Asked Questions
What is CDKTF and why was it sunset? CDKTF (terraform-cdk) let engineers write infrastructure code in general-purpose languages like TypeScript, Python and Go, which compiled to Terraform-compatible JSON. HashiCorp announced in December 2025 that it would stop developing the tool, with no new features and bug fixes.
Is tfgen compatible with OpenTofu as well as Terraform? Yes. tfgen produces standard Terraform-compatible JSON, so it works with both Terraform and OpenTofu without any changes to how you write your configuration.
How is tfgen different from CDKTF? tfgen doesn't generate provider-specific types. Instead of fetching and maintaining strongly typed provider definitions, you set configuration values directly and validate the output using Terraform's own tooling. This removes the extra build step CDKTF required.
Can I use tfgen for providers other than Ably's? Yes. The core package is provider-agnostic. Ably has included a small set of helpers for the Ably Terraform provider as a working example, but tfgen's Resource and Config types work with any Terraform provider.
Do I need to write Go code for every Terraform resource type? No. tfgen doesn't require generated types per resource, so there's nothing provider-specific to install or keep in sync. You express any resource using the generic Resource type and its Config map.
If you find it useful and have any feature requests or questions, please get in touch or raise a GitHub issue .
Build dependable realtime in minutes.
Ably runs the same patterns we write about — in production, at scale. Free tier includes 6M messages a month, no card required.
New posts from the Ably team, monthly.
Provisioning infrastructure at Ably
Twitter / X LinkedIn Y Hacker News Copy link Continue reading
How to connect to Ably directly (and why you probably shouldn't) – Part 2
VPC peering vs Trans

[truncated]
