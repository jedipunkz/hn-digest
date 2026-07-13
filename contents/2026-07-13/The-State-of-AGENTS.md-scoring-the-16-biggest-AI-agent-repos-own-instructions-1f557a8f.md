---
source: "https://fpaul.dev/writing/state-of-agents-md-2026/"
hn_url: "https://news.ycombinator.com/item?id=48892210"
title: "The State of AGENTS.md: scoring the 16 biggest AI agent repos' own instructions"
article_title: "The State of AGENTS.md: Scoring the Agent-Makers' Own Instructions — Franz Paul"
author: "Fpailo"
captured_at: "2026-07-13T14:06:28Z"
capture_tool: "hn-digest"
hn_id: 48892210
score: 1
comments: 1
posted_at: "2026-07-13T13:16:03Z"
tags:
  - hacker-news
  - translated
---

# The State of AGENTS.md: scoring the 16 biggest AI agent repos' own instructions

- HN: [48892210](https://news.ycombinator.com/item?id=48892210)
- Source: [fpaul.dev](https://fpaul.dev/writing/state-of-agents-md-2026/)
- Score: 1
- Comments: 1
- Posted: 2026-07-13T13:16:03Z

## Translation

タイトル: The State of AGENTS.md: 16 の最大の AI エージェント リポジトリ自体の命令のスコアリング
記事のタイトル: AGENTS.md の現状: エージェント作成者自身の指示のスコアリング — Franz Paul
説明: 16 の最大の AI エージェント リポジトリの AGENTS.md ファイルを決定論的エンジンでスコア付けしました。合計スター数は約 146 万、A グレードはゼロ、大会を広めたリポジトリは 16 件中 13 位にランクされています。

記事本文:
AGENTS.md の現状: エージェント作成者自身の指示のスコアリング — Franz Paul コンテンツへスキップ franz paul Projects
AGENTS.md の現状: エージェント作成者自身の指示のスコアリング
エージェント-MD 命令品質の AI-エージェント ツール シュリフ 16 リポジトリ。両者の間には約 146 万人の GitHub スターがいます。ゼロAグレード。
AGENTS.md は、密かにリポジトリレベルのエージェント命令のクロスツール標準になりました。これは、Cursor、Codex、Copilot、Claude Code、およびその他の増え続けるリストが起動時に読み取る単一ファイルです。そこで私は、AI エージェントのエコシステム自体が、それが作成した規則をどのように使用しているかを調べてみました。私は、最も有名な AI コーディング エージェントとエージェント フレームワークのリポジトリを 36 個調べ、ルート AGENTS.md を出荷する少なくとも 20,000 個のスターを持つ 16 個を見つけ、決定論的エンジンですべてのファイルをスコアリングしました。LLM ジャッジはなく、すべてのマシンで同じファイルが同じスコアになり、この投稿の最後にある正確なコミット SHA に対して再現可能です。
結果: 平均は 70.0、グレードの中央値は C、A は 1 つもありませんでした。この大会を普及させたチームのリポジトリは 16 件中 13 位にランクされました。そして、私が見つけた最高の AGENTS.md は、170 個の星が付いたライブラリに属しています。
これは、私自身のマージされた MCP 貢献を採点するために使用したスコアラーと、静的命令ファイルのセキュリティのケースの背後にあるスコアラーです。これはオープン ソース、MIT、stdlib のみです。ルーブリックは、隠されたプロンプトではなく、読むことができる辞書です。
このセットは、ルート AGENTS.md を出荷する 20,000 スター以上のすべてのよく知られた AI コーディング エージェントまたはエージェント フレームワーク プロジェクトであり、それぞれの HEAD は 2026-07-07 にリリースされた schliff==8.5.0 でスコア付けされています。 AGENTS.md ルーブリックは、構造 (重み 0.4)、運用範囲 (0.4)、および効率 (0.2) の 3 つの次元で構成されています。
B が 7 つ、C が 6 つ、D が 2 つ、E が 1 つです。A はなく、S もありません。 † zed は、AGENTS.md をその .rules ファイルへのシンボリックリンクとして出荷します。スコアは解決されたコンテのスコアです

nt であり、シンボリックリンク自体は発見です。生の AGENTS.md を取得するツールは、命令ではなく 6 文字の文字列 .rules を取得します。
1 つのリポジトリは非独立データ ポイントとして除外されました。OpenInterpreter の AGENTS.md は、openai/codex とバイト同一です。これは、プロジェクトがオープン モデルの codex 派生であるため、同じ git blob です。それを2回数えるのは不誠実だろう。
openai/codex — AGENTS.md の普及に最も貢献したチームのリポジトリ — のスコアは 66.4、16 点中 13 位の C です。ファイルは怠惰ではなく、28 の見出しで 22.5 KB あり、レビュー規則と TUI スタイルのルールが実に豊富です。しかし、運用カバレッジの次元では、このファイルはエージェントにリポジトリを操作するための装備を提供するかという狭い質問を投げかけます。 test コマンドと lint コマンドがあります (test だけ、fix だけ)。ビルド コマンドはそうではありません。 Rust と Bazel のモノリポジトリでは、ファイルのどこにも表示されません。コミットや PR の規則もありません。私はギャップを問題として提出し、修正を申し出ました。
このパターンは 16 つすべてに共通しています。構造は均一に強く、ヘディングが簡単なため、75 ～ 95 です。表の上位と下位を分けるのは、運用範囲です。つまり、実際のセットアップ、ビルド、テストのコマンドと、エージェントが実行時に安価に再検出できない問題です。
有名なフレームワーク、空の操作 #
3 つのファイルが D または E を獲得し、生の Markdown を手で読んでも、それぞれが生き残りました。これらはスコアラーのアーティファクトではありません。
crewAI (55.2、D): 1 KB の貢献決まり文句 — 「ベスト プラクティスに従う」、DRY、YAGNI。具体的なコマンドはドキュメント サイト ( mintlify dev ) のみです。実際の Python フレームワークをセットアップ、構築、テストする方法をエージェントに指示するものは何もありません。運用カバレッジ: 100 件中 15 件。
smolagents (50.0, D): ファイル全体は 157 バイトです。一般的なアドバイスの 4 つの箇条書き (「Pythonic であれ」)

」、「OOP 原則に従う」)。運用カバレッジ: 0。
Roo コード (37.0、E): 445 バイト — 本当に役立つ UI 状態の落とし穴が 1 つあり、それ以外は何もありません。セットアップ、ビルド、テストは必要ありません。運用範囲: 0。
運用カバレッジがゼロで高い構造スコアは、プロジェクトを実行する必要があるエージェントのためではなく、人間のスキムのために書かれたファイルの署名です。
サンプル内の最高のファイルには 170 個の星が付いています #
16 個の巨人のうち、A に到達するものは 1 つもありません。私のサンプル全体で最高の AGENTS.md (以下のコーパスとこれら 16 個を加えたもの) は、Postgres でサポートされている 170 個のスターを持つジョブ キューである maxcountryman/underway に属しています: 91.0、グレード A 。これは、セットアップ、ビルド、テスト、lint、およびプロジェクト構造をコピー＆ペースト可能なコマンドとしてカバーする 1,728 個のトークンです。
block/goose (84.2、B) は、大きなリポジトリの中で同じ形状を示します。4.8 KB で、コーデックスのファイルのほぼ 5 倍小さく、セットアップ、コマンド (ビルド / テスト / lint)、構造、および開発ループをカバーしています。オペレーショナルビートの百科事典。 A グレードの指示ファイルには、人数ではなく規律が必要です。
(調整のために、完全に開示されています: スコアラー自身のリポジトリ ファイルのスコアは 91.6 です。これが私が到達する方法を知っている上限です。そして 170 つ星のライブラリが単独でそれに到達しました。これはより興味深い事実です。)
16 番目の前に、GitHub コード検索から取得したパブリック AGENTS.md ファイルの 30 ファイル コーパスをスコア付けしました (2026 年 4 月に収集、8.5.0 で再スコア付け): 平均 61.5、中央値 61.4。 30 人中 17 人は D 以下に着陸します。総合すると、公開されている AGENTS.md の中央値は D-plus、つまり、読むのに問題はなく、何も動作しない散文です。
「点取り屋をこのまま試合にできないのか？」 #
このスコアラーの以前のバージョンは、まさに明らかな攻撃、つまり、フェンスで囲まれたコード ブロックにジャンク コマンドを詰め込んだ整形式の見出しに陥っていました。フェンスで囲まれた bash ブロックの echo 、 ls 、 pwd スコア

d 92.5、A、運用カバレッジ ディメンションが存在する前、これは私が静かにパッチを適用するのではなく仕様に書き込んだ失敗です。現在、その同じファイルのスコアは 56.0、D で、運用カバレッジは 0 です。読み取り専用動詞は何も獲得せず、完全なクレジットにはコマンドの多様性が必要です。
正直な限界が 1 つ残っており、それはテストによって文書化され固定されています。それは、もっともらしい捏造 (整形式のセクションで発明された、しかし現実的に見えるコマンド) は、静的スコアラーによっては実際の最小限のファイルと区別がつかないことです。保証の範囲は正確に定められています。価値のないテキストが運用上のテキストを上回ることはありません。これは「これが品質を評価する」よりも小さな主張であり、意図的に小さくされています。
これで何が測定され、何が測定されないのか #
ルーブリックは独自の意見があり、その重みはバージョン 1 です。 0.4 / 0.4 / 0.2 見出しは、グラウンド トゥルース ラベルなしの 30 ファイル コーパスに対して調整されました。現在の有効性の証拠は、ベンチマーク レベルではなく、実際の修正によってスコアが変化した前後の文書化されたケーススタディ レベルです。重み付けに同意できない場合は、ルーブリックが読みやすい辞書になっているので、その行を指さすことができます。
真実ではなく構造。スコアラーは、文書化されたコマンドが実際に実行されるかどうかを検証できません。行動ではなく、形式を読み取ります。
スナップショット 1 枚。すべてのスコアは、エンジン 8.5.0 での、以下にリストされている正確な HEAD コミットに対するものです。ファイルは変更されます。引用する前に再実行してください。
低いスコアは意図的な選択である可能性があります。最小限のファイルは正当な決定です。それはエージェントを装備しないだけであり、数字が主張しているのはそれだけです。別の監査では、インライン化には負荷がかかるため、14.6k トークンのスキル ファイルは効率の次元を 47 に維持しました。それは欠陥ではなく、情報に基づいた低下でした。
調査範囲は広範囲に及びますが、網羅的ではありません。候補者36名を調べてみました。ルート AGENTS.md で 20,000 スター以上のリポジトリを見逃した場合は、私に教えてください。

それを得点するだろう。
pip インストール シュリフ== 8.5.0
カール -sO https://raw.githubusercontent.com/openai/codex/main/AGENTS.md
シュリフスコア AGENTS.md
または、任意のファイルをプレイグラウンドに貼り付けるか、ライブバッジを独自のリポジトリにドロップします。判定モデルの代わりに決定論的スコアを使用する最大のポイントは、このテーブルがチェック可能であることです。上記のすべての数値は、読み取り、ピン留めし、議論できるルール エンジンから得られるものです。
得点 (16): 上の表。除外 (1): OpenInterpreter/open-interpreter — AGENTS.md バイトは openai/codex と同一です。チェック済み、ルートなし AGENTS.md (19): cline、Aider-AI/aider、Continuedev/Continue、google-gemini/gemini-cli、anthropics/claude-code、microsoft/vscode-copilot-chat、microsoft/autogen、geekan/MetaGPT、microsoft/semantic-kernel、TabbyML/tabby、AntonOsika/gpt-engineer、 SWE-agent/SWE-agent、stackblitz/bolt.new、Pythagora-io/gpt-pilot、reworkd/AgentGPT、TransformerOptimus/SuperAGI、getcursor/cursor、stitionai/devika、plandex-ai/plandex — それ自体が発見であり、私がチェックした最も有名なエージェント リポジトリの半数以上が、リポジトリ ルートにエコシステム独自の標準ファイルを配布していません。
HEAD コミット、2026-07-07: goose f96f62d9 · langchain 2d8100c4 · composio 040ffd49 · opencode 1c25b2f2 · qwen-code 40340ef5 · n8n 66ad8b93 · dify 6edce14e · AutoGPT e2711b17 · OpenHands cc80397e · キロコード b0348cbc · zed fc827a21 · ブラウザ使用 052787f9 · コーデックス cca16a10 · crewAI 799ab0f5 · smolagents 526069c1 · Roo-Code b867ec91 · 進行中 89e9bf9f 。
有名なフレームワーク、空のオペレーション
サンプル内の最高のファイルには 170 個のスターが付いています
「スコアラーをこのようにゲームすることはできませんか？」
これで何が測定され、何が測定されないのか
自分の MCP 貢献度をスコアリングする

## Original Extract

I scored the AGENTS.md files of the 16 biggest AI agent repos with a deterministic engine. ~1.46M combined stars, zero A grades, and the repo that popularized the convention ranks 13th of 16.

The State of AGENTS.md: Scoring the Agent-Makers' Own Instructions — Franz Paul Skip to content franz paul projects
The State of AGENTS.md: Scoring the Agent-Makers' Own Instructions
agents-md instruction-quality ai-agents tooling schliff 16 repos. About 1.46 million GitHub stars between them. Zero A grades.
AGENTS.md has quietly become the cross-tool standard for repo-level agent instructions — a single file that Cursor, Codex, Copilot, Claude Code and a growing list of others read on startup. So I went looking for how the AI-agent ecosystem itself uses the convention it created. I swept 36 of the best-known AI coding-agent and agent-framework repositories, found 16 with at least 20k stars shipping a root AGENTS.md , and scored every file with a deterministic engine — no LLM judge, same file same score on every machine, reproducible against the exact commit SHAs at the bottom of this post.
The result: mean 70.0, median grade C, not a single A. The repo whose team popularized the convention ranks 13th of 16. And the best AGENTS.md I could find anywhere belongs to a library with 170 stars.
This is the same scorer I used to grade my own merged MCP contribution and the one behind the case for static instruction-file security . It is open source, MIT, stdlib-only : the rubric is a dict you can read, not a hidden prompt.
The set is every well-known AI coding-agent or agent-framework project at or above 20k stars that ships a root AGENTS.md , each scored at its HEAD on 2026-07-07 with released schliff==8.5.0 . The AGENTS.md rubric is three dimensions: structure (weight 0.4), operational coverage (0.4), and efficiency (0.2).
Seven B, six C, two D, one E. No A, no S. † zed ships AGENTS.md as a symlink to its .rules file; the score is for the resolved content, and the symlink itself is a finding — a tool that fetches the raw AGENTS.md gets the six-character string .rules , not instructions.
One repo was excluded as a non-independent data point: OpenInterpreter’s AGENTS.md is byte-identical to openai/codex’s — the same git blob, since the project is a codex derivative for open models. Counting it twice would have been dishonest.
openai/codex — the repo whose team did the most to popularize AGENTS.md — scores 66.4, a C, 13th of 16. The file is not lazy: 22.5 KB across 28 headings, genuinely rich in review conventions and TUI style rules. But the operational-coverage dimension asks a narrow question: does this file equip an agent to operate the repo? Test and lint commands are there ( just test , just fix ). A build command is not. It appears nowhere in the file — for a Rust and Bazel monorepo. There are no commit or PR conventions either. I filed the gaps as an issue with an offer to fix .
The pattern generalizes across all sixteen. Structure is uniformly strong — 75 to 95, because headings are easy. What separates the top of the table from the bottom is operational coverage: the real setup, build, and test commands, and the gotchas an agent cannot cheaply rediscover at runtime.
Famous frameworks, empty operations #
Three files earn a D or E, and each survived me reading the raw Markdown by hand — these are not scorer artifacts:
crewAI (55.2, D): 1 KB of contribution platitudes — “follow best practices”, DRY, YAGNI. The only concrete commands are for the docs site ( mintlify dev ). Nothing tells an agent how to set up, build, or test the actual Python framework. Operational coverage: 15 of 100.
smolagents (50.0, D): the entire file is 157 bytes — four bullet points of generic advice (“Be Pythonic”, “follow OOP principles”). Operational coverage: 0.
Roo-Code (37.0, E): 445 bytes — one genuinely useful UI-state gotcha and nothing else. No setup, no build, no test. Operational coverage: 0.
A high structure score with zero operational coverage is the signature of a file written for a human skim, not for an agent that has to run the project.
The best file in the sample has 170 stars #
Not one of the sixteen giants reaches an A. The best AGENTS.md in my entire sample — the corpus below plus these sixteen — belongs to maxcountryman/underway , a Postgres-backed job queue with 170 stars: 91.0, grade A . It is 1,728 tokens covering setup, build, test, lint, and project structure as copy-pasteable commands.
block/goose (84.2, B) shows the same shape among the big repos: 4.8 KB, almost five times smaller than codex’s file, covering Setup, Commands (Build / Test / Lint), Structure, and a Development Loop. Operational beats encyclopedic. An A-grade instruction file costs discipline, not headcount.
(For calibration, and fully disclosed: the scorer’s own repo file scores 91.6. That is the ceiling I know how to reach — and a 170-star library reached it independently, which is the more interesting fact.)
Before the sixteen, I scored a 30-file corpus of public AGENTS.md files pulled from GitHub code search (collected 2026-04, re-scored on 8.5.0): mean 61.5, median 61.4. Seventeen of the thirty land at D or below. In aggregate, the median public AGENTS.md is a D-plus — prose that reads fine and operates nothing.
”Can’t you just game a scorer like this?” #
You could, and an earlier version of this scorer fell for exactly the obvious attack: well-formed headings with junk commands stuffed into fenced code blocks. echo , ls , pwd in a fenced bash block scored 92.5, an A, before the operational-coverage dimension existed — a failure I wrote into the spec rather than quietly patch. Today that same file scores 56.0, a D, with operational coverage 0: read-only verbs earn nothing, and full credit requires command diversity .
One honest limit remains, documented and pinned by a test: a plausible fabrication — invented but realistic-looking commands in well-formed sections — is indistinguishable from a real minimal file by any static scorer. The guarantee is scoped precisely: worthless text cannot outrank operational text. That is a smaller claim than “this measures quality”, and it is deliberately smaller.
What this measures, and what it does not #
The rubric is opinionated, and its weights are version one. The 0.4 / 0.4 / 0.2 headline was calibrated against a 30-file corpus without ground-truth labels. The validity evidence today is case-study-level — documented before-and-afters where the score moved with real fixes — not benchmark-level. If you disagree with a weight, the rubric is a readable dict and you can point at the line.
Structure, not truth. The scorer cannot verify that a documented command actually runs. It reads form, not behavior.
One snapshot. Every score is for the exact HEAD commit listed below, on engine 8.5.0. Files change; re-run before quoting.
A low score can be a deliberate choice. A minimal file is a legitimate decision — it simply will not equip an agent, which is the only thing the number claims. In a separate audit, a 14.6k-token skill file kept its efficiency dimension at 47 because inlining was load-bearing. That was an informed decline, not a defect.
The sweep is broad, not exhaustive. I checked 36 candidates. If I missed a repo above 20k stars with a root AGENTS.md , tell me and I will score it.
pip install schliff== 8.5.0
curl -sO https://raw.githubusercontent.com/openai/codex/main/AGENTS.md
schliff score AGENTS.md
Or paste any file into the playground , or drop the live badge into your own repo. The whole point of using a deterministic score instead of a judge model is that this table is checkable — every number above comes out of a rule engine you can read, pin, and argue with.
Scored (16): the table above. Excluded (1): OpenInterpreter/open-interpreter — AGENTS.md byte-identical to openai/codex’s. Checked, no root AGENTS.md (19): cline, Aider-AI/aider, continuedev/continue, google-gemini/gemini-cli, anthropics/claude-code, microsoft/vscode-copilot-chat, microsoft/autogen, geekan/MetaGPT, microsoft/semantic-kernel, TabbyML/tabby, AntonOsika/gpt-engineer, SWE-agent/SWE-agent, stackblitz/bolt.new, Pythagora-io/gpt-pilot, reworkd/AgentGPT, TransformerOptimus/SuperAGI, getcursor/cursor, stitionai/devika, plandex-ai/plandex — itself a finding: more than half of the best-known agent repos I checked do not ship the ecosystem’s own standard file at the repo root.
HEAD commits, 2026-07-07: goose f96f62d9 · langchain 2d8100c4 · composio 040ffd49 · opencode 1c25b2f2 · qwen-code 40340ef5 · n8n 66ad8b93 · dify 6edce14e · AutoGPT e2711b17 · OpenHands cc80397e · kilocode b0348cbc · zed fc827a21 · browser-use 052787f9 · codex cca16a10 · crewAI 799ab0f5 · smolagents 526069c1 · Roo-Code b867ec91 · underway 89e9bf9f .
Famous frameworks, empty operations
The best file in the sample has 170 stars
"Can't you just game a scorer like this?"
What this measures, and what it does not
Scoring My Own MCP Contribution
