---
source: "https://codedbearder.com/posts/nixidy-part-1-introduction/"
hn_url: "https://news.ycombinator.com/item?id=48446465"
title: "Introduction to nixidy – Kubernetes GitOps with Nix"
article_title: "nixidy part 1: Introduction to nixidy | codedbearder"
author: "granra"
captured_at: "2026-06-08T16:29:08Z"
capture_tool: "hn-digest"
hn_id: 48446465
score: 2
comments: 0
posted_at: "2026-06-08T15:16:19Z"
tags:
  - hacker-news
  - translated
---

# Introduction to nixidy – Kubernetes GitOps with Nix

- HN: [48446465](https://news.ycombinator.com/item?id=48446465)
- Source: [codedbearder.com](https://codedbearder.com/posts/nixidy-part-1-introduction/)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T15:16:19Z

## Translation

タイトル: nixidy の紹介 – Nix を使用した Kubernetes GitOps
記事タイトル: nixidy パート 1: nixidy の概要 |コード化されたベアダー
説明: コードベアダー

記事本文:
nixidy パート 1: nixidy の概要 |コード化されたベアダー
コード化されたベアダー
nixidy パート 1: nixidy の概要
私は ArgoCD を使用して Kubernetes の多くの GitOps リポジトリを管理してきましたが、600 行の YAML である Helm 値オーバーライド ファイルを開いたものの、実際にレンダリングされたマニフェストにどの値が含まれているかがまだわからないのは私だけではないと思います。私は helm template を実行し、 grep を介してパイプし、諦めてとにかくコミットしましたが、ステージングの差分が私が見逃したものをキャッチできることを望みました。
あなたが思っていることとのギャップ
デプロイ中で、実際にクラスターに配置されるのは、まさに nixidy です
閉じることを目的としています。 Helm 値ファイル、KusTOMize オーバーレイ、生の YAML を環境ごとに 1 つの Nix 式に置き換えるためにこれを作成しました。すべてのフィールドは入力され、すべてのビルドは再現可能で、出力は git diff で実行できるプレーンな YAML です。
クラスターに接触する前に。
このパートが終わるまでに、nginx デプロイメントとサービスを定義し、Argo CD アプリケーションを生成する、動作する nixidy プロジェクトが完成します。
マニフェストは自動的に作成され、GitOps を通じてクラスターにデプロイされます。
dev という nixidy 環境を構築します。
Argo CD 経由でクラスターにデプロイされた 1 つのアプリケーションが含まれています。プロジェクト構造は、運用クラスター全体を管理するために拡張するスケルトンになります。
ニックス
フレークを有効にしてインストール (ダウンロード)
Kubernetes クラスター
アルゴCD付き
インストールされています
Git リポジトリ
Kubernetes マニフェスト (GitHub、GitLab など) 用
基本的な知識
Kubernetes のデプロイメント、サービス、および名前空間を使用する
基本的な知識
Argo CD アプリケーションを使用
リソース
nixidy はレンダリングされたマニフェスト パターンを実装します
CI がプレーン YAML を生成し、それを PR でレビューし、Argo CD がそれをデプロイします。以前に生の YAML または KusTOMize で Argo CD を使用したことがある場合、デプロイメント

側面は同一です。違いは完全に YAML の生成方法にあります。
Kubernetes マニフェストを構築する Nix 式
nixidy の背後にある中心的な考え方は、すべての Kubernetes リソースが型指定された Nix オプションであるということです。デプロイメントは YAML の BLOB ではなく、レプリカが配置される構造化された属性セットです。
整数、画像です
は文字列であり、セレクターのタイプミスです
これはビルド エラーであり、実行時の予期せぬ現象ではありません。
プロジェクトを作成することから始めましょう。
mkdir my-cluster && cd my-cluster
gitの初期化
次に、 flake.nix を作成します。これは、 nixidy を Nix フレークに接続するエントリ ポイントです。
{
description = "nixidy で管理された Kubernetes クラスター" ;
入力 = {
nixpkgs 。 url = "github:nixos/nixpkgs/nixos-unstable" ;
フレークユーティリティ 。 URL = "github:numtide/flake-utils" ;
ニクディ 。 URL = "github:arnarg/nixidy" ;
};
出力 = {
nixpkg、
フレークユーティリティ、
ニクディ、
...
}:
フレークユーティリティ 。ライブラリ 。 eachDefaultSystem (システム: let
pkgs = import nixpkgs {システムを継承;};
{で
nixidyEnvs = nixidy 。ライブラリ 。 mkEnvs {
パッケージを継承します。
環境 。開発者。モジュール = [ ./env/dev.nix ];
};
});
}
注目すべき点がいくつかあります:
nixidy.lib.mkEnvs
一連の名前付き環境を取得し、YAML マニフェストを構築する Nix 派生を返します。主要な開発者
は、 .#dev で参照する属性になります。
各環境は、プレーンな .nix である NixOS スタイルのモジュールのリストを受け取ります。
オプションを設定するファイル。これは、NixOS を強化するのと同じモジュール システムです。つまり、 imports 、 lib.mkDefault 、 lib.mkForce 、および期待されるすべての合成プリミティブを取得できます。
以前に NixOS システムを構成したことがある場合、形状は同じです。オプションを設定するモジュールのリストがモジュール システムによってマージされています。違いは、オプションがシステム サービスではなく Kubernetes リソースを説明することです。
1 つのアプリケーションを備えた環境モジュール
次に、環境ディレクトリを作成しましょう。

d 開発モジュール:
mkdir -p 環境
env/dev.nix と書き込みます。リポジトリ URL を必ず独自の URL に置き換えてください (ここで、nixidy は Argo CD にレンダリングされたマニフェストを探すように指示します)。
{
ニクディ 。ターゲット。リポジトリ = "https://github.com/YOUR_USERNAME/my-cluster.git" ;
ニクディ 。ターゲット。ブランチ = "メイン" ;
ニクディ 。ターゲット。 rootPath = "./manifests/dev" ;
アプリケーション。 nginx = {
名前空間 = "nginx" ;
createNamespace = true ;
リソース = {
展開。 nginx 。仕様 = {
レプリカ = 2 ;
セレクター。 matchLabels 。アプリ = "nginx" ;
テンプレート = {
メタデータ 。ラベル 。アプリ = "nginx" ;
仕様。コンテナ。 nginx = {
画像 = "nginx:1.25.1" ;
ポート。 http 。コンテナポート = 80 ;
};
};
};
サービス。 nginx 。仕様 = {
セレクター。アプリ = "nginx" ;
ポート。 http 。ポート = 80 ;
};
};
};
}
これが何を宣言しているかを見てみましょう:
nixidy.target.* : 生成された YAML が Git リポジトリ内で終了する場所。 Argo CD アプリケーション
マニフェストはこのリポジトリ、ブランチ、パスを参照します。
apps.nginx : 1 つの論理アプリケーション。アプリケーションは、出力内に独自のディレクトリと独自の Argo CD アプリケーションを取得します。
マニフェストします。
namespace = "nginx" : このアプリケーションのすべてのリソースは nginx にデプロイされます
名前空間。
createNamespace = true : Nixidy が名前空間を生成します
自動的にマニフェストします。これがないと、アウトオブバンドで名前空間を作成する必要があります。
resource.deployments.nginx : 型指定されたデプロイメント。仕様
属性は Kubernetes デプロイメント仕様に従いますが、Nix 評価時に適用されます。
resource.services.nginx : 型付きサービス、同じ考え方。
タイプエラーはビルドエラーになります。レプリカを「2」に設定します
上記のモジュールと nixidy ビルド内
導入ロールアウト後 15 分ではなく、すぐに失敗します。
構成。 prod.nix を追加するとき
これと同じモジュールをインポートし、replicas = lib.mkForce 10 を設定すると、「同じ」を表現しています

YAML ファイル全体をコピーして 1 つの数値を変更する代わりに、2 行で「アプリ、異なるスケール」を追加します。 NixOS モジュール システム ( imports 、 lib.mkDefault 、 lib.mkForce ) はこれを無料で提供します。これは、複数環境の NixOS 構成を処理するのと同じメカニズムです。
nix run github:arnarg/nixidy -- build .#dev
情報
最初の実行では、nixidy とその依存関係が Nix ストアにダウンロードされます。何も変更されなければ、後続の実行は即座に行われます。
結果/
§── アプリ/
│ └── アプリケーション-nginx.yaml
━─ nginx/
§── デプロイメント-nginx.yaml
§── 名前空間-nginx.yaml
└── サービス-nginx.yaml
生成された Deployment を確認します。
cat result/nginx/Deployment-nginx.yaml
apiVersion : apps/v1
種類：展開
メタデータ:
名前：nginx
名前空間 : nginx
仕様：
レプリカ：2
セレクター:
matchLabels :
アプリ：nginx
テンプレート:
メタデータ:
ラベル:
アプリ：nginx
仕様：
コンテナ:
- 画像：nginx:1.25.1
名前：nginx
ポート:
- コンテナポート: 80
名前： http
そしてArgo CDアプリケーション
nixidy があなたのために生成したもの:
cat result/apps/Application-nginx.yaml
apiバージョン: argoproj.io/v1alpha1
種類 : アプリケーション
メタデータ:
名前：nginx
名前空間 : argocd
仕様：
目的地:
名前空間 : nginx
サーバー: https://kubernetes.default.svc
プロジェクト: デフォルト
ソース:
パス: ./manifests/dev/nginx
リポURL : https://github.com/YOUR_USERNAME/my-cluster.git
targetRevision : main
すべてのアプリケーション。*
ブロックは Argo CD アプリケーションを 1 つだけ生成します
レンダリングされたマニフェストを含むディレクトリを指します。これはレンダリングされたマニフェスト パターンです。Argo CD は、テンプレートや Helm リリースではなく、プレーン YAML を同期し、クラスターの状態と比較できる静的ファイルのみを同期します。
ニクシディスイッチ
コマンドは、ビルドされたマニフェストを rootPath のリポジトリにコピーします。
設定した内容:
nix run github:arnarg/nixidy -- switch .#dev

これにより、./manifests/dev/ が作成されます。
result/ と同じディレクトリツリーを持ちます。コミットしてプッシュします。
git add 。
git commit -m "nixidy 経由で nginx アプリケーションを追加"
gitプッシュ
レンダリングされた YAML がリポジトリに追加されました。アルゴCDで見れます。
Argo CD がすでにクラスター内で実行されている場合、1 つのコマンドで「アプリのアプリ」(親アプリケーション) が作成されます。
すべての nixidy アプリケーションを管理します):
nix run github:arnarg/nixidy -- bootstrap .#dev | kubectl apply -f -
これにより、Argo CD アプリケーションが出力されます。
マニフェスト/dev/apps/ を指すマニフェスト
あなたのリポジトリに。 Argo CD はそのディレクトリを読み取り、 Application-nginx.yaml を検出し、 nginx Application を作成します。これにより、デプロイメント、サービス、および名前空間がクラスターに同期されます。
または: 直接申請します (テスト用)
Argo CD を一時的にスキップしたい場合は、ローカルの種類
たとえばクラスター:
nix run github:arnarg/nixidy -- apply .#dev
これにより、kubectl apply --prune が実行されます。
これにより、nixidy 構成から削除されたリソースは、次回の適用時にクラスターからも削除されます (リソースが削除されている場合)。
現在、1 つの環境に 1 つのアプリケーションがあります。実際のクラスターには、開発、ステージング、本番環境にまたがる多数のアプリケーションがあり、同じデプロイメントを 3 つのファイルにコピーアンドペーストしたくありません。パート 2 では
nginx アプリケーションを共有モジュールにリファクタリングし、レプリカをオーバーライドします
lib.mkDefault を使用した環境ごと
と lib.mkForce を統合し、タイプ セーフティを放棄することなく Helm チャートを統合します。

## Original Extract

codedbearder

nixidy part 1: Introduction to nixidy | codedbearder
codedbearder
nixidy part 1: Introduction to nixidy
I have managed many GitOps repositories for Kubernetes with ArgoCD and I'm sure I'm not alone in having opened a Helm values override file that was 600 lines of YAML and still not being sure which values actually made it into the rendered manifests. I've run helm template , piped it through grep , given up, committed it anyway and hoped the staging diff would catch anything my eyes missed.
That gap between what you think
you're deploying and what actually lands in the cluster is exactly what nixidy
is meant to close. I wrote it to replace Helm value files, Kustomize overlays, and raw YAML with a single Nix expression per environment. Every field is typed, every build is reproducible, and the output is plain YAML you can git diff
before it ever touches a cluster.
By the end of this part we'll have a working nixidy project that defines an nginx Deployment and Service, generates Argo CD Application
manifests automatically, and deploys to your cluster through GitOps.
We're going to build a nixidy environment called dev
containing one application deployed to your cluster via Argo CD. The project structure will be the skeleton you'd extend to manage an entire production cluster.
Nix
installed with flakes enabled ( download )
A Kubernetes cluster
with Argo CD
installed
A Git repository
for your Kubernetes manifests (GitHub, GitLab, etc.)
Basic familiarity
with Kubernetes Deployments, Services, and Namespaces
Basic familiarity
with Argo CD Application
resources
nixidy implements the Rendered Manifests Pattern
where your CI generates plain YAML, you review it in PRs, and Argo CD deploys it. If you've used Argo CD with raw YAML or Kustomize before, the deployment side is identical. The difference is entirely in how the YAML is produced .
A Nix expression that builds a Kubernetes manifest
The core idea behind nixidy is that every Kubernetes resource is a typed Nix option. A Deployment isn't a blob of YAML, it's a structured attribute set where replicas
is an integer, image
is a string, and a typo in selector
is a build error, not a runtime surprise.
Let's start by creating the project:
mkdir my-cluster && cd my-cluster
git init
Now create flake.nix , this is the entry point that wires nixidy into your Nix flake:
{
description = "My Kubernetes cluster managed with nixidy" ;
inputs = {
nixpkgs . url = "github:nixos/nixpkgs/nixos-unstable" ;
flake-utils . url = "github:numtide/flake-utils" ;
nixidy . url = "github:arnarg/nixidy" ;
};
outputs = {
nixpkgs,
flake-utils,
nixidy,
...
}:
flake-utils . lib . eachDefaultSystem (system: let
pkgs = import nixpkgs { inherit system;};
in {
nixidyEnvs = nixidy . lib . mkEnvs {
inherit pkgs;
envs . dev . modules = [ ./env/dev.nix ];
};
});
}
A couple things worth noting:
nixidy.lib.mkEnvs
takes a set of named environments and returns Nix derivations that build YAML manifests. The key dev
becomes the attribute you reference with .#dev .
Each environment takes a list of NixOS-style modules which are plain .nix
files that set options. This is the same module system that powers NixOS, which means you get imports , lib.mkDefault , lib.mkForce , and all the composition primitives you'd expect.
If you've configured a NixOS system before, the shape is identical: a list of modules that set options, merged by the module system. The difference is that the options describe Kubernetes resources instead of system services.
An environment module with one application
Now let's create the environment directory and the dev module:
mkdir -p env
Write env/dev.nix . Make sure to replace the repository URL with your own (this is where nixidy will tell Argo CD to look for rendered manifests):
{
nixidy . target . repository = "https://github.com/YOUR_USERNAME/my-cluster.git" ;
nixidy . target . branch = "main" ;
nixidy . target . rootPath = "./manifests/dev" ;
applications . nginx = {
namespace = "nginx" ;
createNamespace = true ;
resources = {
deployments . nginx . spec = {
replicas = 2 ;
selector . matchLabels . app = "nginx" ;
template = {
metadata . labels . app = "nginx" ;
spec . containers . nginx = {
image = "nginx:1.25.1" ;
ports . http . containerPort = 80 ;
};
};
};
services . nginx . spec = {
selector . app = "nginx" ;
ports . http . port = 80 ;
};
};
};
}
Let me walk through what this declares:
nixidy.target.* : Where the generated YAML ends up in your Git repo. Argo CD Application
manifests will reference this repo, branch, and path.
applications.nginx : One logical application. An application gets its own directory in the output and its own Argo CD Application
manifest.
namespace = "nginx" : All resources in this application are deployed to the nginx
namespace.
createNamespace = true : Nixidy generates a Namespace
manifest automatically. Without this, you'd need to create the namespace out-of-band.
resources.deployments.nginx : A typed Deployment. The spec
attribute follows the Kubernetes Deployment spec , but enforced at Nix evaluation time.
resources.services.nginx : A typed Service, same idea.
Type errors become build errors. Set replicas = "two"
in the module above and nixidy build
fails immediately, not 15 minutes into a deployment rollout.
Composition. When you add a prod.nix
that imports this same module and sets replicas = lib.mkForce 10 , you're expressing "same app, different scale" in two lines instead of copying an entire YAML file and changing one number. The NixOS module system ( imports , lib.mkDefault , lib.mkForce ) gives you this for free, and it's the same mechanism that handles multi-environment NixOS configs.
nix run github:arnarg/nixidy -- build .#dev
info
The first run downloads nixidy and its dependencies into the Nix store. Subsequent runs are instant if nothing changed.
result/
├── apps/
│ └── Application-nginx.yaml
└── nginx/
├── Deployment-nginx.yaml
├── Namespace-nginx.yaml
└── Service-nginx.yaml
Look at the generated Deployment:
cat result/nginx/Deployment-nginx.yaml
apiVersion : apps/v1
kind : Deployment
metadata :
name : nginx
namespace : nginx
spec :
replicas : 2
selector :
matchLabels :
app : nginx
template :
metadata :
labels :
app : nginx
spec :
containers :
- image : nginx:1.25.1
name : nginx
ports :
- containerPort : 80
name : http
And the Argo CD Application
that nixidy generated for you:
cat result/apps/Application-nginx.yaml
apiVersion : argoproj.io/v1alpha1
kind : Application
metadata :
name : nginx
namespace : argocd
spec :
destination :
namespace : nginx
server : https://kubernetes.default.svc
project : default
source :
path : ./manifests/dev/nginx
repoURL : https://github.com/YOUR_USERNAME/my-cluster.git
targetRevision : main
Every applications.*
block produces exactly one Argo CD Application
pointing at the directory containing its rendered manifests. This is the rendered manifests pattern: Argo CD syncs plain YAML, not templates, not Helm releases, just static files it can diff against the cluster state.
The nixidy switch
command copies the built manifests into your repository at the rootPath
you configured:
nix run github:arnarg/nixidy -- switch .#dev
This creates ./manifests/dev/
with the same directory tree as result/ . Commit and push:
git add .
git commit -m "Add nginx application via nixidy"
git push
The rendered YAML is now in your repository. Argo CD can see it.
If Argo CD is already running in your cluster, one command creates an "app of apps" (a parent Application
that manages all your nixidy applications):
nix run github:arnarg/nixidy -- bootstrap .#dev | kubectl apply -f -
This outputs an Argo CD Application
manifest that points at manifests/dev/apps/
in your repo. Argo CD reads that directory, discovers Application-nginx.yaml , creates the nginx Application , which then syncs the Deployment, Service, and Namespace into your cluster.
Or: apply directly (for testing)
If you want to skip Argo CD temporarily, a local kind
cluster for instance:
nix run github:arnarg/nixidy -- apply .#dev
This runs kubectl apply --prune
with the correct label selectors, so resources removed from your nixidy config are also removed from the cluster on the next apply (if resources have been removed).
We now have one application in one environment. Real clusters have a dozen applications across dev, staging, and production and I don't want to copy-paste the same Deployment into three files. In Part 2
we'll refactor the nginx application into a shared module, override replicas
per environment with lib.mkDefault
and lib.mkForce , and integrate a Helm chart without giving up type safety.
