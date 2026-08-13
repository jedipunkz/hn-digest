---
source: "https://github.com/guillaumemeyer/watermarks-remover"
hn_url: "https://news.ycombinator.com/item?id=49282033"
title: "Watermarks-remover: Strip multi-vendor AI provenance marks"
article_title: "GitHub - guillaumemeyer/watermarks-remover: Strip multi-vendor AI provenance marks: Unicode text hygiene, statistical rewrite hooks, and C2PA/metadata from PNG/JPEG/SVG/PDF/DOCX/HTML/MD · GitHub"
author: "thunderbong"
captured_at: "2026-08-13T05:42:01Z"
capture_tool: "hn-digest"
hn_id: 49282033
score: 3
comments: 0
posted_at: "2026-08-13T05:21:36Z"
tags:
  - hacker-news
  - translated
---

# Watermarks-remover: Strip multi-vendor AI provenance marks

- HN: [49282033](https://news.ycombinator.com/item?id=49282033)
- Source: [github.com](https://github.com/guillaumemeyer/watermarks-remover)
- Score: 3
- Comments: 0
- Posted: 2026-08-13T05:21:36Z

## Translation

タイトル: ウォーターマーク除去ツール: マルチベンダー AI 来歴マークの除去
記事のタイトル: GitHub - guillaumemeyer/watermarks-remover: マルチベンダー AI 来歴マークの除去: Unicode テキストの衛生管理、統計的書き換えフック、PNG/JPEG/SVG/PDF/DOCX/HTML/MD からの C2PA/メタデータ · GitHub
説明: マルチベンダー AI 来歴マークの除去: PNG/JPEG/SVG/PDF/DOCX/HTML/MD からの Unicode テキスト衛生、統計的書き換えフック、および C2PA/メタデータ - guillaumemeyer/watermarks-remover

記事本文:
GitHub - guillaumemeyer/watermarks-remover: マルチベンダー AI 来歴マークの除去: Unicode テキストの衛生管理、統計的書き換えフック、PNG/JPEG/SVG/PDF/DOCX/HTML/MD からの C2PA/メタデータ · GitHub
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
ギョーメマイヤー
/
ウォーターマークリムーバー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
38 コミット 38 コミット .figlet .figlet .github .github skill/remove-ai-marks skill/remove-a

i-marks テスト テスト .dockerignore .dockerignore .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile.synthid Dockerfile.synthid LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md pytest.ini pytest.ini 要件-dev.txt 要件-dev.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
_ _ _ ____ ___ ____ ____ _ _ ____ ____ _ _ ____ ____ ____ _ _ ____ _ _ ____ ____
| | | |__| | |___ |__/ |\/| |__| |__/ |_/ [__ _ |__/ |___ |\/| | | | | |__ |__/
|_|_| | | | |__ | \ | | | | | \ | \_ ___] | \ |___ | | |__| \/ |___ | \
ウォーターマークリムーバー
エージェント スキル + stdlib Python スクリプトにより、テキストやファイルからマルチベンダー AI 来歴マークを削除します。これにより、所有するコンテンツのプライバシーと衛生を確保します。
ベンダー / エコシステム (クラスレベル): Claude 、 Gemini / SynthID-Text 、OpenAI 来歴サーフェス、open-LLM Kirchenbauer スタイルのマーク。
スキルパス：skills/remove-ai-marks/
(移行: 以前は Remove-claude-marks ; スラッシュ エイリアス /remove-claude-marks はまだ文書化されています)
# Grok ビルド / プロジェクトローカル
mkdir -p .grok/スキル
ln -sfn " $( pwd ) /skills/remove-ai-marks " .grok/skills/remove-ai-marks
# ユーザーグローバル Grok
mkdir -p ~ /.grok/skills
ln -sfn " $( pwd ) /skills/remove-ai-marks " ~ /.grok/skills/remove-ai-marks
/remove-ai-marks で呼び出すか、「AI ウォーターマーク / C2PA / クロード マーク / SynthID クラスのテキストを削除する」ように依頼します。
オプションのシステム ツール (存在する場合は自動的に使用されます):
コア スクリプトには Python 3.10 以降の stdlib のみが必要です。レイヤ B モデルの呼び出しはオプションです。
SCRIPTS=スキル/AIマークの削除/スクリプト
# 統合検査/クリーニング
python3 " $SCRIPTS /inspect_file.py "draft.md
python3 " $SCRIPTS /clean_file.py "draft.md -odraft.cleaned.md
python3 " $SCRIPTS /clean_file.py " photo.png -o photo.cleaned.png
python3 " $SCRIPTS /clean_file.py " Notes.docx -

o Notes.cleaned.docx
# テキストレイヤーA
python3 " $SCRIPTS /inspect_text.py "draft.md
python3 " $SCRIPTS /clean_text.py "draft.md -odraft.cleaned.md --stats
# レイヤー B 書き換えフック (デフォルト: プロンプトの印刷のみ - モデルは必要ありません)
python3 " $SCRIPTS /rewrite_text.py "draft.md --backend print-prompt --strength paraphrase
# オプションのローカル Ollama (デフォルトではループバックのみ - リモート エンドポイントには必要です)
# WATERMARKS_REWRITE_ALLOW_REMOTE=1 または --allow-remote):
# WATERMARKS_REWRITE_BACKEND=オラマ WATERMARKS_REWRITE_MODEL=ラマ3.2 \
# python3 "$SCRIPTS/rewrite_text.py" ドラフト.md -o ドラフト.rewrite.md
# API キーは WATERMARKS_REWRITE_API_KEY からのみ読み取られます (argv ではありません)。
# 画像
python3 " $SCRIPTS /inspect_image.py " ショット.png
python3 " $SCRIPTS /clean_image.py " ショット.png -o ショット.クリーン.png
オプションの SynthID ピクセル スコアリング
Inspection_image.py と clean_image.py はピクセル ドメインの SynthID を報告できます
外部チェックアウト時の信頼スコア
aloshdenny/reverse-SynthID
利用可能です。スコアラーはバンドルされていません。実行時にスコアラーからロードされます。
チェックアウトし、そのコードは上流プロジェクトの非営利の下に残ります。
研究ライセンス。
オプション 1: 1 つのコマンドによるブートストラップ (Docker なし)
SCRIPTS=スキル/AIマークの削除/スクリプト
# アップストリームのクローンを作成し、venv を作成し、スコアラーのみの依存関係をインストールします。
" $SCRIPTS /setup_synthid.sh "
# 画像をスコアリングします (デフォルトのチェックアウト: ~/reverse-SynthID)。
REVERSE_SYNTHID_DIR= ~ /reverse-SynthID \
~ /reverse-SynthID/.venv/bin/python " $SCRIPTS /score_synthid.py " ショット.png
# または、inspect / clean からスコアを表面化します (同じ venv Python)。
REVERSE_SYNTHID_DIR= ~ /reverse-SynthID \
~ /reverse-SynthID/.venv/bin/python " $SCRIPTS /inspect_image.py " ショット.png
setup_synthid.sh は、 --dir PATH 、 --ref REF 、および --full (
完全なアップストリームのrequirements.txt（トルクを追加）

h/ディフューザー用
このプロジェクトでは使用されないアップストリーム VAE バイパス)。
make docker-synthid-build
# 特権を持たず、読み取り専用の rootfs を使用して実行します。スコアラーは読むだけで済みます
# /data を実行し、stdout/tmp に書き込みます。
docker run --rm \
--user " $( id -u ) : $( id -g ) " \
--読み取り専用 --tmpfs /tmp \
-v " $( pwd ) :/data " \
ウォーターマークリムーバー-synthid-scorer /data/shot.png
イメージは、ビルド時に上流のソースからローカルにビルドされます。そうではありません
公開されるため、上流のコードは再配布されません。
V4 スコアリングでは、上流のチェックアウトからの artifacts/spectral_codebook_v4.npz を使用します。
(約 220 MB)。これは検出/スコアリングのみであり、ピクセルは削除されません。
透かし。
チャンネル
クロード
Gemini/シンセID
OpenAI
オープンLLM
Unicode / 編集ベースのテキスト
レイヤーA
レイヤーA
レイヤーA
レイヤーA
統計的サンプリングのテキスト
レイヤ B ベストエフォート
レイヤ B ベストエフォート
レイヤ B がある場合
レイヤ B ベストエフォート
C2PA / ファイルメタデータ
はい (リストされたフォーマット)
存在する場合ははい
存在する場合ははい
存在する場合ははい
ドット絵マーク
範囲外
オプションの SynthID スコア (外部)。範囲外の削除
範囲外
範囲外
バックドアのトレーニング
範囲外
範囲外
範囲外
範囲外
詳細: skill/remove-ai-marks/references/vendor-notes.md 、mark-classes.md 。
テキストマーキングの仕組み (短い)
最新の LLM 透かしは、非表示の文字だけでなく、トークンが選択される信号 (生成/サンプリング バイアス) を隠すことがよくあります。編集ベースのスキームでは、Unicode またはシノニム ルールが挿入されます。ファイル スキームは C2PA またはジェネレーターのメタデータを添付します。
レイヤ A は、編集ベースの Unicode キャリアを削除します (テスト可能)。
レイヤ B は、大量の書き換えによるウォーターマークのサンプリングを攻撃します (ベストエフォート型、言い換えや逆翻訳などの文献標準の攻撃)。
ファイル クリーナーは、サポートされているコンテナーから C2PA/XMP/props を削除します。
ベンダーが公開検出器とキーを出荷するまで、どのツールもそれを実行できません。

正直に「これは公式チェックに合格しません」と証明します。レポートでは、検証可能な作業とベストエフォート型の作業を区別する必要があります。
レイヤ B には非オリジナル モデルを使用してください (再スタンプを避けたい場合は、クロード テキストをクロードで書き直さないでください)。
免責事項: テキストの透かしの削除にかかる費用
テキストの透かしは文言自体の中に存在します。信号はトークンの選択肢全体に分散されるため、ほぼすべての文に少しずつ透かしが含まれます。 2 つの結果が続きます。これらが、レイヤー B が魔法の消しゴムではなくベストエフォートであると正直に説明される理由です。
削除とは、再構築ではなく、言葉の変更を意味します。段落をシャッフルしたり、見出しを変更したり、軽い修正を加えたりしても、信号はほとんど変わりません。統計マークを削除するには、セクションごとではなく、文ごとに、テキストのかなりの部分を書き直す必要があります。
言い換えるとコピーの品質が低下します。書き換えを行うと、元の単語の選択肢が書き換えモデルの選択肢に置き換えられ、口調、音声、精度が平坦化されます。プロダクションコピー (SEO、マーケティング、クライアントワーク) では、その劣化は現実のものであり、ライティングに最も関心を持っている人々によく見られます。これは、最上位モデルからテキストを取得し、それを能力の低いモデルに最初から書き直すように依頼するようなものです。結果は書き換えモデルの上限を超えることはできません。
これは、次のような正直な質問につながります。
とにかく安価なモデルでテキストを書き換える計画がある場合、そもそもなぜプレミアムモデルにお金を払うのでしょうか?安価なモデルを使用して直接生成する方が簡単かつ安価で、同じかそれ以上の最終結果が得られます。
レイヤー B は、マークフリー テキストへの安価な手段としてではなく、特にプレミアム モデルの思考と草案を必要としていて、衛生やプライバシーの要件を満たすために書き換えパスを受け入れる場合に意味があります。
衛生よりも品質が重要です: ロスレス パスを使用します — レイヤー A Unicode スクラブ プラス

ファイルのメタデータをクリーナーして、元の散文を保持します。
とにかく書き換える: 非オリジン モデルを使用し (オリジン モデルで書き直すとテキストが再スタンプされる可能性がある)、残留リスクが残ることを覚えておいてください。ベンダー検出器が失敗することを保証できるツールはありません。
フォーマット
検査する
クリーン
PNG / JPEG
C2PA チャンク / APP11、AI XMP ヒント
メタデータセグメントを削除する
SVG
<メタデータ> 、XMP
ストリップブロック
PDF
Byte/XMP + オプションツール
exiftool を推奨します。それがないと劣化する
DOCX
docProps / カスタムXml
小道具をスクラブし、customXml をドロップします
ODT
メタ.xml
ドロップジェネレーター / AIっぽいメタ
HTML
メタ、JSON-LD、データ AI*
タグ/属性を削除する
マークダウン
YAML フロントマター AI キー
ドロップキー + レイヤーAボディ
ピクセル ドメインのウォーターマークの削除と C2PA ソフト バインディング (メタデータが削除された後にリモートのコンテンツ資格情報マニフェストを再リンクできるコンテンツ内のウォーターマーク) は引き続き対象外です。ハードバインドされた C2PA を削除しても、これらのチャネルはクリアされません。オプションのローカル SynthID スコアラーは、検出のみに使用できます (上記を参照)。
このツールは、検証可能な削除 (Unicode 数、メタデータ アクション) とベストエフォートのレイヤー B 書き換えをレポートします。ベンダー検出器が失敗することを保証することはできません。
残留信号を自分でチェックするには (オプション、外部):
業界の 2 層コンテキスト (C2PA + 知覚できない透かし): Institute of AI PM ガイド 。
オプション
削除します
注意事項
Unicode スクラブ (レイヤー A)
ZWSP、Bidi、タグ、エキゾチックスペース、…
テキストの安全なデフォルト
リライト（レイヤーB）
統計的トークン マーク (ベストエフォート型)
常にスキルによって提供されます。コストスタイル — 免責事項を参照
コンテナ/メタデータ ストリップ
ファイルの出所
フォーマット表を参照
無差別ウェイトローカルモデル
元のモデルでの再スタンプを避ける
運用上の代替案
マトリックス: skill/remove-ai-marks/references/removal-matrix.md 。
skill/remove-ai-marks/references/ethics.md を参照してください。学術上の不正や不正ではなく、プライバシーとコンテンツの研究のため

「人間が書いた」主張。
python3 -m venv .venv && .venv/bin/pip pytest をインストールします
.venv/bin/python -m pytest # または: make test
スモークを作成する # クイック CLI でフィクスチャをスモークする
変更履歴
v0.3.2 — セキュリティ強化 (安全な書き込み、HTTP クライアント、CI サプライ チェーン)
安全でアトミックな出力書き込み: すべてのクリーナーは、一時ファイル + アトミックな名前変更 (safe_write_bytes /safe_write_text) を介して書き込み、シンボリックリンクされた宛先を拒否し、同じ安全なパスを介して .bak バックアップを作成します。事前に配置されたシンボリックリンク (/tmp またはダウンロード ディレクトリなど) は、クリーンな書き込みを任意のファイルにリダイレクトできなくなります。
rewrite_text.py HTTP クライアントの強化: リダイレクトは完全に拒否されるため、Authorization ヘッダー内の API キーを未検証のホストに再送信することはできません。非ループバック エンドポイントはデフォルトで拒否されます ( --allow-remote または WATERMARKS_REWRITE_ALLOW_REMOTE=1 でオプトインします)。 http スキームのみが受け入れられます。 --api-key が削除されました — キーは WATERMARKS_REWRITE_API_KEY を介して env のみです
リソースの上限: デフォルトの最大入力 1 GiB → 256 MiB、新しい 64 MiB stdin 上限、DOCX/ODT zip バジェット 512 MiB → 128 MiB、および exiftool/c2patool/SynthID サブプロセスに適用される RLIMIT_AS / RLIMIT_FSIZE (すべてのキャップは環境上書き可能)
サプライ チェーン: 権限付き SHA 固定された CI アクション: 内容: 読み取り、固定された dev d

[切り捨てられた]

## Original Extract

Strip multi-vendor AI provenance marks: Unicode text hygiene, statistical rewrite hooks, and C2PA/metadata from PNG/JPEG/SVG/PDF/DOCX/HTML/MD - guillaumemeyer/watermarks-remover

GitHub - guillaumemeyer/watermarks-remover: Strip multi-vendor AI provenance marks: Unicode text hygiene, statistical rewrite hooks, and C2PA/metadata from PNG/JPEG/SVG/PDF/DOCX/HTML/MD · GitHub
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
guillaumemeyer
/
watermarks-remover
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
38 Commits 38 Commits .figlet .figlet .github .github skills/ remove-ai-marks skills/ remove-ai-marks tests tests .dockerignore .dockerignore .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile.synthid Dockerfile.synthid LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md pytest.ini pytest.ini requirements-dev.txt requirements-dev.txt View all files Repository files navigation
_ _ _ ____ ___ ____ ____ _ _ ____ ____ _ _ ____ ____ ____ _ _ ____ _ _ ____ ____
| | | |__| | |___ |__/ |\/| |__| |__/ |_/ [__ __ |__/ |___ |\/| | | | | |___ |__/
|_|_| | | | |___ | \ | | | | | \ | \_ ___] | \ |___ | | |__| \/ |___ | \
watermarks-remover
Agent skill + stdlib Python scripts to strip multi-vendor AI provenance marks from text and files — for privacy and hygiene on content you own .
Vendors / ecosystems (class-level): Claude , Gemini / SynthID-Text , OpenAI provenance surfaces, open-LLM Kirchenbauer-style marks.
Skill path: skills/remove-ai-marks/
(migration: formerly remove-claude-marks ; slash alias /remove-claude-marks still documented)
# Grok Build / project-local
mkdir -p .grok/skills
ln -sfn " $( pwd ) /skills/remove-ai-marks " .grok/skills/remove-ai-marks
# User-global Grok
mkdir -p ~ /.grok/skills
ln -sfn " $( pwd ) /skills/remove-ai-marks " ~ /.grok/skills/remove-ai-marks
Invoke with /remove-ai-marks or ask to “strip AI watermarks / C2PA / Claude marks / SynthID-class text.”
Optional system tools (auto-used when present):
Core scripts need Python 3.10+ stdlib only. Layer B model calls are optional.
SCRIPTS=skills/remove-ai-marks/scripts
# Unified inspect / clean
python3 " $SCRIPTS /inspect_file.py " draft.md
python3 " $SCRIPTS /clean_file.py " draft.md -o draft.cleaned.md
python3 " $SCRIPTS /clean_file.py " photo.png -o photo.cleaned.png
python3 " $SCRIPTS /clean_file.py " notes.docx -o notes.cleaned.docx
# Text Layer A
python3 " $SCRIPTS /inspect_text.py " draft.md
python3 " $SCRIPTS /clean_text.py " draft.md -o draft.cleaned.md --stats
# Layer B rewrite hook (default: print prompt only — no model required)
python3 " $SCRIPTS /rewrite_text.py " draft.md --backend print-prompt --strength paraphrase
# Optional local Ollama (loopback only by default — remote endpoints require
# WATERMARKS_REWRITE_ALLOW_REMOTE=1 or --allow-remote):
# WATERMARKS_REWRITE_BACKEND=ollama WATERMARKS_REWRITE_MODEL=llama3.2 \
# python3 "$SCRIPTS/rewrite_text.py" draft.md -o draft.rewritten.md
# API keys are read from WATERMARKS_REWRITE_API_KEY only (never argv).
# Images
python3 " $SCRIPTS /inspect_image.py " shot.png
python3 " $SCRIPTS /clean_image.py " shot.png -o shot.cleaned.png
Optional SynthID pixel scoring
inspect_image.py and clean_image.py can report a pixel-domain SynthID
confidence score when an external checkout of
aloshdenny/reverse-SynthID
is available. The scorer is not bundled : it is loaded at runtime from your
checkout, and its code remains under the upstream project's non-commercial
Research License.
Option 1: one-command bootstrap (no Docker)
SCRIPTS=skills/remove-ai-marks/scripts
# Clones upstream, creates a venv, and installs scorer-only dependencies.
" $SCRIPTS /setup_synthid.sh "
# Score an image (default checkout: ~/reverse-SynthID).
REVERSE_SYNTHID_DIR= ~ /reverse-SynthID \
~ /reverse-SynthID/.venv/bin/python " $SCRIPTS /score_synthid.py " shot.png
# Or surface the score from inspect / clean (same venv Python).
REVERSE_SYNTHID_DIR= ~ /reverse-SynthID \
~ /reverse-SynthID/.venv/bin/python " $SCRIPTS /inspect_image.py " shot.png
setup_synthid.sh accepts --dir PATH , --ref REF , and --full (install the
full upstream requirements.txt , which adds torch / diffusers for the
upstream VAE bypass this project does not use).
make docker-synthid-build
# Run unprivileged and with a read-only rootfs; the scorer only needs to read
# /data and write to stdout/tmp.
docker run --rm \
--user " $( id -u ) : $( id -g ) " \
--read-only --tmpfs /tmp \
-v " $( pwd ) :/data " \
watermarks-remover-synthid-scorer /data/shot.png
The image is built locally from the upstream source at build time. It is not
published, so it does not redistribute the upstream code.
V4 scoring uses artifacts/spectral_codebook_v4.npz from the upstream checkout
(~220 MB). This is detection/scoring only — it does not remove pixel
watermarks.
Channel
Claude
Gemini/SynthID
OpenAI
Open-LLM
Unicode / edit-based text
Layer A
Layer A
Layer A
Layer A
Statistical sampling text
Layer B best-effort
Layer B best-effort
Layer B if present
Layer B best-effort
C2PA / file metadata
Yes (listed formats)
Yes when present
Yes when present
Yes when present
Pixel image marks
Out of scope
Optional SynthID score (external); removal out of scope
Out of scope
Out of scope
Training backdoors
Out of scope
Out of scope
Out of scope
Out of scope
Details: skills/remove-ai-marks/references/vendor-notes.md , mark-classes.md .
How text marking works (short)
Modern LLM watermarks often hide a signal in which tokens are chosen (generative / sampling bias), not only in invisible characters. Edit-based schemes inject Unicode or synonym rules. File schemes attach C2PA or generator metadata.
Layer A removes edit-based Unicode carriers (testable).
Layer B attacks sampling watermarks via heavy rewrite (best-effort; literature-standard attacks such as paraphrase / back-translation).
File cleaners strip C2PA/XMP/props from supported containers.
Until vendors ship public detectors and keys, no tool can honestly certify “this fails the official check.” Reports must separate verifiable vs best-effort work.
Prefer a non-origin model for Layer B (do not rewrite Claude text with Claude if you are trying to avoid re-stamping).
Disclaimer: what removing a text watermark costs
Text watermarks live in the wording itself : the signal is spread across token choices, so nearly every sentence carries a little of it. Two consequences follow, and they are why Layer B is honestly described as best-effort rather than a magic eraser.
Removal means rewording, not restructuring. Shuffling paragraphs, changing headings, or light touch-ups barely move the signal. Stripping a statistical mark requires rewriting a substantial fraction of the text — sentence by sentence, not section by section.
Rewording degrades the copy. Any rewrite replaces the original word choices with the rewriting model's, which flattens tone, voice, and precision. On production copy (SEO, marketing, client work) that degradation is real and often visible to the people who care most about the writing. It is like taking text from a top-tier model and asking a less capable model to rewrite it from scratch: the result cannot exceed the rewrite model's ceiling.
Which leads to the honest full-circle question:
If the plan is to rewrite the text with a cheaper model anyway, why pay for a premium model in the first place? Generating directly with the cheaper model is simpler, cheaper, and produces the same — or better — end result.
Layer B makes sense when you specifically want the premium model's thinking and drafting and accept a rewrite pass to satisfy a hygiene or privacy requirement — not as a cheap route to mark-free text.
Quality matters more than hygiene: use the lossless path — Layer A Unicode scrub plus the file metadata cleaners — and keep the original prose.
Rewriting anyway: use a non-origin model (rewriting with the origin model can re-stamp the text), and remember residual risk remains — no tool can certify a vendor detector will fail.
Format
Inspect
Clean
PNG / JPEG
C2PA chunks / APP11, AI XMP hints
Drop metadata segments
SVG
<metadata> , XMP
Strip blocks
PDF
Byte/XMP + optional tools
exiftool preferred; degraded without it
DOCX
docProps / customXml
Scrub props, drop customXml
ODT
meta.xml
Drop generator / AI-ish meta
HTML
meta, JSON-LD, data-ai*
Strip tags/attrs
Markdown
YAML frontmatter AI keys
Drop keys + Layer A body
Pixel-domain watermark removal and C2PA soft binding (in-content watermark that can re-link a remote Content Credentials manifest after metadata is stripped) remain out of scope . Stripping hard-bound C2PA does not clear those channels. An optional local SynthID scorer is available for detection only (see above).
This tool reports verifiable removals (Unicode counts, metadata actions) and best-effort Layer B rewrites. It cannot certify that vendor detectors will fail.
To check residual signals yourself (optional, external):
Industry two-layer context (C2PA + imperceptible watermark): Institute of AI PM guide .
Option
Removes
Notes
Unicode scrub (Layer A)
ZWSP, bidi, tags, exotic spaces, …
Safe default for text
Rewrite (Layer B)
Statistical token marks (best-effort)
Always offered by skill; costs style — see Disclaimer
Container/metadata strip
File provenance
See format table
Open-weight local models
Avoid re-stamping with origin model
Operational alternative
Matrix: skills/remove-ai-marks/references/removal-matrix.md .
See skills/remove-ai-marks/references/ethics.md . For privacy and research on your content — not academic fraud or false “human-written” claims.
python3 -m venv .venv && .venv/bin/pip install pytest
.venv/bin/python -m pytest # or: make test
make smoke # quick CLI smoke on fixtures
Changelog
v0.3.2 — security hardening (safe writes, HTTP client, CI supply chain)
Safe, atomic output writes : every cleaner now writes via temp-file + atomic rename ( safe_write_bytes / safe_write_text ), refuses symlinked destinations, and creates .bak backups through the same safe path — pre-placed symlinks (e.g. in /tmp or download dirs) can no longer redirect a clean write onto an arbitrary file
rewrite_text.py HTTP client hardening : redirects are refused outright, so an API key in the Authorization header can never be re-sent to an unvalidated host; non-loopback endpoints are denied by default (opt in with --allow-remote or WATERMARKS_REWRITE_ALLOW_REMOTE=1 ); only http(s) schemes are accepted; --api-key was removed — keys are env-only via WATERMARKS_REWRITE_API_KEY
Resource caps : default max input 1 GiB → 256 MiB, new 64 MiB stdin cap, DOCX/ODT zip budget 512 MiB → 128 MiB, and RLIMIT_AS / RLIMIT_FSIZE applied to exiftool/c2patool/SynthID subprocesses (all caps env-overridable)
Supply chain : CI actions SHA-pinned with permissions: contents: read , pinned dev d

[truncated]
