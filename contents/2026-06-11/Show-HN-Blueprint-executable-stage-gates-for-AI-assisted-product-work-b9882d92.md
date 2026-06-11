---
source: "https://blueprint.ninochavez.co/"
hn_url: "https://news.ycombinator.com/item?id=48489752"
title: "Show HN: Blueprint – executable stage gates for AI-assisted product work"
article_title: "Blueprint"
author: "chavezabelino"
captured_at: "2026-06-11T13:28:00Z"
capture_tool: "hn-digest"
hn_id: 48489752
score: 1
comments: 0
posted_at: "2026-06-11T13:00:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Blueprint – executable stage gates for AI-assisted product work

- HN: [48489752](https://news.ycombinator.com/item?id=48489752)
- Source: [blueprint.ninochavez.co](https://blueprint.ninochavez.co/)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T13:00:03Z

## Translation

タイトル: HN の表示: ブループリント – AI 支援による製品作業のための実行可能なステージ ゲート
記事のタイトル: ブループリント
説明: AI 支援プロジェクトは急速に進み、すぐに腐ってしまいます。ブループリントはチェックポイントと紙の証跡を追加します。エージェントは調査、プロトタイプ、戦略文書、事実確認を実行します。あなたは耐えられる仕事を出荷します。

記事本文:
ブループリント デモ コマンドを学習して戦略に貢献 npm ↗ デモ コマンドを学習して戦略に貢献
エージェント支援による製品イニシアチブ · MIT · npm で
製品イニシアチブを実行し、
端から端まで。
AI 支援プロジェクトは急速に進み、すぐに腐ってしまいます。ブループリントは作業方法です
AI エージェントを使用して製品イニシアチブを実行するための調査 -
展開。エージェントが構築を行います。メソッドはチェックポイントを提供します
そして紙の足跡。どのような船が研究され、試作され、事実確認され、
文書化されています。
ストリップだけ。以下のすべてがさらに深くなります。
面倒な作業を AI に任せて製品作業を行う方法。
AI は文書を調査し、プロトタイプを作成し、草案を作成します。ブループリントは
作業を検証可能に保つ一連のチェックポイント。
AI によって構築された仕事は迅速で説得力がありますが、それが問題なのです。
誰も検証せず、誰も決定を書き留めなかったと自信を持って主張します。
ブループリントは領収書を強制します: すべての請求が追跡され、すべての決定が追跡されます
記録され、次のステージが始まる前にすべてのステージがチェックされます。
製品を出荷する人は誰でも、単独またはチームで AI エージェントと協力します。
ソロ、それは品質管理です。速いということはずさんであるという意味ではありません。チームでは、
それはさらに難しい問題です。AI を維持するための独自の方法は人それぞれです。
正直で、どれも並びません。ブループリントは全体の一方向です
チームで共有できます。
その背後にある考え方 —
私が求めていたキッチン ↗
各ステージでは、次のステージに影響を与えるアーティファクトが生成されます。洞察:
プロトタイプとドキュメントは同時に構築されます —
プロトタイプは決定をテストし、ドキュメントは理論的根拠を捉えます。
01 競合分析、コードベース、比較対象を調査する
02 設計の範囲の定義 - プロトタイプが何を行うか、何を行わないか
03 実際の製品に合わせたHTMLページのプロトタイプ作成
04 情報源と照らし合わせて検証されたすべての主張の事実確認
05 戦略、実現可能性、統合を文書化する

n 件の計画
06 1 つのポータルを展開する — ドキュメント + プロトタイプ + チャット
07 利害関係者のフィードバックを繰り返し、優先順位を付けます
このサイトはブループリント出力です。
ここのすべてのページは単一の blueprint.yml から生成されました
実行するのと同じパイプラインによって。 CLI は npm 上で稼働しています。ポータルは 1 つの構成からデプロイされます。
ブループリントはそれ自体、つまり方法論上で実行され、方法論の製品化に適用されます。
5つのステップ。選択した内容は以下に記載されています。すべてのコマンドは出力されたとおりに実行されます。
ノード 18+ およびクロード コード。それだけです。立ち上がるサーバーもホストするサーバーもありません。
構成に対して検証された、所有するポータルを生成します。 --name が唯一の必須フラグです。その他はすべてデフォルトでエコーバックされます。 2 つの出発点:
新しいアイデア — まだ何も存在しません。これがグリーンフィールド (デフォルト) の意味のすべてです。
既存のアプリ — そのリポジトリから実行します。それがブラウンフィールドです。
blueprint.yml で重要な 3 つのことを設定します。残りは適切なデフォルト値を持っています。
バリアント: ブラウンフィールド # グリーンフィールド |中流 |ブラウンフィールド
層: 1 # 成果物の洗練度 0–2
portal_pattern: A # A = プラットフォーム ポータル · B = 再設計レビュー
クロード・コードでは、ステージはスキルです。それらを順番に実行します。各ゲートはレビュー担当者によって強制されます。
プッシュするとポータルが展開されます。適合性を確認する — 医師はチェックしなかったことについて正直に話します。
すべてを段階ごとに、すべてのレビュー担当者が実際の取り組みで実行したいですか? →
最初の取り組みを実行し、
または、 Learn から始めてください。
7 つのコマンド、すべて本物。 npx @nino-chavez-labs/blueprint-cli <コマンド> — ゼロインストール。
新しいイニシアチブとそのポータルを足場にします。バリアント、階層、パターンの選択は、何かが書き込まれる前に検証されます。
ステージゲートを実行する: ステージをチェックする自動レビューアが実際に完了します。 --list には組み込み + 独自のものが表示されます。
ステージごとの作業量/モデルの構成とテ

lemetry — ダイヤルで支出を増減します。
ブループリントを実行しているプロジェクトごとに 1 つのレポートが作成されます。最新のものと遅れているもの。
プロジェクトを最新の方法論バージョンに移動します。最初にプレビューし、デフォルトで予行演習を行います。
ヘルスチェック: 構成、すべてのレビュアー、ポータルをロードし、チェックされなかった内容を表示します。
エージェントのチームが調整する共有ワークスペースを設定します。ハイブのセットアップでは、最初にその計画が表示されます。 --execute を渡すまでは何もプロビジョニングされません。
最新の状態を保ちます。上流にフィードバックします。
アップデートはあなたを傷つけることなく流れてきます。修正とギャップが逆流します。
ブループリントのアップグレードは、プロジェクトが固定されている方法論のバージョンを読み取り、変更内容を表示して更新します。デフォルトでは、中断はなく、予行演習が行われます。ブループリント フリートは、どの取り組みが遅れているかを示します。
ギャップ、バグ、または機能のニーズに遭遇しましたか?修正または RFC を提出します。トリアージは自動的にルーティングします。つまり、大幅な変更→ RFC が最初、バグ修正→ そのまま PR になります。
学ぶ — 最初の取り組みを実行する
はじめに、完全なエンドツーエンドのチュートリアル、およびロール ルーティング パス。
35 秒のシズルまたはステップスルー ウォークスルー — 実際のトランスクリプト、実際のアーティファクト。
最初の原則、7 つの段階、バリアントと層の選択方法。
10 のアーキテクチャ決定記録 - すべてのプリミティブの背後にある理由。
発見、試行、構築、運用、検査、ロードマップ — より詳細なリファレンス。
懐疑論者が最初に尋ねる質問には、率直に答えられます。
方法論、CLI、テンプレート、レビュー担当者セット。
ブループリント · AI エージェントを使用して製品イニシアチブをエンドツーエンドで実行するための方法論 + CLI。
github ↗ · リリース ↗ · 質問とバグ ↗

## Original Extract

AI-assisted projects move fast — and rot fast. Blueprint adds the checkpoints and the paper trail: an agent runs research, prototype, strategy docs, and fact-check; you ship work that holds up.

Blueprint Learn Demo Commands Contribute Strategy npm ↗ Learn Demo Commands Contribute Strategy
Agent-assisted product initiatives · MIT · on npm
Run a product initiative,
end to end.
AI-assisted projects move fast — and rot fast. Blueprint is a working method
for running a product initiative with an AI agent — research through
deployment. The agent does the building; the method supplies the checkpoints
and the paper trail. What ships is researched, prototyped, fact-checked, and
documented.
Just the strip. Everything below goes deeper.
A method for doing product work with an AI carrying the heavy lifting.
The AI researches, prototypes, and drafts the documents; Blueprint is
the set of checkpoints that keeps the work verifiable.
AI-built work is fast and convincing — and that's the problem.
Confident claims nobody verified, decisions nobody wrote down.
Blueprint forces the receipts: every claim traced, every decision
recorded, every stage checked before the next one starts.
Anyone shipping product work with an AI agent — solo or in a team.
Solo, it's quality control: fast doesn't mean sloppy. In a team,
it's the harder problem: everyone has their own way of keeping AI
honest, and none of them line up. Blueprint is one way the whole
team can share.
the thinking behind it —
The kitchen I'm looking for ↗
Each stage produces artifacts that feed the next. The insight:
prototype and documents are built simultaneously —
the prototype tests the decision, the docs capture the rationale.
01 Research competitive analysis, codebase, comparables
02 Design define scope — what the prototype will and won't do
03 Prototype HTML pages matching the real product
04 Fact-check every claim validated against source
05 Document strategy, feasibility, integration plans
06 Deploy one portal — docs + prototype + chat
07 Iterate stakeholder feedback, triaged
This site is Blueprint output.
Every page here was generated from a single blueprint.yml
by the same pipeline you’d run. The CLI is live on npm; the portal deploys from one config.
Blueprint runs on itself — the methodology, applied to productizing the methodology.
Five steps. The choices you make are written out below — every command runs as printed.
Node 18+ and Claude Code . That’s it — there’s no server to stand up and nothing to host.
Generates a portal you own, validated against your config. --name is the only required flag — everything else defaults and is echoed back. Two starting points:
New idea — nothing exists yet. That’s all greenfield (the default) means.
Existing app — run it from that repo. That’s brownfield .
Set the three things that matter in blueprint.yml — the rest has sane defaults.
variant: brownfield # greenfield | midstream | brownfield
tier: 1 # deliverable sophistication 0–2
portal_pattern: A # A = platform portal · B = redesign review
In Claude Code, the stages are skills. Run them in order — each gate is enforced by a reviewer.
Push and your portal deploys. Verify conformance — doctor is honest about what it didn’t check.
Want the whole thing — stage by stage, every reviewer gate, on a real initiative? →
Run your first initiative ,
or start at Learn .
7 commands, all real. npx @nino-chavez-labs/blueprint-cli <command> — zero install.
Scaffold a new initiative and its portal — your variant, tier, and pattern choices validated before anything is written.
Run a stage gate: an automated reviewer that checks a stage is actually done. --list shows built-in + your own.
Per-stage effort/model configuration and telemetry — dial spend up or down.
One report across every project running Blueprint: which are current, which have fallen behind.
Move a project to the latest methodology version — preview first, dry-run by default.
Health check: loads the config, every reviewer, and the portal — and says what it didn’t check.
Set up the shared workspace a team of agents coordinates through. hive setup shows its plan first — nothing is provisioned until you pass --execute .
Stay current. Feed back upstream.
Updates flow down without breaking you; your fixes and gaps flow back up.
blueprint upgrade reads the methodology version your project is pinned to, shows what changed, and updates it — non-breaking, dry-run by default. blueprint fleet shows which of your initiatives have fallen behind.
Hit a gap, a bug, or a feature need? File an amendment or RFC. Triage routes it automatically — substantial change → RFC first, bug-fix → straight PR.
Learn — run your first initiative
Getting started, the full end-to-end tutorial, and role-routed paths.
35-second sizzle or the step-through walkthrough — real transcripts, real artifacts.
The first principles, the seven stages, how variant and tier get picked.
10 architecture decision records — the why behind every primitive.
Discover · Try · Build · Operate · Inspect · Roadmap — the deeper reference.
The questions a skeptic asks first, answered straight.
The methodology, the CLI, the template, the reviewer set.
Blueprint · A methodology + CLI for running product initiatives end-to-end with an AI agent.
github ↗ · releases ↗ · questions & bugs ↗
