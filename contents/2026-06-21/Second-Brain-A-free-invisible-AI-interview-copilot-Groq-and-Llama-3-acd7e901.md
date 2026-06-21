---
source: "https://github.com/hi2brain/second-brain"
hn_url: "https://news.ycombinator.com/item?id=48615501"
title: "Second Brain – A free, invisible AI interview copilot (Groq and Llama 3)"
article_title: "GitHub - hi2brain/second-brain: A free, invisible AI interview copilot. Powered by Groq, Llama-3, and Whisper to provide real-time, context-aware answers during your job interviews. 🧠💼 · GitHub"
author: "hi2brain"
captured_at: "2026-06-21T04:35:33Z"
capture_tool: "hn-digest"
hn_id: 48615501
score: 2
comments: 0
posted_at: "2026-06-21T03:47:43Z"
tags:
  - hacker-news
  - translated
---

# Second Brain – A free, invisible AI interview copilot (Groq and Llama 3)

- HN: [48615501](https://news.ycombinator.com/item?id=48615501)
- Source: [github.com](https://github.com/hi2brain/second-brain)
- Score: 2
- Comments: 0
- Posted: 2026-06-21T03:47:43Z

## Translation

タイトル: Second Brain – 無料の目に見えない AI 面接副操縦士 (Groq および Llama 3)
記事のタイトル: GitHub - hi2brain/second-brain: 無料の目に見えない AI 面接副操縦士。 Groq、Llama-3、Whisper を利用して、就職面接中にリアルタイムでコンテキストを認識した回答を提供します。 🧠💼 · GitHub
説明: 無料の目に見えない AI 面接副操縦士。 Groq、Llama-3、Whisper を利用して、就職面接中にリアルタイムでコンテキストを認識した回答を提供します。 🧠💼 - hi2brain/セカンドブレイン

記事本文:
GitHub - hi2brain/セカンドブレイン: 無料の目に見えない AI 面接副操縦士。 Groq、Llama-3、Whisper を利用して、就職面接中にリアルタイムでコンテキストを認識した回答を提供します。 🧠💼 · GitHub
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
hi2brain
/
第二の脳
公共
通知

イオン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github/ workflows .github/ workflows electric electric public public src src .env.example .env.example .gitignore .gitignore ライセンス ライセンス README.es.md README.es.md README.fr.md README.fr.md README.it.md README.it.md README.md README.md README.pt.md README.pt.md empty-module.js empty-module.js eslint.config.js eslint.config.jsindex.htmlindex.htmlpackage-lock.json package-lock.json package.json package.jsonremove_bg.cjsremove_bg.cjs Remove_bg_sharp.cjs Remove_bg_sharp.cjs tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
目に見えない AI 面接アシスタント
英語 |ポルトガル語 |スペイン語 |フランセ |イタリアーノ
Second Brain は、世界中のプロフェッショナルが就職面接に合格できるよう支援するために構築された、完全無料の目に見えないデスクトップ アシスタントです。バックグラウンドで静かに実行され、インタビューを聞き、あなた自身の履歴書と職務内容に基づいて、コンテキストを認識した一人称の提案を生成します。
Groq の超高速 Llama-3 推論と Whisper-large-v3 を活用し、音声キャプチャからインテリジェントな提案までのプロセス全体がミリ秒以内に行われます。
リアルタイム音声文字起こし : Groq API 経由で Whisper-v3 を使用し、非常に高い精度と速度で面接官の音声を文字に起こします。
コンテキストに応じた提案 : 進行中の会話、提供された履歴書、職務内容を分析して、可能な限り最善の回答を提案します。
多言語サポート : 完全にローカライズされたインターフェースと AI プロンプト (英語、ブラジルポルトガル語、ヨーロッパポルトガル語)

ゲース、スペイン語、フランス語、イタリア語。
非表示モード: Electron で構築されているため、ビデオ通話に集中しているときに邪魔にならないように最小化できます。
プライバシーを第一に: 履歴書、職務経歴書、および API キーは、ローカル マシンの localStorage にのみ保存されます。
インストーラーをダウンロードします。
最新の Second Brain Setup .exe をリリース ページから入手してください。
無料の Groq API キーを取得します。
console.groq.com にアクセスして、無料の API キーを生成します。これは超高速 AI エンジンに必要です。
履歴書と職務内容を貼り付けます。
「インタビューを開始」をクリックすれば準備完了です。
ソースからビルドするか、貢献したい場合:
git clone https://github.com/2brain/second-brain.git
cd 第二の脳
# 依存関係をインストールする
npmインストール
# 開発サーバーを起動します
npm 実行開発
実稼働用に構築する
# Electron インストーラーをコンパイルしてビルドする
npm 実行距離
インストーラーは dist-installers フォルダー内に生成されます。
Second Brain はオープンソースであり、コミュニティにとっては無料です。このツールが仕事を見つけたり、面接に合格したりするのに役立った場合は、コーヒーをおごってください。
1. 本当に完全に無料ですか?
はい！このアプリケーションはオープンソースで完全に無料です。 API キーを取得するには、Groq で無料アカウントを作成するだけです。
2. 私のデータは安全ですか?私の履歴書を保存してもらえますか？
あなたのプライバシーは私たちの絶対的な優先事項です。データベースはありません。 API キーと履歴書は、コンピューターの localStorage に厳密に保存されます。当社のサーバーには何も送信されません。
3. 「履歴書」と「職務内容」フィールドに記入する最善の方法は何ですか?
履歴書 (PDF または LinkedIn) と職務内容の生のテキスト全体をコピーして貼り付けるだけです。フォーマットについて心配する必要はありません。 Second Brain AI は、メモリ内のデータを自動的に整理して、非常に正確で状況に応じた回答を提供します。
4. なしでシステムをテストするにはどうすればよいですか?

本当のインタビュー中ですか？
今すぐテストできます!アプリを開いて「聞く」をクリックし、模擬面接の YouTube 動画をスピーカーで再生するか、採用担当者のふりをして大声で質問するだけです。
5. システムが音声をキャプチャしていない場合はどうなりますか? (ヘッドセットに関する警告)
Second Brain は、システムのデフォルトのマイクを聞きます。ノイズキャンセリングヘッドセットを使用している場合、ヘッドセットのスピーカーから聞こえる採用担当者の声がマイクに捉えられない可能性があります。面接の前に必ずシステムをテストしてください。採用担当者の音声が聞き取れない場合は、コンピュータのオーディオ出力をスピーカーに変更するか、システムのオーディオ ルーティング設定を調整してください。
6. 採用担当者は AI を見たり聞いたりできますか?
いいえ、Second Brain は画面上で独立したウィンドウとして実行されます。マイクフィードに音声を挿入したり、画面を共有したりしません。面接官には全く見えません。
7. AI が応答するまでに数秒かかるのはなぜですか?
AI は、面接官が完全に考え終わるのを待ってから回答を生成します。スクリプトを提供する前に、質問全体を理解していることを確認するために、7 秒のブロックでリッスンします。
8. Mac または Linux で動作しますか?
現在、自動リリースでは Windows 用の .exe インストーラーが提供されています。ただし、開発者はこのリポジトリのクローンを簡単に作成し、npm run dist を実行して、macOS または Linux 用にネイティブに構築できます。
カスタム非営利ライセンスに基づいて配布されます。ソフトウェアの使用と研究は自由ですが、商用配布や販売は固く禁じられています。詳細については、「ライセンス」を参照してください。
無料の目に見えない AI 面接副操縦士。 Groq、Llama-3、Whisper を利用して、就職面接中にリアルタイムでコンテキストを認識した回答を提供します。 🧠💼
www.buymeacoffee.com/2brain
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0

フォーク
レポートリポジトリ
リリース
1
1.0.0
最新の
2026 年 6 月 21 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A free, invisible AI interview copilot. Powered by Groq, Llama-3, and Whisper to provide real-time, context-aware answers during your job interviews. 🧠💼 - hi2brain/second-brain

GitHub - hi2brain/second-brain: A free, invisible AI interview copilot. Powered by Groq, Llama-3, and Whisper to provide real-time, context-aware answers during your job interviews. 🧠💼 · GitHub
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
hi2brain
/
second-brain
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows electron electron public public src src .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.es.md README.es.md README.fr.md README.fr.md README.it.md README.it.md README.md README.md README.pt.md README.pt.md empty-module.js empty-module.js eslint.config.js eslint.config.js index.html index.html package-lock.json package-lock.json package.json package.json remove_bg.cjs remove_bg.cjs remove_bg_sharp.cjs remove_bg_sharp.cjs tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
Your Invisible AI Interview Assistant
English | Português | Español | Français | Italiano
Second Brain is a 100% free, invisible desktop assistant built to help professionals worldwide ace their job interviews. It runs quietly in the background, listening to the interview, and generates context-aware, first-person suggestions based on your own resume and the job description.
Powered by Groq's lightning-fast Llama-3 inference and Whisper-large-v3, the entire process—from voice capture to intelligent suggestion—happens in milliseconds.
Real-Time Voice Transcription : Uses Whisper-v3 via Groq API to transcribe the interviewer's speech with extremely high accuracy and speed.
Context-Aware Suggestions : Analyzes the ongoing conversation, your provided resume, and the job description to suggest the best possible answers.
Multi-Language Support : Fully localized interface and AI prompting for English, Brazilian Portuguese, European Portuguese, Spanish, French, and Italian.
Invisible Mode : Built with Electron, it can be minimized to stay out of the way while you focus on the video call.
Privacy First : Your resume, job description, and API keys are stored exclusively in your local machine's localStorage .
Download the Installer:
Get the latest Second Brain Setup .exe from our Releases page.
Get a Free Groq API Key:
Go to console.groq.com and generate a free API key. This is required for the ultra-fast AI engine.
Paste your Resume and the Job Description.
Click Start Interview and you're ready to go!
If you want to build from source or contribute:
git clone https://github.com/2brain/second-brain.git
cd second-brain
# Install dependencies
npm install
# Start the development server
npm run dev
Build for Production
# Compile and build the Electron installer
npm run dist
The installer will be generated inside the dist-installers folder.
Second Brain is open-source and free for the community. If this tool helped you land a job or nail an interview, consider buying us a coffee!
1. Is it really 100% free?
Yes! The application is open-source and entirely free. You only need to create a free account on Groq to get your API key.
2. Is my data safe? Do you save my resume?
Your privacy is our absolute priority. We do not have a database. Your API key and your resume are saved strictly in your computer's localStorage . Nothing is sent to our servers.
3. What is the best way to fill in the "Resume" and "Job Description" fields?
Simply copy and paste the entire raw text of your resume (PDF or LinkedIn) and the job description. Don't worry about formatting; the Second Brain AI will automatically organize your data in its memory to provide highly accurate and contextualized answers.
4. How can I test the system without being in a real interview?
You can test it right now! Just open the app, click "Listen", and play a YouTube video of a mock interview on your speakers, or simply pretend to be the recruiter and ask a question out loud.
5. What if the system is not capturing audio? (Headset Warning)
Second Brain listens to your system's default microphone. If you use noise-canceling headsets, your microphone might not capture the recruiter's voice coming from the headset speakers. Always test the system before an interview. If it's not capturing the recruiter's voice, change your computer's audio output to the speakers or adjust your system's audio routing settings.
6. Can the recruiter see or hear the AI?
No. Second Brain runs as an independent window on your screen. It does not inject audio into your microphone feed and does not share your screen. It is completely invisible to the interviewer.
7. Why is the AI taking a few seconds to answer?
The AI waits for the interviewer to finish a complete thought before generating an answer. It listens in 7-second blocks to ensure it understands the full question before giving you the script.
8. Does it work on Mac or Linux?
Currently, the automated releases provide a .exe installer for Windows. However, developers can easily clone this repository and run npm run dist to build it natively for macOS or Linux.
Distributed under a Custom Non-Commercial License. You are free to use and study the software, but commercial distribution or sale is strictly prohibited . See LICENSE for more information.
A free, invisible AI interview copilot. Powered by Groq, Llama-3, and Whisper to provide real-time, context-aware answers during your job interviews. 🧠💼
www.buymeacoffee.com/2brain
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
1.0.0
Latest
Jun 21, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
