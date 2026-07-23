---
source: "https://github.com/rizzdev/claude-detective"
hn_url: "https://news.ycombinator.com/item?id=49024252"
title: "A Claude skill that grills you so hard. Within a nice clean interface"
article_title: "GitHub - rizzdev/claude-detective: Ask a human a whole batch of multiple-choice questions in a slick local web form and get structured answers back — a Claude Code skill (and standalone CLI). · GitHub"
author: "rizzdev"
captured_at: "2026-07-23T17:10:35Z"
capture_tool: "hn-digest"
hn_id: 49024252
score: 1
comments: 0
posted_at: "2026-07-23T16:26:51Z"
tags:
  - hacker-news
  - translated
---

# A Claude skill that grills you so hard. Within a nice clean interface

- HN: [49024252](https://news.ycombinator.com/item?id=49024252)
- Source: [github.com](https://github.com/rizzdev/claude-detective)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T16:26:51Z

## Translation

作品名：あなたを焦らすクロードのスキル。すっきりとしたインターフェイス内
記事のタイトル: GitHub - rizzdev/claude-detective: 洗練されたローカル Web フォームで人間に複数選択の質問を丸ごと質問し、構造化された回答を取得します。これは、クロード コード スキル (およびスタンドアロン CLI) です。 · GitHub
説明: 洗練されたローカル Web フォームで複数選択の質問を人間に質問し、構造化された回答を返します。これは、Claude Code スキル (およびスタンドアロン CLI) です。 - rizzdev/クロード刑事

記事本文:
GitHub - rizzdev/claude-detective: 洗練されたローカル Web フォームで複数選択の質問を人間に質問し、構造化された回答を返します。これは、クロード コード スキル (およびスタンドアロン CLI) です。 · GitHub
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

リズデブ
/
クロード探偵
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
34 コミット 34 コミット .github/ workflows .github/ workflows 資産 アセット ドキュメント ドキュメント テスト test .gitignore .gitignore ライセンス ライセンス README.md README.md SKILL.md SKILL.md Detective.mjs Detective.mjs install.ps1 install.ps1 install.sh install.sh package.json package.json すべてのファイルを表示 リポジトリ ファイルナビゲーション
AI の質問に 1 つずつ答えるのはやめましょう。
Claude Code スキルは、洗練されたローカル Web フォームで多肢選択式の質問のバッチ全体を提供し、すべての選択肢について長所、短所、推奨事項を示し、回答を構造化された JSON として読み取ります。
▶ デモを視聴するか、npx github:rizzdev/rizzdev-detective --demo を実行してください
実際のセッション: あらゆるオプションの長所/短所と推奨事項に加えて、その場で決定、再考、調査、さらなる反発を決定します。
クロードが認証モデル、データ形状、どの機能をカットするかなど、あなたからの多くの決定を必要とする場合、通常はメッセージごとに 1 つの質問をドリップします。それは遅いし、全体像を一度に見ることはできません。
クロード刑事はそれをひっくり返します。クロードは質問セット全体を書き、小さなローカル Web ページを起動します。すると、目の前に並べられた各オプションのトレードオフを使用して、ラジオ ボタン、チェックボックス、はい/いいえの錠剤、「その他」ボックスなど、すべてを 1 回のパスで優先順位付けします。 [送信] をクリックすると、回答がクリーンな JSON として Claude に返されます。
アカウントもクラウドも依存関係もありません。ノードの 1 つのファイル、ローカルホストのみ。
🗂️ バッチ全体を一度に
セクションにグループ化された 1 つのスクロール可能なフォーム - 任意の順序で回答します。
⚖️ すべてのオプションにはトレードオフがあります
それぞれの選択肢には、👍 賛成派と👎 反対派のラインが表示されます。
💡 お勧め

質問ごとに
クロードは提案された選択にマークを付け、その理由を述べます。
❓ 「問題」のコンテキスト
それぞれの質問は、その答えによって何がブロックが解除されるのかを説明できます。
🔬 最初の研究 + 引用された調査結果
--deep / --online を使用すると、Claude は最初に調査し、クリック可能なソースを含む結果説明パネルを表示します。推奨事項は file:line とリンクを引用します。
📡 ライブ適応インタビュー
--live は 1 つのタブを開いたままにし、回答すると新しい質問をストリーミングします。選択肢に基づいて分岐し、過去の回答を修正すると、下流で再質問されます。
🔘 各入力タイプ
単一選択、複数選択、およびコンパクトな Yes/No 錠剤。
🧮 スマートな長いリスト
約 15 個の短いオプションが 2 つのカラムに自動的に流れ、コンパクトさを保ちます。
✍️ 常に避難口
質問ごとの「その他」ボックス + グローバルな「他に何かありますか?」分野。
🎨 ターミナル / CLI の美しさ
等幅、シェルプロンプトヘッダー、@clack スタイルの ◇ ステップガター、○/◉ ラジオと [ ]/[×] チェックボックス、スティッキー ⏎ 送信バー。
📦 依存関係ゼロ
Pure Node の組み込み。 1 つの .mjs ファイル。ローカルホストのみ。
🔌 スキルまたは CLI
クロード コード スキルとして使用するか、スタンドアロンで実行します。
⚡ 5秒で試してください
インストールも構成も不要 — 組み込みデモを GitHub から直接実行します (ノード ≥ 22 が必要):
npx github:rizzdev/rizzdev-detective --demo
ブラウザーでサンプルのインタビューが開きます。調査結果パネル、長所/短所、はい/いいえの丸薬、ドラッグしてランクに移動、キーボード ナビゲーションが表示されます。次に、自分自身の質問に向けてください。
npx github:rizzdev/rizzdev-detectivequestions.json --out results.json
🚀 インストール (クロード コード スキルとして)
git clone https://github.com/rizzdev/rizzdev-detective.git
cd リズデヴ刑事
# Linux / macOS / WSL
./install.sh
# Windows (PowerShell)
。 \.ps1 をインストールします
インストール スクリプトはこのフォルダーを ~/.claude/skills/claude-detective にシンボリックリンクするため、git pull によりすべてのマシンでフォルダーが最新に保たれます。 Claを再起動します

ude Code (または /reload-skills を実行) すれば準備完了です。
rizzdev-detective から名前変更されました。 /rizzdev-detective は、非推奨のエイリアスとして引き続き機能します。 GitHub リポジトリ/npx パスの名前変更は保留中のフォローアップであるため、以下のクローン URL とデモ URL では依然として古い名前が参照されています。
CLAUDE_SKILLS_DIR 環境変数で宛先をオーバーライドします。
クロードにたくさんの質問をするか、次のコマンドを実行してください。
/クロード刑事
クロードが質問を作成し、ブラウザでフォームを開いて待ちます。回答して送信すると、選択が続行されます。それがループ全体です。
それが始まると、「たくさんの質問」、アンケート、調査を求めるか、 /claude-detective を実行します。
それができない場合: 簡単な質問を 1 つ (クロードはチャットで質問するだけです)。
🧩 質問の作成 (スキーマ)
クロードは JSON ファイルを作成し、それに対してサーバーを実行します。形状:
{
"title" : " オプションのページ見出し " ,
「セクション」: [
{
"title" : " 認証 " ,
「質問」: [
{
"id" : " 認証 " ,
"text" : " v1 の認証モデルはどれですか? " ,
"why" : " 第 1 週のうち、設備と配管にどれだけの時間が費やされるかを決定します。 " ,
"タイプ" : " シングル " 、
"recommendation" : { "optionId" : " token " , "why" : " 動作する v1 への最速パス。 " },
"オプション" : [
{ "id" : " トークン " 、 "label" : " 貼り付けられた長期トークン " 、 "pro" : " OAuth 機能なし " 、 "con" : " 手動ローテーション " },
{ "id" : " oauth " 、 "label" : " OAuth " 、 "pro" : " クリーンな UX " 、 "con" : " 数週間の作業 " }
]、
"許可その他" : true
}、
{
"id" : "マージ" ,
"text" : " マスターにマージしてプッシュしますか? " ,
"type" : "はいいいえ" 、
"recommendation" : { "optionId" : " yes " , "why" : " これはリポジトリの通常の同期フローです。 " }
}
】
}
】
}
質問の種類
フィールド — id (一意)、text、なぜ? 、 タイプ？ 、 おすすめ？ ( {optionId?, Why?} )、オプション ( {id, label, pro?, con?} )、allowOther? (シングル/マルチのデフォルトはオン)。トップル

ヴェルが称号を取るのか？ 、任意の所見ですか？研究ブリーフィング ( { summary,sources: [{ label, ref }] } - URL refs リンク、file:line refs がタグ付けされます)、さらにセクションまたはフラットな質問配列のいずれか。
送信時に、サーバーは JSON を stdout (指定されている場合は --out <file>) に出力します。
{
"答え" : {
"auth" : { "selected" : [ " トークン " ]、 "other" : " " }、
"マージ" : { "選択済み" : [ " はい " ]、 "その他" : " " }
}、
"globalNote" : " 全体的なメモ " ,
"submittedAt" : " 2026-07-01T13:05:00.000Z "
}
selected は、選択したオプション ID を保持します (シングル/はいいいえの場合は 0 ～ 1、マルチの場合は 0 ～ n)。もう 1 つは質問ごとのテキスト ボックスです。 globalNote はフォームの終わりのフィールドです。未回答の質問は空の選択状態で返されます。部分的に送信しても問題ありません。
クロードは必要ありません。ノードだけで済みます。
ノード探偵.mjs 質問.json --out results.json
# または GitHub から直接、クローンはありません:
npx github:rizzdev/rizzdev-detectivequestions.json --out results.json
npx github:rizzdev/rizzdev-detective --demo # 組み込みサンプル
空いているポートを選択し、ブラウザを開き (URL の出力に戻ります)、送信するまでブロックし、結果を書き込んで出力して終了します。
📡 ライブモード — 適応型インタビュー
回答に応じて分岐するデシジョン ツリーの場合、クロードはライブ インタビューを実行できます。1 つの開いたタブで、質問が次々と流れてきます。
Detective.mjs --live --out Transcript.json # 永続サーバー (SSE)
Detective.mjs Push <batch.json> # 質問バッチを挿入する
Detective.mjs ユーザーが応答するまで # ブロック待機します
Detective.mjs retract --from <batchId > # 古いダウンストリーム バッチを削除します
探偵.mjs 終了 --out トランスクリプト.json # 終了 + トランスクリプトを出力
クロードはバッチをプッシュし、あなたの回答を待ってから、あなたの発言に基づいて分岐した次の質問をプッシュします。以前の回答の考えを変えると、撤回 + 再‑

下流のすべてに質問します。依存関係は依然としてゼロ (組み込み http を介した SSE)、依然として localhost のみです。
クロードはquestions.jsonを書きます
│
▼
ノード Detective.mjsquestions.json ──► は 127.0.0.1 で 1 ページのフォームを提供します
│ (ブラウザが自動的に開きます)
▼
回答 + 送信 ──► POST /submit ──► 結果が JSON として出力され、サーバーが終了します
│
▼
クロードは答えを読んで続けます
1 回の実行 = 1 回の面接。ステートレスでローカルホストのみであり、自宅に電話をかけることはありません。
ノード --test # 31 ユニット + 統合テスト、ゼロdeps
コードは 1 つのファイル ( Detective.mjs ) で、独立してテストされた小さな要素が含まれています:loadQuestions (検証 + 正規化)、renderPage / renderBatchHtml 、normalizeResults、およびワンショット サーブ + 永続的serveLive サーバーです。元の設計については docs/DESIGN.md を、ビルド プランについては docs/PLAN.md を参照してください。
洗練されたローカル Web フォームで複数選択の質問を人間に質問すると、構造化された回答が返されます。これは、Claude Code スキル (およびスタンドアロン CLI) です。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Ask a human a whole batch of multiple-choice questions in a slick local web form and get structured answers back — a Claude Code skill (and standalone CLI). - rizzdev/claude-detective

GitHub - rizzdev/claude-detective: Ask a human a whole batch of multiple-choice questions in a slick local web form and get structured answers back — a Claude Code skill (and standalone CLI). · GitHub
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
rizzdev
/
claude-detective
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits .github/ workflows .github/ workflows assets assets docs docs test test .gitignore .gitignore LICENSE LICENSE README.md README.md SKILL.md SKILL.md detective.mjs detective.mjs install.ps1 install.ps1 install.sh install.sh package.json package.json View all files Repository files navigation
Stop answering AI questions one at a time.
A Claude Code skill that hands you a whole batch of multiple‑choice questions in a slick local web form — with pros, cons, and a recommendation on every option — then reads your answers back as structured JSON.
▶ watch the demo · or run npx github:rizzdev/rizzdev-detective --demo
A real session: pros/cons and a recommendation on every option, plus in-place you decide · rethink · research · more pushback.
When Claude needs a bunch of decisions from you — auth model, data shape, which features to cut — it usually drips them out one question per message . That's slow, and you can't see the whole picture at once.
claude-detective flips it: Claude writes the entire question set, spins up a tiny local web page, and you triage everything in one pass — radio buttons, checkboxes, yes/no pills, "other" boxes — with the trade‑offs of each option laid out in front of you. Hit Submit and the answers land back in Claude as clean JSON.
No account, no cloud, no dependencies. One file of Node, localhost only.
🗂️ Whole batch at once
One scrollable form, grouped into sections — answer in any order.
⚖️ Trade‑offs on every option
Each choice can show a 👍 pro and 👎 con line.
💡 A recommendation per question
Claude marks its suggested pick and says why .
❓ "The problem" context
Each question can explain what the answer unblocks.
🔬 Research‑first + cited findings
With --deep / --online , Claude investigates first and shows a findings briefing panel with clickable sources; recommendations cite file:line and links.
📡 Live adaptive interviews
--live keeps one tab open and streams new questions in as you answer — branch on your choices, revise a past answer and it re‑asks downstream.
🔘 Every input type
Single‑select, multi‑select, and compact yes/no pills .
🧮 Smart long lists
~15 short options auto‑flow into two columns to stay compact.
✍️ Always an escape hatch
Per‑question "Other" box + a global "anything else?" field.
🎨 Terminal / CLI aesthetic
Monospace, a shell‑prompt header, @clack ‑style ◇ step gutter, ○/◉ radios and [ ]/[×] checkboxes, a sticky ⏎ submit bar.
📦 Zero dependencies
Pure Node built‑ins. One .mjs file. Localhost only.
🔌 Skill or CLI
Use it as a Claude Code skill, or run it standalone.
⚡ Try it in 5 seconds
No install, no config — run the built-in demo straight from GitHub (needs Node ≥ 22):
npx github:rizzdev/rizzdev-detective --demo
It opens a sample interview in your browser — findings panel, pros/cons, yes/no pills, drag-to-rank, keyboard nav. Then point it at your own questions:
npx github:rizzdev/rizzdev-detective questions.json --out results.json
🚀 Install (as a Claude Code skill)
git clone https://github.com/rizzdev/rizzdev-detective.git
cd rizzdev-detective
# Linux / macOS / WSL
./install.sh
# Windows (PowerShell)
. \i nstall.ps1
The install script symlinks this folder into ~/.claude/skills/claude-detective , so a git pull keeps it current on every machine. Restart Claude Code (or run /reload-skills ) and you're set.
Renamed from rizzdev-detective ; /rizzdev-detective still works as a deprecated alias. The GitHub repo / npx path rename is a pending follow-up, so clone and demo URLs below still reference the old name.
Override the destination with the CLAUDE_SKILLS_DIR env var.
Just ask Claude for a lot of questions, or run the command:
/claude-detective
Claude will author the questions, open the form in your browser, and wait. You answer, submit, and it continues with your choices. That's the whole loop.
When it kicks in: you ask for "a lot of questions", a questionnaire, or a survey — or run /claude-detective .
When it won't: a single quick question (Claude just asks in chat).
🧩 Authoring questions (the schema)
Claude writes a JSON file and runs the server against it. The shape:
{
"title" : " Optional page heading " ,
"sections" : [
{
"title" : " Auth " ,
"questions" : [
{
"id" : " auth " ,
"text" : " Which auth model for v1? " ,
"why" : " Determines how much of week 1 goes to plumbing vs features. " ,
"type" : " single " ,
"recommendation" : { "optionId" : " token " , "why" : " Fastest path to a working v1. " },
"options" : [
{ "id" : " token " , "label" : " Pasted long-lived token " , "pro" : " No OAuth work " , "con" : " Manual rotation " },
{ "id" : " oauth " , "label" : " OAuth " , "pro" : " Clean UX " , "con" : " Weeks of work " }
],
"allowOther" : true
},
{
"id" : " merge " ,
"text" : " Merge to master and push? " ,
"type" : " yesno " ,
"recommendation" : { "optionId" : " yes " , "why" : " It's the repo's normal sync flow. " }
}
]
}
]
}
Question types
Fields — id (unique), text , why? , type? , recommendation? ( {optionId?, why?} ), options ( {id, label, pro?, con?} ), allowOther? (defaults on for single/multi). Top level takes a title? , an optional findings? research briefing ( { summary, sources: [{ label, ref }] } — URL refs link, file:line refs get tagged), plus either sections or a flat questions array.
On submit, the server prints JSON to stdout (and to --out <file> if given):
{
"answers" : {
"auth" : { "selected" : [ " token " ], "other" : " " },
"merge" : { "selected" : [ " yes " ], "other" : " " }
},
"globalNote" : " any overall notes " ,
"submittedAt" : " 2026-07-01T13:05:00.000Z "
}
selected holds the chosen option id s (0–1 for single/yesno, 0–n for multi). other is the per‑question text box; globalNote is the end‑of‑form field. Unanswered questions come back with an empty selected — partial submissions are fine.
No Claude required — it's just Node:
node detective.mjs questions.json --out results.json
# or straight from GitHub, no clone:
npx github:rizzdev/rizzdev-detective questions.json --out results.json
npx github:rizzdev/rizzdev-detective --demo # built-in sample
It picks a free port, opens your browser (falls back to printing the URL), blocks until you submit, then writes/prints the results and exits.
📡 Live mode — adaptive interviews
For decision trees that branch on your answers, Claude can run a live interview : one open tab, questions streaming in as you go.
detective.mjs --live --out transcript.json # persistent server (SSE)
detective.mjs push < batch.json > # inject a question batch
detective.mjs wait # block until the user answers
detective.mjs retract --from < batchId > # drop stale downstream batches
detective.mjs finish --out transcript.json # end + print the transcript
Claude pushes a batch, waits for your answers, then pushes the next questions branched on what you said . Change your mind on an earlier answer and it retracts + re‑asks everything downstream. Still zero dependencies (SSE over the built‑in http ), still localhost‑only.
Claude writes questions.json
│
▼
node detective.mjs questions.json ──► serves a one-page form on 127.0.0.1
│ (auto-opens your browser)
▼
you answer + Submit ──► POST /submit ──► results printed as JSON, server exits
│
▼
Claude reads the answers and continues
One run = one interview. It's stateless, localhost‑only, and never phones home.
node --test # 31 unit + integration tests, zero deps
The code is one file ( detective.mjs ) with small, independently‑tested pieces: loadQuestions (validate + normalize), renderPage / renderBatchHtml , normalizeResults , and the one‑shot serve + persistent serveLive servers. See docs/DESIGN.md for the original design and docs/PLAN.md for the build plan.
Ask a human a whole batch of multiple-choice questions in a slick local web form and get structured answers back — a Claude Code skill (and standalone CLI).
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
