---
source: "https://github.com/MicSm/boffin"
hn_url: "https://news.ycombinator.com/item?id=49060279"
title: "Show HN: Boffin – Staff-engineer layer for AI coding agents"
article_title: "GitHub - MicSm/boffin: Staff-engineer layer for AI coding agents: routes per-edit architectural constraints and requires verification. Not another AGENTS.md. · GitHub"
author: "mic_sm"
captured_at: "2026-07-26T17:54:42Z"
capture_tool: "hn-digest"
hn_id: 49060279
score: 2
comments: 0
posted_at: "2026-07-26T17:28:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Boffin – Staff-engineer layer for AI coding agents

- HN: [49060279](https://news.ycombinator.com/item?id=49060279)
- Source: [github.com](https://github.com/MicSm/boffin)
- Score: 2
- Comments: 0
- Posted: 2026-07-26T17:28:03Z

## Translation

タイトル: Show HN: Boffin – AI コーディング エージェントのスタッフ エンジニア層
記事のタイトル: GitHub - MicSm/boffin: AI コーディング エージェントのスタッフ エンジニア層: 編集ごとのアーキテクチャ上の制約をルーティングし、検証が必要です。別のAGENTS.mdではありません。 · GitHub
説明: AI コーディング エージェントのスタッフ エンジニア層: 編集ごとのアーキテクチャ上の制約をルーティングし、検証を必要とします。別のAGENTS.mdではありません。 - MicSm/ボフィン
HN テキスト: AI コーディング エージェントのスタッフ エンジニア層: AI コーディング エージェントの編集ごとのアーキテクチャ上の制約をルーティングする

記事本文:
GitHub - MicSm/boffin: AI コーディング エージェントのスタッフ エンジニア層: 編集ごとのアーキテクチャ上の制約をルーティングし、検証を必要とします。別のAGENTS.mdではありません。 · GitHub
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
マイクスマ
/
棺
公共
通知

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
40 コミット 40 コミット .agents/ ルール .agents/ ルール .claude-plugin .claude-plugin .clinerules .clinerules .codex-plugin .codex-plugin .cursor/ ルール .cursor/ ルール .github .github .kiro/ ステアリング .kiro/ ステアリング .repos .repos .windsurf/ ルール.windsurf/ ルール アセット アセット ビン ビン ドキュメント ドキュメント サンプル サンプル 外部 外部フック フック lib lib パック パック スクリプト スクリプト シグネチャ シグネチャ スキル スキル 仕様 スペック テスト テスト .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md COTRIBUTING.md CONTRIBUTING.md クレジット クレジット ライセンス ライセンス README.md README.md バージョン バージョン gemini-extension.json gemini-extension.json logo.png logo.png package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Boffin (npm: boffinit ) は、AI コーディング用のスタッフ エンジニア制御レイヤーです。
エージェント: エージェントに、正確なファイルのアーキテクチャ上の制約を提供します。
を編集して結果を検証させます -- DuckDB のケーススタディ:
ガイド付きリファクタリングは +17 / -17 行に到達し、2,104 個のアサーションが通過しました。
あなたは 15 行の修正を要求します。エージェントは 500 行を改修して戻ってきました。
Boffin は、エージェントに、ファイルに適用されるアーキテクチャ上の制約を与えます。
が触れようとしていて、その結果を確認させます。
これは別の AGENTS.md ではなく、プロンプト パックでもありません。これらの形式は通常出荷されます
リポジトリ全体に対して 1 つの静的命令ブロック。ボフィンルートのみ
現在の編集に関連する制約。 ParselFire コアを搭載しています。
リポジトリ全体の静的なルール ファイルではありません (AGENTS.md スタイルの 1 ブロックですべてが対応)
即時パックではない、または

システムプロンプトのトリック
リンターや CI ゲートではなく、編集の前後に機能します。
スピードツールではありません -- フレームはレビューセーフティです
エージェントの目の前で編集に関連する制約を選択します
変更に応じた検証が必要
カーソル、クロード コード、コーデックス、および OpenCode を 1 つの npm パッケージから出荷
ParselFire Core を搭載した署名付きポータブル パックを提供します
どう違うのか（正直な比較）
静的ルール ファイル ( AGENTS.md )
ボフィン
配送
通常、リポジトリ全体に対して 1 つの静的ブロック
現在の編集にルーティングされた制約
検証
フォーマットでは必要ありません
必須、変化に比例
証拠
通常は何もありません
事例を数字とともに記録
約束ではなく証拠
公開されているケーススタディでは、実際のオープンソース コードに対する次のガイド付きリファクタリングが記録されています。
アヒルDB : +17 / -17 ; 2,104 件のアサーション
8 つのテスト ファイルに合格しました。明確な継続パスと回復パスがあった
保存されています。
高速API: +16 / -33; 49 のテストに合格しました。
パブリック API の変更はありません。
LangChain : 同期/非同期の境界は
保存される。 4 つのテストに合格しました。
これらは再現可能なケーススタディであり、管理された A/B ベンチマークではありません。
Boffin には Node.js 18 以降が必要です。
npx ボフィニット カーソル
クロード・コード
/プラグイン マーケットプレイス MicSm/boffin を追加
/プラグインのインストール boffin@boffin
コーデックス
コーデックス プラグイン マーケットプレイスに MicSm/boffin を追加
コーデックスプラグイン追加 boffin@boffin
Codex はプラグイン フックを自動的には信頼しません。 Codex 内で /hooks を 1 回実行します
Boffin のフックを確認して信頼すること。それまではプラグインのスキルは機能しますが、
セッションごとの自動アクティベーションはオフのままです。
npx boffinit オープンコード
次に、OpenCode でプロジェクトを開きます。常時オンのガイダンスは、
opencode.json -> .boffin/AGENTS.md 。オンデマンド: /boffin 、
/boffin-review 、または boffin / boffin-review スキル。
インストールの詳細、コマンド、トラブルシューティング:
OpenCode の配信。
機械が欲しいですか？

ParselFire Core がどのように機能するかを読んでください。
類似したコードが常に同じコードであるとは限りません。ボフィンはエージェントに次のような理由を与えます。
実際の特殊なケースをマージし、同期/非同期の境界を曖昧にし、移動する前に停止します。
状態を所有者から遠ざけたり、集中的なタスクをコードベースのツアーに変えたりします。
焦点を絞った変更の場合、要求された範囲を小さく保ち、
編集を証明する最も狭いチェック。
無制限のリファクタリングまたはレビューの場合、最初に読み取り専用の監査が必要です。
その後、検証された結果が一度に 1 つずつ続きます。
クリーンアップが以前の正確性ルールと競合する場合は、正確性が優先されます。
ポイントは、エージェントを怖がらせないことです。高価な細部を作ることです
興味深い午後になる前に、はっきりと伝えてください。
Boffin は AGENTS.md とどう違うのですか?
AGENTS.md は通常、リポジトリ全体に対する 1 つの静的命令ファイルです。
Boffin は、エージェントがファイルに関連するアーキテクチャ上の制約のみをルーティングします。
を編集しようとしているため、変更に応じたチェックが必要です。
制約はどこから来るのでしょうか?
すべてのルールは、読み取り可能なバージョン管理されたマークダウンとしてこのリポジトリに出荷されます。
Packs/ 、パックは GPG 署名されています。インストール時に何も隠されていません
時間: 信頼する前に、パックを開いてすべてのルールを読んでください。編集時ボフィン
タッチされているファイルにこれらのルールのどれを適用するかを選択します。参照
ParselFire Core がルーティング マップに対してどのように機能するか。
私のコーディング エージェントが小さな修正を大規模な書き換えに変えるのはなぜですか?
あなたは小さな修正を求めます。エージェントは改修して戻ってきます。ボフィン
編集前に現在のファイルに耐荷重制約を挿入し、
その後に検証を強制します。
lite 、 full 、 max は何を変更しますか?
彼らは正確さではなく、クリーンアップへの意欲を調整します。
lite はクリーンアップの圧力を低く保ち、最小限の有用な変更を優先します。
max は、タスクが正当な場合に最も強いクリーンアップ プレッシャーを適用します。

それを認めます。
プラグイン ホストで、 /boffin lite 、 /boffin full 、または
/ボフィン最大。オフプロフィールはありません。
プロファイルによって安全フロアは変更されますか?
いいえ。どのプロファイルも同じ初期の正確性段階と拒否ルールを維持します。
信頼境界の検証、データ損失防止、セキュリティ、
アクセシビリティ要件。
Boffin はコマンド サンドボックスまたはセキュリティ ツールですか?
いいえ。Boffin はプロセスを分離したり、シェル コマンドをフィルターしたり、制限したりしません。
ファイルシステムまたはネットワークアクセス。生成されたアーキテクチャ上の決定をガイドします。
コード。コマンド サンドボックスとセキュリティ コントロールを独自のジョブに使用します。ボフィンは
違う仕事。
Cursor または OpenCode 統合をアンインストールするにはどうすればよいですか?
npx boffinit カーソルのアンインストール
npx boffinit opencode アンインストール
各アンインストーラーは、そのホストの管理ファイルのみを削除します。共有
他のホストがまだ残っている場合、.boffin/packs と .boffin/VERSION は残ります。
インストールされています。関連のないプロジェクト ファイルはそのまま残されます。
Boffin はテストやコード レビューを置き換えますか?
いいえ。どの契約が注目に値するかをエージェントに伝え、外部からの要求を提供します。
チェックしますが、リポジトリのテストとレビューのプロセスは依然として権威を持っています。
ポータブル アダプタは、 AGENTS.md 、 CLAUDE.md 、ワークスペースを読み取るホストをカバーします
ルールまたはリポジトリの指示。参照
ホストの配信とアダプター
テクニカルマップ。
リポジトリ: https://github.com/MicSm/boffin
エンジンのドキュメント: ParselFire コア
寄稿: COTRIBUTING.md
Boffin は MIT ライセンスの下で利用可能です。クレジットを参照してください。
AI コーディング エージェントのスタッフ エンジニア層: 編集ごとのアーキテクチャ上の制約をルーティングし、検証を必要とします。別のAGENTS.mdではありません。
Readme MIT ライセンス
貢献 このリポジトリを引用する アクティビティのスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Staff-engineer layer for AI coding agents: routes per-edit architectural constraints and requires verification. Not another AGENTS.md. - MicSm/boffin

Staff-engineer layer for AI coding agents: routes per-edit architectural constraints for AI coding agents

GitHub - MicSm/boffin: Staff-engineer layer for AI coding agents: routes per-edit architectural constraints and requires verification. Not another AGENTS.md. · GitHub
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
MicSm
/
boffin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
40 Commits 40 Commits .agents/ rules .agents/ rules .claude-plugin .claude-plugin .clinerules .clinerules .codex-plugin .codex-plugin .cursor/ rules .cursor/ rules .github .github .kiro/ steering .kiro/ steering .repos .repos .windsurf/ rules .windsurf/ rules assets assets bin bin docs docs examples examples external external hooks hooks lib lib packs packs scripts scripts signatures signatures skills skills spec spec tests tests .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md CREDITS CREDITS LICENSE LICENSE README.md README.md VERSION VERSION gemini-extension.json gemini-extension.json logo.png logo.png package.json package.json View all files Repository files navigation
Boffin (npm: boffinit ) is a staff-engineer control layer for AI coding
agents: it feeds the agent the architectural constraints for the exact file it
is editing and makes it verify the result -- DuckDB case study :
a guided refactor landed at +17 / -17 lines with 2,104 assertions passing.
You ask for a 15-line fix; the agent comes back with a 500-line renovation.
Boffin gives the agent the architectural constraints that apply to the file it
is about to touch, then makes it verify the result.
It is not another AGENTS.md and not a prompt pack. Those formats usually ship
one static instructions block for the whole repository; Boffin routes only the
constraints relevant to the current edit. Powered by ParselFire Core .
Not a static repo-wide rules file ( AGENTS.md -style one block for everything)
Not a prompt pack or system-prompt trick
Not a linter or CI gate -- it acts before and after the edit
Not a speed tool -- the frame is review-safety
Selects the constraints relevant to the edit in front of the agent
Requires verification proportional to the change
Ships for Cursor, Claude Code, Codex, and OpenCode from one npm package
Delivers signed portable packs powered by ParselFire Core
How it differs (honest comparison)
Static rules file ( AGENTS.md )
Boffin
Delivery
Usually one static block for the whole repo
Constraints routed to the current edit
Verification
None required by the format
Required, proportional to the change
Evidence
Usually none
Recorded case studies with numbers
Proof, not promises
The public case studies record these guided refactors on real open-source code:
DuckDB : +17 / -17 ; 2,104 assertions
across 8 test files passed; distinct continuation and recovery paths were
preserved.
FastAPI : +16 / -33 ; 49 tests passed;
no public API change.
LangChain : the sync/async boundary was
preserved; 4 tests passed.
These are reproducible case studies, not a controlled A/B benchmark.
Boffin requires Node.js 18 or newer.
npx boffinit cursor
Claude Code
/plugin marketplace add MicSm/boffin
/plugin install boffin@boffin
Codex
codex plugin marketplace add MicSm/boffin
codex plugin add boffin@boffin
Codex does not trust plugin hooks automatically. Run /hooks once inside Codex
to review and trust Boffin's hooks; until then the plugin's skills work but the
automatic per-session activation stays off.
npx boffinit opencode
Then open the project in OpenCode. Always-on guidance lands via
opencode.json -> .boffin/AGENTS.md . On demand: /boffin ,
/boffin-review , or the boffin / boffin-review skills.
Install details, commands, and troubleshooting:
OpenCode delivery .
Want the machinery? Read how ParselFire Core works .
Similar code is not always the same code. Boffin gives the agent a reason to
stop before it merges a real special case, blurs a sync/async boundary, moves
state away from its owner, or turns a focused task into a tour of the codebase.
For a focused change, it keeps the requested scope small and asks for the
narrowest check that proves the edit.
For an open-ended refactor or review, it requires a read-only audit first,
followed by one verified finding at a time.
When cleanup conflicts with an earlier correctness rule, correctness wins.
The point is not to make the agent timid. It is to make the expensive details
explicit before they become an interesting afternoon.
How is Boffin different from AGENTS.md?
AGENTS.md is usually one static instructions file for the whole repository.
Boffin routes only the architectural constraints relevant to the file the agent
is about to edit, then requires a check proportional to the change.
Where do the constraints come from?
Every rule ships in this repository as readable, versioned markdown under
packs/ , and the packs are GPG-signed. Nothing is hidden at install
time: open any pack and read every rule before trusting it. At edit time Boffin
selects which of those rules apply to the file being touched. See
how ParselFire Core works for the routing map.
Why does my coding agent turn small fixes into huge rewrites?
You ask for a small fix; the agent comes back with a renovation. Boffin
injects the load-bearing constraints for the current file before the edit and
forces verification afterward.
What do lite , full , and max change?
They tune cleanup ambition, not correctness:
lite keeps cleanup pressure low and favors the smallest useful change.
max applies the strongest cleanup pressure when the task justifies it.
On plugin hosts, select a profile with /boffin lite , /boffin full , or
/boffin max . There is no off profile.
Do profiles change the safety floor?
No. Every profile keeps the same early correctness stages and rejection rules,
including trust-boundary validation, data-loss prevention, security, and
accessibility requirements.
Is Boffin a command sandbox or security tool?
No. Boffin does not isolate processes, filter shell commands, or restrict
filesystem or network access. It guides architectural decisions in generated
code. Use command sandboxes and security controls for their own job; Boffin has
a different job.
How do I uninstall the Cursor or OpenCode integration?
npx boffinit cursor uninstall
npx boffinit opencode uninstall
Each uninstaller removes that host's managed files only. Shared
.boffin/packs and .boffin/VERSION stay if the other host is still
installed. Unrelated project files are left alone.
Does Boffin replace tests or code review?
No. It tells the agent which contracts deserve attention and requires external
checks, but your repository's tests and review process remain authoritative.
Portable adapters cover hosts that read AGENTS.md , CLAUDE.md , workspace
rules, or repository instructions. See
host delivery and adapters for
the technical map.
Repository: https://github.com/MicSm/boffin
Engine documentation: ParselFire Core
Contributions: CONTRIBUTING.md
Boffin is available under the MIT License . See credits .
Staff-engineer layer for AI coding agents: routes per-edit architectural constraints and requires verification. Not another AGENTS.md.
Readme MIT license Contributing
Contributing Cite this repository Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
