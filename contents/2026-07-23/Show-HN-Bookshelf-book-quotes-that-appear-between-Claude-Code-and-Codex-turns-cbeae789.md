---
source: "https://github.com/Amal-David/bookshelf"
hn_url: "https://news.ycombinator.com/item?id=49028693"
title: "Show HN: Bookshelf – book quotes that appear between Claude Code and Codex turns"
article_title: "GitHub - Amal-David/bookshelf: book quotes that appear between Claude Code and Codex turns · GitHub"
author: "amaldavid"
captured_at: "2026-07-23T22:51:49Z"
capture_tool: "hn-digest"
hn_id: 49028693
score: 3
comments: 0
posted_at: "2026-07-23T22:06:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Bookshelf – book quotes that appear between Claude Code and Codex turns

- HN: [49028693](https://news.ycombinator.com/item?id=49028693)
- Source: [github.com](https://github.com/Amal-David/bookshelf)
- Score: 3
- Comments: 0
- Posted: 2026-07-23T22:06:47Z

## Translation

タイトル: HN を表示: 本棚 – クロード コードとコーデックスのターンの間に表示される本の引用
記事のタイトル: GitHub - Amal-David/bookshelf: クロード コードとコーデックスのターンの間に表示される本の引用 · GitHub
説明: クロード コードとコーデックスのターンの間に現れる本の引用 - アマルデイヴィッド/本棚
HN テキスト: これは私が自分のために構築した小さな生活の質のプロジェクトで、現在 3 か月以上使用していますが、フォークしてスキルとして個別に公開することにしました。私たち全員が昼も夜もそれに釘付けになっていることを考えると、コードやツール呼び出しでいっぱいの端末を見る代わりに、本の引用を見ることは新鮮な空気の息吹のように感じられました。エージェントが私たちのために何かを大量に出し続けている間、端末に表示されるランダムな引用文によって、時折立ち止まったり幸せな瞬間を過ごしていただければ幸いです。ぜひお気軽に試してみて、ご意見やフィードバックをお聞かせください。

記事本文:
GitHub - Amal-David/bookshelf: クロード コードとコーデックスのターンの間に表示される本の引用 · GitHub
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
アマル・デイヴィッド
/
本棚
公共
通知
通知設定を変更するにはサインインする必要があります
アディ

ナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .codex-plugin .codex-plugin .github/ workflows .github/ workflows アセット アセット 本棚 本棚ドキュメント ドキュメント拡張機能 拡張機能 フック フック プロトコル プロトコル スクリプト スクリプト サイト サイトスキル/ 本棚スキル/ 本棚テスト テスト.gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md DATA.md DATA.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md __init__.py __init__.py package.json package.json plugin.yaml plugin.yaml pyproject.toml pyproject.tomlすべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェント セッション内に表示される書籍の引用。
Bookshelf は、静かで視野を広げる書籍の引用をコーデックス内に配置します。
完了した数ターンごとにクロード コード セッションが行われます。端末を見つめるのではなく
エージェントが作業している間にチャーンすると、ローカルで選択された小さな文学リセットが 1 つ得られます。
プロンプト、コード、トランスクリプトをどこにも送信する必要はありません。
また、完全なターミナル ライブラリ、リーディング リスト、検索、オンデマンドも含まれています。
関連する引用。それらはアンビエントモーメントの背後にあるライブラリであり、メインではありません
イベント。出荷された合計はカタログから取得されます。「カタログ数」を参照してください。
Bookshelf ランディング ページにアクセスします。
アンビエント モードはオプションであり、デフォルトではオフになっています。
本棚アンビエント有効化 --cadence 5 --intent リファクタリング
本棚の周囲の状態
本棚のアンビエントを無効にする
ホストとスクリプトは、プロセスごとにアンビエントの動作をオーバーライドすることもできます。
環境、構成編集は必要ありません: BOOKSHELF_AMBIENT_ENABLED ( 1 / 0 /
true / false )、BOOKSHELF_AMBIENT_CADENCE (正の整数)、および
BOOKSHELF_DATA_HOME (設定と周囲の状態を別の状態にリダイレクトします)
ディレクトリ)。保存された c よりも環境値が優先されます。

オンフィグ。空白または無効な値
無視されます。
エージェント統合をインストールしても、アンビエント モードは有効になりません。アダプターには以下が含まれます
Bookshelf は独自の誤りを抱えていますが、Bookshelf はホストについて絶対的な主張を行っていません。
回ります。
意図的に使用する場合は、 Bookshelf スキルを呼び出すか、 bookshelf quote を実行します。
本棚引用 --intent リファクタリング、または本棚フィードバック up|down 。
Bookshelf quote --intent リファクタリングは明示的なオンデマンド パスです。アンビエント
停止フックは、 ambient Enable --intent によって保存された同様に明示的なテーマを使用します。
完了ターン イベントにタスク コンテキストが含まれているようには見せません。両方のパス
許可リストに登録されたインテントをローカル タグにマッピングします。どちらもコマンド、パス、
プロンプト、トランスクリプト、コード、ツール引数、モデル出力、またはネットワークの作成
電話する。アンビエント配信はオプションで、デフォルトではオフになっており、
安全な境界引用符は使用できません。
関連する代替案は、50 件の最近の引用枠内で繰り返される前にローテーションされます。
アンビエント行は 220 UTF-8 バイトと空白で区切られた 32 ワードに制限されます。
明示的なオンデマンド コマンドは、より広範なコンパクト ディスプレイ バジェットを維持します。
バージョン化された 176 ケースの評価は、そのための作成された回帰契約です。
インテントからタグへのマッピングと決定論的ランカー。精度の高いメトリクスがキャッチ
ランキングの変動。これらは人間が評価した文学的または意味的な主張ではありません
関連性。
Bookshelf には Python 3.10 以降が必要ですが、ランタイムの依存関係はありません。
pipx install git+https://github.com/Amal-David/bookshelf.git
# 最初の PyPI リリース後:
# pipx アンビエント本棚をインストールします
次に、bookshelf を実行して対話型ライブラリを開きます。
コーデックス プラグイン マーケットプレイスに Amal-David/bookshelf を追加
コーデックスプラグイン追加bookshelf@bookshelf
Bookshelf は、スキルとフェイルソフト Stop を備えた Codex プラグインとしてパッケージ化されています。
フック。 Codex は、アプリと CLI サーフェス全体でプラグイン設定を共有します。の
フック

ローカルのリズムが期限になると、コンパクトな systemMessage を返します。
/plugin マーケットプレイス add Amal-David/bookshelf
/プラグインインストールbookshelf@bookshelf
マーケットプレイスはクロード コード 2.1.212 で厳密に検証し、次のように宣言します。
Codex 統合で使用される同じオプションの停止ケイデンス。
pi インストール git:github.com/Amal-David/bookshelf
Pi 0.57 は、正規スキルとそのagent_end拡張機能をインストールしてリストします。
孤立した家から。ライブネイティブ通知は要求されません。
hermes プラグインは Amal-David/bookshelf をインストールします --enable
Hermes 0.16 は、クリーンな作業ディレクトリから Bookshelf をインストールしてロードします。
正規スキルとそのtransform_llm_outputフックを登録します。ライブ
認証されたターンと追加された出力は要求されません。
Bookshelf は、インストール、マニフェスト、ローダー、アダプターの動作を検証します。
Codex、Claude Code、Pi、Hermes のローカルで利用可能な最も強い境界。
Codex フックと Claude フックは、それぞれにインストールされているプラグイン ルートを解決します。
スキルの scripts/quote.py ラッパーは、バンドルされた Python パッケージを相対的に解決します。
それ自体に対して。どちらも、ブックシェルフ コマンドをグローバルにインストールする必要はありません。古い
マニュアルのbookshelf/skill/hook.pyとcodex_notify.pyのパスには互換性があります
アダプタのみであり、新規インストールでは非推奨となります。
protocol/ambient-companion-v1.schema.json とそのサンプルは次のように出荷されます。
インテグレータ向けの不活性プロトコル データ。 Bookshelf は実行時にそれらをインポートしません
また、アンビエント配信はネットワーク要求を行うことはありません。
生成されたカタログ数により、カタログ化されたカタログ数が区別されます。
引用によって参照されている作品の書籍。レガシーレコードは明示的にマークされています
レガシー未検証 ; 585 件のプライマリ ソースにリンクされた v2 レコードが保留中のままです
リポジトリ リンク、ソース ロケータ、抽出スナップショットを確認して保持する
ダイジェスト、抜粋ダイジェスト、権利、およびレビュー保留中のメタデータ。それらのソースリンク

不変コミットに固定されていないため、v2 レコードを記述してはなりません
出典が確認されたものとして。
引用符には、作業に関連するエントリを優先するために使用されるコンテキスト タグが含まれます。
デバッグ、構築、レビュー、または出荷。これはすべてがそうであるという主張ではありません
引用は検証または厳選されます。
出所と修正方針については、DATA.md を参照してください。
お気に入り、読書リスト、統計、フックカウンター、アンビエント設定はそのまま残ります
Bookshelf という名前のプラットフォームの application-data ディレクトリの下にあるローカル。
フル解像度の端末デモをご覧ください。
録音は、Opus 4.8 による認証されたクロード コード 2.1.217 セッションです。
使い捨ての rewrite-prod-in-rust-by-lunch クレート。クロードは本物を読む
README と Rust ソース、危険なロールバック ウィンドウを 0 から 30 に変更
数分、実際の Cargo テスト スイートを実行して両方のテストに合格し、独自のテストを作成します
ドライデプロイジョークを実行し、実際の Bookshelf Stop フックをトリガーします。セットアップ
適度なズームと 4 秒間のハイライトの前に、ツールの作業が圧縮されます
見積もりを保留してください。アカウント識別子は編集されています。エージェントの行動、
それ以外の場合、出力と単一の停止イベントは変更されません。なぜなら、一般人は
プラグインは停止のみで、完了した応答ごとに 1 つの引用を正確に表示します。
ランディング ページのソースは site/ にあります。同じものから作られています
森、紙、バラのビジュアル システムをデモとして使用します。
python3 -m 単体テスト 検出 -s テスト -v
python3 -m COMPILEALL -Q 本棚フック スクリプト テスト
python3 -m ピップホイール。 --no-deps --no-build-isolation --wheel-dir dist
カタログの修正とコードについては CONTRIBUTING.md を参照してください。
貢献。
アプリケーション コードは MIT ライセンスを取得しています。本のタイトル、著者名、概要など
引用されたテキストには、別の出所と権利に関する考慮事項があり、以下に説明されています。
DATA.md 。
クロード・コードとコーデックスのターンの間に現れる本の引用
本棚

lf-8dz.pages.dev トピック
Readme MIT ライセンス
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

book quotes that appear between Claude Code and Codex turns - Amal-David/bookshelf

This is a small quality of life project i built for myself and using it for 3+ months now, decided to fork it out and publish as a skill separately. Instead of seeing terminals filled with code and tool calls, seeing a book quote felt like a breath of fresh air considering all of us are glued to it day and night. Hope the random quotes that pop up in your terminals lets you have a moment of pause or happiness every now and then while the agents keep churning out stuff for us. Please feel free to give it a spin and let me know your thoughts and feedback!

GitHub - Amal-David/bookshelf: book quotes that appear between Claude Code and Codex turns · GitHub
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
Amal-David
/
bookshelf
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .codex-plugin .codex-plugin .github/ workflows .github/ workflows assets assets bookshelf bookshelf docs docs extensions extensions hooks hooks protocol protocol scripts scripts site site skills/ bookshelf skills/ bookshelf tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md DATA.md DATA.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md __init__.py __init__.py package.json package.json plugin.yaml plugin.yaml pyproject.toml pyproject.toml View all files Repository files navigation
Book quotes that appear inside your coding-agent sessions.
Bookshelf puts a quiet, perspective-widening book quote inside your Codex or
Claude Code session every few completed turns. Instead of staring at terminal
churn while an agent works, you get one small literary reset—selected locally,
without sending your prompts, code, or transcript anywhere.
It also includes a full terminal library, reading lists, search, and on-demand
relevant quotes. Those are the library behind the ambient moment, not the main
event. Shipped totals come from the catalog: see catalog counts .
Visit the Bookshelf landing page .
Ambient mode is optional and off by default:
bookshelf ambient enable --cadence 5 --intent refactor
bookshelf ambient status
bookshelf ambient disable
Hosts and scripts can also override ambient behavior per process through the
environment, no config edit needed: BOOKSHELF_AMBIENT_ENABLED ( 1 / 0 /
true / false ), BOOKSHELF_AMBIENT_CADENCE (positive integer), and
BOOKSHELF_DATA_HOME (redirects config and ambient state to another
directory). Environment values win over saved config; blank or invalid values
are ignored.
Installing an agent integration does not enable ambient mode. Adapters contain
their own errors, but Bookshelf does not make an absolute claim about any host
turn.
For deliberate use, invoke the Bookshelf skill or run bookshelf quote ,
bookshelf quote --intent refactor , or bookshelf feedback up|down .
bookshelf quote --intent refactor is the explicit on-demand path. Ambient
Stop hooks use the equally explicit theme saved by ambient enable --intent ;
they do not pretend a completed-turn event contains task context. Both paths
map an allow-listed intent to local tags. Neither reads commands, paths,
prompts, transcripts, code, tool arguments, model output, or makes a network
call. Ambient delivery is optional, off by default, and fails closed when a
safe bounded quote is unavailable.
Relevant alternatives rotate before repeats inside a 50-quote recent window.
Ambient lines are capped at 220 UTF-8 bytes and 32 whitespace-delimited words;
the explicit on-demand command retains its wider compact-display budget.
The versioned 176-case evaluation is an authored regression contract for that
intent-to-tag mapping and the deterministic ranker. Its precision metrics catch
ranking drift; they are not a human-rated claim of literary or semantic
relevance.
Bookshelf requires Python 3.10 or newer and has no runtime dependencies.
pipx install git+https://github.com/Amal-David/bookshelf.git
# After the first PyPI release:
# pipx install ambient-bookshelf
Then run bookshelf to open the interactive library.
codex plugin marketplace add Amal-David/bookshelf
codex plugin add bookshelf@bookshelf
Bookshelf is packaged as a Codex plugin with a skill and a fail-soft Stop
hook. Codex shares plugin configuration across its app and CLI surfaces; the
hook returns a compact systemMessage when the local cadence is due.
/plugin marketplace add Amal-David/bookshelf
/plugin install bookshelf@bookshelf
The marketplace validates strictly with Claude Code 2.1.212 and declares the
same optional Stop cadence used by the Codex integration.
pi install git:github.com/Amal-David/bookshelf
Pi 0.57 installs and lists the canonical skill plus its agent_end extension
from an isolated home. A live native notification is not claimed.
hermes plugins install Amal-David/bookshelf --enable
Hermes 0.16 installs and loads Bookshelf from a clean working directory,
registering the canonical skill plus its transform_llm_output hook. A live
authenticated turn and appended output are not claimed.
Bookshelf validates install, manifest, loader, and adapter behavior at the
strongest locally available boundary for Codex, Claude Code, Pi, and Hermes.
The Codex and Claude hooks resolve their own installed plugin roots, and the
skill's scripts/quote.py wrapper resolves the bundled Python package relative
to itself; neither requires a globally installed bookshelf command. The old
manual bookshelf/skill/hook.py and codex_notify.py paths are compatibility
adapters only and are deprecated for new installs.
protocol/ambient-companion-v1.schema.json and its example are shipped as
inert protocol data for integrators. Bookshelf does not import them at runtime
and ambient delivery never makes network requests.
The generated catalog counts distinguish catalogued
books from works referenced by quotes. Legacy records are explicitly marked
legacy-unverified ; 585 primary-source-linked v2 records remain pending human
review and retain their repository link, source locator, extraction-snapshot
digest, excerpt digest, rights, and review-pending metadata. Those source links
were not pinned to immutable commits, so the v2 records must not be described
as source-verified.
Quotes carry context tags used to prefer relevant entries for work such as
debugging, building, reviewing, or shipping. This is not a claim that all
quotes are verified or curated.
See DATA.md for provenance and the correction policy.
Favorites, reading lists, statistics, hook counters, and ambient settings remain
local under the platform application-data directory named bookshelf .
Watch the full-resolution terminal demo .
The recording is an authenticated Claude Code 2.1.217 session with Opus 4.8 in
the disposable rewrite-prod-in-rust-by-lunch crate. Claude reads the real
README and Rust source, changes the risky rollback window from zero to 30
minutes, runs the real Cargo test suite with both tests passing, writes its own
dry deploy joke, and then triggers the actual Bookshelf Stop hook. The setup
and tool work are compressed before a modest zoom and four-second highlighted
hold on the quote. The account identifier is redacted; the agent's actions,
output, and single Stop event are otherwise untouched. Because the public
plugin is Stop-only, it truthfully shows one quote per completed response.
The landing-page source lives in site/ . It is built from the same
forest, paper, and rose visual system as the demo.
python3 -m unittest discover -s tests -v
python3 -m compileall -q bookshelf hooks scripts tests
python3 -m pip wheel . --no-deps --no-build-isolation --wheel-dir dist
See CONTRIBUTING.md for catalog corrections and code
contributions.
The application code is MIT licensed. Book titles, author names, summaries, and
quoted text have separate provenance and rights considerations described in
DATA.md .
book quotes that appear between Claude Code and Codex turns
bookshelf-8dz.pages.dev Topics
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
