---
source: "https://github.com/trekhleb/yesbrainer"
hn_url: "https://news.ycombinator.com/item?id=48936513"
title: "Yes-Brainer: a councils of LLMs debating to a consencus (open-sourced, BYOK)"
article_title: "GitHub - trekhleb/yesbrainer: 🧠 A council of AI models for the decisions that aren't no-brainers — they answer in parallel, debate to consensus, or get judged to a verdict. Browser-only, open source, bring your own keys (BYOK), no backend. · GitHub"
author: "okso_app"
captured_at: "2026-07-16T17:05:14Z"
capture_tool: "hn-digest"
hn_id: 48936513
score: 2
comments: 0
posted_at: "2026-07-16T16:12:39Z"
tags:
  - hacker-news
  - translated
---

# Yes-Brainer: a councils of LLMs debating to a consencus (open-sourced, BYOK)

- HN: [48936513](https://news.ycombinator.com/item?id=48936513)
- Source: [github.com](https://github.com/trekhleb/yesbrainer)
- Score: 2
- Comments: 0
- Posted: 2026-07-16T16:12:39Z

## Translation

タイトル: Yes-Brainer: コンセンサスに向けて議論する LLM の評議会 (オープンソース、BYOK)
記事のタイトル: GitHub - trekhleb/yesbrainer: 🧠 AI の評議会は、簡単ではない決定をモデル化します。回答は並行して行われ、合意に達するまで議論され、あるいは評決が下されます。ブラウザのみ、オープンソース、Bring Your Own Key (BYOK)、バックエンドなし。 · GitHub
説明: 🧠 AI の評議会は、簡単ではない決定をモデル化します。回答は並行して行われ、合意に達するまで議論され、あるいは評決が下されます。ブラウザのみ、オープンソース、Bring Your Own Key (BYOK)、バックエンドなし。 - トレクレブ/イエスブレイナー

記事本文:
GitHub - trekhleb/yesbrainer: 🧠 AI の評議会は、簡単ではない決定をモデル化します。彼らは並行して回答し、合意に達するまで議論するか、評決に至るまで判断されます。ブラウザのみ、オープンソース、Bring Your Own Key (BYOK)、バックエンドなし。 · GitHub
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
トレクレブ
/
イエスブレイナー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github/ workflows .github/ workflows public public scripts scripts src src testing testing .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc CLAUDE.md CLAUDE.md DEVELOPMENT.md DEVELOPMENT.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md eslint.config.js eslint.config.js Index.html Index.html package-lock.json package-lock.json package.json package.json playwright.config.ts playwright.config.ts tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json tsconfig.tests.json tsconfig.tests.json tsconfig.unit.json tsconfig.unit.json vite.config.ts vite.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Yes-Brainer — 簡単ではない意思決定のための AI モデルの協議会。 1 つの質問が複数の LLM に展開され、ブラウザーのタブを操作する代わりに、1 か所で審議が行われます。独立した回答、ピアレビュー、討論、合成された評決、意見の相違が表示されます。
試してみてください:yesbrainer.ai — このリポジトリから構築された公式デプロイメント。他のドメインは非公式コピーとして扱います。サインアップなし: API キーを貼り付けて質問するか、最初に録画されたデモ評議会を開いてください。キーは必要ありません。
3 つの審議形式 — 🔀 並列回答 (独立、総合なし)、⚖️ 裁判評決 (匿名のピア投票 + 裁判官)、🤝 コンセンサス討論 (複数ラウンドの討論、調停者が収束を促進)。
バックエンドもアカウントも、あなた以外が保持するデータもゼロです。キー、履歴、設定はブラウザー内に存在します。あなたのプロンプトは

nd キーは、選択したモデル プロバイダーにのみ直接送信されます。
Anthropic、OpenAI、Google、Groq、OpenRouter、オプトイン Ollama など、さまざまなモデル プロバイダーが 1 つの協議会でサポートされています。
サブスクリプションやペイウォールはありません。モデルプロバイダーにのみ直接支払います。
モバイルフレンドリーで、PWA としてインストール可能。マルチターンセッション。過去の審議会の補足。
静的バンドル - GitHub Pages / 任意の CDN。ホストするもの、操作するものは何もありません。
フロントページ: 3 つの審議モード、サイドバーに録画されたデモ評議会、およびスタート カード - 1 つの画面でサインアップは不要です。クリックして試してみてください。
そして、評議会の作成が全体のセットアップです。審議モードを選択し、モデルを配置し (能力とコンテキストのサイズが一目でわかる)、誰を審判するかを選択します。または、単に「✨ 最もスマートな利用可能」を押して開始します。
その後、評議会が召集されます。1 つの質問がすべての議席に広げられ、回答が並べられ、意見の相違が一目瞭然になります。その広がりが合図です。
💡 karpathy/llm-council からインスピレーションを得た — オーケストレーションのコンセプトと彼のプロンプトデザイン (ジャッジ統合、ピア評価) が出発点でした。違いは洗練されたものではなく形状です。オリジナル (およびその 80 以上のフォークのほとんど) は、実行する必要があるローカル サーバーの背後で 1 つの固定された回答→レビュー→合成パスを実行しますが、Yes-Brainer は 3 つの審議構造を備えたゼロバックエンドのブラウザ PWA です。これには、コンバージェンス (コンセンサス モード) を参照する実際の複数ラウンドの議論が含まれます。デバイス上の永続性による複数ターンのフォローアップ、合意の広がりが可視化された匿名化されたルーブリック投票、および電話ファーストの UI (他のものとの比較)。
Yes-Brainer であると主張するサイトに API キーを貼り付ける前に、アドレス バーに yesbrainer.ai と表示されていることを確認してください。この README とアプリはそのドメインとそのコピーに同意する必要があります。

ed サイトは URL 以外のすべてを偽装できます。爆発範囲にも上限を設けます。プロバイダーのコンソールで使用量制限を設定したこのアプリ専用の専用キーを使用し、何か問題があると思われる瞬間にそこでキーを取り消します。ソフトウェアは現状のまま提供され、保証はありません (AGPL-3.0 §§15–16) — SECURITY.md は、コードが保護するものと保護できないものを詳しく説明します。
複雑な質問や重要な質問の場合、1 つのモデルの自信に満ちた答えだけでは十分ではありません。複数の独立したテイクとクロスチェックが必要です。現在では、同じプロンプトを複数のブラウザー タブの ChatGPT、Claude、Gemini、およびその他のモデルに貼り付け、違いを観察するという儀式が行われています。このアプリは、それを 1 つのコンポーザーに集約し、選択したすべてのモデルでプロンプトを並行して実行し、審議レイヤーを追加することで、結果が 3 つの意見を並べて表示するのではなく、推論が表示された総合的な推奨事項になります。しかし、モデルが増えるということは真実ではありません。評議会は単一モデルの盲点を減らしますが、すべての答えは依然として AI によって生成され、確実に間違っている可能性があります。評決は専門家のアドバイスではなく、自分の判断の材料として扱い、評決に基づいて下す決定は自分のものとしてください。
「この写真は正確にどこで撮影されたのか?」という単一の答えで解決すべきではない質問についての審議会。 — 添付の写真は、ライブ Web 検索とコード実行を備えた 3 つのモデルで、それぞれに手がかりがリストされています。このまさに評議会はキーレスデモとして出荷されます。
評議会は対話です。セッションには、着席した参加者の名簿、社会構造 (追加の役割が関係し、それらがどのように相互作用するか)、役割モデルごとの割り当て、および順番のスレッドが含まれます。各ターンでは、社会構造によって規定された役割順序に従ってユーザーのメッセージが伝えられます。ユーザーはフォローアップを求め続けることができます。
役割
何をするのか
参加者
テーブルに座っています。質問に答えます。私

n Trial 、他の人の回答 (匿名化) にも投票します。コンセンサス では、各ラウンドで回答を再回答し、匿名化された同僚の議論を踏まえて自身の立場を再検討します。名簿は参加者のリストです。
ホスト
ブラウザ側のオーケストレーション コード: 応答を収集し、投票のために匿名化し、次のロールにルーティングし、デバイス上でイベントを保持します。推論ではなく機械的です。UI には表現されません。
裁判官（オプション）
匿名化された回答と投票を読み取り、単一の総合的な決定を出力します。
メディエーター (オプション)
審判は複数ラウンドの参加者討論を行います。各ラウンドでは参加者が収束したかどうかを判断します。そうでない場合は、意見の相違を抽出（匿名化）して次のラウンドにシードします。合意に関する最終的なコンセンサスの概要、またはラウンドキャップにおける合意点と残りの競合を出力します。
各ロールは、プロンプト テンプレート + I/O コントラクトです。
役割の構成。各評議会は作成時に 1 つを選択します (構造は評議会の存続期間中固定されます。名簿とすべてのレシピ ノブは編集可能なままです)。
並行回答 — 参加者のみ。それぞれが独立して答えます。投票も総合もありません。生の回答を比較します。 (1 人の評議会は、通常の単一モデルのチャットにすぎません。)
裁判の評決 — 参加者が回答 → お互いの回答に匿名で投票 → 裁判官が最終的な評決を総合します。投票はリーダーボードと同意/不同意のビューに反映されます。審査員が発言する前に、モデルがどこで一致し、どこで分裂したかがわかり、ピア評価された最良の回答には★が付けられます。意図的に修正ステップなし: 裁判は権威による判断であり、裁判官は最初の回答案と投票に基づいて決定します。考え方を変えることはコンセンサスに属します。
コンセンサスディベート — 本物の複数ラウンドのディベート。ラウンドごとに、すべての参加者が再回答します。自分自身の以前の回答と他の参加者の順位を確認します。

ションは、ターン内で安定したラベルの下で匿名化されます (ターンをまたいで再シャッフルされるため、ブランドは推測できません)。次に、調停者が収束を判断します。収束 → 最終的な合意の概要を作成します。収束していない → 実際の意見の相違を抽出して、次のラウンドのシードにします。ラウンドには上限があります (デフォルトでは 3 つ、評議会ごとに構成可能)。キャップでは、調停者が合意点と残りの対立点を表面化します。最も高価なモード — およそ参加者 × ラウンドのモデル呼び出しです。
裁判の評決: 匿名化されたピア投票 (同意の指標とピア評価の勝者★付き)、その後、裁判官が決定的な証拠を詳述して判決を下します。
コンセンサスによる議論、その成果: 2 つのラウンドの後、評議会は収束します。そして、「このラウンドで何が変わったのか」には、誰が動いて誰が保持したかが示されます。偽りの調和はありません。ラウンドキャップでは、残りの競合が報告され、平滑化されません。
携帯電話でのフォーカス カルーセル読み取りフローや共有カードなど、モバイルは後付けではなく第一級です。
私たちのサーバーはあなたとモデルの間には存在しません。アプリは HTML/JS/CSS の静的バンドルです。ブラウザからのネットワーク トラフィックは、キーを使用して選択したモデル プロバイダーへの直接呼び出しです。さらに、以下に開示する公式サイトでの 1 回の Cookie なしのページビュー ping です。
Keys パネルには、この README と同じことが書かれており、貼り付けが行われる場所に、取引の一部に名前が付けられています。
カウンターは 1 つで、完全に公開されています。公式サイトは、当社が運営する自己ホスト型分析インスタンス ( stats.yesbrainer.ai ) にページビューを報告します。ページ内で Cookie やクライアント側の識別子、サードパーティのスクリプトが実行されることはありません。送信されるもの: ルート パターン (/council/:id や /settings/keys などの固定リストから — 実際の ID、完全な URL、コンテンツはペイロードの一部ではありません)、画面サイズ、言語、リファレ

いくつかの機能レベルのアクション (評議会の作成/削除、評決の共有、デモのオープン、プロバイダーに追加されたキー - キーではなくプロバイダー名 - Ollama の有効化、PWA のインストール、エクスポート/インポート、ストレージのリセット、永続性の付与/拒否) の数 - 機能の内容ではなく、機能が使用されたこと。どちらのリストも src/analytics/ にある閉じた TypeScript タイプです (100 数行、すぐに読めます)。キーと会話には移動するフィールドがありません。分析で確認できる範囲を広げることは、構成の反転ではなく、目に見えるコードの変更になります。デプロイ時エンドポイント (デフォルトでは任意のフォーク) を使用しないこのリポジトリのビルドでは、何も送信されません。自分のブラウザでレポートをオフにするには、localStorage で yesbrainer:analytics-disabled を 1 に設定します。
クロスデバイス同期はありません。新しいデバイスでアプリを開く → キーを貼り付けて、新たに開始します。 JSON エクスポート/インポート (設定 → ストレージ) は手動ブリッジです。オプションで独自のクラウド (Google Drive / Dropbox) 経由の同期も予定されています。
リセットフローはありません。デバイスを紛失すると、その上のデータも失われます。プロバイダーのダッシュボードはキーの信頼できる情報源です。そこでローテーションします。最後のエクスポートは会話の信頼できる情報源です。
ブラウザはデフォルトでストレージを削除できます。アプリは最初の意味のある書き込み時に永続ストレージを要求するため、ブラウザはデータをキャッシュではなくユーザー データとして扱います。明示的なアクション (cl

[切り捨てられた]

## Original Extract

🧠 A council of AI models for the decisions that aren't no-brainers — they answer in parallel, debate to consensus, or get judged to a verdict. Browser-only, open source, bring your own keys (BYOK), no backend. - trekhleb/yesbrainer

GitHub - trekhleb/yesbrainer: 🧠 A council of AI models for the decisions that aren't no-brainers — they answer in parallel, debate to consensus, or get judged to a verdict. Browser-only, open source, bring your own keys (BYOK), no backend. · GitHub
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
trekhleb
/
yesbrainer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github/ workflows .github/ workflows public public scripts scripts src src tests tests .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc CLAUDE.md CLAUDE.md DEVELOPMENT.md DEVELOPMENT.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md eslint.config.js eslint.config.js index.html index.html package-lock.json package-lock.json package.json package.json playwright.config.ts playwright.config.ts tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json tsconfig.tests.json tsconfig.tests.json tsconfig.unit.json tsconfig.unit.json vite.config.ts vite.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Yes-Brainer — a council of AI models for the decisions that aren’t no-brainers. One question fans out to several LLMs — and instead of juggling browser tabs, you get a deliberation in one place: independent answers, peer review, debate, a synthesized verdict, the disagreements visible.
Try it: yesbrainer.ai — the official deployment, built from this repository; treat any other domain as an unofficial copy. No sign-up: paste an API key and ask — or just open the recorded demo councils first, no key needed.
Three deliberation shapes — 🔀 Parallel answers (independent, no synthesis), ⚖️ Trial verdict (anonymous peer voting + a Judge), 🤝 Consensus debate (multi-round debate, a Mediator drives convergence).
Zero backend, zero accounts, zero data held by anyone but you. Keys, history, and preferences live in your browser; your prompts and keys travel only to the model providers you choose, directly.
Different model providers supported — Anthropic, OpenAI, Google, Groq, OpenRouter, opt-in Ollama — in one council.
No subscription, no paywall — you pay only your model providers, directly.
Mobile-friendly, installable as a PWA ; multi-turn sessions; sidebar of past councils.
A static bundle — GitHub Pages / any CDN; nothing to host, nothing to operate.
The front page: the three deliberation modes, recorded demo councils in the sidebar, and the get-started card — one screen, no sign-up. Click through to try it.
And creating a council is the whole setup: pick the deliberation mode, seat the models (capabilities and context size at a glance), choose who referees — or just hit "✨ Smartest available" and go.
Then the council convenes: one question fanned out to every seat, the answers side by side, the disagreement in plain view — the spread is the signal.
💡 Inspired by karpathy/llm-council — the orchestration concept and his prompt designs (Judge synthesis, peer evaluation) were the starting point. The difference is the shape, not the polish: the original (and most of its 80+ forks) runs one fixed answer → review → synthesize pass behind a local server you have to run, while Yes-Brainer is a zero-backend browser PWA with three deliberation structures — including a real multi-round debate refereed to convergence (consensus mode) — multi-turn follow-ups with on-device persistence, anonymized rubric voting with the agreement spread made visible, and a phone-first UI ( how it compares to everything else ).
Before pasting an API key into any site claiming to be Yes-Brainer, check that the address bar says yesbrainer.ai — this README and the app should agree on that domain, and a copied site can fake everything except its URL. Cap the blast radius too: use a dedicated key just for this app with a spending limit set in your provider's console, and revoke it there the moment anything looks off. The software is provided as-is, with no warranty (AGPL-3.0 §§15–16) — SECURITY.md spells out what the code protects and what it can't.
For complex or important questions, one model's confident answer isn't enough — you want several independent takes and a cross-check. Today the ritual is pasting the same prompt into ChatGPT, Claude, Gemini, and other models in multiple browser tabs and eyeballing the differences. This app collapses that into one composer, runs the prompt across all chosen models in parallel, and adds a deliberation layer so the result is a synthesized recommendation with the reasoning visible — not just three opinions side-by-side. More models isn't truth, though: a council reduces single-model blind spots, but every answer is still AI-generated and can be confidently wrong. Treat verdicts as raw material for your judgment — not professional advice — and own the decisions you make with them.
A council on a question a single answer shouldn't settle: "where exactly was this photo taken?" — the photo attached, three models with live web search and code execution, each listing its clues. This exact council ships as a keyless demo.
A council is a conversation. A session has a roster of seated participants, a social structure (which extra roles are involved and how they interact), per-role model assignments , and a thread of turns . Each turn runs the user's message through the role sequence dictated by the social structure; the user can keep asking follow-ups.
Role
What it does
Participant
Seated at the table; answers the question. In Trial , also votes on others' answers (anonymized). In Consensus , re-answers each round — reconsidering its own position in light of peers' anonymized arguments. The roster is a list of Participants.
Host
Browser-side orchestration code: collects responses, anonymizes for voting, routes to the next role, persists events on-device. Mechanical, not reasoning — not represented in the UI.
Judge (optional)
Reads anonymized answers + votes, emits a single synthesized decision.
Mediator (optional)
Referees a multi-round Participant debate: each round judges whether the Participants have converged; if not, distills their disagreements (anonymized) to seed the next round. Emits the final consensus summary on agreement, or points of agreement + remaining conflicts at the round cap.
Each role is a prompt template + an I/O contract.
Compositions of roles. Each council picks one at creation (the structure is fixed for the council's lifetime; the roster and every recipe knob stay editable).
Parallel answers — Participants only. Each one answers independently; no voting, no synthesis — you compare the raw answers. (A council of one is just a regular single-model chat.)
Trial verdict — Participants answer → vote anonymously on each other's answers → a Judge synthesizes one final verdict. The votes feed a leaderboard and an agreement/disagreement view — you see where the models agreed and where they split before the Judge speaks, and the peer-rated best answer wears a ★. No revision step , on purpose: Trial is judgment from authority — the Judge rules on first-draft answers plus votes. Mind-changing belongs to Consensus.
Consensus debate — a genuine multi-round debate. Every round, every Participant re-answers — seeing its own prior answer plus its peers' positions, anonymized under labels that stay stable within a turn (and reshuffle across turns, so brands can't be inferred). A Mediator then judges convergence: converged → it writes the final consensus summary; not converged → it distills the live disagreements to seed the next round. Rounds are capped (3 by default, configurable per council); at the cap the Mediator surfaces points of agreement + remaining conflicts. The most expensive mode — roughly Participants × rounds model calls.
Trial verdict: anonymized peer votes (with agreement indicators and the peer-rated winner ★), then the Judge rules — with the decisive evidence spelled out.
Consensus debate, the payoff: after two rounds the council converges — and "What changed this round" shows who moved and who held. No fake harmony: at the round cap, remaining conflicts are reported, not smoothed over.
Mobile is first-class, not an afterthought: the focus-carousel reading flow, and the share card, on a phone.
No server of ours sits between you and the models. The app is a static bundle of HTML/JS/CSS. The network traffic from your browser is calls direct to the model providers you chose, with your keys — plus one cookieless pageview ping on the official site, disclosed just below.
The Keys panel says the same thing this README does — and names your part of the deal, right where the pasting happens.
One counter, fully disclosed. The official site reports pageviews to a self-hosted analytics instance we operate ( stats.yesbrainer.ai ) — no cookies, no client-side identifiers, no third-party scripts running in the page. What's sent: the route pattern (from a fixed list like /council/:id or /settings/keys — actual ids, full URLs, and content aren't part of the payload), screen size, language, the referrer that brought you, and counts of a few feature-level actions (council created/deleted, verdict shared, demo opened, a key added for a provider — the provider name, not the key — Ollama enabled, PWA installed, export/import, storage resets, persistence granted/denied) — that a feature was used, not what was in it. Both lists are closed TypeScript types in src/analytics/ (a hundred-odd lines, quick to read): keys and conversations have no field to travel in, and widening what analytics may see would be a visible code change, not a config flip. A build of this repo without the deploy-time endpoint (any fork, by default) sends nothing. To switch reporting off in your own browser, set yesbrainer:analytics-disabled to 1 in localStorage.
No cross-device sync . Open the app on a new device → paste your keys, start fresh; JSON export/import (Settings → Storage) is the manual bridge. Optional sync via your own cloud (Google Drive / Dropbox) is planned.
No reset flow. If you lose a device you lose the data on it. The provider's dashboard is the source of truth for keys — rotate there; your last export is the source of truth for conversations.
Browsers can evict storage by default. The app requests persistent storage on your first meaningful write so the browser treats your data as user data, not cache; explicit actions (cl

[truncated]
