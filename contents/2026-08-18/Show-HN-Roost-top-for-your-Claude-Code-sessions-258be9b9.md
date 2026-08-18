---
source: "https://github.com/gmhoward9289-ops/roost"
hn_url: "https://news.ycombinator.com/item?id=49353452"
title: "Show HN: Roost – top for your Claude Code sessions"
article_title: "GitHub - gmhoward9289-ops/roost: top for Claude Code - live sessions, models, context, and the subagents they spawn · GitHub"
image: "https://opengraph.githubassets.com/634aace57810f9362be25c42322adf514700e13d36fd12bcee8929990c147f24/gmhoward9289-ops/roost"
author: "gmhoward9289"
captured_at: "2026-08-18T22:12:36Z"
capture_tool: "hn-digest"
hn_id: 49353452
score: 1
comments: 0
posted_at: "2026-08-18T22:10:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Roost – top for your Claude Code sessions

- HN: [49353452](https://news.ycombinator.com/item?id=49353452)
- Source: [github.com](https://github.com/gmhoward9289-ops/roost)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T22:10:08Z

## Translation

タイトル: HN を表示: Roost – クロード コード セッションのトップ
記事のタイトル: GitHub - gmhoward9289-ops/roost: クロード コードのトップ - ライブ セッション、モデル、コンテキスト、およびそれらが生成するサブエージェント · GitHub
説明: クロード コードのトップ - ライブ セッション、モデル、コンテキスト、およびそれらが生成するサブエージェント - gmhoward9289-ops/roost

記事本文:
GitHub - gmhoward9289-ops/roost: クロード コードのトップ - ライブ セッション、モデル、コンテキスト、およびそれらが生成するサブエージェント · GitHub
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
gmhoward9289-ops
/
ねぐら
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
69 コミット 69 コミット フォルダーとファイル
.github .github bin bin デモ デモ ドキュメント ドキュメント パッケージング パッケージング スクリプト スクリプト テスト テスト .gitattributes .gitattributes .gitignore .gitignore .re

リース-プリーズ-マニフェスト.json .リリース-プリーズ-マニフェスト.json .repo-visibility .repo-visibility CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md llms.txt llms.txt package.json package.json pyproject.toml pyproject.toml release-please-config.json release-please-config.json roost.1 roost.1 roost.py roost.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード・コードのトップ。すべてのライブ セッション、そのセッションが使用されているモデル、コンテキストの量
それは燃えました。そして、他のものとは異なり、それが生み出したサブエージェントも燃えました。
1 つのファイル、依存関係なし、Python 3.9 以降。 macOS、Linux、Windows 上で動作します。
以下の短いアンビエント ループは、同じプログラムのアイドリングです。
ワーカー モデル CTX トークン トレンド フロー アイドル タスク
限界に近い
デモ-a1 opus-5 85% 170k +4k ..:-=+#+=-.. 12秒 パーサーをリファクタリングする
駐車+費用がかかる
デモ-b2 opus-5 61% 122k +2k ................................ 4 時間 10 分 ビルド スクリプトを監査する
現在作業中
デモ-c3 fable-5 22% 44k +22k ...:=+*#+ 3s 統合テストを追加
開始中
デモ-d4 - - - - -
QUIET (4) デモ-e5 。デモ-f6 。デモ-g7 。デモ-h8
8 人の労働者 | 120万個保有 |過去 8 ターンで 61k | 1 限界近く |寓話-5、作品-5
サブエージェント
ステート エージェント モデル CTX アイドル タスク
作業中 Explore/a812a opus-5 66k/200k 2s 構成ローダーを調査します
アイドル adaffaba4b Sonnet-5 484k/1M 1h22m 移行メモのドラフト
サブエージェント 2 人、稼働中の 1 人
INFRA ollama:11434 アップ qwen2.5-coder:14b (9.2 GB) litellm:4000 アップ openwebui:8080 ダウン
セッションは、サイズではなく、無視する場合のコストによってグループ化されます: NEAR LIMIT
動作を停止しようとしている場合、PARKED + COSTLY は次のコンテキスト全体に請求します
回転すると、静かなすべてが一本の線に崩壊します。
TREND は、セッションの最後の数ターンにわたって追加されたコンテキストです。それは金額ですし、
セッション内のコンテキストは上昇するだけであり、下降するため、スパークラインではありません。
/compa のみ

ct — したがって、シェイプはすべての行に同じランプを描画します。 CTX
セッションがどの程度埋まっているかを示します。 TREND は、どのくらいの速さで埋まっているかを示します。
上記のdemo-c3は22%で注目すべきものとして表示されます。から読み戻されます。
roost の実行中に蓄積されるのではなくトランスクリプトなので、最初の時点で存在します。
フレーム以下 --once 。それがFLOWとの違いでもあります
それ: FLOW は、roost が開始されてから空で開始されてからサンプリングされたスループットをスパークします。
TREND は、roost がファイルを開いたときにすでにファイルに含まれている金額です。
最後の行はフリートを合計します。保持されているコンテキストは、スイープによって再利用されるものです。
brew install gmhoward9289-ops/tap/roost
Debian と Ubuntu、リポジトリから — これにより、適切なアップグレードも可能になります。
カール -fsSL https://gmhoward9289-ops.github.io/roost/roost-archive-keyring.asc | sudo gpg --dearmor -o /usr/share/keyrings/roost-archive-keyring.gpg
echo " deb [signed-by=/usr/share/keyrings/roost-archive-keyring.gpg] https://gmhoward9289-ops.github.io/roost 安定した main " | sudo tee /etc/apt/sources.list.d/roost.list
sudo aptアップデート
sudo apt install roost
または、リポジトリを追加せずに、roost_<version>_all.deb を取得します。
最新リリース
そのファイルを直接インストールします。例:
sudo apt install ./roost_ <バージョン> _all.deb
どちらも、通常のリポジトリのインストールとまったく同じように python3 を解決します。レポは
専用のキーで署名されます (個人情報には関連付けられません)。その公共の
半分はこのリポジトリの package/apt/pubkey.asc であり、変更せずに公開されています。
上記のroost-archive-keyring.asc。
pipx インストール roost-top
PyPI は、裸の名前 roost を予備として保持します。以前のプロジェクトの名前が保持されます。
削除後 — したがって、パッケージは roost-top になります。インストールするコマンドは
普通のねぐら。リポジトリから直接インストールすると、インデックスが完全にスキップされます。
pipx インストール git+https://github.com/gmhoward9289-ops/roost
npm を使用すると、ワンチャ

Windows に実際の roost コマンドを提供する nnel:
npm install -g roost-top
または、インストールせずに実行します: npx roost-top 。 npm パッケージはラッパーであり、ラッパーではありません。
ポート: 同じ roost.py を同梱し、それを実行する Python を見つけます ( py -3
最初は Windows、他の場所では python3、いずれにしても 3.9 以降)。 Pythonにはまだあります
PATH 上にあること — npm はスクリプトを配信し、 roost を PATH 上に置きますが、そうではありません。
通訳を連れてくる。 PyPI と同様に、裸の名前 roost はすでに取得されているため、
パッケージは roost-top 、コマンドはプレーン roost です。
または、ファイルを取得するだけです。これは 1 つのスクリプトであり、stdlib のみであり、依存関係はありません。
カール -o roost https://raw.githubusercontent.com/gmhoward9289-ops/roost/main/roost.py
chmod +x roost && ./roost
ウイングゲットの場合:
winget インストール gmhoward9289 - ops.roost
これにより、凍結された Windows 実行可能ファイルがインストールされます。Python は必要ありません。それは構築されており、
roost-<version>-windows-x64.zip とともにタグ付けされたすべてのリリースに添付されます。
リリース ページで、
ジッパーを直接掴んだほうがいいでしょう。
または、そのようなことを行わずに、 roost.py を保存して実行します。 .PY は PATHEXT にあります。
したがって、 roost.py は PATH 上のどこからでも動作します。そのルートとその上の npm
どちらも PATH 上に Python インタープリターが必要です。 wingetのインストールはそれです
冷凍されたexeが出荷されるため、そうではありません。
man ページ ( man roost ) は、Homebrew および .deb のインストールに同梱されています。 pipx
install はそれを venv 自身の share/man の下に置きますが、これはデフォルトにはありません
マンパス ;それをその場で読んでください
man "$(pipx 環境 --value PIPX_LOCAL_VENVS)/roost-top/share/man/man1/roost.1" 。
ねぐらのライブ、毎秒リフレッシュ
roost -w 5 更新が遅くなる
roost -1 1 フレーム、その後終了
roost --json 結合レコード、パイプ用
実行中: 今すぐスペースを更新 · アドバイス パネル · サブエージェント パネル · m ローカル モデル パネル · u 使用状況パネル · g ゲートウェイ パネル · r リモート パネル

・hとか？私は何を見ているのですか · インタラクティブに腕を立てます · Q 終了します
The FLOW column is a sparkline of each session's recent token throughput —
更新ごとのコンテキストの増加。最も忙しい瞬間に正規化され、最も新しい時点で
そうです。 Ａ．流れのないサンプルです。ランプ :-=+*# はアクティビティを増加させています。
わざとASCIIにしてあります（Windowsの楷書文字化け）
コンソール)、空で開始します: 履歴は roost の開始時に始まり、何も持続しません。
サブエージェントの CTX 列には 48k/200k、つまりウィンドウ上にロードされたトークンが表示されます。
ウィンドウが推測されます (観察された使用量が適合する最小の標準層)。
トランスクリプトにはセッションが開かれたウィンドウが記録されていないためです。
と。 AGENT 列には、エージェントのタイプ ( Explore/a812a ) が表示されます。
parent has recorded it — which only happens when the agent finishes, so a
まだ実行中のエージェントには 16 進 ID が表示されます。
実験的で、デフォルトではオフになっています。私はインタラクティブモードを武装します -
cursor, x , y , and the EXPERIMENTAL marker in the top-right corner all
それは、x が実際のプロセスを終了することを意味します。を読む
ダッシュボードが危険な部分になったことは一度もありません。私は、を描画する 1 つのキーです
ライン。もう一度押すと解除されます。カーソルが下がり、x / y / j / k が停止します。
もう一度押すまで応答します。
準備完了したら、j / k (または矢印キー) でカーソルを上げます。上げると拡大します
QUIET グループ。何時間もアイドル状態になるセッションはまさにスイープに相当するため
を探していますが、折りたたまれている間はアクセスできません。
これから行動することが分かっている場合は、roost --interactive で準備を始めてください。
すぐにセッション。
x はプロセスを終了します。 It does not compact, save, or otherwise negotiate with the
セッション - 実行中のクロード コードへのローカル制御チャネルはありません
session なので、セッションの外からはこれ以上穏やかなものはありません。 Unix では、これは
シグターム

そしてセッションは自然に終了します。 Windows にはありません
クロスプロセスに相当するため、TerminateProcess のハードキルとなります。成績証明書
は一度に 1 ターンずつ書き込まれるため、失われるのは最大でも実行中のターンです。
どちらのキーも、押したときに画面上にあった行オブジェクトに作用します。
その後再解決されたインデックス上で。セッションが進むにつれてフレーム間で行の順序が変更されます
静かで、フレームを超えて存続するインデックスは、最終的には間違ったインデックスを停止します。
roost は、自身のプロセスやその親プロセスを停止することを拒否し、内部から実行します。
それがポイントされているセッションで、カーソルはあなたのセッションを所有する行に着くことができます。
ターミナル。
一度に開くパネルは 1 つだけです: a 、 s 、 m 、 u 、 g 、 r 、および h の間で反転します。
スタックするのではなく、アドバイス、サブエージェント、ローカル モデル、使用法、ゲートウェイ、リモート、ヘルプ。 24 回のセッションで
画面上では、積み重ねられた 2 番目のパネルが端末の底部の下に着地します。
キーが機能しないのと区別がつきません。フレームも同じ理由で
静かに切り詰めるのではなく、... N 行を下に追加します。
それとも？ HELP パネルを開きます。キーバインドの参照ではありません (フッターのヒント)
すでにキーがリストされています）が、各画面に何が表示されているかを 1 行ずつまとめています。
表示の意味: インフラ、ワーカー、サブエージェント、アドバイス、ローカルモデル、使用状況、
ゲートウェイ、リモート。
roost は非常に小さいので、これがマニュアル全体になります。
INFRA 行には、Ollama が現在 VRAM に常駐しているもののみが表示されます。
ollama:11434 up qwen-coder-16k:latest (5.5 GB) — と書かれているため
/api/ps 。インストールされているがアイドル状態のモデルはそのラインから完全に外れます。
これは、「現在ロードされていません」ではなく、「roost がそれを認識していません」と解釈します。
全体像を表示するには m を押してください: オラマのリストが知っているすべてのモデル、それぞれ
ディスク サイズ、常駐、ロード時の VRAM、Ollama までの時間を示す行
それを降ろします。 /api/tags と読み取られます

/api/ps と統合されました。オラマがいないなら
実行すると、パネルには何も表示されずにそのように表示されます。
u は USAGE パネルを開きます: 過去 1 週間のモデルごとの 1 日あたりのトークンの集計
ディスク上のトランスクリプトから (各アシスタント ターンからの入力 + 出力、キャッシュ)
読み取りは意図的に除外されます。読み取りは別の方法で請求され、制限されます。
それらをカウントすると、変更されていないコンテキストの再読み取りで数が膨大になってしまいます)。
USAGE 観察されたトランスクリプト トークン (入力 + 出力) -- 人間のメーターではなく推定値
2026-08-02 2.9M opus-5 1.6M、sonnet-5 965k、fable-5 253k <- 今日
2026-08-01 9.1M opus-5 4.5M、fable-5 3.3M、sonnet-5 984k
2026-07-29 190 万 gemma4-32k (ローカル) 530 万、opus-5 190 万
今日は 290 万 | 7 日 2,480 万クラウド / 6,000 万予算 (41%)
このパネルには 2 つの正直さのルールが組み込まれています。まず、これは概算ではなく、
人間測定器 — 実際の計画を記録するローカル ファイルや API はありません
レート制限残高のため、予算は自分で設定した数値になります。
export ROOST_WEEKLY_BUDGET=60M # または 850k、またはプレーン トークン数
Claude Code 内で /usage を 1 回実行し、表示される内容と一致する番号を選択します。
そしてそれ以降、パネルはあなたの火傷を追跡します。設定を解除すると、集計結果が表示されます
予算の端数なしで。第二に、ローカルモデルは何でも無料です
クロードモデル名なし (Ollama via LiteLLM)

[切り捨てられた]

## Original Extract

top for Claude Code - live sessions, models, context, and the subagents they spawn - gmhoward9289-ops/roost

GitHub - gmhoward9289-ops/roost: top for Claude Code - live sessions, models, context, and the subagents they spawn · GitHub
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
gmhoward9289-ops
/
roost
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
69 Commits 69 Commits Folders and files
.github .github bin bin demo demo docs docs packaging packaging scripts scripts tests tests .gitattributes .gitattributes .gitignore .gitignore .release-please-manifest.json .release-please-manifest.json .repo-visibility .repo-visibility CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md llms.txt llms.txt package.json package.json pyproject.toml pyproject.toml release-please-config.json release-please-config.json roost.1 roost.1 roost.py roost.py View all files Repository files navigation
top for Claude Code. Every live session, the model it is on, how much context
it has burned — and, unlike anything else, the subagents it spawned .
One file, no dependencies, Python 3.9+. Runs on macOS, Linux and Windows.
The short ambient loop below is the same program, idling:
WORKER MODEL CTX TOKENS TREND FLOW IDLE TASK
NEAR LIMIT
demo-a1 opus-5 85% 170k +4k ..:-=+#+=-.. 12s refactor the parser
PARKED + COSTLY
demo-b2 opus-5 61% 122k +2k .............. 4h10m audit the build scripts
WORKING NOW
demo-c3 fable-5 22% 44k +22k ...:=+*#+ 3s add integration tests
STARTING
demo-d4 - - - - -
QUIET (4) demo-e5 . demo-f6 . demo-g7 . demo-h8
8 worker(s) | 1.2M held | 61k last 8 turns | 1 near limit | fable-5, opus-5
SUBAGENTS
STATE AGENT MODEL CTX IDLE TASK
working Explore/a812a opus-5 66k/200k 2s survey the config loaders
idle adaffaba4b sonnet-5 484k/1M 1h22m draft the migration notes
2 subagent(s), 1 working
INFRA ollama:11434 up qwen2.5-coder:14b (9.2 GB) litellm:4000 up openwebui:8080 DOWN
Sessions are grouped by what it costs to ignore them, not by size: NEAR LIMIT
is about to stop working, PARKED + COSTLY bills its whole context on the next
turn, and everything quiet collapses to a single line.
TREND is context added over the session's last few turns. It is an amount and
not a sparkline because context inside a session only ever rises — it drops
solely on /compact — so a shape would draw the same ramp on every row. CTX
says how full a session is; TREND says how fast it is filling, which is how
demo-c3 above shows as the one to watch at 22%. It is read back out of the
transcript rather than accumulated while roost runs, so it is there on the first
frame and under --once . That is also what separates it from FLOW alongside
it: FLOW sparks throughput sampled since roost started and begins empty,
TREND is an amount already in the file when roost opens it.
The last line totals the fleet. Context held is what a sweep would reclaim.
brew install gmhoward9289-ops/tap/roost
Debian and Ubuntu, from the repo — this also gets you apt upgrade :
curl -fsSL https://gmhoward9289-ops.github.io/roost/roost-archive-keyring.asc | sudo gpg --dearmor -o /usr/share/keyrings/roost-archive-keyring.gpg
echo " deb [signed-by=/usr/share/keyrings/roost-archive-keyring.gpg] https://gmhoward9289-ops.github.io/roost stable main " | sudo tee /etc/apt/sources.list.d/roost.list
sudo apt update
sudo apt install roost
Or, without adding a repo — grab roost_<version>_all.deb from the
latest release
and install that file directly, e.g.:
sudo apt install ./roost_ < version > _all.deb
Both resolve python3 exactly as a normal repo install would. The repo is
signed with a dedicated key (not tied to any personal identity); its public
half is packaging/apt/pubkey.asc in this repo, published unchanged as
roost-archive-keyring.asc above.
pipx install roost-top
PyPI holds the bare name roost in reserve — a prior project's name, retained
after deletion — so the package is roost-top ; the command it installs is
plain roost . Installing straight from the repo skips the index entirely:
pipx install git+https://github.com/gmhoward9289-ops/roost
With npm — the one channel that gives Windows a real roost command:
npm install -g roost-top
Or run it without installing: npx roost-top . The npm package is a wrapper, not
a port: it ships the same roost.py and finds a Python to run it with ( py -3
first on Windows, python3 elsewhere, 3.9 or newer either way). Python still has
to be on PATH — npm delivers the script and puts roost on PATH , it does not
bring an interpreter. As on PyPI, the bare name roost was already taken, so the
package is roost-top and the command is plain roost .
Or just take the file. It is one script, stdlib only, no dependencies:
curl -o roost https://raw.githubusercontent.com/gmhoward9289-ops/roost/main/roost.py
chmod +x roost && ./roost
With winget:
winget install gmhoward9289 - ops.roost
That installs a frozen Windows executable — no Python required. It's built and
attached to every tagged release alongside a roost-<version>-windows-x64.zip
on the Releases page , if
you'd rather grab the zip directly.
Or, without any of that: save roost.py and run it — .PY is in PATHEXT ,
so roost.py works from anywhere on PATH . That route and the npm one above
both still need a Python interpreter on PATH ; the winget install is the one
that doesn't, since it ships a frozen exe.
The man page ( man roost ) ships with the Homebrew and .deb installs. A pipx
install puts it under the venv's own share/man , which is not on the default
MANPATH ; read it in place with
man "$(pipx environment --value PIPX_LOCAL_VENVS)/roost-top/share/man/man1/roost.1" .
roost live, refreshing every second
roost -w 5 slower refresh
roost -1 one frame, then exit
roost --json joined records, for piping
While running: space refresh now · a advice panel · s subagents panel · m local models panel · u usage panel · g gateway panel · r remote panel · h or ? what am I looking at · i arm interactive · q quit
The FLOW column is a sparkline of each session's recent token throughput —
context growth per refresh, normalised to its own busiest moment, newest at the
right. A . is a sample with no flow; the ramp :-=+*# is increasing activity.
It is ASCII on purpose (block-drawing characters mojibake in the Windows
console) and starts empty: history begins when roost starts, nothing persists.
The subagent CTX column reads 48k/200k — tokens loaded over the window.
The window is inferred (the smallest standard tier the observed usage fits in),
because nothing in a transcript records which window the session was opened
with. The AGENT column shows the agent's type ( Explore/a812a ) once the
parent has recorded it — which only happens when the agent finishes, so a
still-running agent shows its hex id.
Experimental, and off by default. i arms interactive mode — the
cursor, x , y , and the EXPERIMENTAL marker in the top-right corner all
come alive together, and it means it: x ends a real process. Reading the
dashboard has never been the risky half; i is the one key that draws the
line. Press it again to disarm — the cursor drops and x / y / j / k stop
responding until you press it once more.
Once armed, j / k (or the arrow keys) raise a cursor. Raising it expands the
QUIET group, because a session idle for hours is exactly what a sweep is
looking for and it is unreachable while collapsed.
Start already armed with roost --interactive if you know you'll be acting on
a session right away.
x ends a process. It does not compact, save, or otherwise negotiate with the
session — there is no local control channel into a running Claude Code
session , so nothing gentler is available from outside it. On Unix that is a
SIGTERM and the session exits on its own terms; on Windows there is no
cross-process equivalent, so it is a TerminateProcess hard kill. Transcripts
are written a turn at a time, so at most an in-flight turn is lost.
Both keys act on the row object that was on screen when you pressed them, never
on an index re-resolved afterwards. Rows reorder between frames as sessions go
quiet, and an index that outlived its frame would eventually stop the wrong one.
roost refuses to stop its own process or its parent — run it from inside the
session it is pointed at and the cursor can land on the row that owns your
terminal.
Only one panel is open at a time: a , s , m , u , g , r , and h flip between
ADVICE, SUBAGENTS, LOCAL MODELS, USAGE, GATEWAY, REMOTE, and HELP rather than stacking. With two dozen sessions
on screen a stacked second panel lands below the bottom of the terminal, which
is indistinguishable from the key not working. For the same reason the frame
now says ... N more line(s) below instead of quietly truncating.
h or ? opens a HELP panel — not a keybinding reference (the footer hint
already lists the keys), but a one-line-each rundown of what each screen on
the display means: INFRA, WORKERS, SUBAGENTS, ADVICE, LOCAL MODELS, USAGE,
GATEWAY, REMOTE.
roost is small enough that this is the whole manual.
The INFRA line only ever shows what Ollama currently has resident in VRAM —
ollama:11434 up qwen-coder-16k:latest (5.5 GB) — because it reads
/api/ps . A model that is installed but idle drops out of that line entirely,
which reads as "roost doesn't see it" rather than "it isn't loaded right now."
Press m for the full picture: every model ollama list knows about, each
row showing disk size, residency, VRAM when loaded, and how long until Ollama
unloads it. It reads /api/tags merged with /api/ps ; if Ollama isn't
running, the panel says so instead of showing nothing.
u opens the USAGE panel: tokens per day per model over the last week, tallied
from the transcripts on disk (input + output from each assistant turn; cache
reads are excluded on purpose — they are billed and limited differently, and
counting them would swamp the number with re-reads of unchanged context).
USAGE observed transcript tokens (input+output) -- an estimate, not the Anthropic meter
2026-08-02 2.9M opus-5 1.6M, sonnet-5 965k, fable-5 253k <- today
2026-08-01 9.1M opus-5 4.5M, fable-5 3.3M, sonnet-5 984k
2026-07-29 1.9M gemma4-32k (local) 5.3M, opus-5 1.9M
today 2.9M | 7d 24.8M cloud / 60.0M budget (41%)
Two honesty rules bake into this panel. First, it is an estimate, not the
Anthropic meter — there is no local file or API that records a plan's real
rate-limit balance, so the budget is a number you set yourself:
export ROOST_WEEKLY_BUDGET=60M # or 850k, or a plain token count
Run /usage inside Claude Code once, pick a number that matches what it shows,
and the panel tracks your burn against it from then on. Unset, the tallies show
without the budget fraction. Second, local models are free — anything
without a claude- model name (Ollama via LiteLLM

[truncated]
