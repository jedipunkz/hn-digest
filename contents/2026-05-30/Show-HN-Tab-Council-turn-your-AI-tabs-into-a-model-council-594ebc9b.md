---
source: "https://github.com/vaddisrinivas/tab-council"
hn_url: "https://news.ycombinator.com/item?id=48323971"
title: "Show HN: Tab Council – turn your AI tabs into a model council"
article_title: "GitHub - vaddisrinivas/tab-council: Chrome MV3 extension that turns AI tabs into a structured model council · GitHub"
author: "srinivasvaddi"
captured_at: "2026-05-30T11:46:21Z"
capture_tool: "hn-digest"
hn_id: 48323971
score: 4
comments: 0
posted_at: "2026-05-29T15:02:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tab Council – turn your AI tabs into a model council

- HN: [48323971](https://news.ycombinator.com/item?id=48323971)
- Source: [github.com](https://github.com/vaddisrinivas/tab-council)
- Score: 4
- Comments: 0
- Posted: 2026-05-29T15:02:25Z

## Translation

タイトル: Show HN: Tab Council – AI タブをモデル評議会に変える
記事タイトル: GitHub - vaddisrinivas/tab-council: AI タブを構造化モデル評議会に変える Chrome MV3 拡張機能 · GitHub
説明: AI タブを構造化モデル評議会に変える Chrome MV3 拡張機能 - vaddisrinivas/tab-council

記事本文:
GitHub - vaddisrinivas/tab-council: AI タブを構造化モデル評議会に変える Chrome MV3 拡張機能 · GitHub
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
ヴァディスリニバス
/
タブ評議会
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .github .github アセット アセット docs docs scripts scripts src src test test video/ tab-council-demo video/ tab-council-demo .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンスライセンス PRIVACY.md PRIVACY.md README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md SECURITY.md SECURITY.md manifest.json マニフェスト.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すでに使用している AI タブを構造化モデル評議会に変えます。
Tab Council は、ChatGPT、Claude、Gemini、Perplexity、Merlin、Grok にわたる重要なプロンプトをクロスチェックするための Chrome MV3 拡張機能です。 tab-council という名前のタブ グループを作成し、その中にプロバイダー タブを配置し、独立した回答、批評、総合判断、およびオプションの承認/拒否フローを通じて 1 つのプロンプトを実行します。
API キーはありません。バックエンドはありません。アカウント、タブ、ローカル Chrome ストレージ。
v0.1.0-alpha のデモビデオを見る
模範解答が 1 つあると便利です。小さな議会のほうが安全です。
1 つのプロンプトを多数の AI ツールに送信し、回答を並べて表示するブラウザ拡張機能がすでに存在します。 Tab Council は、より狭い範囲のことを実行しようとしています。つまり、すでに開いている AI タブ上で完全なモデル評議会ワークフローを実行します。
これは、発売レビュー、製品決定、アーキテクチャ チェック、調査総合、ポリシー草案、および答えを信頼する前に独立した推論と目に見える反対意見が必要なあらゆる質問向けに構築されています。
既存の AI タブを使用します。OpenRouter セットアップ、プロバイダー API キー、コピーされたモデル名簿はありません。
Chrome タブ グループから開始します。tab-council の下にグループ化したタブが評議会です。
放送だけでなく議会ラウンドを運営：独立

回答、批評、意見の相違チェック、オプションの最終的な立場、総合。
明示的な役割: メンバー、裁判官、検証者、オブザーバー、除外。
構造化された出力を処理します: JSON 解析、1 回の修復パス、サイレントな不良トランスクリプトの代わりに大規模なエラー。
最終レビューを追加: 承認/拒否により、裁判官ではないタブが合成評決に異議を申し立てることができます。
このアルファ版は、拡張機能ベースのオーケストレーションに焦点を当てています。
tab-council という名前のアクティブな Chrome タブ グループを検出します。
グループ化された各 AI タブを明示的なモデル シートとして使用します。
プロバイダー UI が公開する、表示されているモデル ラベルを読み取ります。
すべてのタブに役割を持たせます: Member 、 Judge 、 Validator 、 Observer 、または Exclude 。
ラウンド 1 の独立した回答、ラウンド 2 の丁寧な批評、オプションのラウンド 3、統合、およびオプションの承認/拒否を実行します。
すべてのフェーズで構造化された JSON が必要です。
プロバイダーが有効な JSON ではなくほぼ JSON を返した場合に、1 つの JSON 修復パスを実行します。
実行状態と履歴は chrome.storage.local にのみ保存されます。
ローカルの Markdown トランスクリプトをエクスポートします。
ChatGPT: chatgpt.com、chat.openai.com
Perplexity: perplexity.ai、www.perplexity.ai
マーリン: getmerlin.in 、 www.getmerlin.in 、 extension.getmerlin.in
Grok: grok.com 、 x.com 、 *.x.ai
ローカル/開発の汎用 AI ページ: localhost および 127.0.0.1
不明なパブリック Web ページは、ストアセーフ マニフェストでは広く有効になっていません。既知のアダプターは、プロバイダー UI が変更されると目に見えて障害が発生します。
必要な AI プロバイダーのタブを開きます。
各プロバイダー タブ内のモデル (Claude Sonnet、ChatGPT モデル、Gemini、Perplexity、Grok など) を選択します。
タブカウンシルのサイドパネルを開きます。
役割を割り当ててモードを選択します。
ラウンド、最終評決、承認/拒否の結果を確認し、エクスポートします。
メンバーはディベートラウンドに参加し、審査員に選ばれることができます。
ジャッジはメンバーとして参加しており、合成には優先されます。
バリデーターは議論をスキップし、評価のみを行います

最終的な評決に対して異議を申し立てるか拒否権を発動します。
オブザーバーは名簿に含まれますが、自動化されることはありません。
除外は実行では無視されます。
Quick : 現在のスレッドを使用してより高速なパス。
バランスのとれた : 新鮮なスレッド、批評、判断者の総合。
徹底 : 合成前に最終ステートメントを強制します。
厳格 : 最終声明と批准/拒否権を強制します。
手動コントロールを使用すると、モードをオーバーライドできます。つまり、ラウンド 3 を強制し、ジャッジ モードを選択し、承認/拒否を有効にし、生のエクスポートを含め、出力ビューを選択します。
構造化出力と JSON 修復
すべてのプロバイダーは、各フェーズで厳密な JSON を返すように求められます。これにより、サイドパネルの信頼性が高まります。信頼スコア、意見の相違、評決、および拒否権は、散文からかき出すのではなく、解析することができます。
JSON 修復とは次のことを意味します。モデルがほぼ有効な JSON (マークダウン フェンス、末尾のコメント、不正な形式の引用エスケープなど) を返した場合、Tab Council はその同じタブに修復プロンプトを送信し、有効な JSON のみを要求します。修復された応答は UI でマークされます。修復が失敗すると、タブは静かにトランスクリプトを破損するのではなく、大きな音で失敗します。
Tab Council には、拡張機能が所有するサーバー、分析、または追跡ピクセルはありません。
プロンプトは、選択したプロバイダーのタブに貼り付けられます。
後のラウンドでは、選択したプロバイダー タブ全体で成功したモデルの出力を意図的に共有します。
ローカルの実行状態は chrome.storage.local にあります。
エクスポートはブラウザによってローカルに生成されます。
[ローカルをクリア] は、保存されている拡張機能の実行データを削除します。
プロバイダー Web サイトは、独自のタブに送信されたコンテンツを引き続き受信します。正確な境界については PRIVACY.md を参照してください。
GitHub から v0.1.0-alpha リリース zip をダウンロードします。
github.com/vaddisrinivas/tab-council/releases
リリースを解凍し、manifest.json が含まれるフォルダーを選択します。
Chrome で AI プロバイダーにサインインします。
これらのタブを tab-council という名前のグループに入れます。
タブカウンシル拡張アイコンをクリックします。
npm

インストールする
npm実行チェック
決定的な偽のプロバイダー ハーネスを実行します。
npm は偽 AI を実行します
http://localhost:4173/ を 2 つ以上のタブで開き、それらをタブ - カウンシル グループに入れ、localhost 権限を付与してカウンシルを実行します。このページは、フェーズごとに構造化された JSON を返します。
manifest.json : MV3 権限、サイド パネル、サービス ワーカー、プロバイダー ホストの許可リスト。
src/background.js : タブグループの検出、権限、実行ステートマシン、プロンプトオーケストレーション。
src/contentScript.js : プロバイダー ページ自動化アダプター。
src/sidepanel.* : オペレーター UI、ロール、実行コントロール、トランスクリプト、エクスポート。
src/prompts.js : 構造化プロンプトビルダーとマークダウンエクスポート。
src/shared.js : プロバイダーのレジストリ、設定、解析、ヘルパー コントラクト。
詳細については、 docs/ARCHITECTURE.md を参照してください。
AI タブを構造化モデル評議会に変える Chrome MV3 拡張機能
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

Chrome MV3 extension that turns AI tabs into a structured model council - vaddisrinivas/tab-council

GitHub - vaddisrinivas/tab-council: Chrome MV3 extension that turns AI tabs into a structured model council · GitHub
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
vaddisrinivas
/
tab-council
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .github .github assets assets docs docs scripts scripts src src test test video/ tab-council-demo video/ tab-council-demo .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md SECURITY.md SECURITY.md manifest.json manifest.json package.json package.json View all files Repository files navigation
Turn the AI tabs you already use into a structured model council.
Tab Council is a Chrome MV3 extension for cross-checking important prompts across ChatGPT, Claude, Gemini, Perplexity, Merlin, and Grok. Create a tab group named tab-council , place your provider tabs inside it, and run one prompt through an independent answer, critique, judge synthesis, and optional Ratify/Veto flow.
No API keys. No backend. Your accounts, your tabs, your local Chrome storage.
Watch the v0.1.0-alpha demo video
One model answer is useful. A small council is safer.
There are already browser extensions that send one prompt to many AI tools and show the answers side by side. Tab Council is trying to do something narrower: run the full model-council workflow over the AI tabs you already have open.
It is built for launch reviews, product decisions, architecture checks, research synthesis, policy drafts, and any question where you want independent reasoning plus visible dissent before you trust the answer.
Uses existing AI tabs : no OpenRouter setup, no provider API keys, no copied model roster.
Starts from a Chrome tab group : the tabs you group under tab-council are the council.
Runs council rounds, not just broadcast : independent answers, critique, disagreement check, optional final positions, synthesis.
Has explicit roles : Member, Judge, Validator, Observer, Exclude.
Handles structured output : JSON parsing, one repair pass, and loud failures instead of silent bad transcripts.
Adds final review : Ratify/Veto lets non-judge tabs challenge the synthesized verdict.
This alpha focuses on extension-based orchestration:
Detects the active Chrome tab group named tab-council .
Uses each grouped AI tab as an explicit model seat.
Reads visible model labels where the provider UI exposes them.
Lets every tab take a role: Member , Judge , Validator , Observer , or Exclude .
Runs Round 1 independent answers, Round 2 polite critique, optional Round 3, synthesis, and optional Ratify/Veto.
Requires structured JSON from every phase.
Runs one JSON repair pass when a provider returns almost-JSON instead of valid JSON.
Stores run state and history only in chrome.storage.local .
Exports a local Markdown transcript.
ChatGPT: chatgpt.com , chat.openai.com
Perplexity: perplexity.ai , www.perplexity.ai
Merlin: getmerlin.in , www.getmerlin.in , extension.getmerlin.in
Grok: grok.com , x.com , *.x.ai
Local/dev generic AI pages: localhost and 127.0.0.1
Unknown public web pages are not broadly enabled in the store-safe manifest. Known adapters fail visibly when a provider UI changes.
Open the AI provider tabs you want.
Select the model inside each provider tab, such as Claude Sonnet, a ChatGPT model, Gemini, Perplexity, or Grok.
Open the Tab Council side panel.
Assign roles and choose a mode.
Review the rounds, final verdict, Ratify/Veto results, and export.
Member participates in debate rounds and can be selected as judge.
Judge participates as a member and is preferred for synthesis.
Validator skips debate and only ratifies or vetoes the final verdict.
Observer appears in the roster but is never automated.
Exclude is ignored for the run.
Quick : faster pass using current threads.
Balanced : fresh threads, critique, judge synthesis.
Thorough : forces final statements before synthesis.
Rigorous : forces final statements and Ratify/Veto.
Manual controls let you override the mode: force Round 3, choose judge mode, enable Ratify/Veto, include raw export, and select output views.
Structured Output and JSON Repair
Every provider is asked to return strict JSON for each phase. That makes the side panel reliable: confidence scores, disagreements, verdicts, and vetoes can be parsed instead of scraped from prose.
JSON repair means this: if a model returns almost valid JSON, for example markdown fences, trailing comments, or malformed quote escaping, Tab Council sends that same tab a repair prompt and asks for valid JSON only. The repaired response is marked in the UI. If repair fails, the tab fails loud instead of silently corrupting the transcript.
Tab Council has no extension-owned server, analytics, or tracking pixels.
Prompts are pasted into the provider tabs you selected.
Later rounds intentionally share successful model outputs across selected provider tabs.
Local run state lives in chrome.storage.local .
Exports are generated locally by your browser.
Clear Local removes stored extension run data.
Provider websites still receive the content sent into their own tabs. See PRIVACY.md for the exact boundary.
Download the v0.1.0-alpha release zip from GitHub:
github.com/vaddisrinivas/tab-council/releases
Unzip the release and select the folder containing manifest.json .
Sign into your AI providers in Chrome.
Put those tabs in a group named tab-council .
Click the Tab Council extension icon.
npm install
npm run check
Run the deterministic fake provider harness:
npm run fake-ai
Open http://localhost:4173/ in two or more tabs, put them in a tab-council group, grant localhost permission, and run a council. The page returns structured JSON for every phase.
manifest.json : MV3 permissions, side panel, service worker, provider host allowlist.
src/background.js : tab-group discovery, permissions, run state machine, prompt orchestration.
src/contentScript.js : provider page automation adapters.
src/sidepanel.* : operator UI, roles, run controls, transcript, export.
src/prompts.js : structured prompt builders and Markdown export.
src/shared.js : provider registry, settings, parsing, helper contracts.
Read more in docs/ARCHITECTURE.md .
Chrome MV3 extension that turns AI tabs into a structured model council
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
