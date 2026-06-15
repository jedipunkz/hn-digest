---
source: "https://github.com/themartiano/luz"
hn_url: "https://news.ycombinator.com/item?id=48538833"
title: "Show HN: I wrote a C++ ray tracer from scratch without AI"
article_title: "GitHub - themartiano/luz: C++ Path Tracer from scratch with zero third-party libraries. · GitHub"
author: "martiano"
captured_at: "2026-06-15T11:11:28Z"
capture_tool: "hn-digest"
hn_id: 48538833
score: 3
comments: 1
posted_at: "2026-06-15T09:34:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I wrote a C++ ray tracer from scratch without AI

- HN: [48538833](https://news.ycombinator.com/item?id=48538833)
- Source: [github.com](https://github.com/themartiano/luz)
- Score: 3
- Comments: 1
- Posted: 2026-06-15T09:34:10Z

## Translation

タイトル: Show HN: AI を使用せずに C++ レイ トレーサーを最初から作成しました
記事のタイトル: GitHub - themartiano/luz: サードパーティ ライブラリを使用せずにゼロから作成した C++ パス トレーサ。 · GitHub
説明: サードパーティ ライブラリを一切使用せずにゼロから作成した C++ パス トレーサ。 - テマルティアーノ/ルス

記事本文:
GitHub - themartiano/luz: サードパーティ ライブラリを使用せずにゼロから作成した C++ パス トレーサ。 · GitHub
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
テマルティアーノ
/
ルス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店T

ags ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
555 コミット 555 コミット .github/ ワークフロー .github/ ワークフロー ベンチマーク ベンチマーク docker/ ベンチマーク docker/ ベンチマーク ドキュメント ドキュメント サンプル/ シーン サンプル/ シーン インクルード/ luz インクルード/ luz スクリプト スクリプト src src テスト テスト ツール ツール .gitignore .gitignore CMakeLists.txt CMakeLists.txt COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス Makefile Makefile README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Luz は、サードパーティへの依存関係を一切持たずにゼロから開発された C++20 パス トレーサです。
モンテカルロ パス トレーシング、グローバル イルミネーション、BVH アクセラレーション、適応サンプリング、ノイズ除去、大気散乱、カスタム シーン ファイル、および Blender-to-Luz エクスポーターをサポートしています。
球、平面、長方形、三角形、立方体、ボリューム、OBJ メッシュ
ランバーシアン、金属、誘電体、放射、等方性マテリアル
エリア、ポイント、球、指向性ライト
CLI またはシーン ファイル経由で完全にカスタマイズ可能なレンダリング パラメータ
BVH アクセラレーション（ビン化された SAH 構造とニアファーストトラバーサルを備えたパックメッシュ BVH を含む）
散乱を伴う大気シミュレーション
被写界深度、アンチエイリアス、露出、コントラスト、トーンマッピング、ガンマ補正、ブルーム
レンダリング、ノイズ除去、ポストプロセス、スコアの内訳を含む決定論的なベンチマーク ハーネス
Python 3、オプションのツール/スクリプトのみ
作る
バンドルされたサンプル シーンをレンダリングします。
./Luz --ファイル例/シーン/blender_monkey.luz --サンプル 50 --解像度 300x300
デフォルトの出力は render.bmp です。シーン ファイルは、outputfilename=... を設定でき、CLI は共通のレンダリング設定をオーバーライドできます。
Luz には、レンダリング、ノイズ除去、ポストプロセス、および
総合スコアの比較。
ベンチマークを作成 BENCH_CPUS=1 BENCH_THREADS=1 > before.csv
BENCH_CPUS=1 BENCH_THREADS=1 > 後でベンチマークを作成します

r.csv
ベンチマーク比較を行う BEFORE=before.csv AFTER=after.csv
詳細については、 docs/benchmarks.md を参照してください。
CMake ビルドも利用できます。
cmake -S 。 -B ビルド
cmake --build ビルド
ctest --test-dir ビルド
プラットフォームのサポート
macOS および Linux では、Makefile がプライマリ パスです。 Windows では、CMake を使用します。
MSVC または MinGW ベースの Makefile ターゲット:
窓を作る
WSL は Linux ビルド環境としてもサポートされています。
リリース ビルドは、デフォルトでビルドを実行するマシンに合わせて調整されます。の
Makefile により、 -O3 、 -march=native を使用したネイティブ CPU チューニング、およびリンク時間が有効になります
-flto による最適化。また、高速浮動小数点モードも有効になります。
コンパイラ/プラットフォームがそれをサポートしています。 CMake は同じリリース インテントを使用します: -O3 、ネイティブ
CPU チューニング、およびプロシージャ間の最適化/LTO (サポートされている場合)。
これらのデフォルトではローカル レンダーが高速化されますが、バイナリは次のようにビルドされます。
-march=native は古い CPU や異なる CPU では実行できない可能性があり、LTO が公開する可能性があります。
ツールチェーン固有のリンカーの問題。不正な命令のクラッシュに遭遇した場合、
リンカー エラーが発生した場合、またはより移植性の高いバイナリが必要な場合は、積極的なオプションを無効にし、
クリーンなオブジェクトから再構築します。
きれいにする
NATIVE=0 LTO=0にする
CMake ビルドの場合は、最適化トグルをオフにして構成します。
cmake -S 。 -B ビルド -DLUZ_NATIVE_OPTIMIZATIONS=OFF -DLUZ_ENABLE_LTO=OFF
cmake --build build --clean-first
CLI
使用法: ./Luz [オプション]
-f, --file PATH .luz シーン ファイルをロードします。
-r, --resolution WxH レンダリング解像度をオーバーライドします
-s, --samples N ピクセルごとのサンプルをオーバーライドします
--adaptive [true|false] ピクセルごとの適応サンプリングを有効にする
--no-adaptive アダプティブ サンプリングを無効にする
--adaptive-min-samples N 適応停止前の最小サンプル数
--adaptive-threshold F 相対的な適応ノイズしきい値
--adaptive-check-interval N 適応型収束チェック間隔
-mlb, --maxLightBounces N 最大ライト バウンスをオーバーライドします。

--max-light-bounces N --maxLightBounces のエイリアス
-t, --threads N N ワーカースレッドでレンダリングします
--seed N シードのランダム サンプリング
--gamma true|false ガンマ補正の切り替え
-tm, --tonemapping true|false トーン マッピングを切り替えます
--bloom true|false ブルームの切り替え
--exposure EV 絞り時の露出補正
--contrast F 表示コントラスト乗数
--denoise [true|false] ノイズ除去されたコンパニオン レンダーを作成します
--no-denoise ノイズ除去を無効にする
-o, --output PATH レンダリング出力パスをオーバーライドします
--denoise-output PATH ノイズ除去された出力パスをオーバーライドします
--render-time renderTime.bmp を書き込みます
--benchmark 組み込みのベンチマーク シーンを実行します。
--benchmark-case 名前 ベンチマーク ケース: デフォルト、多オブジェクト、メッシュ BVH、拡散、ポストプロセス、雰囲気、ライト、エミッシブ ジオメトリ、プリミティブ マテリアル、ボリューム、オブジェクト メッシュ
適応サンプリング
--adaptive は、--samples をピクセルあたりの最大サンプルとして扱います。それぞれのピクセル
プログレッシブピクセルごとのサンプルシーケンスを使用し、少なくともレンダリング
--adaptive-min-samples 、輝度と RGB 信頼性を定期的にチェックします
間隔。非常に暗いピクセルは、停止する前に保守的な最小値を使用するため、
まれな光の寄与は、収束した黒と間違われる可能性が低くなります。
しきい値を低くすると、より詳細な情報が保持され、より多くの時間がかかります。最終的なレンダリングの場合は、開始します
最大サンプル数を高くし、次のような値で調整します。
./Luz --file exports/stormtroopers.luz --samples 4096 --adaptive --adaptive-min-samples 512 --adaptive-check-interval 64 --adaptive-threshold 0.005 --denoise
ノイズ除去
--denoise は、Luz の NFOR スタイルの機能バッファー デノイザーを有効にし、
別のコンパニオン画像。デフォルトでは、render.bmp は次のようになります。
render_denoized.bmp ; --denoise-output PATH を使用して正確なパスを選択します。
デノイザーには厳密な最小解像度やサンプル数はありませんが、十分な解像度が必要です
有用な色と特徴の統計を推定するための信号

スティックス。ピクセルごとに 1 つのサンプルは次のとおりです。
主にストレス テスト: ピクセルごとの分散推定がないため、ノイズ除去された
画像はほとんど変化していないように見えたり、間違った細部が滑らかになったりすることがあります。少なくとも
プレビューの場合はピクセルあたりのサンプルが少なく、ピクセルあたり約 16 以上のサンプルが望ましい
デノイザーの品質を判断するとき。非常に低い解像度も評価の対象となります
各ローカル フィルター ウィンドウが画像の大部分をカバーしすぎるため、誤解を招きます。
サンプルシーンは、examples/scenes/ にあります。メッシュ アセットは、assets/objects/ に存在します。シーン ファイルの形式は docs/scene-files.md に文書化されています。
.luz ファイル内のオブジェクト パスは、最初にシーン ファイルに対して相対的に解決され、次に現在の作業ディレクトリに対して相対的に解決され、次に、assets/objects/ の下で解決されます。つまり、examples/scenes/blender_monkey.luz は ../../assets/objects/blender_monkey.obj を参照でき、引き続きリポジトリ ルートから実行できます。
OBJ メッシュをオフセットしてシーン マテリアルを割り当てることもできます。
obj=メッシュ.obj,(x,y,z),マテリアル[
金属=(0.8,0.8,0.8),0.1
】
ブレンダーエクスポーター
Blender のシーンは、Blender の Python API を通じてエクスポートできます。
" /Applications/Blender.app/Contents/MacOS/Blender " -b scene.blend --python tools/blender_export_luz.py -- --output exports/scene.luz
./Luz --file エクスポート/scene.luz --threads 8
エクスポーターは、.luz ファイルと OBJ メッシュを書き込みます。使用法と現在の忠実度
制限については docs/blender-exporter.md に記載されています。
include/luz/ パブリックヘッダー
src/core/ 数学、ジオメトリ、マテリアル、画像、サンプリング コード
src/renderer/ レンダリングの実装
src/scene/ シーン モデルとシーン ヘルパー
src/io/ シーンファイル、OBJ、BMP、TIFFの読み込み/書き込み
src/cli/ コマンドラインのエントリポイントとフラグ
サンプル/シーン/ .luz シーン ファイルの例
例で使用されるアセット/オブジェクト/OBJ アセット
docs/images/ 圧縮されたショーケース画像
ツール/エクスポートおよびユーティリティ スクリプト
テスト/標準ライブラリ

y のみのテスト プログラム
ドッカー/ベンチマークコンテナ
ショーケース
Ray Tracing in One Weekend 書籍シリーズに心より感謝いたします。これは、特に AI が登場する前の時代であったため、Luz の開発の大部分においてインスピレーションと情報の大きな源でした。
BlendSwap 上の @ ScottGraham によるストームトルーパーのシーン。
BlendSwap の @geoffreymarchal による胸像。
サードパーティ ライブラリを使用せずにゼロから作成した C++ パス トレーサ。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.1.1
最新の
2026 年 6 月 12 日
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

C++ Path Tracer from scratch with zero third-party libraries. - themartiano/luz

GitHub - themartiano/luz: C++ Path Tracer from scratch with zero third-party libraries. · GitHub
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
themartiano
/
luz
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
555 Commits 555 Commits .github/ workflows .github/ workflows benchmarks benchmarks docker/ benchmark docker/ benchmark docs docs examples/ scenes examples/ scenes include/ luz include/ luz scripts scripts src src tests tests tools tools .gitignore .gitignore CMakeLists.txt CMakeLists.txt CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md View all files Repository files navigation
Luz is a C++20 Path Tracer developed from scratch with zero third-party dependencies.
It supports Monte Carlo path tracing, global illumination, BVH acceleration, adaptive sampling, denoising, atmospheric scattering, custom scene files, and a Blender-to-Luz exporter.
Spheres, planes, rectangles, triangles, cubes, volumes, and OBJ meshes
Lambertian, metal, dielectric, emissive, and isotropic materials
Area, point, sphere and directional lights
Fully customizable render parameters via CLI or scene file
BVH acceleration, including packed mesh BVHs with binned SAH construction and near-first traversal
Atmospheric simulation w/ scattering
Depth of field, antialiasing, exposure, contrast, tone mapping, gamma correction, and bloom
Deterministic benchmark harness with render, denoise, post-process, and score breakdowns
Python 3, only for optional tools/scripts
make
Render a bundled example scene:
./Luz --file examples/scenes/blender_monkey.luz --samples 50 --resolution 300x300
The default output is render.bmp . Scene files can set outputfilename=... , and the CLI can override common render settings.
Luz includes deterministic benchmarks for render, denoise, post-process, and
overall score comparisons.
make benchmark BENCH_CPUS=1 BENCH_THREADS=1 > before.csv
make benchmark BENCH_CPUS=1 BENCH_THREADS=1 > after.csv
make benchmark-compare BEFORE=before.csv AFTER=after.csv
For details, see docs/benchmarks.md .
A CMake build is also available:
cmake -S . -B build
cmake --build build
ctest --test-dir build
Platform Support
On macOS and Linux, the Makefile is the primary path. On Windows, use CMake with
MSVC or the MinGW-based Makefile target:
make windows
WSL is also supported as a Linux build environment.
Release builds are tuned for the machine doing the build by default. The
Makefile enables -O3 , native CPU tuning with -march=native , and link-time
optimization with -flto . It also enables a fast floating-point mode where the
compiler/platform supports it. CMake uses the same release intent: -O3 , native
CPU tuning, and interprocedural optimization/LTO when supported.
These defaults produce faster local renders, but binaries built with
-march=native may not run on older or different CPUs, and LTO can expose
toolchain-specific linker issues. If you hit an illegal-instruction crash,
linker error, or need a more portable binary, disable the aggressive options and
rebuild from clean objects:
make clean
make NATIVE=0 LTO=0
For CMake builds, configure with the optimization toggles off:
cmake -S . -B build -DLUZ_NATIVE_OPTIMIZATIONS=OFF -DLUZ_ENABLE_LTO=OFF
cmake --build build --clean-first
CLI
Usage: ./Luz [options]
-f, --file PATH Load a .luz scene file
-r, --resolution WxH Override render resolution
-s, --samples N Override samples per pixel
--adaptive [true|false] Enable adaptive per-pixel sampling
--no-adaptive Disable adaptive sampling
--adaptive-min-samples N Minimum samples before adaptive stopping
--adaptive-threshold F Relative adaptive noise threshold
--adaptive-check-interval N Adaptive convergence check interval
-mlb, --maxLightBounces N Override maximum light bounces
--max-light-bounces N Alias for --maxLightBounces
-t, --threads N Render with N worker threads
--seed N Seed random sampling
--gamma true|false Toggle gamma correction
-tm, --tonemapping true|false Toggle tone mapping
--bloom true|false Toggle bloom
--exposure EV Exposure compensation in stops
--contrast F Display contrast multiplier
--denoise [true|false] Write a denoised companion render
--no-denoise Disable denoising
-o, --output PATH Override render output path
--denoise-output PATH Override denoised output path
--render-times Write renderTime.bmp
--benchmark Run the built-in benchmark scene
--benchmark-case NAME Benchmark case: default, many-objects, mesh-bvh, diffuse, postprocess, atmosphere, lights, emissive-geometry, primitives-materials, volumes, obj-mesh
Adaptive Sampling
--adaptive treats --samples as the maximum samples per pixel. Each pixel
uses a progressive per-pixel sample sequence, renders at least
--adaptive-min-samples , then periodically checks luminance and RGB confidence
intervals. Very dark pixels use a conservative minimum before they can stop, so
rare light contributions are less likely to be mistaken for converged black.
Lower thresholds keep more detail and cost more time. For final renders, start
with a high max sample count and tune with values like:
./Luz --file exports/stormtroopers.luz --samples 4096 --adaptive --adaptive-min-samples 512 --adaptive-check-interval 64 --adaptive-threshold 0.005 --denoise
Denoising
--denoise enables Luz's NFOR-style feature-buffer denoiser and writes a
separate companion image. By default, render.bmp becomes
render_denoised.bmp ; use --denoise-output PATH to choose the exact path.
The denoiser has no hard minimum resolution or sample count, but it needs enough
signal to estimate useful color and feature statistics. One sample per pixel is
mainly a stress test: there is no per-pixel variance estimate, so the denoised
image can look almost unchanged or can smooth the wrong details. Use at least a
few samples per pixel for previews, and prefer roughly 16+ samples per pixel
when judging denoiser quality. Very low resolutions also make evaluation
misleading because each local filter window covers too much of the image.
Example scenes live in examples/scenes/ . Mesh assets live in assets/objects/ . The scene-file format is documented in docs/scene-files.md .
Object paths in .luz files are resolved relative to the scene file first, then relative to the current working directory, then under assets/objects/ . This means examples/scenes/blender_monkey.luz can reference ../../assets/objects/blender_monkey.obj and still run from the repository root.
OBJ meshes can also be offset and assigned a scene material:
obj=mesh.obj,(x,y,z),material[
metal=(0.8,0.8,0.8),0.1
]
Blender Exporter
Blender scenes can be exported through Blender's Python API:
" /Applications/Blender.app/Contents/MacOS/Blender " -b scene.blend --python tools/blender_export_luz.py -- --output exports/scene.luz
./Luz --file exports/scene.luz --threads 8
The exporter writes a .luz file plus OBJ meshes. Usage and current fidelity
limits are documented in docs/blender-exporter.md .
include/luz/ Public headers
src/core/ Math, geometry, materials, image, and sampling code
src/renderer/ Rendering implementation
src/scene/ Scene model and scene helpers
src/io/ Scene-file, OBJ, BMP, and TIFF loading/writing
src/cli/ Command-line entry point and flags
examples/scenes/ Example .luz scene files
assets/objects/ OBJ assets used by examples
docs/images/ Compressed showcase images
tools/ Export and utility scripts
tests/ Standard-library-only test program
docker/ Benchmark container
Showcase
Special thanks to the Ray Tracing in One Weekend book series. It was a great source of inspiration and information during a big part of the development of Luz, specially since those were times before AI.
Stormtrooper Scene by @ ScottGraham on BlendSwap .
Bust Statue by @ geoffreymarchal on BlendSwap .
C++ Path Tracer from scratch with zero third-party libraries.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.1.1
Latest
Jun 12, 2026
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
