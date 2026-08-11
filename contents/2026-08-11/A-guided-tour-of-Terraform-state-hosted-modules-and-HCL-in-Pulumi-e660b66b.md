---
source: "https://www.pulumi.com/blog/terraform-to-pulumi-cloud-hands-on/"
hn_url: "https://news.ycombinator.com/item?id=49265772"
title: "A guided tour of Terraform state, hosted modules, and HCL in Pulumi"
article_title: "A guided tour of Terraform state, hosted modules, and HCL in Pulumi | Pulumi Blog"
author: "cnunciato"
captured_at: "2026-08-11T23:29:57Z"
capture_tool: "hn-digest"
hn_id: 49265772
score: 1
comments: 0
posted_at: "2026-08-11T23:14:59Z"
tags:
  - hacker-news
  - translated
---

# A guided tour of Terraform state, hosted modules, and HCL in Pulumi

- HN: [49265772](https://news.ycombinator.com/item?id=49265772)
- Source: [www.pulumi.com](https://www.pulumi.com/blog/terraform-to-pulumi-cloud-hands-on/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T23:14:59Z

## Translation

タイトル: Pulumi の Terraform 状態、ホストされたモジュール、HCL のガイド付きツアー
記事のタイトル: Pulumi の Terraform 状態、ホストされたモジュール、および HCL のガイド付きツアー |プルミのブログ
説明: Terraform と OpenTofu の新しいサポート (状態、リモート実行、モジュール相互運用、HCL のファーストクラス サポートなど) の実践的なチュートリアルです。

記事本文:
Pulumi の Terraform 状態、ホストされたモジュール、HCL のガイド付きツアー |プルミブログ メインコンテンツにスキップ
最新リリース: Terraform 状態、言語間モジュール、およびファーストクラス言語としての HCL を完全にサポート。リリースを参照
製品
プラットフォームの概要
プラットフォーム エンジニアリング チームがクラウド インフラストラクチャを構築、保護、拡張するために必要なものすべて
コードとしてのインフラストラクチャ
あらゆるクラウド、あらゆる言語の IaC — Node.js、Python、Go、.NET、Java、YAML
AIインフラエージェント
AI を活用したインフラストラクチャ エンジニアリング エージェント Neo をご紹介します
秘密と設定
環境、シークレット、構成管理
ディスカバリーとガバナンス
クラウド上の資産管理、コンプライアンス修復、AI の洞察
社内開発者プラットフォーム
クラウド インフラストラクチャを提供するための最速かつ最も安全な方法
始めましょう
ステップバイステップのガイドに従って、Pulumi をすばやく学習します
ドキュメント
完全なガイドと API リファレンス
レジストリ
170 以上のクラウド プロバイダーとパッケージを参照
テンプレート
共通のアーキテクチャをあらゆるクラウドに導入
チュートリアル
Pulumi のコンセプトを実際に体験してみよう
イベントやワークショップ
ライブセッションとワークショップ
コミュニティ
Slack で 10,000 人以上の開発者に参加しましょう
エンジニアはプルミを愛しています
エンジニアが当社を愛している理由を聞く
プルミガイド
クラウドプロバイダー向けの実践的なインフラストラクチャパターン
リリース
チームによる主要なプラットフォームのアップデート
エンタープライズソリューション
セキュリティ、コンプライアンス、チームのサポート
ケーススタディ
Snowflake、Mercedes-Benz などが Pulumi をどのように使用しているか
デモをリクエストする
Pulumi がチームにどのように役立つかをご覧ください
プロフェッショナルサービス
導入に関して専門家のサポートを受ける
営業担当者にお問い合わせください
ニーズについて当社のチームにご相談ください
私たちについて
私たちの目的と価値観
受賞歴
報道機関やアナリストからの評価
プラットフォームの概要
プラットフォーム エンジニアリング チームがクラウド インフラストラクチャを構築、保護、拡張するために必要なものすべて

e
コードとしてのインフラストラクチャ
あらゆるクラウド、あらゆる言語の IaC — Node.js、Python、Go、.NET、Java、YAML
AIインフラエージェント
AI を活用したインフラストラクチャ エンジニアリング エージェント Neo をご紹介します
秘密と設定
環境、シークレット、構成管理
ディスカバリーとガバナンス
クラウド上の資産管理、コンプライアンス修復、AI の洞察
社内開発者プラットフォーム
クラウド インフラストラクチャを提供するための最速かつ最も安全な方法
始めましょう
ステップバイステップのガイドに従って、Pulumi をすばやく学習します
ドキュメント
完全なガイドと API リファレンス
レジストリ
170 以上のクラウド プロバイダーとパッケージを参照
テンプレート
共通のアーキテクチャをあらゆるクラウドに導入
チュートリアル
Pulumi のコンセプトを実際に体験してみよう
イベントやワークショップ
ライブセッションとワークショップ
コミュニティ
Slack で 10,000 人以上の開発者に参加しましょう
エンジニアはプルミを愛しています
エンジニアが当社を愛している理由を聞く
プルミガイド
クラウドプロバイダー向けの実践的なインフラストラクチャパターン
リリース
チームによる主要なプラットフォームのアップデート
エンタープライズソリューション
セキュリティ、コンプライアンス、チームのサポート
ケーススタディ
Snowflake、Mercedes-Benz などが Pulumi をどのように使用しているか
デモをリクエストする
Pulumi がチームにどのように役立つかをご覧ください
プロフェッショナルサービス
導入に関して専門家のサポートを受ける
営業担当者にお問い合わせください
ニーズについて当社のチームにご相談ください
私たちについて
私たちの目的と価値観
受賞歴
報道機関やアナリストからの評価
チュートリアル Pulumi の Terraform 状態、ホストされたモジュール、HCL のガイド付きツアー
Terraform プロジェクトから始める
状態を Pulumi Cloud に移行する
Pulumi プログラムからモジュールを使用する
本日の大きなリリースには、Terraform および OpenTofu エコシステムとのシームレスな相互運用性を実現するために設計されたまったく新しい機能セットが含まれており、その機能のすべてを理解するのは難しいほどたくさんあります。しかし、それは一般的に 3 つの主要なカテゴリに分類されます

リース:
人間の承認によるリモート実行を含む、Terraform 状態バックエンドとしての Pulumi Cloud のサポート
Pulumi Cloud の Terraform モジュール レジストリにより、言語の境界を超えてモジュールを公開、文書化、共有できます。
Pulumi エンジンのオーサリング言語としての HCL のファーストクラスのサポート
このリリースを全体的に理解しやすくするために、すべてを完全にカバーしているわけではありませんが、主要な部分はカバーしており、すべてがどのようにまとめられているかを理解できるように、簡単なエンドツーエンドのウォークスルーをまとめました。まず、AWS にデプロイする単純な Terraform プロジェクトから始めて、次にそれを Pulumi Cloud に取り込み、これらの各新機能を試していきます。少し時間がかかりますが、必要なのは無料の Pulumi アカウントと、S3 バケットを AWS にデプロイする機能だけです。
カバーしなければならないことがたくさんあるので、早速始めましょう。
Terraform プロジェクトから始める
私たちのツアーは、後で公開するローカルで定義されたモジュールを使用して単一の Amazon S3 バケットをプロビジョニングする小さな Terraform プロジェクトから始まります。このプロジェクトは GitHub でテンプレートとして入手でき、GitHub CLI を使用するのが最も簡単な方法です。
$ gh リポジトリ作成 my-tf-project \
--template cnunciato/simple-tf-template \
--パブリック\
--clone && cd my-tf-project
まず、ローカルの Terraform バックエンドを使用します。 AWS 認証情報を設定し (できれば環境変数を使用して)、Terraform または OpenTofu を使用してプロジェクトをデプロイします。 (このチュートリアルでは terraform CLI を使用しますが、お好みに応じて tofu に置き換えることもできます。)
$ terraform init && terraform apply
...
応募完了！リソース: 2 追加、0 変更、0 破壊。
出力:
bucket_arn = "arn:aws:s3:::my-tf-project-bucket-14d19ece"
バケット名 = "my-tf-プロジェクト-バケット-14d19ece"
今

このプロジェクトを Pulumi Cloud に移動する方法を見てみましょう。
状態を Pulumi Cloud に移行する
まず、Pulumi Cloud アカウントをまだお持ちでない場合は作成し (個人の場合は無料です)、Pulumi コンソールにサインインします。次に、バックエンド ブロックを main.tf に追加し、 <your-org> を自分の Pulumi Cloud アカウントまたは組織名に置き換えます。
テラフォーム {
# ...
バックエンド「リモート」{
ホスト名 = "tf.pulumi.com"
組織 = "<あなたの組織>"
ワークスペース {
名前 = "my-tf-project_dev"
}
}
}
ワークスペース名は、使用するプロジェクトの名前 (ここでは my-tf-project ) とスタック ( dev ) を表すアンダースコアで区切られた文字列です。 1 つの Pulumi プロジェクトには、好きなだけスタックを含めることができます。
次に、Terraform CLI を使用して Pulumi Cloud にサインインします。
$ terraform ログイン tf.pulumi.com
プロンプトが表示されたら「はい」を選択すると、Pulumi Cloud に移動して個人用アクセス トークンを作成します。これをプロンプトに貼り付けて認証できます。
成功! Terraform Enterprise (tf.pulumi.com) にログインしました
バックエンド ブロックを配置し、terraform CLI で Pulumi Cloud にサインインすると、移行を完了する準備が整います。
$ terraform init -移行状態
Terraform はバックエンドの変更を検出し、既存の状態のコピーを提案する必要があります。
既存の状態を新しいバックエンドにコピーしますか?
以前の「ローカル」バックエンドの移行中に既存の状態が見つかりました
新しく構成された「リモート」バックエンドに接続します。 [...] 「yes」と入力してコピーし、
"no" は空の状態から開始します。
値を入力してください: はい
[はい] を選択すると完了です。
バックエンド「リモート」が正常に構成されました。 Terraform は自動的に
バックエンド構成が変更されない限り、このバックエンドを使用してください。
ここでは、デプロイされたインフラストラクチャについては何も変更されていないことに注意してください。私たちが行ったのは、ローカル状態を Pulumi Cloud に移行することだけであり、そのプロセスは次のとおりです。

S3、Azure、Google Cloud、HCP Terraform または Terraform Enterprise から移行する場合も同様です。詳細については、「Pulumi Cloud への Terraform 状態の保存」を参照してください。
ここで、Pulumi Cloud コンソールに移動し、 [スタック] を選択すると、リストに新しいスタックとその最初の更新が表示されます。
デフォルトでは実行はリモートで行われます
もう 1 つ注意すべき点は、Pulumi Cloud によってサポートされる Terraform スタックは、HCP Terraform や Terraform Enterprise と同様に、デフォルトでリモートで実行されることです。計画と適用は Pulumi Cloud でホストされるランナー上で実行され、出力は端末とブラウザに同時にストリーミングされます。 main.tf に小さな変更を加えて (タグを追加するなど)、 terraform apply を実行してみてください。
タグ = {
環境 = "開発"
+ 所有者 = 「私」
}
$ terraform 適用
...
リモート バックエンドで apply を実行します。出力はここにストリーミングされます。
この実行をブラウザで表示するには、次の場所にアクセスしてください。
https://tf.pulumi.com/app/cnunciato/my-tf-project_dev/runs/run-8a635542-334f-4367-b188-8b4b442b2550
...
もう一度「はい」で確認すると、適用が失敗することがわかります。ランナーがまだ AWS 認証情報を持っていないことを考えると、これは当然のことです。そこでPulumi ESCの出番です。
ESC を使用してクラウド認証情報を構成する
新しい Terraform スタックを作成すると、Pulumi Cloud は同じ名前でリンクされた ESC 環境を自動的に作成します。この環境を使用して、クラウド資格情報やその他の暗号化されたシークレットを含むあらゆる種類の設定を構成し、それらの設定を実行時にスタックで利用できるようにすることができます。
スタックの [概要] タブで、リンクされた my-tf-project/dev 環境をクリックして環境エディターを開きます。そこで AWS 認証情報を直接設定するか、[新規追加] → [ログインプロバイダー/構成] を選択して AWS の手順に従って新しい接続を確立できます。
私は以来、

すべてのスタックで AWS 認証情報を共有するための環境 (default/personal という名前を付けました) がすでに構成されているので、その環境をこの環境にインポートするだけで使用できます。
輸入品 :
- デフォルト/個人用
これを自分で試してから、再度適用を実行し、成功することを確認します。
$ terraform apply -auto-approve
...
OpenTofu は次のアクションを実行します。
# module.s3-bucket.aws_s3_bucket.this はインプレースで更新されます
~ リソース "aws_s3_bucket" "this" {
id = "my-tf-プロジェクト-バケット-b64b8e37"
~ タグ = {
「環境」=「開発」
+「所有者」=「私」
}
~ タグすべて = {
+「所有者」=「私」
}
}
計画: 0 で追加、1 で変更、0 で破壊。
...
応募完了！リソース: 0 追加、1 変更、0 破壊。
出力:
bucket_arn = "arn:aws:s3:::my-tf-project-bucket-b64b8e37"
バケット名 = "my-tf-プロジェクト-バケット-b64b8e37"
リモート実行は Pulumi Deployments によって強化されています。つまり、Terraform スタックは VCS イベント (プル リクエストや GitHub、GitLab などへのプッシュ)、および人間による手動の承認によってトリガーすることもできます。配線も非常に簡単です。
Pulumi コンソールの「管理」→「バージョン管理」で、VCS プロバイダーを構成し、my-tf-project リポジトリを追加します。
スタックの [設定] タブで、 [デプロイ] を選択し、リポジトリとビルド元のベース ブランチ (例: main ) を構成し、保存します。
これを行うと、[デプロイメント] タブに新しいプランが表示され、それが完了すると、実行を承認または拒否するように求められます。
[確認] をクリックすると、起動して実行できるようになります。
それで終わりです！ Terraform スタックは Pulumi Cloud の第一級市民となり、アクセス コントロール 、 Neo コード レビュー 、および Pulumi ポリシーをすべて利用できるようになりました。監査ポリシーはどの Terraform スタックでも実行可能で、予防ポリシーを使用して非準拠をブロックできます。

変化が起こる前に変化します。
チームが Terraform をしばらく使用している場合は、Terraform モジュールも作成している可能性が高いため、Pulumi Cloud に移行する際にそれらを保管する場所が必要になります。 Pulumi Cloud には、Terraform 状態バックエンドとしての役割に加えて、Terraform モジュールをホストし、組織全体で利用できるようにするプライベート レジストリが含まれるようになりました。
モジュールの公開には Enterprise プランまたは Business Critical プランが必要なため、この部分ではいずれかのプランを備えた組織が必要になります。ただし、まだプランをお持ちでない場合は、無料トライアルを利用して簡単に作成できます。 Pulumi コンソールで、組織メニューを開き、 [組織の作成] を選択し、名前を付ければ準備完了です。
Pulumi Cloud の Terraform レジストリ API は、HCP Terraform のレジストリ API と有線互換性があります。つまり、公開にすでに使用しているツール (たとえば、 go-tfe ライブラリや hach​​icorp/tfe Terraform プロバイダー) は、 tf.pulumi.com に指定するだけで済みます。これを簡単にするために、私たちのプロジェクトには go-tfe を使用してそれを行う小さな Go プログラムが含まれています。ターミナルに Pulumi アクセス トークンを設定し (新しいトークンを取得する必要がある場合があります)、組織名を渡してプロジェクト ルートから実行します。
$export PULUMI_ACCESS_TOKEN = <アクセストークン>
# https://go.dev/doc/install で Go をインストールします。

[切り捨てられた]

## Original Extract

A hands-on walkthrough of our new support for Terraform and OpenTofu, including state, remote execution, module interop, and first-class support for HCL.

A guided tour of Terraform state, hosted modules, and HCL in Pulumi | Pulumi Blog Skip to main content
Latest release: Full support for Terraform state, cross-language modules, and HCL as a first-class language. See the release
Product
Platform overview
Everything platform engineering teams need to build, secure, and scale cloud infrastructure
Infrastructure as code
IaC for any cloud, in any language — Node.js, Python, Go, .NET, Java, and YAML
AI infrastructure agent
Meet Neo, our AI-powered infrastructure engineering agent
Secrets & configuration
Environments, secrets, and configuration management
Discovery & governance
Asset management, compliance remediation, and AI insights over the cloud
Internal developer platform
The fastest, most secure way to deliver cloud infrastructure
Get started
Follow a step-by-step guide to quickly learn Pulumi
Documentation
Complete guides and API references
Registry
Browse 170+ cloud providers and packages
Templates
Deploy common architectures on any cloud
Tutorials
Get hands-on with Pulumi concepts
Events and workshops
Live sessions and workshops
Community
Join 10k+ developers on Slack
Engineers love Pulumi
Hear from engineers why they love us
Pulumi guides
Practical infrastructure patterns for cloud providers
Releases
Major platform updates from the team
Enterprise solutions
Security, compliance, and support for teams
Case studies
How Snowflake, Mercedes-Benz, and others use Pulumi
Request a demo
See how Pulumi can help your team
Professional services
Get expert help with your implementation
Contact sales
Talk to our team about your needs
About us
Our purpose and values
Awards
Recognition from press and analysts
Platform overview
Everything platform engineering teams need to build, secure, and scale cloud infrastructure
Infrastructure as code
IaC for any cloud, in any language — Node.js, Python, Go, .NET, Java, and YAML
AI infrastructure agent
Meet Neo, our AI-powered infrastructure engineering agent
Secrets & configuration
Environments, secrets, and configuration management
Discovery & governance
Asset management, compliance remediation, and AI insights over the cloud
Internal developer platform
The fastest, most secure way to deliver cloud infrastructure
Get started
Follow a step-by-step guide to quickly learn Pulumi
Documentation
Complete guides and API references
Registry
Browse 170+ cloud providers and packages
Templates
Deploy common architectures on any cloud
Tutorials
Get hands-on with Pulumi concepts
Events and workshops
Live sessions and workshops
Community
Join 10k+ developers on Slack
Engineers love Pulumi
Hear from engineers why they love us
Pulumi guides
Practical infrastructure patterns for cloud providers
Releases
Major platform updates from the team
Enterprise solutions
Security, compliance, and support for teams
Case studies
How Snowflake, Mercedes-Benz, and others use Pulumi
Request a demo
See how Pulumi can help your team
Professional services
Get expert help with your implementation
Contact sales
Talk to our team about your needs
About us
Our purpose and values
Awards
Recognition from press and analysts
Tutorials A guided tour of Terraform state, hosted modules, and HCL in Pulumi
Start with a Terraform project
Migrate the state to Pulumi Cloud
Consume the module from a Pulumi program
Today’s big release contains a whole new set of features designed for seamless interoperability with the Terraform and OpenTofu ecosystems, and there’s a lot there — so much that it can be tough to get your head around all of it. But it generally falls into three major categories:
Support for Pulumi Cloud as a Terraform state backend , including remote execution with human approvals
A Terraform module registry in Pulumi Cloud that lets you publish, document, and share your modules even across language boundaries
First-class support for HCL as an authoring language in the Pulumi engine
To make this release a little easier to appreciate holistically, I’ve put together a quick end-to-end walkthrough that doesn’t quite cover everything , but does cover the big stuff, and should give you a sense of how it all comes together. We’ll start with a simple Terraform project that you’ll deploy to AWS, and then one step at a time, bring it into Pulumi Cloud and kick the tires on each of these new features as we go. It’ll take a bit, but all you’ll need are a free Pulumi account and the ability to deploy an S3 bucket to AWS.
We’ve got a bunch to cover, so let’s jump right in.
Start with a Terraform project
Our tour begins with a tiny Terraform project that provisions a single Amazon S3 bucket using a locally defined module that we’ll publish later. The project is available on GitHub as a template, and the easiest way to use it is with the GitHub CLI:
$ gh repo create my-tf-project \
--template cnunciato/simple-tf-template \
--public \
--clone && cd my-tf-project
We’ll use the local Terraform backend to start. Set your AWS credentials (preferably with environment variables ), then deploy the project with Terraform or OpenTofu. (This walkthrough uses the terraform CLI, but you can swap in tofu if that’s your preference.)
$ terraform init && terraform apply
...
Apply complete! Resources: 2 added, 0 changed, 0 destroyed.
Outputs:
bucket_arn = "arn:aws:s3:::my-tf-project-bucket-14d19ece"
bucket_name = "my-tf-project-bucket-14d19ece"
Now let’s see how to move this project into Pulumi Cloud.
Migrate the state to Pulumi Cloud
First, create a Pulumi Cloud account if you don’t already have one (it’s free for individuals) and sign in to the Pulumi console. Then, add a backend block to main.tf , swapping <your-org> for your own Pulumi Cloud account or organization name:
terraform {
# ...
backend "remote" {
hostname = "tf.pulumi.com"
organization = "<your-org>"
workspaces {
name = "my-tf-project_dev"
}
}
}
The workspace name is an underscore-delimited string that expresses the name of the project you’d like to use (here, my-tf-project ) and the stack ( dev ). A single Pulumi project can have as many stacks as you like.
Next, sign in to Pulumi Cloud with the Terraform CLI:
$ terraform login tf.pulumi.com
Choose yes when prompted, and you’ll be taken to Pulumi Cloud to create a personal access token , which you can paste into the prompt to authenticate:
Success! Logged in to Terraform Enterprise (tf.pulumi.com)
With the backend block in place and your terraform CLI signed in to Pulumi Cloud, you’re ready to complete the migration:
$ terraform init -migrate-state
Terraform should detect the backend change and offer to copy your existing state:
Do you want to copy existing state to the new backend?
Pre-existing state was found while migrating the previous "local" backend
to the newly configured "remote" backend. [...] Enter "yes" to copy and
"no" to start with an empty state.
Enter a value: yes
Choose yes , and you’re done:
Successfully configured the backend "remote"! Terraform will automatically
use this backend unless the backend configuration changes.
Note that nothing about your deployed infrastructure has changed here; all we did was migrate your local state to Pulumi Cloud, and the process is identical whether you’re moving from S3, Azure, Google Cloud, or HCP Terraform or Terraform Enterprise. See Store Terraform state in Pulumi Cloud for details.
Now hop over to the Pulumi Cloud console, choose Stacks , and you’ll see your new stack in the list, along with its first update:
Runs happen remotely by default
Another thing to note is that Terraform stacks backed by Pulumi Cloud run remotely by default, just as they do in HCP Terraform and Terraform Enterprise. Plans and applies are executed on a Pulumi Cloud-hosted runner, and output is streamed simultaneously to your terminal and to the browser. Try making a small change to main.tf — adding a tag, say — and run terraform apply :
tags = {
Environment = "dev"
+ Owner = "me"
}
$ terraform apply
...
Running apply in the remote backend. Output will stream here.
To view this run in a browser, visit:
https://tf.pulumi.com/app/cnunciato/my-tf-project_dev/runs/run-8a635542-334f-4367-b188-8b4b442b2550
...
Confirm with another yes , and you’ll notice the apply fails — which makes sense, considering the runner doesn’t have your AWS credentials yet. That’s where Pulumi ESC comes in.
Configure cloud credentials with ESC
When you create a new Terraform stack, Pulumi Cloud automatically creates a linked ESC environment for you with the same name. You can use this environment to configure settings of all kinds, including cloud credentials and other encrypted secrets, and make those settings available to the stack at runtime.
In the stack’s Overview tab, click the linked my-tf-project/dev environment to open the environment editor, where you can either set your AWS credentials directly or wire up a new connection by choosing Add new → Login provider/configuration and following the steps for AWS.
Since I’ve already configured an environment for sharing my AWS credentials across all of my stacks (which I’ve named default/personal ), I can just import that environment into this one to use it:
imports :
- default/personal
Try this yourself, then run the apply again, and see that it succeeds:
$ terraform apply -auto-approve
...
OpenTofu will perform the following actions:
# module.s3-bucket.aws_s3_bucket.this will be updated in-place
~ resource "aws_s3_bucket" "this" {
id = "my-tf-project-bucket-b64b8e37"
~ tags = {
"Environment" = "dev"
+ "Owner" = "me"
}
~ tags_all = {
+ "Owner" = "me"
}
}
Plan: 0 to add, 1 to change, 0 to destroy.
...
Apply complete! Resources: 0 added, 1 changed, 0 destroyed.
Outputs:
bucket_arn = "arn:aws:s3:::my-tf-project-bucket-b64b8e37"
bucket_name = "my-tf-project-bucket-b64b8e37"
Remote execution is powered by Pulumi Deployments , which means your Terraform stacks can also be triggered by VCS events — pull requests and pushes to GitHub, GitLab, and others — as well as manual human approvals. Wiring that up is pretty simple as well:
Under Management → Version Control in the Pulumi console, configure your VCS provider and add the my-tf-project repository.
In your stack’s Settings tab, choose Deploy , configure the repository and the base branch to build from (e.g., main ), and save.
When you do that, you’ll see a new plan in the Deployments tab, and once that finishes, you’ll be prompted to approve or decline the run:
Click Confirm , and you’re off and running.
And that’s it! Your Terraform stacks are now first-class citizens in Pulumi Cloud, with access control , Neo code reviews , and Pulumi Policies all available to them. Audit policies are runnable on any Terraform stack, and preventative policies can be used to block non-compliant changes before they happen.
If your team’s been using Terraform for a while, chances are you’ve also written some Terraform modules, so you’ll need somewhere to keep them as you move to Pulumi Cloud. In addition to its role as a Terraform state backend, Pulumi Cloud now includes a private registry that can host your Terraform modules and make them available across your organization.
Module publishing requires an Enterprise or Business Critical plan, so you’ll need an organization with one of those for this part — but you can easily create one with a free trial if you don’t have one already. In the Pulumi console, open the organization menu, choose Create organization , give it a name, and you’re good to go.
Pulumi Cloud’s Terraform registry API is wire-compatible with HCP Terraform’s, which means the tools you already use for publishing — e.g., the go-tfe library or the hashicorp/tfe Terraform provider — need only be pointed at tf.pulumi.com . To make this easy, our project includes a small Go program that uses go-tfe to do just that. Set a Pulumi access token in your terminal (you may need to obtain a new one ), then run it from the project root, passing your organization name:
$ export PULUMI_ACCESS_TOKEN = <your-access-token>
# Install Go at https://go.dev/doc/install if

[truncated]
