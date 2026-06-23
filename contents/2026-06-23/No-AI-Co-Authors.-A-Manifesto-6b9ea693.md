---
source: "https://no-ai-coauthors.dev"
hn_url: "https://news.ycombinator.com/item?id=48651494"
title: "No AI Co-Authors. A Manifesto"
article_title: "no-ai-coauthors — the byline is not shared"
author: "zethraeus"
captured_at: "2026-06-23T21:32:08Z"
capture_tool: "hn-digest"
hn_id: 48651494
score: 7
comments: 3
posted_at: "2026-06-23T21:09:25Z"
tags:
  - hacker-news
  - translated
---

# No AI Co-Authors. A Manifesto

- HN: [48651494](https://news.ycombinator.com/item?id=48651494)
- Source: [no-ai-coauthors.dev](https://no-ai-coauthors.dev)
- Score: 7
- Comments: 3
- Posted: 2026-06-23T21:09:25Z

## Translation

タイトル: AI の共著者はいません。マニフェスト
記事のタイトル: no-ai-coauthors — 署名欄は共有されません
説明: AI アトリビューション トレーラーを拒否する commit-msg フックと GitHub アクション。これまで同様、著者が寄稿者です。

記事本文:
no-ai-共著者
マニフェスト
何が引っかかるのか
インストール
GitHub ↗
commit-msgフック・githubアクション
commit-msg フックと、コミットのキャリーを拒否する GitHub アクション
AI アトリビューション トレーラー — 署名欄が唯一の場所であるため
責任を共有することはできません。
署名欄は共有リソースではありません。
AI の共同作成者の git-commit 署名欄は、印象的なマーケティングのフックです —
しかし、それらは出所を示す有用なシグナルを提供しません。
2026年半ばです。書かれたすべてのコードは、高度なツールのホストによって容易に作成されます。
これらの最新のものは AI の「共著者」です。これらは信じられないほどの有用性を持っています。
ただし、開発ツールチェーンの他のツールと同様に、
彼らは自分たちの成果に対して意味のある責任を負っていません。
せいぜい、コードを生成したボットを「共同作成者」としてタグ付けするだけで、コードにノイズが追加されます。
開発における単一の最も重要な説明責任と帰属メカニズム
階層。最悪の場合、この慣行は同様に責任を分散させることを意味します。
コードの作成という行為が変化するのと同じくらい、作成に関する社会的なルールも変化しています
私たちは自分の貢献に対して責任があります。
これまでと同様、「著者」は貢献者です。
コミットメッセージから企業スパムを削除します。
工具に当たります。それは人を維持します。
表示名と電子メールの両方で照合が実行され、正規化されるため、ボットの名前が変更されます
または、数値の応答アドレスはすり抜けません。人間というのは、
爆発範囲内には決して入ってはいけません。
# コミットは作成されません
# コミットはそのまま通過します
Claude という名前のボットをブロックします。
クロード・シャノンという人間を迎え入れる。
ピン 1.0.0 。フックマネージャーを選択してください。
すべてのパスはパブリック リポジトリとタグ付けされたリリースを指しているため、インストールは繰り返し可能です。アップグレードするときにタグを交換します。
リポジトリを .pre-commit-config.yaml に追加し、commit-msg ステージをインストールします。
デフォルト_i

nstall_hook_types : [コミット前、コミットメッセージ]
リポジトリ:
- リポジトリ: https://github.com/GoodHatsLLC/no-ai-coauthors
リビジョン : 1.0.0
フック:
- id : no-ai-coauthors
端末コピー
$ pre-commit install --hook-type commit-msg
prek は同じ構成を読み取るか、ネイティブの prek.toml エントリを使用します。
default_install_hook_types = [ "pre-commit" , "commit-msg" ]
[[ リポジトリ ]]
リポジトリ = "https://github.com/GoodHatsLLC/no-ai-coauthors"
rev = "1.0.0"
フック = [{ id = "no-ai-coauthors" }]
端末コピー
$ prek install --hook-type commit-msg
これを Lefthook リモート構成として使用し、インストールします。
リモコン：
- git_url : https://github.com/GoodHatsLLC/no-ai-coauthors
参照: 1.0.0
構成:
- 左フック.yml
端末コピー
$ 左フックのインストール
リポジトリを複合アクションとして直接使用します。プッシュ時にイベント ペイロードを読み取ります。 pull_request では PR コミットを読み取ります。
名前 : no-ai-共著者
に:
プルリクエスト:
プッシュ：
ジョブ:
no-ai-共著者:
実行: ubuntu-最新
権限:
内容：読む
プルリクエスト : 読み取り
手順:
- 使用: GoodHatsLLC/no-ai-coauthors@1.0.0
フック マネージャーはまったくありません。スクリプトをダウンロードして、core.hooksPath を指定します。
$ mkdir -p .githooks
$カール-fsSL\
https://raw.githubusercontent.com/GoodHatsLLC/no-ai-coauthors/1.0.0/hooks/no-ai-coauthors \
-o .githooks/commit-msg
$ chmod +x .githooks/commit-msg
$ git config core.hooksPath .githooks
このパッケージは、ノード指向マネージャー向けに AI のない共同作成者ビンを公開します。
no-ai-共著者「$1」
package.json · simple-git-hooks のコピー
{
"simple-git-hooks" : {
"commit-msg" : "no-ai-coauthors \"$1\""
}
}
あなたの名前がコミットにあります。それは意味します。

## Original Extract

A commit-msg hook and GitHub Action that rejects AI attribution trailers. The author is, as ever, the contributing person.

no-ai-coauthors
Manifesto
What it catches
Install
GitHub ↗
commit-msg hook · github action
A commit-msg hook and GitHub Action that rejects commits carrying
AI attribution trailers — because the byline is the one place
accountability can’t be shared.
The byline is not a shared resource.
AI co-author git-commit bylines are an impressive marketing hook —
but they provide no useful signal of provenance.
It is mid-2026. All code written is facilitated by a host of sophisticated tools.
The most recent of these are AI ‘co-authors.’ These have incredible utility,
but — as with the other tools in the development toolchain —
they are not meaningfully accountable for their output.
At best, tagging the bot that generated code as a ‘co-author’ adds noise to the
single most important accountability and attribution mechanism in the development
hierarchy. At worst, the practice implies similarly diffused responsibility.
As much as the act of code authorship is changing, the social rules for authorship
must not slip: we are responsible for our contributions.
The ‘author’ is, as ever, the contributing person.
Remove corporate spam from commit messages.
It strikes the tool. It keeps the person.
Matching runs on both the display name and the email, normalized — so a renamed bot
or a numeric noreply address doesn’t slip through. A human is
never in the blast radius.
# the commit will not be created
# the commit goes through, untouched
It blocks the bot named Claude .
It welcomes the human named Claude Shannon .
Pin 1.0.0 . Pick your hook manager.
Every path points at the public repo and a tagged release, so installs are repeatable. Swap the tag when you upgrade.
Add the repo to .pre-commit-config.yaml , then install the commit-msg stage.
default_install_hook_types : [pre-commit, commit-msg]
repos :
- repo : https://github.com/GoodHatsLLC/no-ai-coauthors
rev : 1.0.0
hooks :
- id : no-ai-coauthors
terminal copy
$ pre-commit install --hook-type commit-msg
prek reads the same config, or use a native prek.toml entry.
default_install_hook_types = [ "pre-commit" , "commit-msg" ]
[[ repos ]]
repo = "https://github.com/GoodHatsLLC/no-ai-coauthors"
rev = "1.0.0"
hooks = [{ id = "no-ai-coauthors" }]
terminal copy
$ prek install --hook-type commit-msg
Consume it as a Lefthook remote config, then install.
remotes :
- git_url : https://github.com/GoodHatsLLC/no-ai-coauthors
ref : 1.0.0
configs :
- lefthook.yml
terminal copy
$ lefthook install
Use the repo directly as a composite action. On push it reads the event payload; on pull_request it reads the PR commits.
name : no-ai-coauthors
on :
pull_request :
push :
jobs :
no-ai-coauthors :
runs-on : ubuntu-latest
permissions :
contents : read
pull-requests : read
steps :
- uses : GoodHatsLLC/no-ai-coauthors@1.0.0
No hook manager at all — download the script and point core.hooksPath at it.
$ mkdir -p .githooks
$ curl -fsSL \
https://raw.githubusercontent.com/GoodHatsLLC/no-ai-coauthors/1.0.0/hooks/no-ai-coauthors \
-o .githooks/commit-msg
$ chmod +x .githooks/commit-msg
$ git config core.hooksPath .githooks
The package exposes a no-ai-coauthors bin for Node-oriented managers.
no-ai-coauthors "$1"
package.json · simple-git-hooks copy
{
"simple-git-hooks" : {
"commit-msg" : "no-ai-coauthors \"$1\""
}
}
Your name is on the commit. Mean it.
