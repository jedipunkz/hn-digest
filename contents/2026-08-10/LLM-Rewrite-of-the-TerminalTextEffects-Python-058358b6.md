---
source: "https://github.com/omacom-io/ttfx"
hn_url: "https://news.ycombinator.com/item?id=49243229"
title: "LLM Rewrite of the TerminalTextEffects Python"
article_title: "GitHub - omacom-io/ttfx: Terminal text effects as a single static binary — a parity-exact Rust port of terminaltexteffects · GitHub"
author: "m-novikov"
captured_at: "2026-08-10T14:14:28Z"
capture_tool: "hn-digest"
hn_id: 49243229
score: 2
comments: 1
posted_at: "2026-08-10T13:17:16Z"
tags:
  - hacker-news
  - translated
---

# LLM Rewrite of the TerminalTextEffects Python

- HN: [49243229](https://news.ycombinator.com/item?id=49243229)
- Source: [github.com](https://github.com/omacom-io/ttfx)
- Score: 2
- Comments: 1
- Posted: 2026-08-10T13:17:16Z

## Translation

タイトル: TerminalTextEffects Python の LLM 書き換え
記事のタイトル: GitHub - omacom-io/ttfx: 単一の静的バイナリとしてのターミナル テキスト効果 — ターミナルテキストエフェクトのパリティ正確な Rust ポート · GitHub
説明: 単一の静的バイナリとしてのターミナル テキスト効果 — ターミナルテキストエフェクトのパリティ正確な Rust ポート - omacom-io/ttfx

記事本文:
GitHub - omacom-io/ttfx: 単一の静的バイナリとしてのターミナル テキスト効果 — ターミナルテキストエフェクトのパリティ正確な Rust ポート · GitHub
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
オマコム-io
/
ttfx
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
90 コミット 90 コミット .github/ workflows .github/ workflows docs docs パッケージ化 パッケージ化 src src テスト t

ests ツール ツール .gitattributes .gitattributes .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス通知 通知 README.md README.md logo_end-54.png logo_end-54.png logo_end-55.png logo_end-55.png logo_end-56.png logo_end-56.png logo_end-57.png logo_end-57.png logo_end-58.png logo_end-58.png logo_end-59.png logo_end-59.png logo_end-60.png logo_end-60.png logo_end-61.png logo_end-61.png logo_end-62.png logo_end-62.png logo_mid.png logo_mid.png plan.md plan.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ターミナル テキストは単一の静的バイナリとして機能します。テキストをパイプで入力し、効果を選択します。
ls -la | ttfx復号化
猫バナー.txt | ttfx ビーム
運勢 | ttfx --ランダム効果
git log --oneline -10 | ttfx マトリックス
支払われるべきところにクレジットを与える
これは TerminalTextEffects のポートです
(TTE) ChrisBuilds 著。あらゆるエフェクト、アニメーション エンジン、
コマンドライン インターフェイスは彼らの設計です。このプロジェクトはその機能を Rust に変換します。
そして芸術にそれ自体は何も加えません。ここにあるものが気に入ったら、オリジナルにスターを付けてください。
TTE は MIT ライセンスを取得しており、このポートも同様です。オリジナルの著作権は以下に保存されます
ライセンスと通知。エフェクトのアイデアは、それが所属する上流にファイルしてください。
TTE は Python パッケージです。それは図書館にとっては正しい要求ですが、そこに住む貝殻のおもちゃにとっては正しいことです。
プロンプト パイプラインとは、インタープリター、インストール ステップ、および実行前の最大 90 ミリ秒のインポートを意味します。
最初のフレーム。 ttfx は、約 1 ミリ秒で開始される依存関係のないバイナリです。
その違いこそが、これが存在する理由なのです。フルスクリーンのキャンバスではより重いエフェクトがかかります
Python では高いフレーム レートを維持できません。
37 のエフェクトすべてで、スピードアップの中央値は 9.6 倍 (範囲 4.5 ～ 21.6 倍) です。
これはパリティ ポートであり、精神的な再実装ではありません。同じ入力、構成、および
ランダムな描画、ttfx はバイト同一のフレームを生成します

Python のオリジナルに準拠 — 検証済み
目視ではなく、固定された上流チェックアウト (v0.15.0) に対して CI で機械的に。
それを可能にするということは、アップストリームの癖を「修正」するのではなく、意図的に再現することを意味します。
Python のバンカーの丸め、float ではなく整数のフロア除算から構築されたグラデーション
補間、最終セグメントを削除するベジェ弧長近似、およびループ
ティックごとに完了したと報告するシーン。それらはカタログ化されています
プラン.md ; Python の順序なし反復を特定する必要があった場所は次のとおりです。
docs/ordering-inventory.md にあります。
2 つの意図的な違い。乱数生成は CPython とビット互換性がありません —
ttfx は xohiro256++ を使用するため、 --seed は ttfx 内で再現可能ですが、Python のものとは一致しません
メルセンヌ・ツイスター。 (パリティ ハーネスは共有 PRNG を両側に交換します。これにより、
フレーム比較はまったく可能です。) また、Python プラグイン エフェクトはサポートされていません。
それらをロードするインタープリターはありません。
<プロデューサー> | ttfx [ターミナルオプション] <効果> [効果オプション]
ttfx --help # 37 個すべてのエフェクトとターミナル オプション
ttfx <effect> --help # 1 つのエフェクトのオプション
ttfx --random-effect # 驚かせてください (--include-effect / --exclude-effect をフィルタリングします)
ttfx --print-completion bash|zsh
端末オプション (キャンバス サイズとアンカー、カラー処理、フレーム レート、テキストの折り返し) が追加されました。
エフェクト名の前。その後のエフェクトオプション。オプション名とデフォルトは一致しますので、
既存の呼び出しは、バイナリ名が交換された状態で機能します。
カーゴビルド --release
カーゴ ビルド --release --target x86_64-unknown-linux-musl # 静的、~3.3 MB
パリティ スイートを実行するには、python3 とアップストリームのコピーが必要です。
./tools/parity/fetch_reference.sh # ピン留めされたコミットで TTE をクローンします
./tools/parity/run_suite.sh
アップストリームはここでは販売されていません - ハーネス

それは彼らのコードなので、それを取得します。
Linux と macOS。元々はオマルキーのために作られました。何もターゲットにしない
特定の libc と CI は両方のプラットフォームでテストと CLI コーパスを実行します。バイト正確
パリティ スイートは Linux/glibc に固定されたまま — Apple の libm はいくつかの超越関数を丸めます
last-ulp は異なります。量子化は実際のフレームに隠蔽されますが、ビット正確な比較が行われます。
表面化するだろう。
MIT — このプロジェクトの著作権とオリジナルの両方が含まれる LICENSE を参照してください。
著作権は TerminalTextEffects にあり、完全な帰属については通知してください。
単一の静的バイナリとしてのターミナル テキスト効果 — ターミナルテキストエフェクトのパリティ正確な Rust ポート
Readme MIT ライセンス アクティビティ カスタム プロパティ スター
4 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Terminal text effects as a single static binary — a parity-exact Rust port of terminaltexteffects - omacom-io/ttfx

GitHub - omacom-io/ttfx: Terminal text effects as a single static binary — a parity-exact Rust port of terminaltexteffects · GitHub
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
omacom-io
/
ttfx
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
90 Commits 90 Commits .github/ workflows .github/ workflows docs docs packaging packaging src src tests tests tools tools .gitattributes .gitattributes .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE NOTICE NOTICE README.md README.md logo_end-54.png logo_end-54.png logo_end-55.png logo_end-55.png logo_end-56.png logo_end-56.png logo_end-57.png logo_end-57.png logo_end-58.png logo_end-58.png logo_end-59.png logo_end-59.png logo_end-60.png logo_end-60.png logo_end-61.png logo_end-61.png logo_end-62.png logo_end-62.png logo_mid.png logo_mid.png plan.md plan.md View all files Repository files navigation
Terminal text effects as a single static binary. Pipe text in, pick an effect:
ls -la | ttfx decrypt
cat banner.txt | ttfx beams
fortune | ttfx --random-effect
git log --oneline -10 | ttfx matrix
Credit where it's due
This is a port of TerminalTextEffects
(TTE) by ChrisBuilds . Every effect, the animation engine,
and the command-line interface are their design — this project translates that work to Rust
and adds nothing of its own to the art. If you like what you see here, star the original.
TTE is MIT licensed and so is this port; the original copyright is preserved in
LICENSE and NOTICE . Please file effect ideas upstream, where they belong.
TTE is a Python package. That's the right call for a library, but for a shell toy that lives in
your prompt pipeline it means an interpreter, an install step, and ~90 ms of import before the
first frame. ttfx is one dependency-free binary that starts in ~1 ms.
That difference is the whole reason this exists. On a fullscreen canvas the heavier effects
can't hold a high frame rate under Python:
Across all 37 effects the median speedup is 9.6× (range 4.5×–21.6×).
This is a parity port , not a reimplementation-in-spirit. Given the same input, config, and
random draws, ttfx produces byte-identical frames to the Python original — verified
mechanically in CI against a pinned upstream checkout (v0.15.0), not by eyeballing.
Making that possible meant reproducing upstream's quirks deliberately, not "fixing" them:
Python's banker's rounding, gradients built from integer floor division rather than float
interpolation, a bezier arc-length approximation that drops its final segment, and looping
scenes that report themselves complete on every tick. They're catalogued in
plan.md ; the places where Python's unordered iteration had to be pinned down are
in docs/ordering-inventory.md .
Two deliberate differences. Random number generation is not bit-compatible with CPython —
ttfx uses xoshiro256++, so --seed is reproducible within ttfx but won't match Python's
Mersenne Twister. (The parity harness swaps a shared PRNG into both sides, which is what makes
frame comparison possible at all.) And Python plugin effects aren't supported, since there's
no interpreter to load them.
<producer> | ttfx [terminal options] <effect> [effect options]
ttfx --help # all 37 effects and the terminal options
ttfx <effect> --help # options for one effect
ttfx --random-effect # surprise me (--include-effects / --exclude-effects to filter)
ttfx --print-completion bash|zsh
Terminal options (canvas size and anchoring, color handling, frame rate, text wrapping) go
before the effect name; effect options after it. Option names and defaults match tte , so
existing invocations work with the binary name swapped.
cargo build --release
cargo build --release --target x86_64-unknown-linux-musl # static, ~3.3 MB
Running the parity suites needs python3 and a copy of upstream:
./tools/parity/fetch_reference.sh # clones TTE at the pinned commit
./tools/parity/run_suite.sh
Upstream is not vendored here — the harness fetches it, because it's their code.
Linux and macOS. Built for Omarchy originally; nothing targets a
specific libc, and CI runs the tests and CLI corpus on both platforms. The byte-exact
parity suites stay pinned to Linux/glibc — Apple's libm rounds a few transcendentals a
last-ulp differently, which quantization hides in real frames but a bit-exact comparison
would surface.
MIT — see LICENSE , which carries both this project's copyright and the original
TerminalTextEffects copyright, and NOTICE for the attribution in full.
Terminal text effects as a single static binary — a parity-exact Rust port of terminaltexteffects
Readme MIT license Activity Custom properties Stars
4 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
