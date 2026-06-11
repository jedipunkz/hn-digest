---
source: "https://github.com/outgate-ai/og-local"
hn_url: "https://news.ycombinator.com/item?id=48495934"
title: "Local Privacy Filter for Claude Code"
article_title: "GitHub - outgate-ai/og-local: A local privacy proxy for coding agents. Runs PII and secret detection in-process, redacts before your prompts leave your machine. · GitHub"
author: "alikh31"
captured_at: "2026-06-11T21:45:49Z"
capture_tool: "hn-digest"
hn_id: 48495934
score: 2
comments: 0
posted_at: "2026-06-11T20:24:48Z"
tags:
  - hacker-news
  - translated
---

# Local Privacy Filter for Claude Code

- HN: [48495934](https://news.ycombinator.com/item?id=48495934)
- Source: [github.com](https://github.com/outgate-ai/og-local)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T20:24:48Z

## Translation

タイトル: クロード コードのローカル プライバシー フィルター
記事のタイトル: GitHub - outgate-ai/og-local: コーディング エージェント用のローカル プライバシー プロキシ。 PII とシークレット検出をプロセス内で実行し、プロンプトがマシンから送信される前に編集します。 · GitHub
説明: コーディング エージェント用のローカル プライバシー プロキシ。 PII とシークレット検出をプロセス内で実行し、プロンプトがマシンから送信される前に編集します。 - outgate-ai/og-local

記事本文:
GitHub - outgate-ai/og-local: コーディング エージェント用のローカル プライバシー プロキシ。 PII とシークレット検出をプロセス内で実行し、プロンプトがマシンから送信される前に編集します。 · GitHub
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
アウトゲートアイ
/
OGローカル
公共
通知
ch にサインインする必要があります

アンジュ通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
24 コミット 24 コミット .githooks .githooks .github .github cmd/ ogl cmd/ ogl docs docs 内部 内部スクリプト スクリプト .codecov.yml .codecov.yml .commitlintrc.yaml .commitlintrc.yaml .coverageignore .coverageignore .editorconfig .editorconfig .gitignore .gitignore .golangci.yml .golangci.yml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sumすべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェント用のローカル プライバシー プロキシ。
When your coding agent reads a file, the file gets shipped to a third-party LLM.多くの場合それで問題ありません。 The file is open-source, or your team has a vendor agreement that covers it.そうでない場合もあります。差分に紛れ込んだ .env、テスト フィクスチャ内の顧客の電子メール、コメント内の API キー、プライベート サービスからのスタック トレース。
og-local は、マシン上で実行される単一のバイナリで、エージェントが行う API 呼び出しをインターセプトし、localhost を離れる前にプロンプ​​ト本文内の PII とシークレットを検出し、それらを不透明なプレースホルダーと交換し、編集されたプロンプトを上流に転送し、応答内のオリジナルを透過的に復元します。エージェントには違いが分かりません。上流のプロバイダーがシークレットを閲覧することはありません。
Detection runs in-process via the openai/privacy-filter ONNX model. There's no cloud round-trip and no network call to anywhere except the upstream provider you were already calling.モデルは ogl model pull を使用して一度フェッチされます。それ以降はすべてローカルです。
curl -fsSL https://raw.githubusercontent.com/outgate-ai/og-local/main/scripts/install.sh |しー
Windows (PowerShell):
irm https://raw.g

ithubusercontent.com / outgate - ai / og - local / main / scripts / install.ps1 |アイエックス
これにより、ogl バイナリがインストールされ、リダクションをサポートするプラットフォームでは、バンドルされた ONNX ランタイムが ogl が予期する場所に配置されます。次に、検出モデルを一度ダウンロードします。
ogl モデルは # ~840MB を ~/.cache/og-local にプルします。欠落している場合は ONNX ランタイムも取得します
以上です。ogl claude "..." と ogl codex "..." が編集されるようになりました。
最初の実行時に不足しているものがあった場合、ogl はエージェントを起動する前にその場でダウンロードすることを提案します (予想されるサイズが表示されます)。非対話型セッションでは、代わりに明示的なエラーが保持されます。
マニュアルのダウンロード。 Releases から署名付きアーカイブを取得します。
ogl_<バージョン>_<os>_<arch>.tar.gz (Windows では .zip)。編集可能なプラットフォームでは、アーカイブにはバイナリと lib/libonnxruntime.{so,dylib} (Windows では lib\onnxruntime.dll) が含まれています。その lib を ~/.cache/og-local/runtime/<os>-<arch>/ にコピーするか、 OGL_ONNXRUNTIME_LIB を指定します。
インストールに行きます。 go install github.com/outgate-ai/og-local/cmd/ogl@latest はパススルー ビルドのみを生成します。編集できません (cgo やバンドルされたモデル ランタイムはありません)。編集にはインストール スクリプトまたはリリース アーカイブを使用します。
プラットフォーム
墨消し
Linux / amd64
✅いっぱい
Linux / arm64
✅いっぱい
macOS / arm64 (アップルシリコン)
✅いっぱい
Windows / amd64
✅いっぱい
macOS / amd64 (インテル)
⚠️パススルーのみ
リダクションには 2 つのネイティブ ライブラリが必要です。 daulet/tokenizers (Windows staticlib は、アップストリームが公開していないため、リリース時にピン留めされたソースからビルドされます) と ONNX Runtime です。ONNX Runtime には Intel-macOS バイナリが同梱されていないため、パススルー ターゲットは 1 つです。パススルー プラットフォーム ogl claude / ogl codex では、プロンプトを保護せずに転送するのではなく、明確な「このビルドは編集できません」というメッセージを表示して終了します。
macOS の初回実行: バイナリおよびバンドルされたライブラリは公証されていません

ただし、ゲートキーパーがそれらを隔離する可能性があります。 xattr -d com.apple.quarantine $(command -v ogl) (および ~/.cache/og-local/runtime/ の下の lib) でクリアするか、右クリック→一度開くを選択します。
Windows の初回実行時: .exe は署名されていないため、SmartScreen が警告する場合があります。 「詳細」→「とにかく実行」を選択するか、PowerShell の Unblock-File を使用してファイルのブロックを解除します。
リリースには、checksums.txt とキーレス署名が同梱されます。
sha256sum -c checksums.txt # または: shasum -a 256 -c checksums.txt
検証 BLOB に署名 \
--certificate-oidc-issuer https://token.actions.githubusercontent.com \
--certificate-identity-regexp ' ^https://github.com/outgate-ai/og-local/.github/workflows/release.yml@refs/tags/v.*$ ' \
--署名チェックサム.txt.sig \
チェックサム.txt
クイックスタート
ogl モデル プル # 1 回、~800MB
# アントロピック風味のエージェント
ogl claude " cmd/server で失敗したテストを修正します "
# OpenAI 風味のエージェント
ogl codex " この PR を確認してください "
ogl は、ランダムなループバック ポート上でローカル プロキシを起動し、そのプロキシに子エージェントを指定し、エージェントを子プロセスとして実行します。完全な環境が子に転送され、エージェントはすでに持っている資格情報を使用し続けます。ogl はリクエストの送信先のみをリダイレクトします。エージェントが終了すると、ogl も終了します。
ほとんどのエージェントは *_BASE_URL 環境変数を使用してリダイレクトされます。 Codex はその変数を無視するため、ogl codex は代わりに、Codex をプロキシに指定する専用のプロバイダー構成を ~/.codex/ogl ( CODEX_HOME 経由) に書き込みます。独自の ~/.codex/config.toml はそのまま残されます。
ogl codex は、両方の Codex サインイン モードで動作します。 API キー ( OPENAI_API_KEY 、または ~/.codex/auth.json の auth_mode = "apikey") を使用して、 api.openai.com に転送します。 ChatGPT サブスクリプション ログインでは、ログインのトークンがスコープされるエンドポイントである chatgpt.com/backend-api/codex に転送され、それらのリクエストが api.openai.com に送信されます。

失敗するだろう。モードは ~/.codex/auth.json から読み取られ、 OPENAI_API_KEY が優先されます。どちらの方法でも、プロキシは既存の Codex 認証情報を転送し、その間のプロンプト本文を編集します。
デーモンも PID ファイルもグローバル状態もありません。
送信リクエストごとに、ogl はユーザー指定のコンテンツ フィールド (messages[].content 、 system 、ツール呼び出し入力、およびツール結果) を抽出し、ONNX ベースの PII 検出器を各フィールドに対してローカルで個別に実行し、検出されたスパンを不透明なプレースホルダー ( OG_PRIVATE_EMAIL_<hex> 、 OG_SECRET_<hex> など) に置き換え、書き換えられた本文を上流に転送し、反転します。プレースホルダーが SSE イベント間で分割される可能性があるストリーミング応答を含む、応答の置換。リクエスト フレーム フィールド (モデル、温度、ツール スキーマ、ID) は変更されずに渡されます。プレースホルダーと値のマッピングは、単一のリクエストの間のみ存続します。永続的なボールトはありません。プレースホルダー自体は、ogl セッションの存続期間中は決定的です。同じ値はすべてのリクエストで同じプレースホルダーにリダクトされるため、再送信された会話履歴はバイト同一のままであり、プロバイダー側​​のプロンプト キャッシュは機能し続けます。
OpenAI チャット コンプリーション ( /v1/chat/completions )、ストリーミングおよび非ストリーミング
OpenAI Responses ( /v1/responses ) および ChatGPT サブスクリプション Codex ( /backend-api/codex/responses )
Anthropic メッセージ ( /v1/messages )、ストリーミングおよび非ストリーミング (ツールの使用を含む)
他の道は手付かずに通過します
変数
目的
OGL_CACHE_DIR
モデル + ランタイム キャッシュ ディレクトリをオーバーライドします (デフォルト: ~/.cache/og-local )。
OGL_DEBUG
1 プロキシ アクティビティをファイルに記録します (PII 値はありません)。パスによってファイルが選択されます。起動時にパスが出力される
OGL_ONNXRUNTIME_LIB
ONNX ランタイム共有ライブラリへのパス。デフォルトのキャッシュ ルックアップをオーバーライドします。
貢献する

COTRIBUTING.md を参照してください。 TL;DR:
git clone https://github.com/outgate-ai/og-local
cd og-local
make setup # one-time: git フックをインストールします
make ci # lint + テスト + カバレッジ + ビルド
メインに対する PR には、合格した CI の実行と 1 回のレビューが必要です。従来のコミットの件名行には CI が適用されます。
ビジネス ソース ライセンス 1.1。2030 年 6 月 8 日に Apache 2.0 に自動的に変換されます。正確な条件については、「ライセンス」を参照してください。
平たく言えば、商用ソフトウェアを含め、自由に使用、変更、再配布できます。変更日までの 1 つの制限は、OG-local (または実質的に同様のサービス) をホスト型マルチテナント サービスとしてサード パーティに提供できないことです。変更日以降、その制限は解除され、単なる Apache 2.0 になります。
ライセンサー: Gatewise UG (haftungsbeschränkt)。商用代替品や質問については、support@outgate.ai までお問い合わせください。
コーディング エージェント用のローカル プライバシー プロキシ。 PII とシークレット検出をプロセス内で実行し、プロンプトがマシンから送信される前に編集します。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
5
v0.4.0
最新の
2026 年 6 月 11 日
+ 4 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A local privacy proxy for coding agents. Runs PII and secret detection in-process, redacts before your prompts leave your machine. - outgate-ai/og-local

GitHub - outgate-ai/og-local: A local privacy proxy for coding agents. Runs PII and secret detection in-process, redacts before your prompts leave your machine. · GitHub
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
outgate-ai
/
og-local
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
24 Commits 24 Commits .githooks .githooks .github .github cmd/ ogl cmd/ ogl docs docs internal internal scripts scripts .codecov.yml .codecov.yml .commitlintrc.yaml .commitlintrc.yaml .coverageignore .coverageignore .editorconfig .editorconfig .gitignore .gitignore .golangci.yml .golangci.yml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum View all files Repository files navigation
A local privacy proxy for coding agents.
When your coding agent reads a file, the file gets shipped to a third-party LLM. Often that's fine. The file is open-source, or your team has a vendor agreement that covers it. Sometimes it isn't: a .env slipped into a diff, a customer email in a test fixture, an API key in a comment, a stack trace from a private service.
og-local is a single binary that runs on your machine, intercepts the API calls your agent makes, detects PII and secrets in the prompt body before it leaves localhost, swaps them with opaque placeholders, forwards the redacted prompt upstream, and transparently restores the originals in the response. The agent never sees the difference. The upstream provider never sees the secrets.
Detection runs in-process via the openai/privacy-filter ONNX model. There's no cloud round-trip and no network call to anywhere except the upstream provider you were already calling. The model is fetched once with ogl model pull ; everything after that is local.
curl -fsSL https://raw.githubusercontent.com/outgate-ai/og-local/main/scripts/install.sh | sh
Windows (PowerShell):
irm https: // raw.githubusercontent.com / outgate - ai / og - local / main / scripts / install.ps1 | iex
This installs the ogl binary and, on platforms that support redaction, places the bundled ONNX Runtime where ogl expects it. Then download the detection model once:
ogl model pull # ~840MB into ~/.cache/og-local; also fetches the ONNX Runtime if missing
That's it — ogl claude "..." and ogl codex "..." now redact.
If anything is missing on first run, ogl offers to download it on the spot (showing the expected size) before launching the agent; in non-interactive sessions it keeps the explicit error instead.
Manual download. Grab a signed archive from Releases :
ogl_<version>_<os>_<arch>.tar.gz (or .zip on Windows). On a redaction-capable platform the archive contains the binary plus lib/libonnxruntime.{so,dylib} ( lib\onnxruntime.dll on Windows); copy that lib to ~/.cache/og-local/runtime/<os>-<arch>/ , or point OGL_ONNXRUNTIME_LIB at it.
go install . go install github.com/outgate-ai/og-local/cmd/ogl@latest produces a passthrough build only — it cannot redact (no cgo, no bundled model runtime). Use the install script or a release archive for redaction.
Platform
Redaction
linux / amd64
✅ full
linux / arm64
✅ full
macOS / arm64 (Apple Silicon)
✅ full
Windows / amd64
✅ full
macOS / amd64 (Intel)
⚠️ passthrough only
Redaction needs two native libraries: daulet/tokenizers (the Windows staticlib is built from the pinned source during release, since upstream doesn't publish one) and ONNX Runtime , which ships no Intel-macOS binary — hence the one passthrough target. On a passthrough platform ogl claude / ogl codex exit with a clear "this build cannot redact" message rather than forwarding your prompt unprotected.
macOS first run: the binary and bundled library aren't notarized yet, so Gatekeeper may quarantine them. Clear it with xattr -d com.apple.quarantine $(command -v ogl) (and the lib under ~/.cache/og-local/runtime/ ), or right-click → Open once.
Windows first run: the .exe is unsigned, so SmartScreen may warn. Choose "More info" → "Run anyway", or unblock the file with Unblock-File in PowerShell.
Releases ship checksums.txt and a keyless cosign signature:
sha256sum -c checksums.txt # or: shasum -a 256 -c checksums.txt
cosign verify-blob \
--certificate-oidc-issuer https://token.actions.githubusercontent.com \
--certificate-identity-regexp ' ^https://github.com/outgate-ai/og-local/.github/workflows/release.yml@refs/tags/v.*$ ' \
--signature checksums.txt.sig \
checksums.txt
Quickstart
ogl model pull # one-time, ~800MB
# Anthropic-flavored agent
ogl claude " fix the failing test in cmd/server "
# OpenAI-flavored agent
ogl codex " review this PR "
ogl starts a local proxy on a random loopback port, points the child agent at it, and exec s the agent as a child process. Your full environment forwards to the child, and the agent keeps using whatever credentials it already has — ogl only redirects where the requests go. When the agent exits, ogl exits.
Most agents are redirected with their *_BASE_URL env var. Codex ignores that variable, so ogl codex instead writes a dedicated provider config under ~/.codex/ogl (via CODEX_HOME ) pointing Codex at the proxy; your own ~/.codex/config.toml is left untouched.
ogl codex works with both Codex sign-in modes. With an API key ( OPENAI_API_KEY , or auth_mode = "apikey" in ~/.codex/auth.json ) it forwards to api.openai.com . With a ChatGPT subscription login it forwards to chatgpt.com/backend-api/codex , the endpoint that login's token is scoped for — sending those requests to api.openai.com would fail. The mode is read from ~/.codex/auth.json , with OPENAI_API_KEY taking precedence; either way the proxy forwards your existing Codex credentials and redacts the prompt body in between.
No daemon, no PID file, no global state.
For each outbound request, ogl extracts the user-supplied content fields ( messages[].content , system , tool-call inputs, and tool results), runs the ONNX-based PII detector locally over each field independently, replaces detected spans with opaque placeholders ( OG_PRIVATE_EMAIL_<hex> , OG_SECRET_<hex> , and the like), forwards the rewritten body upstream, and inverts the substitution on the response, including streaming responses where placeholders may split across SSE events. Request frame fields ( model , temperature , tool schemas, ids) are passed through unchanged. The placeholder↔value mapping lives only for the duration of a single request — there is no persistent vault. Placeholders themselves are deterministic for the lifetime of an ogl session: the same value redacts to the same placeholder on every request, so re-sent conversation history stays byte-identical and provider-side prompt caching keeps working.
OpenAI Chat Completions ( /v1/chat/completions ), streaming and non-streaming
OpenAI Responses ( /v1/responses ) and ChatGPT-subscription Codex ( /backend-api/codex/responses )
Anthropic Messages ( /v1/messages ), streaming and non-streaming, including tool use
Other paths pass through untouched
Variable
Purpose
OGL_CACHE_DIR
Override the model + runtime cache directory (default: ~/.cache/og-local )
OGL_DEBUG
1 logs proxy activity to a file (no PII values); a path chooses the file. The path is printed at startup
OGL_ONNXRUNTIME_LIB
Path to the ONNX Runtime shared library, overriding the default cache lookup
Contributing
See CONTRIBUTING.md . The TL;DR:
git clone https://github.com/outgate-ai/og-local
cd og-local
make setup # one-time: installs git hooks
make ci # lint + tests + coverage + build
PRs against main require a passing CI run and one review. Conventional-commits subject lines are CI-enforced.
Business Source License 1.1, converting automatically to Apache 2.0 on 2030-06-08 . See LICENSE for the precise terms.
In plain English: free to use, modify, and redistribute, including in commercial software. The one restriction until the change date is that you can't offer og-local (or a substantially-similar service) to third parties as a hosted multi-tenant service. After the change date, that restriction lifts and it's just Apache 2.0.
Licensor: Gatewise UG (haftungsbeschränkt). For commercial alternatives or questions: support@outgate.ai .
A local privacy proxy for coding agents. Runs PII and secret detection in-process, redacts before your prompts leave your machine.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
5
v0.4.0
Latest
Jun 11, 2026
+ 4 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
