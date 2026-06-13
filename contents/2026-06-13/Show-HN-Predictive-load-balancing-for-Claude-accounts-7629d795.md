---
source: "https://github.com/yasyf/cc-pool"
hn_url: "https://news.ycombinator.com/item?id=48515767"
title: "Show HN: Predictive load balancing for Claude accounts"
article_title: "GitHub - yasyf/cc-pool: Load-balancing for `claude`: CLAUDE_CONFIG_DIR=$(ccp select) claude · GitHub"
author: "yasyfm"
captured_at: "2026-06-13T12:43:50Z"
capture_tool: "hn-digest"
hn_id: 48515767
score: 1
comments: 0
posted_at: "2026-06-13T10:37:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Predictive load balancing for Claude accounts

- HN: [48515767](https://news.ycombinator.com/item?id=48515767)
- Source: [github.com](https://github.com/yasyf/cc-pool)
- Score: 1
- Comments: 0
- Posted: 2026-06-13T10:37:34Z

## Translation

タイトル: HN の表示: クロード アカウントの予測負荷分散
記事のタイトル: GitHub - yasyf/cc-pool: `claude` の負荷分散: CLAUDE_CONFIG_DIR=$(ccp select) クロード · GitHub
説明: `claude` のロードバランシング: CLAUDE_CONFIG_DIR=$(ccp select) claude - yasyf/cc-pool

記事本文:
GitHub - yasyf/cc-pool: `claude` の負荷分散: CLAUDE_CONFIG_DIR=$(ccp select) クロード · GitHub
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
ヤシフ
/
cc-プール
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
112 コミット 112 コミット .claude .claude .github/ workflows .github/ workflows .superset .superset Casks Casks Formula Formula cmd/ cc-pool cmd/ cc-pool docs docs 内部 内部 launchd launchd ウィジェット ウィジェット .gitignore .gitignore .mcp.json .mcp.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md STYLEGUIDE.md STYLEGUIDE.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイル ナビゲーション
複数の Claude サブスクリプションにわたる予測的なセッション開始時の負荷分散 (macOS 向け)。
いくつかの Claude Max/Pro サブスクリプションを実行すると、毎週同じ壁にぶつかります。
1 つのアカウントは 5 時間または 1 週間の制限に固定されている一方で、別のアカウントはアイドル状態のままです。
cc-pool は、選択された最も空のアカウントですべてのクロード コード セッションを起動します。
セッション開始前の 5 時間/7 日間のライブ使用量から — 予測的ではなく予測的
反応的な。リクエスト パスにプロキシはなく、手動切り替えも待機もありません。
レート制限エラーにより、選択が間違っていたことがわかります。
セットアップ後、claude は ccp run のエイリアスとなり、プーリングが消えます
背景に。 2 つの厳格な保証が設計を保証します: plain claude
~/.claude はそのまま動作し続けます — プールは決してログアウトできません — そして
シークレットは macOS キーチェーンに残り、cc-pool のデータベースには決して残りません。なんと
worksに詳細が載っています。
醸造タップ yasyf/cc-pool https://github.com/yasyf/cc-pool
醸造インストール yasyf/cc-pool/cc-pool
macOSのみ。バイナリは、ccp シンボリックリンクを使用して cc-pool としてインストールされます。デフォルト
ビルドは純粋な Go です。
2 つのサブスクリプションをプールし、最も空いたサブスクリプションでクロードを約 5 回以内に起動します
分。
1. ccp を実行します。空のプールでは、それぞれのログイン手順が説明されます。
サブスクリプション — すべてのアカウントは独自の claude /login を取得します。
$ccp
✓ このマシンに cc-pool をセットアップします。
H

ログインしますか?
> この端末で今すぐログインしてください
# クロードはここで /login フローを開きます。それを完了すると、CCPはあなたのためにクロードを閉じます
このアカウントの名前 (オプション)
> work@example.com
✓ work@example.com を追加しました。
別のアカウントを追加しますか?
> はい
# 2 回目のログインも同じ方法で実行されます
✓ Personal@example.com を追加しました。
プールには 2 つのアカウントが存在します。
最も空のアカウントでクロードを起動します。
CCP実行
常に最も空のアカウントで起動するように「claude」をラップしますか?
> はい
✓ ラップされた `claude` — エイリアスを ~/.zshrc に追加しました。
今すぐシェルを再起動するか、「source ~/.zshrc」を実行して使用してください。
プレーン ~/.claude に対して `command claude` を実行します。
2. プールを確認します。 ccp status --plain はテーブルを 1 回出力します。裸の
ccp status (または、アカウントが存在する場合は裸の ccp ) により、ライブ TUI が開きます。
アカウントごとのスコアの内訳:
$ ccp ステータス --プレーン
アカウントスコア 5時間使用 7日使用 ライブリセット
▸ work@example.com 68.1 22% 46% 0 6:00 PM
Personal@example.com 34.8 61% 70% 0 4:30 PM
▸ = 次のピック · スコアが高い = 空っぽ · 5 時間/7 日 = 使用率
午後3時58分更新
3. 起動します。クロードを実行します (現在はラップされています)。標準エラー出力での選択を発表します。
次に、実際のクロードを実行します。以前に cc-pool がプロセス ツリーから削除されます。
Claude Code は最初のフレームを描画します。
$クロード
選択した work@example.com · 5 時間 22% 使用 · 7 日 46% 使用
[選択された行] には、ステップ 2 で ▸ とマークされたアカウントの名前が付けられています。つまり、一致するのは次のとおりです。
あなたの確認。 From here on, every claude lands on whichever account
最も大きなヘッドルームを持っています。
クイックスタートは、claude をエイリアス claude='ccp run' でラップします。コマンドクロード
それをバイパスします — ~/.claude 上のプレーンなクロード、1 キーストロークの距離です。を好む
クロードをそのままにしておきますか？プロンプトを拒否します (または ccp add --no-alias を渡します)
自分の名前を選択してください:
エイリアス cl= ' ccp run '
引数を渡す
ccp run はすべての引数をそのままクロードに転送します — no -- 区切り文字
必要です。 f 付きの裸の ccp

lags は同じことの短縮形です。
ccp run --resume
ccp -p " このリポジトリを要約する " # `ccp run -p ...` に自動変換します
強制と固定
CCP_ACCOUNT=2 ccp run は自動選択ではなくアカウント 2 を強制します。繰り返し
同じディレクトリから起動し、プロンプトキャッシュ用に 1 つのアカウントに固執します
継続性。これらは、代わりに work@example.com (固定) を再利用することを発表しています。
選択されました。
選択は、ほぼリセットされたウィンドウを書き込むことを拒否します: 使い果たされたアカウント
どのアカウントにも余裕がある間は、プランウィンドウが選択されることはありません。
ccp select --wait はブロックされるまでブロックします。プール全体が使い果たされた場合、
起動は最も悪いアカウントにフォールバックし、stderr で大声で警告します。
追加使用が有効になっている場合、そのセッションは従量課金制クレジットを請求します。または
ウィンドウがリセットされるまでレート制限が適用されます。
ccp select は、選択した構成ディレクトリを標準出力に出力するだけで、他には何も出力しません。を設定します。
プラグイン ルートも同様であるため、セッションは正規のプラグイン パスを共有ディレクトリに書き込みます。
~/.claude/plugins (ccp run と ccp env の両方がこれを行います):
CLAUDE_CODE_PLUGIN_CACHE_DIR= " $HOME /.claude/plugins " CLAUDE_CONFIG_DIR= $( ccp select ) クロード
仕組み
~/.claude は移動されず、プール アカウントとして登録されることもありません。それ
正規の設定ディレクトリのままなので、プレーンクロードは正確に動作し続けます
before, and is the shared base every pooled account mirrors.の
プールは、プレーン クロードの資格情報やログイン ID には決して触れません。どのアカウントも、
メインのサブスクリプションを含め、独自のクロード /login で結合するため、
トークン チェーンはプレーン クロードのトークン チェーンから完全に独立しています。
アカウントごとに 1 つの実際の構成ディレクトリ
Claude Code は、設定ディレクトリごとにキーチェーン認証情報の名前空間を設定します: デフォルト
~/.claude は項目 Claude Code-credentials を使用します。カスタム
CLAUDE_CONFIG_DIR は、接尾辞付き項目 Claude Code-credentials-<hash> を取得します。
cc-pool は、各アカウントに実際の一意のディレクトリを与えます (

~/.cc-pool/accounts/acct-NN)
したがって、それぞれが独自のキーチェーン項目、独自の独立した OAuth 付与 (独自の
リフレッシュ トークン チェーン)、API ではなく独自のサブスクリプションで実行されます。
請求。各アカウント ディレクトリには ~/.claude.json のコピーがシードされます。
ID が削除される (アカウント自身のログインがその ID を書き込む) ため、プールされる
セッションは設定、MCP サーバー、プロジェクトごとのツールの承認を継承します。
初回実行オンボーディングを実行する代わりに。
各アカウント ディレクトリには、 ~/.claude (projects/、skills/、など) のすべてが表示されます。
plan/ 、 settings.json 、history.jsonl 、ロット — 書き込み渡し付き
すぐに戻るため、すべてのセッションが同じワークスペースとプランモードの計画を共有します
プールされたセッション間で持続します。 2 つのプロバイダー:
シンボリックリンク (デフォルト、依存関係なし): の各トップレベルエントリにシンボリックリンクを作成します。
~/.claude をアカウント ディレクトリに追加します。新しいトップレベルのエントリがピックアップされます
起動時、デーモン、および ccp Doctor --fix によって自動的に実行されます。
ヒューズ (オプション、ライブ ミラー): 経由で取り付けられたパススルー ミラー
ヒューズ-t — kext なし、としてマウント
あなたは root ではなく、切り離された cc-pool mount-holder プロセスによってホストされているため、
デーモンの再起動とアップグレードによってライブ セッションのマウントが妨げられることはありません。が必要です
-tags は、ビルド (cgo) と 1 回限りのネットワーク ボリューム プライバシー許可を融合します。
いくつかのエントリは共有されずにアカウントごとに残ります: daemon/ と ide/
(クロードの PID キー付きスーパーバイザーと IDE ロック/ソケット ファイルは衝突します。
同時セッション間)、バックアップ/ (各アカウントのローテーション バックアップ)
.claude.json )、ID ファイルと認証情報ファイル .claude.json および
.credentials.json 、 .last-update-result.json (インスタンスローカル自動更新)
state)、およびremote-settings.json (クロードのサブスクリプションごとにキャッシュされたもの)
設定）。
ただし、アカウントごとの .claude.json は設定のフォークを意味するものではありません。共有可能です。
トップレベルのキー — ID を除くすべて、p ごと

プロジェクトの状態と起動
カウンター — ~/.claude.json からプールされたセッションに流れるため、設定は
バニラ・クロードの変化はすべてのアカウントに届きます。 1 つの注意点:
デフォルトのシンボリックリンクオーバーレイでは、フローは一方向です（起動時にマージされ、ベースが優先します）。
そのため、プールされたセッション内で共有可能な設定を変更すると、次回のセッションで元に戻ります。
起動 — バニラ クロードで共有設定を管理するか、ヒューズ オーバーレイを使用します。
whose live merged view writes shareable changes back to ~/.claude.json
(双方向)。
ベースライン (ウィンドウがリセットから程遠いときの正確な値) は次のとおりです。
スコア = 0.70・(100−util_5h) + 0.25・(100−util_7d)
− 2・active_sessions − 100・rate_limited − 20・stale_or_refresh_failed
3 つの条件により、ランキングは端付近で正直に保たれます: 差し迫ったリセット
ウィンドウがリセットされるまでの時間 (90% が使用されたウィンドウ) に比例してクレジットを獲得します。
10 分以内にリセットするとランクが上がります (ランクは下がりません)。頭上空間の低いバリア
ほぼ使い果たされた 7 日間のウィンドウが 5 時間のヘッドルームによって隠蔽されるのを防ぎます。
バーンレート条件は、アクティブに排出されているアカウントをランクダウンします。
select は argmax を選択します。使用法はクロード自身の /api/oauth/usage から取得されます。
エンドポイント。
brew services start cc-pool (Homebrew installs) または ccp service install
(ソース ビルド) ユーザー LaunchAgent を実行します — ルート デーモンがあなたのファイルを読み取ることができませんでした
ログインキーチェーン。指数バックオフを使用して、約 3 分ごとに使用状況をポーリングします。
アイドル状態のアカウントのトークンが期限切れになる前に更新します (チェックアウトされたセッション)
独自のリフレッシュを所有します。デーモンは、ローテーションされたトークンをすべて採用します。
チェックイン）、スコアをキャッシュし、ヒューズ オーバーレイを使用して、
マウントを所有する切り離されたマウント ホルダー (デーモンが再起動してアップグレードされる)
決して彼らの邪魔をしないでください）。 ccp add と ccp init は自動的に開始します。もしそうなら
が実行されていない場合、ccp select が自動生成するか、ライブサンプルを生成します。
シークレットは cc-p に保存されません。

ool のデータベース — macOS キーチェーンは
唯一の秘密の店。
これら 2 つのテーブルは、ユーザーに面する面全体をカバーします。 ccp ヘルプ <コマンド>
コマンドごとに同じものを出力します。
ccp service uninstall # デーモンを停止し、ホルダーをマウントし、ヒューズ オーバーレイをアンマウントします
# (ライブセッションでは拒否; --force オーバーライド)
ccp service uninstall --purge # ...すべてのプールアカウント/ディレクトリ/状態を削除します
醸造アンインストールccプール
~/.claude とその資格情報には決して触れません。
CGO_ENABLED=0 でビルドします。 go build ./cmd/cc-pool ; go test ./... で合格します
ネットワーク、キーチェーン、デーモンはありません。手動によるエンドツーエンドのテスト マトリックスは、
docs/VERIFICATION.md 、リリース履歴は次のとおりです。
CHANGELOG.md 、および AGENTS.md の規約。
PolyForm-Noncommercial-1.0.0 © Yasyf Mohamedali — 非営利使用は無料です。
LICENSE またはライセンスのテキストを参照してください
オンライン。
「claude」のロードバランシング: CLAUDE_CONFIG_DIR=$(ccp select) クロード
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
37
v0.25.0
最新の
2026 年 6 月 12 日
+ 36 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Load-balancing for `claude`: CLAUDE_CONFIG_DIR=$(ccp select) claude - yasyf/cc-pool

GitHub - yasyf/cc-pool: Load-balancing for `claude`: CLAUDE_CONFIG_DIR=$(ccp select) claude · GitHub
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
yasyf
/
cc-pool
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
112 Commits 112 Commits .claude .claude .github/ workflows .github/ workflows .superset .superset Casks Casks Formula Formula cmd/ cc-pool cmd/ cc-pool docs docs internal internal launchd launchd widget widget .gitignore .gitignore .mcp.json .mcp.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md STYLEGUIDE.md STYLEGUIDE.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Predictive, start-of-session load-balancing across multiple Claude subscriptions — for macOS.
Run several Claude Max/Pro subscriptions and you hit the same wall every week:
one account pegged at its 5-hour or weekly limit while another sits idle.
cc-pool launches every Claude Code session on the emptiest account, picked
from live 5-hour / 7-day usage before the session starts — predictive, not
reactive. No proxy in the request path, no manual switching, no waiting for a
rate-limit error to learn you picked wrong.
After setup, claude is an alias for ccp run and the pooling disappears
into the background. Two hard guarantees underwrite the design: plain claude
on ~/.claude keeps working untouched — the pool can never log it out — and
secrets stay in the macOS Keychain, never in cc-pool's database. How it
works has the details.
brew tap yasyf/cc-pool https://github.com/yasyf/cc-pool
brew install yasyf/cc-pool/cc-pool
macOS only. The binary installs as cc-pool with a ccp symlink; the default
build is pure Go.
Pool two subscriptions and launch Claude on the emptiest one, in about five
minutes.
1. Run ccp . On an empty pool it walks you through logging in each
subscription — every account gets its own claude /login :
$ ccp
✓ Set up cc-pool on this machine.
How do you want to log in?
> Log in now, in this terminal
# claude opens its /login flow right here; finish it and ccp closes claude for you
Name for this account (optional)
> work@example.com
✓ Added work@example.com.
Add another account?
> Yes
# the second login runs the same way
✓ Added personal@example.com.
Your pool now has 2 accounts.
Launch Claude on the emptiest account:
ccp run
Wrap `claude` to always launch on the emptiest account?
> Yes
✓ Wrapped `claude` — added an alias to ~/.zshrc.
Restart your shell or run `source ~/.zshrc` to use it now.
Run `command claude` for plain ~/.claude.
2. Check the pool. ccp status --plain prints the table once; bare
ccp status (or bare ccp , now that accounts exist) opens a live TUI with
per-account score breakdowns:
$ ccp status --plain
ACCOUNT SCORE 5h used 7d used LIVE RESETS
▸ work@example.com 68.1 22% 46% 0 6:00 PM
personal@example.com 34.8 61% 70% 0 4:30 PM
▸ = next pick · score higher = emptier · 5h/7d = % used
updated 3:58 PM
3. Launch. Run claude (now wrapped). It announces its pick on stderr,
then execs the real claude — cc-pool is gone from the process tree before
Claude Code draws its first frame:
$ claude
Selected work@example.com · 5h 22% used · 7d 46% used
The Selected line names the account marked ▸ in step 2 — that match is
your verification. From here on, every claude lands on whichever account
has the most headroom.
The quickstart wraps claude with alias claude='ccp run' . command claude
bypasses it — plain claude on ~/.claude , one keystroke away. Prefer to
leave claude untouched? Decline the prompt (or pass ccp add --no-alias )
and pick your own name:
alias cl= ' ccp run '
Passing arguments
ccp run forwards every argument to claude verbatim — no -- separator
needed. Bare ccp with flags is shorthand for the same thing:
ccp run --resume
ccp -p " summarize this repo " # auto-converts to `ccp run -p ...`
Forcing and pinning
CCP_ACCOUNT=2 ccp run forces account 2 instead of auto-selecting. Repeated
launches from the same directory stick to one account for prompt-cache
continuity; those announce Reusing work@example.com (pinned) instead of
Selected .
Selection refuses to burn a nearly-reset window: an account with an exhausted
plan window is never picked while any account has headroom, and
ccp select --wait blocks until one does. If the whole pool is exhausted, the
launch falls back to the least-bad account and warns loudly on stderr —
that session bills pay-as-you-go credits if extra usage is enabled, or
rate-limits until the window resets.
ccp select prints the chosen config dir on stdout and nothing else. Set the
plugin root too, so the session writes canonical plugin paths into the shared
~/.claude/plugins ( ccp run and ccp env both do this for you):
CLAUDE_CODE_PLUGIN_CACHE_DIR= " $HOME /.claude/plugins " CLAUDE_CONFIG_DIR= $( ccp select ) claude
How it works
~/.claude is never moved and never registered as a pool account. It
stays the canonical config dir, so plain claude keeps working exactly as
before, and is the shared base every pooled account mirrors. The
pool never touches plain claude's credential or login identity. Every account,
including your main subscription, joins with its own claude /login , so its
token chain is fully independent of plain claude's.
One real config dir per account
Claude Code namespaces its Keychain credential per config dir : the default
~/.claude uses the item Claude Code-credentials ; a custom
CLAUDE_CONFIG_DIR gets a suffixed item Claude Code-credentials-<hash> .
cc-pool gives each account a real, unique dir ( ~/.cc-pool/accounts/acct-NN )
so each gets its own Keychain item, its own independent OAuth grant (its own
refresh-token chain), and runs on its own subscription — never API
billing. Each account dir is seeded with a copy of your ~/.claude.json with
the identity stripped (the account's own login writes its identity), so pooled
sessions inherit your settings, MCP servers, and per-project tool approvals
instead of running first-run onboarding.
Each account dir presents all of ~/.claude — projects/ , skills/ ,
plans/ , settings.json , history.jsonl , the lot — with writes passing
straight back, so every session shares the same workspace and plan-mode plans
persist across pooled sessions. Two providers:
symlink (default, zero-dependency): symlinks each top-level entry of
~/.claude into the account dir. New top-level entries are picked up
automatically at launch, by the daemon, and by ccp doctor --fix .
fuse (optional, live mirror): a passthrough mirror mounted via
fuse-t — kext-less, mounted as
you, no root — and hosted by a detached cc-pool mount-holder process, so
daemon restarts and upgrades never disturb live sessions' mounts. Requires a
-tags fuse build (cgo) and a one-time Network Volumes privacy grant.
A few entries stay per-account instead of shared: daemon/ and ide/
(Claude's PID-keyed supervisor and IDE lock/socket files, which would collide
across concurrent sessions), backups/ (rotating backups of each account's
.claude.json ), the identity and credential files .claude.json and
.credentials.json , .last-update-result.json (instance-local auto-update
state), and remote-settings.json (claude's cached per-subscription
settings).
Per-account .claude.json doesn't mean settings fork, though: its shareable
top-level keys — everything except identity, per-project state, and startup
counters — flow from ~/.claude.json into pooled sessions, so a setting you
change in vanilla claude reaches every account. One caveat: under the
default symlink overlay the flow is one-way (merged in at launch, base wins),
so changing a shareable setting inside a pooled session reverts at the next
launch — manage shared settings in vanilla claude , or use the fuse overlay,
whose live merged view writes shareable changes back to ~/.claude.json
(two-way).
The baseline — exact when windows are far from a reset — is:
score = 0.70·(100−util_5h) + 0.25·(100−util_7d)
− 2·active_sessions − 100·rate_limited − 20·stale_or_refresh_failed
Three terms keep the ranking honest near the edges: an imminent reset
earns credit in proportion to how soon the window resets (a 90%-used window
resetting in 10 minutes ranks up , not down); a low-headroom barrier
stops a nearly-exhausted 7-day window from being masked by 5-hour headroom;
and a burn-rate term downranks an account being actively drained.
select picks argmax. Usage comes from Claude's own /api/oauth/usage
endpoint.
brew services start cc-pool (Homebrew installs) or ccp service install
(source builds) runs a user LaunchAgent — a root daemon couldn't read your
login Keychain. It polls usage every ~3 min with exponential backoff,
refreshes idle accounts' tokens before they expire (a checked-out session
owns its own refresh; the daemon adopts whatever token it rotated to on
check-in), caches scores, and — with the fuse overlay — supervises the
detached mount holder, which owns the mounts (so daemon restarts and upgrades
never disturb them). ccp add and ccp init start it automatically; if it
isn't running, ccp select auto-spawns it or samples live.
No secrets are ever stored in cc-pool's database — the macOS Keychain is the
only secret store.
These two tables cover the full user-facing surface; ccp help <command>
prints the same per command.
ccp service uninstall # stop the daemon + mount holder, unmount fuse overlays
# (refuses under live sessions; --force overrides)
ccp service uninstall --purge # ...and remove all pool accounts/dirs/state
brew uninstall cc-pool
~/.claude and its credential are never touched.
Build with CGO_ENABLED=0 go build ./cmd/cc-pool ; go test ./... passes with
no network, Keychain, or daemon. The manual end-to-end test matrix lives in
docs/VERIFICATION.md , release history in
CHANGELOG.md , and conventions in AGENTS.md .
PolyForm-Noncommercial-1.0.0 © Yasyf Mohamedali — free for noncommercial use.
See LICENSE or the license text
online .
Load-balancing for `claude`: CLAUDE_CONFIG_DIR=$(ccp select) claude
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
37
v0.25.0
Latest
Jun 12, 2026
+ 36 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
