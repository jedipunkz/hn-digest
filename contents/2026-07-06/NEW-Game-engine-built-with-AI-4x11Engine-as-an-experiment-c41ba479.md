---
source: "https://github.com/rwusmm-dc/4x11Engine"
hn_url: "https://news.ycombinator.com/item?id=48807757"
title: "(NEW) Game engine built with AI (4x11Engine) as an experiment"
article_title: "GitHub - rwusmm-dc/4x11Engine: 4x11 Engine is a game engine made using only AI as a experiment, supports DX10/11 · GitHub"
author: "rwusmm"
captured_at: "2026-07-06T17:47:03Z"
capture_tool: "hn-digest"
hn_id: 48807757
score: 1
comments: 1
posted_at: "2026-07-06T17:24:23Z"
tags:
  - hacker-news
  - translated
---

# (NEW) Game engine built with AI (4x11Engine) as an experiment

- HN: [48807757](https://news.ycombinator.com/item?id=48807757)
- Source: [github.com](https://github.com/rwusmm-dc/4x11Engine)
- Score: 1
- Comments: 1
- Posted: 2026-07-06T17:24:23Z

## Translation

タイトル：（NEW）AIを使ったゲームエンジン（4x11Engine）を実験的に構築
記事タイトル: GitHub - rwusmm-dc/4x11Engine: 4x11 Engineは実験的にAIだけを使って作られたゲームエンジン、DX10/11をサポート · GitHub
説明: 4x11 Engine は実験的に AI のみを使用して作成されたゲーム エンジンで、DX10/11 をサポート - rwusmm-dc/4x11Engine

記事本文:
GitHub - rwusmm-dc/4x11Engine: 4x11 Engineは実験的にAIのみを使用して作られたゲームエンジンで、DX10/11をサポート · GitHub
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
rwusmm-dc
/
4x11エンジン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

n 個のオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット 例/ コントローラー-1.0 例/ コントローラー-1.0 LuaJIT LuaJIT imgui imgui スクリプト スクリプト src src ツール ツール zstd zstd .gitignore .gitignore ライセンス ライセンス README.md README.md build.bat build.bat build_game.bat build_game.bat imgui.ini imgui.ini libzstd.dll libzstd.dll すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI 支援ソフトウェア開発の実験として開発された、Direct3D 10 および 11 をサポートする軽量のゲーム エンジンとエディター。
4x11Engine は、次の機能を備えた完全なゲーム開発フレームワークです。
依存関係が最小限の C++ コア
ランタイム バックエンド選択による Direct3D 10/11 のサポート
シーン階層、プロパティインスペクター、ギズモ操作を備えた組み込みエディター
弾丸物理学 4xPhys (バージョン v0.2)
LuaJIT (4xLang) によるスクリプトのサポート
ダイナミック スカイ 独自のスカイ シェーダー (実験的) を作成およびカスタマイズできます。
エンティティ コンポーネント システム (ECS) アーキテクチャ
4xエンジン/ソース/
§── コア/
│ §── ComPtr.h - スマート COM ポインタ ラッパー
│ §── CullingSystem.* - 錐台とオクルージョンのカリング
│ §── FPSCamera.* - 一人称視点のカメラ コントローラー
│ §── Window.* - Win32 ウィンドウ管理
│ └─ Project.* - プロジェクト管理
§── ecs/
│ §── ECS.h - エンティティコンポーネントシステム
│ └── ECS.cpp - エンティティのシリアル化とシーン管理
§── d3d10/ - Direct3D 10 実装
§── d3d11/ - Direct3D 11 実装
│ └── skybox.* - 太陽シミュレーションを備えたダイナミックな空のシステム
§── phy/ - 物理システム (AABB 衝突、インパルス分解能)
§── script/ - 4xLang Lua スクリプト エンジン
§── io/ - アーカイブおよびファイル I/O
§── ui/ - ImGui ベースのエディター インターフェイス
│ §── オーバーレイ.* - 舞

n エディタウィンドウとツール
│ §── CodeEditor.* - 構文を強調表示するコード エディター
│ §── Gizmo.* - 3D 変換ギズモ
│ └── ProjectManagerUI.* - プロジェクト管理 UI
└── gizmo/ - 変換ギズモの実装
組み立て説明書
C++ をサポートする MinGW-w64 または Visual Studio をインストールする
DirectX SDK または Windows SDK がインストールされていることを確認してください
DirectXMath を使用する場合は、TDM-GCC に DirectXMath がインストールされていることを確認してください ( https://github.com/rwusmm-dc/directx-examples/releases/download/dxmath/tdm-gcc-directxmathInstaller.exe )。
提供されたスクリプトを使用してビルドします。
game.exe - スタンドアロン ゲーム ランチャー
main.exeを起動してプロジェクトマネージャーを開きます。
新しいプロジェクトを作成するか、既存のプロジェクトを開きます
プロジェクトはDocuments/4xEngine/Projects/に保存されます。
階層ウィンドウ : 右クリックして、プリミティブ (立方体、球、カプセル、平面、三角形、八角形、雪だるま)、ライト、カメラ、およびスクリプトを作成します。エンティティをドラッグして親を変更します。
プロパティ ウィンドウ : 変換、物理学、色、ライト設定、およびスクリプト パスを編集します。
ギズモ : 移動 (T)、回転 (R)、スケール (S) モードを使用してエンティティを変換します。ローカル (L) 空間とワールド (W) 空間を切り替えます。
パフォーマンス ウィンドウ : FPS、描画コール、CPU/GPU 使用率、およびカリング統計を監視します。
次の API を使用して .4xs スクリプト ファイルを作成します。
EntityService.findByName(name) - エンティティを名前で検索します
EntityService.findById(id) - ID でエンティティを検索
EntityService.getAll() - すべてのエンティティを取得する
EntityService.create(name) - 新しいエンティティを作成する
EntityService.remove(entity) - エンティティを削除します
:getPosition() / :setPosition({x,y,z})
:getRotation() / :setRotation({x,y,z})
:getScale() / :setScale({x,y,z})
:getVelocity() / :setVelocity({x,y,z})
:isLight() 、:isCamera() 、:isStatic()
LightService.createPoint(名前)
LightService.createDirectional(名前)
LightService.setColor(entity, {r,g,b})

LightService.setIntensity(エンティティ, 値)
LightService.setRange(エンティティ, 値)
CameraService.getActive() - アクティブなカメラを取得します
CameraService.setActive(entity) - アクティブなカメラを設定します
CameraService.create(name) - カメラの作成
WorldService.setSkyColor({r,g,b})
WorldService.setSunBrightness(値)
WorldService.getSunBrightness()
Input.key(name) - キーの状態を確認する
Input.mouseX() 、Input.mouseY()
Input.mouseDelta() - dx、dy を返します。
SyncService.getDelta() - フレームデルタ時間
SyncService.getDiagnostics() - エンジンの統計情報
シーン内に少なくとも 1 つのカメラ エンティティが存在することを確認してください
エディター メニューから [ファイル] > [エクスポート] を選択します。
エクスポーターは、すべてのシーン データがパッケージ化されたスタンドアロンの game.exe を作成します。
ショートカット
アクション
Ctrl+S
シーンを保存する
Ctrl+Z
元に戻す
Ctrl+Y
やり直し
削除
選択したエンティティを削除します
W/A/S/D
カメラを移動する
宇宙
カメラを上に移動
×
カメラを下に移動します
右クリック
カメラ制御用のマウスのロック/ロック解除
クレジット:
https://github.com/ocornut/imgui
https://github.com/luajit/luajit
カリング : オプションの D3D11 オクルージョン クエリを使用した錐台カリング
物理学 : 衝撃分解能、摩擦、復元を備えた AABB ベースの衝突
シリアル化 : GAF 形式はスクリプトを含む完全なエンティティ状態をサポートします
照明 : 時刻シミュレーション付きのダイナミックな太陽、減衰付きのポイント ライト
スクリプト: 各スクリプトは分離のために独自の Lua 状態で実行されます。
このプロジェクトは教育および実験目的で現状のまま提供されており、まだ実際に使用する準備はできていませんが、デモには使用できます。自己責任で使用してください。プロジェクトは MIT ライセンスに基づいてアルファステートライセンスとして提供されています。
4x11 EngineはAIだけを使って実験的に作られたゲームエンジンで、DX10/11に対応しています
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フット

ナビゲーション
私の個人情報を共有しないでください

## Original Extract

4x11 Engine is a game engine made using only AI as a experiment, supports DX10/11 - rwusmm-dc/4x11Engine

GitHub - rwusmm-dc/4x11Engine: 4x11 Engine is a game engine made using only AI as a experiment, supports DX10/11 · GitHub
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
rwusmm-dc
/
4x11Engine
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits Examples/ Controller-1.0 Examples/ Controller-1.0 LuaJIT LuaJIT imgui imgui scripts scripts src src tools tools zstd zstd .gitignore .gitignore LICENSE LICENSE README.md README.md build.bat build.bat build_game.bat build_game.bat imgui.ini imgui.ini libzstd.dll libzstd.dll View all files Repository files navigation
A lightweight game engine and editor supporting Direct3D 10 and 11, developed as an experiment in AI-assisted software development.
4x11Engine is a complete game development framework featuring:
C++ core with minimal dependencies
Direct3D 10/11 support with runtime backend selection
Built-in editor with scene hierarchy, property inspector, and gizmo manipulation
Bullet physics 4xPhys (Version v0.2)
Scripting support via LuaJIT (4xLang)
Dynamic sky Allows creating and customizing your own sky shader (experimental).
Entity Component System (ECS) architecture
4xEngine/src/
├── core/
│ ├── ComPtr.h - Smart COM pointer wrapper
│ ├── CullingSystem.* - Frustum and occlusion culling
│ ├── FPSCamera.* - First-person camera controller
│ ├── Window.* - Win32 window management
│ └── Project.* - Project management
├── ecs/
│ ├── ECS.h - Entity Component System
│ └── ECS.cpp - Entity serialization and scene management
├── d3d10/ - Direct3D 10 implementation
├── d3d11/ - Direct3D 11 implementation
│ └── skybox.* - Dynamic sky system with sun simulation
├── phy/ - Physics system (AABB collision, impulse resolution)
├── script/ - 4xLang Lua scripting engine
├── io/ - Archive and file I/O
├── ui/ - ImGui-based editor interface
│ ├── Overlay.* - Main editor windows and tools
│ ├── CodeEditor.* - Syntax-highlighting code editor
│ ├── Gizmo.* - 3D transformation gizmo
│ └── ProjectManagerUI.* - Project management UI
└── gizmo/ - Transformation gizmo implementation
Build Instructions
Install MinGW-w64 or Visual Studio with C++ support
Ensure DirectX SDK or Windows SDK is installed
Ensure you have DirectXMath installed in your TDM-GCC if you use that ( https://github.com/rwusmm-dc/directx-examples/releases/download/dxmath/tdm-gcc-directxmathInstaller.exe )
Build with the provided script:
game.exe - Standalone game launcher
Launch main.exe to open the Project Manager
Create a new project or open an existing one
Projects are stored in Documents/4xEngine/Projects/
Hierarchy Window : Right-click to create primitives (cube, sphere, capsule, plane, triangle, octagon, snowman), lights, cameras, and scripts. Drag entities to reparent.
Properties Window : Edit transform, physics, colors, light settings, and script paths.
Gizmo : Transform entities using translate (T), rotate (R), and scale (S) modes. Toggle between local (L) and world (W) space.
Performance Window : Monitor FPS, draw calls, CPU/GPU usage, and culling statistics.
Create .4xs script files with the following API:
EntityService.findByName(name) - Find entity by name
EntityService.findById(id) - Find entity by ID
EntityService.getAll() - Get all entities
EntityService.create(name) - Create new entity
EntityService.remove(entity) - Remove entity
:getPosition() / :setPosition({x,y,z})
:getRotation() / :setRotation({x,y,z})
:getScale() / :setScale({x,y,z})
:getVelocity() / :setVelocity({x,y,z})
:isLight() , :isCamera() , :isStatic()
LightService.createPoint(name)
LightService.createDirectional(name)
LightService.setColor(entity, {r,g,b})
LightService.setIntensity(entity, value)
LightService.setRange(entity, value)
CameraService.getActive() - Get active camera
CameraService.setActive(entity) - Set active camera
CameraService.create(name) - Create camera
WorldService.setSkyColor({r,g,b})
WorldService.setSunBrightness(value)
WorldService.getSunBrightness()
Input.key(name) - Check key state
Input.mouseX() , Input.mouseY()
Input.mouseDelta() - Returns dx, dy
SyncService.getDelta() - Frame delta time
SyncService.getDiagnostics() - Engine stats
Ensure at least one Camera entity exists in the scene
Select File > Export from the editor menu
The exporter creates a standalone game.exe with all scene data packaged
Shortcut
Action
Ctrl+S
Save scene
Ctrl+Z
Undo
Ctrl+Y
Redo
Delete
Delete selected entity
W/A/S/D
Move camera
Space
Move camera up
X
Move camera down
Right-click
Lock/unlock mouse for camera control
Credits:
https://github.com/ocornut/imgui
https://github.com/luajit/luajit
Culling : Frustum culling with optional D3D11 occlusion queries
Physics : AABB-based collision with impulse resolution, friction, and restitution
Serialization : GAF format supports full entity state including scripts
Lighting : Dynamic sun with time-of-day simulation, point lights with attenuation
Scripting : Each script runs in its own Lua state for isolation
This project is provided as-is for educational and experimental purposes, not ready for real world use yet but can be used for demos, use at your own risk, the project is provided as alpha-state licensed under MIT license.
4x11 Engine is a game engine made using only AI as a experiment, supports DX10/11
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
