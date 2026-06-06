---
source: "https://github.com/c0decave/peers"
hn_url: "https://news.ycombinator.com/item?id=48422239"
title: "New version of \"peers\" – the AI couple doing things"
article_title: "GitHub - c0decave/peers: Peers is a Claude+Codex framework, both working together to write, research and fix code. · GitHub"
author: "dash0r"
captured_at: "2026-06-06T07:10:44Z"
capture_tool: "hn-digest"
hn_id: 48422239
score: 1
comments: 0
posted_at: "2026-06-06T07:08:54Z"
tags:
  - hacker-news
  - translated
---

# New version of "peers" – the AI couple doing things

- HN: [48422239](https://news.ycombinator.com/item?id=48422239)
- Source: [github.com](https://github.com/c0decave/peers)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T07:08:54Z

## Translation

タイトル: 新バージョンの「peers」 – AI カップルが何かをする
記事のタイトル: GitHub - c0decave/peers: Peers は Claude+Codex フレームワークであり、両方が連携してコードの作成、調査、修正を行います。 · GitHub
説明: Peers は Claude+Codex フレームワークであり、両方が連携してコードの作成、調査、修正を行います。 - c0decave/ピア

記事本文:
GitHub - c0decave/peers: Peers は Claude+Codex フレームワークであり、両方が連携してコードの作成、調査、修正を行います。 · GitHub
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
c0decave
/
仲間
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション操作

ション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット auth-proxy auth-proxy docs docs proxy プロキシ scripts scripts src src testing testing .gitignore .gitignore Containerfile Containerfile Makefile Makefile README.md README.md README_DE.md README_DE.md compose.yaml compose.yaml pyproject.toml pyproject.toml pytest.ini pytest.ini すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントは 1 人より 2 人の方が優れています。それを証明させるのであれば。
ピアは、n ≥ 2 の AI コーディング CLI (Claude Code、Codex など) を連携として駆動します。
タスクが完了したことに同意するだけでなく、困難をクリアしなければならない仲間、
最初に測定可能なゲート: テストは合格、カバレッジは保持、回帰なし、なし
TODO/スタブ/スキップされたテスト、シークレットはクリーンです。一方のピアが実装し、もう一方のピアが
ブラインドレビュー（最初の人のメモを見ずに）、そして敵対的レビュー
「完了」が受け入れられる前に、懐疑的な再監査が行われます。無人で実行され、
予算制限付き 、およびコンテナーサンドボックス付き 。
ループ上の単一エージェントに勝る理由:
バイブベースではなく、ゲート付き。 「完了したように見える」は決して収束しない — ゲートは緑色 +
スケプティッククリーンはそうします。コンバージェンスシアターはありません。
ブラインド査読はゴム印を捕まえる - 独立した 2 番目のペア
目の構造による。
敵対的な懐疑論者は、テストで見逃されるエッジケースを探し出します。
無人&安全: アイドルタイムアウト監視、USD/ティックバジェットキャップ、
ルートレスキャップドロップコンテナ、出力許可リスト。
インストルメント化された診断では、ピアが式言語インタープリターを構築しました
グリーンフィールドとブラウンフィールドの両方を 50,000 回のランダム テストで欠陥ゼロに達成
プログラム — 植え付けられた回帰を捕捉し、エッジケースのバグを自己発見する
アクセプタンススイートは決して調査されませんでした。
ドイツ語バージョン: README_DE.md 。
HOWTO: 既存アプリの完全な監査と修正: docs/HOWTO-audit-and-fix.md — ドイツ語解説
実装モード (ビルド

d PLAN.md の機能) : docs/MODES_IMPLEMENT.md — DE
セキュリティ モデル: docs/SECURITY.md — DE
クイックスタート (無人、コントローラー経由)
パス A — 新しいプロジェクトから開始します (ワンショット)
Peers-ctl 新しい mything --modes=audit --spec ./mything-spec.md
$EDITOR ~ /c0de/peers-c0de/mything/.peers/goals.yaml # プロジェクト固有のゲートをトリミングする
Peers-ctl start mything --max-ticks 20 --max-usd 5
利用可能なモード: 「peers-ctl モード リスト」を参照してください。複数をスタックする
--modes=監査、徹底的 。現在の組み込みモード:
# 監査 + 完全 (既存のコードベースの推奨デフォルト):
Peers-ctl new myapp --modes=audit,through
# ベア監査:
ピア-ctl 新しい myapp --modes=audit
# 最初にドキュメントを作成し、後で監査します (2 回に分けて実行):
Peers-ctl new myapp --modes=describe # 1 を実行
Peers-ctl new myapp-audit --modes=audit,thorough # 2 を実行
# PLAN.md から機能を実装します (スタンドアロン — コンポーザブルではありません):
ピア-ctl 新しい myfeature --container --modes=implement --plan ./PLAN.md
# PLAN.md スキーマとエスケープバルブについては、docs/MODES_IMPLEMENT.md を参照してください。
自動フック (オプトアウト フラグ):
recon pre-tick (デフォルトはオン): サブストレートはティック 1 の前にリポジトリを 1 回スキャンし、.peers/recon.md (検出された言語、キードキュメント、エントリポイント候補、トップレベルツリー) を書き込みます。無料 + 高速 — LLM コールは不要です。 「ブラインド ティック 1」ペナルティを廃止します。オプトアウト:peers-ctl start <name> --without-recon 。
codemap pre-tick (デフォルトはオン): サブストレートは AST から構造的な CODEMAP を構築し、.peers/CODEMAP.yaml (機械可読: すべてのパブリック シンボル、そのファイル: 行と署名) と .peers/codemap.md (コンテキストとして読み取られるコンパクトでバイトキャップされたダイジェスト ピア) を書き込みます。無料 + 高速 — LLM コールは不要です。 Recon のファイル レベルのビューに加えて、ティック 1 の前にコードベースのパブリック API シェイプを使用してピアをプライムします。オプトアウト:peers-ctl start <name> --no-codemap 。
自動懐疑的な投稿

-convergence (デフォルトはオン):Continuous_clean_ticks >= N が convergence-reached を起動する場合、オーケストレーターは重要な再監査プロンプトとともに 1 つの追加のティックを実行します。懐疑的なダニがきれいなままであれば→本当に末期です。新しいブロックバグが表面化した場合→カウンターがリセットされ、ループが継続します。オプトアウト:peers-ctl start <name> --without-post-convergence-sketic 。
ディレクトリが見つからない場合は作成します (ディレクトリへのスキャフォールディングを拒否します)
--force を除く空でないディレクトリ);
裸の名前 ( / なし) は $PEERS_PROJECTS_ROOT の下に配置されます (デフォルト)
~/c0de/peers-c0de/<名前> 。 / を含むパスはそのまま解釈されます。
git init + 初期スキャフォールドコミット;
--force が使用されている場合でも、トップレベルの README.md が存在することを保証します。
既存の Git リポジトリに対して。
--spec 引数を SPEC.md にコピーします (既存のファイル パスは
読む; ./typo.md などのパスを探す欠損値は拒否されます)。
ピア init を実行します ( .peers/ を書き込み、ピアにタグを付けます)。
.gitignore をコミットし、 .peers/log/runs.jsonl を作成します);
--modes=audit を指定すると、6 つの監査チェック スクリプトと 1 つの監査チェック スクリプトがインストールされます。
監査対応の Goals.yaml ; --lang=js 、 --lang=rust 、または
--lang=go スタック固有のチェック エントリポイントの場合。
プロジェクトをpeers-ctlに登録し、コントローラログを作成します
Peers-ctl config ディレクトリの下にあります。
別のプロジェクト ルートを使用するには (例: プロジェクト固有のルート)
ディスク): PEERS_PROJECTS_ROOT=/work/peers/ を 1 回エクスポートし、その後は裸にします
名前がそこに着陸します。 Peers-ctl Doctor はアクティブなルートを出力します。
パス B — 独自の既存プロジェクトを持ち込む (最初の監査)
cd /パス/to/ターゲットプロジェクト
ピア init # .peers/ を書き込み + .gitignore をコミット
$EDITOR .peers/goals.yaml # `placeholder-replace-me` を削除し、実際のゲートを書き込みます
python3 - << ' PY '
ハッシュライブラリ、パスライブラリをインポート
p = pathlib.Path(".peers")
(p / "goals.sha256").write_text(hashlib.sha256((p / "goals.yaml").read_bytes()).hexdigest() + "\n")
ピィ
$編集

R .peers/config.yaml # コーデックスにカスタム argv パスが必要な場合のみ
ピア情報 # サニティチェック: ピア、目標、予算、健康状態
Peers-ctl add /path/to/your-target-project --name mything
Peers-ctl Doctor # ツール + プロジェクトごとの構成を確認します
Peers-ctl start mything --max-ticks 20 --max-usd 5
パス C — 既存のプロジェクトをさまざまなモードで再監査する
モードは、スキャフォールド時に .peers/goals.yaml にベイクされます。再実行するには
異なるモードが設定された同じプロジェクト (例: 最初に監査を実行した)
そして今は監査が必要です、一番上まで):
# バリアント 1: 適切な場所で再初期化します (破壊的 — Goals.yaml + チェックを上書きします)
Peers-ctl 新しい mything /path/to/your-project \
--modes=監査、徹底的な --force
# その後、通常どおりに開始します。
Peers-ctl start mything --container --max-ticks 30
# バリエーション 2: 個別のワークツリー (非破壊、推奨)
git -C /path/to/your-project ワークツリー add \
/path/to/your-project-through HEAD
Peers-ctl 新しい mything-thorough /path/to/your-project-thorough \
--container --modes=監査、徹底
Peers-ctl start --container mything-through
# 完了したら、実質的な修正を厳選してメインのワークツリーに戻します。
バリアント 2 は、反復監査に推奨されるパターンです。それぞれ
ワークツリーのクローンの監査を実行します。修正はマージによって厳選されます
レビュー後に --no-ff を付けます。ワークツリー パターンにより、既存の
監査履歴 ( .peers/state.json 、 .peers/log/runs.jsonl ) はそのままです。
ピア CTL ステータス神話 # スナップショット
Peers-ctl ダッシュボード # 登録されているすべてのプロジェクトを一度に
Peers-ctl ダッシュボード --live # アラート/イベントによる継続的な再描画
Peers-ctl ダッシュボード --project mything # ドリルダウン: 最近の実行 + バグ
Peers-ctl tail mything # live tail (切り離すには Ctrl-C)
tail -f /path/to/your-target-project/.peers/log/runs.jsonl # 豊富なティックごとの監査
ピア -C /path/to/your-target-project 再生 3 # 検査ティック 3

それが完了したら（または停止したい場合）
Peers-ctl stop mything #優雅な SIGTERM → 10s → SIGKILL
ピア -C /path/to/your-target-project レポート # .peers/REPORT.md を書き込みます
Peers-ctl report mything # コントローラー REPORT-mything.md を書き込みます
Peers-ctl レビューの神話 # 最新のハンドオフのセルフレビュー
CI ガードレールは .gitea/workflows/test.yml として利用可能です。
スクリプト/pre-push.sh ; makehooks-install を使用してローカルフックをインストールします。
コントローラーはステートレスです。プロジェクト独自の .peers/state.json
そして、runs.jsonl は永続的なレコードです。ホストが再起動した場合
実行中に、peers-ctl リストはプロジェクトがクラッシュしたことをマークします。できます
もう一度開始すると、ループは保存された反復から再開されます。
プロジェクトの状態は、peers-ctl list によって表示されます。
モードは、監査目標とチェック スクリプトの再利用可能なバンドルです。
Peers-ctl new --modes=… は .peers/ にあります。モードは次のとおりです。
stackable (カンマ区切りリスト) — ただし、 description を除きます。
監査/セキュリティ モードとは相互に排他的です (ドキュメントを書き込みますが、
コードを監査します)。
監査 (基礎 - ほとんどの場合必須)
ハードゲート: ハンドオフ時のセルフレビュー、テスト合格、
テスト-カバー-ハッピーエッジ-悲しい 、テスト-不当なスキップ-または失敗
(ピアはすべての @pytest.mark.skip / xfail を正当化する必要があります)、
lint-clean 、 type-clean 、 bug-hunt-clean 、 tdd-reproductions-bug 、
no-secrets-committed 、 deps-justified 、 api-stable 、
事前回帰なし、解決ごとの差分サイズ。
ソフト目標: bug-hunt-round-1-deep 、 bug-hunt-round-2-cross-review 、
テスト-3-クラス-レビュー 。
常に使用してください。他のモードは、 Audit のハードゲートがアクティブであると想定します。
そして「クリーン」の意味をしっかりと理解してください。
徹底的（監査の上に積み重ねる）
収束到達 (ハード、N=3 デフォルト): N 連続クリーン
新しいクリティカル/高/中程度のバグレポートのないティック — 基板が拒否する
N 個の静止の証明なしで成功を宣言する。
all-peers-healthy (ハード): d を拒否します。

成功を宣言します
ピアは利用できない状態です (停止パターン ヒット)。
sketic-pass (ソフト、両方のピア、間隔 1): すべてのティックを再監査します
余計な疑いを持って。 5+を文書化せずに通過を拒否する
ファイル/モジュールごとに除外される障害モード。
アグレッシブ正直 (ソフト、両方のピア、間隔 3): src ごと
最上位パス: 3 つ以上の障害モードがチェックされ、2 つ以上のセキュリティ カテゴリ、
1 つのテスト カバレッジ ギャップが明示的に指定されています。
徹底だけでは (監査なしで) 不完全です。収束に達したかどうかを知るには (監査からの) バグハント - クリーンに依存します。
「きれい」という意味です。常に監査とスタックします: --modes=audit,through 。
説明する (ドキュメントを作成し、監査はしない)
ピアはプロジェクトの仕様ドキュメント (SPEC.md + ARCHITECTURE.md +) を作成します。
DESIGN.md) N=2 連続の非実体ドキュメントまで反復
コミットします。ハードゲート:
description-files-present : 3 つのファイルがすべて存在し、それぞれ ≥500 文字
description-sections-present : SPEC には ## 脅威モデル + があります
## 不変条件 + ## API ; ARCH には ## コンポーネント + ## データ フローがあります。デザインには ## 決定 + ## トレードオフがあります。各セクション
本文 ≥50 文字
description-converged : 3 つのファイルに対する最後の N コミットは非
実質的 (新しい ## セクションなし、追加行 100 行未満、削除 50% 未満)
監査モードでは構成できません - 書き込み、監査攻撃を記述します。
次のリポジトリで --modes=describe FIRST を実行します。

[切り捨てられた]

## Original Extract

Peers is a Claude+Codex framework, both working together to write, research and fix code. - c0decave/peers

GitHub - c0decave/peers: Peers is a Claude+Codex framework, both working together to write, research and fix code. · GitHub
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
c0decave
/
peers
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits auth-proxy auth-proxy docs docs proxy proxy scripts scripts src src tests tests .gitignore .gitignore Containerfile Containerfile Makefile Makefile README.md README.md README_DE.md README_DE.md compose.yaml compose.yaml pyproject.toml pyproject.toml pytest.ini pytest.ini View all files Repository files navigation
Two AI coding agents are better than one — if you make them prove it.
peers drives n ≥ 2 AI coding CLIs (Claude Code, Codex, …) as cooperating
peers that don't just agree a task is done — they have to clear hard,
measurable gates first: tests pass, coverage holds, no regression, no
TODO/stub/skipped-test, secrets clean. One peer implements, the other
blind-reviews (without seeing the first's notes), and an adversarial
skeptic re-audits before any "done" is accepted. Runs unattended ,
budget-capped , and container-sandboxed .
Why it beats a single agent on a loop:
Gated, not vibes-based. "Looks done" never converges — gates green +
skeptic-clean does. No convergence theater.
Blind peer review catches rubber-stamping — an independent second pair
of eyes, by construction.
An adversarial skeptic hunts the edge cases your tests miss.
Unattended & safe: idle-timeout supervision, USD/tick budget caps,
rootless cap-dropped container, egress allow-listing.
In an instrumented diagnostic, peers built an expression-language interpreter
both greenfield and brownfield to 0 defects over 50,000 random test
programs — catching planted regressions and self-finding edge-case bugs the
acceptance suite never probed.
Deutsche Version: README_DE.md .
HOWTO: full audit + fix on an existing app : docs/HOWTO-audit-and-fix.md — deutsche Anleitung
implement mode (build a feature from PLAN.md) : docs/MODES_IMPLEMENT.md — DE
Security model: docs/SECURITY.md — DE
Quickstart (unattended, via the controller)
Path A — start from a fresh project (one shot)
peers-ctl new mything --modes=audit --spec ./mything-spec.md
$EDITOR ~ /c0de/peers-c0de/mything/.peers/goals.yaml # trim project-specific gates
peers-ctl start mything --max-ticks 20 --max-usd 5
Available modes: see peers-ctl modes list . Stack multiple with
--modes=audit,thorough . Current built-in modes:
# audit + thorough (recommended default for an existing codebase):
peers-ctl new myapp --modes=audit,thorough
# bare audit:
peers-ctl new myapp --modes=audit
# write docs first, audit later (two separate runs):
peers-ctl new myapp --modes=describe # run 1
peers-ctl new myapp-audit --modes=audit,thorough # run 2
# implement a feature from a PLAN.md (standalone — not composable):
peers-ctl new myfeature --container --modes=implement --plan ./PLAN.md
# see docs/MODES_IMPLEMENT.md for the PLAN.md schema + escape valves.
Automatic hooks (opt-out flags):
recon pre-tick (default on): substrate scans the repo once before tick 1 and writes .peers/recon.md (detected languages, key docs, entry-point candidates, top-level tree). Free + fast — no LLM call. Eliminates the "blind tick 1" penalty. Opt out: peers-ctl start <name> --without-recon .
codemap pre-tick (default on): substrate builds a structural CODEMAP from the AST and writes .peers/CODEMAP.yaml (machine-readable: every public symbol, its file:line and signature) plus .peers/codemap.md (a compact, byte-capped digest peers read as context). Free + fast — no LLM call. Primes peers with the codebase's public-API shape before tick 1, on top of recon's file-level view. Opt out: peers-ctl start <name> --no-codemap .
auto-skeptic post-convergence (default on): when consecutive_clean_ticks >= N would fire convergence-reached , the orchestrator runs ONE extra tick with a critical re-audit prompt. If the skeptic-tick stays clean → really terminal. If it surfaces a new blocking bug → counter resets, loop continues. Opt out: peers-ctl start <name> --without-post-convergence-skeptic .
creates the directory if missing (refuses to scaffold into a
non-empty dir unless --force );
bare name (no / ) lands under $PEERS_PROJECTS_ROOT , default
~/c0de/peers-c0de/<name> . Path with / is taken verbatim;
git init + initial scaffold commit;
ensures a top-level README.md exists, even when --force is used
against an existing Git repo;
copies the --spec argument to SPEC.md (existing file paths are
read; path-looking missing values such as ./typo.md are rejected);
runs peers init (which writes .peers/ , tags peers-baseline ,
commits .gitignore , and creates .peers/log/runs.jsonl );
with --modes=audit , installs six audit check scripts and an
audit-ready goals.yaml ; use --lang=js , --lang=rust , or
--lang=go for stack-specific check entrypoints;
registers the project with peers-ctl and creates the controller log
under the peers-ctl config directory.
To use a different projects root (e.g. on a project-specific
disk): export PEERS_PROJECTS_ROOT=/work/peers/ once, then bare
names land there. peers-ctl doctor prints the active root.
Path B — bring your own existing project (first audit)
cd /path/to/your-target-project
peers init # writes .peers/ + commits .gitignore
$EDITOR .peers/goals.yaml # delete `placeholder-replace-me`, write real gates
python3 - << ' PY '
import hashlib, pathlib
p = pathlib.Path(".peers")
(p / "goals.sha256").write_text(hashlib.sha256((p / "goals.yaml").read_bytes()).hexdigest() + "\n")
PY
$EDITOR .peers/config.yaml # only if codex needs a custom argv path
peers info # sanity-check: peers, goals, budget, health
peers-ctl add /path/to/your-target-project --name mything
peers-ctl doctor # confirms tooling + per-project config
peers-ctl start mything --max-ticks 20 --max-usd 5
Path C — re-audit an existing project with different modes
Modes are baked into .peers/goals.yaml at scaffold-time. To re-run
the SAME project with a DIFFERENT mode set (e.g. you ran audit first
and now want audit,thorough on top):
# Variant 1: re-init in place (DESTRUCTIVE — overwrites goals.yaml + checks)
peers-ctl new mything /path/to/your-project \
--modes=audit,thorough --force
# Then start as usual:
peers-ctl start mything --container --max-ticks 30
# Variant 2: separate worktree (NON-DESTRUCTIVE, recommended)
git -C /path/to/your-project worktree add \
/path/to/your-project-thorough HEAD
peers-ctl new mything-thorough /path/to/your-project-thorough \
--container --modes=audit,thorough
peers-ctl start --container mything-thorough
# Cherry-pick the substantive fixes back to your main worktree when done.
Variant 2 is the recommended pattern for iterative audits. Each
run audits a worktree clone; fixes are cherry-picked back via merge
with --no-ff after review. The worktree pattern keeps your existing
audit history ( .peers/state.json , .peers/log/runs.jsonl ) intact.
peers-ctl status mything # snapshot
peers-ctl dashboard # all registered projects at once
peers-ctl dashboard --live # continuous redraw with alerts/events
peers-ctl dashboard --project mything # drilldown: recent runs + bugs
peers-ctl tail mything # live tail (Ctrl-C to detach)
tail -f /path/to/your-target-project/.peers/log/runs.jsonl # rich per-tick audit
peers -C /path/to/your-target-project replay 3 # inspect tick 3
When it's done (or you want to stop)
peers-ctl stop mything # graceful SIGTERM → 10s → SIGKILL
peers -C /path/to/your-target-project report # writes .peers/REPORT.md
peers-ctl report mything # writes controller REPORT-mything.md
peers-ctl review mything # latest handoff self-review
CI guardrails are available as .gitea/workflows/test.yml plus
scripts/pre-push.sh ; install the local hook with make hooks-install .
The controller is stateless; the project's own .peers/state.json
and runs.jsonl are the durable record. If the host reboots
mid-run, peers-ctl list will mark the project crashed ; you can
start it again and the loop resumes from the saved iteration.
Project states shown by peers-ctl list :
A mode is a reusable bundle of audit goals + check scripts that
peers-ctl new --modes=… lays down in .peers/ . Modes are
stackable (comma-separated list) — except describe , which is
mutually exclusive with audit/security modes (it writes docs, not
audits code).
audit (foundation — almost always required)
Hard gates: self-review-on-handoff , tests-pass ,
tests-cover-happy-edge-sad , tests-no-unjustified-skip-or-fail
(peers must justify every @pytest.mark.skip / xfail ) ,
lint-clean , type-clean , bug-hunt-clean , tdd-reproduces-bug ,
no-secrets-committed , deps-justified , api-stable ,
no-prior-regression , diff-size-per-resolve .
Soft goals: bug-hunt-round-1-deep , bug-hunt-round-2-cross-review ,
tests-3-class-review .
Use it always. Other modes assume audit 's hard-gates are active
and tighten what „clean" means.
thorough (stacks ON TOP of audit)
convergence-reached (hard, N=3 default): N consecutive clean
ticks without new crit/high/med bug-reports — the substrate refuses
to declare success without N proofs of stillness.
all-peers-healthy (hard): refuses to declare success while any
peer is in unavailable state (halt-pattern hit).
skeptic-pass (soft, both peers, interval 1): every tick re-audits
with extra suspicion; refuses to pass without documenting 5+
failure modes excluded per file/module.
aggressive-honesty (soft, both peers, interval 3): per src
top-level path: 3+ failure modes checked, 2+ security categories,
1 test-coverage gap explicitly named.
thorough alone (without audit ) is incomplete — convergence- reached depends on bug-hunt-clean (from audit) to know what
„clean" means. Always stack with audit: --modes=audit,thorough .
describe (write docs, don't audit)
Peers WRITE the project's spec docs (SPEC.md + ARCHITECTURE.md +
DESIGN.md) iteratively until N=2 consecutive non-substantive doc
commits. Hard gates:
description-files-present : all 3 files exist, ≥500 chars each
description-sections-present : SPEC has ## Threat Model +
## Invariants + ## API ; ARCH has ## Components + ## Data Flow ; DESIGN has ## Decisions + ## Tradeoffs ; each section
body ≥50 chars
description-converged : last N commits to the 3 files are non-
substantive (no new ## section, <100 lines added, <50% deletion)
Not composable with audit modes — describe writes, audit attacks.
Run --modes=describe FIRST on a repo that

[truncated]
