---
source: "https://github.com/Reality-Shifting-Tech/deepsight"
hn_url: "https://news.ycombinator.com/item?id=49186867"
title: "DeepSight – give text-only LLMs eyes and hands (zero tokens, on-device)"
article_title: "GitHub - Reality-Shifting-Tech/deepsight: DeepSight: vision-session proxy for text-only LLMs — sketch + targeted tool calls (look/zoom/ocr) with content-addressed caching. OpenAI-compatible. · GitHub"
author: "eliaseffects"
captured_at: "2026-08-05T19:15:18Z"
capture_tool: "hn-digest"
hn_id: 49186867
score: 1
comments: 0
posted_at: "2026-08-05T18:25:54Z"
tags:
  - hacker-news
  - translated
---

# DeepSight – give text-only LLMs eyes and hands (zero tokens, on-device)

- HN: [49186867](https://news.ycombinator.com/item?id=49186867)
- Source: [github.com](https://github.com/Reality-Shifting-Tech/deepsight)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T18:25:54Z

## Translation

タイトル: DeepSight – テキストのみの LLM に目と手を与えます (トークンゼロ、デバイス上)
記事のタイトル: GitHub - Reality-Shifting-Tech/deepsight: DeepSight: テキスト専用 LLM 用のビジョン セッション プロキシ — コンテンツ アドレス指定のキャッシュを使用したスケッチ + 対象を絞ったツール呼び出し (look/zoom/ocr)。 OpenAI対応。 · GitHub
説明: DeepSight: テキスト専用 LLM 用のビジョン セッション プロキシ — コンテンツ アドレス指定のキャッシュを使用したスケッチ + ターゲット ツール呼び出し (look/zoom/ocr)。 OpenAI対応。 - 現実を変える技術/ディープサイト

記事本文:
GitHub - Reality-Shifting-Tech/deepsight: DeepSight: テキスト専用 LLM 用のビジョンセッション プロキシ — コンテンツ アドレス指定のキャッシュを使用したスケッチ + ターゲット ツール呼び出し (look/zoom/ocr)。 OpenAI対応。 · GitHub
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
現実を変えるテクノロジー
/
深視力
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
41 コミット 41 コミット

.github/ ワークフロー .github/ ワークフロー .hermes/ プラン .hermes/ プラン アセット アセット ベンチ ベンチ ドキュメント ドキュメント eval eval スクリプト スクリプト スキル/ ディープサイト スキル/ ディープサイト src/ ディープサイト src/ ディープサイト テスト テスト .editorconfig .editorconfig .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
DeepSeek (またはテキストのみのモデル) に目と手を与えてください。 DeepSight は、既存の LLM セットアップを現実世界に接続します。送信された画像を確認したり、デスクトップのスクリーンショットを撮ったり、画面上のテキストを読んだり、ボタンをクリックしたり、フィールドに入力したり、アプリを開いたり、Web を検索して事実を確認したりすることができます。すべてのビジョンはデバイス上で実行されます。macOS では Apple Vision、Windows では PIL + オプションの Tesseract OCR です。トークンも GPU もゼロで、画像データがマシンから流出することはありません。
pip インストールディープサイトビジョン
deepsight setup # ビジョンバイナリ (macOS) をコンパイルするか、環境 (Windows) を検証します
deepsight ドクター # すべてが機能することを確認
setup コマンドは、macOS 上の Swift からvision_eyes をコンパイルし (Xcode SDK が必要)、デフォルトで .env を作成し、次のステップを出力します。 Windows ではプラットフォームを検出し、PIL とオプションの Tesseract OCR を使用します。または、UV 同期を使用してソースからインストールします (下記のエージェント統合を参照)。
AI エージェントの場合: 1 つのコマンドでスキルをインストールします
DeepSight は、ロード可能なエージェント スキルとして出荷されます。つまり、エージェントがオンデマンドでロードできる手順知識 (セットアップ、配線、ツール プロトコル、トラブルシューティング) です。 1 つのコマンドでインストールします。ユーザーによる手順は必要ありません。
カール -fsSL https://raw.githubusercontent.com/Reality-Shifting-Tech/deepsight/main/scripts/install-skill.sh |バッシュ
それ

'座る。スクリプトはエージェントを検出し、適切なスキル ディレクトリ、Hermes ( ~/.hermes/skills/ )、Claude Code ( ~/.claude/skills/ )、Codex ( ~/.codex/skills/ )、または明示的な DEEPSIGHT_SKILLS_DIR にインストールします。アカウント、API キー、スキル自体の設定は必要ありません。以前のコピーは deepsight.old として保存され、削除されることはありません。
実行後、エージェントはスキルをロードし、その指示に従う必要があります (Hermes: skill_view(name='deepsight') ; 任意のエージェント: インストールされている SKILL.md を読み取ります)。ビジョン バイナリのコンパイル、推論ループの配線、16 ツール プロトコル、補助ビジョンの統合、およびトラブルシューティングについて説明します。
スキルの信頼できる情報源: このリポジトリの skill/deepsight/。チェックアウトからのオフライン インストール: cp -R skill/deepsight ~/.hermes/skills/ 。
Deepsight を最初からセットアップする 1 つのコマンド:
git clone https://github.com/Reality-Shifting-Tech/deepsight.git
CDディープサイト
UV同期
uv run deepsight setup # バイナリをコンパイルする (macOS) または環境を検証する (Windows)
uv run deepsight Doctor # すべてが機能することを確認します
それだけです。 setup コマンドは、ビジョン バイナリ (macOS) をコンパイルし、デフォルトで .env ファイルを作成し、次の手順を出力します。 Windows の場合、プラットフォームが自動的に検出され、代わりに PIL + オプションの Tesseract OCR が使用されます。
Python 統合 (エージェント用):
深視力から。バックエンドインポート NativeVisionBackend 、 ReasoningBackend 、 \
ComputerUseBackend 、 SearchBackend
深視力から。オーケストレーターのインポート オーケストレーター
深視力から。設定のインポート get_settings
設定 = get_settings ()
ビジョン = NativeVisionBackend ( bin_path = settings .vision_bin )
推論 = ReasoningBackend (
Base_url = 設定。 reasoning_base_url 、
api_key = 設定。 reasoning_api_key 、
モデル＝設定。推論モデル 、
)
# ビジョンのみのセッション
エージェント = オーケストレーター (ビジョン = visi

上、推論＝推論）
結果 = エージェント。 run ( "data:image/png;base64,..." , "この画像の説明" )
# デスクトップ オートメーション (macOS) または Windows オートメーションを使用する場合
深視力から。バックエンドインポート ComputerUseBackend
エージェント = オーケストレーター (
ビジョン＝ビジョン、リーズン＝推論、
コンピューター = ComputerUseBackend ()、
)
結果 = エージェント。 run ( "data:image/png;base64,..." ,
「ターミナルを開き、「npm run dev」を実行し、結果をキャプチャします。」 )
Windows では、 NativeVisionBackend の代わりに WindowsVisionBackend を使用します。
深視力から。バックエンド インポート WindowsVisionBackend
ビジョン = WindowsVisionBackend ()
特長
DeepSight は、4 つのレイヤーに編成された 16 のツールを推論モデルに公開します。
👁️ ビジョンツール — 画面を参照
ツール
何をするのか
見てください
画像の長方形領域を記述します
ocr
領域内のすべてのテキストを、書かれたとおりに正確に転写します。
ズーム
領域を拡大して細部を検査します
数える
領域内の説明に一致するオブジェクトを数える
見つける
説明によってオブジェクトを検索し、その境界ボックス (x%、y%、w%、h%) を返します。
すべてのビジョン ツールはゼロトークンです。macOS では Apple Vision (コンパイルされた Swift バイナリ経由)、または Windows では PIL + オプションの Tesseract OCR を使用します。ネットワーク、GPU、API 呼び出しはありません。
📸 ライブキャプチャツール — 現在何が起こっているかを確認します
ツール
何をするのか
キャプチャ
画面 (または特定のウィンドウ) のスクリーンショットを撮り、完全なビジョン パイプラインで分析します。
時計
画面を経時的に監視します。一定の間隔でキャプチャし、知覚的ハッシュを使用して同一のフレームをスキップし、変更のタイムラインを返します。ターゲットテキストが表示されるときにパラメータが停止するまではオプション
キャプチャ後、後続のすべてのビジョン ツールはキャプチャされた画面上で動作します。モデルはキャプチャ、検査、動作を行った後、再度キャプチャを行うことができます。
🌐 グラウンディングツール — 事実を確認する
ツール
何をするのか
地面
Web でクレームやエンティティ、フェッチなどを検索する

h 一番上の結果を返し、引用を含む検証の概要を返します。
Brave Search によって提供されます。 DEEPSIGHT_SEARCH_API_KEY によってゲートされる — 設定を解除すると正常に機能が低下します。
🖱️ アクション ツール — デスクトップと対話します
ツール
何をするのか
クリック
位置をクリックします (x%、y% — 検索出力と一致します)
タイプ
フォーカスされた入力フィールドにテキストを入力します
キー
キーボード ショートカット ( cmd+s 、 return 、 Escape 、 ctrl+c ) を押します。
スクロール
アクティブなウィンドウをスクロールします (方向、クリック数)
開く
アプリケーションを名前で起動またはアクティブ化する
フォーカス
タイトルを一致させてウィンドウを最前面に移動します
アプリ
表示されているすべてのアプリケーションとそのウィンドウ タイトルをリストします。
窓
% 画面座標を使用してウィンドウのサイズ変更または位置変更を行う
アクション ツールは、macOS osascript (組み込み) または cliclick (推奨: brew install cliclick ) を使用します。システム設定のアクセシビリティ権限が必要です。
深視力から。オーケストレーターのインポート オーケストレーター
深視力から。バックエンドインポート NativeVisionBackend 、 ReasoningBackend 、 ComputerUseBackend
# エージェントをセットアップする
ビジョン = NativeVisionBackend ( bin_path = "vision_eyes" )
推論 = ReasoningBackend (
Base_url = "https://api.deepseek.com/v1" ,
api_key = "sk-..." ,
モデル = "ディープシーク-v4-フラッシュ" ,
)
エージェント = オーケストレーター (ビジョン = ビジョン、推論 = 推論、コンピューター = ComputerUseBackend ())
# モデルは 16 個のツールすべてを自律的に使用します。
代理店。走って（
image_url = "データ:画像/png;base64,..." ,
user_text = "ターミナルを開き、新しいゲームプロジェクトを作成してビルドします。"
「その後、結果をキャプチャして、コンパイルされたかどうかを教えてください。」 、
応答形式 = { "タイプ" : "json_object" },
)
モデルは、ターミナルを開き、コマンドを入力し、画面をキャプチャして出力を確認し、エラーを見つけて修正し、再構築し、結果を報告します。
git clone https://github.com/Reality-Shifting-Tech/deepsight.git
CDディープサイト
UV同期
ビルドアイを作る
輸出深いため息

T_VISION_BIN= " $PWD /scripts/vision_eyes "
uv run deepsight 記述パス/to/image.jpg
UV ラン ディープサイト ドクター
窓
git clone https://github.com / 現実 - 変化 - 技術 / deepsight.git
CDディープサイト
UV同期
winget インストール UB - Mannheim.TesseractOCR
uv run python - m deepsight パス / to / image.jpg を記述します
UV ラン ディープサイト ドクター
Windows は、PIL ベースのシーン分析 (色、明るさ、テクスチャ) とオプションの Tesseract OCR に WindowsVisionBackend を使用します。デスクトップ オートメーションはネイティブ Windows API ( user32.dll 、PowerShell SendKeys) を使用します。追加のツールをインストールする必要はありません。
uv run deepsight 記述パス/to/image.jpg
出力: OCR テキスト、シーン分類、顔/人間/動物の数、検出されたスポーツ、カラー パレット、検出されたすべてのオブジェクトの境界ボックス。
ビジョンセッション（推論ループ）
深視力から。バックエンドインポート NativeVisionBackend 、 ReasoningBackend
深視力から。オーケストレーターのインポート オーケストレーター
ビジョン = NativeVisionBackend ( bin_path = "vision_eyes" )
推論 = ReasoningBackend (
Base_url = "https://api.deepseek.com/v1" ,
api_key = "sk-..." ,
モデル = "ディープシーク-v4-フラッシュ" ,
)
セッション = オーケストレーター (ビジョン = ビジョン、推論 = 推論)
結果 = セッション 。走って（
image_url = "https://example.com/screenshot.png" ,
user_text = "画面には何が表示されていますか? テキストを見つけてレイアウトを説明してください。" 、
)
print (結果と内容)
建築
推論モデルは、ユーザーのリクエストと 16 個のツールすべてのツール定義を受け取ります。
ビジョン ツール (look、ocr、zoom、count、locate) は、Perception モジュールを経由してルーティングされ、ゼロトークン分析のために、vision_eyes (コンパイルされた Apple Vision バイナリ) をシェルします。
ライブ キャプチャ ( Capture 、 watch ) は、macOS スクリーンキャプチャを使用して画面を取得し、結果をアクティブな画像として保存し、その上で完全なビジョン パイプラインを実行します。
アクションツール (クリック、

タイプ、キー、スクロール、オープン、フォーカス、アプリ、ウィンドウ） ComputerUseBackend を介してルーティングされます。これは、macOS の osascript またはデスクトップ オートメーションの cliclick を使用します。
Grounding ( ground ) は、SearchBackend を使用して Brave Search API 経由で Web を検索します。
構造化出力 — JSON スキーマに制約された回答を取得するには、response_format を渡します。
クロスキャプチャ メモリ — 知覚ハッシュ (ダッシュ) + OCR セット diff は、キャプチャ間での変更を追跡します。
知覚キャッシュは、セッション内で繰り返されるビジョン クエリを重複排除します。
画像の領域を検査します。すべての座標はパーセンテージ (0 ～ 100) です。そこにあるものの説明を返します。
領域内のテキストを転写します。改行を含む正確な文字起こし。
領域を拡大して詳細を検査します。
説明に一致するオブジェクトをカウントします。 「人」、「赤い車」、「ボタン」などの文字列を渡します。
説明によってオブジェクトを検索します。正規化された境界ボックスの座標に信頼度を加えた値を返します。 Apple Vision のオンデバイス検出 (顔、人間、動物、テキスト、四角形、顕著なオブジェクト) を使用します。例:locate("ログイン ボタン") は Login (85%): x=40% y=60% w=20% h=8% を返します。
スクリーンショットを撮ります。オプションの領域: 「screen」 (デフォルト) またはウィンドウタイトルの部分文字列 (例: 「ターミナル」 、 「Safari」 )。 OCR、検出されたオブジェクト、および変更を使用した完全なシーン分析を返します。

[切り捨てられた]

## Original Extract

DeepSight: vision-session proxy for text-only LLMs — sketch + targeted tool calls (look/zoom/ocr) with content-addressed caching. OpenAI-compatible. - Reality-Shifting-Tech/deepsight

GitHub - Reality-Shifting-Tech/deepsight: DeepSight: vision-session proxy for text-only LLMs — sketch + targeted tool calls (look/zoom/ocr) with content-addressed caching. OpenAI-compatible. · GitHub
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
Reality-Shifting-Tech
/
deepsight
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
41 Commits 41 Commits .github/ workflows .github/ workflows .hermes/ plans .hermes/ plans assets assets bench bench docs docs eval eval scripts scripts skill/ deepsight skill/ deepsight src/ deepsight src/ deepsight tests tests .editorconfig .editorconfig .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md pyproject.toml pyproject.toml View all files Repository files navigation
Give DeepSeek (or any text-only model) eyes and hands. DeepSight connects your existing LLM setup to the real world — it can look at images you send, take screenshots of your desktop, read text on screen, click buttons, type into fields, open apps, and search the web to verify facts. All vision runs on-device: Apple Vision on macOS, PIL + optional Tesseract OCR on Windows. Zero tokens, zero GPU, no image data ever leaves your machine.
pip install deepsight-vision
deepsight setup # compiles the vision binary (macOS) or verifies env (Windows)
deepsight doctor # confirms everything works
The setup command compiles vision_eyes from Swift on macOS (Xcode SDK required), creates a .env with defaults, and prints next steps. On Windows it detects the platform and uses PIL + optional Tesseract OCR. Or install from source with uv sync (see Agent Integration below).
For AI agents: install the skill in one command
DeepSight ships as a loadable agent skill : procedural knowledge (setup, wiring, tool protocols, troubleshooting) that any agent can load on demand. Install it with one command — no user steps:
curl -fsSL https://raw.githubusercontent.com/Reality-Shifting-Tech/deepsight/main/scripts/install-skill.sh | bash
That's it. The script detects the agent and installs into the right skills directory — Hermes ( ~/.hermes/skills/ ), Claude Code ( ~/.claude/skills/ ), Codex ( ~/.codex/skills/ ) — or an explicit DEEPSIGHT_SKILLS_DIR . No accounts, no API keys, no configuration required for the skill itself. A previous copy is preserved as deepsight.old , never deleted.
After it runs, the agent should load the skill and follow its instructions (Hermes: skill_view(name='deepsight') ; any agent: read the installed SKILL.md ). It covers compiling the vision binary, reasoning-loop wiring, the 16-tool protocol, auxiliary.vision integration, and troubleshooting.
Skill source of truth: skill/deepsight/ in this repo. Offline install from a checkout: cp -R skill/deepsight ~/.hermes/skills/ .
One command to set up deepsight from scratch:
git clone https://github.com/Reality-Shifting-Tech/deepsight.git
cd deepsight
uv sync
uv run deepsight setup # compiles binary (macOS) or verifies env (Windows)
uv run deepsight doctor # confirms everything works
That's it. The setup command compiles the vision binary (macOS), creates a .env file with defaults, and prints next steps. For Windows, it detects the platform automatically and uses PIL + optional Tesseract OCR instead.
Python integration (for agents):
from deepsight . backends import NativeVisionBackend , ReasoningBackend , \
ComputerUseBackend , SearchBackend
from deepsight . orchestrator import Orchestrator
from deepsight . config import get_settings
settings = get_settings ()
vision = NativeVisionBackend ( bin_path = settings . vision_bin )
reasoning = ReasoningBackend (
base_url = settings . reasoning_base_url ,
api_key = settings . reasoning_api_key ,
model = settings . reasoning_model ,
)
# Vision-only session
agent = Orchestrator ( vision = vision , reasoning = reasoning )
result = agent . run ( "data:image/png;base64,..." , "Describe this image" )
# With desktop automation (macOS) or Windows automation
from deepsight . backends import ComputerUseBackend
agent = Orchestrator (
vision = vision , reasoning = reasoning ,
computer = ComputerUseBackend (),
)
result = agent . run ( "data:image/png;base64,..." ,
"Open Terminal, run 'npm run dev', capture the result" )
On Windows, use WindowsVisionBackend instead of NativeVisionBackend :
from deepsight . backends import WindowsVisionBackend
vision = WindowsVisionBackend ()
Features
DeepSight exposes 16 tools to the reasoning model, organized into four layers.
👁️ Vision tools — see the screen
Tool
What it does
look
Describe a rectangular region of the image
ocr
Transcribe all text in a region, exactly as written
zoom
Zoom into a region for small-detail inspection
count
Count objects matching a description in a region
locate
Find an object by description and return its bounding box (x%, y%, w%, h%)
All vision tools are zero-token — they use Apple Vision on macOS (via a compiled Swift binary) or PIL + optional Tesseract OCR on Windows. No network, no GPU, no API calls.
📸 Live capture tools — see what's happening now
Tool
What it does
capture
Screenshot the screen (or a specific window) and analyze it with the full vision pipeline
watch
Monitor the screen over time — captures at an interval, uses perceptual hashing to skip identical frames, returns a timeline of changes. Optional until param stops when target text appears
After capture , all subsequent vision tools operate on the captured screen. The model can capture, inspect, act, then capture again.
🌐 Grounding tools — verify facts
Tool
What it does
ground
Search the web for a claim or entity, fetch the top result, and return a verification summary with citations
Powered by Brave Search. Gated by DEEPSIGHT_SEARCH_API_KEY — degrades gracefully when unset.
🖱️ Action tools — interact with the desktop
Tool
What it does
click
Click at a position (x%, y% — matches locate output)
type
Type text into the focused input field
key
Press keyboard shortcuts ( cmd+s , return , escape , ctrl+c )
scroll
Scroll the active window (direction, clicks)
open
Launch or activate an application by name
focus
Bring a window to front by matching its title
apps
List all visible applications and their window titles
window
Resize or reposition a window using % screen coordinates
Action tools use macOS osascript (built-in) or cliclick (recommended: brew install cliclick ). Requires Accessibility permission in System Settings.
from deepsight . orchestrator import Orchestrator
from deepsight . backends import NativeVisionBackend , ReasoningBackend , ComputerUseBackend
# Set up the agent
vision = NativeVisionBackend ( bin_path = "vision_eyes" )
reasoning = ReasoningBackend (
base_url = "https://api.deepseek.com/v1" ,
api_key = "sk-..." ,
model = "deepseek-v4-flash" ,
)
agent = Orchestrator ( vision = vision , reasoning = reasoning , computer = ComputerUseBackend ())
# The model uses all 16 tools autonomously:
agent . run (
image_url = "data:image/png;base64,..." ,
user_text = "Open a terminal, create a new game project, build it, "
"then capture the result and tell me if it compiled." ,
response_format = { "type" : "json_object" },
)
The model will: open Terminal, type commands, capture the screen to check output, locate errors, fix them, rebuild, and report the result.
git clone https://github.com/Reality-Shifting-Tech/deepsight.git
cd deepsight
uv sync
make build-eyes
export DEEPSIGHT_VISION_BIN= " $PWD /scripts/vision_eyes "
uv run deepsight describe path/to/image.jpg
uv run deepsight doctor
Windows
git clone https: // github.com / Reality - Shifting - Tech / deepsight.git
cd deepsight
uv sync
winget install UB - Mannheim.TesseractOCR
uv run python - m deepsight describe path / to / image.jpg
uv run deepsight doctor
Windows uses WindowsVisionBackend for PIL-based scene analysis (colors, brightness, texture) and optional Tesseract OCR. Desktop automation uses native Windows APIs ( user32.dll , PowerShell SendKeys) — no additional tools to install.
uv run deepsight describe path/to/image.jpg
Output: OCR text, scene classification, face/human/animal counts, detected sports, color palette, bounding boxes for every detected object.
Vision session (reasoning loop)
from deepsight . backends import NativeVisionBackend , ReasoningBackend
from deepsight . orchestrator import Orchestrator
vision = NativeVisionBackend ( bin_path = "vision_eyes" )
reasoning = ReasoningBackend (
base_url = "https://api.deepseek.com/v1" ,
api_key = "sk-..." ,
model = "deepseek-v4-flash" ,
)
session = Orchestrator ( vision = vision , reasoning = reasoning )
result = session . run (
image_url = "https://example.com/screenshot.png" ,
user_text = "What's on the screen? Find any text and describe the layout." ,
)
print ( result . content )
Architecture
Reasoning model receives the user's request plus tool definitions for all 16 tools.
Vision tools (look, ocr, zoom, count, locate) route through the Perception module, which shells vision_eyes — the compiled Apple Vision binary — for zero-token analysis.
Live capture ( capture , watch ) uses macOS screencapture to grab the screen, stores the result as the active image, and runs the full vision pipeline on it.
Action tools (click, type, key, scroll, open, focus, apps, window) route through ComputerUseBackend , which uses macOS osascript or cliclick for desktop automation.
Grounding ( ground ) uses SearchBackend to search the web via Brave Search API.
Structured output — pass response_format to get JSON-schema-constrained answers.
Cross-capture memory — perceptual hashing (dhash) + OCR set diff tracks what changed between captures.
Perception cache deduplicates repeated vision queries within a session.
Inspect a region of the image. All coordinates are percentages (0-100). Returns a description of what's there.
Transcribe text in a region. Exact transcription including line breaks.
Upscale and inspect a region for small details.
Count objects matching a description. Pass a what string like "people", "red cars", "buttons".
Find an object by description. Returns normalized bounding box coordinates plus confidence. Uses Apple Vision's on-device detection (faces, humans, animals, text, rectangles, salient objects). Example: locate("the login button") returns Login (85%): x=40% y=60% w=20% h=8% .
Take a screenshot. Optional region : "screen" (default) or a window title substring (e.g. "Terminal" , "Safari" ). Returns a full scene analysis with OCR, detected objects, and chang

[truncated]
