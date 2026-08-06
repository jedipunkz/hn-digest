---
source: "https://nesbitt.io/2026/08/06/a-year-of-ai-disclosure-in-critical-packages.html"
hn_url: "https://news.ycombinator.com/item?id=49194705"
title: "A year of AI disclosure in critical packages"
article_title: "A year of AI disclosure in critical packages | Andrew Nesbitt"
author: "omkar-foss"
captured_at: "2026-08-06T10:28:32Z"
capture_tool: "hn-digest"
hn_id: 49194705
score: 1
comments: 0
posted_at: "2026-08-06T10:06:11Z"
tags:
  - hacker-news
  - translated
---

# A year of AI disclosure in critical packages

- HN: [49194705](https://news.ycombinator.com/item?id=49194705)
- Source: [nesbitt.io](https://nesbitt.io/2026/08/06/a-year-of-ai-disclosure-in-critical-packages.html)
- Score: 1
- Comments: 0
- Posted: 2026-08-06T10:06:11Z

## Translation

タイトル: 重要なパッケージにおける AI 開示の 1 年
記事のタイトル: 重要なパッケージにおける AI 開示の 1 年 |アンドリュー・ネスビット
説明: 支援者: Daniel Stenberg

記事本文:
重要なパッケージにおける AI 開示の 1 年 |アンドリュー・ネスビット
アンドリュー・ネスビット
重要なパッケージにおける AI 開示の 1 年
Stephen O'Grady 氏による RedMonk によるオープンソース コードの記述者に関する分析では、2026 年上半期の 15 件の大規模プロジェクトへのコミットを調査し、宣言された AI 関与の 2 つの形態をカウントしました。それは、コミット作成者としての既知の自律エージェント、または Co-Authored-By トレーラー内の既知の AI ID です。結果は 1 パーセント未満で、床としてフレーム化されました。
私は、16 のレジストリにわたって最も依存しているパッケージの背後にある 5,682 の GitHub リポジトリである、packages.ecosyste.ms クリティカル セットに対して同じ測定のより広範なバージョンを実行しました。CHAOSS 公開ライブラリを使用して、2 種類ではなく 4 種類の明示的なシグナルを検出しました。同じ 6 か月間の比率は 4.13% でした。 2026 年 7 月 29 日までの 1 年間では、その割合は 2.93% (非マージ コミット 589,798 件中 17,279 件) で、昨年 8 月の 0.48% から今年 7 月には 5.32% に上昇しました。
これらは、誰かが git メタデータに明示的なマーカーを残したコミットの数です。未申告の使用は計測されず、1行変更しても1万行変更してもコミットは1単位となります。
サンプルの選択と検出器の選択
2 つのシグナルのみを使用して RedMonk の 15 のリポジトリに対してスキャナを実行したところ、マージを含む 23,346 の前半コミットのうち、RedMonk の「約 24,000 件のコミット」に対して 94 件の一致 (0.40%) が見つかり、一致数は「数十件」でした。マージを除外すると、17,323 個のコミットと同じ 94 個の一致 (0.54%) が残ります。 espressif/esp-idf と openssl/openssl はそのうち 71 個を提供しており、RedMonk が報告した 2 つのプロジェクトの集中率 73% と一致します。
2 つの追加の信号タイプを追加すると、どちらのサンプルでもレートが約 0.5 パーセント ポイント移動しました。サンプルを変更すると、サンプルが 3 ポイント移動しました。 RedMonkの15人はコントロールによって選ばれました

O'Grady 氏の言葉を借りれば、ibutor ベースのサイズは C に意図的に偏っています。重要なパッケージ セットは、各レジストリの依存関係グラフの最上位にあるものであり、より小規模で新しい、会社が運営するリポジトリが多数取り込まれます。
重要なスナップショットには 8,605 個のパッケージが含まれており、リポジトリ URL とメタデータは、Weekend at Bernie’s 用に構築したのと同じパッケージ キャッシュから取得されました。リポジトリを共有するパッケージをマージし、名前を変更し、GitHub に制限し、不正な形式の URL を削除することで、5,707 個の候補が残りました。 5,682 個のクローン作成が成功しました。残りの 25 件は削除または非公開でした。 2026 年 7 月 29 日に終了する年度に、3,533 社が少なくとも 1 つの非マージ コミットを行いました。
各リポジトリは、ツリー フィルターと浅い日付境界を使用して裸のままクローンされ、公開ライブラリを通じてストリーミングされ、削除されました。フルパスは約 1 GB 転送され、保持されるチェックポイントは 16 MB のリポジトリごとのサマリーと一致するコミット SHA です。
次の名前を変更すると、GitHub の安定したリポジトリ ID とリダイレクトがチェックされます。 npm パッケージ ベースには、依然としてリポジトリとして node-base/base がリストされています。 GitHub は組織名を再利用したため、そのパスは Base ブロックチェーン モノリポジトリにリダイレクトされるようになりました。これにより、9 年前の npm ユーティリティに 3,135 のコミットと 273 の AI シグナルが提供されたことになります。 IDチェックは対象外でした。
マージ以外のすべてのコミットは以下についてチェックされました。
著者またはコミッターとして既知の AI エージェント
Co-Authored-By の既知の AI アイデンティティ
AI ツールまたはモデルに名前を付ける Assisted-By トレーラー
開示がサポートするツール固有の帰属形式
マージは除外されるため、スカッシュ、リベース、またはマージを行うプロジェクトは同じ基準でカウントされます。コミットは、変更が現在のブランチに到達したときのコミッター時間によってバケット化されます。通常のコミット文でのツール名の言及は無視されます。トレーラーも同様であるため、Assisted-By 値が検証されます。

人物に使用: 生の一致が含まれます。支援者: Daniel Stenberg および支援者: 自動ツール、人間によるレビュー。クローンは refs/notes/ai をフェッチしなかったため、 git note として記録された宣言は存在しません。
月次金利は２月に３％、３月に５％を超え、その後７月まで４．５８％─５．３２％の間で推移した。コミットの代わりにリポジトリをカウントすると、2025 年 8 月にはコミットのある 1,734 個のリポジトリのうち 41 個 (2.4%)、2026 年 7 月には 1,793 個のうち 276 個 (15.4%) でシグナルが発生しました。
17,279 件の調査結果のうち、4,625 件は自律エージェント ID のみを伝え、12,628 件は宣言された支援信号のみを伝え、26 件は両方を伝えました。宣言された支援は、8月のコミットメントの0.08％から7月には4.92％に増加しました。エージェントのオーサーシップは 0.40% で始まり 0.41% で終了し、その間のピークは 1.33% でした。これらのコミットのうち 4,613 件には GitHub Copilot のエージェントが作成者として含まれており、38 件には Devin のエージェントが含まれており、その間の 25 件は Claude、Cursor、Codex、および Amazon Q アカウントです。 Copilot エージェントのコミットは 3 月に 85 のリポジトリで 745 に達しましたが、7 月には 35 のリポジトリで 208 に減少しました。個々のプロジェクトでエージェントが短期間に実行されました。pycqa/isort は 3 月に 49 でその後はゼロで、azure/azure-sdk-for-net は 2 月に 275、3 月に 28 でした。
合計の 2 月と 3 月のステップは、Claude Code Co-Authored-By トレーラーです。これらは、12 月の 97 件のコミットから、その後の 3 か月で 325、753、2,037 件に増加し、39 の異なるリポジトリから 190 件に増加しました。同じ月に Cursor の共著者トレーラーは 1 から 48 に増加し、Copilot は 1 から 2 に増加したため、このステップは開示慣行の一般的な変更ではなく、1 つのツールに固有のものです。 Anthropic は 2 月 5 日に Claude Opus 4.6 を、2 月 17 日に Sonnet 4.6 をリリースしました。 3 月は両方が利用できる最初の丸月です。
調査結果には、17,392 oc にわたって 231 の異なる宣言されたツール文字列が含まれています。

通貨。それらをクライアント ファミリごとにグループ化し、クライアントの名前が指定されていないモデルまたはプロバイダーごとにグループ化します。
宣言された生の文字列は概要 JSON にあります。グループ化は私によるもので、2 つのクライアントを指定する値は両方にカウントされます。
リポジトリ レベルでは、3,533 のアクティブ リポジトリのうち 687 が年間少なくとも 1 つのシグナルを記録したため、アクティブ リポジトリの割合の中央値はゼロです。検出結果が最も多い 10 のリポジトリが全体の 40.8% を占め、上位 100 のリポジトリが 84.9% を占めます。
go-git は、追加のディテクタが 1 つのリポジトリに何を追加するかを示します。その 731 コミットのうち 269 は検証済みのシグナルを伝送し、そのうち 12 は RedMonk の狭いルールに一致します。残りは、Assisted-By トレーラーとツールの帰属です。
16 のエコシステムのうち 9 つでは、ウィンドウ内に少なくとも 30,000 のコミットがありました。
残りの 7 つは、13,581 コミットの CocoaPods から 682 コミットの Julia まで、サマリー JSON に含まれています。 Julia の 682 件のコミットにおける 51 件の結果は、最小のサンプルで 7.47% というセット内で最高の割合を示しています。
NuGet の 6.84% は Microsoft のデプロイメントです。 aspnet 、 azure 、 azuread 、 dotnet 、 Microsoft 、および nuget のリポジトリは、NuGet の検出結果 2,842 件のうち 2,716 件 (95.6%) を提供し、そのうち 2,634 件は自律エージェント ID でした。これらの所有者を削除すると、NuGet は 14,283 コミット中 126 件の検出結果 (0.88%) に下がり、Maven よりも下回ります。 NuGet に対して所有者の除外を実行しただけです。リポジトリごとの CSV には、他の CSV を実行するために必要なものが含まれています。
同じ 5,682 のデフォルト ブランチ ヘッドを介した別のパスで、コーディング エージェントへのコミットされた命令がチェックされました: AGENTS.md 、 CLAUDE.md 、 GEMINI.md 、および文書化された Copilot、Cursor、Cline、Windsurf、および Continue ルール パス。 353 のリポジトリ (6.21%) に少なくとも 1 つのリポジトリがあり、それらの間に 1,091 個のファイルが保持されています。ファイルの存在は、誰かがエージェントに対するガイダンスを設定したことを記録します。それは何もありません

どのコミットに対してもngです。
各リポジトリは、最も古い生存命令ファイルが追加された月に 1 回カウントされるため、118 個のファイルを持つ Cypress は、1 個のファイルを持つプロジェクトと同じようにカウントされます。現在のヘッドに存在するファイルのみが表示されるため、追加されたものや後で削除されたものはタイムラインには表示されません。
353 のリポジトリのうち 228 にも、その年に公開されたコミットがあります。そのうち 94 人は最初に公開されたコミットの前に最初の命令ファイルを追加し、98 人はその後に追加し、36 人は同日に追加しました。中央値ギャップはゼロです。残りの 125 には命令ファイルがありますが、ウィンドウにはコミットが公開されていません。
スキャナーとレポート ジェネレーターは andrew/critical-ai-scan にあります。概要 JSON には、全体、月次、エコシステム、シグナル、ツール、および主要なリポジトリの数が含まれます。リポジトリ CSV には、正確なデフォルト ブランチ ヘッドが使用された成功したスキャンごとに 1 行が含まれるため、再クローンを作成せずに個々のケースを確認できます。命令ファイルのレポートには、一致したすべてのパスがそのカテゴリおよびそれを追加したコミットとともにリストされます。
2026 年の CHAOSS 指標
2026 年 5 月 27 日 CHAOSS メトリクスが人間の速度に貢献するように調整されました
メンテナーへのインタビュー
Jul 24, 2026 緑の広場第214話。
RFC: オープンソースへの人為的貢献者
2026 年 5 月 21 日 対象ステータス: 現在のベストプラクティス。
中心性は活力ではない
2026 年 5 月 14 日 依存関係グラフの PageRank に自動的に到達しない
セキュリティ上の問題ではない
2026 年 5 月 12 日 Curl の開示ポリシーが AI スキャナーの結果をソースでどのようにフィルタリングしたか

## Original Extract

Assisted-By: Daniel Stenberg

A year of AI disclosure in critical packages | Andrew Nesbitt
Andrew Nesbitt
A year of AI disclosure in critical packages
Stephen O’Grady’s RedMonk analysis of who is writing open source code looked at commits to fifteen large projects during the first half of 2026 and counted two forms of declared AI involvement: a known autonomous agent as the commit author, or a known AI identity in a Co-Authored-By trailer. The result was under one percent, framed as a floor.
I ran a wider version of the same measurement over the packages.ecosyste.ms critical set: 5,682 GitHub repositories behind the most-depended-on packages across sixteen registries, using the CHAOSS disclosure library to detect four kinds of explicit signal instead of two. Over the same six months the rate was 4.13%. Over the year ending 29 July 2026 it was 2.93% (17,279 of 589,798 non-merge commits), rising from 0.48% last August to 5.32% this July.
These are counts of commits where someone left an explicit marker in git metadata. Undeclared use is not measured, and a commit is one unit regardless of whether it changed one line or ten thousand.
Sample selection versus detector choice
Running my scanner against RedMonk’s fifteen repositories with only their two signals found 94 matches in 23,346 first-half commits including merges, or 0.40%, against RedMonk’s “~24K commits” and a match count “in the dozens”. Excluding merges leaves 17,323 commits and the same 94 matches, or 0.54%; espressif/esp-idf and openssl/openssl supply 71 of them, matching RedMonk’s reported 73% concentration in two projects.
Adding the two extra signal types moved the rate by about half a percentage point on either sample. Changing the sample moved it by three points. RedMonk’s fifteen were chosen by contributor-base size with, in O’Grady’s words, a deliberate bias towards C; the critical package set is whatever sits at the top of each registry’s dependency graph, which pulls in a lot of smaller, newer, company-run repositories.
The critical snapshot contained 8,605 packages, with repository URLs and metadata pulled from the same package cache I built for Weekend at Bernie’s . Merging packages that share a repository, following renames, restricting to GitHub, and dropping malformed URLs left 5,707 candidates. 5,682 cloned successfully; the other 25 were deleted or private. 3,533 had at least one non-merge commit in the year ending 29 July 2026.
Each repository was cloned bare with a tree filter and a shallow date boundary, streamed through the disclosure library, and deleted. The full pass transferred about 1 GB and the retained checkpoint is 16 MB of per-repository summaries and matched commit SHAs.
Rename following checks GitHub’s stable repository ID as well as the redirect. The npm package base still lists node-base/base as its repository. GitHub reused the org name, so that path now redirects to the Base blockchain monorepo, which would have contributed 3,135 commits and 273 AI signals to a nine-year-old npm utility. The ID check excluded it.
Every non-merge commit was checked for:
a known AI agent as author or committer
a known AI identity in Co-Authored-By
an Assisted-By trailer naming an AI tool or model
a tool-specific attribution format that disclosure supports
Merges are excluded so projects that squash, rebase, or merge count on the same basis. Commits are bucketed by committer time, when the change landed on the current branch. Mentions of tool names in ordinary commit prose are ignored. Assisted-By values are validated because the trailer is also used for people: raw matches included Assisted-By: Daniel Stenberg and Assisted-By: Automated Tooling, Human Reviewed. The clones did not fetch refs/notes/ai , so declarations recorded as git notes are absent.
The monthly rate passed 3% in February and 5% in March, then held between 4.58% and 5.32% through July. Counting repositories instead of commits, a signal appeared in 41 of the 1,734 repositories with commits in August 2025 (2.4%) and 276 of 1,793 in July 2026 (15.4%).
Of the 17,279 findings, 4,625 carry only an autonomous-agent identity, 12,628 carry only a declared-assistance signal, and 26 carry both. Declared assistance went from 0.08% of commits in August to 4.92% in July. Agent authorship started at 0.40% and ended at 0.41%, peaking at 1.33% in between; 4,613 of those commits have GitHub Copilot’s agent as author, 38 have Devin’s, and Claude, Cursor, Codex, and Amazon Q account for 25 between them. Copilot agent commits reached 745 in March across 85 repositories and fell to 208 across 35 in July, with individual projects running the agent in short bursts: pycqa/isort had 49 in March and none after, azure/azure-sdk-for-net had 275 in February and 28 in March.
The February and March step in the total is Claude Code Co-Authored-By trailers. Those went from 97 commits in December to 325, 753, and 2,037 over the following three months, and from 39 distinct repositories to 190. Cursor’s co-author trailers rose from 1 to 48 over the same months and Copilot’s from 1 to 2, so the step is specific to one tool rather than a general change in disclosure practice. Anthropic released Claude Opus 4.6 on 5 February and Sonnet 4.6 on 17 February; March is the first full month with both available.
The findings carry 231 distinct declared tool strings across 17,392 occurrences. Grouping them by client family, and separately by model or provider where no client is named:
The raw declared strings are in the summary JSON; the grouping is mine and a value naming two clients counts in both.
At the repository level, 687 of the 3,533 active repositories recorded at least one signal over the year, so the median active repository’s rate is zero. The ten repositories with the most findings account for 40.8% of the total and the top hundred for 84.9%.
go-git shows what the extra detectors add in one repository: 269 of its 731 commits carry a validated signal, 12 of which match RedMonk’s narrow rules. The rest are Assisted-By trailers and tool attributions.
Nine of the sixteen ecosystems had at least 30,000 commits in the window:
The other seven, from CocoaPods at 13,581 commits down to Julia at 682, are in the summary JSON. Julia’s 51 findings in 682 commits give it the highest rate in the set at 7.47%, on the smallest sample.
NuGet’s 6.84% is a Microsoft deployment. Repositories under aspnet , azure , azuread , dotnet , microsoft , and nuget supplied 2,716 of the 2,842 NuGet findings (95.6%), and 2,634 of those are autonomous-agent identities. Remove those owners and NuGet falls to 126 findings in 14,283 commits, or 0.88%, below Maven. I have only run that owner exclusion for NuGet; the per-repository CSV has what’s needed to do it for the others.
A separate pass over the same 5,682 default-branch heads checked for committed instructions to coding agents: AGENTS.md , CLAUDE.md , GEMINI.md , and the documented Copilot, Cursor, Cline, Windsurf, and Continue rule paths. 353 repositories (6.21%) have at least one, holding 1,091 files between them. A file’s presence records that someone set up guidance for an agent; it attributes nothing to any commit.
Each repository is counted once, in the month its earliest surviving instruction file was added, so Cypress with 118 files counts the same as a project with one. Only files present on current heads are visible, so anything added and later deleted is absent from the timeline.
228 of the 353 repositories also have a disclosed commit in the year. 94 of those added their earliest instruction file before their first disclosed commit, 98 added it after, and 36 on the same day; the median gap is zero. The other 125 have an instruction file and no disclosed commit in the window.
The scanner and report generator are at andrew/critical-ai-scan . The summary JSON has the overall, monthly, ecosystem, signal, tool, and leading-repository counts. The repository CSV has one row per successful scan with the exact default-branch head used, so individual cases can be checked without recloning. The instruction-file report lists every matched path with its category and the commit that added it.
CHAOSS Metrics in 2026
May 27, 2026 CHAOSS metrics were calibrated for human-speed contribution
Interview with a Maintainer
Jul 24, 2026 Episode 214 of Green Squares.
RFC: Artificial Contributors to Open Source
May 21, 2026 Intended status: Best Current Practice.
Centrality is not vitality
May 14, 2026 Don't automatically reach for PageRank on dependency graphs
Not a Security Issue
May 12, 2026 How curl's disclosure policy filtered an AI scanner's findings at source
