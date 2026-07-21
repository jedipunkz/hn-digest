---
source: "https://github.com/overflowy/herdr-claude-gpt-adversarial-review-skill"
hn_url: "https://news.ycombinator.com/item?id=48993960"
title: "Show HN: Adversarial code review setup with herdr, Claude and GPT-5.6-sol"
article_title: "GitHub - overflowy/herdr-claude-gpt-adversarial-review-skill: Claude Code skill for adversarial code review - spawns a GPT‑5.6 Sol reviewer in a herdr split pane to attack your diff, then interrogates its findings before relaying verdicts. · GitHub"
author: "overflowy"
captured_at: "2026-07-21T16:16:34Z"
capture_tool: "hn-digest"
hn_id: 48993960
score: 2
comments: 1
posted_at: "2026-07-21T15:51:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Adversarial code review setup with herdr, Claude and GPT-5.6-sol

- HN: [48993960](https://news.ycombinator.com/item?id=48993960)
- Source: [github.com](https://github.com/overflowy/herdr-claude-gpt-adversarial-review-skill)
- Score: 2
- Comments: 1
- Posted: 2026-07-21T15:51:35Z

## Translation

タイトル: HN を表示: herdr、Claude、GPT-5.6-sol を使用した敵対的コード レビューのセットアップ
記事のタイトル: GitHub - overflowy/herdr-claude-gpt-adversarial-review-skill: 敵対的コード レビュー用のクロード コード スキル - Herdr 分割ペインに GPT‑5.6 Sol レビュアーを生成して差分を攻撃し、評決を伝える前にその結果を尋問します。 · GitHub
説明: 敵対的コードレビュー用の Claude Code スキル - GPT‑5.6 Sol レビュアーを Herdr 分割ペインに生成して差分を攻撃し、評決を伝える前にその結果を尋問します。 - overflowy/herdr-claude-gpt-adversarial-review-skill

記事本文:
GitHub - overflowy/herdr-claude-gpt-adversarial-review-skill: 敵対的コード レビュー用のクロード コード スキル - Herdr 分割ペインに GPT‑5.6 Sol レビュアーを生成して差分を攻撃し、評決を伝える前にその結果を尋問します。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。 rにリロード

セッションを更新してください。
アラートを閉じる
{{ メッセージ }}
あふれた
/
herdr-claude-gpt-adversarial-review-skill
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
overflowy/herdr-claude-gpt-adversarial-review-skill
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .claude-plugin .claude-plugin skill/ adversarial-review skill/ adversarial-review .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
2 番目の独立したモデル GPT‑5.6 Sol をレビューアとして使用して、敵対的コード レビューを実行するクロード コード スキル。 Claude は、ライブ Herdr 分割ペインでレビューアーを生成し、差分 (または計画) と指定されたインテントを渡し、完了するのを待ち、その後、評決を伝える前に実際のコードに対してすべての検出結果を尋問します。
レビューアーは独自のペインでインタラクティブに実行されるため、レビューアーの動作を観察したり、飛びついて操作したりすることができます。
herdr、CLIProxyAPI、agent-safehouse、safecodex シェル関数の 4 つを順番に設定する必要があります。次にスキル自体をインストールします。
敵対的レビューデモ.mp4
1.herdrをインストールする
herdr はスキルが駆動するエージェント マルチプレクサです。レビューア ペインを作成し、入力を送信し、アイドル/動作/ブロック状態を待機します。
カール -fsSL https://herdr.dev/install.sh |しー
herdr 内でターミナル セッションを実行します (クイック スタートを参照)。このスキルは、実行されるクロード コード セッションが herdr ペインに存在することを前提としているため、その隣にある兄弟ペインを分割できます。
2. CLIProxyAPI をインストールして構成する
CLIProxyAPI は、他のプロバイダーのログイン (ここでは、GPT-5.6 Sol を提供する OpenAI/Codex ログイン) によってサポートされる Anthropic 互換 API を公開するローカル プロキシです。これにより、ストック クロード CLI が非人類の MO と通信できるようになります。

デル。
インストール - リリース ページからバイナリをダウンロードします (または Docker を使用します。ガイドを参照してください)。
~/.config/cli-proxy-api/config.yml に構成を作成します。
ポート : 8317
認証ディレクトリ: " ~/.cli-proxy-api "
APIキー:
- " <任意のランダム文字列> " # 例:出力: openssl rand -hex 24
API キーは、プロキシとクロード CLI 間の単なる共有秘密です。作成してください。どこかに保存してください。後で、safecodex 関数で必要になります。
OpenAI アカウントでログインします (Codex OAuth フロー):
cli-proxy-api --config ~ /.config/cli-proxy-api/config.yml --codex-login
サーバーを起動し、実行したままにします。
cli-proxy-api --config ~ /.config/cli-proxy-api/config.yml
スキルは http://localhost:8317 上でそれを期待します。
Agent-safehouse は、macOS ネイティブのカーネルレベルのサンドボックスを提供します。レビューアーは --dangerously-skip-permissions で実行されるため、サンドボックスによってレビューアーはプロジェクト ディレクトリに限定されます。SSH キー、他のリポジトリ、またはワークスペースの外部のものには触れることができません。
brew install eugene1g/safehouse/agent-safehouse
次に、安全なラッパーを ~/.zshrc に追加します (シェルに合わせて調整します)。
安全 () {
セーフハウス --env --add-dirs= " /tmp " " $@ "
}
/tmp 付与はこのスキルにとって重要です。これはハンドオフ ファイルが存在する場所です (スキルはレビュー プロンプトを /tmp/adversarial-review/ に書き込み、レビュー担当者はその結果をそこに書き込みます)。これがないと、レビュー担当者はプロンプトを読んだり、出力を生成したりできません。
4.セーフコーデックス関数を定義する
~/.zshrc に追加します (シェルに合わせて調整します):
import CLI_PROXY_API_KEY= " <config.yml のキー> "
セーフコーデックス () {
ANTHROPIC_BASE_URL=http://localhost:8317 \
ANTHROPIC_AUTH_TOKEN= " $CLI_PROXY_API_KEY " \
ANTHROPIC_MODEL=gpt-5.6-sol \
ANTHROPIC_SMALL_FAST_MODEL=gpt-5.6-sol \
CLAUDE_CODE_SUBAGENT_MODEL=gpt-5.6-sol \
CLAUDE_CODE_ALWAYS_ENABLE_EFFORT=1 \
CLAUDE_CODE_MAX_TOOL_USE_CONCURRENCY=3 \

ENABLE_TOOL_SEARCH=false \
安全なクロード --model ' gpt-5.6-sol[1m] ' --dangerously-skip-permissions " $@ "
}
シェルをリロードし( exec zsh )、対話型シェル構成で関数を定義する必要があることに注意してください。スキルは herdr pane run 経由で関数を起動します。これにより対話型シェルが正確に実行されるため、エイリアスと関数が解決されます。
このリポジトリは Claude Code プラグイン マーケットプレイスであるため、Claude Code 内から直接インストールできます。
/plugin Marketplace add overflowy/herdr-claude-gpt-adversarial-review-skill
/plugin install adversarial-review@herdr-claude-gpt-adversarial-review-skill
あるいは、スキルをスキル ディレクトリに手動でコピーします。
mkdir -p ~ /.claude/skills/adversarial-review
cp スキル/adversarial-review/SKILL.md ~ /.claude/skills/adversarial-review/
使用法
レビューしたいプロジェクトの herdr ペイン内で実行されている Claude Code セッションから:
/敵対的レビュー
または、自然言語で尋ねるだけです。「この差分に穴を開けて」、「この計画をレッドチームに」、「出荷する前にセカンドオピニオンを求めて」などです。クロードは次のようにします:
作業の意図 (変更によって何を達成しようとしているのか) を述べます。
diff + レビュー担当者の料金を /tmp/adversarial-review/ の下のプロンプト ファイルに書き込みます。
herdr ペインを分割し、その中でsafecodexを起動します。レビュアーのTUIはそこに存在します。
レビューが完了するまで待ちます (herdr はエージェントのステータスをネイティブに追跡します)。
実際のコードと照らし合わせてすべての発見事項を調査し、議論のあるものについてはまだ開いているレビューセッションに参加させ、それぞれを確認済み、議論済み、または未検証として報告します。
最後のステップが重要です。査読者は指示によって敵対的であり、時には問題を誇張したり捏造したりすることがあります。得られるのは生の出力ではなく評決であり、「レビュー担当者が見つけたものは何も支持されない」という結果が有効です。
初めて使用する前に、各レイヤーを確認してください。
# プロキシが起動し、キーを受け入れます

curl -s http://localhost:8317/v1/models -H " x-api-key: $CLI_PROXY_API_KEY "
#safecodex はエンドツーエンドで動作します (ワンショット、TUI なし)
safecodex -p " OK と言う "
トラブルシューティング
レビューアー ペインは開きますが、すぐにエラーが発生します。通常はプロキシです。8317 で実行されていること、および CLI_PROXY_API_KEY が config.yml の api-keys のエントリと一致することを確認してください。
モデルからの認証エラー - Codex OAuth トークンの有効期限が切れている可能性があります。 --codex-login ステップを再実行します。
safecodex: コマンドがペインに見つかりません - 関数が対話型シェル構成にないか、1 つのターミナルでのみエクスポートされました。 ~/.zshrc に存在する必要があります。
サンドボックスの拒否 - エージェント セーフハウスは、プロジェクト ディレクトリへの読み取り/書き込みのみを許可します。レビューに別のパスが正当に必要な場合は、そのパスをセーフ ラッパーの --add-dirs リストに追加します。
敵対的コード レビュー用の Claude Code スキル - GPT‑5.6 Sol レビュアーを Herdr 分割ペインに生成して差分を攻撃し、評決を伝える前にその結果を尋問します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude Code skill for adversarial code review - spawns a GPT‑5.6 Sol reviewer in a herdr split pane to attack your diff, then interrogates its findings before relaying verdicts. - overflowy/herdr-claude-gpt-adversarial-review-skill

GitHub - overflowy/herdr-claude-gpt-adversarial-review-skill: Claude Code skill for adversarial code review - spawns a GPT‑5.6 Sol reviewer in a herdr split pane to attack your diff, then interrogates its findings before relaying verdicts. · GitHub
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
overflowy
/
herdr-claude-gpt-adversarial-review-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
overflowy/herdr-claude-gpt-adversarial-review-skill
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .claude-plugin .claude-plugin skills/ adversarial-review skills/ adversarial-review .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
A Claude Code skill that runs an adversarial code review using a second, independent model - GPT‑5.6 Sol - as the reviewer. Claude spawns the reviewer in a live herdr split pane, hands it your diff (or plan) plus the stated intent, waits for it to finish, then interrogates every finding against the actual code before relaying a verdict.
The reviewer runs interactively in its own pane, so you can watch it work or jump in and steer.
Four things need to be set up, in order: herdr, CLIProxyAPI, agent-safehouse, and the safecodex shell function. Then install the skill itself.
adversarial-review-demo.mp4
1. Install herdr
herdr is the agent multiplexer the skill drives - it creates the reviewer pane, sends it input, and waits on its idle/working/blocked status.
curl -fsSL https://herdr.dev/install.sh | sh
Run your terminal sessions inside herdr (see the quick start ). The skill assumes the Claude Code session it runs in lives in a herdr pane, so it can split a sibling pane next to it.
2. Install and configure CLIProxyAPI
CLIProxyAPI is a local proxy that exposes an Anthropic-compatible API backed by other providers' logins - here, an OpenAI/Codex login serving GPT‑5.6 Sol. This is what lets the stock claude CLI talk to a non-Anthropic model.
Install - download a binary from the releases page (or use Docker; see the guides ).
Create the config at ~/.config/cli-proxy-api/config.yml :
port : 8317
auth-dir : " ~/.cli-proxy-api "
api-keys :
- " <any-random-string> " # e.g. output of: openssl rand -hex 24
The API key is just a shared secret between the proxy and the claude CLI - make one up. Save it somewhere; the safecodex function needs it later.
Log in with your OpenAI account (Codex OAuth flow):
cli-proxy-api --config ~ /.config/cli-proxy-api/config.yml --codex-login
Start the server and leave it running:
cli-proxy-api --config ~ /.config/cli-proxy-api/config.yml
The skill expects it on http://localhost:8317 .
agent-safehouse provides macOS-native kernel-level sandboxing. The reviewer runs with --dangerously-skip-permissions , so the sandbox is what keeps it confined to the project directory - it can't touch your SSH keys, other repos, or anything outside its workspace.
brew install eugene1g/safehouse/agent-safehouse
Then add the safe wrapper to ~/.zshrc (adjust for your shell):
safe () {
safehouse --env --add-dirs= " /tmp " " $@ "
}
The /tmp grant matters for this skill: it's where the handoff files live (the skill writes the review prompt to /tmp/adversarial-review/ and the reviewer writes its findings back there). Without it, the reviewer can't read its prompt or produce output.
4. Define the safecodex function
Add to ~/.zshrc (adjust for your shell):
export CLI_PROXY_API_KEY= " <the key from your config.yml> "
safecodex () {
ANTHROPIC_BASE_URL=http://localhost:8317 \
ANTHROPIC_AUTH_TOKEN= " $CLI_PROXY_API_KEY " \
ANTHROPIC_MODEL=gpt-5.6-sol \
ANTHROPIC_SMALL_FAST_MODEL=gpt-5.6-sol \
CLAUDE_CODE_SUBAGENT_MODEL=gpt-5.6-sol \
CLAUDE_CODE_ALWAYS_ENABLE_EFFORT=1 \
CLAUDE_CODE_MAX_TOOL_USE_CONCURRENCY=3 \
ENABLE_TOOL_SEARCH=false \
safe claude --model ' gpt-5.6-sol[1m] ' --dangerously-skip-permissions " $@ "
}
Reload your shell ( exec zsh ) and note that the function must be defined in your interactive shell config - the skill launches it via herdr pane run , which runs an interactive shell precisely so aliases and functions resolve.
The repo is a Claude Code plugin marketplace, so you can install it directly from within Claude Code:
/plugin marketplace add overflowy/herdr-claude-gpt-adversarial-review-skill
/plugin install adversarial-review@herdr-claude-gpt-adversarial-review-skill
Alternatively, copy the skill into your skills directory manually:
mkdir -p ~ /.claude/skills/adversarial-review
cp skills/adversarial-review/SKILL.md ~ /.claude/skills/adversarial-review/
Usage
From a Claude Code session running inside a herdr pane, in the project you want reviewed:
/adversarial-review
or just ask in natural language - "poke holes in this diff", "red-team this plan", "get a second opinion before I ship". Claude will:
State the intent of the work (what the change is trying to achieve).
Write the diff + reviewer charge to a prompt file under /tmp/adversarial-review/ .
Split a herdr pane and launch safecodex in it - the reviewer's TUI is live there.
Wait for the review to finish (herdr tracks the agent's status natively).
Interrogate every finding against the actual code, pressing the still-open reviewer session on anything disputed, and report each one as Confirmed, Disputed, or Unverified.
The last step matters: the reviewer is adversarial by instruction and will sometimes overstate or manufacture problems. You get verdicts, not raw output - and "nothing the reviewer found holds up" is a valid outcome.
Before first use, verify each layer:
# Proxy is up and accepting your key
curl -s http://localhost:8317/v1/models -H " x-api-key: $CLI_PROXY_API_KEY "
# safecodex works end to end (one-shot, no TUI)
safecodex -p " Say ok "
Troubleshooting
Reviewer pane opens but errors immediately - usually the proxy: check it's running on 8317 and that CLI_PROXY_API_KEY matches an entry in config.yml 's api-keys .
Auth errors from the model - the Codex OAuth token may have expired; re-run the --codex-login step.
safecodex: command not found in the pane - the function isn't in your interactive shell config, or was only exported in one terminal. It must live in ~/.zshrc .
Sandbox denials - agent-safehouse grants read/write to the project directory only. If the review legitimately needs another path, add it to the safe wrapper's --add-dirs list.
Claude Code skill for adversarial code review - spawns a GPT‑5.6 Sol reviewer in a herdr split pane to attack your diff, then interrogates its findings before relaying verdicts.
There was an error while loading. Please reload this page .
0
forks
Report repository
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
