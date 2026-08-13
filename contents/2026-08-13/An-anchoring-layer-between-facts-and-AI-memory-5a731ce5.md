---
source: "https://github.com/Anchorstate-Lab/GMR"
hn_url: "https://news.ycombinator.com/item?id=49292758"
title: "An anchoring layer between facts and AI memory"
article_title: "GitHub - Anchorstate-Lab/GMR: Grounded Memory Runtime — a runtime tool that links the Agent’s memory and facts · GitHub"
author: "Zongming"
captured_at: "2026-08-13T23:31:51Z"
capture_tool: "hn-digest"
hn_id: 49292758
score: 1
comments: 1
posted_at: "2026-08-13T22:52:33Z"
tags:
  - hacker-news
  - translated
---

# An anchoring layer between facts and AI memory

- HN: [49292758](https://news.ycombinator.com/item?id=49292758)
- Source: [github.com](https://github.com/Anchorstate-Lab/GMR)
- Score: 1
- Comments: 1
- Posted: 2026-08-13T22:52:33Z

## Translation

タイトル: 事実と AI メモリの間の固定層
記事のタイトル: GitHub - Anchorstate-Lab/GMR: Grounded Memory Runtime — エージェントの記憶と事実をリンクするランタイム ツール · GitHub
説明: Grounded Memory Runtime — エージェントの記憶と事実をリンクするランタイム ツール - Anchorstate-Lab/GMR

記事本文:
GitHub - Anchorstate-Lab/GMR: グラウンデッド メモリ ランタイム — エージェントの記憶と事実をリンクするランタイム ツール · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
アンカーステートラボ
/
GMR
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
218 コミット 218 コミット .anchor .anchor .github/ workflows .github/ workflows バッテリー バッテリー クレート クレート

dist dist docs docs ドメイン/ コーディング ドメイン/ コーディング メモリ メモリ ツール ツール .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md accepts.sh accepts アーキテクチャ.toml アーキテクチャ.toml デモ.sh デモ.sh ゲート.sh ゲート.sh release-plz.toml release-plz.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
主観的な判断をリポジトリ内の再計算可能な観察に固定することで追跡します。
GMR は、メモを事実にバインドし、観察を再生するための軽量ランタイムです。
それらを生み出したものであり、世界があなたの方向に沿って動くとき、あなたに警告します。
と宣言した。
GMR は、次の作業に役立つ CLI ツールです。
リポジトリの観察可能な部分に判決またはメモを添付する
その判断を再現可能な観察に結びつけておく
関連する世界の状態が変化したときを検出する
移動したアンカーにバインドされたノートを返す
これはリンターやルール エンジンではありません。メモをアンカーするためのランタイムです。
観察可能な状態とそのバインディングを時間の経過とともに維持します。
手動の判断に責任を持たせたい場合は、GMR を使用します。
リファクタリング後も契約は保持される必要があります
コード変更後も動作上の仮定が真である必要がある
依存するものが移動すると、メモは再び表示されるはずです
GMR は、判断を依存する事実に付随させ、それが事実であるかどうかを報告します。
事実が変わる。
npm (エンドユーザーに推奨)
npm パッケージが公開されている場合は、ラッパー パッケージをインストールします。
npm install -g @zongming_he/gmr
npm ラッパーは、利用可能な場合、一致するプラットフォーム バンドルをロードします。
ご使用のオペレーティング システムまたはアーキテクチャが公開されているバージョンでサポートされていない場合は、
バンドルの場合、ラッパーはフォールバック メッセージを出力し、直接の
インストールスクリプト。
リポジトリには、ビルド済みリリース用の簡単なインストール スクリプトが含まれています。
もしあなたがw

npm を使用せずに直接バイナリをインストールすることもできます。
カール -fsSL https://raw.githubusercontent.com/Anchorstate-Lab/GMR/main/dist/install.sh |しー
フォークまたはミラーを使用している場合は、Anchorstate-Lab/GMR をリポジトリ パスに置き換えます。
このリポジトリから CLI を自分で構築する場合は、次のようにします。
カーゴインストール --path ドメイン/コーディング/cli --locked --root ~ /.local
これにより、現在のワークスペースから構築された、組み立てられた gmr CLI がインストールされます。
GMR は、Rust が含まれていないリポジトリを含むプロジェクト ディレクトリに対して動作します。
--repo <path> を指定してコマンドを実行します (デフォルトは . )。
gmr --repo /path/to/project init
これにより、ローカルの .anchor/ レイアウトが作成され、組み込みのプローブ バンドルがインストールされます。
そして、(初回のみ) クロード コードのスキル ドキュメントを書き込みます。
.claude/skills/gmr/SKILL.md — --global を渡して書き込みます
代わりに ~/.claude/skills/gmr/SKILL.md を使用します。アンカーやメモは作成されません
あなたのために。安全に再実行できます。すでに存在するファイルを上書きすることはありません。
2. 座標を固定し、メモリを書き込みます
gmr --repo /パス/プロジェクト アンカー src/auth.ts#createSession \
-m "セッションは 30 分後に期限切れになります。理由は..."
アンカーは座標をプローブ、形状、位置にルーティングし、
記憶の下のメモ/その記憶を使用して、アンカーを開き、メモをバインドします
すべてをワンステップで実現します。 -m を省略すると、メモは自分で書き込めるようになります。
あなたが書くまでは書かれていないものとして報告されます。同様に、次のように手書きすることもできます。
同じフロントマターを使用していることに注意し、gmr sync を実行してそれを開きます。
---
について: src/auth.ts#createSession
---
# セッションはサービス境界内で作成する必要があります
座標を指定せずに gmr アンカーを実行して、宣言や内容を開きます。
メモはすでに求めています - ジャーナルには必要ないため、新しいクローンに必要なもの
リポジトリと一緒に移動します。
gmr --repo /path/to/project check --json
チェックは適切なアンカーを評価します

メモで質問された軸がある場合は、rs および終了 1
移動し、移動したものに綴じられたメモを返します。アンカーにもフラグを立てます
その宣言がライブ基準とアンカーに一致しなくなった
別のプローブ機器が取得した読み取り値に立つ - 問題を解決する
(後述の accept --criteria と rebase を参照)、静かな結果を信頼する前に。
gmr --repo /path/to/project accept src/auth.ts#createSession --why " ... "
見てみると、チェックで示されたのは新しいベースラインです。受け入れるとクリアされます
ベクトルを作成し、その理由をジャーナルに永久に封印します。宣言の場合
変更も保留中です。 --baseline または --criteria を明示的に渡します。
別々の判断をし、同じ理由を共有しないでください。
5. 監視されているものすべてを確認する
gmr --repo /パス/to/プロジェクトステータス --json
status は、すべてのアンカー、その軸、およびそれにバインドされているノートを報告します。読み取り専用。
正面玄関 — 6 つの動詞 gmr --help は次を示します。
init — .anchor/ をセットアップし、バンドルされたプローブをインストールし、スキルドキュメントを作成します
アンカー — 座標を観察し、それに伴う記憶を書き込む
ステータス — 何が、どの軸で、どの記憶とともに見られているか
チェックしてください - 記憶が尋ねた軸上で何かが動きましたか?
accept — アンカーが現在表示しているものを新しいベースラインとして採用するか、
変更された宣言の基準 ( --baseline / --criteria / --all --criteria )
close — アンカーを永久に廃止します
他のすべては引き続き機能し、 gmr help <name> を通じてアクセスできます。
プローブ リスト / プローブ ビルド — 利用可能なプローブをリストまたはビルドします
sync — ノートによって宣言されたアンカーを開き、バインディングを整列させます（どのアンカー
座標実行なし)
open — アンカーを手動で直接作成します
観察 / 渡す / 読み取り — アンカーを評価し、移動したノートを返す、または
生のアンカー状態を読み取る
再プローブ / 再遷移 / リターミナル / リベース / 再ステート — ハンドドライブ
アンカー基準の一部

a;それぞれに --why が必要であり、ジャーナルに封印されています
binding / reaffirm / cobound — 参照バインディングを管理する
link — 参照間の関係を記録する
エッジ — ある時点以降のジャーナル遷移の読み取り
ヘルス - アンカーの活性度を検査する
ドクター - アンカーは観察されなかった、メモなし、解決不能、またはメモなし
lint の問題あり (未請求、ベアキー、ロングハンド、引退)
requeue — アンカーを強制的にキューに戻します
エクスポート / インポート — ストアのコンテンツのスナップショットと再生
ほとんどのコマンドは --repo <path> と --json をサポートしています。
このリポジトリで作業していて、それを構築または検証したい場合は、次のようにします。
カーゴビルド --release -p コーディングアンカー
貨物テスト --ワークスペース
sh ゲート.sh
sh 受け入れ.sh
ドキュメント
docs/GMR.md — アーキテクチャとデザインの信頼できる情報源
docs/architect.md — リポジトリの階層化とパッケージの責任
CLAUDE.md — 設計上の決定とリポジトリの基準
思い出/ — リポジトリのメモと根拠のある記録
グラウンデッド メモリ ランタイム — エージェントの記憶と事実をリンクするランタイム ツール
Readme MIT ライセンス アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Grounded Memory Runtime — a runtime tool that links the Agent’s memory and facts - Anchorstate-Lab/GMR

GitHub - Anchorstate-Lab/GMR: Grounded Memory Runtime — a runtime tool that links the Agent’s memory and facts · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Anchorstate-Lab
/
GMR
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
218 Commits 218 Commits .anchor .anchor .github/ workflows .github/ workflows batteries batteries crates crates dist dist docs docs domains/ coding domains/ coding memories memories tools tools .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md acceptance.sh acceptance.sh architecture.toml architecture.toml demo.sh demo.sh gate.sh gate.sh release-plz.toml release-plz.toml View all files Repository files navigation
Track subjective judgments by anchoring them to recomputable observations in your repository.
GMR is a lightweight runtime for binding notes to facts, replaying the observation
that generated them, and warning you when the world moves along a direction you
declared.
GMR is a CLI tool that helps you:
attach a judgment or note to an observable part of a repository
keep that judgment bound to a reproducible observation
detect when the relevant world state has changed
return the notes bound to anchors that moved
It is not a linter or rule engine. It is a runtime for anchoring notes to
observable state and preserving that binding over time.
Use GMR when you want to make a manual judgment accountable:
a contract should still hold after a refactor
a behavioural assumption should still be true after code changes
a note should be surfaced again when the thing it depends on moves
GMR keeps the judgment attached to the fact it depends on and reports if that
fact changes.
npm (recommended for end users)
If the npm packages are published, install the wrapper package:
npm install -g @zongming_he/gmr
The npm wrapper will load the matching platform bundle, if available.
If your operating system or architecture is not supported by the published
bundle, the wrapper prints a fallback message and points you to the direct
installation script.
The repository includes a simple install script for prebuilt releases.
If you want a direct binary install without npm:
curl -fsSL https://raw.githubusercontent.com/Anchorstate-Lab/GMR/main/dist/install.sh | sh
If you are using a fork or mirror, replace Anchorstate-Lab/GMR with your repository path.
If you want to build the CLI yourself from this repository:
cargo install --path domains/coding/cli --locked --root ~ /.local
This installs the assembled gmr CLI built from the current workspace.
GMR works against a project directory, including repositories with no Rust in them.
Run commands with --repo <path> (default is . ).
gmr --repo /path/to/project init
This creates the local .anchor/ layout, installs the built-in probe bundle,
and (the first time only) writes a Claude Code skill doc to
.claude/skills/gmr/SKILL.md — pass --global to write it to
~/.claude/skills/gmr/SKILL.md instead. It does not create anchors or notes
for you. Safe to rerun; it never overwrites a file that already exists.
2. Anchor a coordinate and write the memory
gmr --repo /path/to/project anchor src/auth.ts#createSession \
-m " sessions expire after 30 minutes because ... "
anchor routes the coordinate to a probe, a shape and a position, writes a
note under memories/ with that memory, opens the anchor, and binds the note
to it — all in one step. Omit -m and the note is left for you to write,
reported as unwritten until you do. Equivalently, you can hand-write the
note yourself with the same frontmatter and run gmr sync to open it:
---
about : src/auth.ts#createSession
---
# Sessions must still be created inside the service boundary
Run gmr anchor with no coordinate to open whatever the declarations and
notes already ask for — what a fresh clone needs, since the journal doesn't
travel with the repository.
gmr --repo /path/to/project check --json
check evaluates due anchors and exits 1 if any axis a note asked about
moved, handing back the notes bound to what moved. It also flags anchors
whose declaration no longer matches their live criteria, and anchors
standing on a reading a different probe instrument took — resolve those
(see accept --criteria and rebase below) before trusting a quiet result.
gmr --repo /path/to/project accept src/auth.ts#createSession --why " ... "
You looked, and what check showed is the new baseline; accept clears the
vector and seals the reason permanently into the journal. If a declaration
change is pending too, pass --baseline or --criteria explicitly — they're
separate judgments and don't share one reason.
5. See everything being watched
gmr --repo /path/to/project status --json
status reports every anchor, its axes, and the notes bound to it. Reads only.
The front door — six verbs gmr --help shows:
init — set up .anchor/ , install bundled probes, write the skill doc
anchor — watch a coordinate and write the memory that goes with it
status — what is being watched, on which axes, with which memories
check — did anything move on an axis a memory asked about?
accept — take what an anchor now shows as the new baseline, or take a
changed declaration's criteria ( --baseline / --criteria / --all --criteria )
close — retire an anchor permanently
Everything else still works, reachable through gmr help <name> :
probes list / probes build — list or build available probes
sync — open anchors declared by notes and align bindings (what anchor
with no coordinate runs)
open — create an anchor directly by hand
observe / pass / read — evaluate due anchors, return moved notes, or
read raw anchor state
reprobe / retransition / reterminal / rebase / restate — hand-drive
one part of an anchor's criteria; each needs --why , sealed into the journal
bind / reaffirm / cobound — manage reference bindings
link — record a relationship between references
edges — read journal transitions since a point
health — inspect anchor liveness
doctor — anchors never observed, with no note, unresolvable, or notes
with lint problems ( unclaimed , bare-key , long-hand , retired )
requeue — force an anchor back onto the due queue
export / import — snapshot and replay store contents
Most commands support --repo <path> and --json .
If you are working in this repository and want to build or verify it:
cargo build --release -p coding-anchor
cargo test --workspace
sh gate.sh
sh acceptance.sh
Documentation
docs/GMR.md — architecture and design source of truth
docs/architect.md — repository layering and package responsibilities
CLAUDE.md — design decisions and repository norms
memories/ — repository notes and grounded records
Grounded Memory Runtime — a runtime tool that links the Agent’s memory and facts
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
