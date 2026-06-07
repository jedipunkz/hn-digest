---
source: "https://github.com/wavever/CCLimitPing"
hn_url: "https://news.ycombinator.com/item?id=48434613"
title: "Show HN: LimitPing – Keep Claude Code and Codex rate-limit windows continuous"
article_title: "GitHub - wavever/CCLimitPing: Keep your Claude Code / Codex / GLM 5h rate-limit windows back-to-back — auto-pings each provider the moment its window resets. | 在 5 小时限额窗口重置的瞬间自动触发,让 Claude Code / Codex / GLM 的窗口背靠背、不留空档。 · GitHub"
author: "wavever"
captured_at: "2026-06-07T15:35:35Z"
capture_tool: "hn-digest"
hn_id: 48434613
score: 2
comments: 0
posted_at: "2026-06-07T13:26:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LimitPing – Keep Claude Code and Codex rate-limit windows continuous

- HN: [48434613](https://news.ycombinator.com/item?id=48434613)
- Source: [github.com](https://github.com/wavever/CCLimitPing)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T13:26:51Z

## Translation

タイトル: Show HN: LimitPing – クロード コードとコーデックスのレート制限ウィンドウを継続的に維持する
記事のタイトル: GitHub - wavever/CCLimitPing: Claude Code / Codex / GLM 5 時間のレート制限ウィンドウを連続して維持する — ウィンドウがリセットされた瞬間に各プロバイダーに自動 ping を送信します。 · GitHub
説明: Claude Code / Codex / GLM の 5 時間のレート制限ウィンドウを連続して維持します。ウィンドウがリセットされた瞬間に各プロバイダーに自動 ping を送信します。 - wavever/CCLimitPing

記事本文:
GitHub - wavever/CCLimitPing: Claude Code / Codex / GLM 5 時間のレート制限ウィンドウを連続して維持します。ウィンドウがリセットされた瞬間に各プロバイダーに自動 ping を送信します。 | 5 時間制限のクロード ポート再配置の瞬間自動起動で、Claude Code / Codex / GLM のクロード バックバック、空隙のないことを確認します。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロード先

セッションをリフレッシュしてください。
アラートを閉じる
{{ メッセージ }}
ウェイバー
/
CCLimitPing
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .github .github cmd/ 制限 cmd/ 制限内部 内部 .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude コード、コーデックス、GLM (Zhipu / Z.ai コーディング プラン) を保管してください。
レート制限ウィンドウを連続して実行します。
これらのプロバイダーは、5 時間のローリング ウィンドウ (プラス週ごとの上限) に基づいて料金を請求します。
最初のメッセージから 5 時間のウィンドウが始まります。すぐに何も送信しない場合
ウィンドウがリセットされると、そのギャップは無駄になります。次のウィンドウは、次のウィンドウが開始されるまで開始されません。
たまたまそのツールを再び使用してしまい、日常と同期しなくなってしまいます。
制限は各プロバイダーを監視し、5 時間のウィンドウがリセットされた瞬間にプロバイダーを送信します。
次のウィンドウをすぐに開始するための最小限のメッセージ - ウィンドウをそのままにします
継続的かつ予測可能。
クロード ✓ ping されました (6.6 秒)
codex ✓ ping (13.6 秒、16,862 トークン (入力 16,814 / 出力 48)、$0.0098)
ハイライト
アイドル状態のギャップを漂わせるのではなく、5 時間のプロバイダー ウィンドウを継続的に維持します。
あなたのスケジュール。
ゼロクォータのエンドポイントから使用状況を読み取り、
可能な場合は常に公式プロバイダー ツールを使用してください。
Claude Code、Codex、およびオプトイン GLM/Z.ai コーディング プランのモニタリングをサポートします。
ドライラン モード、毎週の制限ガード、リセット バッファ、ローカル設定、
テレメトリーはありません。
カール -fsSL https://raw.githubuserconten

t.com/wavever/CCLimitPing/main/install.sh |しー
構成の初期化を制限する
ping の制限 --dry-run
限界状態
リミットウォッチ
必要に応じて、最初に制限 ping --dry-run または制限ウォッチ --dry-run を使用します。
プロバイダー クォータを消費せずに何が起こるかを検査します。
プロバイダー
読み取り使用量 (ゼロクォータ)
トリガー
認証
クロード・コード
…/api/oauth/usage
インタラクティブなクロード コード CLI
OAuth (キーチェーン / ~/.claude )
コーデックス
…/backend-api/wham/usage
コーデックス実行
OAuth ( ~/.codex/auth.json )
GLM (Zhipu / Z.ai)
…/api/monitor/usage/quota/limit
最小限のチャット完了
APIキー（config / env）
注記
GLM はデフォルトでオフになっており、ライブ プランではまだ検証されていません - を参照してください。
GLM を有効にする前に。
ウォッチは 5 時間のウィンドウがリセットされたことを確認すると、まずすでに実行中の
クロード/コーデックス CLI タスク。実行中の場合、待機と再読み取りの使用を制限します
独自の ping を送信する代わりに、そのタスクの次のモデル リクエストが
新しいウィンドウが自然に起動します。
クロード : を使用して GET https://api.anthropic.com/api/oauth/usage を読み取ります。
macOS キーチェーンからの OAuth トークン ( Claude Code-credentials ) または
~/.claude/.credentials.json 。トリガーには TTY ベースのインタラクティブを使用します
クロード "<プロンプト>" セッションなので、引き続きクロードが開始されます。
ヘッドレス印刷コマンドがエージェントに移動した後のサブスクリプションに基づくウィンドウ
SDK/API クレジット。
Codex : を使用して GET https://chatgpt.com/backend-api/wham/usage を読み取ります。
~/.codex/auth.json からの OAuth トークン。
GLM : GET …/api/monitor/usage/quota/limit を読み取ります (api.z.ai または
open.bigmodel.cn ) コーディング プラン API キーを使用します。 GLM にはスタンドアロン CLI がありません。
したがって、トリガーは直接の最小限のチャット完了です。
…/api/coding/paas/v4/chat/completions ではなく、シェルアウトです。
クロード/コーデックス トークンは公式ツールから再利用されます (個別のログインはありません)。
401 で更新されました。GLM は静的 API キー (c から) を使用します。

onfig または env) — 以下を参照してください。
シップを単一の自己完結型バイナリとして制限します。Go は必要ありません。
一行スクリプト (macOS / Linux):
カール -fsSL https://raw.githubusercontent.com/wavever/CCLimitPing/main/install.sh |しー
適切なビルド済みバイナリを
最新リリースへ
/usr/local/bin (または ~/.local/bin )。 LIMITPING_INSTALL_DIR でオーバーライドします。
アップグレード — インストールされているバイナリを最新リリースに置き換えます。
アップグレードの制限
別名: 制限アップ 、 制限更新 。
アンインストール — インストールされているバイナリと構成/キャッシュを削除します。
アンインストールの制限
別名: rm を制限する、remo を制限する。
制限アンインストール --keep-config を使用して ~/.config/limitping (または
$XDG_CONFIG_HOME/limitping )。
手動ダウンロード — プラットフォームのアーカイブを次の場所から取得します。
リリース ページ ( .tar.gz
macOS/Linux、Windows の場合は .zip):
tar -xzf 制限_darwin_arm64.tar.gz
sudo mv 制限 /usr/local/bin/
Homebrew (macOS / Linux) — brew install wavever/tap/limitping
(Homebrew タップが設定されると機能します — .goreleaser.yaml を参照してください)。
ソースから (開発者、Go 1.25 以降が必要):
github.com/wavever/CCLimitPing/cmd/limitping@latest をインストールしてください
# または、クローンから:
go build -o bin/limitping ./cmd/limitping
有効にする各プロバイダーには、独自の認証情報 (claude / codex CLI) が必要です。
ログイン (Claude / Codex)、またはコーディング プラン API キー (GLM)。
制限設定の初期化 # write ~/.config/limitping/config.toml
ステータスの制限 # 5 時間/週の表示 % + カウントダウンのリセット (エイリアス: s)
ステータスを制限する -v # 生の JSON も出力します
ping の制限 # すべての有効なプロバイダーを今すぐトリガーします (エイリアス: p)
ping クロードを制限 # クロードのみ
ping コーデックスの制限 # コーデックスのみ
ping glm の制限 # GLM のみ
ping を制限 --dry-run # 送信せずにコマンドを表示します
制限監視 # フォアグラウンド デーモン: リセット時に各ウィンドウに ping (

別名: w)
制限監視 claude # 1 つのプロバイダーのみを監視 (claude|codex|glm)
制限 watch --dry-run # ping が送信されない場合にログを記録します
バージョンの制限 # バージョンを出力します (エイリアス: v、ver)
アップグレードの制限 # 最新の GitHub リリースへの更新 (エイリアス: up、update)
制限アンインストール # 制限プラス設定/キャッシュを削除します (エイリアス: rm、削除)
config コマンドには短いエイリアスも使用できます。 c i を制限する
config init と config path の cp を制限します。
制限 --help は、エイリアスをインラインでリストします (例: ping、p )。
ping は正確なコマンド、ライブ タイマー (端末上のスピナー)、
トークンの使用状況 (codex --json またはから解析された、利用可能な場合に消費された ping)
GLM API 応答)、および利用可能な場合は USD コスト:
クロード → クロード --model 俳句 。
クロード ✓ ping されました (6.6 秒)
codex → codex exec --skip-git-repo-check --json -c model_reasoning_effort=low -m gpt-5.4-mini ok
codex ✓ ping (13.6 秒、16,862 トークン (入力 16,814 / 出力 48)、$0.0098)
コストソース:
クロード対話モードは呼び出しごとの機械可読性を公開しません
使用量/コストなので、トークン/コストの接尾辞は表示されません。
Codex (サブスクリプション) は米ドルのコストを返さないため、CodexBar/ccusage のように
— 同等の API レート コストを次の式から導き出します。
LiteLLM 価格データセット
(コスト = キャッシュされていない入力 × 入力 + キャッシュされた入力 × キャッシュ読み取り + 出力 × 出力)。
データセットは ~/.config/limitping/litellm_prices.json (24 時間 TTL) にキャッシュされます。
モデルエイリアス/日付サフィックスフォールバックを使用します。 [codex].model をそのように設定する必要があります
レートを調べることができます。
GLM はプロンプトごとのサブスクリプションであるため、通話ごとの USD コストは表示されません。
トークンの数。
クロードのトリガーは依然として少量のクロード サブスクリプション クォータを消費します。
ただし、対話型 CLI では、ping ごとの正確なトークン数は公開されません。
クロード
5時間[█████░░░]

░░] 51.0% が 3 時間 14 分でリセット (日 00:10)
毎週 [█████░░░░░] 54.0% が 7 時間 04 分にリセット (日曜日 04:00)
コーデックス（プラス）
5時間 [██░░░░░░░░] 3時間15分で24.0%がリセット (日 00:11)
毎週 [████░░░░░░] 111 時間 57 分で 37.0% がリセット (木曜日 12:53)
構成
~/.config/limitping/config.toml ( $XDG_CONFIG_HOME を優先):
Weekly_threshold = 0.99 # 毎週の使用量 >= これ (0..1) の場合、毎週リセットされるまで ping をスキップします
replace_buffer = " 10s " # リセット後、ping を送信する前にこの時間待機します (ロールオーバーを保証します)
通知 = true # ping/スキップ/失敗時の macOS 通知
[クロード]
有効 = true
プロンプト = " . "
model = " haiku " # 最安層;トリガーには SOTA モデルは必要ありません
extra_args = [] # 追加のクロード CLI 引数; print/headless-only フラグは無視されます
align_start = " " # 最初のウィンドウのオプションの RFC3339 アンカー。空 = できるだけ早く開始
[コーデックス]
有効 = true
プロンプト = " OK "
model = " gpt-5.4-mini " # トリガー用の最も安価な Codex モデル
reasoning_effort = " low " # web_search/image_gen ツールが有効な場合、「minimal」は拒否されます
extra_args = []
align_start = " "
[ グラム ]
有効 = false # オプトイン: プラン + API キーを取得したら有効にします
プロンプト = " OK "
model = " glm-4.6 " # 最安の標準モデル;フラッグシップ GLM-5/5.1 は乗数がかかります
プラットフォーム = "グローバル" # "グローバル" = api.z.ai、"cn" = open.bigmodel.cn (Zhipu)
api_key = " " # empty = $ZAI_API_KEY (グローバル) / $ZHIPU_API_KEY (cn) から読み取ります
align_start = " "
最上位のキー:
Weekly_threshold — 週単位のウィンドウがこれ以上になると監視が停止します
ping を送信し、毎週のリセットを待ちます (使用可能なクレジットが存在しない場合)。
reset_buffer — ウィンドウのリセット時間の後に待機する時間
ピーンと鳴っているので、窓は間違いなく横転しています。
align_start (プロバイダーごと) — ウィンドウのフェーズを固定します: に設定します。
未来R

FC3339 それまで最初の ping を遅らせる時間です。その後窓
~5 時間ごとに自動的にチェーンします。
ウィンドウのトリガーはモデルに依存しません。請求可能なリクエストはすべて開始されます。
5 時間クロック — したがって、ping は各プロバイダーの最も安価なモデルを使用して、消費量を最小限に抑えます。
あなたの予算:
クロード → 俳句 : 毎週の別個の Opus バケットも避けます。
Codex → gpt-5.4-mini : ミニのバリアント (~/.codex/models_cache.json を参照)
プランの内容に応じて)。
GLM → glm-4.6 : 標準モデル。フラッグシップ GLM-5/5.1 の控除枠を
倍率は 2 ～ 3 倍なので、単なるトリガーとして使用するのは避けてください。
Claude/Codex は実行時にモデルごとの価格を公開しません (Anthropic のローカルコスト)
キャッシュは空です。 Codex のモデル キャッシュには価格フィールドがないため)、最も安いモデルは
ライブ価格検索ではなく、賢明なデフォルトです。プロバイダーごとにモデルを上書きする
お好みであれば。
GLM（Zhipu / Z.aiコーディングプラン）
GLM は、Claude/Codex と同じ 5 時間 + 週の構造を使用しますが、次の 2 つの点が異なります。
違います:
Auth は OAuth ではなく、静的な API キーです。 [glm].api_key に入れるか、そのままにしておきます
空にして、ZAI_API_KEY (グローバル) / ZHIPU_API_KEY (CN) をエクスポートします。使用状況の読み取りがヒットしました
…/api/monitor/usage/quota/limit ;キーは Authorization ヘッダーに入れられます
Bearer プレフィックスなし (エンドポイントはこれを予期します)。
GLM にはスタンドアロン CLI がないため、トリガーは直接 API 呼び出しです。それ
ワントークンのチャット完了を …/api/coding/paas/v4/chat/completions に送信します。
ライブプランでは未検証。 GLM は d によってオフになります

[切り捨てられた]

## Original Extract

Keep your Claude Code / Codex / GLM 5h rate-limit windows back-to-back — auto-pings each provider the moment its window resets. | 在 5 小时限额窗口重置的瞬间自动触发,让 Claude Code / Codex / GLM 的窗口背靠背、不留空档。 - wavever/CCLimitPing

GitHub - wavever/CCLimitPing: Keep your Claude Code / Codex / GLM 5h rate-limit windows back-to-back — auto-pings each provider the moment its window resets. | 在 5 小时限额窗口重置的瞬间自动触发,让 Claude Code / Codex / GLM 的窗口背靠背、不留空档。 · GitHub
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
wavever
/
CCLimitPing
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .github .github cmd/ limitping cmd/ limitping internal internal .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum install.sh install.sh View all files Repository files navigation
Keep your Claude Code , Codex , and GLM (Zhipu / Z.ai Coding Plan)
rate-limit windows back-to-back.
These providers bill on a 5-hour rolling window (plus a weekly cap), and the
5h window starts on your first message . If you don't send anything right when
a window resets, that gap is wasted — the next window only starts whenever you
happen to use the tool again, drifting out of sync with your day.
limitping watches each provider and, the moment a 5h window resets, sends one
minimal message to start the next window immediately — so your windows stay
continuous and predictable.
claude ✓ pinged (6.6s)
codex ✓ pinged (13.6s, 16,862 tok (in 16,814 / out 48), $0.0098)
Highlights
Keeps 5-hour provider windows continuous instead of letting idle gaps drift
your schedule.
Reads usage from zero-quota endpoints and triggers windows through the
official provider tools whenever possible.
Supports Claude Code, Codex, and opt-in GLM/Z.ai Coding Plan monitoring.
Includes dry-run modes, weekly-limit guards, reset buffers, local config, and
no telemetry.
curl -fsSL https://raw.githubusercontent.com/wavever/CCLimitPing/main/install.sh | sh
limitping config init
limitping ping --dry-run
limitping status
limitping watch
Use limitping ping --dry-run or limitping watch --dry-run first if you want
to inspect what would happen without consuming provider quota.
Provider
Read usage (zero-quota)
Trigger
Auth
Claude Code
…/api/oauth/usage
interactive Claude Code CLI
OAuth (Keychain / ~/.claude )
Codex
…/backend-api/wham/usage
codex exec
OAuth ( ~/.codex/auth.json )
GLM (Zhipu / Z.ai)
…/api/monitor/usage/quota/limit
minimal chat completion
API key (config / env)
Note
GLM is off by default and not yet verified on a live plan — see
GLM before enabling it.
When watch sees a 5h window has reset, it first checks for an already-running
Claude/Codex CLI task. If one is running, limitping waits and re-reads usage
instead of sending its own ping, because that task's next model request will
start the new window naturally.
Claude : reads GET https://api.anthropic.com/api/oauth/usage using the
OAuth token from the macOS Keychain ( Claude Code-credentials ) or
~/.claude/.credentials.json . Triggering uses a TTY-backed interactive
claude "<prompt>" session, so it continues to start the Claude
subscription-backed window after the headless print command moves to Agent
SDK/API credits.
Codex : reads GET https://chatgpt.com/backend-api/wham/usage using the
OAuth token from ~/.codex/auth.json .
GLM : reads GET …/api/monitor/usage/quota/limit (on api.z.ai or
open.bigmodel.cn ) using your Coding Plan API key. GLM has no standalone CLI,
so the trigger is a direct minimal chat completion to
…/api/coding/paas/v4/chat/completions rather than a shell-out.
Claude/Codex tokens are reused from the official tools (no separate login) and
refreshed on 401. GLM uses a static API key (from config or env) — see below.
limitping ships as a single self-contained binary — no Go required .
One-line script (macOS / Linux):
curl -fsSL https://raw.githubusercontent.com/wavever/CCLimitPing/main/install.sh | sh
Downloads the right prebuilt binary from the
latest release into
/usr/local/bin (or ~/.local/bin ). Override with LIMITPING_INSTALL_DIR .
Upgrade — replace the installed binary with the latest release:
limitping upgrade
Aliases: limitping up , limitping update .
Uninstall — remove the installed binary plus config/cache:
limitping uninstall
Aliases: limitping rm , limitping remove .
Use limitping uninstall --keep-config to preserve ~/.config/limitping (or
$XDG_CONFIG_HOME/limitping ).
Manual download — grab the archive for your platform from the
Releases page ( .tar.gz for
macOS/Linux, .zip for Windows):
tar -xzf limitping_darwin_arm64.tar.gz
sudo mv limitping /usr/local/bin/
Homebrew (macOS / Linux) — brew install wavever/tap/limitping
(works once the Homebrew tap is set up — see .goreleaser.yaml ).
From source (developers, needs Go 1.25+):
go install github.com/wavever/CCLimitPing/cmd/limitping@latest
# or, from a clone:
go build -o bin/limitping ./cmd/limitping
Each provider you enable needs its own credentials: the claude / codex CLIs
logged in (Claude / Codex), or a Coding Plan API key (GLM).
limitping config init # write ~/.config/limitping/config.toml
limitping status # show 5h/weekly % + reset countdowns (alias: s)
limitping status -v # also print raw JSON
limitping ping # trigger all enabled providers now (alias: p)
limitping ping claude # Claude only
limitping ping codex # Codex only
limitping ping glm # GLM only
limitping ping --dry-run # show the commands without sending
limitping watch # foreground daemon: ping each window at reset (alias: w)
limitping watch claude # watch only one provider (claude|codex|glm)
limitping watch --dry-run # log when pings would fire, without sending
limitping version # print the version (aliases: v, ver)
limitping upgrade # update to the latest GitHub release (aliases: up, update)
limitping uninstall # remove limitping plus config/cache (aliases: rm, remove)
Short aliases are also available for config commands: limitping c i for
config init and limitping c p for config path .
limitping --help lists aliases inline, for example ping, p .
ping shows the exact command, a live timer (a spinner on a terminal), the
token usage the ping consumed where available (parsed from codex --json or
the GLM API response), and a USD cost where available:
claude → claude --model haiku .
claude ✓ pinged (6.6s)
codex → codex exec --skip-git-repo-check --json -c model_reasoning_effort=low -m gpt-5.4-mini ok
codex ✓ pinged (13.6s, 16,862 tok (in 16,814 / out 48), $0.0098)
Cost sources:
Claude interactive mode does not expose per-invocation machine-readable
usage/cost, so no token/cost suffix is shown.
Codex (subscription) doesn't return a USD cost, so — like CodexBar/ccusage
— we derive the equivalent API-rate cost from the
LiteLLM pricing dataset
( cost = non-cached-input × input + cached-input × cache-read + output × output ).
The dataset is cached at ~/.config/limitping/litellm_prices.json (24h TTL),
with model-alias/date-suffix fallbacks. Requires [codex].model to be set so
the rate can be looked up.
GLM is a per-prompt subscription, so no per-call USD cost is shown — only
the token count.
Claude triggering still consumes a small amount of Claude subscription quota,
but the interactive CLI does not expose the exact per-ping token count.
claude
5h [█████░░░░░] 51.0% resets in 3h14m (Sun 00:10)
weekly [█████░░░░░] 54.0% resets in 7h04m (Sun 04:00)
codex (plus)
5h [██░░░░░░░░] 24.0% resets in 3h15m (Sun 00:11)
weekly [████░░░░░░] 37.0% resets in 111h57m (Thu 12:53)
Configuration
~/.config/limitping/config.toml (honors $XDG_CONFIG_HOME ):
weekly_threshold = 0.99 # skip pinging when weekly usage >= this (0..1), until weekly reset
reset_buffer = " 10s " # wait this long after a reset before pinging (ensures rollover)
notify = true # macOS notifications on ping/skip/failure
[ claude ]
enabled = true
prompt = " . "
model = " haiku " # cheapest tier; triggering doesn't need a SOTA model
extra_args = [] # extra Claude CLI args; print/headless-only flags are ignored
align_start = " " # optional RFC3339 anchor for the first window; empty = start ASAP
[ codex ]
enabled = true
prompt = " ok "
model = " gpt-5.4-mini " # cheapest Codex model for triggering
reasoning_effort = " low " # "minimal" is rejected when web_search/image_gen tools are enabled
extra_args = []
align_start = " "
[ glm ]
enabled = false # opt-in: enable once you have a plan + API key
prompt = " ok "
model = " glm-4.6 " # cheapest standard model; flagship GLM-5/5.1 cost a multiplier
platform = " global " # "global" = api.z.ai, "cn" = open.bigmodel.cn (Zhipu)
api_key = " " # empty = read from $ZAI_API_KEY (global) / $ZHIPU_API_KEY (cn)
align_start = " "
Top-level keys:
weekly_threshold — when the weekly window is at/above this, watch stops
pinging and waits for the weekly reset (unless usable credits exist).
reset_buffer — how long to wait after a window's reset time before
pinging, so the window has definitely rolled over.
align_start (per provider) — pin the phase of your windows: set to a
future RFC3339 time to delay the very first ping until then; afterwards windows
chain automatically every ~5h.
Triggering a window doesn't depend on the model — any billable request starts
the 5h clock — so the ping uses each provider's cheapest model to eat the least of
your budget:
Claude → haiku : also avoids the separate weekly Opus bucket.
Codex → gpt-5.4-mini : the mini variant (see ~/.codex/models_cache.json
for what your plan offers).
GLM → glm-4.6 : a standard model; the flagship GLM-5/5.1 deduct quota at a
2–3× multiplier, so avoid them for a mere trigger.
Claude/Codex don't expose per-model prices at runtime (Anthropic's local cost
cache is empty; Codex's model cache has no price field), so the cheapest model is
a sensible default rather than a live price lookup. Override model per provider
if you prefer.
GLM (Zhipu / Z.ai Coding Plan)
GLM uses the same 5h + weekly structure as Claude/Codex, but two things
differ:
Auth is a static API key , not OAuth. Put it in [glm].api_key , or leave it
empty and export ZAI_API_KEY (global) / ZHIPU_API_KEY (CN). Usage reads hit
…/api/monitor/usage/quota/limit ; the key goes in the Authorization header
without a Bearer prefix (that's how the endpoint expects it).
The trigger is a direct API call , because GLM has no standalone CLI. It
sends a one-token chat completion to …/api/coding/paas/v4/chat/completions .
Unverified on a live plan. GLM is off by d

[truncated]
