---
source: "https://github.com/Run-Labs-com/agent-skills/tree/main"
hn_url: "https://news.ycombinator.com/item?id=49259589"
title: "Show HN: Add the official EU AI-content labels to images and videos"
article_title: "GitHub - Run-Labs-com/agent-skills: Agent skills for pi / Claude Code / Codex — including eu-ai-label, which burns the official EU AI-content labels (AI Act Art. 50) into images and videos · GitHub"
author: "mnewme"
captured_at: "2026-08-11T15:51:47Z"
capture_tool: "hn-digest"
hn_id: 49259589
score: 1
comments: 0
posted_at: "2026-08-11T15:08:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Add the official EU AI-content labels to images and videos

- HN: [49259589](https://news.ycombinator.com/item?id=49259589)
- Source: [github.com](https://github.com/Run-Labs-com/agent-skills/tree/main)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T15:08:02Z

## Translation

タイトル: HN を表示: 公式 EU AI コンテンツ ラベルを画像とビデオに追加します
記事のタイトル: GitHub - Run-Labs-com/agent-skills: pi / Claude Code / Codex のエージェント スキル — 公式 EU AI コンテンツ ラベル (AI 法第 50 条) を画像やビデオに焼き付ける eu-ai-label を含む · GitHub
説明: pi / Claude Code / Codex のエージェント スキル — 公式 EU AI コンテンツ ラベル (AI 法第 50 条) を画像やビデオに焼き付ける eu-ai-label を含む - Run-Labs-com/agent-skills

記事本文:
GitHub - Run-Labs-com/agent-skills: pi / Claude Code / Codex のエージェント スキル — 公式 EU AI コンテンツ ラベル (AI 法第 50 条) を画像やビデオに焼き付ける eu-ai-label を含む · GitHub
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
ランラボコム
/
エージェントのスキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット スキル スキー

lls .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
これは AI のスロップです。使用は自己責任で行ってください。私にとってはうまくいきます:)
私が実際に使用するエージェント スキルは、 エージェント スキル標準 に準拠しています。
pi 、Claude Code、Codex、および SKILL.md を読み取るハーネスで動作します。
まずコマンドライン ツールをインストールします。スキルは ffmpeg 、curl 、jq &co にシェルアウトされます。
そしてそれらなしでは機能しません。 「要件」を参照してください。
AI 生成コンテンツにラベルを付けるために EU の公式アイコンを焼きます (AI 法第 50 条の透明性)
ffmpeg 経由で画像やビデオに変換します。 12 個の公式アイコン ファイルすべてが同梱され、黒/白が自動選択されます。
コントラストを調整し、推奨される代替テキストを出力します。
$ npx skill add Run-Labs-com/agent-skills --skill eu-ai-label
名前の可用性
ブランド/製品名が使用可能かどうかを確認します: TLD 全体でのドメインの可用性
(RDAP + DNS + Whois であるため、.de などの ccTLD は無料であると誤って報告されず、委任されていません
.zzz のような TLD は利用可能であると誤って報告されていません）に加えて、商標パスがまとめられています。
判決表。
$ npx skill add Run-Labs-com/agent-skills --skill name-availability
商標調査
パブリック TMview を通じて EU および各国の商標 (EUIPO、DE/AT/CH、WIPO など) を検索します
JSON API。 Nice クラス、ステータスと所有者、および同一/ライブ/クラスを含むマークを返します。
競合判定 — ページネーションを認識し、生存性と偽の Office フィルターについてのフェールセーフを実現します。
$ npx skill add Run-Labs-com/agent-skills --skill 商標検索
3 つすべてを一度にインストールします。
$ npx skill add Run-Labs-com/agent-skills --skill '*'
要件
スキルをインストールするとファイルがコピーされるだけで、コマンドライン ツールは取り込まれません。
スクリプト呼び出し。これらを最初にインストールしないと、スキルは実行時に失敗します。
npx スキルの追加には、さらに Node.js ≥ 18 が必要です。 ffprobe には ffmpeg が同梱されています。
macOS エリア

dy は bash 、curl 、dig 、whois 、column および python3 を提供します
(コマンドラインツール経由) したがって、実際には ffmpeg と jq のみが必要です。
# macOS (自作)
醸造インストールffmpeg jq
# Debian / Ubuntu
sudo apt update && sudo apt install -y ffmpeg jqcurl python3 whois dnsutils bsdextrautils
# Fedora / RHEL
sudo dnf install -y ffmpeg jqカールpython3 whoisバインド-utils util-linux
# アーチ
sudo pacman -S --needed ffmpeg jqカールpythonwhoisバインドutil-linux
古い Debian/Ubuntu リリースでは、列は bsdextrautils ではなく bsdmainutils に存在します。
そして、新しいものについてはbind9-dnsutilsを掘り下げてください。どちらもオプションです。
ffmpeg ffprobe python3curl jq dig whois 列の t について;する
printf ' %-9s %s\n ' " $t " " $( command -v " $t " || echo ' MISSING ' ) "
完了しました
jq -n ' IN(1) ' > /dev/null 2>&1 && echo ' jq >= 1.6 ok ' || echo ' jq が古すぎます (1.6 以上が必要) '
2 つのシェル スキルにはネットワーク アクセス (RDAP、DNS、whois、TMview API) が必要です。オフラインの彼らは
推測ではなく不明を報告します。彼らは使用可能な名前を決して報告しません。
上記の npx skill add コマンドはすでに pi ( -a pi ) をサポートしており、次の場所にインストールされます。
~/.pi/エージェント/スキル/ 。代わりにクローンから接続するには:
git clone https://github.com/Run-Labs-com/agent-skills ~ /src/agent-skills
設定 — ~/.pi/agent/settings.json 内のリポジトリ全体を指す pi
(スキルドキュメント):
{ "スキル" : [ " ~/src/エージェント-スキル/スキル " ] }
シンボリックリンク — 個々のスキルをグローバルにロードします。
eu-ai-label name-availability traderade-search の s については、する
ln -s ~ /src/agent-skills/skills/ $s ~ /.pi/agent/skills/ $s
完了しました
1 回限り — pi --skill ~/src/agent-skills/skills/eu-ai-label
クロード コード: ~/.claude/skills/ へのシンボリックリンク · コーデックス: ~/.codex/skills/ へのシンボリックリンク ·
~/.agents/skills : symlink を読み取るハーネスは代わりにそこにあります。
すべてのスキルは、実行できるスクリプトを含むプレーンなディレクトリです。

自分自身:
git clone https://github.com/Run-Labs-com/agent-skills && cd エージェント-スキル
python3 skill/eu-ai-label/scripts/ai-label.py Clip.mp4 -v generated
Skills/name-availability/scripts/check-domains.sh acme com ai de
skill/trademark-search/scripts/tmview-search.sh " acme " -o DACH
ライセンス
コードについては MIT (ライセンスを参照)。 Skills/eu-ai-label/assets/ の EU アイコンは次のとおりです。
欧州委員会によって出典なしで自由に使用できるように公開されています - を参照してください。
アセット/SOURCE.md 。
pi / Claude Code / Codex のエージェント スキル — 公式 EU AI コンテンツ ラベル (AI 法第 50 条) を画像やビデオに焼き付ける eu-ai-label を含む
Readme ライセンス アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agent skills for pi / Claude Code / Codex — including eu-ai-label, which burns the official EU AI-content labels (AI Act Art. 50) into images and videos - Run-Labs-com/agent-skills

GitHub - Run-Labs-com/agent-skills: Agent skills for pi / Claude Code / Codex — including eu-ai-label, which burns the official EU AI-content labels (AI Act Art. 50) into images and videos · GitHub
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
Run-Labs-com
/
agent-skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits skills skills .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
This is AI slop, use at your own risk. Works for me though :)
Agent skills I actually use, following the Agent Skills standard .
Works with pi , Claude Code, Codex, and any harness that reads SKILL.md .
Install the command-line tools first — the skills shell out to ffmpeg , curl , jq &co.
and will not work without them. See Requirements .
Burns the official EU icons for labelling AI-generated content (AI Act Art. 50 transparency)
into images and videos via ffmpeg. Ships all 12 official icon files, auto-picks black/white for
contrast, and prints the recommended alt text.
$ npx skills add Run-Labs-com/agent-skills --skill eu-ai-label
name-availability
Checks whether a brand/product name is usable: domain availability across TLDs
(RDAP + DNS + whois, so ccTLDs like .de aren't falsely reported free, and non-delegated
TLDs like .zzz aren't falsely reported available) plus a trademark pass, summarised in a
verdict table.
$ npx skills add Run-Labs-com/agent-skills --skill name-availability
trademark-search
Searches EU & national trademarks (EUIPO, DE/AT/CH, WIPO, …) through the public TMview
JSON API. Returns marks with Nice classes, status and owner, and an identical/live/class
conflict verdict — pagination-aware and fail-safe about liveness and bogus office filters.
$ npx skills add Run-Labs-com/agent-skills --skill trademark-search
Install all three at once:
$ npx skills add Run-Labs-com/agent-skills --skill '*'
Requirements
Installing a skill only copies files — it does not pull in the command-line tools the
scripts call. Install those first, or the skills fail at runtime.
npx skills add additionally needs Node.js ≥ 18 . ffprobe ships with ffmpeg .
macOS already provides bash , curl , dig , whois , column and python3
(via the Command Line Tools), so in practice you only need ffmpeg and jq .
# macOS (Homebrew)
brew install ffmpeg jq
# Debian / Ubuntu
sudo apt update && sudo apt install -y ffmpeg jq curl python3 whois dnsutils bsdextrautils
# Fedora / RHEL
sudo dnf install -y ffmpeg jq curl python3 whois bind-utils util-linux
# Arch
sudo pacman -S --needed ffmpeg jq curl python whois bind util-linux
On older Debian/Ubuntu releases column lives in bsdmainutils instead of bsdextrautils ,
and dig in bind9-dnsutils on newer ones. Both are optional.
for t in ffmpeg ffprobe python3 curl jq dig whois column ; do
printf ' %-9s %s\n ' " $t " " $( command -v " $t " || echo ' MISSING ' ) "
done
jq -n ' IN(1) ' > /dev/null 2>&1 && echo ' jq >= 1.6 ok ' || echo ' jq too old (need >= 1.6) '
The two shell skills need network access (RDAP, DNS, whois, the TMview API). Offline they
report UNKNOWN rather than guessing — they never report a taken name as available.
The npx skills add command above already supports pi ( -a pi ) and installs to
~/.pi/agent/skills/ . To wire it up from a clone instead:
git clone https://github.com/Run-Labs-com/agent-skills ~ /src/agent-skills
Settings — point pi at the whole repo in ~/.pi/agent/settings.json
( skills docs ):
{ "skills" : [ " ~/src/agent-skills/skills " ] }
Symlinks — load individual skills globally:
for s in eu-ai-label name-availability trademark-search ; do
ln -s ~ /src/agent-skills/skills/ $s ~ /.pi/agent/skills/ $s
done
One-off — pi --skill ~/src/agent-skills/skills/eu-ai-label
Claude Code: symlink into ~/.claude/skills/ · Codex: symlink into ~/.codex/skills/ ·
any harness that reads ~/.agents/skills : symlink there instead.
Every skill is a plain directory with scripts you can run yourself:
git clone https://github.com/Run-Labs-com/agent-skills && cd agent-skills
python3 skills/eu-ai-label/scripts/ai-label.py clip.mp4 -v generated
skills/name-availability/scripts/check-domains.sh acme com ai de
skills/trademark-search/scripts/tmview-search.sh " acme " -o DACH
Licence
MIT for the code (see LICENSE ). The EU icons in skills/eu-ai-label/assets/ are
published by the European Commission for free use without attribution — see
assets/SOURCE.md .
Agent skills for pi / Claude Code / Codex — including eu-ai-label, which burns the official EU AI-content labels (AI Act Art. 50) into images and videos
Readme License Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
