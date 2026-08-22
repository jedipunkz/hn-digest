---
source: "https://skillworks.kynth.studio"
hn_url: "https://news.ycombinator.com/item?id=49398601"
title: "Show HN: SkillWorks – every Claude Code skill, scored on whether it loads"
article_title: "Every Claude Code skill, scored on whether it actually loads · SkillWorks"
image: "https://skillworks.kynth.studio/og.jpg"
author: "kyisaiah47"
captured_at: "2026-08-22T12:18:58Z"
capture_tool: "hn-digest"
hn_id: 49398601
score: 2
comments: 0
posted_at: "2026-08-22T11:26:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: SkillWorks – every Claude Code skill, scored on whether it loads

- HN: [49398601](https://news.ycombinator.com/item?id=49398601)
- Source: [skillworks.kynth.studio](https://skillworks.kynth.studio)
- Score: 2
- Comments: 0
- Posted: 2026-08-22T11:26:04Z

## Translation

タイトル: HN を表示: SkillWorks – すべてのクロード コード スキル、ロードされたかどうかでスコア付けされます。
記事のタイトル: すべてのクロード コード スキル、実際にロードされるかどうかでスコア付けされる · SkillWorks
説明: 私たちが見つけることができるすべての Claude Code スキル、サブエージェント、プラグイン、およびマーケットプレイスのリポジトリには、実際に読み込まれるかどうかについて 0 ～ 100 のスコアが公開されています。 GitHub と skill.sh から毎晩再構築されます。何も手で選ばれることはなく、スポンサーが数字を動かすことはできません。

記事本文:
すべてのクロード コード スキル、実際に読み込まれるかどうかに基づいてスコア付けされます · SkillWorks SkillWorks Skills
すべてのクロード コード スキル、スコア
ソースから読み取り、毎晩 4 つの加重コンポーネントで 0 ～ 100 のスコアを獲得しました。
503,570 のリスト · 392,227 のスキル · 80,636 のサブエージェント · 23,442 のプラグイン · 7,265 のマーケットプレイス · 追跡されたインストール数を持つ 6,267 · ロードされない 48,190 · 再構築 2026-08-22
4 つのタイプすべてにわたるインデックス内の最もスコアの高いリスト。各スコアは、それを生成する 4 つの重み付けされたコンポーネントに分類されます。ここでは何も厳選されず、スポンサーが数字を動かすことはできません。アーティファクト スコア タイプ 動作数 40 メンテナンス 25 採用 20 ドキュメント 15 リーチ 最終コミット スキルクリエイター アンソロピクス/スキル 99 スキル 100 94 89 100 今日のインストール数 358,827 インストール数 find-skills vercel-labs/skills 98 スキル 100 100 94 74 3,055,082 インストール数 3日前 ui-ux-pro-max nextlevelbuilder/ui-ux-pro-max-skill 97 スキル 100 100 87 100 昨日のインストール数 325,271 pptx 人文学/スキル 97 スキル 100 94 86 91 今日のインストール数 206,139 pdf 人文学/スキル 97 スキル 100 94 86 91 今日のインストール数 182,719 docx anthropics/skills 97 スキル 100 94 85 91 今日のインストール数 175,275 mcp-builder anthropics/skills 97 スキル 100 94 83 92 今日のインストール数 104,949turborepo vercel/turborepo 97 スキル100 100 77 90 64,154 今日のインストール数 vercel-optimize vercel-labs/agent-skills 97 スキル 100 94 76 100 54,920 今日のインストール数 Sandbox-bench vercel/next.js 97 スキル 100 100 71 100 今日の 142,000 個のスター Microsoft-foundry microsoft/azure-skills 96 スキル 100 100 79 100 今日のインストール数 544,819 個 azure-quotas Microsoft/azure-skills 96 スキル 100 100 78 100 今日のインストール数 406,483 504k をすべて参照 すべての行は、合格したチェック、失敗したチェック、および各コンポーネントが計算された事実にリンクしています。から。数字で表したインデックス
クロードコードがロードできる 4 つのもの
えー

ch は、異なるファイル形式であり、異なる方法で破壊されるため、それぞれが個別にインデックス付けされ、スコア付けされます。
クロード コードがオンデマンドでロードする命令のフォルダー。SKILL.md ファイルによって定義されます。
名前付きエージェントには独自のプロンプト、モデル、およびツールの許可リストがあり、agents/ の下のマークダウン ファイルによって定義されます。
スキル、エージェント、フック、MCP サーバー、および LSP サーバーのパッケージ化されたバンドル。マーケットプレイスのリポジトリからインストールされます。
ルートで .claude-plugin/marketplace.json を公開するリポジトリ。これは、Claude Code が /plugin Marketplace add を使用してプラグイン ソースとして追加します。
README リストでは分からないこと
代わりとなるのは、誰も再チェックしない素晴らしいリストと、星を数えるスクレーパーです。どちらも「これは存在しますか？」と答えます。どちらも「これは読み込まれますか？」とは答えません。
4 つの重み付けされたコンポーネント。4 つすべてがリポジトリから直接読み取ります。実行もモデルも何も判断せず、手作業でキュレーションすることもなく、ポイントを購入する方法もありません。
解析するのか、名前と説明を宣言するのか、実際に存在するファイルを参照するのか。ここでの失敗は、ロードしないという判定を決定するものです。
それを公開するリポジトリでの最後のコミットからの日数 (急な曲線上)、およびそのリポジトリがアーカイブされているかどうか。
インストール数は、skills.sh レジストリが存在する場合に取得されます。リポジトリスターと、そうでない場合の 30 日間のスター速度。
手順、実際に動作するサンプル、バンドルされたスクリプトとリファレンスの長さ - モデルが実際にどれだけ進めなければならないか。
人気があり、スターがいるが、構造的に壊れている。これは他のディレクトリが公開していないインデックスの半分であり、これがこのディレクトリが存在する主な理由です。これらが最も多くの視聴者を抱える 6 つです。
決定論的なキーワード ルールによってアーティファクト自体の名前、説明、リポジトリ トピックから割り当てられ、毎晩新たに適用されます。カウントは、paのリストの数です

すべての構造チェックを行ってください。
他の誰もファイルを再読み取りしません
クロード コード スキルのすべてのリストは一度組み立てられ、それ以来朽ち続けています。リポジトリはアーカイブされ、フロントマター ブロックは解析されなくなるものに編集され、参照されたファイルの名前は変更されます。そして、推奨するリストは元の場所にそのまま残ります。何もチェックされていないため、まだ機能するかどうかは何もわかりません。
このインデックスはファイルを読み取ります。すべての SKILL.md、すべてのエージェント マークダウン ファイル、すべてのプラグインとマーケットプレイスのマニフェストは、それを公開するリポジトリからフェッチされ、Claude Code 独自のローダーが適用するものと同じ構造チェックが 1 晩に 1 回実行されます。ここにある 503,570 件のリストのうち 48,190 件はそのうちの 1 つが失敗し、失敗したチェックがその行に名前を付けてリストに残ります。それ自体の悪い結果を隠すディレクトリが再びリストになります。
スコアは 4 つの測定されたコンポーネントであり、それ以外は何もありません。スキルが優れているかどうかはモデルに問われず、サンドボックスでは何も実行されず、誰もキュレーションしません。スポンサーシップはランキングの外にあり、数字を動かすことはできません。これらすべては、他に何かを言う前に /methodology に書き留められます。
リポジトリを送信する これも Kynth Studios から
SkillWorks と同じ人向けに構築
AI エージェント ツールの独立した測定
パブリック shadcn レジストリを項目ごとに検索
SaaS スターター キット、チェックできる内容のスコア付け
1 つの製品を月に 1 回分解
Kynth Studios は、毎月 1 つの出荷済み製品をプルしてオープンします。その機能、構築にかかるコスト、その背後にあるパイプラインがどのようなものなのか、そして数値はどうなったのか。メールは月に 1 通、その間は何もありません。
ダブル オプトイン — 確認リンクは 1 つ送信され、それをクリックするまで他には何も送信されません。
Toolproof 、 BlockDex 、 KitGrade の背後にあるスタジオ
Toolproof の一部、AI エージェント ツールの測定レイヤー

g
スキル サブエージェント プラグイン マーケットプレイス CLAUDE.md の例 ランキング
最高のサブエージェント 最高のプラグイン 最も多くインストールされている インデックスは読み込まれません
検索カテゴリ スコアリング方法 リファレンス
リポジトリを送信する セクションのスポンサーになる リスト JSON プライバシー © 2026 SkillWorks. Kynth Studios の製品。変更履歴
1 つの製品を月に 1 回分解
Kynth Studios は、毎月 1 つの出荷済み製品をプルしてオープンします。その機能、構築にかかるコスト、その背後にあるパイプラインがどのようなものなのか、そして数値はどうなったのか。メールは月に 1 通、その間は何もありません。
ダブル オプトイン — 確認リンクは 1 つ送信され、それをクリックするまで他には何も送信されません。
Toolproof 、 BlockDex 、 KitGrade の背後にあるスタジオ
Toolproof の一部、AI エージェント ツールの測定レイヤー
スキル サブエージェント プラグイン マーケットプレイス CLAUDE.md の例 ランキング
最高のサブエージェント 最高のプラグイン 最も多くインストールされている インデックスは読み込まれません
検索カテゴリ スコアリング方法 リファレンス
リポジトリを送信する セクションのスポンサーになる リスト JSON プライバシー © 2026 SkillWorks. Kynth Studios の製品。変更履歴
1 つの製品を月に 1 回分解
Kynth Studios は、毎月 1 つの出荷済み製品をプルしてオープンします。その機能、構築にかかるコスト、その背後にあるパイプラインがどのようなものなのか、そして数値はどうなったのか。メールは月に 1 通、その間は何もありません。
ダブル オプトイン — 確認リンクは 1 つ送信され、それをクリックするまで他には何も送信されません。
Toolproof 、 BlockDex 、 KitGrade の背後にあるスタジオ
Toolproof の一部、AI エージェント ツールの測定レイヤー
スキル サブエージェント プラグイン マーケットプレイス CLAUDE.md の例 ランキング
最高のサブエージェント 最高のプラグイン 最も多くインストールされている インデックスは読み込まれません
検索カテゴリ スコアリング方法 リファレンス
リポジトリを送信する セクションのスポンサーになる リスト JSON プライバシー © 2026 SkillWorks. Kynth Studios の製品。変更履歴

## Original Extract

Every Claude Code skill, subagent, plugin and marketplace repo we can find, each with a published 0–100 score for whether it actually loads. Rebuilt from GitHub and skills.sh every night. Nothing is hand-picked and a sponsor cannot move a number.

Every Claude Code skill, scored on whether it actually loads · SkillWorks SkillWorks Skills
Every Claude Code skill, scored
Read from the source and scored 0–100 on four weighted components, every night.
503,570 listings · 392,227 skills · 80,636 subagents · 23,442 plugins · 7265 marketplaces · 6,267 with a tracked install count · 48190 that do not load · rebuilt 2026-08-22
The highest-scoring listings in the index, across all four types , each score broken into the four weighted components that produced it. Nothing here is hand-picked and a sponsor cannot move a number. Artifact Score Type works 40 maint 25 adopt 20 docs 15 Reach Last commit skill-creator anthropics/skills 99 Skill 100 94 89 100 358,827 installs today find-skills vercel-labs/skills 98 Skill 100 100 94 74 3,055,082 installs 3 days ago ui-ux-pro-max nextlevelbuilder/ui-ux-pro-max-skill 97 Skill 100 100 87 100 325,271 installs yesterday pptx anthropics/skills 97 Skill 100 94 86 91 206,139 installs today pdf anthropics/skills 97 Skill 100 94 86 91 182,719 installs today docx anthropics/skills 97 Skill 100 94 85 91 175,275 installs today mcp-builder anthropics/skills 97 Skill 100 94 83 92 104,949 installs today turborepo vercel/turborepo 97 Skill 100 100 77 90 64,154 installs today vercel-optimize vercel-labs/agent-skills 97 Skill 100 94 76 100 54,920 installs today sandbox-bench vercel/next.js 97 Skill 100 100 71 100 142k stars today microsoft-foundry microsoft/azure-skills 96 Skill 100 100 79 100 544,819 installs today azure-quotas microsoft/azure-skills 96 Skill 100 100 78 100 406,483 installs today Browse all 504k Every row links to the checks it passed, the checks it failed, and the facts each component was computed from. The index, in numbers
Four things Claude Code can load
Each is a different file shape with a different set of ways to be broken, so each is indexed and scored separately.
Folders of instructions Claude Code loads on demand, defined by a SKILL.md file.
Named agents with their own prompt, model and tool allow-list, defined by a markdown file under agents/.
Packaged bundles of skills, agents, hooks, MCP servers and LSP servers, installed from a marketplace repo.
Repos that publish a .claude-plugin/marketplace.json at their root, which Claude Code adds as a plugin source with /plugin marketplace add.
What a README list cannot tell you
The alternatives are an awesome-list nobody re-checks and a scraper that counts stars. Both answer “does this exist”. Neither answers “will this load”.
Four weighted components, all four read straight from the repository. No execution, no model judging anything, no hand curation, and no way to buy a point.
Does it parse, does it declare a name and a description, does it reference files that are actually there. A failure here is what sets the will-not-load verdict.
Days since the last commit on the repository that publishes it, on a steep curve, and whether that repository has been archived.
Install counts from the skills.sh registry where they exist; repository stars and 30-day star velocity where they do not.
Length of the instructions, worked examples, bundled scripts and references — how much a model actually has to go on.
Popular, starred, and structurally broken. This is the half of the index the other directories do not publish, and it is most of the reason this one exists — these are the six with the largest audiences.
Assigned from the artifact’s own name, description and repository topics by deterministic keyword rules, applied fresh every night. Counts are of listings that pass every structural check.
Nobody else re-reads the files
Every list of Claude Code skills was assembled once and has been decaying ever since. A repository is archived, a frontmatter block is edited into something that no longer parses, a referenced file is renamed — and the list recommending it stays exactly where it was. Nothing on it says whether any of it still works, because nothing on it ever checked.
This index reads the file. Every SKILL.md, every agent markdown file, every plugin and marketplace manifest is fetched from the repository that publishes it and put through the same structural checks Claude Code’s own loader applies, once a night. 48190 of the 503,570 listings here fail one of them, and they stay listed, with the check they failed named on the row — a directory that hides its own bad results is a list again.
The score is four measured components and nothing else. No model is asked whether a skill is good, nothing is executed in a sandbox, nobody curates, and sponsorship sits outside the ranking and cannot move a number. All of that is written down at /methodology before it says anything else.
Submit a repository Also from Kynth Studios
Built for the same person as SkillWorks
Independent measurement of AI agent tooling
Search public shadcn registries, item by item
SaaS starter kits, scored on what can be checked
One product, taken apart, once a month
Kynth Studios pulls one shipped product open every month — what it does, what it cost to build, what the pipeline behind it looks like, and what the numbers did. One email a month, nothing in between.
Double opt-in — we send one confirmation link and nothing else until you click it.
the studio behind Toolproof , BlockDex and KitGrade
part of Toolproof , the measurement layer for AI agent tooling
Skills Subagents Plugins Marketplaces CLAUDE.md examples Rankings
Best subagents Best plugins Most installed Will not load The index
Search Categories How we score Reference
Submit a repository Sponsor a section Listings JSON Privacy © 2026 SkillWorks. A Kynth Studios product. Changelog
One product, taken apart, once a month
Kynth Studios pulls one shipped product open every month — what it does, what it cost to build, what the pipeline behind it looks like, and what the numbers did. One email a month, nothing in between.
Double opt-in — we send one confirmation link and nothing else until you click it.
the studio behind Toolproof , BlockDex and KitGrade
part of Toolproof , the measurement layer for AI agent tooling
Skills Subagents Plugins Marketplaces CLAUDE.md examples Rankings
Best subagents Best plugins Most installed Will not load The index
Search Categories How we score Reference
Submit a repository Sponsor a section Listings JSON Privacy © 2026 SkillWorks. A Kynth Studios product. Changelog
One product, taken apart, once a month
Kynth Studios pulls one shipped product open every month — what it does, what it cost to build, what the pipeline behind it looks like, and what the numbers did. One email a month, nothing in between.
Double opt-in — we send one confirmation link and nothing else until you click it.
the studio behind Toolproof , BlockDex and KitGrade
part of Toolproof , the measurement layer for AI agent tooling
Skills Subagents Plugins Marketplaces CLAUDE.md examples Rankings
Best subagents Best plugins Most installed Will not load The index
Search Categories How we score Reference
Submit a repository Sponsor a section Listings JSON Privacy © 2026 SkillWorks. A Kynth Studios product. Changelog
