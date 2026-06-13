---
source: "https://github.com/Neelagiri65/equiv"
hn_url: "https://news.ycombinator.com/item?id=48515830"
title: "Equiv, check that an AI refactor did not change what your code does"
article_title: "GitHub - Neelagiri65/equiv: Deterministic checker for behaviour-preserving code changes. Signed, re-runnable receipts; PR gate; single static binary. · GitHub"
author: "Srinathprasanna"
captured_at: "2026-06-13T12:43:34Z"
capture_tool: "hn-digest"
hn_id: 48515830
score: 1
comments: 0
posted_at: "2026-06-13T10:46:39Z"
tags:
  - hacker-news
  - translated
---

# Equiv, check that an AI refactor did not change what your code does

- HN: [48515830](https://news.ycombinator.com/item?id=48515830)
- Source: [github.com](https://github.com/Neelagiri65/equiv)
- Score: 1
- Comments: 0
- Posted: 2026-06-13T10:46:39Z

## Translation

タイトル: Equiv、AI リファクタリングによってコードの動作が変更されていないことを確認してください
記事のタイトル: GitHub - Neelagiri65/equiv: 動作を保持するコード変更のための決定論的チェッカー。署名済みの再発行可能な領収書。 PRゲート。単一の静的バイナリ。 · GitHub
Description: Deterministic checker for behaviour-preserving code changes.署名済みの再発行可能な領収書。 PRゲート。単一の静的バイナリ。 - ニーラギリ65/相当

記事本文:
GitHub - Neelagiri65/equiv: 動作を維持するコード変更のための決定論的チェッカー。署名済みの再発行可能な領収書。 PRゲート。単一の静的バイナリ。 · GitHub
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
ニーラギリ65
/
相当する
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
この GitHub アクションをプロジェクトで使用する このアクションを既存のワークフローに追加するか、新しいワークフローを作成します マーケットプレイスで表示する メイン ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .github/ workflows .github/ workflows 資産 アセット 適合性/ 有効な適合性/ 有効 クレート クレート ドキュメント ドキュメント サンプル サンプル スクリプト スクリプト .equiv-review .equiv-review .equiv-review.example .equiv-review.example .gitignore .gitignore AGENTS.md AGENTS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md action.yml action.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM で書かれたコードをレビューするのは LLM だけではありません。
equiv は、変更された関数を以前のバージョンに対して同じ環境で実行します。
決定的に生成された入力と、動作が変化したかどうかをレポートします。もし
そうであれば、異なる部分の正確な入力を取得できます。どちらの方法でも、
再現可能な署名付き領収書: どのマシンでもチェックを再実行すると、同じ結果が得られます。
モデルの意見を信用することなく、バイトごとに答えます。
現在、ほとんどのコードは AI によって記述され、AI によってレビューされます。 「これっぽいね」というモデル
大丈夫です」は検証ではありません。自分で再実行できる確定的なチェックです。
マニフェスト内の PR 全体にわたって動作を保持する必要がある関数をリストします。
リポジトリのルートにあります。各行の形式は次のとおりです。
<file> : <function> : <arg type> 、ここで arg の型は int 、 str 、または
list[int] 、カンマ区切り:
src/math.py : 合計 : int
ワークフローを .github/workflows/equiv-review.yml に追加します。
上: pull_request
権限: { 内容: 読み取り、プルリクエスト: 書き込み、ID トークン: 書き込み }
ジョブ:
レビュー：
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
と: { フェッチ深さ: 0 }
- 使用: Neelagiri65/equiv@v0.1.0

付き: { キーレス: "true" }
@main ではなくリリースされたタグ ( @v0.1.0 ) にピン留めすることで、実行を再現可能にします
そしてあなたの下で変わらないでください。
各 PR はコメントを受け取ります。変更されたすべての関数はそのバージョンに対してテストされます
根元の枝に。動作を維持する変更はパスされます。実現する変化
not は、2 つのバージョンを区別する入力とともに報告されます。それ
チェックに失敗します。領収書は Sigstore キーレス署名で署名されます。
鍵がありません。これらは cosign で検証できます。
カール --proto '=https' --tlsv1.2 -LsSf \
https://github.com/Neelagiri65/equiv/releases/latest/download/equiv-cli-installer.sh |しー
equiv reviewcandidate.pyreference.py <関数> <引数の種類>
equiv verify-receipt <signed-receipt-hex>
終了コード: 0 は同等、1 は出力された反例と分岐、2
確認できませんでした。
equiv は、参照に対して関数の動作の等価性をチェックします。
決定的に生成された入力。これは制限されたランダム テストではなく、
徹底的な検証: 合格は、生成されたデータに相違が見つからなかったことを意味します。
入力。次のような入力に対してのみ現れるエッジケースを見逃す可能性があります。
は生成されませんでした。意図、アーキテクチャ、セキュリティはチェックされません。それ
比較する基準がない新しい機能を判断することはできません。あ
結果が合格したということは、テストされた入力で動作が維持されたことを意味します。そうではありません
変更が正しいことを意味します。 Supported input types in this version are int ,
str と list[int] 。
入力の生成と判定は、固定シードから Rust で計算されます。の
言語ランタイムは評価器としてのみ使用され、何かを決定することはありません。
レシートに届きます。受領書はホスト間で同一です。領収書は次のとおりです。
ローカルの ed25519 キーまたはキーレス Sigstore (OIDC) で署名されます。キーレス
パスは署名を検証可能な CI ID rath にバインドします。

保管されているものよりも
秘密。このツールは、実行時の依存関係のない単一の静的バイナリです。
macOS、Linux、Windows 用に事前構築されています。
docs/signing-model.md : ed25519 とキーレス Sigstore を使用した領収書の署名。
docs/RELEASING.md : Cargo-dist を使用して事前に構築されたバイナリを構築します。
crates/ : Rust ワークスペース ( equiv-core 、 equiv-engine 、 equiv-review 、 equiv-cli )。
動作を維持するコード変更のための決定論的チェッカー。署名済みの再発行可能な領収書。 PRゲート。単一の静的バイナリ。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
v0.1.0
最新の
2026 年 6 月 13 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Deterministic checker for behaviour-preserving code changes. Signed, re-runnable receipts; PR gate; single static binary. - Neelagiri65/equiv

GitHub - Neelagiri65/equiv: Deterministic checker for behaviour-preserving code changes. Signed, re-runnable receipts; PR gate; single static binary. · GitHub
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
Neelagiri65
/
equiv
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .github/ workflows .github/ workflows assets assets conformance/ valid conformance/ valid crates crates docs docs examples examples scripts scripts .equiv-review .equiv-review .equiv-review.example .equiv-review.example .gitignore .gitignore AGENTS.md AGENTS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md action.yml action.yml View all files Repository files navigation
An LLM should not be the only thing reviewing LLM-written code.
equiv runs a changed function against its previous version on the same
deterministically generated inputs and reports whether the behaviour changed. If
it did, you get the exact input where they differ. Either way you get a
reproducible, signed receipt: re-run the check on any machine and you get the same
answer, byte for byte, without trusting any model's opinion.
Most code is now written by AI and reviewed by AI. A model saying "this looks
fine" is not verification. A deterministic check you can re-run yourself is.
List the functions whose behaviour must be preserved across a PR in a manifest
at the repository root. The format of each line is
<file> : <function> : <arg types> , where arg types are int , str , or
list[int] , comma separated:
src/math.py : total : int
Add the workflow at .github/workflows/equiv-review.yml :
on : pull_request
permissions : { contents: read, pull-requests: write, id-token: write }
jobs :
review :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
with : { fetch-depth: 0 }
- uses : Neelagiri65/equiv@v0.1.0
with : { keyless: "true" }
Pin to a released tag ( @v0.1.0 ) rather than @main so runs are reproducible
and do not change under you.
Each PR receives a comment. Every changed function is tested against its version
on the base branch. A change that preserves behaviour passes. A change that does
not is reported with the input that distinguishes the two versions. That
fails the check. Receipts are signed with Sigstore keyless signing, which stores
no key. They can be verified with cosign .
curl --proto '=https' --tlsv1.2 -LsSf \
https://github.com/Neelagiri65/equiv/releases/latest/download/equiv-cli-installer.sh | sh
equiv review candidate.py reference.py <function> <arg types>
equiv verify-receipt <signed-receipt-hex>
Exit codes: 0 equivalent, 1 diverges with a printed counterexample, 2
could not check.
equiv checks behavioural equivalence of a function against a reference, on
deterministically generated inputs. This is bounded random testing, not
exhaustive verification: a pass means no divergence was found on the generated
inputs. It can still miss an edge case that only shows up for an input that
was not generated. It does not check intent, architecture, security. It
cannot judge new functionality that has no reference to compare against. A
passing result means behaviour was preserved on the tested inputs. It does not
mean the change is correct. Supported input types in this version are int ,
str and list[int] .
Input generation and the verdict are computed in Rust from a fixed seed. The
language runtime is used only as an evaluator and never decides anything that
reaches the receipt. Receipts are identical across hosts. Receipts can be
signed with a local ed25519 key or with keyless Sigstore (OIDC). The keyless
path binds the signature to a verifiable CI identity rather than a stored
secret. The tool is a single static binary with no runtime dependencies,
prebuilt for macOS, Linux and Windows.
docs/signing-model.md : receipt signing with ed25519 and keyless Sigstore.
docs/RELEASING.md : building prebuilt binaries with cargo-dist.
crates/ : the Rust workspace ( equiv-core , equiv-engine , equiv-review , equiv-cli ).
Deterministic checker for behaviour-preserving code changes. Signed, re-runnable receipts; PR gate; single static binary.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
v0.1.0
Latest
Jun 13, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
