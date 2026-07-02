---
source: "https://github.com/tomascupr/sous-chef"
hn_url: "https://news.ycombinator.com/item?id=48768400"
title: "Sous-Chef, a Claude Code plugin where Fable reviews, Codex implements"
article_title: "GitHub - tomascupr/sous-chef: Claude orchestrates and reviews; Codex CLI implements. A Claude Code plugin for the two-model kitchen: /serve tasks end to end, /simmer goals until green. Your head chef doesn't chop onions. · GitHub"
author: "tomcupr"
captured_at: "2026-07-02T23:09:30Z"
capture_tool: "hn-digest"
hn_id: 48768400
score: 1
comments: 0
posted_at: "2026-07-02T22:50:59Z"
tags:
  - hacker-news
  - translated
---

# Sous-Chef, a Claude Code plugin where Fable reviews, Codex implements

- HN: [48768400](https://news.ycombinator.com/item?id=48768400)
- Source: [github.com](https://github.com/tomascupr/sous-chef)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T22:50:59Z

## Translation

タイトル: Sous-Chef、Fable がレビューし、Codex が実装されるクロード コード プラグイン
記事のタイトル: GitHub - tomascupr/sous-chef: Claude がオーケストレーションとレビューを行います。 Codex CLI を実装します。 2 つのモデルのキッチン用の Claude Code プラグイン: / タスクをエンドツーエンドで提供し、 / 目標を緑色になるまで煮込みます。あなたの料理長は玉ねぎを刻みません。 · GitHub
説明: クロードはオーケストレーションとレビューを行います。 Codex CLI を実装します。 2 つのモデルのキッチン用の Claude Code プラグイン: / タスクをエンドツーエンドで提供し、 / 目標を緑色になるまで煮込みます。あなたの料理長は玉ねぎを刻みません。 - トマスキュル/副料理長

記事本文:
GitHub - tomascupr/副料理長: クロードがオーケストレーションとレビューを行います。 Codex CLI を実装します。 2 つのモデルのキッチン用の Claude Code プラグイン: / タスクをエンドツーエンドで提供し、 / 目標を緑色になるまで煮込みます。あなたの料理長は玉ねぎを刻みません。 · GitHub
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
トマスカップ

/
副料理長
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .claude-plugin .claude-plugin .claude .claude codex codex docs docs skill skill templates templates .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Fable 5 は調整とレビューを行います。 GPT-5.5 xhigh を実装します。あなたの料理長は玉ねぎを刻みません。
キッチンのように 2 つのフロンティア モデル間でコーディングを分割する Claude Code プラグイン
分割作業。クロードは計画を立て、チケットを書き、すべての差分を 1 行ずつ確認し、
チェック自体を再実行します。 Codex CLI はサンドボックス内で実装を行います。
バックグラウンドがあり、何を出荷するかについては何も言えません。判定とコーデックスにクロード トークンを使う
バルク上のトークン: このパターンが構築される測定されたセットアップでは、
Codex は実装作業を最大 20 倍行いました
オーケストレーションの往復あたり、2 つの中間層のサブスクリプションが 1 つの最上位層のサブスクリプションを上回ることがよくあります。
> /sous-chef:serve 認証モジュールを非推奨のセッション API から移行します
クロード「gpt-5.5 でのフルサービス: 実装、クロスレビュー、修正、検証。
2 ～ 3 回の Codex 実行、約 15 ～ 45 分。メッキができたら報告します。」
クロードはチケットを書きます - タッチするファイル、タッチしないファイル、いつ完了するか
基準、検証コマンド - そしてそれを起動します (バックグラウンド)
コーデックスが実装されています。 11個のファイルが変更されました。 「すべてのテストに合格した」と主張します。
クロードは diff を 1 行ずつ確認し、pnpm テストと tsc 自体を再実行します。
Codex は読み取り専用の差分をクロスレビューします。 3 つの発見
クロードはそれらをコードに照らして検証します: 2 つは確認され、1 つは反駁されました
クロードは 2 つの確認された結果を範囲指定された修正チケットとして再発行します
クロードは、引用された場所で各発見を再検証します。緑のチェック
Cl

aude "提供済み: 認証移行、11 ファイル、42 テスト合格、2 レビュー
調査結果は修正され、1 件は反論されました。 2 つのモデルは一致します。お皿はあなたのものよ。」
1 つのコマンドが入力され、1 つのメッキ結果が出力されます。コーデックスで「テストに合格」と言うのは文章です。
pnpm テストの出力は事実です - クロードはすべてを自分自身で再実行します。
/sous-chef:serve はタスク形式の作業用であり、エンドツーエンドで実行されます。実装、
クロスレビューし、発見事項を修正し、検証します。事前の発表は 1 つ、会場での報告は 1 つ
最終的には、その間に 5 つの Codex を実行するという厳しい予算がかかります。こちらは日替わりドライバーです。
/sous-chef:simmer は目標に基づいた作業用で、コマンドが通過するまでループされます。
「スイートを環境に優しいものにする」、「200 ミリ秒未満のベンチマークを取得する」。毎周ごとに新しいコーデックスが実行され、
クロードは、専用ブランチ上で実際のコマンド出力を使用してラップごとに判断します。
キャップと進行なしの検出。労働者は自分で宿題を採点することはありません。
経験則: タスクを遂行し、目標を達成します。サーブが予算を使い果たした場合、
残ったものは目標の形をしており、煮込み続けることを提案します。
自分でステーションを運転したい場合のアラカルト:
要件: Codex CLI ≥ 0.134、
認証済み (コーデックス ログイン - ChatGPT サブスクリプションで十分です。API キーは必要ありません)。
/plugin Marketplace add tomascupr/sous-chef
/プラグインインストール sous-chef@sous-chef
( sous-chef@sous-chef は plugin@marketplace です - ここでは両方とも同じ名前です。) 次に、
リポジトリ内:
/副料理長:mise
分割の仕組み
あなた ── 「/sous-chef:serve maigration auth」 ──▶ CLAUDE (料理長)
│ チケット: ファイル ±、完了時、
│ 検証コマンド
▼
┌─────────────────────┐
│ codex exec --profile 副料理長 │ 背景;
│ ワークスペース書き込みサンドボックス · 承認オフ │ セッションメモリなし。
│ AGENTS.md を読み取る · チケットを実装する │ ハード境界

━━━━━━━━━━━━━━━━━━┘
▼ 差分
CLAUDE は検証自体をレビューし、再実行します
▼
CODEX クロスレビューは読み取り専用 ──▶ CLAUDE が調査結果を検証
▼
確認された所見は一度再発行された ──▶ 確認された ──▶ 提供された
ハードブロックではなくソフトルーティング。 CLAUDE.md のルーティング ポリシーとスキル
委任を最も抵抗の少ない道にする。クロードは今でも小規模なコンテンツを直接編集しています
外科的修正 - ハードブロッキングの編集/書き込みにより、エージェントが周囲を迂回することが証明されています。
代わりにブロックしてください。難しい境界線: 委任された Codex の実行は、
workspace-write サンドボックスは承認がオフになっており、Codex レビューは読み取り専用で実行されます。 (
オプションの GLM Claude-worker ルートには、その下に OS サンドボックスがありません。医師には治療するように書かれています
それに応じて、信頼できるリポジトリまたはブランチ/ワークツリーのみ)。
標準に関する信頼できる情報源の 1 つ。リポジトリ規約は AGENTS.md にあります。
Codex は、非インタラクティブな codex exec を含む実行ごとに再読み取りされます。クロードは読む
CLAUDE.md の @AGENTS.md インポートを介して同じファイルを作成します。タスクごとの指示の移動
チケットに記載されています。スタンディングオーダーはファイルに残ります。
バックグラウンドは常に、ポーリングは決して行われません。委任された実行は run_in_background 経由で実行されます
そのため、Bash タイムアウトの上限によりジョブの途中でそれらを強制終了することはできず、完了するとクロードが再び呼び出されます。
主張は証拠ではありません。委任された実行のたびに、クロードは差分ラインをレビューします
行ごとに検証コマンド自体を再実行します。
あらゆる負荷のかかる決定は、文書化されたインシデント、公式文書、または
測定された比較 - 雰囲気ではありません。サンプル:
バックグラウンドで常時実行される理由: 実行中の Codex ジョブに対する単一のポーリング ループが書き込まれる
約 12 時間で毎週のクロード割り当ての 27% が何も生成されない
( anthropics/claude-code#54143 )。
なぜブロックではなくソフトルーティングなのか

king 編集/書き込み: エージェント、によって 3 回ブロックされました
フック、Python ファイルでルーティング - Bash 経由で書き込み
( anthropics/claude-code#29709 )。
保持できないハード ブロックは、正直なルーティング ポリシーよりも悪いです。
調査結果が検証される理由: 20 レビューのフィールド テストで、Codex レビュー 20 件中最大 3 件
サイレントに失敗し、敵対モードではサーキット ブレーカーが見つからないというフラグが立てられました。
500行のcronスクリプト。
これらおよびその他すべての決定に関する完全なソース: docs/design.md 。
これは OpenAI の公式コーデックス プラグインとどう違うのですか? 3つの意図的な
相違点、それぞれ docs/design.md に領収書あり: (1) 停止時間なし
review Gate - OpenAI 自身の README は、「長期にわたるクロード/コーデックスを作成する可能性がある」と警告しています。
ループし、使用制限がすぐになくなる可能性があります。」ここで、パス内の実行をレビューします
明示的に注文され、ハードランの予算の下で、すべてのストップではありません。 (2) 所見取得
表示する前に実際のコードに対して検証 - 生のクロスモデルレビュー
オーバーフラグを設定し、検証で誤検知をフィルタリングします。 (3) すき間を埋める/煮る
公式プラグインも ralph-loop もカバーしていません。
ループの外側に独立した裁判官がいる。
これにはどのような費用がかかりますか? 2 つのサブスクリプション: Claude Code の任意の Claude プランと、
Codex 用の ChatGPT プラン - codex ログイン、API キーは必要ありません。サブスクリプション認証は、
ヘッドレス実行のファーストクラス パス: codex exec は保存されたログイン、トークンを再利用します。
実行中でも自動更新され、fire により 2 つの環境変数 ( CODEX_API_KEY 、
CODEX_ACCESS_TOKEN ) を使用すると、実行をトークンごとの課金にサイレントに切り替えることができます。
委任のオーバーヘッドは往復あたり最大 5 ～ 7,000 のクロード トークンであるため、タスクが小規模になります
クロードと一緒にいてください。
調理中に何が見えるでしょうか？最初に発表: 何が委任されたのか、
予想される所要時間 (通常、高度な推論努力で Codex を実行するたびに 5 ～ 20 分以上)、
そしてログ

パス。あなたは働き続けます。ジョブが終了すると、Claude が再度呼び出されます。キャンセル
いつでも - クロードはジョブを強制終了し、保持または元に戻す部分的な変更を表示します。
クロードはコードを書くのをやめますか?いいえ。小さな修正、プロトタイプ、その他何でも
設計が曖昧なため、Claude に留まります - ルーティング ルール自体がそう示しています。代表団
とアナウンスされますが、決して沈黙することはありません。
どのモデルですか? ~/.codex/config.toml の内容に関係なく、出荷されたプロファイル
サンドボックスと承認ポリシーのみを固定します。推奨: gpt-5.5 付き
model_reasoning_effort = "xhigh" 。 GLM-5.2 はオプトインの 2 番目の実装者として出荷されます
(「GLM で起動」): SWE-bench Pro で GPT-5.5 のベンチマークをわずかに上回ります。
ただし、トークンの消費量は最大 3.3 倍になります。テンプレートとしての 2 つのルート
(ヘッドレス Claude ワーカーを介した GLM コーディング プラン、または Codex を介した OpenRouter)。 /mise
持っているキーを設定します。クロード側では、モデルに依存しません。のために作られた
Fable 5 でドッグフーディングされています。
なぜMCPではないのでしょうか？ Bash 上でプレーンな codex を実行すると、サンドボックス フラグ、終了が表示されます。
コード、プロンプトの標準入力、および余分な可動部分のないバックグラウンド実行。つまり
副料理長が永続的な MCP サーバーの代わりに薄いラッパーを使用する理由。
ウィンドウズ？スニペットは POSIX です。 Claude Code の Git Bash では動作するはずです。
しかし、これは macOS ではドッグフード化されています。
スキル/サービス/自律パイプライン: 発射、試食、再発射、検証、報告
スキル/煮る/ループ: コーデックスは機能します、クロードは判断します、ゴールが通過するまで
スキル/ファイア/委任スキル + チケットテンプレート + GLM ルート
スキル/テイスト/クロスレビュー スキル + レビュー プロンプト テンプレート
スキル/再発射/修正スキル: 確認された結果は範囲指定された修正実行になります
スキル/ミセ/セットアップスキル
codex/ Codex プロファイル → ~/.codex/ (sous-chef デフォルト、sous-chef-glm)
templates/ AGENTS.md スキャフォールド、CLAUDE.md ルーティング ブロック、GLM ワーカー構成
docs/design.md 領収書

ts: あらゆる設計上の決定のソース
アンインストール
/plugin uninstall sous-chef はスキルを削除します (そして
/plugin Marketplace 副料理長の登録を削除します)。 /mise を実行した場合、
以下も作成しました (使い終わったら手動で削除してください)。
~/.codex/sous-chef.config.toml および ~/.codex/sous-chef-glm.config.toml
~/.sous-chef/glm-claude/ (分離された GLM ワーカー構成)
~/.claude/CLAUDE.md に追加される「分業 (副料理長)」ブロック
設定したリポジトリ内の AGENTS.md スキャフォールドおよび/または @AGENTS.md インポート行 (これら
これであなたのものになります - プラグインに関係なく便利です)
フィールドレポートは歓迎します - 特に Windows、特に矛盾する領収書
docs/design.md ;それは修正されることを意味します。
クロードは調整とレビューを行います。 Codex CLI を実装します。 2 つのモデルのキッチン用の Claude Code プラグイン: / タスクをエンドツーエンドで提供し、 / 目標を緑色になるまで煮込みます。あなたの料理長は玉ねぎを刻みません。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude orchestrates and reviews; Codex CLI implements. A Claude Code plugin for the two-model kitchen: /serve tasks end to end, /simmer goals until green. Your head chef doesn't chop onions. - tomascupr/sous-chef

GitHub - tomascupr/sous-chef: Claude orchestrates and reviews; Codex CLI implements. A Claude Code plugin for the two-model kitchen: /serve tasks end to end, /simmer goals until green. Your head chef doesn't chop onions. · GitHub
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
tomascupr
/
sous-chef
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .claude-plugin .claude-plugin .claude .claude codex codex docs docs skills skills templates templates .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Fable 5 orchestrates and reviews; GPT-5.5 xhigh implements. Your head chef doesn't chop onions.
A Claude Code plugin that splits coding between two frontier models the way a kitchen
splits work. Claude plans, writes the ticket, reviews every diff line by line, and
re-runs the checks itself. Codex CLI does the implementation - in a sandbox, in the
background, with no say over what ships. Spend Claude tokens on judgment and Codex
tokens on bulk: in the measured setup this pattern is built on,
Codex did ~20x the implementation work
per orchestration round trip, and two mid-tier subscriptions often beat one top-tier one.
> /sous-chef:serve migrate the auth module off the deprecated session API
Claude "Full serve at gpt-5.5: implement, cross-review, fix, verify.
2-3 Codex runs, ~15-45 min. I'll report when it's plated."
Claude writes the ticket - files to touch, files NOT to touch, done-when
criteria, verification commands - and fires it (background)
Codex implements. 11 files changed. Claims "all tests pass."
Claude reviews the diff line by line, re-runs pnpm test + tsc itself
Codex cross-reviews the diff read-only; 3 findings
Claude validates them against the code: 2 confirmed, 1 refuted
Claude refires the 2 confirmed findings as a scoped fix ticket
Claude re-verifies each finding at its cited location; checks green
Claude "Served: auth migration, 11 files, 42 tests pass, 2 review
findings fixed, 1 refuted. Two models agree; plate's yours."
One command in, one plated result out. Codex saying "tests pass" is a sentence;
pnpm test output is a fact - Claude re-runs everything itself.
/sous-chef:serve is for task-shaped work, done end to end: implement,
cross-review, fix the findings, verify. One announcement up front, one report at the
end, a hard budget of five Codex runs in between. This is the daily driver.
/sous-chef:simmer is for goal-shaped work, looped until a command passes:
"make the suite green", "get the benchmark under 200ms". A fresh Codex run each lap,
Claude judging every lap with real command output, on a dedicated branch, with lap
caps and no-progress detection. The worker never grades its own homework.
Rule of thumb: serve a task, simmer a goal. If a serve runs out of budget and
what remains is goal-shaped, it offers to continue as a simmer.
À la carte, when you want to drive the stations yourself:
Requirements: Codex CLI ≥ 0.134,
authenticated ( codex login - a ChatGPT subscription is enough; no API key needed).
/plugin marketplace add tomascupr/sous-chef
/plugin install sous-chef@sous-chef
( sous-chef@sous-chef is plugin@marketplace - same name for both here.) Then,
inside a repo:
/sous-chef:mise
How the split works
you ── "/sous-chef:serve migrate auth" ──▶ CLAUDE (head chef)
│ ticket: files ±, done-when,
│ verification commands
▼
┌────────────────────────────────────────────┐
│ codex exec --profile sous-chef │ background;
│ workspace-write sandbox · approvals off │ no session memory;
│ reads AGENTS.md · implements the ticket │ hard boundary
└───────────────────┬────────────────────────┘
▼ diff
CLAUDE reviews + re-runs verification itself
▼
CODEX cross-reviews read-only ──▶ CLAUDE validates findings
▼
confirmed findings refired once ──▶ verified ──▶ served
Soft routing, not hard blocks. A routing policy in CLAUDE.md plus skills that
make delegation the path of least resistance. Claude still edits directly for small
surgical fixes - hard-blocking Edit/Write provably makes agents route around the
block instead. The boundary that IS hard: delegated Codex runs execute in a
workspace-write sandbox with approvals off, and Codex reviews run read-only . (The
optional GLM Claude-worker route has no OS sandbox underneath; its docs say to treat
it accordingly - trusted repos or a branch/worktree only.)
One source of truth for standards. Repo conventions live in AGENTS.md , which
Codex re-reads on every run - including non-interactive codex exec . Claude reads
the same file via an @AGENTS.md import in CLAUDE.md . Per-task instructions travel
on the ticket; standing orders stay in the file.
Background always, polling never. Delegated runs execute via run_in_background
so the Bash timeout ceiling can't kill them mid-job, and completion re-invokes Claude.
Claims are not evidence. After every delegated run, Claude reviews the diff line
by line and re-runs the verification commands itself.
Every load-bearing decision traces to a documented incident, an official doc, or a
measured comparison - not vibes. A sample:
Why background-always: a single polling loop against a running Codex job burned
27% of a weekly Claude quota in ~12 hours producing nothing
( anthropics/claude-code#54143 ).
Why soft routing, not blocking Edit/Write: an agent, blocked three times by a
hook, routed around it with a Python file-write via Bash
( anthropics/claude-code#29709 ).
A hard block that can't hold is worse than an honest routing policy.
Why findings get validated: in a 20-review field test, ~3 of 20 Codex reviews
failed silently, and adversarial mode flagged missing circuit breakers on a
500-line cron script.
Full sources for these and every other decision: docs/design.md .
How is this different from OpenAI's official codex plugin? Three deliberate
divergences, each with receipts in docs/design.md : (1) no stop-time
review gate - OpenAI's own README warns it "can create a long-running Claude/Codex
loop and may drain usage limits quickly"; here, review runs inside a pass you
explicitly ordered, under a hard run budget, not on every stop. (2) findings get
validated against the actual code before you see them - raw cross-model reviews
over-flag, and validation filters the false positives. (3) /simmer fills a gap
neither the official plugin nor ralph-loop covers: a delegated implementer inside the
loop with an independent judge outside it.
What does this cost me? Two subscriptions: any Claude plan for Claude Code, and a
ChatGPT plan for Codex - codex login , no API key needed. Subscription auth is the
first-class path for headless runs: codex exec reuses the saved login, tokens
auto-refresh even mid-run, and fire unsets the two env vars ( CODEX_API_KEY ,
CODEX_ACCESS_TOKEN ) that could silently switch a run to per-token billing.
Delegation overhead is ~5-7k Claude tokens per round trip, which is why small tasks
stay with Claude.
What do I see while it cooks? An announcement first: what was delegated, the
expected duration (typically 5-20+ minutes per Codex run at high reasoning effort),
and the log path. You keep working; Claude is re-invoked when the job exits. Cancel
anytime - Claude kills the job and shows you any partial changes to keep or revert.
Does Claude stop writing code? No. Small fixes, prototypes, and anything
design-ambiguous stay with Claude - the routing rules themselves say so. Delegation
is announced, never silent.
Which models? Whatever your ~/.codex/config.toml says - the shipped profile
pins only sandbox and approval policy. Recommended: gpt-5.5 with
model_reasoning_effort = "xhigh" . GLM-5.2 ships as an opt-in second implementer
("fire with GLM"): it slightly out-benchmarks GPT-5.5 on SWE-bench Pro at a fraction
of the per-token price, though ~3.3x more token-hungry. Two routes as templates
(GLM Coding Plan via a headless Claude worker, or OpenRouter through Codex); /mise
sets up whichever key you have. On the Claude side it's model-agnostic; built for and
dogfooded with Fable 5.
Why not MCP? Plain codex exec over Bash gives you the sandbox flag, the exit
code, stdin for prompts, and background execution with no extra moving parts. That is
why sous-chef uses a thin wrapper instead of a persistent MCP server.
Windows? The snippets are POSIX; under Claude Code's Git Bash they should work,
but this is dogfooded on macOS.
skills/serve/ the autonomous pipeline: fire, taste, refire, verify, report
skills/simmer/ the loop: Codex works, Claude judges, until the goal passes
skills/fire/ delegation skill + ticket template + GLM routes
skills/taste/ cross-review skill + review prompt template
skills/refire/ fix skill: confirmed findings become a scoped fix run
skills/mise/ setup skill
codex/ Codex profiles → ~/.codex/ (sous-chef default, sous-chef-glm)
templates/ AGENTS.md scaffold, CLAUDE.md routing block, GLM worker config
docs/design.md the receipts: sources for every design decision
Uninstall
/plugin uninstall sous-chef removes the skills (and
/plugin marketplace remove sous-chef the registration). If you ran /mise , it may
also have created (remove by hand if you're done with them):
~/.codex/sous-chef.config.toml and ~/.codex/sous-chef-glm.config.toml
~/.sous-chef/glm-claude/ (isolated GLM worker config)
a "Division of labor (sous-chef)" block appended to ~/.claude/CLAUDE.md
an AGENTS.md scaffold and/or @AGENTS.md import line in repos you set up (these
are yours now - they're useful regardless of the plugin)
Field reports welcome - especially Windows, and especially receipts that contradict
docs/design.md ; it's meant to be corrected.
Claude orchestrates and reviews; Codex CLI implements. A Claude Code plugin for the two-model kitchen: /serve tasks end to end, /simmer goals until green. Your head chef doesn't chop onions.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
