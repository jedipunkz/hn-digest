---
source: "https://github.com/dylanp12/proctor"
hn_url: "https://news.ycombinator.com/item?id=48650355"
title: "Show HN: Proctor – signed isolation bundles for AI coding-agent benchmarks"
article_title: "GitHub - dylanp12/proctor: Tamper-proof execution sandbox for trustworthy AI coding-agent benchmarks · GitHub"
author: "dp12"
captured_at: "2026-06-23T20:42:15Z"
capture_tool: "hn-digest"
hn_id: 48650355
score: 2
comments: 0
posted_at: "2026-06-23T19:48:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Proctor – signed isolation bundles for AI coding-agent benchmarks

- HN: [48650355](https://news.ycombinator.com/item?id=48650355)
- Source: [github.com](https://github.com/dylanp12/proctor)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T19:48:28Z

## Translation

タイトル: Show HN: Proctor – AI コーディング エージェント ベンチマーク用の署名付き分離バンドル
記事タイトル: GitHub - dylanp12/proctor: 信頼できる AI コーディング エージェント ベンチマークのための改ざん防止実行サンドボックス · GitHub
説明: 信頼できる AI コーディング エージェント ベンチマークのための改ざん防止実行サンドボックス - dylanp12/proctor

記事本文:
GitHub - dylanp12/proctor: 信頼できる AI コーディング エージェント ベンチマークのための改ざん防止実行サンドボックス · GitHub
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
ディランプ12
/
監督官
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メートル

ain ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
80 コミット 80 コミット .cargo .cargo .github .github corpus corpus crates crates docs docs scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md action.yml action.yml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Proctor は、AI コーディング エージェントのベンチマーク実行を署名付きで独立して検証可能なものに変換します
整合性バンドル。
回答が分離された Linux サンドボックスでエージェントを実行します。そこでは、構成された非表示のテスト、修正が行われます。
履歴、およびネットワーク出口に到達できない場合は、評決と対象となる文書に署名します。
アクセス禁止のタイムライン。
AI コーディング エージェントのベンチマークが試されています。 2026 年 4 月、UPenn の研究者 (スタイン、
ブラウン、ハッサニ、ナイク、ウォン）は、人気エージェントに対する広範な不正行為を文書化した
ベンチマーク : 1,000 以上のハーネスレベルの不正行為
主要なベンチマーク全体のトレース — Terminal-Bench 2 と HAL USACO に集中 — さらに
～ 30 件のタスク レベルのケース (メソッド ペーパー: 多数のエージェントにわたる安全性違反の検出)
トレース、arXiv 2604.11806 )。チートはありふれたものです:
エージェントはテストオラクルを読み取ります (削除された Terminal-Bench 2 送信の 1 つ、415/429)
成功したトレースは /tests のプレーンなファイルシステム読み取りでした)、修正のために Git ログを調べました
コミットするか、ソリューションをカールするか、採点者の報酬ファイルを事前に書き込みます。これらのそれぞれは、
モデリングの問題ではなく、サンドボックス/アクセス制御の失敗。
この研究で最も劇的な数字 - 上位の投稿が 1 位から 14 位に転落
一度チートが解除されると、別のチート クラスからのものです。回答キーは、
エージェントの足場 ( AGENTS.md )。これはサンドボックスの外から届くため、OS が分離されます。

に
単独ではそれを止めることはできません。これは Proctor の v0.2 来歴作業のターゲットです。
私たちはその境界に、
静かに主張しています。
Proctor は、強制的な OS レベルの分離の下でベンチマーク タスクを実行します。
評価者のアーティファクト (Oracle/テスト、修正履歴、ネットワーク) はエージェントのからアクセスできません。
サンドボックスを使用し、署名済みの評決と、対象となる禁止されたアクセスの改ざん明示ログを出力します。
試行（マスクされたファイルの読み取り、ブロックされた出力）。
プロクター run --task ./task --agent "my-agent --solve" --policy ./policy.yaml
# -> verdict.json { "pass": true, "status": "compromized", "violations_count": 1, ... }
# -> Violations.jsonl { "ステップ": 8, "種類": "masked_read", "パス": "/oracle/answer.txt", ... }
仕組み (v1)
エージェントは、新しいユーザー / マウント / PID / 内の再実行された Sandbox-init プロセスで実行されます。
ネットワーク / IPC / UTS 名前空間、完全に特権なし:
oracle/test/solution ファイルはそのマウント名前空間にありません - 空の名前空間でマスクされています
オーバーレイ ワークスペース上の読み取り専用 tmpfs。 open() → ENOENT、構造上
ネットワーク名前空間は空です。直接出力は ENETUNREACH で終了します。ホワイトリスト
モードは、UNIX ソケット上のホスト側の CONNECT/フォワード プロキシを介して、承認されたホストをブリッジします。
すべてのプロキシの許可/拒否の決定は、署名されたタイムラインに記録されます。
リポジトリはベースコミット時に実体化されます。その後の (修正) 履歴は決して残りません。
転送されたため、git log がアクセスできません
seccomp ユーザー通知モニターは、オープン試行を記録します ( open / openat / openat2 )
禁止されたパスと直接出力接続により、ハッシュチェーン化された改ざん防止機能が確立されます。
タイムラインにアクセスすると、常に CONTINUE が応答します。そのため、マウントと netns によって分離が強制されます。
決してモニター越しではありません。施行は建設までに完了します。監査の対象となるのは、
通知リスト上の syscall ではなく、

考えられるあらゆるバリエーション
グレーダーは、真のオラクルであるエージェントに対して、2 番目の分離されたサンドボックスで実行されます。
見たこともない。判定は、RFC-8785 正規 JSON に対する ed25519 署名 +
環境ダイジェスト
すべての実行では、ポータブルな Bundle.json (署名済みの判定 + 違反) も生成されます。
レコード + エージェント ログ ハッシュ、すべて 1 つの署名の下にバインドされます。プロクター検証バンドル
署名、違反チェーン (判定に関連付けられている)、およびログ ハッシュを再チェックします。
安定したオペレーターキー (proctor keygen / PROCTOR_SIGNING_SEED) を使用すると、それが証明されます。
どの演算子が結果を生成したか
設計目標は、ベンチマークに依存しない一般的な標準です。ターミナルベンチ（港）
format) は最初のアダプター ( proctor run-tb ) です。 SWEベンチアダプター
( proctor run-swebench ) 修正履歴を含むベースコミットでリポジトリを実体化します
削除され、 --grade は、分離されたグレーダーを通じてインスタンスのテストを実行します。
CI 上のホスト ネットワーク — を参照してください。
グレーディングレポート、これも文書化します
境界: インスタンスごとに忠実に解決されたグレーディングには、SWE ベンチの固定が必要です
一方、Proctor の再現可能な信号は改ざん防止の完全性を示します。
評決 (git-mining チートはブロックされ、侵害されたとフラグが立てられます)。 --イメージが実行されます
インスタンスの固定された SWE ベンチ イメージ内のエージェント + グレーダー (デーモンレス)
podman/docker fetch)、gitsan のリポジトリは /testbed にオーバーレイされたままになります。
Proctor はサンドボックス内での不正アクセスをブロックします。ファイルシステムを通じて答えに到達します。
git 履歴、ネットワーク、またはプロセス テーブル。到着する回答をブロックしません
サンドボックスの外側から (エージェントのプロンプトに回答キーを挿入する足場、
またはエージェント バイナリ内に密かに密輸されたソリューション) — それらには送信来歴ポリシーが必要です。
v0.2 の焦点 — 採点者の不正行為 (PASS -greps、ハードコードされた出力、モック) ではありません。
どれが緯度ですか

えーフェーズ。クラスごとの完全な情報については、corpus/RESULTS.md を参照してください。
テーブル。
v1 が実装およびリリースされました (Linux、Rust、非特権)。エクスプロイト コーパス
( corpus/ ) は、文書化されたサンドボックス内アクセスチートクラスを再現します。
それぞれがブロックされてログに記録されており、ストック GitHub の CI ではスイート全体が緑色であることをアサートします。
ランナー —
そのため、サンドボックスは開発ボックス上だけでなく、オフマシンに確立される可能性があります。上に発送されました
核心:
署名付きのポータブルな実行バンドル —bundle.json (判定 + 違反 + ログ ハッシュ)
1 つの署名の下で）; proctor verify-bundle はすべてを再チェックします。安定したオペレータキー。
実際のベンチマーク タスク、エンドツーエンド: ターミナル-ベンチ 2 タスク (リファレンス ソリューション →
クリーンパス。 oracle 読み取り → ブロック + ログ）および SWE ベンチ インスタンス
( proctor run-swebench ; --grade は、CI 内の分離されたグレーダーを通じてテストを実行します)。
GitHub Action ( action.yml ) + 事前に構築された v0.1.1 バイナリとしての proctor なので、
ベンチマークの CI は、Proctor の下で数行で実行できます。
ここは新しいですか？まず「なぜプロクターなのか」を読んでから、
最初のタスクを実行するために使用します。完全な設計と脅威モデルについては、を参照してください。
設計仕様と
よくある質問 。バンドル仕様では正確に定義されています
署名された実行から検証者が結論できること、そしてできないこと、検証可能なもの
例のバンドル。
v0.2 — 証明された提出物の出所。プロクターがまだできていない、文書化された最大の不正行為
stop はサンドボックス外の回答密輸です: エージェントの足場を介して注入された回答キー
( AGENTS.md ) またはエージェント バイナリにコンパイルされたソリューション — 研究の背後にあるクラス
1→14ドロップ。 OS 分離では、送信者が持ち込んだ回答を確認できません。v0.2 はそれを閉じます。
反対側から: プロクターは、エージェントに与えられたすべての入力をキャプチャしてコンテンツアドレスにします。
(スキャフォールド、命令ファイル、エージェントバイナリ、環境) をバインドし、署名付きの改ざん防止ファイルをバインドします。
提出マニフェストへの

バンドルを実行する - レビュー担当者が何が入っているかを正確に検証できるように、
エージェントが到達したものだけではありません。違反ログと同じ哲学: を証明します。
入力は信用しないでください。
後で (実際の需要に引っ張られて): PASS に対するグレーダーの強化 -greps / ハードコード
出力/モックされたライブラリ。追加のベンチマーク アダプター。ピン留めされた画像 SWE ベンチ
解決されたグレーディング パス。
ビルド済みバイナリ (Linux x86_64、glibc ≥ 2.35):
gh リリース ダウンロード v0.1.1 --repo dylanp12/proctor \
--パターン 'proctor-x86_64-unknown-linux-gnu.tar.gz*'
sha256sum -c proctor-x86_64-unknown-linux-gnu.tar.gz.sha256
tar -xzf proctor-x86_64-unknown-linux-gnu.tar.gz
sudo インストール proctor-x86_64-unknown-linux-gnu/proctor /usr/local/bin/
プロクター --バージョン
libseccomp2 (ランタイム ライブラリ) が必要です - ほとんどのデバイスにはデフォルトでインストールされています
ディストリビューション (それ以外の場合は sudo apt-get install -y libseccomp2)。 Ubuntu 24.04 (および任意の)
特権のないユーザーの名前空間を制限するディストリビューション) を 1 回有効にするか、実行するたびに失敗します。
sudo sysctl -w kernel.apparmor_restrict_unprivileged_userns=0 。プロクタープローブを実行して、
ホストがサンドボックス化できることを確認します。
sudo apt-get install -y libseccomp-dev # リンク時 libseccomp
カーゴインストール --git https://github.com/dylanp12/proctor proctor-cli
自分で確認してください (60 秒)
コーパスが証拠です: 文書化された 5 つのサンドボックス内チート クラス。それぞれがテストとして再生されます。
これは、ランダムなノンスを「答え」として植え付け、エージェントがそれを決して目にしないと主張します。
git clone https://github.com/dylanp12/proctor && cd プロクター
./scripts/dev-setup.sh # ビルド用に libseccomp をリンクします
# Ubuntu 24.04 (GitHub CI ランナーを含む) は、次の方法で非特権ユーザーの名前空間を無効にします。
# デフォルト — 1 回有効にしないと、サンドボックスの実行がすべて失敗します。
sudo sysctl -w kernel.apparmor_restrict_unprivileged_userns=0
カーゴ テスト -p proctor-cli --test corpus_test -- --nocapture
えー

ch test はタスクを構築し、チートを試行するエージェントを実行し、ブロックされたことを表明します。
(そして、マスクされたリソースに対して syscall が発行された場合、ログに記録されます)。参照
クラスごとのテーブルの corpus/RESULTS.md。
./scripts/dev-setup.sh # ビルド用に libseccomp をリンクします (1 回限り)
カーゴテスト --workspace # ユニット + 分離統合テスト
カーゴ run -p proctor-cli -- プローブ # ホストがサンドボックスできることを確認します
特権のないユーザー名前空間を備えた Linux ≥ 5.11、C libseccomp ≥ 2.5 が必要です
ランタイム、および git 。 Ubuntu 24.04 / CI では、まず特権のないユーザーを有効にします。
sudo sysctl -w kernel.apparmor_restrict_unprivileged_userns=0 。
信頼できる AI コーディング エージェント ベンチマークのための改ざん防止実行サンドボックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
v0.1.1
最新の
2026 年 6 月 23 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Tamper-proof execution sandbox for trustworthy AI coding-agent benchmarks - dylanp12/proctor

GitHub - dylanp12/proctor: Tamper-proof execution sandbox for trustworthy AI coding-agent benchmarks · GitHub
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
dylanp12
/
proctor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
80 Commits 80 Commits .cargo .cargo .github .github corpus corpus crates crates docs docs scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md action.yml action.yml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Proctor turns AI coding-agent benchmark runs into signed, independently verifiable
integrity bundles.
It runs agents in an answer-isolated Linux sandbox where the configured hidden tests, fix
history, and network egress are not reachable, then signs the verdict and the covered
forbidden-access timeline.
AI coding-agent benchmarks are being gamed. In April 2026, UPenn researchers (Stein,
Brown, Hassani, Naik & Wong) documented widespread cheating on popular agent
benchmarks : 1,000+ harness-level cheating
traces across major benchmarks — concentrated in Terminal-Bench 2 and HAL USACO — plus
~30 task-level cases (method paper: Detecting Safety Violations Across Many Agent
Traces , arXiv 2604.11806 ). The cheats are mundane:
agents read the test oracle ( in one removed Terminal-Bench 2 submission, 415 of 429
successful traces were plain filesystem reads of /tests ), mine git log for the fix
commit, curl the solution, or pre-write the grader's reward file. Every one of these is a
sandboxing / access-control failure, not a modeling one.
The study's most dramatic single number — a top submission falling from 1st to 14th
once de-cheated — came from a different cheat class: answer keys injected through the
agent's scaffold ( AGENTS.md ). That arrives from outside the sandbox, so OS isolation
alone can't stop it — it's the target of Proctor's v0.2 provenance work , and
we name that boundary plainly (see Honest claim scope ) rather than
quietly claiming it.
Proctor runs a benchmark task under enforced OS-level isolation so the configured hidden
evaluator artifacts (oracle/tests, fix history, network) are not reachable from the agent's
sandbox, and emits a signed verdict plus a tamper-evident log of covered forbidden-access
attempts (masked-file reads, blocked egress).
proctor run --task ./task --agent "my-agent --solve" --policy ./policy.yaml
# -> verdict.json { "pass": true, "status": "compromised", "violations_count": 1, ... }
# -> violations.jsonl { "step": 8, "kind": "masked_read", "path": "/oracle/answer.txt", ... }
How it works (v1)
The agent runs in a re-exec'd sandbox-init process inside fresh user / mount / PID /
network / IPC / UTS namespaces, fully unprivileged:
oracle/test/solution files aren't in its mount namespace — masked by an empty
read-only tmpfs over an overlay workspace; open() → ENOENT, by construction
the network namespace is empty — direct egress dies with ENETUNREACH ; allowlist
mode bridges approved hosts through a host-side CONNECT/forward proxy over a unix socket,
and every proxy allow/deny decision is recorded in the signed timeline
the repo is materialized at the base commit — later (fix) history is never
transferred, so git log can't reach it
a seccomp user-notification monitor records attempted opens ( open / openat / openat2 )
of forbidden paths and direct egress connect s into a hash-chained, tamper-evident
timeline, then always replies CONTINUE — so isolation is enforced by the mounts and netns,
never by the monitor. Enforcement is complete by construction; the audit covers the
syscalls on the notify list, not every conceivable variant
the grader runs in a second isolated sandbox, against the true oracle the agent
never saw; the verdict is an ed25519 signature over RFC-8785 canonical JSON + an
environment digest
every run also emits a portable bundle.json — the signed verdict + the violation
records + agent-log hashes, all bound under one signature. proctor verify-bundle
re-checks the signature, the violation chain (bound to the verdict), and the log hashes;
with a stable operator key ( proctor keygen / PROCTOR_SIGNING_SEED ) it proves
which operator produced the result
The design goal is a general, benchmark-agnostic standard . Terminal-Bench (Harbor
format) is the first adapter ( proctor run-tb ); a SWE-bench adapter
( proctor run-swebench ) materializes the repo at the base commit with fix history
stripped, and --grade runs the instance's tests through the isolated grader over the
Host network on CI — see the
grading report , which also documents
the boundary: faithful per-instance resolved-grading needs SWE-bench's pinned
environment, while Proctor's reproducible signal is the tamper-evident integrity
verdict (the git-mining cheat is blocked + flagged compromised ). --image runs
the agent + grader inside the instance's pinned SWE-bench image (daemonless
podman/docker fetch) with the gitsan'd repo still overlaid at /testbed .
Proctor blocks in-sandbox access cheats — reaching the answer through the filesystem,
git history, the network, or the process table. It does not block answers that arrive
from outside the sandbox (a scaffold that injects answer keys into the agent's prompt,
or solutions smuggled inside the agent binary) — those need submission-provenance policy,
the focus of v0.2 — nor grader-fooling ( PASS -greps, hardcoded outputs, mocks),
which is a later phase. See corpus/RESULTS.md for the full per-class
table.
v1 implemented and released (Linux, Rust, unprivileged). The exploit corpus
( corpus/ ) replays the documented in-sandbox access-cheat classes it covers and
asserts each is blocked and logged, and the full suite is green in CI on a stock GitHub
runner —
so the sandbox provably establishes off-machine, not just on a dev box. Shipped on top of
the core:
Signed, portable run bundles — bundle.json (verdict + violations + log hashes
under one signature); proctor verify-bundle re-checks everything; stable operator keys.
Real benchmark tasks, end-to-end: a Terminal-Bench 2 task (reference solution →
clean pass; oracle read → blocked + logged) and a SWE-bench instance
( proctor run-swebench ; --grade runs the tests through the isolated grader in CI).
proctor as a GitHub Action ( action.yml ) + a prebuilt v0.1.1 binary, so a
benchmark's CI can run under Proctor in a few lines.
New here? Read Why Proctor first, then
usage to run your first task. For the full design and threat model see
the design spec and the
FAQ . The bundle spec defines exactly
what a verifier can — and cannot — conclude from a signed run, with a verifiable
example bundle .
v0.2 — attested submission provenance. The biggest documented cheat Proctor can't yet
stop is out-of-sandbox answer smuggling: answer keys injected through the agent's scaffold
( AGENTS.md ) or a solution compiled into the agent binary — the class behind the study's
1st→14th drop. OS isolation can't see an answer the submitter carries in. v0.2 closes it
from the other side: Proctor captures and content-addresses every input the agent was given
(scaffold, instruction files, agent binary, environment) and binds a signed, tamper-evident
submission manifest into the run bundle — so a reviewer can verify exactly what went in,
not just what the agent reached for. Same philosophy as the violation log: attest the
inputs, don't trust them.
Later (pulled by real demand): grader hardening against PASS -greps / hardcoded
outputs / mocked libraries; additional benchmark adapters; a pinned-image SWE-bench
resolved-grading path.
Prebuilt binary (Linux x86_64, glibc ≥ 2.35):
gh release download v0.1.1 --repo dylanp12/proctor \
--pattern 'proctor-x86_64-unknown-linux-gnu.tar.gz*'
sha256sum -c proctor-x86_64-unknown-linux-gnu.tar.gz.sha256
tar -xzf proctor-x86_64-unknown-linux-gnu.tar.gz
sudo install proctor-x86_64-unknown-linux-gnu/proctor /usr/local/bin/
proctor --version
Needs libseccomp2 (the runtime library) present — installed by default on most
distributions ( sudo apt-get install -y libseccomp2 otherwise). On Ubuntu 24.04 (and any
distro that restricts unprivileged user namespaces) enable them once or every run fails:
sudo sysctl -w kernel.apparmor_restrict_unprivileged_userns=0 . Run proctor probe to
confirm your host can sandbox.
sudo apt-get install -y libseccomp-dev # link-time libseccomp
cargo install --git https://github.com/dylanp12/proctor proctor-cli
Verify it yourself (60 seconds)
The corpus is the proof: five documented in-sandbox cheat classes, each replayed as a test
that plants a random nonce as the "answer" and asserts the agent never sees it.
git clone https://github.com/dylanp12/proctor && cd proctor
./scripts/dev-setup.sh # links libseccomp for the build
# Ubuntu 24.04 (incl. the GitHub CI runner) disables unprivileged user namespaces by
# default — enable once, or every sandbox run fails:
sudo sysctl -w kernel.apparmor_restrict_unprivileged_userns=0
cargo test -p proctor-cli --test corpus_test -- --nocapture
Each test builds a task, runs an agent that tries the cheat, and asserts it's blocked
(and, where a syscall is issued against a masked resource, logged). See
corpus/RESULTS.md for the per-class table.
./scripts/dev-setup.sh # links libseccomp for the build (one-time)
cargo test --workspace # unit + isolation integration tests
cargo run -p proctor-cli -- probe # check the host can sandbox
Requires Linux ≥ 5.11 with unprivileged user namespaces, a C libseccomp ≥ 2.5
runtime, and git . On Ubuntu 24.04 / CI, enable unprivileged userns first:
sudo sysctl -w kernel.apparmor_restrict_unprivileged_userns=0 .
Tamper-proof execution sandbox for trustworthy AI coding-agent benchmarks
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
v0.1.1
Latest
Jun 23, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
