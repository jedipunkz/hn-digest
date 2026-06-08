---
source: "https://github.com/dhruvmehra/makememe"
hn_url: "https://news.ycombinator.com/item?id=48443328"
title: "Show HN: Makememe – a meme CLI for your Claude Code"
article_title: "GitHub - dhruvmehra/makememe: Agent-friendly meme CLI over memegen.link — pip install makememe · GitHub"
author: "dhruvme"
captured_at: "2026-06-08T10:40:30Z"
capture_tool: "hn-digest"
hn_id: 48443328
score: 1
comments: 1
posted_at: "2026-06-08T10:02:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Makememe – a meme CLI for your Claude Code

- HN: [48443328](https://news.ycombinator.com/item?id=48443328)
- Source: [github.com](https://github.com/dhruvmehra/makememe)
- Score: 1
- Comments: 1
- Posted: 2026-06-08T10:02:46Z

## Translation

タイトル: HN を表示: Makememe – クロード コード用のミーム CLI
記事のタイトル: GitHub - dhruvmehra/makememe: memegen.link を介したエージェントフレンドリーなミーム CLI — pip install makememe · GitHub
説明: memegen.link を介したエージェントフレンドリーな meme CLI — pip install makememe - dhruvmehra/makememe

記事本文:
GitHub - dhruvmehra/makememe: memegen.link を介したエージェントフレンドリーなミーム CLI — pip install makememe · GitHub
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
ドゥルブメーラ
/
メイクメメ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ママ

ブランチ内 タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット .github/ workflows .github/ workflows docs docs サンプル サンプル src/ memecli src/ memecli テスト テスト .DS_Store .DS_Store .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
makememe — コーディングエージェントが操作できるミーム CLI
エージェントと CI をコーディングするための、依存関係のない小さなミーム ジェネレーター。 APIなし
キー、サインアップなし、stdlib のみ — 無料のファイルをラップします
memegen.link API を使用して、公開画像 URL を返します。
![meme](url) を使用してどこにでも埋め込むことができます。
コーディング エージェントにミーム ボタンを与えます。 2 つのコマンドを実行したら、次のように尋ねます。
uv ツール install makememe # または: pipx install makememe / pip install makememe
meme --install-skill # バンドルされているクロード コード スキルをインストールします
Claude Code を再起動し、通常どおりに通信します。
あなた: ロッドがダウンしていることについて「これで大丈夫です」ミームを作成する
クロード: → https://api.memegen.link/images/fine/prod_is_down/this_is_fine.png
このスキルは、ジョークに適合するテンプレートを選択するようエージェントに教えます（ジョークだけでなく、
ドレイク）、キャプションを書き、言及した名前をすべて残しておき、URL を返します
PR、Slack、またはその他の場所でインラインレンダリングします。 (コーデックスと他のエージェントは機能します
meme --help / meme --list 経由ですぐに使用可能 — スキルは必要ありません。)
meme drake " 手動デプロイ " " ci/cd " # PNG を一時フォルダーに保存します
meme -t 問題ありません " prod がダウンしています " " これは問題ありません " --print-url # または共有可能な URL を出力するだけです
meme --list # すべてのテンプレート ID を参照します
試すのにインストールは必要ありません: uvx --from makememe meme drake "a" "b" 。
次に、GitHub アクションをコピー＆ペーストして、すべての PR に CI をミームします。
# インストール (1 つ選択)
uv ツールのインストール makememe # 推奨 — 分離され、`meme` を PATH に置きます
pipx インストール makememe
python3 -m pip install makememe
# インストールする

バンドルされた Claude Code スキルを使用してから、Claude Code を再起動します
meme --install-skill # ~/.claude/skills/meme/ (すべてのプロジェクト)
meme --install-skill --project # ./.claude/skills/meme/ (このリポジトリのみ)
アップグレード:
uv ツールのアップグレード makememe # pipx: pipx アップグレード makememe | pip: pip install -U makememe
meme --install-skill # ⚠️ これを再実行すると、スキルも更新されます
⚠️ パッケージをアップグレードしても、既にコピーされているスキルは更新されません
~/.claude/skills/meme/ 。アップグレードするたびに meme --install-skill を再実行します (その後、
Claude Code を再起動します)、またはエージェントは古いスキルを使用し続けます。
meme --version # CLI バージョン
ミーム --リスト | head # CLI が実行され、テンプレートに到達します
ls ~ /.claude/skills/meme/SKILL.md # スキルがインストールされています
使用法
meme < テンプレート > " 上の行 " " 下の行 " [-o out.png]
旗
意味
-t、--テンプレート
フラグとしてのテンプレート ID (位置指定の代替)。単一の権限承認ですべてのテンプレートをカバーできるように、エージェントはこれを使用する必要があります。
-o、--out
出力ファイル (デフォルト: 一時フォルダー内の固有のファイルなので、現在のディレクトリには書き込まれません)
--bg URL
テンプレートの代わりにカスタムの背景画像を使用する
--ext
png (デフォルト)、jpg、webp、または gif
--スタイル / --フォント
テンプレート スタイル バリアント/フォント オーバーライド
--open
完成した画像をデフォルトのビューアで開きます
--print-url
画像の URL を印刷し、ダウンロードしないでください
--json
機械可読出力 (スクリプト/エージェント用)
--リスト
利用可能なテンプレート ID をリストする
例
ミーム・ドレイク「手動デプロイ」「ci/cd」
ミーム同じ「売った後」「もし私が持っていたら」「同じ写真」
meme --bg https://example.com/pic.png " _ " " 回避 "
ミーム後悔 " SOLD @ 620 " " NOW 780 (+26%) " --print-url
ミーム --リスト
ミーム --list --json
統合
共有 (CI、チャット、コメント) の場合は、通常、URL ではなくパブリック URL が必要です。
ローカル ファイル — --print-url は、永続的な memegen.link URL を返します。

埋め込み
![meme](url) を含む任意の場所。ダウンロードや画像のホスティングは必要ありません。
これを .github/workflows/pr-meme.yml にドロップし、1 つのテスト行を次の行に変更します。
コマンドを実行すると、すべての PR が成功するかどうかに基づいて「これで大丈夫です」ミームが得られます。
テストに合格しました。これが設定全体です。秘密はありません (GITHUB_TOKEN が組み込まれています)。
名前 : prミーム
上: pull_request
権限:
pull-requests : # と書いてコメントを投稿します
ジョブ:
ミーム :
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
- 名前 : テストの実行
ID : テスト
run : echo "これをテストコマンド (pytest、npm test など) に置き換えます"
- 名前 : 結果をミームにコメントします
if : always() # テストが失敗した場合でも実行
実行: |
pip install --quite makememe
if [ "${{ ステップ.テスト.結果 }}" = "成功" ];それから
url=$(meme -t 成功 "テスト" "合格" --print-url)
それ以外の場合
url=$(meme -t 問題ありません "テストは失敗しました" "これは問題ありません" --print-url)
フィ
gh pr コメント "${{ github.event.pull_request.number }}" --body "![meme]($url)"
環境:
GH_TOKEN : ${{ Secrets.GITHUB_TOKEN }}
(examples/pr-meme.yml にも保存されます。)
Slack/Discord: --print-url URL をチャネルに投稿します。自動的に展開されます
プレビュー。
url= $( meme -t success " build " " pass " --print-url ) # -> パブリック URL、どこにでも埋め込みます
- で始まるテキスト
キャプション行が - で始まる場合 (例: "-26%" )、行の前に -- を置きます。
フラグとして解析されません。
ミーム後悔 --json -- " -26% " " なぜ "
( --json / -o のようなフラグを -- の前に置きます。)
エージェント向け（クロードコード/コーデックス/スクリプト）
このツールは解析できるように設計されています。
プレーン モードでは、出力パスのみが stdout に出力されます (ステータスは stderr になります)。
したがって、 path=$(meme drake "a" "b") は問題なく機能します。
--json モードは、単一の JSON オブジェクトを出力します。
{ "パス" : " /tmp/makememe/meme-ab12cd.png " 、 "バイト" : 12345 、 "url" : " https://api.memegen.link/... " }
--list --json はテンプレート cat を返します

JSON としてのログ。 --print-url --json
{"url": "..."} を返します。失敗すると、{"error": "...", "url": "..."} が返されます。
ゼロ以外の終了コード。
meme --list --json # テンプレート ID を検出する
meme drake " old way " " new way " --json # 生成、パスをキャプチャ
クロードコードのスキル
パッケージにはクロード コード スキルがバンドルされています。インストール後、1 つのコマンドを実行して、
クロード コードにツールを自動検出させます。
meme --install-skill # ~/.claude/skills/meme/ (すべてのプロジェクト) にインストールします
meme --install-skill --project # または ./.claude/skills/meme/ (このリポジトリのみ)
Claude Code を再起動して、話しかけるだけです。テンプレートやフラグに名前を付ける必要はありません。
スキルはジョークに合ったテンプレートを選択し、キャプションを書き、保持します。
あなたが挙げた名前は何でも。言えること:
会話型です — 「もっと面白い」、「代わりに 2 つのボタンのテンプレートを使用してください」、または
「収益を短くする」はすべてフォローアップとして機能します。
(Codex などの他のエージェントはこのスキル形式を使用しません。彼らはすべてを発見します
meme --help および meme --list を使用します。これらはすぐに使用できるようになります)。
CLI は、生のトレースバックを使用せずに正常に失敗するように構築されています。
ネットワーク エラー、デッド ホスト、不正なテンプレート ID (404)、特大テキスト (414)、
画像以外の背景 (415) はすべて 1 行のメッセージとともにゼロ以外で終了します。
部分的/ゴミファイルは書き込まれません。
Ctrl-C は正常に終了します (コード 130)。ヘッドなどに配管しても噴出しません。
壊れたパイプエラー 。
任意のテキスト — 絵文字、CJK、% # & / ? " \ 、タブ、制御文字、10k 文字
行 — は安全にエスケープされます。
テスト スイートを実行します (stdlib のみ、ネットワークは必要ありません)。
Python -m Unittest Discover -s テスト
仕組み
テンプレートとテキストから memegen.link URL を構築します (面倒な処理はすべて処理します)。
パスセグメントのエスケープ — スペース、_、-、? 、 / 、 % など)、
画像を選択して保存します。それがすべてのトリックです。
エージェント-金

memegen.link 経由の endly meme CLI — pip install makememe
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
9
v0.1.9
最新の
2026 年 6 月 8 日
+ 8 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agent-friendly meme CLI over memegen.link — pip install makememe - dhruvmehra/makememe

GitHub - dhruvmehra/makememe: Agent-friendly meme CLI over memegen.link — pip install makememe · GitHub
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
dhruvmehra
/
makememe
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits .github/ workflows .github/ workflows docs docs examples examples src/ memecli src/ memecli tests tests .DS_Store .DS_Store .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
makememe — a meme CLI your coding agent can drive
A tiny, zero-dependency meme generator for coding agents and CI . No API
key, no signup, stdlib-only — it wraps the free
memegen.link API and hands back a public image URL you
can embed anywhere with ![meme](url) .
Give your coding agent a meme button. Two commands, then just ask:
uv tool install makememe # or: pipx install makememe / pip install makememe
meme --install-skill # installs the bundled Claude Code skill
Restart Claude Code and talk to it normally:
you: make a "this is fine" meme about prod being down
Claude: → https://api.memegen.link/images/fine/prod_is_down/this_is_fine.png
The skill teaches the agent to pick the template that fits the joke (not just
drake), write the caption, and keep any names you mention — then hand back a URL
that renders inline in a PR, Slack, or anywhere. (Codex and other agents work
out of the box via meme --help / meme --list — no skill needed.)
meme drake " manual deploys " " ci/cd " # saves a PNG to a temp folder
meme -t fine " prod is down " " this is fine " --print-url # or just print a shareable URL
meme --list # browse all template ids
No install needed to try it: uvx --from makememe meme drake "a" "b" .
Next: meme your CI on every PR with a copy-paste GitHub Action.
# install (pick one)
uv tool install makememe # recommended — isolated, puts `meme` on PATH
pipx install makememe
python3 -m pip install makememe
# install the bundled Claude Code skill, then restart Claude Code
meme --install-skill # ~/.claude/skills/meme/ (all projects)
meme --install-skill --project # ./.claude/skills/meme/ (this repo only)
Upgrading:
uv tool upgrade makememe # pipx: pipx upgrade makememe | pip: pip install -U makememe
meme --install-skill # ⚠️ re-run this so the skill updates too
⚠️ Upgrading the package does not refresh the skill already copied into
~/.claude/skills/meme/ . Re-run meme --install-skill after every upgrade (then
restart Claude Code), or the agent keeps using the old skill.
meme --version # CLI version
meme --list | head # CLI runs and reaches templates
ls ~ /.claude/skills/meme/SKILL.md # skill is installed
Usage
meme < template > " top line " " bottom line " [-o out.png]
Flag
Meaning
-t, --template
template id as a flag (alternative to the positional). Agents should use this so a single permission approval covers every template
-o, --out
output file (default: a unique file in a temp folder, so it never writes into your current directory)
--bg URL
use a custom background image instead of a template
--ext
png (default), jpg , webp , or gif
--style / --font
template style variant / font override
--open
open the finished image in your default viewer
--print-url
print the image URL, don't download
--json
machine-readable output (for scripts/agents)
--list
list available template ids
Examples
meme drake " manual deploys " " ci/cd "
meme same " after I sold " " if I held " " same picture "
meme --bg https://example.com/pic.png " _ " " DODGED "
meme regret " SOLD @ 620 " " NOW 780 (+26%) " --print-url
meme --list
meme --list --json
Integrations
For sharing (CI, chat, comments) you usually want the public URL , not a
local file — --print-url returns a permanent memegen.link URL you can embed
anywhere with ![meme](url) . No download, no image hosting.
Drop this in .github/workflows/pr-meme.yml , change the one test line to your
command, and every PR gets a success / "this is fine" meme based on whether
tests passed. That's the whole setup — no secrets ( GITHUB_TOKEN is built in):
name : pr-meme
on : pull_request
permissions :
pull-requests : write # to post the comment
jobs :
meme :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- name : Run tests
id : tests
run : echo "replace this with your test command (e.g. pytest, npm test)"
- name : Comment a meme with the result
if : always() # run even when tests fail
run : |
pip install --quiet makememe
if [ "${{ steps.tests.outcome }}" = "success" ]; then
url=$(meme -t success "tests" "passed" --print-url)
else
url=$(meme -t fine "tests failed" "this is fine" --print-url)
fi
gh pr comment "${{ github.event.pull_request.number }}" --body "![meme]($url)"
env :
GH_TOKEN : ${{ secrets.GITHUB_TOKEN }}
(Also saved at examples/pr-meme.yml .)
Slack/Discord: post a --print-url URL to a channel; it auto-unfurls into a
preview.
url= $( meme -t success " build " " passed " --print-url ) # -> public URL, embed it anywhere
Text that starts with -
If a caption line begins with - (e.g. "-26%" ), put -- before your lines so
it isn't parsed as a flag:
meme regret --json -- " -26% " " WHY "
(Put flags like --json / -o before the -- .)
For agents (Claude Code / Codex / scripts)
The tool is designed to be parsed:
Plain mode prints only the output path to stdout (status goes to stderr),
so path=$(meme drake "a" "b") just works.
--json mode prints a single JSON object:
{ "path" : " /tmp/makememe/meme-ab12cd.png " , "bytes" : 12345 , "url" : " https://api.memegen.link/... " }
--list --json returns the template catalog as JSON; --print-url --json
returns {"url": "..."} ; failures return {"error": "...", "url": "..."}
with a non-zero exit code.
meme --list --json # discover template ids
meme drake " old way " " new way " --json # generate, capture the path
Claude Code skill
The package bundles a Claude Code skill. After installing, run one command to
make Claude Code auto-discover the tool:
meme --install-skill # installs into ~/.claude/skills/meme/ (all projects)
meme --install-skill --project # or into ./.claude/skills/meme/ (this repo only)
Restart Claude Code, then just talk to it. You don't name templates or flags —
the skill picks the template that fits the joke, writes the caption, and keeps
any names you mention. Things you can say:
It's conversational — "funnier" , "use the two-buttons template instead" , or
"now make the bottom line shorter" all work as follow-ups.
(Other agents like Codex don't use this skill format — they discover everything
through meme --help and meme --list , which already works out of the box.)
The CLI is built to fail gracefully, never with a raw traceback:
Network errors, dead hosts, bad template ids (404), oversized text (414), and
non-image backgrounds (415) all exit non-zero with a one-line message — and
no partial/garbage file is written .
Ctrl-C exits cleanly (code 130); piping into head etc. won't spew a
BrokenPipeError .
Arbitrary text — emoji, CJK, % # & / ? " \ , tabs, control chars, 10k-char
lines — is escaped safely.
Run the test suite (stdlib only, no network needed):
python -m unittest discover -s tests
How it works
It builds a memegen.link URL from your template + text (handling all the fiddly
path-segment escaping — spaces, _ , - , ? , / , % , etc.), downloads the
image, and saves it. That's the whole trick.
Agent-friendly meme CLI over memegen.link — pip install makememe
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
9
v0.1.9
Latest
Jun 8, 2026
+ 8 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
