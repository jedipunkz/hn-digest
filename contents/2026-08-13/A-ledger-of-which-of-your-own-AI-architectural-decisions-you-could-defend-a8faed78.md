---
source: "https://github.com/Dupflo/decision-ledger"
hn_url: "https://news.ycombinator.com/item?id=49284900"
title: "A ledger of which of your own AI architectural decisions you could defend"
article_title: "GitHub - Dupflo/decision-ledger: Three Claude Code agent skills that record which structural decisions you could actually defend, and give it back as a map of your project. · GitHub"
author: "dupflo"
captured_at: "2026-08-13T12:45:58Z"
capture_tool: "hn-digest"
hn_id: 49284900
score: 1
comments: 1
posted_at: "2026-08-13T12:26:05Z"
tags:
  - hacker-news
  - translated
---

# A ledger of which of your own AI architectural decisions you could defend

- HN: [49284900](https://news.ycombinator.com/item?id=49284900)
- Source: [github.com](https://github.com/Dupflo/decision-ledger)
- Score: 1
- Comments: 1
- Posted: 2026-08-13T12:26:05Z

## Translation

タイトル: 自分自身の AI アーキテクチャ上の決定のうちどれを擁護できるかを示す台帳
記事のタイトル: GitHub - Dupflo/決定台帳: どの構造的決定を実際に擁護できるかを記録し、それをプロジェクトの地図として返す 3 つのクロード コード エージェント スキル。 · GitHub
説明: どの構造的決定を実際に防御できるかを記録し、それをプロジェクトの地図として返す 3 つのクロード コード エージェント スキル。 - Dupflo/決裁台帳

記事本文:
GitHub - Dupflo/意思決定台帳: どの構造的決定を実際に擁護できるかを記録し、それをプロジェクトの地図として返す 3 つのクロード コード エージェント スキル。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
デュフロ
/
決裁台帳
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14 コミット 14 コミット .claude-plugin .claude-plugin .mastery .mastery アセット アセット bin bin ledger-audit ledger-audit

元帳レポート 元帳レポート 元帳 元帳参照 参照スクリプト スクリプト .gitignore .gitignore ライセンス ライセンス README.md README.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Claude Code の 3 つのエージェント スキル: /ledger 、 /ledger-audit 、 /ledger-report 。
2 人の開発者が今週同じ機能をリリースします。 3週間後、まだ説明できるのはそのうちの1人だけです。
diff には、あなたがどちらであるかを示すものは何もありません。 Decision Ledger は、記録を保持する Claude Code の 3 つのエージェント スキルです。構造上の決定ごとに、ポジションを取るように求められ、それを保持できるかどうかをメモし、プロジェクトの地図として渡します。この四半期には 12,000 回線が出荷されました。見直しの中で、これらの決定のうち何件を擁護できますか?
ステータス:初期。これは実験であり、製品ではありません。まだ使用状況データはありません。私はまず自分のプロジェクトでそれを実行しており、誰も摩擦を許容しないという答えがあるかどうかも含め、見つけたものを公開します。
これらは、通常のクロード コード スキルとして ~/.claude/skills にインストールされます。それぞれ 1 つの Promise があり、最初の Promise だけが自動的にロードされます。他の 2 つは、他のスラッシュ コマンドと同様に、名前で呼び出します。
自分でコマンドを実行することはありません。状態は、スキルが必要とする小さなスクリプトの背後に存在します。 CLI を学ぶ必要はありません。
Next.js アプリに認証を追加しています。何かを書く前に:
httpOnly Cookie のセッションですか、それともクライアント側で保持されているトークンですか?そしてなぜそれが？
クッキーの方が安全だと言いましたね。それは正しい判断ですが、これでは理屈の半分なので、もう一度尋ねます。
XSS の角度については同意。モバイル クライアントが同じ API にアクセスすると何が起こるでしょうか?
あなたはそれについて考えていませんでした。ブロックすることはなく、 auth/session-storage:partial を記録するため、とにかく Cookie のバージョンを書き込みます。
3 週間後、プル リクエストで次のようになりました。
マスタリーマップ: 24 の決定を追跡、21

日
auth/ ██████░░░░ 防御された 5 つのエリアのうち 2 つ <- 最も薄いエリア
データモデル/ █████████░ 5 個中 4 個、1 個が朽ちた
ui/ ██████████ 7/7
求人/ ████░░░░░░ 5 件中 2 件、1 件不明
観察された転送: 2 (セッション-ストレージ、キャッシュ-無効化)
このリポジトリ自体の台帳、実際には
この例は例示的なものです。これはそうではありません。これは .mastery/index.json で、ここでコミットされ、ツール自体を実行することによって生成されます。
スキル/スキルの粒度は 1 つまたは 3 つのスキルに相当
state/personal-layer-split で調整されたプロジェクト ファイルとローカル ファイル
状態/崩壊範囲の部分的な面積当たりの半減期、未解決
配布/インストール範囲を調整したグローバルデフォルトとプロジェクトごと
state/decay-scope:partial がポイントです。半減期は領域ごとに計算されるため、未知のものを auth/ で記録すると、正当化された近隣の信頼が更新されます。それは間違いであり、私はそれを修正したり反論したりしていません。それは、作成者自身のプロジェクトのファイル内に公開されています。
無防備な意思決定を測定し、それ自体を何も持たずに出荷するツールには何の価値もありません。
クロード コード ユーザー用のネイティブ パス。このリポジトリをプラグイン マーケットプレイスとして追加し、プラグインをインストールします。 3 つのスキルは、1 つの共有個人レイヤーを持つ 1 つの製品であるため、一緒に提供されます。
/プラグイン マーケットプレイス Dupflo/意思決定台帳を追加
/プラグインのインストール決定-ledger@dupflo
/plugin は、Claude Code ユーザーが実際にスキルを探す場所であるため、これがスキルを表示するパスになります。マニフェストは .claude-plugin/ にあります。1 つは 3 つのスキル ディレクトリを指す plugin.json で、もう 1 つはプラグインをリストする Marketplace.json です。
エージェント間、skills.sh ディレクトリ (Claude Code、Codex、Cursor など) 経由。これはスキルが発揮されるもう 1 つの場所です。ディレクトリはインストール テレメトリから自動的にインデックス付けされるため、

これらのいずれかを追加すると、少し高くなります。
npx スキルで Dupflo/意思決定台帳を追加
3 つのスキル ディレクトリはリポジトリ ルートにあり、自己完結型で、この CLI が期待するレイアウトとまったく同じです。 npxスキル追加。 --list は 3 つすべてを検索します。
または、どちらのディレクトリも使用せずに、スクリプトによって同じ 3 つのスキルを使用します。
カール -fsSL https://raw.githubusercontent.com/Dupflo/decion-ledger/main/install.sh |バッシュ
デフォルトではグローバル、 ~/.claude/skills 内にあり、すべてのリポジトリで利用可能です。これは意図的なものです。個人レイヤーはプロジェクトにまたがるため、あるコードベースで擁護したコンセプトは次のコードベースでは再質問されず、プロジェクトごとのインストールによってそのコンセプトは破棄されます。
./install.sh --project # このリポジトリのみ、./.claude/skills
./install.sh --uninstall # シンボリックリンクを削除し、すべての台帳を保持します
再実行すると、その場で更新されます。 3 つのスキルは 1 つのチェックアウトにシンボリックリンクされているため、バラバラになることはありません。 PATH に python3 が必要です。他には何も必要なく、パッケージも必要ありません。
次に、インターセプトをセッション間で持続させるために、プロジェクトの CLAUDE.md に次の内容を追加します。
## 意思決定台帳
構造的に重要な決定（データモデル、国家所有権、
認証戦略、モジュール境界、広がる依存関係、キャッシュ、
再試行セマンティクス)、台帳スキルをロードしてそれに従ってください。完全にスキップ
セッションが渡されるとき --ship。
状態はリポジトリのルートにある .mastery/index.json にあります。コミットします。台帳はプロジェクトに属しており、プロジェクトとともに移動します。記録された最初の決定によってファイルが作成されるため、セットアップ手順はありません。
これをコミットすると、1 つの注意が必要になります。無防備なセキュリティ上の決定を説明するメモには、プロジェクトの弱点も説明されています。プライベート リポジトリには、まさにあなたが書き留めておきたい内容が含まれています。公的なものでは、地図を描かずに決定に名前を付けるようにメモを書きます。 「すべての書き込みは行レベルのセキュリティをバイパスし、調停されることはありません」

十分なことを言い、適切な人にそれを言います。
機能するが後でコストがかかるコードを表す言葉が「技術的負債」です。理解せずに出荷した決定のためのものはありません。
静かに増えていきます。エージェントに認証を要求し、認証を取得すると、機能します。 3 週間後、誰かがセッションが Cookie 内に存在する理由を尋ねましたが、答えがないことがわかりました。その選択が間違っていたからではなく、それを選択したのはあなたではなかったからです。
このリポジトリの背後にある賭けは、生産性を感じていることと有能であるという感覚は静かに乖離する可能性があり、適切なタイミングで質問すれば、その乖離は測定可能であるということです。クイズではなく、コードが存在する前に、自分のプロジェクトで実際の選択を行った場合です。
これらは制約であり、機能ではありません。それらのほとんどは、このツールの明白なバージョンが初日にアンインストールされるために存在します。
作業が妨げられることはありません。クイズもレッスンもゲートもありません。コードは常に同じ順番で書かれます。 --ship はセッション中のすべてを無効にします。
文脈を無視してではなく、決定時に尋ねます。データ モデル、状態の所有権、認証戦略、モジュールの境界、広がる依存関係、キャッシュ、再試行セマンティクスなど、元に戻すのにコストがかかる選択のみ。質問は、ツールごとに 1 つではなく、インストールしたすべてのツールにわたってセッションごとに 1 つまでです。
「わかりません」が有効な答えです。それは記録され、コードが書かれ、実行中に推論が説明されます。このコンセプトは、後で隣接する決定に反映されます。あなたが何もせずにそれに答えると、それは転送、つまりファイル内で最も強い信号です。
それは曲線ではなく地図です。上昇するラインは最適化を促し、エンゲージメント指標になります。これは、これが拒否するために存在するものです。自信も失われます。触れていない領域では半減期が 90 日になります。
それはあなた自身ではなく、プロジェクトについて説明します。グローバルスコアなし、タイミングなし

あ、ランキングはなし。これらのフィールドは意図的に省略されています。 References/index-schema.md を参照してください。
あなたの歴史はあなたを追っています。それはあなたを追いかけません。あるプロジェクトで擁護したコンセプトが、次のプロジェクトでゼロから問い直されることはありません。これは ~/.decion-ledger/personal.json にあります。これはローカルのみ、モード 0600 で、決してコミットされず、エクスポート コマンドはありません。プロジェクト台帳は共有可能であり、誰が何を理解したかについては何も記載されていません。個人用ファイルはマシンから離れることはなく、忘れるように要求すると、そのファイルが存在する 1 つのファイルが削除されます。
コンセプトが古くなると、新しいプロジェクトで一度質問され、別のコードベースで正しく答えることが、ツールが収集できる最も強力な移転証拠となります。
このプロジェクトが下した決定
構造上の決定を守るためのツールは、あなた自身が持つ義務があります。これらはそれぞれ上記の台帳に記載されており、それぞれに何らかの費用がかかります。
スキルは 1 つではなく 3 つです。インターセプトは要求されずに起動する必要があるため、広範な説明が必要であり、ルーティングに関して広範な説明が互いに競合します。監査とレポートに簡潔なコマンド スタイルの記述を与えることで、競合から遠ざけ、決定ごとに実際に必要なものまでファイルをロードし続けることができます。代償は現実的です。「1 つのスキル、1 つの約束」のほうがすっきりした話であり、3 つのスキルは 3 つの同期を維持する必要があります。
個人レイヤーは 2 番目のファイルです。単一の元帳のほうがシンプルです。ただし、プロジェクト ファイルは、誰が何を理解したかを言わなくてもチームが読み取れる必要があり、個人履歴はエクスポートできなくてもプロジェクト間で存続する必要があります。これら 2 つの要件は 1 つのファイル内に存在できないため、存在しません。
自信が失墜する。上昇するだけのインデックスはバッジであり、バッジはエンゲージメント指標であり、まさにこれが拒否するものです。モジュールを開いていない人

3か月後の彼らは彼らほど理解していませんでした。減衰した概念が再び質問可能になります。これは間隔をあけた検索であり、正しく再回答することは、最初の回答よりも強力な証拠となります。
ここには門など何もない。モデルが回答が薄いと判断したために自分のプロジェクトで誰かをブロックすると、最初の誤検知で確実にアンインストールされます。プル リクエストでの警告 (「340 行追加、認証で 2 つの決定/記録には何も残らない」) は、密室で警告するよりも正直で効果的です。
すでに存在するプロジェクトから始める
どのプロジェクトも初めてのことです。 /ledger-audit は、実際のスキーマ、モジュール、依存関係に基づいて、構築中にスキルが尋ねるであろう質問を再構築するため、空のファイルから開始する必要はありません。
docs/decions/ または docs/adr/ が存在する場合、それはそこから始まります。それらの決定はすでに分離され、範囲が設定され、日付が付けられています。エージェントによって作成された ADR は、可能な限り最も鋭い質問です。なぜなら、それはプロジェクトが決定したことを文書化する一方で、それを擁護できるかどうかについては何も証明しないからです。
CLAUDE.md および AGENTS.md も読み取ります。理由もなくそこに記載されている規約は、「継承」とマークされ、コードベース全体に適用され、誰も擁護しません。多くの場合、それらは別のプロジェクトからコピーされたものであり、理解不能です。

[切り捨てられた]

## Original Extract

Three Claude Code agent skills that record which structural decisions you could actually defend, and give it back as a map of your project. - Dupflo/decision-ledger

GitHub - Dupflo/decision-ledger: Three Claude Code agent skills that record which structural decisions you could actually defend, and give it back as a map of your project. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Dupflo
/
decision-ledger
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits .claude-plugin .claude-plugin .mastery .mastery assets assets bin bin ledger-audit ledger-audit ledger-report ledger-report ledger ledger references references scripts scripts .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
Three agent skills for Claude Code: /ledger , /ledger-audit , /ledger-report .
Two developers ship the same feature this week. In three weeks, only one of them can still explain it.
Nothing in the diff tells you which one you are. Decision Ledger is three agent skills for Claude Code that keep the record: at each structural decision they ask you to take a position, note whether you could hold it, and hand it back as a map of your project. You shipped twelve thousand lines this quarter. How many of those decisions could you defend in review?
Status: early. This is an experiment, not a product. There is no usage data yet. I'm running it on my own projects first and will publish what I find, including if the answer is that nobody tolerates the friction.
They install as normal Claude Code skills, in ~/.claude/skills . One promise each, and only the first one loads on its own. The other two you invoke by name, like any slash command.
You never run a command yourself. State lives behind a small script the skills call for you; there is no CLI to learn.
You're adding auth to a Next.js app. Before writing anything:
Session in an httpOnly cookie or a token held client-side? And why that one?
You say the cookie is safer. That's the right call, but it's half the reasoning, so it asks once more:
Agreed on the XSS angle. What happens when a mobile client hits the same API?
You hadn't thought about it. It writes the cookie version anyway, because it never blocks you, and records auth/session-storage: partial .
Three weeks later, at a pull request:
Mastery map: 24 decisions tracked, 21 days
auth/ ██████░░░░ 2 of 5 defended <- thinnest area
data-model/ █████████░ 4 of 5, 1 decayed
ui/ ██████████ 7 of 7
jobs/ ████░░░░░░ 2 of 5, 1 unknown
Transfers observed: 2 (session-storage, cache-invalidation)
This repo's own ledger, for real
That example is illustrative. This one is not. It is .mastery/index.json , committed here, produced by running the tool on itself:
skills/skill-granularity justified one skill or three
state/personal-layer-split justified project file vs local file
state/decay-scope partial per-area half-life, unresolved
distribution/install-scope justified global default vs per project
state/decay-scope: partial is the point. The half-life is computed per area, so recording an unknown in auth/ refreshes the confidence of its justified neighbours. That is wrong, and I have not fixed it or argued it away. It sits in the file, in public, on the author's own project.
A tool that measures undefended decisions and ships without any of its own would be worth nothing.
The native path, for Claude Code users. Add this repo as a plugin marketplace, then install the plugin; the three skills arrive together, since they are one product with one shared personal layer.
/plugin marketplace add Dupflo/decision-ledger
/plugin install decision-ledger@dupflo
/plugin is where Claude Code users actually look for skills, so this is the path that gets it seen. The manifests live in .claude-plugin/ : one plugin.json that points at the three skill directories, one marketplace.json that lists the plugin.
Cross-agent, via the skills.sh directory (Claude Code, Codex, Cursor, and the rest). This is the other place skills get seen: the directory indexes automatically from install telemetry, so every one of these adds it a little higher.
npx skills add Dupflo/decision-ledger
The three skill directories sit at the repo root, self-contained, exactly the layout this CLI expects; npx skills add . --list finds all three.
Or, without either directory, the same three skills by script:
curl -fsSL https://raw.githubusercontent.com/Dupflo/decision-ledger/main/install.sh | bash
Global by default, in ~/.claude/skills , available in every repo. That is deliberate: the personal layer spans projects, so a concept you defended on one codebase is not re-asked on the next, and a per-project install throws that away.
./install.sh --project # this repo only, ./.claude/skills
./install.sh --uninstall # removes the symlinks, keeps every ledger
Re-running it updates in place. The three skills are symlinked into one checkout, so they never drift apart. Requires python3 on your PATH, nothing else, no packages.
Then, for the interception to persist across sessions, add to your project's CLAUDE.md :
## decision-ledger
Before any structurally significant decision (data model, state ownership,
auth strategy, module boundaries, a dependency that will spread, caching,
retry semantics), load the ledger skill and follow it. Skip entirely
when the session is passed --ship.
State lives in .mastery/index.json at your repo root. Commit it: the ledger belongs to the project, and it travels with it. There is no setup step, since the first decision recorded creates the file.
One caution follows from committing it. A note describing an undefended security decision also describes where the project is weak. On a private repo that is exactly what you want written down. On a public one, write the note so it names the decision without drawing the map. "All writes bypass row-level security, never arbitrated" says enough, and says it to the right people.
We have a word for code that works but costs us later: technical debt. We don't have one for the decisions we shipped without understanding them.
It accrues quietly. You ask an agent for auth, you get auth, it works. Three weeks later someone asks why the session lives in a cookie and you find you have no answer. Not because the choice was wrong, but because you were never the one who made it.
The bet behind this repo is that feeling productive and being capable can drift apart silently, and that the drift is measurable if you ask at the right moment: not in a quiz, but on a real choice, in your own project, before the code exists.
These are the constraints, not features. Most of them exist because the obvious version of this tool gets uninstalled on day one.
It never blocks your work. No quizzes, no lessons, no gates. The code always gets written in the same turn. --ship disables everything for a session.
It asks at the decision, not out of context. Only on choices that are expensive to reverse: data model, state ownership, auth strategy, module boundaries, a dependency that will spread, caching, retry semantics. At most one question per session, across every tool you have installed, not one per tool.
"I don't know" is a valid answer. It gets recorded, the code gets written, and the reasoning is explained in the doing. The concept comes back later on a neighbouring decision. If you answer it then unaided, that's a transfer , the strongest signal in the file.
It's a map, not a curve. A rising line invites optimisation and becomes an engagement metric, which is the thing this exists to reject. Confidence also decays: 90-day half-life on areas you haven't touched.
It describes the project, not you. No global score, no timing data, no ranking. Those fields are deliberately absent. See references/index-schema.md .
Your history follows you; it doesn't follow you around. A concept you defended on one project isn't re-asked from scratch on the next. That lives in ~/.decision-ledger/personal.json , which is local only, mode 0600 , never committed, with no export command. The project ledger is shareable and says nothing about who understood what; the personal one never leaves your machine, and asking for it to be forgotten deletes the one file it lives in.
Once a concept has gone stale, it gets asked once in the new project, and answering it correctly in a different codebase is the strongest transfer evidence the tool can collect.
The decisions this project made
A tool about defending structural decisions owes you its own. Each of these is in the ledger above, and each cost something.
Three skills, not one. Interception must fire without being asked for, so it needs a broad description, and broad descriptions compete with each other for routing. Giving the audit and the report terse, command-style descriptions keeps them out of that competition, and keeps the file loaded at every decision down to what is actually needed there. The cost is real: "one skill, one promise" was a cleaner story, and three skills are three things to keep in sync.
The personal layer is a second file. A single ledger would have been simpler. But the project file has to be readable by a team without saying who understood what, and a personal history has to survive across projects without ever being exportable. Those two requirements cannot live in one file, so they don't.
Confidence decays. An index that only rises is a badge, and badges are engagement metrics, the exact thing this rejects. Someone who hasn't opened a module in three months understands it less well than they did. A decayed concept becomes askable again, which is spaced retrieval, and re-answering correctly is stronger evidence than the first answer was.
Nothing here is a gate. Blocking someone on their own project because a model judged their answer thin is a guaranteed uninstall on the first false negative. A warning at a pull request ("340 lines added, two decisions in auth/ with nothing on record") is more honest and more effective than a closed door.
Starting on a project that already exists
Which is every project, the first time. /ledger-audit reconstructs the questions the skill would have asked while you were building, grounded in your actual schema, modules and dependencies, so you don't start from an empty file.
Where docs/decisions/ or docs/adr/ exists, it starts there: those decisions are already isolated, scoped and dated, and an ADR written by an agent is the sharpest possible question, because it documents what the project decided while proving nothing about whether you can defend it.
It also reads CLAUDE.md and AGENTS.md . Conventions stated there without a reason are marked inherited : applied across the whole codebase, defended by nobody. Often they were copied from another project, which is comprehension debt

[truncated]
