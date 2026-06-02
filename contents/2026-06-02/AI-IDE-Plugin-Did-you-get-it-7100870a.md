---
source: "https://github.com/bugthesystem/dygit"
hn_url: "https://news.ycombinator.com/item?id=48357700"
title: "AI IDE Plugin: Did you get it?"
article_title: "GitHub - bugthesystem/dygit: Did you get it? — An AI IDE plugin that quietly fixes messy prompts — typos, dropped words, thumb-typed, half-awake — so Claude gets it the first time · GitHub"
author: "ziyasal"
captured_at: "2026-06-02T00:38:22Z"
capture_tool: "hn-digest"
hn_id: 48357700
score: 2
comments: 1
posted_at: "2026-06-01T14:59:37Z"
tags:
  - hacker-news
  - translated
---

# AI IDE Plugin: Did you get it?

- HN: [48357700](https://news.ycombinator.com/item?id=48357700)
- Source: [github.com](https://github.com/bugthesystem/dygit)
- Score: 2
- Comments: 1
- Posted: 2026-06-01T14:59:37Z

## Translation

タイトル: AI IDE プラグイン: 理解しましたか?
記事のタイトル: GitHub - bugthesystem/dygit: 理解できましたか? — 乱雑なプロンプト (タイプミス、単語の脱落、親指での入力、中途半端な入力) を静かに修正する AI IDE プラグイン。クロードが初めて理解できるように · GitHub
説明: 分かりましたか? — 乱雑なプロンプト (タイプミス、単語の脱落、親指での入力、半分起きている状態) を静かに修正する AI IDE プラグイン - クロードが初めて理解できるようにする - bugthesystem/dygit

記事本文:
GitHub - bugthesystem/dygit: わかりましたか? — 乱雑なプロンプト (タイプミス、単語の脱落、親指での入力、中途半端な入力) を静かに修正する AI IDE プラグイン。クロードが初めて理解できるように · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
バグザシスト

えっと
/
ダイギット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット .claude-plugin .claude-plugin .cursor .cursor .github/ workflows .github/ workflows Formula Formula アセット アセット bin bin コマンド コマンド クレート クレート フック フック オープンコード オープンコード スクリプト スクリプト .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルナビゲーション
AI エディターがプロンプトのタイプミスを認識する前に、プロンプトのタイプミスを修正します。ローカルでは AI もトークンコストもありません。
クロード コード、カーソル、OpenCode で動作します。
aut hbug wehn usr lgs ut を修正してください
モデルには次のように書かれています。
ユーザーがログアウトしたときの認証バグを修正
オリジナルのテキストはトランスクリプトに残ります。どこにも何も送信されません。
小さな Rust バイナリが各プロンプトで実行され、2 つのパスで修正されます。
明確なキーボード スリップの厳選された表 — teh → the 、wehn → when 、
usr → user — オフラインで即座に修正されます。
それ以外のすべてについては、82,000 語の頻度辞書と比較して symspell を実行します。
実際の単語は辞書に含まれているため、そのまま残されます: form、route、および
安定したものは、他のものに「修正」されることはありません。
スペース修復パスも分割トークンを再結合します ( aut hbug → auth bug )。
両方の半分が実際の単語の場合。あいまいな入力はそのまま残されます - バイナリ
推測はしないだろう。
クロード コードとカーソルでは、クリーン化された読み取り値が次のようにモデルに渡されます。
サイドチャネル コンテキストなので、元のプロンプトは表示されたままになります。 OpenCode (
そのようなチャネルはありません) メッセージはインラインで書き換えられ、信頼性の高いメッセージに対してのみ書き換えられます。
修正します。
82,000 語の辞書のロードには約 1/2 秒かかるため、常駐デーモンがそれを 1 回ロードし、
Unix ソケット経由で 1 ミリ秒以内に応答します。暖かくなるまで、テーブル
あなたをカバーします。どれでも

エラー バイナリはサイレントのままで、プロンプトは通過します
手付かずの。
スペルチェックに AI はありません。これはルックアップ テーブルと編集距離アルゴリズムです。
モデルがタイプミスを修正するように求められることはありません。
トークンコストはかかりません。補正はローカル計算です。乱雑なプロンプトにより短いメッセージが追加される
すでに送信していたリクエストへのヒント。クリーンなプロンプトでは何も追加されません。
プライベート。あなたのマシン上で動作します。ネットワーク、テレメトリ、API キーはありません。
1 つのバイナリ、3 つのエディタ。すべての統合の背後にある同じエンジン。
バイナリを入手します (macOS および Linux、arm64 + x64):
醸造タップ bugthesystem/dygit https://github.com/bugthesystem/dygit
醸造インストールdygi
または、 ソース からビルドします。次に、それをエディターに接続します。
クロード プラグイン マーケットプレイス add bugthesystem/dygit
クロード プラグインのインストール Did-you-get-it@dygit-local
プロンプトフックと 4 つのスラッシュコマンド (以下) をインストールします。
Cursor 1.7 以降には、同等のプロンプト フックがあります。ポイント .cursor/hooks.json (プロジェクト) または
ラッパーの ~/.cursor/hooks.json (グローバル):
{
「バージョン」: 1 、
「フック」: {
"beforeSubmitPrompt" : [
{ "コマンド" : " bash /absolute/path/to/dygit/hooks/run.sh " }
]
}
}
すぐにコピーできる .cursor/hooks.json がこのリポジトリに同梱されています。
cp opencode/dygi.js ~ /.config/opencode/plugins/ # またはプロジェクトごとに .opencode/plugins/
OpenCode はメッセージをインラインで書き換えるので (サイドチャネルなし)、
信頼性の高い修正。 opencode/README.md を参照してください。
編集者
ステータス
仕組み
クロード・コード
いっぱい
UserPromptSubmit フック →AdditionalContext
カーソル
いっぱい
beforeSubmitPrompt フック →AdditionalContext
オープンコード
いっぱい
chat.message プラグイン → インライン書き換え (高信頼性のみ)
VSコード
まだ ( #1 )
Chat API には事前送信フックがありません。参加者は明示的な @ment でのみ起動します
ウィンドサーフィン
まだ ( #2 )
pre_user_prompt フックはブロックまたはログを記録できますが、プロンプトを変更することはできません
バイナリはエディターに依存しません - VS Code または Windsur の場合

f 事前送信フックを追加すると、
コンテキストを挿入したり、プロンプトを書き換えたりするのは、小さなシムです。どちらも
追跡されており、入手可能です — #1 と #2 を参照してください。
状態は ~/.claude/plugins/data/did-you-get-it/ にあります。
このバイナリはスタンドアロンでも動作します。
echo "バグを修正してください" |ディギ正しい
# {"original":"バグを修正","cleaned":"バグを修正","verdict":"trivial","changed":true}
ビルドする
./scripts/build-all.sh # すべてのプラットフォーム (クロスツールチェーンが必要)
バイナリは bin/ に配置されます。フックは uname から正しいものを選択します。 CI がすべてを構築する
各タグに 4 つのプラットフォーム。
macOS はバイナリのコピー時にそのアドホック署名を削除し、その後カーネルがバイナリのアドホック署名を削除します。
起動時にそれを殺します。 codesign --force --sign - <binary> で再署名します。
build-all.sh は、darwin ターゲットに対してこれを実行します。
分かりましたか？ — 乱雑なプロンプト（タイプミス、単語の脱落、親指での入力、半分起きている状態）を静かに修正する AI IDE プラグインなので、クロードは初めてそれを理解できます
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
v0.2.0
最新の
2026 年 6 月 1 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Did you get it? — An AI IDE plugin that quietly fixes messy prompts — typos, dropped words, thumb-typed, half-awake — so Claude gets it the first time - bugthesystem/dygit

GitHub - bugthesystem/dygit: Did you get it? — An AI IDE plugin that quietly fixes messy prompts — typos, dropped words, thumb-typed, half-awake — so Claude gets it the first time · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
bugthesystem
/
dygit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits .claude-plugin .claude-plugin .cursor .cursor .github/ workflows .github/ workflows Formula Formula assets assets bin bin commands commands crate crate hooks hooks opencode opencode scripts scripts .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Fixes typos in your prompts before your AI editor sees them. Locally, no AI, no token cost.
Works in Claude Code, Cursor, and OpenCode.
fix teh aut hbug wehn usr lgs ut
The model reads:
fix the auth bug when user logs out
Your original text stays in the transcript. Nothing is sent anywhere.
A small Rust binary runs on each prompt and corrects it in two passes:
A curated table of unambiguous keyboard slips — teh → the , wehn → when ,
usr → user — fixed instantly, offline.
symspell against an 82k-word frequency dictionary for everything else.
Real words are in the dictionary, so they're left alone: form , route , and
stable are never "corrected" into something else.
A space-repair pass also re-joins split tokens ( aut hbug → auth bug ), but only
when both halves are real words. Ambiguous input is left untouched — the binary
won't guess.
In Claude Code and Cursor the cleaned reading is passed to the model as
side-channel context, so your original prompt stays visible. In OpenCode (which has
no such channel) the message is rewritten inline, and only for high-confidence
fixes.
The 82k-word dictionary takes ~½s to load, so a resident daemon loads it once and
answers over a Unix socket in under a millisecond. Until it's warm, the table
covers you. On any error the binary stays silent and your prompt goes through
untouched.
No AI in the spell-check. It's a lookup table and an edit-distance algorithm.
The model is never asked to fix a typo.
No token cost. Correction is local computation. A messy prompt adds a short
hint to the request you were already sending; a clean prompt adds nothing.
Private. Runs on your machine. No network, no telemetry, no API key.
One binary, three editors. Same engine behind every integration.
Get the binary (macOS & Linux, arm64 + x64):
brew tap bugthesystem/dygit https://github.com/bugthesystem/dygit
brew install dygi
Or build from source . Then wire it into your editor.
claude plugin marketplace add bugthesystem/dygit
claude plugin install did-you-get-it@dygit-local
Installs the prompt hook and four slash commands ( below ).
Cursor 1.7+ has an equivalent prompt hook. Point .cursor/hooks.json (project) or
~/.cursor/hooks.json (global) at the wrapper:
{
"version" : 1 ,
"hooks" : {
"beforeSubmitPrompt" : [
{ "command" : " bash /absolute/path/to/dygit/hooks/run.sh " }
]
}
}
A ready-to-copy .cursor/hooks.json ships in this repo.
cp opencode/dygi.js ~ /.config/opencode/plugins/ # or .opencode/plugins/ per project
OpenCode rewrites the message inline (no side-channel), so it only acts on
high-confidence fixes. See opencode/README.md .
Editor
Status
Mechanism
Claude Code
full
UserPromptSubmit hook → additionalContext
Cursor
full
beforeSubmitPrompt hook → additionalContext
OpenCode
full
chat.message plugin → inline rewrite (high-confidence only)
VS Code
not yet ( #1 )
the Chat API has no pre-send hook; participants only fire on explicit @mention
Windsurf
not yet ( #2 )
the pre_user_prompt hook can block or log, but cannot modify the prompt
The binary is editor-agnostic — if VS Code or Windsurf add a pre-send hook that can
inject context or rewrite the prompt, wiring them up is a small shim. Both are
tracked and up for grabs — see #1 and #2 .
State lives in ~/.claude/plugins/data/did-you-get-it/ .
The binary also works standalone:
echo " fix teh bug " | dygi correct
# {"original":"fix teh bug","cleaned":"fix the bug","verdict":"trivial","changed":true}
Build
./scripts/build-all.sh # all platforms (needs cross toolchains)
Binaries land in bin/ ; the hook picks the right one from uname . CI builds all
four platforms on every tag.
macOS strips a binary's adhoc signature when it's copied, and the kernel then
kills it on launch. Re-sign with codesign --force --sign - <binary> ;
build-all.sh does this for the darwin targets.
Did you get it? — An AI IDE plugin that quietly fixes messy prompts — typos, dropped words, thumb-typed, half-awake — so Claude gets it the first time
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
v0.2.0
Latest
Jun 1, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
