---
source: "https://cretezy.com/2026/aur-ai-security/"
hn_url: "https://news.ycombinator.com/item?id=49149119"
title: "Reviewing the Arch User Repository with AI"
article_title: "Reviewing the Arch User Repository with AI — Cretezy"
author: "CraftThatBlock"
captured_at: "2026-08-02T22:45:30Z"
capture_tool: "hn-digest"
hn_id: 49149119
score: 1
comments: 1
posted_at: "2026-08-02T22:35:13Z"
tags:
  - hacker-news
  - translated
---

# Reviewing the Arch User Repository with AI

- HN: [49149119](https://news.ycombinator.com/item?id=49149119)
- Source: [cretezy.com](https://cretezy.com/2026/aur-ai-security/)
- Score: 1
- Comments: 1
- Posted: 2026-08-02T22:35:13Z

## Translation

タイトル: AI を使用した Arch ユーザー リポジトリのレビュー
記事タイトル: AI を使用した Arch ユーザー リポジトリのレビュー — Cretezy
説明: AI 支援による AUR パッケージのセキュリティ レビュー パイプライン。バージョン履歴、リポジトリの差分、検査できる証拠が含まれます。

記事本文:
AI を使用した Arch ユーザー リポジトリのレビュー — Cretezy Cretezy — テクノロジーとプログラミング 2026 年 8 月 2 日 開発者ツール Arch Linux AI AI を使用した Arch ユーザー リポジトリのレビュー
Arch ユーザー リポジトリは、私が興味を持った理由の 1 つです。
Arch Linux を使用するのと同じです。ソフトウェアが存在する場合、可能性は十分にあります
誰かがすでにそれ用のPKGBUILDを作成しています。
この利便性は、特定の信頼モデルによって実現されます。PKGBUILD はシェル コードです。
別のユーザーによって書かれたものであり、パッケージのインストールはそれを実行することを意味します。 6月
2026 年、それは抽象的な警告ではなくなりました。
6 月 12 日、Arch Linux は
アクティブな悪意のあるパッケージに関する公式通知
事件、
大量の悪意のあるパッケージの導入と更新について説明します。チーム
悪意のあるコミットを削除し、アカウント作成中に新しいコミットを阻止しようとしていました
作成、パッケージの更新、およびパッケージの採用が制限されていました。
それは一度限りの出来事ではありませんでした。 7月31日、
Phoronix は、悪意のあるパッケージの新たな波と Arch が停止したことを報告しました
パッケージ採用。
前回のキャンペーンでは 1,500 以上のパッケージが影響を受けましたが、新たなキャンペーンは
すでにさらに数十件が含まれています。チームが対応している間、Arch は養子縁組を無効にしました
流入が激しく、ユーザーに警戒を続けるよう改めて呼びかけた。
Arch のアドバイスは直接的でした。PKGBUILD とインストール スクリプトの変更をすべて確認してください。
AUR パッケージを更新するとき、特にインシデントがアクティブな間。
それらの出来事が私がこのプロジェクトを始めた具体的な理由です。問題は
誰も聞いたことのない明らかに偽物の新しいパッケージに限定されません。おなじみの
または孤立したパッケージが採用され、悪意のあるアップデートを受信して到着する可能性があります。
退屈なバージョンアップと同じ更新ワークフローを通じて。
すべての更新を手動で読み取ると、拡張性があまり高くありませんが、回答が減少します
「AI が安全だと言いました」も役に立ちません。私

別のレビュー層が欲しかった
まさにこれらの変更については、現在のパッケージの更新を自動的に検査します。
レビューされた内容を正確に保存し、証拠を簡単に検査できるようにする
その後。
AUR のほとんどの変更は依然として完全に通常のバージョン バンプです。時々、
ダウンロードが別のドメインに移動するか、リポジトリの所有者が変更されるか、または新しいコマンドが実行される
梱包ステップで表示されます。事前に気づいておきたい変化です
アップデートをインストールしています。
それはこうなった
aur_ai_security 、小さな Rust
AUR パッケージ バージョンのインデックスを作成し、AI を使用して各パッケージをレビューするサービス
の Git diff を更新し、結果を Web インターフェイスに提供します。
現在の結果を参照できます。
ライブデモ。
プロジェクトはまだ初期段階にありますが、インデックスからレビューまでの完全なフローは機能しています。これ
この投稿では、記録される内容、チェックの仕組み、さらに時間がかかった部分について説明します。
予想以上の繰り返し。
PKGBUILD のみをレビューする場合の問題
現在の PKGBUILD を確認することは出発点としては役に立ちますが、最も重要な点が失われます。
重要な背景: 何が変わったのか?
確立されたアップストリーム GitHub リリースからバイナリをダウンロードするパッケージは、
普通の。同じパッケージを無関係なドメインから突然ダウンロードするのは非常に危険です
もっと面白い。バージョンが変更されると、新しいチェックサムが期待されます。新しい
リポジトリの所有者または配信メカニズムは、たとえシェルであっても注目に値します。
コードはまだきれいに見えます。
AUR はすでに各パッケージ ベースを Git に保存しているため、レビューの便利な単位
は単なるファイルではありません。これはパッケージのバージョン、そのリポジトリのコミット、完全なバージョンです。
PKGBUILD 、およびそれを導入した差分。
aur_ai_security には、インデクサー、AI チェッカー、Web という 3 つの主要な部分があります。
結果をカタログ化してレビューするために使用されるインターフェイス。
update-index コマンドは、AUR メタデータ インデックスをダウンロードし、新たに追加します。
SQLite のパッケージ バージョンを確認しました:
カーゴラン -p aur_

ai_security -- インデックスの更新
パッケージのメタデータのスコープはパッケージ名とバージョンです。インデックスは AUR を維持します
パッケージおよびパッケージベースの ID、送信者、最終変更時刻、スナップショットのパス、
人気。最新のインデックスに存在するバージョンは現行としてマークされますが、古いものとしてマークされます。
行は履歴チェックに引き続き使用できます。
パッケージのインデックスとチェック結果は現在 SQLite データベースに保存されています
CLI と Web サイトで共有されます。
check コマンドは、まだチェックされていない現在のバージョンを選択します。
選択したプロバイダーとモデルで。正確なパッケージ名に限定したり、
タイムスタンプの後に変更されたパッケージ:
カーゴ ラン -p aur_ai_security -- チェック \
--プロバイダー コーデックス \
--モデル gpt-5.6-luna \
-- 24 時間以降 \
--filter netcatty-bin zed-preview-bin
選択されたパッケージごとに、チェッカーは次のことを行います。
AUR Git リポジトリのクローンを作成します。
最初にPKGBUILDを使用してコミット差分を構築し、次にもう一方の変更されたコミットを構築します
ファイル。
そのコンテキストを選択した AI プロバイダーに送信します。
判定、説明、ソース、差分、コミット、プロバイダー、モデルを保存します。
プロジェクトでは AI エージェントとして Rig を使用しています
ループにより、チェッカーを 1 つに結び付けることなく、複数のプロバイダーをサポートできるようになります。
API。現在、OpenAI、Anthropic、OpenRouter がサポートされています。エージェントが持っているのは、
単一の read_file ツールにより、複製されたリポジトリ内の他のファイルを検査できます
diff にさらにコンテキストが必要な場合。
Codex CLI もサポートしているため、
別の API アカウントの代わりに ChatGPT サブスクリプションを使用します。コーデックスは内部から実行されます
クローン作成された AUR リポジトリのシェル ツールと Web 検索は無効になっています。
Web アプリケーションは次のように構築されています
Topcoat 、サーバーレンダリングされた Rust Web
フレームワークとテイルウィンド。パッケージの検索可能なカタログを提供し、
各判定の背後にあるソースと差分を確認できるようにチェックします。
同じSQLを読み込みます

ite データベースをインデクサーおよびチェッカーとして使用します。次のように実行します。
トップコート開発 --package aur_ai_security_web
そこから、パッケージを検索して最新の評価を確認できます。
履歴、PKGBUILD、更新差分をブラウザで確認できます。
これらのチェックを行うために、paru フォークに 2 つの実験的なブランチも構築しました。
パッケージ インストール ワークフローに追加します。
aur-ai-セキュリティ-リモート
既存の評価についてホストされた検索 API をクエリします。
aur-ai-セキュリティ-ローカル
チェッカー クレートを直接統合し、設定されたプロバイダーを実行し、
現地でモデルを作ります。
どちらも、paru が AUR リポジトリをダウンロードした後、AUR リポジトリを開始する前に実行されます。
事前ビルド コマンドまたはパッケージ ビルド。
各チェックは、次の 3 つの判定のいずれかを返します。
通常の梱包動作では安全です。
具体的な変更が人間によるレビューに値する場合は疑わしい。
悪意のある行為の強い証拠がある場合は危険です。
疑わしく危険な結果には説明が必要です。安全な結果は得られません
通常のパッケージが正常に見えるというフィラーテキストが必要です。
プロンプトは、同じ上流ソースからのバージョンとチェックサムのバンプを次のように扱います。
通常の動作に加えて、検証済みのアップストリーム バイナリを使用して補完を生成します
またはメタデータ。ドメイン、リポジトリ所有者、またはダウンロード メカニズムの変更
もっと注目を集めてください。
これらの判定はいずれも、荷物が安全であることを証明するものではありません。モデルは悪意のあるものを見逃す可能性がある
動作、誤検知を生成するか、それ自体を持つアップストリーム バイナリを信頼するか
侵害されました。プロジェクトは各結果の背後にある証拠を保存するため、
検査を受ける。中身ではなく、パッケージとサプライチェーンのシグナルをレビューします
ダウンロードされたバイナリの。
当面の作業は、チェッカーを継続的に実行し、通知を追加することです。
危険な結果を検出し、モデルが存在する場所のプロンプトを改善します。
常にうるさすぎるか、信頼しすぎます。
目標はいっぱいではありません

荷物が安全かどうかの判断は AI を信頼します。それ
それを早期警告レイヤーとして使用することです。アップデートが明らかに悪意のあるものである場合、
すぐに知りたい、判決の背後にある正確な情報源を確認したい、そしてできるようになりたい
さらに多くのユーザーが影響を受ける前に対処してください。
こんにちは！私の名前は Cretezy で、ソフトウェア開発者です。プログラミングや個人的なことについて書いています
プロジェクトなど。

## Original Extract

An AI-assisted security review pipeline for AUR packages, with version history, repository diffs, and evidence you can inspect.

Reviewing the Arch User Repository with AI — Cretezy Cretezy — Technology & Programming Aug 2, 2026 Developer Tools Arch Linux AI Reviewing the Arch User Repository with AI
The Arch User Repository is one of the reasons I
like using Arch Linux. If a piece of software exists, there is a good chance
someone has already written a PKGBUILD for it.
That convenience comes with a specific trust model: a PKGBUILD is shell code
written by another user, and installing a package means running it. In June
2026, that stopped being an abstract warning.
On June 12, Arch Linux published an
official notice about an active malicious-packages
incident ,
describing a high volume of malicious package adoptions and updates. The team
was removing malicious commits and trying to prevent new ones while account
creation, package updates, and package adoption were restricted.
It was not a one-off incident. On July 31,
Phoronix reported another wave of malicious packages and that Arch had halted
package adoptions .
The previous campaign had affected more than 1,500 packages, and the new wave
already included dozens more. Arch disabled adoptions while the team handled
the influx and again asked users to stay vigilant.
Arch’s advice was direct: review every PKGBUILD and install-script change
when updating AUR packages, especially while the incident was active.
Those incidents are the specific reason I started this project. The problem was
not limited to an obviously fake new package nobody had heard of. A familiar
or orphaned package could be adopted, receive a malicious update, and arrive
through the same update workflow as a boring version bump.
Reading every update by hand does not scale very well, but reducing the answer
to “an AI said it was safe” is not useful either. I wanted another review layer
for exactly these changes: automatically inspect current package updates,
preserve exactly what was reviewed, and make the evidence easy to inspect
afterward.
Most AUR changes are still completely ordinary version bumps. Occasionally a
download moves to another domain, a repository changes owners, or a new command
appears in a packaging step. Those are the changes I want to notice before
installing an update.
That became
aur_ai_security , a small Rust
service that indexes AUR package versions, uses AI to review each package
update’s Git diff, and serves the results in a web interface.
You can browse the current results in the
live demo .
The project is still early, but the complete index-to-review flow works. This
post goes over what it records, how a check works, and the parts that took more
iteration than I expected.
The problem with reviewing only a PKGBUILD
Looking at the current PKGBUILD is a useful start, but it loses the most
important context: what changed?
A package downloading a binary from its established upstream GitHub release is
normal. The same package suddenly downloading from an unrelated domain is much
more interesting. A new checksum is expected when the version changes. A new
repository owner or delivery mechanism deserves attention even if the shell
code still looks clean.
The AUR already stores each package base in Git, so the useful unit of review
is not just a file. It is a package version, its repository commit, the full
PKGBUILD , and the diff that introduced it.
aur_ai_security has three main parts: the indexer, the AI checker, and the web
interface used to catalog and review the results.
The update-index command downloads the AUR metadata index and appends newly
seen package versions to SQLite:
cargo run -p aur_ai_security -- update-index
Package metadata is scoped to package name and version. The index keeps the AUR
package and package-base IDs, submitter, last-modified time, snapshot path, and
popularity. Versions present in the latest index are marked current, while old
rows remain available for historical checks.
The package index and check results are currently stored in a SQLite database
shared by the CLI and website.
The check command selects current versions that have not already been checked
with the chosen provider and model. It can be limited to exact package names or
to packages modified after a timestamp:
cargo run -p aur_ai_security -- check \
--provider codex \
--model gpt-5.6-luna \
--since 24h \
--filter netcatty-bin zed-preview-bin
For each selected package, the checker:
clones its AUR Git repository;
builds a commit diff with PKGBUILD first, followed by the other changed
files;
sends that context to the selected AI provider;
stores the verdict, explanation, source, diff, commit, provider, and model.
The project uses Rig for its AI agent
loop, which lets it support multiple providers without tying the checker to one
API. OpenAI, Anthropic, and OpenRouter are currently supported. The agent has a
single read_file tool so it can inspect other files in the cloned repository
when the diff needs more context.
It also supports the Codex CLI, making it possible to run checks using a
ChatGPT subscription instead of a separate API account. Codex runs from inside
the cloned AUR repository with its shell tool and web search disabled.
The web application is built with
Topcoat , a server-rendered Rust web
framework, and Tailwind. It provides a searchable catalog of packages and
checks, with the source and diff behind each verdict available for review.
It reads the same SQLite database as the indexer and checker. Run it with:
topcoat dev --package aur_ai_security_web
From there, packages can be searched and their latest assessment, check
history, PKGBUILD , and update diff can be reviewed in the browser.
I also built two experimental branches in my paru fork to bring these checks
into the package installation workflow:
aur-ai-security-remote
queries the hosted lookup API for existing assessments;
aur-ai-security-local
integrates the checker crate directly and runs the configured provider and
model locally.
Both run after paru downloads the AUR repositories and before it starts any
pre-build commands or package builds.
Each check returns one of three verdicts:
safe for ordinary packaging behavior;
suspicious when a concrete change deserves human review;
dangerous when there is strong evidence of malicious behavior.
Suspicious and dangerous results require an explanation. Safe results do not
need filler text saying that normal packaging looks normal.
The prompt treats a version and checksum bump from the same upstream source as
ordinary, along with using a verified upstream binary to generate completions
or metadata. Changes to domains, repository owners, or download mechanisms
receive more attention.
None of these verdicts prove that a package is safe. A model can miss malicious
behavior, produce false positives, or trust an upstream binary that has itself
been compromised. The project stores the evidence behind each result so it can
be inspected; it reviews packaging and supply-chain signals, not the contents
of downloaded binaries.
The immediate work is running the checker continuously, adding notifications
for dangerous results, and improving the prompt where the models are
consistently too noisy or too trusting.
The goal is not to fully trust AI with deciding whether a package is safe. It
is to use it as an early-warning layer: when an update is clearly malicious, I
want to know quickly, see the exact source behind the verdict, and be able to
address it before more users are affected.
Hello! My name is Cretezy and I am a software developer. I write about programming, personal
projects, and more.
