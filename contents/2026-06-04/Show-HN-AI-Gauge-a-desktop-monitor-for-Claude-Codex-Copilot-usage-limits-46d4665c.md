---
source: "https://github.com/jpajak/ai-gauge"
hn_url: "https://news.ycombinator.com/item?id=48392849"
title: "Show HN: AI Gauge, a desktop monitor for Claude/Codex/Copilot usage limits"
article_title: "GitHub - jpajak/ai-gauge: Compact cross-platform monitor for Claude.ai, ChatGPT Codex, and GitHub Copilot usage limits. · GitHub"
author: "jpajak"
captured_at: "2026-06-04T04:35:40Z"
capture_tool: "hn-digest"
hn_id: 48392849
score: 2
comments: 0
posted_at: "2026-06-04T02:18:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Gauge, a desktop monitor for Claude/Codex/Copilot usage limits

- HN: [48392849](https://news.ycombinator.com/item?id=48392849)
- Source: [github.com](https://github.com/jpajak/ai-gauge)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T02:18:46Z

## Translation

タイトル: Show HN: AI Gauge、クロード/コーデックス/コパイロットの使用制限用デスクトップ モニター
記事のタイトル: GitHub - jpajak/ai-gauge: Claude.ai、ChatGPT Codex、および GitHub Copilot の使用制限用のコンパクトなクロスプラットフォーム モニター。 · GitHub
説明: Claude.ai、ChatGPT Codex、および GitHub Copilot の使用制限を監視するコンパクトなクロスプラットフォーム モニター。 - jpajak/ai-gauge
HN テキスト: こんにちは HN、新しいアカウントですが長年の読者です。私がこれを自分用に構築したのは、Claude、Codex、Copilot 全体の使用状況を手動でチェックし続け、セッションと毎週の使用状況をすべて 1 か所で追跡したかったからです。これにより、AI サブスクリプションをより適切に追跡し、最大限に活用できるようになりました。これは、エージェント コーディングによって、以前は時間をかける価値もなかったような便利なツールの作成や、必要最小限のソリューションではなく、より完全なユーティリティの構築が容易になったことの良い証拠でもあります。他の人が役に立つと思った場合に備えて公開します。フィードバックをお待ちしております。

記事本文:
GitHub - jpajak/ai-gauge: Claude.ai、ChatGPT Codex、および GitHub Copilot の使用制限用のコンパクトなクロスプラットフォーム モニター。 · GitHub
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
jpajak
/
aiゲージ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

イオンオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
49 コミット 49 コミット .github .github docs/ スクリーンショット docs/ スクリーンショット src/ aigauge src/ aigauge テスト テスト ツール ツール .gitignore .gitignore AI Gauge-datasheet.md AI Gauge-datasheet.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンスライセンス README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md build.ps1 build.ps1 build.sh build.sh pyinstaller_entry.py pyinstaller_entry.py pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
複数の AI サブスクリプションの料金を支払い、使用状況を頻繁に確認する場合は、AI ゲージが役立つ可能性があります。セッションと週ごとの使用状況、リセット時間、アカウント残高、支出が常に表示されるコンパクトなビューで表示されるため、支払っている金額を最大限に活用できます。
Claude.ai 、 ChatGPT Codex 、 GitHub Copilot 、 OpenRouter の使用のためのコンパクトなモニター。各 OS のプラットフォーム ネイティブ UI を使用した手動 + 自動更新:
Windows / Linux — 常に一番上に表示されるドラッグ可能なフレームレス ウィジェットとシステム トレイ アイコン。
macOS — 統計スタイルのメニューバー項目 ( ● 42% ● 78% ● 15% )。パネルをクリックすると、ポップオーバーとしてパネルが開きます。
Python 3.11以降が必要です。シークレットは、OS ネイティブの資格情報ストア (Windows Credential Manager / DPAPI、macOS キーチェーン、Linux Secret Service) に存在します。自動起動では、プラットフォームの標準メカニズム (Run キー / LaunchAgent / ~/.config/autostart ) を使用します。
現在のバージョン: 0.5.9 。リリースノートについては、CHANGELOG.md を参照してください。
AI Gauge は独立したオープンソース プロジェクトであり、非公式のローカル デスクトップです。
ユーティリティ。 Anthropic、OpenAI、GitHub、Microsoft、
OpenRouter、またはその他のプロバイダー。プロバイダーのページと API は変更される可能性があります。
気づいてください。
Windows / Linux — 常に最前面に表示されるフローティング ウィジェット、フル パネルおよびコレクション

apsed ピル モード:
macOS — プロバイダーごとに色付きのドットが表示される統計スタイルのメニューバー項目。クリックしてパネルをポップオーバーとして開きます。
各リリースの事前構築済みバイナリは、「リリース」ページで公開されています。 OS のアーカイブを選択し、解凍して実行します。
SHA256 合計は各アーカイブとともに公開されます。ビルドは署名されていません。SmartScreen / Gatekeeper の処理については、以下の初回起動時の警告セクションを参照してください。
py -m venv .venv
.\.venv\Scripts\ python.exe - m pip install - e .[ dev ]
.\.venv\Scripts\ python.exe - m aigauge
macOS / Linux (bash):
python3 -m venv .venv
./.venv/bin/python -m pip install -e ' .[dev] '
./.venv/bin/python -m aigauge
初めて起動すると、ウィジェットが有効なプロバイダー タイルとともに表示されます。クロードとコーデックスはサインイン フローを使用します。 GitHub Copilot と OpenRouter は、API 認証情報を使用して [設定] から構成されます。 [設定] を開いて、使用しないプロバイダーを無効にするか、クロード/コーデックス アカウントを追加します。
プロバイダー
セットアップ
クロードアイ
サインイン (推奨): 埋め込みブラウザが開きます。 [Google で続行] をクリックしないでください。Google は埋め込みブラウザ内での認証を拒否します。アカウントが Google にリンクされている場合は、同じメールアドレスを [メールアドレスを入力してください] ボックスに入力し、受信トレイに送信されたマジック リンクを使用します。 Cookie を貼り付けます。マジックリンクが使用できない場合はフォールバックします。以下を参照してください。 「設定」→「クロード」からクロードのサブスクリプションを追加します。
ChatGPT コーデックス
クロードと同じ - 埋め込みブラウザで電子メールとマジック リンクを使用するか、フォールバックとして Cookie を貼り付けます。 OpenAI アカウントが Google またはパスキーを介してルーティングされる場合は、「Cookie の貼り付け」を使用します。埋め込みブラウザでは、これらのフローを完了できないことがよくあります。 [設定] → [コーデックス] から追加のコーデックス サブスクリプションを追加します。
GitHub コパイロット
https://github.com/settings/personal-access-tokens/new で詳細な PAT を作成します。個人プランの場合は、「アカウント権限」→「プラン」→「読み取り」を追加します。パス

設定に入ります。毎月の AI クレジット許容量を設定します (Pro=1,500、Pro+=7,000、Max=20,000)。 Copilot が組織を通じて請求される場合は、請求先組織に入り、組織の請求アクセス権と組織権限を持つトークン/アカウントを使用します → [管理] → [読み取り] 。
オープンルーター
https://openrouter.ai/keys で推論 API キーを作成し、設定に貼り付けます。アカウント残高とモデルのアクティビティを表示するには、 https://openrouter.ai/settings/provisioning-keys で管理キーも作成します。管理キーは推論には使用できません。 AI Gauge はそれを個別に保存し、OpenRouter 管理エンドポイントにのみ使用します。 1 日の支出予算はオプションです。
複数のクロード / コーデックス アカウント
Claude と Codex は、一度に複数のサブスクリプションを追跡できます。 [設定] → [Claude] または [設定] → [コーデックス] を開き、 [別の追加] をクリックしてアカウントに短い名前を付け、その特定の行に対して [サインイン] または [Cookie の貼り付け] を使用します。デフォルトのアカウントは Claude または Codex として表示されます。名前付きアカウントは、 Claude (Work) 、Codex (Account 2) などとして表示されます。
[全般] タブでは、プロバイダー グループを制御します。 Claude がチェックされている場合、設定されているすべての Claude アカウントが表示されます。 Codex がチェックされている場合、設定されているすべての Codex アカウントが表示されます。セカンダリ アカウントはプロバイダー タブから削除できます。各クロード/コーデックス アカウントは、個別の Cookie ストレージ、ブラウザー プロファイル データ、ウィジェット タイルの状態、および履歴レコードを使用します。
セッションは、OS ごとの app-data ディレクトリの下で次の実行まで維持されます。
AI ゲージにはテレメトリやバックエンド サービスは含まれません。プロバイダーのリクエスト
ローカル アプリから構成されたプロバイダーに対して行われます。参照
SECURITY.md はセキュリティとプライバシーに関するメモです。
埋め込みブラウザのサインインが機能しない場合 (例: アカウントに Google サインインやパスキー認証が必要な場合、またはマジックリンク パスが使用できない場合)、既存のセッション ファイルをコピーしてください。

通常のブラウザからアプリにアクセスしてください。クッキーは数週間以内に貼り直す必要があります。
通常どおり、Chrome / Edge / Firefox で https://claude.ai (または https://chatgpt.com ) にサインインします。
ChatGPT の場合は、 F12 → Network を押し、ページをリロードして、
chatgpt.com リクエストを作成し、完全なリクエスト ヘッダー → Cookie: 値をコピーします。
これには、分割セッション Cookie とコンパニオン認証 Cookie が含まれます。
__セキュア・オアイです。
Claude の場合は、 F12 → Network を押し、 https://claude.ai/new#settings/usage をリロードします。
claude.ai リクエストをクリックし、完全なリクエスト ヘッダー → Cookie をコピーします。
値。 sessionKey を含める必要があります。
アプリ内: [設定] → [Claude] または [設定] → [コーデックス] → [アカウントの横にある [Cookie] を貼り付け] をクリックし、貼り付け、保存します。
Windows / Linux: ウィジェットはデフォルトで他のウィンドウの上にフロートします。どこにでもドラッグして移動します。閉じる (✕) トレイに非表示になります。トレイ アイコンを右クリックして、[更新] / [設定] / [終了] を選択します。左クリックするとウィジェットの表示/非表示が切り替わります。最高のタイル読み取り値に基づいて、トレイ アイコンが黄色 ≥75% / 赤 ≥90% に変わります。
macOS: メニュー バー項目には、有効なプロバイダー/アカウント タイルの色付きのステータス ドットが表示されます。クリックするとパネルがポップオーバーとして開きます。外側をクリックして閉じる。右クリックすると、同じ更新 / 設定 / 終了メニューが表示されます。
システム トレイのない Linux (ストック GNOME): フローティング ウィジェットは表示されたままで、ウィジェットを右クリックすると、同じ表示 / 更新 / 設定 / 終了メニューが表示されます。
折りたたむ / 展開する: ウィジェット ヘッダーの - ボタンをクリックすると、コンパクトなピル ビューに縮小されます。有効なプロバイダー/アカウント チップは、必要に応じて追加の行にラップされ、スペースを節約するためにアカウント名のみを使用する名前付きセカンダリ クロード/コーデックス アカウントが使用されます。
未使用のプロバイダーを非表示にする: 設定で Claude / Codex / Copilot / OpenRouter のチェックを外して、それらのグループをウィジェットから削除します。1 つまたは 2 つだけを使用する場合に便利です。

彼ら。
自動更新は適応型です。手動更新または変更された使用法はアクティブになります。
ケイデンスが変化すると、変更されていない結果は、設定された最大間隔に向かってバックオフします。
デフォルトはアクティブ時間が 5 分、アイドル時間が最大 60 分です。
毎日のユーティリティとして実行する場合は、[設定] で [ログイン時に開始] を有効にします。
ほとんどのユーザーにとって、事前に構築されたダウンロードの方が簡単です。このセクションは、ローカルで構築するか、リリースをカットするメンテナー向けです。ビルド マシンには Python 3.11 以降と、 pip install -e .[dev] がすでに実行されている .venv が必要です。結果として得られるバイナリは、ターゲット マシン上で Python を必要としません。
v* に一致するタグ付きコミットは、リリース ワークフローを自動的に実行します。これにより、3 つのプラットフォームすべてが CI で構築され、メンテナーが公開できるドラフト GitHub リリースとしてアップロードされます。
Chromium ランタイムが同梱されているため、バンドルは約 150 ～ 200 MB です。ユーザー データは依然としてバンドルの外、OS ごとの app-data ディレクトリの下に存在します。
単一ファイルのバイナリ (初回起動が遅い) の場合は、-OneFile (PowerShell) または --onefile (bash) を渡します。 macOS では、単一ファイル形式よりも .app バンドルの方が推奨されます。
署名付き OS バンドル システムでの初回起動時の警告 - リリース アーティファクトは署名されていません。
Windows: SmartScreen → [詳細情報] → [とにかく実行]。
macOS: 最初の起動時にゲートキーパーがブロックします。 .app を右クリック→最初に開くか、xattr -dr com.apple.quarantine ai-gauge.app を 1 回実行します。
Linux: 署名層はありません。 ai-gauge が実行可能になっていない場合は、実行可能にするだけです。
メンテナのリリース手順については、RELEASING.md を参照してください。
.\.venv\Scripts\ python.exe - m pytest # Windows
。 / .venv / bin / python -m pytest # macOS / Linux
テストの対象となるのは、構成ラウンドトリップ、Copilot および OpenRouter REST ヘルパー (モック化された HTTP を使用)、ウィジェットの動作、およびスナップショット モデルです。プロバイダー スクレイパー (Claude/Codex) にはライブ ブラウザー セッションが必要で、手動で検証されます。
バグリポジトリ

rts、プロバイダー レイアウトの修正、PR は大歓迎です。参照
COTRIBUTING.md (環境設定、テスト コマンド、および
使用する課題テンプレート。
なぜ Chrome Cookie を読み取る代わりに埋め込みブラウザを使用するのでしょうか? Chrome 127 以降では、すべての外部 Python ライブラリによる Chrome/Edge Cookie の復号化をブロックするアプリバインド暗号化 (2024 年半ば) が追加されました。ブラウザ セッションを自分で所有することが唯一の信頼できる回避策です。
クロード/コーデックスのレイアウトは変更される場合があります。アップストリームの UI 更新後にプロバイダー タイルに「エラー」が表示された場合、src/aigauge/providers/{claude,codex}.py のページ抽出 JS を調整する必要があります。アプリの残りの部分は動作し続けます。
Copilot REST エンドポイントは、現在の暦月の請求使用量を返します。ウィジェットは、含まれている許容量に対して消費された総 AI クレジットを追跡します。正味数量/金額は、請求可能な超過料金のみです。リセットは翌月 1 日として計算されます。 GitHub は現在、信頼できる個人プランの許容量フィールドを公開していないため、設定ではカスタム フォールバックを備えたプラン ドロップダウンを使用します。年間/リクエストベースのアカウントは、従来のプレミアムリクエストフォールバックで処理されます。
コパイロットの使用は上流に遅れをとっています。 Copilot REST エンドポイントの更新は、Claude や Codex よりも著しく遅く、クレジット カウントが最近のアクティビティを反映するまでに数時間かかる場合があります。ウィジェットには最新の値が表示されます。

[切り捨てられた]

## Original Extract

Compact cross-platform monitor for Claude.ai, ChatGPT Codex, and GitHub Copilot usage limits. - jpajak/ai-gauge

Hi HN, new account but long-time reader. I built this for myself because I kept manually checking usage across Claude, Codex, and Copilot, and wanted to track the session and weekly usage all in one place. It's helping me better track and get the most out of my AI subscriptions. It's also a good testament to agentic coding making it easy to create useful tools that wouldn't have been worth the time before, or to build more complete utilities instead of bare-minimum solutions. Putting it out there in case anyone else finds it useful. Happy to take feedback.

GitHub - jpajak/ai-gauge: Compact cross-platform monitor for Claude.ai, ChatGPT Codex, and GitHub Copilot usage limits. · GitHub
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
jpajak
/
ai-gauge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
49 Commits 49 Commits .github .github docs/ screenshots docs/ screenshots src/ aigauge src/ aigauge tests tests tools tools .gitignore .gitignore AI Gauge-datasheet.md AI Gauge-datasheet.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md build.ps1 build.ps1 build.sh build.sh pyinstaller_entry.py pyinstaller_entry.py pyproject.toml pyproject.toml View all files Repository files navigation
If you pay for multiple AI subscriptions and frequently check your usage, AI Gauge might help. It shows session and weekly usage, reset times, account balances, and spend in a compact always-visible view, so you can get the most out of what you're paying for.
Compact monitor for Claude.ai , ChatGPT Codex , GitHub Copilot , and OpenRouter usage. Manual + auto refresh, with a platform-native UI on each OS:
Windows / Linux — always-on-top draggable frameless widget plus a system-tray icon.
macOS — Stats-style menu-bar item ( ● 42% ● 78% ● 15% ); the panel opens as a popover when you click it.
Requires Python 3.11+. Secrets live in the OS-native credential store (Windows Credential Manager / DPAPI, macOS Keychain, Linux Secret Service). Auto-start uses the platform's standard mechanism (Run key / LaunchAgent / ~/.config/autostart ).
Current version: 0.5.9 . See CHANGELOG.md for release notes.
AI Gauge is an independent open-source project and unofficial local desktop
utility. It is not affiliated with Anthropic, OpenAI, GitHub, Microsoft,
OpenRouter, or any other provider. Provider pages and APIs may change without
notice.
Windows / Linux — always-on-top floating widget, in full panel and collapsed pill modes:
macOS — Stats-style menu-bar item with per-provider tinted dots; click to open the panel as a popover:
Pre-built binaries for each release are published on the Releases page . Pick the archive for your OS, extract, and run:
SHA256 sums are published alongside each archive. Builds are unsigned — see the first-launch warnings section below for SmartScreen / Gatekeeper handling.
py - m venv .venv
.\.venv\Scripts\ python.exe - m pip install - e .[ dev ]
.\.venv\Scripts\ python.exe - m aigauge
macOS / Linux (bash):
python3 -m venv .venv
./.venv/bin/python -m pip install -e ' .[dev] '
./.venv/bin/python -m aigauge
On first launch the widget appears with enabled provider tiles. Claude and Codex use a Sign in flow; GitHub Copilot and OpenRouter are configured from Settings with API credentials. Open Settings to disable providers you don't use or to add more Claude/Codex accounts.
Provider
Setup
Claude.ai
Sign in (recommended): opens an embedded browser. Don't click "Continue with Google" — Google refuses to authenticate inside embedded browsers. If your account is Google-linked, just type that same email into the Enter your email box and use the magic link sent to your inbox. Paste cookie: fallback if magic-link is unavailable; see below. Add extra Claude subscriptions from Settings → Claude .
ChatGPT Codex
Same as Claude — use email + magic link in the embedded browser, or paste cookie as a fallback. If your OpenAI account routes through Google or a passkey, use Paste cookie ; embedded browsers often cannot complete those flows. Add extra Codex subscriptions from Settings → Codex .
GitHub Copilot
Create a fine-grained PAT at https://github.com/settings/personal-access-tokens/new . For personal plans, add Account permissions → Plan → Read . Paste into Settings; set your monthly AI credit allowance (Pro=1,500, Pro+=7,000, Max=20,000). If Copilot is billed through an organization, enter the billing org and use a token/account with org billing access and Organization permissions → Administration → Read .
OpenRouter
Create an inference API key at https://openrouter.ai/keys and paste it into Settings. To show account balance and model activity, also create a management key at https://openrouter.ai/settings/provisioning-keys . Management keys cannot be used for inference; AI Gauge stores it separately and only uses it for OpenRouter management endpoints. Daily spend budget is optional.
Multiple Claude / Codex accounts
Claude and Codex can track more than one subscription at a time. Open Settings → Claude or Settings → Codex , click Add another , give the account a short name, then use Sign in or Paste cookie for that specific row. The default account displays as Claude or Codex ; named accounts display as Claude (Work) , Codex (Account 2) , etc.
The General tab controls provider groups. If Claude is checked, all configured Claude accounts appear; if Codex is checked, all configured Codex accounts appear. Secondary accounts can be removed from their provider tab. Each Claude/Codex account uses separate cookie storage, browser profile data, widget tile state, and history records.
Sessions persist between runs under the per-OS app-data directory:
AI Gauge does not include telemetry or a backend service. Provider requests
are made from the local app to the configured providers. See
SECURITY.md for security and privacy notes.
If the embedded-browser sign-in doesn't work for you (e.g. your account requires Google sign-in, passkey authentication, or you can't use the magic-link path), copy your existing session cookie from your normal browser into the app. Cookies last weeks before they need re-pasting.
Sign into https://claude.ai (or https://chatgpt.com ) in Chrome / Edge / Firefox as you normally do.
For ChatGPT, press F12 → Network , reload the page, click a
chatgpt.com request, and copy the full Request Headers → Cookie: value.
This includes split session cookies plus companion auth cookies such as
__Secure-oai-is .
For Claude, press F12 → Network , reload https://claude.ai/new#settings/usage ,
click a claude.ai request, and copy the full Request Headers → Cookie:
value. It must include sessionKey .
In the app: Settings → Claude or Settings → Codex → click Paste cookie next to the account, paste, Save.
Windows / Linux: the widget floats above other windows by default. Drag anywhere to move; close (✕) hides to tray. Right-click the tray icon for Refresh / Settings / Quit. Left-click toggles widget visibility. Tray icon turns yellow ≥75% / red ≥90% based on the highest tile reading.
macOS: the menu-bar item shows tinted status dots for enabled provider/account tiles. Click it to open the panel as a popover; click outside to dismiss. Right-click for the same Refresh / Settings / Quit menu.
Linux without a system tray (stock GNOME): the floating widget stays visible and serves the same Show / Refresh / Settings / Quit menu via right-click on the widget.
Collapse / expand: click the − button in the widget header to shrink to the compact pill view. Enabled provider/account chips wrap onto additional rows when needed, with named secondary Claude/Codex accounts using just the account name to save space.
Hide unused providers: uncheck Claude / Codex / Copilot / OpenRouter in Settings to remove their group from the widget — useful if you only use one or two of them.
Auto-refresh is adaptive: manual refresh or changed usage enters the active
cadence, then unchanged results back off toward the configured max interval.
Defaults are 5 min active and 60 min idle max.
Enable Start at login in Settings if you want it to run as a daily utility.
For most users the pre-built downloads are easier — this section is for building locally or for maintainers cutting releases. The build machine needs Python 3.11+ and a .venv with pip install -e .[dev] already run; the resulting binary does not require Python on the target machine.
Tagged commits matching v* automatically run the release workflow , which builds all three platforms in CI and uploads them as a draft GitHub Release for the maintainer to publish.
Bundles are ~150-200 MB because the Chromium runtime ships inside. User data still lives outside the bundle, under the per-OS app-data directory.
For a single-file binary (slower first launch), pass -OneFile (PowerShell) or --onefile (bash). On macOS the .app bundle is recommended over the single-file form.
First-launch warnings on signed-OS-bundle systems — release artifacts are unsigned:
Windows: SmartScreen → "More info" → "Run anyway".
macOS: Gatekeeper blocks on first launch. Either right-click the .app → Open the first time, or run xattr -dr com.apple.quarantine ai-gauge.app once.
Linux: no signing layer; just make ai-gauge executable if it isn't already.
See RELEASING.md for maintainer release steps.
.\.venv\Scripts\ python.exe - m pytest # Windows
. / .venv / bin / python - m pytest # macOS / Linux
Tests cover: config round-trip, Copilot and OpenRouter REST helpers (with mocked HTTP), widget behavior, and snapshot models. Provider scrapers (Claude/Codex) require a live browser session and are validated manually.
Bug reports, provider-layout fixes, and PRs are welcome. See
CONTRIBUTING.md for environment setup, test commands, and
the issue templates to use.
Why an embedded browser instead of reading Chrome cookies? Chrome 127+ added App-Bound Encryption (mid-2024) that blocks every external Python library from decrypting Chrome/Edge cookies. Owning the browser session ourselves is the only reliable workaround.
Claude / Codex layouts may change. If a provider tile shows "error" after a UI update upstream, the page-extractor JS in src/aigauge/providers/{claude,codex}.py needs adjusting — the rest of the app keeps working.
The Copilot REST endpoint returns the current calendar month of billing usage. The widget tracks gross AI credits consumed against the included allowance; net quantity/amount is only the billable overage. Reset is computed as the 1st of the next month. GitHub does not currently expose a reliable personal-plan allowance field, so Settings uses a plan dropdown with a Custom fallback. Annual/request-based accounts are handled with a legacy premium-request fallback.
Copilot usage lags upstream. The Copilot REST endpoint updates noticeably slower than Claude or Codex — credit counts can take hours to reflect recent activity. The widget shows the most recent value GitHub

[truncated]
