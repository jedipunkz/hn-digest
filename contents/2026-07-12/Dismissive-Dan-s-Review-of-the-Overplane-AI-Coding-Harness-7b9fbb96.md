---
source: "https://www.overplane.dev/overview/dan/"
hn_url: "https://news.ycombinator.com/item?id=48877337"
title: "Dismissive Dan's Review of the Overplane AI Coding Harness"
article_title: "Overplane Overview | Dismissive Dan"
author: "mayank"
captured_at: "2026-07-12T01:47:44Z"
capture_tool: "hn-digest"
hn_id: 48877337
score: 3
comments: 0
posted_at: "2026-07-12T01:02:09Z"
tags:
  - hacker-news
  - translated
---

# Dismissive Dan's Review of the Overplane AI Coding Harness

- HN: [48877337](https://news.ycombinator.com/item?id=48877337)
- Source: [www.overplane.dev](https://www.overplane.dev/overview/dan/)
- Score: 3
- Comments: 0
- Posted: 2026-07-12T01:02:09Z

## Translation

タイトル: Overplane AI コーディング ハーネスに関する Dismissive Dan のレビュー
記事のタイトル: オーバープレーンの概要 |否定的なダン
説明: 好みの読み方で、Overplane を平易に説明します。

記事本文:
コンテンツにスキップ
オーバープレーンの概要サンプル ガイド リファレンス ダウンロード v0.0.7 概要
Overplane は、フォルダーを変換する、小さなオープンソースのシングルバイナリ ツールです。
プレーンな Markdown 仕様を実際に動作するソフトウェアに変換: クロスチェックします
SMT ソルバーを使用して仕様を決定し、ロックダウンされた内部で AI コーディング エージェントを駆動します
ローカルコンテナ。
Overplane の概要については、スタイルを選択してください。
ユーザー マニュアル フィールド ブリーフィング AI 基調講演 否定的なダン 不機嫌な脳 dan_ | 47点 | 63件のコメント
まず、チームの出荷に心からおめでとうございます。
常に尊敬に値します。そうは言っても、私が言いたいのはこれです
建設的には、製品を見るのに苦労しています。シングルです
番号付き Markdown ファイルを読み取り、ローカル ファイルを構築する Go バイナリ
Docker または Podman イメージ、Claude Code、Codex、Gemini CLI、または
OpenCode は独自の API キーでヘッドレス内部にエージェントを持ちます
いくつかの SMT ファイルを発行し、Z3 がそれをチェックして、
モデルはコードディレクトリに書き込まれます。素敵ですね、
しかし、開発者なら誰でも、これを簡単に実現できます。
Dockerfile、読み取り専用バインド マウント、短いシェル ループ、および cron
再構築のための仕事。私のものは約 40 行の bash と z3 です。
呼び出し。これを読んでいるほとんどの人が何かを持っていると思います
似たような。
「確認済み」については温かく言った
パイプライン (raise → verify → codegen) がエンドツーエンドで実行されることを喜んで許可します。
エージェントが各仕様を IR および SMT-LIB に引き上げ、Z3 がそれぞれをチェックします
spec のモデルとそれらすべてを結合したモデル、そして
矛盾があるとビルドに失敗します。それはもちろん魅力的ですが、
LLM がアサートを発行し、それらを Z3 にパイプするのは次のとおりです。
私たちの多くが Makefile に入れてきたのと同じようなものです。
SMT-LIB 2規格が定着しました。耐荷重観察、
穏やかに提供されます。パイプラインは両端でヒューリスティックです。ソルバー
忠実にチェックする

王様、あなたの散文の間違った形式化は、
自信を感じるための非常に効果的な方法であり、チームの真の気持ちを高めることができます。
彼らはこのページを含め、あらゆる場所でこれを開示していると信じています。のために
ほとんどの読者がそうしていると思いますが、私たち人間は、小さな絵をスケッチします。
ステートフル以前の TLA+ モデル、以下の違い
一貫性のチェックと証明は第二の性質です。実際に検証済み
ソフトウェアは seL4 または CompCert で、費用は数十人です。それはです
彼らに配慮して、新しい人のためにそれを詳しく説明します。
エージェントは、リポジトリが読み取り専用でマウントされたコンテナ内で実行されます。
書き込みは出力マウントに限定されます。思いやりのあるタッチ —
確かに誰もがすでにエージェントを使い捨てで実行していますが、
読み取り専用マウントの VM。本当に驚かれると思います
それ以外のことを学びます。
正規化されたトークンとコストを備えた 4 つのエージェント CLI 上に 1 つのドライバー
会計。まだ同じものを書いていない場合は便利です
自分自身をシムする。ほとんどのチームがアイドル状態でそれを行うと思います
午後。
コンテンツでハッシュされたコンテナ イメージ、コンテンツでアドレス指定された出力
ファイルセット。非常に上品です — 当然のことながら、この多くは次のようなものから得られます。
Nix は、今ではほとんどのショップで使用されていると思います。
Apache-2.0 (アカウントなし) はローカルで実行されます。賭け金が記載されていますが、
丁寧に。
敬意を表しますが、おそらくそうではありません。おそらくここにいる皆さんと同じように、
すでに独自のサンドボックス スクリプトを保守している場合、請求
ノーマライザー、再現可能なコンテナ パイプライン、そして少し
ソルバー実行のための仕様から SMT へのハーネス。インタラクティブな作業の場合、
クロード・コードとカーソルは相変わらず素敵です。証拠が必要な場合は、
今日のプログラムの正確さ、TLA+、Dafny、または Lean は快適です
週末の読書 — ここに出荷されるものは仕様の一貫性をチェックするものであり、チェックするものではありません
コード。 GitHub Spec Kit と AWS Kiro は同じスペック主導で動作します
コンテナ隔離のない通路。温かく認めます、
それは稀なチームにとっては

まだすべてを構築するまでには至っていない
これ自体、ここでのパッケージ化 - 無人、再現可能、
ソルバー ゲートとエージェントの移植性を備えたサンドボックス化された仕様からコードへの変換 —
本当に賢明です。
このツールは意図的に退屈なものです。影響力はあなたのスペックにあります —
粒度、精度、モデルをどの程度許容するか —
実行ごと ( --agent ) または仕様ごとにどのエージェントを固定するか (
フロントマターのagent_config)、およびサンドボックス
overplane.yaml で作成します (ベースイメージ、追加
パッケージ、エージェント、環境パススルー)。仕様書作成
ソルバーがチェックするのに十分な精度は、幸いにも私たち全員が持っているものです
自然にできるので、中央値のチームには何の困難もないと予想しています。
温かくブックマークしました。私はソルバーフェーズがこうなるだろうと予想していましたが、
私は彼らに喜んでいます。
試してみる準備はできましたか?ガイドが説明します
システムを準備し、最初のプロジェクトを構築します。を参照することを好みます。
表面積が先ですか？を参照してください。
参考にしてください。

## Original Extract

A plain-language explanation of Overplane, in a reading style of your choice.

Skip to content
Overplane Overview Examples Guide Reference Download v0.0.7 Overview
Overplane is a small, open-source, single-binary tool that turns a folder
of plain Markdown specs into working software: it cross-checks
the specs with SMT solvers, then drives AI coding agents inside locked-down
local containers.
Choose your style for an overview of Overplane:
User's Manual Field Briefing AI Keynote Dismissive Dan grug brain dan_ | 47 points | 63 comments
First, sincere congratulations to the team on shipping — that
always deserves respect. That said, and I mean this
constructively, I'm struggling to see the product. It's a single
Go binary that reads numbered Markdown files, builds a local
Docker or Podman image, runs Claude Code, Codex, Gemini CLI, or
OpenCode headlessly inside it on your own API keys, has the agent
emit some SMT files that Z3 then checks, and swaps whatever the
model wrote into a code directory. Which is lovely,
but any developer can already get this quite trivially with a
Dockerfile, a read-only bind mount, a short shell loop, and a cron
job for the rebuilds. Mine is about forty lines of bash, plus a z3
invocation. I assume most people reading this have something
similar.
On "verified", said with warmth
I will happily grant that the pipeline — raise → verify → codegen — now runs end to end:
an agent lifts each spec into IR and SMT-LIB, Z3 checks each
spec's model and a merged model of all of them, and a
contradiction fails the build. Which is charming, though of course
an LLM emitting assert s and piping them through Z3 is
the sort of thing many of us have had in a Makefile since the
SMT-LIB 2 standard settled down. The load-bearing observation,
offered gently: the pipeline is heuristic at both ends. A solver
faithfully checking the wrong formalization of your prose is a
very efficient way to feel confident, and to the team's genuine
credit they disclose this everywhere, including on this page. For
those of us who — as I imagine most readers do — sketch a small
TLA+ model before anything stateful, the distinction between
consistency checking and proof is second nature. Actually verified
software is seL4 or CompCert and costs person-decades; it is
considerate of them to spell that out for the newer folks.
The agent runs in a container with your repo mounted read-only
and writes confined to an output mount. A thoughtful touch —
though surely everyone already runs their agents in a throwaway
VM with a read-only mount; I would be genuinely surprised to
learn otherwise.
One driver over four agent CLIs with normalized token and cost
accounting. Convenient, if you haven't already written the same
shim yourself, which I'd have assumed most teams did in an idle
afternoon.
Content-hashed container images, content-addressed output
filesets. Very tasteful — naturally you'd get much of this from
Nix, which I understand most shops use by now.
Apache-2.0, no account, runs locally. Table stakes, but stated
politely.
Respectfully, probably not — if, like presumably everyone here,
you already maintain your own sandbox scripts, billing
normalizers, a reproducible container pipeline, and a little
spec-to-SMT harness for the solver runs. For interactive work,
Claude Code and Cursor remain lovely. If you need proofs of
program correctness today, TLA+, Dafny, or Lean are a pleasant
weekend of reading — what ships here checks spec consistency, not
the code. GitHub Spec Kit and AWS Kiro sit in the same spec-driven
aisle without the container isolation. I will concede, warmly,
that for the rare team that hasn't gotten around to building all
of this themselves, the packaging here — unattended, reproducible,
sandboxed spec-to-code with a solver gate and agent portability —
is genuinely sensible.
The tool is deliberately boring; the leverage is in your specs —
granularity, precision, how much latitude you leave the model — in
which agent you pin per run ( --agent ) or per spec (
agent_config in the frontmatter), and in the sandbox
you compose in overplane.yaml (base image, extra
packages, which agents, env passthrough). Writing specifications
precise enough for a solver to check is, happily, something we all
do naturally, so I anticipate no difficulties for the median team.
Bookmarked, with warmth; I did predict the solver phases would
ship, and I am delighted for them.
Ready to try it? The guide walks you through
preparing your system and building your first project. Prefer to browse the
surface area first? See the
reference .
