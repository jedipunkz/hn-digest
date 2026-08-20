---
source: "https://github.com/CaydenChik/doover"
hn_url: "https://news.ycombinator.com/item?id=49371211"
title: "Show HN: Do-over, undo for AI agent shell commands"
article_title: "GitHub - CaydenChik/doover: Undo for AI agent shell commands. Snapshots files before Claude Code runs destructive bash (rm -rf, git reset, rsync), so agent mistakes are reversible, even for files git never tracked. · GitHub"
image: "https://repository-images.githubusercontent.com/1290762482/95fdd3bc-53da-4b44-a3f8-d632513a9a0a"
author: "Cayden27"
captured_at: "2026-08-20T07:33:18Z"
capture_tool: "hn-digest"
hn_id: 49371211
score: 2
comments: 1
posted_at: "2026-08-20T06:42:36Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Do-over, undo for AI agent shell commands

- HN: [49371211](https://news.ycombinator.com/item?id=49371211)
- Source: [github.com](https://github.com/CaydenChik/doover)
- Score: 2
- Comments: 1
- Posted: 2026-08-20T06:42:36Z

## Translation

タイトル: HN の表示: AI エージェント シェル コマンドのやり直し、元に戻す
記事のタイトル: GitHub - CaydenChik/doover: AI エージェントのシェル コマンドを元に戻すクロード コードが破壊的な bash (rm -rf、git restart、rsync) を実行する前にファイルのスナップショットを作成するため、git で追跡されなかったファイルであっても、エージェントのミスを元に戻すことができます。 · GitHub
説明: AI エージェントのシェル コマンドを元に戻す。クロード コードが破壊的な bash (rm -rf、git restart、rsync) を実行する前にファイルのスナップショットを作成するため、git で追跡されなかったファイルであっても、エージェントのミスを元に戻すことができます。 - ケイデンチック/ドゥーバー
HN テキスト: こんにちは、HN。プロジェクトの整理中にコーディング エージェントが私のファイルを削除した後、これをビルドしました。 Do-over は Claude Code のフックにプラグインし、コマンドが実行される直前にそのコマンドが触れようとしている内容をスナップショットします。あなたが何を考えているかはわかります。もっと頻繁にコミットしないのはなぜですか?このツールは git に代わるものではありません。それが行うことは次のとおりです。
[1] git が保護しないものを保護します: 追跡されていない/無視されたファイルと、リポジトリの外に存在するファイル (これが、ある男が 50 GB のデータを失った方法です - Claude code issue tracker)
[2] git 自体がエージェントによって武器化された場合に保護します: git checkout .、git clean -fd、および git replace --hard (3 年前の Unity プロジェクトがこの方法で削除されました - Claude Code issue tracker)
[3] 忘れっぽくてコミットメントを怠った場合にあなたを守る これは、潜在的な頭痛に対する簡単な予防策です。そして、これは現在の対策では確保できないものを確保する最後のワンマイルの取り組みです。これには、サンドボックス、チェックポイント、git が含まれます。

記事本文:
GitHub - CaydenChik/doover: AI エージェントのシェル コマンドを元に戻す。クロード コードが破壊的な bash (rm -rf、git restart、rsync) を実行する前にファイルのスナップショットを作成するため、git で追跡されなかったファイルであっても、エージェントのミスを元に戻すことができます。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ケイデンチック
/
いたずら
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
75 コミット 75 コミット フォルダーとファイル
.github .github アセット アセット ベンチ ベンチ クレート クレート

es docs docs パッケージング/ homebrew パッケージング/ homebrew スクリプト スクリプト テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LAUNCH.md LAUNCH.md LICENSE LICENSE Makefile Makefile NOTICE.md NOTICE.md README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-LICENSES.txt THIRD-PARTY-LICENSES.txt Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
どのエージェントもやり直す価値がある。
AI エージェントのシェル コマンドを元に戻します。 doover スナップショット ファイルを作成する前に、
エージェントの破壊的なコマンドが実行され、実行されたすべての記録が記録されます。
エージェントが外部で触れたファイルも含めて、実際に元に戻すことができます
あなたのプロジェクトや、git が見たことのないものすべてについて。
クリップ内のエージェント コマンドは、正確な動作を実行する代役です。
フックフロークロードコードドライブ。これは実際のクロード コードの録音です
セッションはエンドツーエンドで同じことを行います。
$ claude "ビルドアーティファクトをクリーンアップ"
⏺ Bash(rm -rf dist/ photos/) # ...2 番目のやつは痛かったです。
$ 道を越えたログ
#42 破壊的rm -rf dist/ photos/完成
$ やり直し元に戻す
アクション #42 の取り消しが完了しました: 2 つのパスが復元されました
$ ls 写真/
Birthday.jpg Wedding.jpg # 戻って、バイトごとに。
なぜこれが存在するのか
コーディング エージェントは 1 日中シェル コマンドを実行しますが、シェル コマンドには元に戻す機能がありません。
既存の安全機構はそれぞれ同じ場所の手前で停止します。
Claude Code チェックポイントは、ファイル ツールを通じて行われた編集を巻き戻します。
ただし、Bash ツールを通じて行われた変更にはチェックポイントが設定されません。 rm -rf
永遠です。
サンドボックス (Codex スタイル) は、爆発範囲をワークスペースに制限します。
便利ですが、ワークスペース内での削除は依然として回復できません。
git はコミットしたものを保護します。追跡されていないファイルに対しては何も行われません。
無視されたファイル ( .env 、ローカル データベース、そのフォルダー

テストデータ)、または任意の
リポジトリではないディレクトリ。エージェント自体も実行可能
git チェックアウト 。または git clean -fd 、コミットされていない作業を破棄します
gitを使用して。
doover には欠けているレイヤー、つまりエージェント シェル アクションのトランザクション ログがあります。
許可を求めたり、何もブロックしたりしません。それは、
代わりに危険なコマンドを元に戻すことができます。
macOS または Linux (WSL は機能しますが、ネイティブ Windows は機能しません)。
Cargo (Rust 1.85+ が必要; --locked は監査されたものを正確にインストールします
リリースのビルドとテストに使用された依存関係セット):
$ カーゴインストール doover --locked
自作:
$ ブリュータップ ケイデンチック/ドゥーバー
$ brew trust caydenchik/doover # newer Homebrew はサードパーティのタップごとに 1 回質問します
$ brew install doover
すべてのプラットフォーム用に事前に構築されたバイナリは、
リリースページ、
SHA256SUMS を確認します。
$ doover init # ~/.claude/settings.json にフックを追加します
$ doover Doctor # すべてをエンドツーエンドで検証します
doover init --project を使用して単一プロジェクトにインストールします
グローバルではなく ( ./.claude/settings.json )。 init はあなたのものとマージします
既存の設定はそのままであり、それ自体が複製されることはありません。いつでも医者にかかりなさい
何か違和感を感じます。
アンインストールするには、設定ファイルから 2 つの doover フック エントリを削除します。
スナップショットは、それも削除するまで ~/.doover に残ります。
ほとんどの場合、いたずらには気付かないでしょう。これは Claude Code の PreToolUse / の背後にあります。
PostToolフックを使用し、コマンドごとに数ミリ秒を追加し、音声のみを送信します
何かを完全に守ることができなかったとき。それからある日：
知っておく価値のあるいくつかの動作:
元に戻すには競合チェックが行われます。アクションの後にファイルが変更された場合、
元に戻す（後のコマンド、またはあなた）、ドゥーバーは拒否（終了コード 3）し、指示します
なぜだ。 --force はとにかく続行します。 --dry-run は最初に計画を示します。
元に戻すこと自体がジャーナルに記録されます。元に戻す操作を元に戻すことは、やり直しの仕組みです。歴史
追加専用です。何も黙って書き換えられることはありません。
やり直し元に戻します

実際に何かを変更した最後のコマンドを対象とします。
読み取り専用コマンドは、それを中心に doover スナップショットが作成された場合でもスキップされます。
Undo は冪等であり、自身のレコードよりもディスクを信頼します。元に戻す
同じアクションを 2 回実行し、2 回目は何もしません。しかし、何かが起こった場合、
アクションの効果を元に戻します (後のアクションを強制的に元に戻すことができます)。
すでに元に戻したことを伝えるのではなく、もう一度元に戻します。ファイルが
doover の簿記ではなく、あなたのファイルについて質問があります。
ディレクトリ全体を復元すると、ディレクトリが置き換えられます。あなたの殻が座っているなら
そのディレクトリ内で cd を実行します。その後リフレッシュします。戸惑いが告げる
これが起こったときのあなた。
.git は git に任せます。ツリー全体のスナップショットは .git を通り過ぎます (それは
多くの場合、作業ツリーよりも大きいため、「元に戻す」を決して巻き戻してはなりません
git の背後にあるリポジトリ履歴);元に戻すと、そのままになります。
コマンドを直接指定すると ( rm -rf .git )、完全にキャプチャされます。
ビルド ディレクトリはスキップされますが、これは git が同意した場合に限ります。行き帰りのとき
target/ 、node_modules/ 、.venv/ を通過するツリー全体のスナップショットを作成します。
そして友達: ビルドはそれらを再作成し、それらをキャプチャするには次の時間を費やします。
ソースではなく成果物に全時間の予算を置きます。ディレクトリはただ
その名前がそのリストにあり、git がすでにそれを無視している場合はスキップされます。
実際のソースを build/ というフォルダーに保存し、git がそれを追跡すると、
他のものと同じようにキャプチャされます。コマンドをまっすぐに向ける
( rm -rf target ) 完全にキャプチャされます。ドゥーバーショーリスト何でも
はスキップされ、元に戻すとそれらのフォルダーはそのまま残ります。
部分的なスナップショットは部分的に復元されます。スナップショットが切り取られた場合
短い (以下の制限を参照)、元に戻すは置き換えるのではなくデフォルトで拒否されます。
部分コピーを含む完全なツリー。
エージェントの実行: rm -rf build/
│
▼
PreToolUse フック ── バーを解析する

sh ── レジストリに照らして分類する
│ rm → 破壊的、スコープ: build/
▼
スナップショット build/ into ~/.doover/store (コピーオンライト、コンテンツアドレス指定)
│
▼
アクションを記録する (SQLite) ── その後、コマンドが実際に実行されます
│
▼後ほど…
PostToolUse フックは事後の状態を記録します ──── doover undo returns build/
興味深い部分:
実際の bash パーサー (正規表現ではない) は、各コマンドが触れる内容を解決します。
&& チェーン、パイプ、リダイレクト、グロブ、引用を介して。何でも
(コマンド置換、 eval 、不明なツール) を完全には説明できません
潜在的に破壊的なものとして扱われ、決して安全であるとは考えられていません。
152 CC0 ライセンスの可逆性レジストリ
エージェントが実際に実行するコマンドを安全なものから安全なものまで分類する YAML ルール
不可逆: rm 、 mv 、 git checkout 、 rsync --delete 、 gzip 、
wget -O を危険にさらすか、どのパスをキャプチャするか。読み取り専用であることが証明されたコマンド
そのように分類されているため、費用はかかりません。興味深いエントリは、
無害に見えて実際はそうではないもの、たとえば、静かに切り捨てるものなど
出力ファイル。
コピーオンライトのスナップショット。 APFS/Btrfs/XFS では、ファイルを「コピー」する前に
削除ではディスク ブロックが共有されるため、1 GB ディレクトリのスナップショットの作成にコストがかかります
オリジナルが実際に変更されるまではほとんど何もありません。ファイルが保存されます
BLAKE3 ハッシュによってアドレス指定され、復元するたびに再度検証されます。
復元は段階的に行われます。 doover は復元されたツリーをその隣に構築します。
ターゲットにして丸ごと交換します。復元中にクラッシュするとファイルが残る
まさにその通りでした。
パーサーが証明できる内容に応じて 3 つの層:
その中間層は内部化するものです。doover が解析できないコマンドの場合、
保護は作業ディレクトリのみをカバーします。を削除するスクリプト
~/something-else は、静的解析で確認できる範囲外にあります。
Apple Silicon / APFS で測定 (bench/hook_latency.py を自分で実行):
～

スナップショットを必要とするものが何もない場合、コマンドごとに 5 ～ 10 ミリ秒。
コマンド ( ls 、 cat 、 git status 、 builds、tests) — 再測定
2026 年 8 月 15 日のライブ エージェント トライアルでのエンドツーエンド (開始前最大 6 ミリ秒 + 終了後最大 3 ミリ秒、
プロセスの生成を含む)。
スナップショットのコストはバイト数ではなくファイル数に応じて変化します。ファイルあたり約 0.19 ミリ秒。
1 つの 100 MB ファイルのコストは最大 70 ミリ秒です。
スナップショットは 5 秒で停止する (構成可能) ため、巨大なツリーが停止することはありません。
エージェントを失速させます。日誌には、捕獲が部分的であったことが記録されている。
すべては環境変数です。デフォルトはそのままにしておくべきです
一人で。
固定されたアクション ( doover pin <id> ) は、クリーンアップおよび最新のアクションに存続します。
1 時間の歴史がスペースのために追い出されることはありません。
悪意のあるエージェントに対する防御ではありません。 doover 分析コマンド
静的に;敵がそれを回避したい場合は回避できます。から保護します
これはエージェントが実際に生み出すものであり、攻撃に対するものではありません。
金庫ではなく、シートベルトのように扱ってください。
バックアップツールではありません。履歴には制限があります (デフォルトでは 7 日間 / 5 GiB)。
同じディスク上に存在します。実際のバックアップを保管してください。
リモートエフェクトを元に戻すことはできません。データベースの削除、ポッドの削除、
強制的に押し出された枝: ドアオーバーはそれが起こったことを教えてくれます。それを元に戻すことはできません。
保存時には暗号化されません。スナップショットはファイルのコピーであり、読み取りのみ可能です
ユーザーアカウント (0700 / 0600) によって。あなたのアカウントまたは root を持っている人なら誰でも、
それらを読んでください。コマンドはジャーナルに書き込まれる前に編集されます。
したがって、Doover が認識する資格情報はディスクに到達しませんが、一致は
DLP エンジンではなく、パターンベースの衛生管理: 非常に珍しい秘密が得られます
を通して。スナップショット ストアも、その中のファイルと同様に機密性の高いものとして扱い続けてください。
所有権ではなく、ファイルの内容。スナップショットはファイルの内容を復元します。
権限 (モード)、タイムスタンプ、および拡張属性。そうではありません
所有権の取得または復元 (uid/gid): 元に戻す

chown / chgrp は、
データは戻されるが所有者は戻されず、所有権のみの変更は自動選択されません
裸の doover undo によって ( doover undo <id> で明示的に名前を付けます)。あなたの
データは常に安全です。所有者フィールドは復元時にそのまま残される唯一のものです
それを見つけます。
git、チェックポイント、サンドボックスの代わりとなるものではありません。それはレイヤーです
それらはすべて開いたままになります。 3つとも使い続けてください。
YAML ファイルを ~/.doover/registry.d/ にドロップして、doover に自分のことを教えます
ツール:
ルール：
- ID : my.dbtool
一致 : { コマンド: dbtool、サブコマンド: ワイプ }
効果 : 破壊的
スコープ : { パス: 位置 }
元に戻す: スナップショットの復元
オーバーレイによりコマンドを追加し、分類を強化できます。彼らにはできません
出荷されたものを弱める: rm が安全であるというルールは無視され、
どのような表現であっても警告です。
レジストリ データは正確に CC0 (パブリック ドメイン) であるため、他のツールが盗むことができます
それ。あるコマンドが実際に何を破壊するかを計画している場合は、PR を送信してください。それ
知識はこのプロジェクトで最も再利用可能な部分です。
クロード・コード以外のエージェントでも機能しますか?
コアはエージェントに依存しません。フック配線は現在クロード コードをターゲットにしています。
フックイベント。他のハーネス用のアダプターは当然のことです。
doover フック pre は標準入力の JSON を読み取るだけです。
複数のエージェントを同時に?
はい。日記はデザインです

[切り捨てられた]

## Original Extract

Undo for AI agent shell commands. Snapshots files before Claude Code runs destructive bash (rm -rf, git reset, rsync), so agent mistakes are reversible, even for files git never tracked. - CaydenChik/doover

Hello HN, I built this after a coding agent deleted my files while tidying a project. Do-over plugs into Claude Code's hooks and snapshots what a command is about to touch right before it runs. I know what you're thinking. Why don't you just commit more often? This tool doesn't seek to replace git. What it does is:
[1] Protect the stuff git doesn't protect: your untracked/ignored files and files that exist outside your repo (which is how a guy lost 50GB of data - Claude code issue tracker)
[2] Protects you when git itself is weaponized by your agent: git checkout ., git clean -fd, and git reset --hard (a three year old Unity project was deleted this way - Claude Code issue tracker)
[3] Protects you in the case when you're just forgetful and don't commit It really is a simple safeguard to potential headaches. And it's a last mile effort that secures what current measures can't. That includes sandboxes, checkpoints, git.

GitHub - CaydenChik/doover: Undo for AI agent shell commands. Snapshots files before Claude Code runs destructive bash (rm -rf, git reset, rsync), so agent mistakes are reversible, even for files git never tracked. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
CaydenChik
/
doover
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
75 Commits 75 Commits Folders and files
.github .github assets assets bench bench crates crates docs docs packaging/ homebrew packaging/ homebrew scripts scripts tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LAUNCH.md LAUNCH.md LICENSE LICENSE Makefile Makefile NOTICE.md NOTICE.md README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-LICENSES.txt THIRD-PARTY-LICENSES.txt rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Every agent deserves a do-over.
Undo for your AI agent's shell commands. doover snapshots files before your
agent's destructive commands run, keeps a journal of everything it did, and
gives you a real undo , including for files your agent touched outside
your project and for everything git never saw.
The agent command in the clip is a stand-in that drives the exact
hook flow Claude Code drives. Here is a recording of a real Claude Code
session doing the same thing end to end.
$ claude " clean up the build artifacts "
⏺ Bash(rm -rf dist/ photos/) # ...that second one hurt.
$ doover log
#42 completed destructive rm -rf dist/ photos/
$ doover undo
undo of action #42 complete: 2 path(s) restored
$ ls photos/
birthday.jpg wedding.jpg # back, byte for byte.
Why this exists
Coding agents run shell commands all day, and shell commands have no undo.
The existing safety mechanisms each stop short of the same spot:
Claude Code checkpoints rewind edits made through its file tools,
but changes made through the Bash tool aren't checkpointed . rm -rf
is forever.
Sandboxes (Codex-style) confine the blast radius to your workspace.
Useful, but inside the workspace deletion still has no recovery.
git protects what you committed. It does nothing for untracked files,
ignored files ( .env , local databases, that folder of test data), or any
directory that isn't a repo. And the agent itself can run
git checkout . or git clean -fd , which destroy uncommitted work
using git.
doover is the missing layer: a transaction log for agent shell actions.
It doesn't ask for permission and it doesn't block anything. It makes the
dangerous commands reversible instead.
macOS or Linux (WSL works; native Windows doesn't).
Cargo (needs Rust 1.85+; --locked installs the exact audited
dependency set the release was built and tested with):
$ cargo install doover --locked
Homebrew:
$ brew tap caydenchik/doover
$ brew trust caydenchik/doover # newer Homebrew asks once per third-party tap
$ brew install doover
Prebuilt binaries for every platform are on the
releases page , with
SHA256SUMS to verify.
$ doover init # adds hooks to ~/.claude/settings.json
$ doover doctor # verifies everything end to end
Use doover init --project to install for a single project
( ./.claude/settings.json ) instead of globally. init merges with your
existing settings and never duplicates itself; run doctor any time
something feels off.
To uninstall, remove the two doover hook entries from your settings file.
Your snapshots stay in ~/.doover until you delete that too.
You mostly won't notice doover. It sits behind Claude Code's PreToolUse /
PostToolUse hooks, adds a few milliseconds per command, and speaks up only
when it couldn't fully protect something. Then one day:
A few behaviors worth knowing:
Undo is conflict-checked. If a file changed after the action you're
undoing (a later command, or you), doover refuses (exit code 3) and tells
you why. --force proceeds anyway; --dry-run shows the plan first.
Undo is itself journaled. Undoing an undo is how redo works. History
is append-only; nothing is ever silently rewritten.
doover undo targets the last command that actually changed something.
Read-only commands are skipped, even when doover snapshotted around them.
Undo is idempotent, and it trusts your disk over its own records. Undo
the same action twice and the second is a no-op. But if something puts an
action's effect back (a forced undo of a later action can), doover will
undo it again rather than tell you it already did. Whether your files are
there is a question about your files, not about doover's bookkeeping.
Restoring a whole directory replaces it. If your shell is sitting
inside that directory, run cd . afterwards to refresh it. doover tells
you when this happens.
.git is left to git. Whole-tree snapshots walk past .git (it is
often larger than the working tree, and undo should never rewind
repository history behind git's back); undo leaves it exactly as it is.
Point a command straight at it ( rm -rf .git ) and it is captured in full.
Build directories are skipped, but only if git agrees. When doover
snapshots a whole tree it walks past target/ , node_modules/ , .venv/
and friends: a build recreates them, and capturing them would spend the
whole time budget on artifacts instead of your source. A directory is only
skipped when its name is on that list and git already ignores it — so
if you keep real source in a folder called build/ and git tracks it, it
is captured like anything else. Point a command straight at one
( rm -rf target ) and it is captured in full. doover show lists whatever
was skipped, and undo leaves those folders exactly as they are.
Partial snapshots restore partially, and say so. If a snapshot was cut
short (see limits below), undo refuses by default rather than replace a
full tree with a partial copy.
agent runs: rm -rf build/
│
▼
PreToolUse hook ── parse the bash ── classify against the registry
│ rm → destructive, scope: build/
▼
snapshot build/ into ~/.doover/store (copy-on-write, content-addressed)
│
▼
journal the action (SQLite) ── then the command actually runs
│
▼ later…
PostToolUse hook records the after-state ──── doover undo restores build/
The interesting parts:
A real bash parser (not regexes) resolves what each command touches,
through && chains, pipes, redirects, globs, and quoting. Anything it
can't fully account for (command substitution, eval , unknown tools) is
treated as potentially destructive, never assumed safe.
A reversibility registry of 152 CC0-licensed
YAML rules classifying the commands agents actually run, from safe to
irreversible : what rm , mv , git checkout , rsync --delete , gzip ,
wget -O put at risk, and which paths to capture. Commands proven read-only
are classified as such so they cost nothing; the interesting entries are the
ones that look harmless and aren't, like sort -o quietly truncating its
output file.
Copy-on-write snapshots. On APFS/Btrfs/XFS, "copying" a file before
deletion shares its disk blocks, so snapshotting a 1 GB directory costs
almost nothing until the original actually changes. Files are stored
once, addressed by BLAKE3 hash, verified again before every restore.
Restores are staged. doover builds the restored tree next to the
target and swaps it in whole. A crash mid-restore leaves your files
exactly as they were.
Three tiers, depending on what the parser can prove:
That middle tier is the one to internalize: for commands doover can't parse,
protection covers your working directory only . A script that deletes
~/something-else is outside what static analysis can see.
Measured on Apple Silicon / APFS (run bench/hook_latency.py yourself):
~5–10 ms per command when nothing needs snapshotting, which is most
commands ( ls , cat , git status , builds, tests) — re-measured
end-to-end in the 2026-08-15 live-agent trial (~6 ms pre + ~3 ms post,
including process spawn).
Snapshot cost scales with file count , not bytes: ~0.19 ms per file;
a single 100 MB file costs ~70 ms.
Snapshots stop at 5 seconds (configurable) so a huge tree can never
stall your agent. The journal records that the capture was partial.
Everything is an environment variable; the defaults are meant to be left
alone.
Pinned actions ( doover pin <id> ) survive any cleanup, and the most recent
hour of history is never evicted for space.
Not a defense against a malicious agent. doover analyzes commands
statically; an adversary who wants to evade it can. It protects against
mistakes, which is what agents actually produce, not against attacks.
Treat it like a seatbelt, not a vault.
Not a backup tool. History is bounded (7 days / 5 GiB by default) and
lives on the same disk. Keep real backups.
Not able to undo remote effects. Dropped databases, deleted pods,
force-pushed branches: doover tells you it happened; it can't reverse it.
Not encrypted at rest. Snapshots are copies of your files, readable only
by your user account ( 0700 / 0600 ). Anyone with your account, or root, can
read them. Commands are redacted before they are written to the journal,
so credentials doover recognizes never reach the disk, but the matching is
pattern-based hygiene, not a DLP engine: an exotic enough secret gets
through. Keep treating your snapshot store as sensitive as the files in it.
File content, not ownership. A snapshot restores a file's contents,
permissions (mode), timestamps, and extended attributes. It does not
capture or restore ownership (uid/gid): undoing a chown / chgrp brings the
data back but not the owner, and an ownership-only change isn't auto-selected
by a bare doover undo (name it explicitly with doover undo <id> ). Your
data is always safe; the owner field is the one thing a restore leaves as it
finds it.
Not a replacement for git, checkpoints, or sandboxes. It's the layer
they all leave open. Keep using all three.
Drop YAML files in ~/.doover/registry.d/ to teach doover about your own
tools:
rules :
- id : my.dbtool
match : { command: dbtool, subcommand: wipe }
effect : destructive
scope : { paths: positional }
undo : snapshot-restore
Overlays can add commands and strengthen classifications. They can't
weaken a shipped one: a rule that says rm is safe is ignored, with a
warning, no matter how it's phrased.
The registry data is CC0 (public domain) precisely so other tools can steal
it. If you map out what some command really destroys, send a PR. That
knowledge is the most reusable part of this project.
Does it work with agents other than Claude Code?
The core is agent-agnostic; the hook wiring currently targets Claude Code's
hook events. Adapters for other harnesses are a natural contribution;
doover hook pre just reads JSON on stdin.
Multiple agents at once?
Yes. The journal is desig

[truncated]
