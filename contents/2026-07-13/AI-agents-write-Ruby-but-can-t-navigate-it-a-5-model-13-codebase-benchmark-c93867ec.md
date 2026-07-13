---
source: "https://github.com/luuuc/sense/blob/main/bench/verticals/ruby-rails/report-for-humans.md"
hn_url: "https://news.ycombinator.com/item?id=48896651"
title: "AI agents write Ruby but can't navigate it: a 5-model, 13-codebase benchmark"
article_title: "sense/bench/verticals/ruby-rails/report-for-humans.md at main · luuuc/sense · GitHub"
author: "luuuc"
captured_at: "2026-07-13T19:12:59Z"
capture_tool: "hn-digest"
hn_id: 48896651
score: 2
comments: 0
posted_at: "2026-07-13T18:21:35Z"
tags:
  - hacker-news
  - translated
---

# AI agents write Ruby but can't navigate it: a 5-model, 13-codebase benchmark

- HN: [48896651](https://news.ycombinator.com/item?id=48896651)
- Source: [github.com](https://github.com/luuuc/sense/blob/main/bench/verticals/ruby-rails/report-for-humans.md)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T18:21:35Z

## Translation

タイトル: AI エージェントは Ruby を書くがナビゲートできない: 5 モデル、13 コードベースのベンチマーク
記事のタイトル: sense/bench/verticals/ruby-rails/report-for-humans.md at main · luuuc/sense · GitHub
説明: AI コーディング エージェント用の MCP サーバー - クロード コード、OpenCode、カーソル、および Codex CLI にコードベースの構造的理解を提供します (シンボル グラフ、爆発半径、セマンティック検索、規則)。 - メインの sense/bench/verticals/ruby-rails/report-for-humans.md · luuuc/sense

記事本文:
メインの sense/bench/verticals/ruby-rails/report-for-humans.md · luuuc/sense · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ルウク
/
感覚
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする

原因 さらなるファイルアクション 最新のコミット
履歴 履歴 136 行 (89 loc) · 11.2 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルをコピー raw ファイルをダウンロード 概要 編集と raw アクション AI エージェントは Ruby をナビゲートできますか?結果を分かりやすく言うと
AI エージェントは Ruby をうまく書きます。このベンチマークは別の質問をします: 彼らはそれをナビゲートできるでしょうか?実際のコードベースと中央モデルに関わる変更が与えられた場合、エージェントは、変更によってコードベースが静かに壊れる前に、コードベースに依存するすべての場所を見つけることができるでしょうか?
13 の実際のコードベースと 5 つのモデルにわたって、答えはきれいに 2 つに分かれます。小さくて読みやすい gem では、エージェントの通常のツールで十分です。実際のアプリケーションでは、コールド エージェントは依存関係の一部を見つけて、監査が完了したことを自信を持って宣言します。同じエージェントにコードベースの構造マップを与えると、ギャップは縮まります。最も強いモデルは引用再現率で平均 +0.26 を獲得し、その獲得はほぼ完全に難しい部分、つまり何が依存するかに集中します (依存グループでは +0.48)。
開示: 私は「sense」アームのコードマップである Sense を構築していますが、ベンチは、手作業で構築された回答キー、ピン留めされたコミット、機械的に検証された引用、公開された差異、再生可能なトランスクリプトなどを気にする必要がないように設計されています。このページのすべての数値は、生のレポートと最後にリンクされているモデルごとのレポートに基づいています。
1 つのシナリオ シェイプ、すべてのリポジトリ:
中央モデルの解体方法を変えようとしています。それに触れる前に、それに依存するコードベース内のすべての場所を見つけてください。
Chatwoot の受信箱。マストドンのステータス 。 GitLab の MergeRequest 。談話のアップロード 。これは、上級エンジニアが危険なリファクタリングを行う前に尋ねる質問であり、グリーン レビューが失敗したデプロイになるかどうかを決定する質問です。
各リポジトリは、同じモデル、同じプロンプトでタスクを 2 回実行します。

、同じ固定されたコミット:
ベースライン アーム: エージェントの通常のツール (ファイル読み取り、grep、シェル)
センスアーム: 同じツールに加えて、リポジトリの構造マップをクエリする機能: そのシンボル、コールエッジ、およびフレームワークレベルの関係の永続化されたグラフ。ファイルウォークの代わりに単一のクエリで回答可能
実行前に、各シナリオは必ず見つけなければならないセットを取得します。つまり、完全な答えを明らかにするために必要な正確なコードの場所であり、ソースから手作業で構築されます。エージェントはそれを見ることはありません。
見出しの指標は、「そのセットのうち、エンジニアが直接ジャンプできる正確なパス:ラインに固定された回答のシェア」を引用しています。参考文献を意識した審査員が、主要な内容に照らしてカバレッジを採点するため、自信に満ちたように聞こえるが不完全な解答は、省略されている部分で減点されます。次に、答えが出力されたすべての path:line が、固定されたコミットのリポジトリに対して機械的に解決されます。解決しないものは、公開引用チェックでモデルごとにフラグが立てられます。
なぜリコールが引用されるのか?重要な故障モードは正直な漏れだからです。どちらの腕も幻覚はあまりありません。コールド エージェントは 11 人の依存関係者のうち 2 人を見つけ、監査が終了したと判断します。散文の質を評価するという指標は、まさにそれを見逃してしまうだろう。
5 つのモデル、それぞれ 13 リポジトリ、両腕。ヘッドライン アームでは、すべてのリポジトリがアームごとに 2 回実行されました。実行間のスプレッドは、リポジトリごとの分散で公開されます。デルタはセンスからベースラインを引いたものであるため、正の値はマップが役に立ったことを意味します。
ヘッドラインアーム (Opus): 13 回のレポで 12 勝 1 引き分け 0 敗。
この表の 2 つのことが合計よりも重要です。
リフトは硬い柱に着地します。すべてのモデルの依存リフトは全体のリフトを上回ります。モデルのパブリック API に名前を付けるのは簡単です。ポリモーフィックな関連付けを通じてそれにタッチしたワーカーを見つけるのは、午前 3 時に誰かを呼び出す部分であり、そこがマップの役割です。
得られた最も弱いモデル

最強と同じくらい。 24B オープンモデルの Devstral は、Opus の +0.26 に対して +0.25 を記録し、フロアははるかに低かった。マップは、モデルが作業メモリに保持できないものを補っており、保持できるものが少なくなるほど、マップの価値は高くなります。
リポジトリをサイズごとに並べ替えると、パターンが微妙ではなくなります。これらは、各リポジトリのインデックス付きファイル数に対する Opus デルタです。
約 2,500 ファイルを超えるすべてのリポジトリは少なくとも +0.25 増加しました。 2,000 ファイル未満のすべてのリポジトリは最大で +0.17 増加しました。マップは、モデルが自分自身で読み取ることができないコードベースのスライスに相当し、1 行以上の価値はありません。
4 行で全体のストーリーがわかります。
チャットウート、+0.68。冷静に、エージェントは受信トレイ モデルの散在する 11 個の依存ファイルのうち 2 個を見つけ、監査が完了したと宣言しました。再実行では 0/11 が見つかりました。マップでは 11/11 で、両方とも実行されます。見逃した依存関係は、モデルと grep トークンをコールド共有しません。これらは、関心事、サービス、および多態性の関連付けを通じてそれに到達し、grep は関係を確認できません。
ギットラボ、+0.41。 36,829 ファイル、インデックス内の 1,121,147 エッジ。エージェントの 1 つの sense_blast MergeRequest 呼び出しにより、解決された依存セット (影響を受ける 354 個のシンボル) が 1 ステップで返されました。この規模では、「コードを読むだけ」という計画はありません。
レール、+0.25。最も有益な行。コールド、このモデルは、Rails を効果的に記憶しているため、大規模なアプリケーションよりもはるかに強力です。両方のコールド ランで、ActiveRecord::Relation の 5 つのコントラクト項目すべてが固定されています。思い出すことができなかったのは、チュートリアルがこれまでに書いたことのないクエリ コンパイラの内部です。最悪のコールド ランでは 6 つの内部依存関係のうちの 1 つが引用されました。マップでは、両方の実行で 6 のうち 5 が得られました。有名なコードに対するエージェントの信頼は思い出すことですが、それはまさに書かれた内容の境界で切れてしまいます。プライベート リポジトリは完全にその境界の向こう側にあります。
私

lm.rb、-0.00。唯一の同点、意図的にボード上に保持されます。全体を読むのに十分な小さな宝石には地図は必要ありません。センスアームは、約 4 分の 1 少ない請求トークンで同じ質問に回答し、それだけです。測定対象を褒めるだけのベンチマークはパンフレットです。
逆方向に切断する細胞
リポジトリごとの完全なモデルテーブルには 65 個のセルがあり、そのすべてが成功するわけではありません。 GPT-5.5 は、langchainrb に -0.03 と 3 つのフラット ゼロをポストします。 Qwen は redmine に -0.03 を投稿します。 Devstral は、セット内で最小のリポジトリである raix に -0.20 を投稿します。このリポジトリは 38 ファイルからなり、マップのオーバーヘッドがツール オーケストレーションに弱いモデルの値を上回ります。とにかくパターンは当てはまります。リフトはコードベースのサイズに応じてスケールされ、負のセルは小さい端に位置し、緑色のセルのみを持つ表は広告であるため、それらは表内に表示されます。
シリーズを走ることでベンチの両サイドも研ぎ澄まされ、それがベンチマークの目的だ。 Mastodon はシンボルのあいまいさの行き詰まり (Status は Ruby モデルと React コンポーネントの両方) を暴露し、GitLab はハブ スケールでの非決定的な結果の上限を暴露し、Solidus は gem 全体で再開されたクラスのリゾルバー ギャップを暴露しました。各修正は次のリポジトリが実行される前に出荷され、現在は各修正がすべてのユーザーに配布されています。グレーディング側も同様に強化しました。最初の裁判官は 44% 完了した監査を「網羅的」と評価したため、手作りキーに対する参照を意識したグレーディングに置き換えられました。これらの変更はいずれも数値をより正確なものにし、そのうち 3 つは製品を改善しました。
思い出すことが目標です。トークンは副作用です。 13 のリポジトリすべてにわたるモデルごとの平均を計算すると、請求されたトークン (実際に支払う金額) はマップとともに次のように移動しました。
3 つのモデルが安くなり、1 つは横ばい、1 つは値上がりしました。提示された取引は少額ではないため、どちらの方法でも報告されます。 GitLab では、+ を購入した請求トークンが 19% 増加しました

36,829 ファイルのモノリスで引用リコールが 0.41 上昇。これは、あなたが経験しなかったインシデントに対する丸め誤差です。
(トークンの数値はモデル内のみです。モデル間で比較しないでください。ハーネスとプロバイダーではカウントが異なります。)
作者が構築したベンチマークが信頼できる理由
そうする必要はありません。これが設計上の制約でした。
実行前にソースから手動で作成された回答キー。エージェントがそれらを目にすることはありません。
コミットは固定されています。すべての実行は、固定された再現可能なツリーに対して行われます。
引用は機械的に検証されます。すべての回答のすべての path:line はリポジトリに対して解決されます。障害はモデルごとに公開されます。
差異が公開されました。各リポジトリの実行間スプレッドはリポジトリごとの分散であるため、ヘッドラインは安定している場合にのみ信頼されます。
失敗作が公開されました。同点、マイナスセル、交代されたジャッジ、ベンチの3つのツールの欠陥が表面化した。そのすべてが上記にあり、すべてが生データの中にあります。
すべてがリプレイ可能です。トランスクリプト、シナリオ、スクリプトはこのリポジトリにあります。 LSP アームや別のメトリックが状況を変えると思われる場合は、ハーネスがここにあります。
独自のリポジトリで 10 分で実行できます
今日の午後、法廷の主張はコードベース上で反証可能です。
冷静にエージェントに尋ねてください。「[最も多忙なモデル] の廃止方法を変更する前に、それに依存している場所をすべて見つけてください。」リストを保存します。
Sense をインストールします: README のワンライナー、またはソースからビルドします。
リポジトリ ルートでセンス スキャンを実行し、エージェントに接続するためのセンス セットアップを実行します。
一言一句同じ質問をしてください。 2 つのリストを、真実であるとわかっているものと比較してください。
Sense は無料、オープンソース、1 つのバイナリです。すべてがマシン上に留まり、どこにも何も送信されません。
リポジトリが小規模な gem の場合は、同数になることが予想されます。それがアプリケーションの場合、最初のリストには、最も検出したくない依存関係が含まれていないことが予想されます。

えー、本番中です。
生のレポート: 方法論、完全な 5 モデル × 13 リポジトリ デルタ マトリックス、効率テーブル
リポジトリごとの絶対スコアと引用チェックを含むモデルごとのレポート: Claude Opus 4.8 · GPT-5.5 · Kim K2.7 · Devstral Small 24B · Qwen3 Coder Next
リポジトリごとの差異: 各見出しは安定していますか、それとも幸運な抽選ですか?
シナリオと必須セット: 質問と回答キー
このガイドは、プロジェクト全体のドキュメント マップです。つまり、Sense の機能、コードベースでエージェントが Sense を使用する方法、コントリビューションが属する場所、および 1 バイナリ コード インデックスがどのように機能するかについて説明します。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

MCP server for AI coding agents - gives Claude Code, OpenCode, Cursor, and Codex CLI structural understanding of your codebase: symbol graph, blast radius, semantic search, conventions. - sense/bench/verticals/ruby-rails/report-for-humans.md at main · luuuc/sense

sense/bench/verticals/ruby-rails/report-for-humans.md at main · luuuc/sense · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
luuuc
/
sense
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 136 lines (89 loc) · 11.2 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions Can an AI agent navigate Ruby? The results, in plain language
AI agents write Ruby well. This benchmark asks a different question: can they navigate it? Given a real codebase and a change that touches a central model, can an agent find every place that depends on it, before the change silently breaks one?
Across 13 real codebases and 5 models, the answer splits cleanly in two. On small, readable gems, the agent's normal tools are enough. On real applications, the cold agent finds a fraction of the dependents and confidently declares the audit done. Give the same agent a structural map of the codebase and the gap closes: the strongest model gained a mean +0.26 in cited recall, and the gain concentrates almost entirely on the hard part, what depends on what ( +0.48 on the dependents group).
Disclosure: I build Sense , the code map in the "sense" arm, and the bench is designed so you don't have to care: hand-built answer keys, pinned commits, mechanically verified citations, published variance, replayable transcripts. Every number on this page comes from the raw report and the per-model reports linked at the end.
One scenario shape, every repo:
You are about to change how a central model is torn down. Before you touch it, find every place in the codebase that depends on it.
Chatwoot's Inbox . Mastodon's Status . GitLab's MergeRequest . Discourse's Upload . This is the question a senior engineer asks before any risky refactor, and it is the question that decides whether a green review becomes a broken deploy.
Each repo runs the task twice with the same model, same prompt, same pinned commit :
baseline arm: the agent's normal tools (file reads, grep, shell)
sense arm: the same tools, plus the ability to query a structural map of the repo: a persisted graph of its symbols, call edges, and framework-level relationships, answerable in single queries instead of file walks
Before any run, each scenario gets a must-find set : the exact code locations a complete answer has to surface, built by hand from source. The agent never sees it.
The headline metric is cited recall : of that set, the share the answer pinned to an exact path:line an engineer could jump straight to. A reference-aware judge grades coverage against the key, so a confident-sounding but incomplete answer is penalized for what it leaves out. Then every path:line the answer printed is mechanically resolved against the repo at the pinned commit; anything that doesn't resolve is flagged, per model, in a public citation check .
Why cited recall? Because the failure mode that matters is honest omission . Neither arm hallucinates much. The cold agent finds 2 of 11 dependents and calls the audit finished. A metric that rewarded prose quality would miss exactly that.
Five models, 13 repos each, both arms. On the headline arm every repo ran twice per arm; run-to-run spread is published in per-repo variance . The delta is sense minus baseline, so positive means the map helped.
On the headline arm (Opus): 12 wins, 1 tie, 0 losses across the 13 repos.
Two things in that table matter more than the totals.
The lift lands on the hard column. Every model's dependents lift beats its overall lift. Naming a model's public API is easy; finding the worker that touches it through a polymorphic association is the part that pages someone at 3am, and that is where the map pays.
The weakest model gained about as much as the strongest. Devstral, a 24B open model, posted +0.25 against Opus's +0.26, from a far lower floor. The map compensates for what a model cannot hold in working memory, and the less it holds, the more the map is worth.
Sort the repos by size and the pattern is not subtle. These are the Opus deltas against the indexed file count of each repo:
Every repo above ~2,500 files gained at least +0.25. Every repo below 2,000 files gained at most +0.17. The map is worth exactly the slice of the codebase the model cannot afford to read for itself, and not a line more.
Four rows tell the whole story:
Chatwoot, +0.68. Cold, the agent found 2 of 11 scattered dependents of the Inbox model and declared the audit done. On the rerun it found 0 of 11. With the map, 11 of 11, both runs. The dependents it missed cold share no grep token with the model; they reach it through concerns, services, and polymorphic associations, and grep cannot see relationships.
GitLab, +0.41. 36,829 files, 1,121,147 edges in the index. The agent's single sense_blast MergeRequest call came back with the resolved dependent set, 354 affected symbols, in one step. At this size, "just read the code" is not a plan.
Rails, +0.25. The most instructive row. Cold, the model is far stronger here than on any of the big applications, because it has effectively memorized Rails: it pinned all 5 contract items of ActiveRecord::Relation in both cold runs. What it could not recall are the query-compiler internals no tutorial ever wrote about: its worst cold run cited 1 of the 6 internal dependents; with the map, 5 of 6 in both runs. An agent's confidence on famous code is recall, and it runs out exactly at the boundary of what got written about. Your private repo is entirely on the far side of that boundary.
llm.rb, -0.00. The lone tie, kept on the board on purpose. A gem small enough to read whole needs no map; the sense arm answered the same question on about a quarter fewer billed tokens, and that's all. A benchmark that only ever flatters the thing it measures is a brochure.
The cells that cut the other way
The full model-by-repo table has 65 cells and not all of them are wins. GPT-5.5 posts -0.03 on langchainrb and three flat zeros. Qwen posts -0.03 on redmine. Devstral posts -0.20 on raix , the smallest repo in the set: a 38-file gem where the map's overhead outweighs its value for a model weak at tool orchestration. The pattern holds anyway. Lift scales with codebase size, the negative cells sit at the small end, and they are in the table because a table with only green cells is an ad.
Running the series also sharpened both sides of the bench, which is what a benchmark is for. Mastodon exposed a symbol-ambiguity dead end ( Status is both a Ruby model and a React component), GitLab exposed a nondeterministic result cap at hub scale, and Solidus exposed a resolver gap on classes reopened across gems; each fix shipped before the next repo ran, and each now ships to every user. The grading side hardened the same way: the first judge rated a 44%-complete audit "exhaustive," so it was replaced with reference-aware grading against the hand-built keys. Every one of those changes made the numbers more honest, and three of them made the product better.
Recall is the goal; tokens are a side effect. Averaged per model across all 13 repos, billed tokens (what you actually pay for) moved like this with the map:
Three models got cheaper, one stayed flat, one paid more. Reported either way, because the trade on offer is not a smaller bill. On GitLab, 19% more billed tokens bought a +0.41 lift in cited recall on a 36,829-file monolith. That is a rounding error against the incident you didn't have.
(Token figures are within-model only. Never compare them across models; the harnesses and providers count differently.)
Why you can trust a benchmark its author built
You shouldn't have to, which was the design constraint:
Answer keys built by hand from source, before any run. The agent never sees them.
Commits pinned. Every run is against a fixed, reproducible tree.
Citations mechanically verified. Every path:line in every answer is resolved against the repo; failures are published per model.
Variance published. Each repo's run-to-run spread is in per-repo variance , so a headline is trusted only when it is stable.
Failures published. The tie, the negative cells, the judge that got replaced, the three tool defects the bench surfaced. All of it is above, and all of it is in the raw data.
Everything replayable. Transcripts, scenarios, and scripts are in this repo. If you think an LSP arm or a different metric would change the picture, the harness is sitting right here.
Run it on your own repo in 10 minutes
The bench's claim is falsifiable on your codebase this afternoon:
Ask your agent, cold: "Before I change how [your busiest model] is torn down, find every place that depends on it." Save the list.
Install Sense: one-liner in the README , or build from source.
sense scan in the repo root, then sense setup to connect your agent.
Ask the same question, word for word. Diff the two lists against what you know is true.
Sense is free, open source, one binary. Everything stays on your machine and sends nothing anywhere.
If your repo is a small gem, expect a tie. If it is an application, expect the first list to be missing the dependents you'd least like to discover in production.
Raw report : methodology, the full 5-model × 13-repo delta matrix, efficiency tables
Per-model reports with absolute scores per repo and citation checks: Claude Opus 4.8 · GPT-5.5 · Kimi K2.7 · Devstral Small 24B · Qwen3 Coder Next
Per-repo variance : is each headline stable or a lucky draw?
Scenarios and must-find sets : the questions and the answer keys
The guide is the doc map for the whole project: what Sense does, how an agent uses it on a codebase, where a contribution belongs, and how a one-binary code index works.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
