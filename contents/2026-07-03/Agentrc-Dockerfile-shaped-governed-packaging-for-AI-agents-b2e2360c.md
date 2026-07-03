---
source: "https://agentrc.ai/"
hn_url: "https://news.ycombinator.com/item?id=48771202"
title: "Agentrc – Dockerfile-shaped, governed packaging for AI agents"
article_title: "agentrc · agentrc"
author: "adeelahmadch"
captured_at: "2026-07-03T05:52:33Z"
capture_tool: "hn-digest"
hn_id: 48771202
score: 1
comments: 0
posted_at: "2026-07-03T05:36:32Z"
tags:
  - hacker-news
  - translated
---

# Agentrc – Dockerfile-shaped, governed packaging for AI agents

- HN: [48771202](https://news.ycombinator.com/item?id=48771202)
- Source: [agentrc.ai](https://agentrc.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T05:36:32Z

## Translation

タイトル: Agentrc – AI エージェント用の Dockerfile 形式の管理されたパッケージ化
記事のタイトル: エージェントrc · エージェントrc
説明: Agent Run Config: 1 つの AI エージェントを移植可能な管理されたアーティファクトとしてパッケージ化するためのオープン仕様。

記事本文:
作業草案 — Agentrc 0.1.0-draft.6 は進化中の仕様草案であり、完成した標準ではありません。破壊的な変更を期待してください。変更履歴 →
エージェントrc_
エージェントアークとは何ですか?
スペック
ドキュメント
例
プロフィール
CLI
GitHub
★スター
オープン仕様
ポータブルで管理された AI エージェント _
bashrc や zshrc と似ていますが、エージェント用です。 Agentfile は、1 つの AI エージェントの ID、機能、システム プロンプト、ツールに加えて、モデル、リソース、ネットワークに対するリクエストを、セキュリティ チームが確認できる型付きポリシーとして宣言します。 OCI アーティファクトとしてパッケージ化します。互換性のあるランナーがそれを実行し、強制します。ランタイム、クラウド、モデルプロバイダー、エージェントフレームワークではありません。
01 > { } _ /mnt
=> ポリシー 10
アークビルド..
0x1f 4a2b >_
アイデンティティ::
{ エージェント } //
能力1
SOP -> 実行
ラベル0.1.0
$ アークリント _
ネットワーク:443
許可|狭い
oci://ghcr..
04a2 > 1101
拒否デフォルト
>_ 杉は大丈夫
Agentrc:~$cat Agentfile
# syntax=agentrc.agentfile/v0.1
Python から:3.11-slim
IDENTITY 名前=hello バージョン=0.1 著者=acme
IDENTITY description= "最小限のagentrcエージェント"
能力のテキスト
SOP あなたは最小限のエージェントの例です。読む
求められたらファイルします。他には何もしません。
CMD Python ./agent.py
# ツール (ローカル、組み込み) → /mnt/tools/
COPY --chmod=755 ./tools/file_read /mnt/tools/file_read
# リクエスト: プラットフォームは許可、絞り込み、または拒否します。
ポリシー モデル名 claude-sonnet-4
ポリシー Agent.tool_timeout 30 秒
ポリシーネットワーク dns:api.example.com:443
HEALTHCHECK --interval=60s CMD /mnt/tools/file_read --agentrc-schema
arc lint: OCI にコンパイルされます + ai.agentrc.* ラベル ポリシーはレビュー可能です
宣言的で再現可能
1 つの Agentfile は、ID、機能、ポリシー、ツール、およびリソースをキャプチャし、標準の Dockerfile キーワードに加えて、エージェント ネイティブの 4 つのキーワード ( IDENTITY 、 CAPABILITY 、 SOP 、 POLICY ) を再利用します。
POLICY 行は、モデル、リソース、または制約を要求します。プラットフォーム

それを許可、絞り込む、または拒否し、Cedar による決定をデフォルトで拒否します。
ビルドは、インテントを名前空間付きの ai.agentrc.* OCI ラベルに変換します。プラットフォームはエージェントファイルではなくラベルを読み取るため、エージェントは他のコンテナイメージと同様に出荷、署名、ミラーリングを行います。
1 つのバイナリ — Agentrc (エイリアス arc )。 Agentfile のスキャフォールディング、検証、構築を行い、エージェントの要求内容を検査し、アーティファクトをバックエンドのデプロイ構成に変換します。
カール -fsSL https://agentrc.ai/install.sh |しー
自作
醸造インストール
アディーラマッド/タップ/エージェントrc
ゴー 1.25+
インストールに行く
github.com/adeelahmad/agentrc/cmd/agentrc@latest
ソースから
git clone https://github.com/adeelahmad/agentrc
cd Agentrc && go build -o arc ./cmd/agentrc
macOS および Linux (amd64 / arm64) 用の事前構築済み、チェックサム検証済みバイナリ。アークバージョンで確認してください。最初に読んだほうがいいですか? curl -fsSL https://agentrc.ai/install.sh を実行して調べます。
エージェントをスキャフォールディング、検証して移植可能な OCI アーティファクトにコンパイルし、ローカル ランナーが実行するものを正確にプレビューします。
Scaffold arc init › 書き込み ./Agentfile
arc lint Agentfile の検証 › ID、ポリシー、スキーマ
アーク build -t ghcr.io/you/hello:0.1 をビルドします。 › OCI アーティファクト
実行アークのプレビュー run ghcr.io/you/hello:0.1 --backend local --dry-run
arc build は、実際の OCI イメージを生成します (docker build およびagentrc BuildKit フロントエンド経由)。 --dry-run は、ランナーが適用する構成を出力します。agentrc は宣言して変換します。独自のランタイムは同梱されていません。
同じアーティファクトをクラウドに送信する
ビルドは ai.agentrc.* ラベルを 1 回書き込みます。ポイント アークは任意のバックエンドで実行され、これらのラベルをそのプラットフォームのデプロイ フォームに変換します。
アークプッシュ
ghcr.io/you/hello:0.1
→ 任意の OCI レジストリ
アークラン …こんにちは:0.1
--バックエンドの基盤 --ドライラン
→ CreateAgentRuntime JSON
アークラン …こんにちは:0.1
--バックエンド kubernetes --dry-ru

n
→ マニフェストをデプロイする
分離エージェントが作成する
プロジェクトは標準スタイルのリポジトリとして公開されます。最初に仕様、次にリファレンス ツールです。

## Original Extract

Agent Run Config: an open specification for packaging one AI agent as a portable, governed artifact.

Working Draft — agentrc 0.1.0-draft.6 is an evolving specification draft, not a finished standard. Expect breaking changes. Changelog →
agentrc _
What is agentrc?
Spec
Docs
Examples
Profiles
CLI
GitHub
★ Star
Open specification
portable, governed ai agents _
Like bashrc or zshrc , but for an agent. An Agentfile declares one AI agent's identity, capabilities, system prompt, and tools, plus its requests for models, resources, and network — as typed policy a security team can review. Package it as an OCI artifact; compatible runners execute and enforce it. Not a runtime, cloud, model provider, or agent framework.
01 > { } _ /mnt
=> POLICY 10
arc build ..
0x1f 4a2b >_
IDENTITY ::
{ agent } //
CAPABILITY 1
SOP -> run
label 0.1.0
$ arc lint _
network:443
grant|narrow
oci://ghcr..
04a2 > 1101
deny-default
>_ cedar ok
agentrc:~$ cat Agentfile
# syntax=agentrc.agentfile/v0.1
FROM python:3.11-slim
IDENTITY name=hello version=0.1 author=acme
IDENTITY description= "Minimal agentrc agent"
CAPABILITY text
SOP You are a minimal example agent. Read a
file when asked; do nothing else.
CMD python ./agent.py
# Tool (local, embedded) → /mnt/tools/
COPY --chmod=755 ./tools/file_read /mnt/tools/file_read
# Requests: the platform grants, narrows, or rejects
POLICY model.name claude-sonnet-4
POLICY agent.tool_timeout 30s
POLICY network dns:api.example.com:443
HEALTHCHECK --interval=60s CMD /mnt/tools/file_read --agentrc-schema
arc lint: ok compiles to OCI + ai.agentrc.* labels policy reviewable
Declarative & reproducible
One Agentfile captures identity, capability, policy, tools, and resources — reusing standard Dockerfile keywords plus four agent-native ones: IDENTITY , CAPABILITY , SOP , POLICY .
A POLICY line requests a model, resource, or constraint. The platform grants, narrows, or rejects it and enforces the decision with Cedar, deny-by-default.
The build translates intent into namespaced ai.agentrc.* OCI labels. Platforms read the labels — never the Agentfile — so agents ship, sign, and mirror like any container image.
One binary — agentrc (alias arc ). It scaffolds, validates, and builds Agentfiles, inspects what an agent requests, and translates an artifact into a backend's deploy config.
curl -fsSL https://agentrc.ai/install.sh | sh
Homebrew
brew install
adeelahmad/tap/agentrc
Go 1.25+
go install
github.com/adeelahmad/agentrc/cmd/agentrc@latest
From source
git clone https://github.com/adeelahmad/agentrc
cd agentrc && go build -o arc ./cmd/agentrc
Prebuilt, checksum-verified binaries for macOS & Linux (amd64 / arm64). Confirm with arc version . Prefer to read first? curl -fsSL https://agentrc.ai/install.sh and inspect it.
Scaffold, validate, and compile an agent into a portable OCI artifact, then preview exactly what a local runner would execute.
Scaffold arc init › writes ./Agentfile
Validate arc lint Agentfile › identity, policy & schema
Build arc build -t ghcr.io/you/hello:0.1 . › OCI artifact
Preview the run arc run ghcr.io/you/hello:0.1 --backend local --dry-run
arc build produces a real OCI image (via docker build and the agentrc BuildKit frontend). --dry-run prints the config a runner would apply — agentrc declares and translates; it ships no runtime of its own.
Ship the same artifact to the cloud
The build writes ai.agentrc.* labels once. Point arc run at any backend to translate those labels into that platform's deploy form.
arc push
ghcr.io/you/hello:0.1
→ any OCI registry
arc run …hello:0.1
--backend bedrock --dry-run
→ CreateAgentRuntime JSON
arc run …hello:0.1
--backend kubernetes --dry-run
→ deploy manifests
The separation agentrc creates
The project is published as a standards-style repository: specification first, reference tooling second.
