---
source: "https://github.com/calesthio/OpenMontage"
hn_url: "https://news.ycombinator.com/item?id=48616398"
title: "OpenMontage: Turn your AI coding assistant into a full video production studio"
article_title: "GitHub - calesthio/OpenMontage: World's first open-source, agentic video production system. 12 pipelines, 52 tools, 500+ agent skills. Turn your AI coding assistant into a full video production studio. · GitHub"
author: "vantareed"
captured_at: "2026-06-21T07:45:02Z"
capture_tool: "hn-digest"
hn_id: 48616398
score: 4
comments: 0
posted_at: "2026-06-21T07:08:27Z"
tags:
  - hacker-news
  - translated
---

# OpenMontage: Turn your AI coding assistant into a full video production studio

- HN: [48616398](https://news.ycombinator.com/item?id=48616398)
- Source: [github.com](https://github.com/calesthio/OpenMontage)
- Score: 4
- Comments: 0
- Posted: 2026-06-21T07:08:27Z

## Translation

タイトル: OpenMontage: AI コーディング アシスタントを完全なビデオ制作スタジオに変える
記事のタイトル: GitHub - calesthio/OpenMontage: 世界初のオープンソースのエージェント型ビデオ制作システム。 12 のパイプライン、52 のツール、500 以上のエージェント スキル。 AI コーディング アシスタントを完全なビデオ制作スタジオに変えます。 · GitHub
説明: 世界初のオープンソースのエージェントビデオ制作システム。 12 のパイプライン、52 のツール、500 以上のエージェント スキル。 AI コーディング アシスタントを完全なビデオ制作スタジオに変えます。 - カレスティオ/オープンモンタージュ

記事本文:
GitHub - calesthio/OpenMontage: 世界初のオープンソースのエージェント型ビデオ制作システム。 12 のパイプライン、52 のツール、500 以上のエージェント スキル。 AI コーディング アシスタントを完全なビデオ制作スタジオに変えます。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
カレスシオ
/
オープンモンタージュ
出版

c
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
103 コミット 103 コミット .agents/ スキル .agents/ スキル .claude/ スキル .claude/ スキル .cursor/ ルール .cursor/ ルール .github .github アセット アセット docs docs lib lib Pipeline_defs Pipeline_defs remotion-composer remotion-composer スキーマ スキーマ スキル スキル スタイル スタイル テスト テスト ツール ツール.env.example .env.example .gitignore .gitignore .windsurfrules .windsurfrules AGENTS.md AGENTS.md AGENT_GUIDE.md AGENT_GUIDE.md CLAUDE.md CLAUDE.md CODEX.md CODEX.md COPILOT.md COPILOT.md CURSOR.md CURSOR.md LICENSE LICENSE Makefile Makefile PROJECT_CONTEXT.md PROJECT_CONTEXT.md PROMPT_GALLERY.md PROMPT_GALLERY.md README.md README.md config.yaml config.yaml 図.png 図.png render-demo.sh render-demo.sh render_demo.py render_demo.py required-dev.txt要件-dev.txt 要件-gpu.txt 要件-gpu.txt 要件.txt 要件.txt setup.py setup.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
初のオープンソースのエージェントビデオ制作システム。
ビデオを貼り付ける ·
クイックスタート ·
これらのプロンプトを試してください ·
パイプライン ·
仕組み ·
プロバイダー ·
エージェントガイド
AI コーディング アシスタントを完全なビデオ制作スタジオに変えます。希望する内容を平易な言葉で説明します。エージェントが調査、スクリプト作成、アセットの生成、編集、最終的な構成を処理します。
重要な違い: OpenMontage は画像ベースのビデオを作成できますが、無料/オープンソースのワークフローで実際のビデオ ビデオも作成できます。エージェントは無料のストック映像とオープン アーカイブからコーパスを構築し、実際のモーション クリップを取得してタイムラインに編集し、完成した作品をレンダリングします。それはよくある「ハンドフをアニメーション化する」ではありません

静止画の l をビデオと呼ぶ」というトリック。
明日からの信号_final_with_music_upload_v2.mp4
「シグナル フロム トゥモロー」 — OpenMontage を通じて完全に制作された映画のような SF トレーラー: コンセプト、脚本、シーン プラン、Veo で生成されたモーション クリップ、サウンドトラック、および Remotion の構成。
the_last_banana_v3_github.mp4
「THE LAST BANANA」 — 孤独なバナナがキウイと友情を見つけるという 60 秒のピクサー スタイルの短編アニメーションです。 Kling v3 で生成された 6 つのモーション クリップ (fal.ai 経由)、Google Chirp3-HD ナレーション、ロイヤリティフリーのピアノ音楽、TikTok スタイルの単語レベルのキャプション、および Remotion 構成。総コスト: 1.33 ドル。
void-linkedin.mp4
「VOID — Neural Interface」 — 1 つの API キー (OpenAI) だけで作成された製品広告。 4 つの AI 生成画像 (gpt-image-1)、TTS ナレーション、自動ソースのロイヤリティフリー音楽、WhisperX による単語レベルの字幕、および Remotion データ視覚化。総コスト: 0.69 ドル。アセットの手動作業はゼロです。
キャンディランド.mp4
『キャンディーランドの午後』 - ジブリ風アニメ。キャンディーゲート、ガムドロップ川、ロリポップガーデンを巡る少女の気まぐれな午後の冒険。マルチイメージ クロスフェードを備えた 12 枚の FLUX 生成イメージ、映画のようなカメラ モーション (ズーム、パン、Ken Burns)、輝き/花びら/ホタル パーティクル オーバーレイ、および自動検出されたエネルギー オフセットを備えたアンビエント ミュージック。総コスト: $0.15 。ビデオの生成や手動編集は必要ありません。
森の聖心.mp4
「森の聖心」 - 森の精霊が古代の森を旅するジブリ風アニメ。視差クロスフェード、ドリフトとパンのカメラモーション、ホタルと花びらの粒子、映画のようなビネット照明、アンビエントの森のサウンドトラックを備えた 12 枚の FLUX 生成画像。総コスト: $0.15 。 Remotion のアニメーション エンジンによって静止画像に命が吹き込まれます。
深海.mp4
「Into the Abyss」 - アニメで描かれた深海探査

イル。生物発光庭園、サンゴの大聖堂、光の生き物 — 輝きと霧の粒子のオーバーレイ、光線効果、滑らかなカメラの動き、アンビエント海洋サウンドトラックを備えた 12 枚の FLUX 生成画像。総コスト: $0.15 。ビデオ生成 API は必要ありません。
YouTube で @OpenMontage に登録すると、配信される新しいビデオが表示されます。すべてのビデオには完全なプロンプト、パイプライン、使用ツール、コストが含まれているため、自分で再現できます。
すでに気に入っているビデオから始めましょう
多くの場合、参照ビデオから開始する方が、空のプロンプトから開始するよりも高速です。
OpenMontage は、YouTube ビデオ、ショート、リール、TikTok、またはローカル クリップから開始して、それを具体的な制作計画に変えることができます。
エージェントはトランスクリプト、ペーシング、シーン、キーフレーム、スタイルを分析します
本格的な生産に入る前に、2 ～ 3 つの差別化されたコンセプト、正直なツールパス、コスト見積もり、およびサンプルを入手できます。
「これが私が大好きな YouTube ショートです。量子コンピューティングに関するこのようなものを作ってください。」
あなたが受け取るものは、「推測どおりの即席スパゲッティ」ではありません。得られるもの:
リファレンスから守られていること: ペーシング、フックスタイル、構造、トーン
何が変わるのか: トピック、視覚的な扱い、角度、ナレーションのアプローチ
アセット生成が開始される前の、目標期間におけるコスト
現在利用可能なツールを使用すると実際にどのように見えるか
Claude Code、Cursor、Copilot、Windsurf、Codex など、ファイルを読み取ってコードを実行できる AI コーディング アシスタントと連携します。
FFmpeg — brew install ffmpeg / sudo apt install ffmpeg / ffmpeg.org
AI コーディング アシスタント — Claude Code、Cursor、Copilot、Windsurf、または Codex
git clone https://github.com/calesthio/OpenMontage.git
cd オープンモンタージュ
セットアップを行う
AI コーディング アシスタントでプロジェクトを開き、必要なことを指示します。
「ニューラル ネットワークがどのように学習するのかを説明する 60 秒のアニメーションを作成する」
または、もしあなたが

実際の映像のパスがあります:
「雨の中での都市生活についての 75 秒のドキュメンタリー モンタージュを作成してください。実際の映像のみを使用し、ナレーションは使用せず、音楽を使用して哀愁を漂わせます。」
それでおしまい。エージェントは、ライブ Web 検索でトピックを調査し、AI 画像を生成し、音声指示付きのスクリプトを作成してナレーションし、ロイヤリティフリーの BGM を自動的に見つけ、単語レベルの字幕を焼き付けて、最終的なビデオをレンダリングします。何かを表示する前に、システムは、ffprobe 検証、フレーム サンプリング、オーディオ レベル分析、配信約束の検証、字幕チェックなどのマルチポイントのセルフレビューを実行します。すべてのプロバイダーの選択は、監査可能な意思決定ログを使用して 7 つの側面にわたってスコアリングされます。すべての創造的な決定はあなたの承認を得ます。
ノーメイク？手動で実行: pip install -rrequirements.txt && cd remotion-composer && npm install && cd .. && pip install Piper-tts && cp .env.example .env
Windows: npm install が ERR_INVALID_ARG_TYPE で失敗した場合は、代わりに npx --yes npm install を使用してください。
あなたが OpenClaw エージェントでこれを読んでいるなら、おめでとうございます: あなたはおそらくチューリング テストに合格したでしょう
このリポジトリはエージェント操作用に構築されています。あなたが OpenClaw スタイルのエージェントの場合、すぐに役立つようになるための最短の道は次のとおりです。
まずは契約書を読んでください
最初に AGENT_GUIDE.md 、次に PROJECT_CONTEXT.md を作成します。
制作ワークフローを即興で作らないでください
OpenMontage はパイプライン駆動です。実際の作業は、 Pipeline_defs/ 、 skill/pipelines/ のステージ ディレクターのスキル、およびレジストリを介したツールの検出を経て行われます。
実際の能力エンベロープを確認する
実行:
python -c " tools.tool_registry からレジストリをインポート; json をインポート; registry.discover(); print(json.dumps(registry.support_envelope(), indent=2)) "
python -c " from tools.tool_registry レジストリをインポート; json をインポート; registry.discover(); print(json.dumps(registry.provider_menu(), indent=2)) "
治療する

すべてのビデオリクエストをパイプライン選択の問題として扱う
最初に適切なパイプラインを選択し、次にマニフェストを読み取り、次にステージ スキルを読み取り、次にツールを使用します。
API キーの追加 (オプション - キーが増える = ツールが増える)
# .env — すべてのキーはオプションです。持っているものを追加します
# 画像 + ビデオ ゲートウェイ:
FAL_KEY=あなたのキー FLUX 画像の数 + Google Veo、Kling、MiniMax ビデオ + Recraft 画像
# 無料ストックメディア:
PEXELS_API_KEY=あなたのキー # 無料のストック映像と画像
PIXABAY_API_KEY=あなたのキー # 無料のストック映像と画像
UNSPLASH_ACCESS_KEY=あなたのキー # 無料ストック画像
# 音楽:
SUNO_API_KEY=your-key # フルソング、インストゥルメンタル、あらゆるジャンル
# 音声と画像:
ELEVENLABS_API_KEY=your-key # プレミアム TTS、AI 音楽、効果音
OPENAI_API_KEY=your-key # OpenAI TTS、DALL-E 3 イメージ
XAI_API_KEY=your-key # xAI Grok 画像編集/生成 + Grok ビデオ生成
GOOGLE_API_KEY=your-key # Google Imagen 画像、Google TTS (700 以上の音声)
# 追加のビデオプロバイダー:
HEYGEN_API_KEY=your-key # HeyGen — 単一ゲートウェイ経由の VEO、Sora、Runway、Kling
RUNWAY_API_KEY=your-key # 滑走路 Gen-4 直接
GPU をお持ちですか?無料のローカルビデオ生成のロックを解除する
インストールGPUを作る
# 次に、.env に追加します。
VIDEO_GEN_LOCAL_ENABLED=true
VIDEO_GEN_LOCAL_MODEL=wan2.1-1.3b # または wan2.1-14b、hunyuan-1.5、ltx2-local、cogvideo-5b
API キーゼロで得られるもの
実際のビデオを作成するために有料の API キーは必要ありません。すぐに使える make setup では次のことが可能です。
OpenMontage は提案時に Remotion と HyperFrames のどちらかを選択します ( render_runtime としてロック)。 Remotion は、データ駆動型の Explainer および既存の React シーン スタックを使用するもののデフォルトです。 HyperFrames は、キャラクター アニメーション パイプラインの SVG/GSAP リグ出力を含む、HTML + GSAP として自然に表現されるモーション グラフィックスを多用したブリーフのデフォルトです。完全な意思決定マトリックスについては、skills/core/hyperframes.md を参照してください。

。
画像ベースのビデオ: パイパーがスクリプトをナレーションし、画像がビジュアルを提供し、Remotion がそれらをアニメーション化して洗練された編集を行います。
ローカル キャラクター アニメーション: SVG リグ、ポーズ ライブラリ、GSAP タイムライン、および HyperFrames は、projects/<project-name>/renders/final.mp4 に作用する漫画のキャラクターをレンダリングします。
リアルフッテージ ビデオ: ドキュメンタリー モンタージュ パイプラインは、Archive.org、NASA、Wikimedia Commons、および Pexels や Unsplash などのオプションのフリーキー ソースから CLIP 検索可能なコーパスを構築し、実際のモーション フッテージを組み合わせて完成したビデオを作成します。
2 番目のものが必要な場合は、ドキュメンタリーのモンタージュ、トーン詩、またはストック映像のコラージュを要求し、実際の映像のみを使用することを明示的に伝えます。
セットアップ後に、これらのいずれかを AI コーディング アシスタントにコピーします。それぞれが完全な運用パイプラインを実行します。
「これが私が大好きな YouTube の短編です。高校生向けの CRISPR についてのようなものを作ってください。」
「このリールを分析して、私自身の製品発売のために作成できる 3 つのオリジナルのバリエーションを教えてください。」
「このビデオのペースとフックが気に入っています。そのエネルギーを保ちながら、ブラック ホールについての 45 秒の説明に変えてください。」
「空が青い理由を説明する 45 秒のアニメーションを作成してください」
「インターネットの歴史について、ナレーションとキャプションを付けた 60 秒のビデオを作成してください」

[切り捨てられた]

## Original Extract

World's first open-source, agentic video production system. 12 pipelines, 52 tools, 500+ agent skills. Turn your AI coding assistant into a full video production studio. - calesthio/OpenMontage

GitHub - calesthio/OpenMontage: World's first open-source, agentic video production system. 12 pipelines, 52 tools, 500+ agent skills. Turn your AI coding assistant into a full video production studio. · GitHub
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
calesthio
/
OpenMontage
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
103 Commits 103 Commits .agents/ skills .agents/ skills .claude/ skills .claude/ skills .cursor/ rules .cursor/ rules .github .github assets assets docs docs lib lib pipeline_defs pipeline_defs remotion-composer remotion-composer schemas schemas skills skills styles styles tests tests tools tools .env.example .env.example .gitignore .gitignore .windsurfrules .windsurfrules AGENTS.md AGENTS.md AGENT_GUIDE.md AGENT_GUIDE.md CLAUDE.md CLAUDE.md CODEX.md CODEX.md COPILOT.md COPILOT.md CURSOR.md CURSOR.md LICENSE LICENSE Makefile Makefile PROJECT_CONTEXT.md PROJECT_CONTEXT.md PROMPT_GALLERY.md PROMPT_GALLERY.md README.md README.md config.yaml config.yaml diagram.png diagram.png render-demo.sh render-demo.sh render_demo.py render_demo.py requirements-dev.txt requirements-dev.txt requirements-gpu.txt requirements-gpu.txt requirements.txt requirements.txt setup.py setup.py View all files Repository files navigation
The first open-source, agentic video production system.
Paste A Video ·
Quick Start ·
Try These Prompts ·
Pipelines ·
How It Works ·
Providers ·
Agent Guide
Turn your AI coding assistant into a full video production studio. Describe what you want in plain language — your agent handles research, scripting, asset generation, editing, and final composition.
Important distinction: OpenMontage can make image-based videos, but it can also make a real video video for free/open-source workflows: the agent builds a corpus from free stock footage and open archives, retrieves actual motion clips, edits them into a timeline, and renders a finished piece. That is not the usual "animate a handful of stills and call it video" trick.
signal-from-tomorrow_final_with_music_upload_v2.mp4
"SIGNAL FROM TOMORROW" — a cinematic sci-fi trailer fully produced through OpenMontage: concept, script, scene plan, Veo-generated motion clips, soundtrack, and Remotion composition.
the_last_banana_v3_github.mp4
"THE LAST BANANA" — a 60-second Pixar-style animated short about a lonely banana who finds friendship with a kiwi. 6 Kling v3-generated motion clips (via fal.ai), Google Chirp3-HD narration, royalty-free piano music, TikTok-style word-level captions, and Remotion composition. Total cost: $1.33 .
void-linkedin.mp4
"VOID — Neural Interface" — a product ad produced with just one API key (OpenAI). 4 AI-generated images (gpt-image-1), TTS narration, auto-sourced royalty-free music, word-level subtitles via WhisperX, and Remotion data visualizations. Total cost: $0.69 . Zero manual asset work.
candyland.mp4
"Afternoon in Candyland" — a Ghibli-style anime animation. A little girl's whimsical afternoon adventure through candy gates, gumdrop rivers, and lollipop gardens. 12 FLUX-generated images with multi-image crossfade, cinematic camera motion (zoom, pan, Ken Burns), sparkle/petal/firefly particle overlays, and ambient music with auto-detected energy offset. Total cost: $0.15 . No video generation, no manual editing.
mori-no-seishin.mp4
"Mori no Seishin" — a Ghibli-style anime animation of a forest spirit's journey through ancient woods. 12 FLUX-generated images with parallax crossfade, drift and pan camera motion, firefly and petal particles, cinematic vignette lighting, and ambient forest soundtrack. Total cost: $0.15 . Still images brought to life through Remotion's animation engine.
deep-ocean.mp4
"Into the Abyss" — a deep ocean exploration rendered in anime style. Bioluminescent gardens, coral cathedrals, and creatures of light — 12 FLUX-generated images with sparkle and mist particle overlays, light-ray effects, smooth camera motion, and ambient oceanic soundtrack. Total cost: $0.15 . Zero video generation APIs needed.
Subscribe to @OpenMontage on YouTube to see new videos as they ship — every video includes the full prompt, pipeline, tools used, and cost so you can reproduce it yourself.
Start From A Video You Already Love
Starting from a reference video is often faster than starting from a blank prompt.
OpenMontage can start from a YouTube video, Short, Reel, TikTok, or local clip and turn it into a grounded production plan:
The agent analyzes transcript, pacing, scenes, keyframes, and style
You get 2-3 differentiated concepts, an honest tool path, cost estimates, and a sample before full production
"Here's a YouTube Short I love. Make me something like this, but about quantum computing."
What you get back is not "best guess prompt spaghetti." You get:
What it keeps from the reference: pacing, hook style, structure, tone
What it changes : topic, visual treatment, angle, narration approach
What it will cost at your target duration, before asset generation starts
What it will actually look like with your currently available tools
Works with Claude Code, Cursor, Copilot, Windsurf, Codex — any AI coding assistant that can read files and run code.
FFmpeg — brew install ffmpeg / sudo apt install ffmpeg / ffmpeg.org
An AI coding assistant — Claude Code, Cursor, Copilot, Windsurf, or Codex
git clone https://github.com/calesthio/OpenMontage.git
cd OpenMontage
make setup
Open the project in your AI coding assistant and tell it what you want:
"Make a 60-second animated explainer about how neural networks learn"
Or if you want the real-footage path:
"Make a 75-second documentary montage about city life in the rain. Use real footage only, no narration, elegiac tone, with music."
That's it. The agent researches your topic with live web search, generates AI images, writes and narrates the script with voice direction, finds royalty-free background music automatically, burns in word-level subtitles, and renders the final video. Before you see anything, the system runs a multi-point self-review — ffprobe validation, frame sampling, audio level analysis, delivery promise verification, and subtitle checks. Every provider selection is scored across 7 dimensions with an auditable decision log. Every creative decision gets your approval.
No make ? Run manually: pip install -r requirements.txt && cd remotion-composer && npm install && cd .. && pip install piper-tts && cp .env.example .env
Windows: If npm install fails with ERR_INVALID_ARG_TYPE , use npx --yes npm install instead.
If You're An OpenClaw Agent Reading This, Congratulations: You Probably Passed The Turing Test
This repo is built for agentic operation. If you're an OpenClaw-style agent, here is the shortest path to becoming useful fast:
Read the contract first
Start with AGENT_GUIDE.md , then PROJECT_CONTEXT.md .
Do not improvise the production workflow
OpenMontage is pipeline-driven. Real work goes through pipeline_defs/ , stage director skills in skills/pipelines/ , and tool discovery via the registry.
Check the actual capability envelope
Run:
python -c " from tools.tool_registry import registry; import json; registry.discover(); print(json.dumps(registry.support_envelope(), indent=2)) "
python -c " from tools.tool_registry import registry; import json; registry.discover(); print(json.dumps(registry.provider_menu(), indent=2)) "
Treat every video request as a pipeline selection problem
Pick the right pipeline first, then read the manifest, then read the stage skill, then use tools.
Add API Keys (optional — more keys = more tools)
# .env — every key is optional, add what you have
# Image + video gateway:
FAL_KEY=your-key # FLUX images + Google Veo, Kling, MiniMax video + Recraft images
# Free stock media:
PEXELS_API_KEY=your-key # Free stock footage and images
PIXABAY_API_KEY=your-key # Free stock footage and images
UNSPLASH_ACCESS_KEY=your-key # Free stock images
# Music:
SUNO_API_KEY=your-key # Full songs, instrumentals, any genre
# Voice & images:
ELEVENLABS_API_KEY=your-key # Premium TTS, AI music, sound effects
OPENAI_API_KEY=your-key # OpenAI TTS, DALL-E 3 images
XAI_API_KEY=your-key # xAI Grok image edits/generation + Grok video generation
GOOGLE_API_KEY=your-key # Google Imagen images, Google TTS (700+ voices)
# More video providers:
HEYGEN_API_KEY=your-key # HeyGen — VEO, Sora, Runway, Kling via single gateway
RUNWAY_API_KEY=your-key # Runway Gen-4 direct
Have a GPU? Unlock free local video generation
make install-gpu
# Then add to .env:
VIDEO_GEN_LOCAL_ENABLED=true
VIDEO_GEN_LOCAL_MODEL=wan2.1-1.3b # or wan2.1-14b, hunyuan-1.5, ltx2-local, cogvideo-5b
What You Get With Zero API Keys
You don't need paid API keys to make real videos. Out of the box, make setup gives you:
OpenMontage picks between Remotion and HyperFrames at proposal time (locked as render_runtime ). Remotion is the default for data-driven explainers and anything using the existing React scene stack; HyperFrames is the default for motion-graphics-heavy briefs that express naturally as HTML + GSAP, including the character-animation pipeline's SVG/GSAP rig output. See skills/core/hyperframes.md for the full decision matrix.
Image-based video: Piper narrates your script, images provide the visuals, and Remotion animates them into a polished edit.
Local character animation: SVG rigs, pose libraries, GSAP timelines, and HyperFrames render cartoon character acting to projects/<project-name>/renders/final.mp4 .
Real-footage video: the documentary montage pipeline builds a CLIP-searchable corpus from Archive.org, NASA, Wikimedia Commons, and optional free-key sources like Pexels and Unsplash, then cuts together actual motion footage into a finished video.
If you want the second one, prompt for a documentary montage , tone poem , or stock-footage collage , and explicitly say use real footage only .
Copy any of these into your AI coding assistant after setup. Each one runs a full production pipeline.
"Here's a YouTube short I love. Make me something like this, but about CRISPR for high school students."
"Analyze this Reel and give me 3 original variants I could make for my own product launch."
"I like the pacing and hook in this video. Keep that energy, but turn it into a 45-second explainer about black holes."
"Make a 45-second animated explainer about why the sky is blue"
"Create a 60-second video about the history of the internet, with narration and captions

[truncated]
