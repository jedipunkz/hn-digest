---
source: "https://rsgm.dev/post/ai-dev-platform/"
hn_url: "https://news.ycombinator.com/item?id=48542433"
title: "My Homelab AI Dev Platform"
article_title: "My Homelab AI Dev Platform • Rsgm's Blog"
author: "rsgm"
captured_at: "2026-06-15T16:43:11Z"
capture_tool: "hn-digest"
hn_id: 48542433
score: 28
comments: 6
posted_at: "2026-06-15T15:09:40Z"
tags:
  - hacker-news
  - translated
---

# My Homelab AI Dev Platform

- HN: [48542433](https://news.ycombinator.com/item?id=48542433)
- Source: [rsgm.dev](https://rsgm.dev/post/ai-dev-platform/)
- Score: 28
- Comments: 6
- Posted: 2026-06-15T15:09:40Z

## Translation

タイトル: My Homelab AI 開発プラットフォーム
記事タイトル: My Homelab AI 開発プラットフォーム • Rsgm のブログ
説明: GitOps スタイルのホームラボ変更用の自己ホスティング OpenCode Web。

記事本文:
My Homelab AI 開発プラットフォーム • Rsgm のブログ
ホームタグ
検索 --> メニューを閉じる
2026 年 6 月 14 日公開
ホームラボの管理を容易にするために、Git アクセスを使用して OpenCode Web UI をセットアップしました。
OpenCode が Git にプッシュされ、私が PR を承認し、GitOps が変更をデプロイします。
何よりも、OpenCode は、デバイス間で同期された永続的なコーディング セッションを備えたサーバーとして実行されます。
私のホームラボのセットアップをすぐに共有します。
私が管理するサービスには、約 12 個の Docker Compose スタックがあります。
最近、GitOps で管理/デプロイできるように、それらを Arcane に移動しました。
次の当然のステップは、AI ツールを使用してサービスの維持を支援することでした。
最初に思いついた用途は、AI を使用してコンテナーの更新を支援することでした。
以前は、時間をかけて各サービスのリリース ノートを検索し、重大な変更がないか確認し、アップデートを実行していました。
各サービスに問題がないか手動でチェックします。
これには数時間を費やします。リリース ノートの概要を数分で読むことができるようになり、バージョン アップグレードがより簡単かつ安全になりました。
さらに、AI を使用してほとんどのコンテナにヘルスチェックを追加し、問題をより迅速に発見できるようにしました。
私は主に Claude Code を使用していましたが、最近 AI プロバイダーはトークン制限によって顧客から価値を搾り取っています。
そこで、この機会に他の選択肢を検討してみました。
ベンダーに依存せず、主要なプラグインでサポートされるものが欲しかったのです。
OpenCode で終了しました。
おそらく他にも適切なコーディング環境があると思いますが、これが私が試した中で私のお気に入りでした。
その後、組み込みの Web サーバーと Web UI が付属していることがわかり、アイデアが得られました。
基本的な開発ツールを使用して Truenas ホスト上に単純な VM をセットアップし、OpenCode Web サーバーを systemd ユニットとして追加しました。
ターミナル、ファイル ブラウザ、git が組み込まれた堅牢な環境です

差分、
同時に複数のコーディング セッションを管理するための git worktree サポートも含まれます。
さらに、OpenCode は、私がこれまでに見た中で最高のモバイル Web UI の質問/回答ポップアップを備えていました。
専用の SSH キーを使用して、Git サーバー上で OpenCode に独自のユーザーを与えました。
プロジェクトのクローンを作成してブランチをプッシュすることはできますが、デプロイ ブランチに直接プッシュすることはできません。
私のワークフローでは、AI が PR レビューの背後にあります。 OpenCode が変更を書き込み、私がそれを PR にマージします。
これはかわいいと思いますが、さらに重要なのは、レビューされていないコードがデプロイされるのを防ぐことです。
VM はインターネットにアクセスでき、Git サーバーにアクセスできますが、実際のサービスにはアクセスできません。
影響範囲が狭いため、ビルド ツールをインストールしたり依存関係をテストしたりする必要がある場合に、VM に OpenCode ルートを与えることに抵抗がありません。
これを実稼働開発者プラットフォームに組み込むことができそうです。開発者はプリインストールされたツールを使用して一時コンテナを利用できます。
ガードレールと監査ログにアクセスします。しかし、私にとっては、可動部分があまり多くなくても、必要なことは十分に機能します。
OpenCode の機能または改善を計画する (仕様、実装計画、セルフレビュー)
可能であれば変更をテストまたは検証します
気に入らない点については OpenCode を反復処理する
OpenCode は変更を機能ブランチにプッシュします
この支店の PR を開始します
満足できたら PR を統合します
GitOps がそこから引き継ぎます
- Docker サービスの変更には Arcane、ホーム アシスタントの設定変更には GitOps プラグイン、ブログの変更には Cloudflare Pages ワーカー
サービスを Truenas から Arcane GitOps プロジェクトに移行しました。
これは主に、以前 Truenas で実行していたすべての docker compose スタック用に git-backed ストレージを確保するためでした。
これが OpenCode の追加と連携してうまく機能することに驚きました。
すべてのコンテナにわたるネットワークを更新できること。たとえば、

私の携帯電話のおかげで、スプロールの管理がずっと簡単になりました。
以前は、すべての構成スタックを調べてネットワーク接続を追跡するには何時間もかかりました。
これで、目的を持ったコードベースで OpenCode を指定し、結果の PR 変更を確認してマージできるようになりました。
主に欠けている部分は CI フィードバックです。
GitHub では、コーディング エージェントにアクション ログを指定して、失敗したテスト、リンター エラー、スタック トレース、および IaC 計画の変更を診断できるようにするのが好きです。
これは、単体テストではカバーできない変更に対する高速フィードバック ループを維持するのに役立ちます。
Forgejo はそれをさらに難しくします。
Forgejo Actions は、パブリック API を通じてジョブログを公開しません。
文書化されていない API もありますが、私はそれらを中心に構築するつもりはありません。
この設定により、AI が変更中のサービスに直接アクセスできるようにすることなく、どのデバイスからでも自宅のインフラを変更できるようになります。
自分のコンピューターから変更を開始し、携帯電話から PR を確認し、GitOps にデプロイを処理させることができます。

## Original Extract

Self-hosting OpenCode Web for GitOps style homelab changes.

My Homelab AI Dev Platform • Rsgm's Blog
Home Tags
Search --> Close Show Menu
Published Jun 14, 2026
I set up OpenCode Web UI with Git access to make my homelab easier to manage.
OpenCode pushes to Git, I approve the PRs, GitOps deploys the changes.
Best of all, OpenCode runs as a server with persistent coding sessions synced across devices.
I’ll share my homelab setup soon.
There are about a dozen docker compose stacks for the services that I manage.
I recently moved them to Arcane so I can manage/deploy them with GitOps.
The next logical step was using AI tooling to help maintain my services.
The first use that came to mind was using AI to help with container updates.
Previously, I would spend time looking up the release notes for each of the services, checking for any breaking changes, running the updates,
and manually checking each of the services for issues.
I would spend a few hours on this. Now I can read a summary of the release notes in a few minutes, making version upgrades easier and safer.
On top of that, I’ve used AI to add healthchecks to most of the containers to make it faster to spot issues.
I mainly used Claude Code, but AI providers have been really squeezing the value out of customers recently through token limits,
so I took the opportunity to look into other options.
I wanted something that was vendor agnostic and supported by the major plugins.
I ended on OpenCode .
There are probably other decent coding environments, but this was my favorite of the ones I tried.
Then I found it ships with a built in webserver and web UI , which gave me an idea.
I set up a simple VM on the Truenas host with basic dev tooling and added OpenCode webserver as a systemd unit.
It’s a solid environment with a built in terminal, file browser, and git diffs,
as well as git worktree support for managing multiple coding sessions at the same time.
Plus, OpenCode had the best the question/answer popups in the mobile web UI that I’ve seen.
I gave OpenCode its own user on my Git server with dedicated SSH keys.
It can clone projects and push branches, but it cannot push straight to the deploy branch.
My workflow keeps the AI behind PR review. OpenCode writes the change and I merge it myself in a PR.
I think it’s cute, but more importantly, it keeps unreviewed code from getting deployed.
The VM has internet access and access to my Git server, but it cannot reach my actual services.
Because the blast radius is small, I am comfortable giving OpenCode root on the VM when it needs to install build tools or test dependencies.
I could see building this into a production developer platform. Ephemeral containers available to developers with preinstalled tooling,
access guardrails, and audit logs. But for me, it does what I need it to without too many moving parts.
Plan out a feature or improvement in OpenCode (spec, implementation plan, and self-reviews)
I’ll test or verify changes if possible
Iterate with OpenCode on things I don’t like
OpenCode pushes changes to a feature branch
I’ll open a PR for this branch
I’ll merge the PR once I’m happy
GitOps takes over from there
- Arcane for docker service changes, GitOps plugin for Home Assistant config changes, Cloudflare Pages worker for blog changes
I migrated my services from Truenas to Arcane GitOps projects.
This was mainly to have git-backed storage for all the docker compose stacks I was running in Truenas previously.
I was surprised how well this worked in conjunction with adding OpenCode.
Being able to update the networking across all containers, for example, from my phone makes the sprawl much easier to manage.
Before it would take hours to comb through all of the compose stacks, tracing out network connectivity.
Now I can point OpenCode at the codebase with a goal, check the resulting PR changes, and merge.
The main missing piece is CI feedback.
On GitHub, I like pointing a coding agent at Actions logs so it can diagnose failing tests, linter errors, stack traces, and IaC plan changes.
This helps maintain a fast feedback loop for changes that unit tests don’t cover.
Forgejo makes that harder.
Forgejo Actions does not expose job logs through the public API.
There are undocumented APIs, but I would rather not build around those.
This setup lets me make home infra changes from any device without giving AI direct access to the services it’s changing.
I can start a change from my computer, review the PR from my phone, and let GitOps handle the deploy.
