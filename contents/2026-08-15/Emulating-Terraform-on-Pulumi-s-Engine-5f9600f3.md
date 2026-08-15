---
source: "https://www.pulumi.com/blog/terraforms-data-model-on-pulumis-engine/"
hn_url: "https://news.ycombinator.com/item?id=49314291"
title: "Emulating Terraform on Pulumi's Engine"
article_title: "Emulating Terraform on Pulumi's Engine | Pulumi Blog"
author: "cnunciato"
captured_at: "2026-08-15T21:12:03Z"
capture_tool: "hn-digest"
hn_id: 49314291
score: 1
comments: 0
posted_at: "2026-08-15T21:00:21Z"
tags:
  - hacker-news
  - translated
---

# Emulating Terraform on Pulumi's Engine

- HN: [49314291](https://news.ycombinator.com/item?id=49314291)
- Source: [www.pulumi.com](https://www.pulumi.com/blog/terraforms-data-model-on-pulumis-engine/)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T21:00:21Z

## Translation

タイトル: Pulumi のエンジンで Terraform をエミュレートする
記事のタイトル: Pulumi のエンジンで Terraform をエミュレートする |プルミのブログ
説明: Terraform をマッピングする方法

記事本文:
Pulumi のエンジンで Terraform をエミュレートする |プルミブログ メインコンテンツにスキップ
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
コードとしてのインフラストラクチャ
IaC

あらゆるクラウド、あらゆる言語 — Node.js、Python、Go、.NET、Java、YAML
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
Pulumi のエンジンで Terraform をエミュレートするエンジニアリング
Pulumi の HCL サポートの主な約束は、既存の Terraform 構成とモジュールを持ち込んで、pulumi がそれらを実行できることです。 OpenTofu では動作するが Pulumi では動作しない場合は、それを修正したいと考えています。その目標を考慮すると、HCL インタープリターは HCL を入力として受け取り、tofu が同じ入力を解釈する方法と意味的に一致する命令を Pulumi エンジンに発行する必要があります。これは、Pulumi と O という事実によってさらに困難になります。

penTofu は根本的に異なるエンジン セマンティクスとプロバイダー エコシステムを備えています。このブログ投稿では、上位の Terraform モジュールの 96% 1 を Pulumi 上で動作させるために、そのマッピングをどのように実装したかを探ります。 Pulumi の HCL インタープリターが Terraform のリソース セマンティクス、プロバイダー、モジュールをどのように処理するかを簡単に説明します。また、Pulumi の HCL サポートにより、Terraform や OpenTofu では許可されないことが可能になる場合についても説明します。
Pulumi と Terraform には両方ともプロバイダーがありますが、同じプロバイダーはありません。 Terraform にないプロバイダーもありますが、Pulumi は、Pulumi の紛らわしい名前の terraform-provider プロバイダーを使用して、いつでも Terraform プロバイダーを解決できます。 2 これは、 pulumi package add terraform-provider ... を使用して、別の Pulumi プログラムで任意の Terraform Provider を使用できるのと同じプロバイダーです。 Terraform プロバイダー プロバイダーはリレーとして機能します。Pulumi のプロトコルを Pulumi エンジンに伝え、Terraform のプロバイダー プロトコルを立ち上げた Terraform プロバイダーに伝えます。 Pulumi HCL はすべての Pulumi プロバイダーと連携する必要があり、terraform-provider により Pulumi HCL が Pulumi プロトコル経由で Terraform プロバイダーと通信できるため、Pulumi HCL は実際には Pulumi プロトコルのみを直接通信します。
フローチャート LR
サブグラフ n2Entry[" "]
n2["テラフォームプロバイダー"]
終わり
サブグラフプロバイダボックス["Pulumi プロバイダ"]
TD方向
n2エントリー
n3["Terraform プロバイダー"]
終わり
n0["Pulumi HCL"] <--> n1["Pulumi エンジン"]
n1 <--> n2エントリ
n2 <--> n3
n2@{形状: 長方形}
n3@{形状: 長方形}
n0@{形状: 長方形}
n1@{形状: 長方形}
スタイル n2Entry 塗りつぶし:透明、ストローク:透明
github.com/pulumi/pulumi-hcl の内部はすべて Pulumi の言語プロトコルで実装されています。
terraform プロバイダーは Terraform のバージョン範囲をネイティブに理解し、デフォルトで Ope を使用するためです。

nTofu レジストリを使用すると、通常のプロバイダー リクエストを terraform-provider に直接変換できます。このインターフェイスは元々コマンド ライン用に設計されているため、引数を型なしの文字列リストとして受け取ります。いくつかの簡単な例を見てみましょう。
リソース "aws_s3_bucket" "example" {
バケット = "my-bucket-123"
}
Pulumi の HCL インタープリターは、terraform.required_providers ブロックが存在しないことを認識するため、最初の _ でリソース トークンを切り取り、デフォルトのレジストリと名前空間を使用します。 Pulumi エンジンに送信されるリクエストは次のようになります。
プルミルPC 。パッケージ仕様 {
ソース: "terraform-provider" 、
パラメータ: [] string { "registry.opentofu.org/bashicorp/aws" },
}
バージョンが指定されていないため、バージョンの決定は terraform-provider に任せます。最新バージョンが使用されます。
Terraform と同じように、required_providers ブロックでこれをオーバーライドできます。
テラフォーム {
必須プロバイダー {
例 = {
ソース = "my.custom.registry/me/example"
バージョン = "~> 5.0"
}
}
}
リソース "example_resource" "another_example" {
}
同じ機械翻訳を実行します。ソースは完全に指定されており、バージョンがあるため、それを Pulumi エンジンに渡し、Pulumi エンジンはそれを terraform-provider に渡します。
プルミルPC 。パッケージ仕様 {
ソース: "terraform-provider" 、
パラメータ: [] string { "my.custom.registry/me/example" , "~> 5.0" },
}
Pulumi HCL は、terraform プロバイダー トランスレーターを通じて、ほぼすべてのプロバイダー呼び出しを自動的に翻訳します。代わりに、ソース pulumi/* を持つプロバイダーは直接ルーティングされます。 Pulumi HCL でネイティブ Pulumi プロバイダーを使用する方法は次のとおりです。
テラフォーム {
必須プロバイダー {
kubernetes = {
ソース = "pulumi/kubernetes"
バージョン = "4.33.0"
}
}
}
リソース "kubernetes_yaml_config_file" "app" {
ファイル = "app.yaml"
}
弊社の HCL 解釈

ter はこれを Pulumi Kubernetes パッケージに直接ルーティングします。
プルミルPC 。パッケージ仕様 {
出典:「kubernetes」、
バージョン：「4.33.0」、
}
バージョンが「パラメータ」から「バージョン」フィールドに移動したことを確認します。これは、source = "pulumi/*" で記述されたプロバイダーがプラグインについて直接話しているためです。
Pulumi プログラムと Terraform config は両方とも、リソース グラフを表現するために存在します。それが彼らの主な目的です。すべてを説明する時間もピクセル数もないので、ここでは 3 つのサブトピックに限定します。
プロバイダーの背後にあるスキーマを見たことがあれば、名前がキャメルケースであることに気づくでしょう。
"aws:s3/バケット:バケット" : {
"プロパティ" : {
"accelerationStatus" : { "type" : "string" , "description" : "..." , "deprecationMessage" : "..." },
"bucketDomainName" : { "type" : "string" , "description" : "..." },
"bucketNamespace" : { "type" : "string" , "description" : "..." },
...
私たちの codegen は実際にこの事実に依存しています。これは、terraform-provider によって生成されたスキーマでもキャメルケースのプロパティ名を使用することを意味します。賢い人は問題に気づくでしょう。Terraform は慣例により、snake_case を使用します。 Pulumi HCL がプロパティ値を含むリソース リクエストを送信する場合、エンジン用にスネークケース テキストからキャメルケースに変換し、エンジンから返された値に再度変換する必要があります。
シーケンス図
HCL プログラムとしての参加者プログラム<br/>(snake_case)
HCL 通訳者としての参加者 Interp
Pulumi エンジンとしての参加者エンジン<br/>(キャメルケース)
Interp->>Engine: GetSchema("example_resource")
エンジン -->>Interp: スキーマ (プロパティ名 + 型)
プログラム->>Interp: some_value = "true"
Interp->>Engine: RegisterResource { someValue: true }
Engine-->>Interp: 出力 { someValue: true, computedValue: "..." }
Interp-->>プログラム: example_resource.example.computed_value

これを正しく行うには、実際にはリソースのスキーマを認識する必要があります。生の値では十分ではありません。その理由を理解するために例を見てみましょう。
リソース "example_resource" "example" {
some_value = "true"
ブロック{
inner_value = false
}
別の属性 = {
「トラップワン」= 1
「トラップツー」= 2
}
}
ここでは、上記の「例」の入力に対応する 2 つの異なる値 BLOB (JSON として表されます) を示します。 "example_resource" のタイプに応じて、どちらも有効です。
{
"someValue" : "true" ,
"ブロック" : {
"innerValue" : false
}、
"別の属性" : {
"trap_one" : 1 、
"trap_two" : 2
}
}
{
"someValue" : true 、
"ブロック" : [{
"innerValue" : "false"
}]、
"別の属性" : {
"トラップワン" : "1" 、
「トラップツー」：2
}
}
2 つの潜在的な値の blob の違いを詳しく見てみましょう。
HCL はスカラー値をプロバイダーが予期する型に変換するため、 someValue が文字列として型指定されている場合は、 "true" を維持します。ブール値として入力された場合は、変換してプロバイダーに変換後の値 true を渡します。
ブロックがオブジェクトとして型指定されている場合、Pulumi プロバイダーはネットワーク上にオブジェクトが返されることを期待します。 block がオブジェクトのリストとして型指定されている場合、Pulumi プロバイダーは、それが 1 つのリストであっても、オブジェクトのリストを返す必要があります。
innerValue は、 someValue と同じ型変換を行うことができます。型変換はトップレベルで止まりません。
anotherAttribute がマップ タイプの場合、そのキーはユーザー指定の値であるため、そのままにしておく必要があります。オブジェクト型の場合、そのキーをキャメルケースに戻す必要があります。
したがって、Pulumi の HCL インタプリタは型を認識します。変換対象の各プロバイダーのスキーマをエンジンに問い合わせて、正しい変換を実行し、変換の値を追跡するときにターゲットの型を追跡します。これはリソースに限定されません。同じ翻訳を行います

データ ソース、モジュール、プロバイダー ブロックのオプション。
Terraform プロバイダーからブリッジされる Pulumi プロバイダーの場合は、追加の手順を適用します。ブリッジプロバイダーは、 pulumi Convert --from terraform によって使用される明示的な命名テーブルを公開します。 Pulumi HCL は、このマッピングについてもエンジンにクエリを実行し、それを使用して、プロパティが Pulumi で名前変更されているか非表示になっている場合でも、リソース名とプロパティ名が基になる Terraform プロバイダーと正確に一致していることを確認します。
Terraform プロビジョナーを使用すると、Terraform 構成でリソースのライフサイクルのさまざまな時点で実行するコマンドを指定できます。条件でも同じことができますが、代わりにブール式が評価されます。たとえば:
リソース "example_resource" "example" {
入力 = var 。入力
ライフサイクル {
前提条件 {
条件 = var 。入力 > 3
error_message = "var.input は 3 を超える必要があります"
}
}
プロビジョナー "local-exec" {
command = "echo 出力例は ${self.output} >> file.txt です"
}
}
Pulumi には、リソースのライフサイクルにフックするメカニズムが 1 つだけあります: フック 。 Pulumi のリソース フックは、プロビジョナーと条件の両方を表現するのに十分強力です。これを表現する最も簡単な方法は、他の言語と比較することです。 TypeScriptを使用します。上記の構成ブロックの例は、この TypeScript プログラムと同等です。

[切り捨てられた]

## Original Extract

How we map Terraform

Emulating Terraform on Pulumi's Engine | Pulumi Blog Skip to main content
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
Engineering Emulating Terraform on Pulumi's Engine
The core promise of Pulumi’s HCL support is that you can bring your existing Terraform configuration and modules, and pulumi will run them. If it works in OpenTofu and doesn’t work in Pulumi, we would like to fix that. Given that goal, our HCL interpreter needs to take HCL as input and emit instructions to the Pulumi engine that semantically match how tofu would interpret the same input. This is made harder by the fact that Pulumi and OpenTofu have fundamentally different engine semantics and provider ecosystems. This blog post will explore how we have implemented that mapping well enough to get 96% 1 of our top Terraform modules working on Pulumi. We’ll briefly walk through how Pulumi’s HCL interpreter handles Terraform’s resource semantics, providers, and modules. It will also call out where Pulumi’s HCL support lets you do things that Terraform and OpenTofu will not allow.
Both Pulumi and Terraform have providers, but they don’t have the same providers. While there are providers that Terraform does not have , Pulumi can always resolve a Terraform provider using Pulumi’s confusingly named terraform-provider provider. 2 This is the same provider that lets you consume Any Terraform Provider in another Pulumi program with pulumi package add terraform-provider ... . The terraform-provider provider acts as a relay: it speaks Pulumi’s protocol to the Pulumi engine, and speaks Terraform’s provider protocol to the Terraform provider it stands up. Because Pulumi HCL needs to work with all Pulumi providers and because terraform-provider lets Pulumi HCL speak to Terraform providers via the Pulumi protocol, Pulumi HCL actually only speaks Pulumi protocols directly:
flowchart LR
subgraph n2Entry[" "]
n2["terraform-provider"]
end
subgraph providerBox["Pulumi Provider"]
direction TD
n2Entry
n3["Terraform Provider"]
end
n0["Pulumi HCL"] <--> n1["Pulumi Engine"]
n1 <--> n2Entry
n2 <--> n3
n2@{ shape: rect}
n3@{ shape: rect}
n0@{ shape: rect}
n1@{ shape: rect}
style n2Entry fill:transparent,stroke:transparent
All internals of github.com/pulumi/pulumi-hcl are implemented in Pulumi’s language protocol.
Because the terraform-provider natively understands Terraform version ranges and defaults to the OpenTofu registry, we can directly translate normal provider requests to the terraform-provider . It takes its arguments as an untyped list of strings, since the interface was originally intended for the command line. Let’s walk through some simple examples:
resource "aws_s3_bucket" "example" {
bucket = "my-bucket-123"
}
Pulumi’s HCL interpreter sees that there is no terraform.required_providers block, so it cuts the resource token at the first _ and uses the default registry and namespace. This is what the request that goes to the Pulumi engine looks like:
pulumirpc . PackageSpec {
Source : "terraform-provider" ,
Parameters : [] string { "registry.opentofu.org/hashicorp/aws" },
}
Because no version was specified, we leave it to terraform-provider to determine the version. It will use the latest version .
Just like Terraform, you can override this with a required_providers block:
terraform {
required_providers {
example = {
source = "my.custom.registry/me/example"
version = "~> 5.0"
}
}
}
resource "example_resource" "another_example" {
}
We perform the same mechanical translation. The source is fully specified, and there is a version, so we pass it along to the Pulumi engine , which passes it along to terraform-provider :
pulumirpc . PackageSpec {
Source : "terraform-provider" ,
Parameters : [] string { "my.custom.registry/me/example" , "~> 5.0" },
}
Pulumi HCL will automatically translate almost all provider calls through the terraform-provider translator. Providers with the source pulumi/* will instead be routed directly. This is how you can use a native Pulumi provider in Pulumi HCL:
terraform {
required_providers {
kubernetes = {
source = "pulumi/kubernetes"
version = "4.33.0"
}
}
}
resource "kubernetes_yaml_config_file" "app" {
file = "app.yaml"
}
Our HCL interpreter routes this directly to the Pulumi Kubernetes package:
pulumirpc . PackageSpec {
Source : "kubernetes" ,
Version : "4.33.0" ,
}
Observe that the version moved from Parameters to the Version field. That’s because providers written with source = "pulumi/*" are talking about the plugin directly.
Both Pulumi programs and Terraform config exist to express a resource graph. It is their primary purpose. I don’t have the time or the pixels to explain everything, so I’ll restrict myself to three sub-topics here:
If you’ve ever looked at the schema behind any of our providers, you will observe that names are camelCase.
"aws:s3/bucket:Bucket" : {
"properties" : {
"accelerationStatus" : { "type" : "string" , "description" : "..." , "deprecationMessage" : "..." },
"bucketDomainName" : { "type" : "string" , "description" : "..." },
"bucketNamespace" : { "type" : "string" , "description" : "..." },
...
Our codegen actually relies on this fact, which means that even the schemas produced by terraform-provider use camelCase property names. The clever among you will see the problem: Terraform uses snake_case by convention. When Pulumi HCL sends a resource request with property values, it needs to translate from the snake_case text to camelCase for the engine, and then back again on values returned by the engine.
sequenceDiagram
participant Program as HCL program<br/>(snake_case)
participant Interp as HCL interpreter
participant Engine as Pulumi engine<br/>(camelCase)
Interp->>Engine: GetSchema("example_resource")
Engine-->>Interp: schema (property names + types)
Program->>Interp: some_value = "true"
Interp->>Engine: RegisterResource { someValue: true }
Engine-->>Interp: outputs { someValue: true, computedValue: "..." }
Interp-->>Program: example_resource.example.computed_value
Doing this correctly actually requires you to be aware of the resource’s schema; the raw value is not enough. Let’s work through an example to understand why:
resource "example_resource" "example" {
some_value = "true"
block {
inner_value = false
}
another_attribute = {
"trap_one" = 1
"trap_two" = 2
}
}
Here are two different value blobs (represented as JSON) that correspond to the inputs for "example" above. Both are valid, depending on the type of "example_resource" :
{
"someValue" : "true" ,
"block" : {
"innerValue" : false
},
"anotherAttribute" : {
"trap_one" : 1 ,
"trap_two" : 2
}
}
{
"someValue" : true ,
"block" : [{
"innerValue" : "false"
}],
"anotherAttribute" : {
"trapOne" : "1" ,
"trapTwo" : 2
}
}
Let’s break down the difference between the two potential value blobs:
HCL converts scalar values to the type the provider expects, so if someValue is typed as a string, we keep the "true" . If it’s typed as a boolean, we convert and pass the provider the converted value: true .
If block is typed as an object, the Pulumi provider expects to see an object back on the wire. If block is typed as a list of objects, then the Pulumi provider needs to see a list of objects back, even if it’s a list of one.
innerValue can have the same type conversion as someValue . Type conversion doesn’t stop at the top level.
If anotherAttribute is a map type, then its keys are user-provided values and need to be kept as they are. If it’s an object type, then its keys need to be shifted back to camelCase.
Pulumi’s HCL interpreter is thus type aware. It queries the engine for the schema of each provider it translates for and does the correct conversion, tracking the target type as it walks the value for translation. This is not limited to resources; it does the same translation for data sources, modules, and provider blocks.
For Pulumi providers that are bridged from Terraform providers, we apply an additional step. Bridged providers expose an explicit naming table used by pulumi convert --from terraform . Pulumi HCL queries the engine for this mapping also, and uses that to make sure that resource and property names line up exactly with the underlying Terraform provider, even if the property was renamed or hidden in Pulumi.
Terraform provisioners allow the Terraform config to specify commands to run at various points in the resource lifecycle. Conditions allow the same, but for boolean expressions to be evaluated instead. For example:
resource "example_resource" "example" {
input = var . input
lifecycle {
precondition {
condition = var . input > 3
error_message = "var.input must exceed 3"
}
}
provisioner "local-exec" {
command = "echo An example output is ${self.output} >> file.txt"
}
}
Pulumi only has one mechanism to hook into the resource lifecycle: hooks . Pulumi’s resource hooks are powerful enough to express both provisioners and conditions. The easiest way I can express this is by comparison to another language. I’ll use TypeScript. The example configuration block above is equivalent to this TypeScript progr

[truncated]
