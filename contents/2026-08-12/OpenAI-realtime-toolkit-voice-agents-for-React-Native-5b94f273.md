---
source: "https://github.com/switchboard-sdk/openai-realtime-toolkit"
hn_url: "https://news.ycombinator.com/item?id=49276238"
title: "OpenAI-realtime-toolkit: voice agents for React Native"
article_title: "GitHub - switchboard-sdk/openai-realtime-toolkit: openai-realtime-toolkit is a React Native SDK that provides easy integration with OpenAI's voice-to-voice agents, along with on-device VAD, local turn handling and barge-in detection as first-class primitives for building voice-first agentic applicat\n[truncated]"
author: "chaudharygee"
captured_at: "2026-08-12T18:48:53Z"
capture_tool: "hn-digest"
hn_id: 49276238
score: 2
comments: 0
posted_at: "2026-08-12T17:54:42Z"
tags:
  - hacker-news
  - translated
---

# OpenAI-realtime-toolkit: voice agents for React Native

- HN: [49276238](https://news.ycombinator.com/item?id=49276238)
- Source: [github.com](https://github.com/switchboard-sdk/openai-realtime-toolkit)
- Score: 2
- Comments: 0
- Posted: 2026-08-12T17:54:42Z

## Translation

タイトル: OpenAI-realtime-toolkit: React Native 用の音声エージェント
記事のタイトル: GitHub - switchboard-sdk/openai-realtime-toolkit: openai-realtime-toolkit は、OpenAI の音声対音声エージェントとの簡単な統合を提供する React Native SDK であり、音声ファースト エージェント アプリケーションを構築するためのファーストクラス プリミティブとしてオンデバイス VAD、ローカル ターン処理、バージイン検出も提供します。
[切り捨てられた]
説明: openai-realtime-toolkit は、OpenAI の音声間エージェントとの簡単な統合を提供する React Native SDK であり、音声ファースト エージェント アプリケーションを構築するためのファーストクラス プリミティブとして、オンデバイス VAD、ローカル ターン処理、およびバージイン検出とともに提供されます。 - switchboard-sdk/openai-realtime-toolkit

記事本文:
GitHub - switchboard-sdk/openai-realtime-toolkit: openai-realtime-toolkit は、OpenAI の voice-to-voice エージェントとの簡単な統合を提供する React Native SDK であり、オンデバイス VAD、ローカル ターン処理、および音声ファースト エージェント アプリケーションを構築するためのファーストクラス プリミティブとしてのバージイン検出も備えています。 · GitHub
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
配電盤-SDK
/
openai-リアルタイム-ツールキット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のn

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット .github/ workflows .github/ workflows android android cpp cpp サンプル サンプル ios ios プラグイン プラグイン スクリプト スクリプト src src .clang-format .clang-format .editorconfig .editorconfig .gitignore .gitignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.js .prettierrc.js CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE.txt LICENSE.txt OpenAIRealtimeToolkit.podspec OpenAIRealtimeToolkit.podspec README.md README.md app.plugin.js app.plugin.js babel.config.js babel.config.js eslint.config.js eslint.config.js jest.config.js jest.config.js jest.setup.js jest.setup.js パッケージロック.json パッケージロック.json パッケージ.json パッケージ.json 反応ネイティブ.config.js 反応ネイティブ.config.js tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
openai-realtime-toolkit は、OpenAI の音声間エージェントとの簡単な統合を提供する React Native SDK であり、オンデバイス VAD、ローカル方向転換検出、バージイン処理、および音声ファースト エージェント アプリケーションを構築するためのファーストクラス プリミティブとしてのツール呼び出しを提供します。スイッチボード SDK を利用します。
プラットフォーム
ステータス
iOS
サポートされています
アンドロイド
サポートされています
インストール
npm install @synervoz/openai-realtime-toolkit
要件
要件
最小値
ネイティブに反応する
0.76+
新しいアーキテクチャ
必須 (有効)
iOS
13.4+
アンドロイドNDK
29 — Android を参照
Node.js
22歳以上
openai-realtime-toolkit は C++ TurboModule であり、新しいアーキテクチャが必要です。 React Native CLI と Expo アプリの両方で動作しますが、Expo Go ではネイティブ コードが同梱されているため動作しません。
アプリの ios/ ディレクトリから:
cd ios && ポッドのインストール
アプリの Info.plist にマイク使用量の文字列を追加します (必須、それなしの場合は
AP

マイクが要求されると p がクラッシュします):
< key >NSMicrophoneUsageDescription</ key >
< string >音声アシスタントに使用されます。</ string >
アンドロイド
ここには 2 つのパスがあります。アプリに一致するものを選択してください。
Expo → このセクションを完全にスキップしてください。構成プラグインは 4 つすべてを適用します
プリビルド中に以下の手順を実行します。エキスポに行きます。
React Native CLI (ベア React Native) → 以下の 4 つの手順をアプリに適用します。
手。
openai-realtime-toolkit の C++ TurboModule はアプリのネイティブ ビルドでコンパイルされるため、
アプリは、(a) Switchboard Maven リポジトリを認識し、(b) Prefab を有効にし、(c) を使用してビルドする必要があります。
NDK 29 。 3 つすべてが必要です。
1.Maven リポジトリ。これをアプリのルート android/build.gradle (public、no) に追加します。
資格情報）。プロジェクト レベルで宣言します — React Native の Gradle プラグインにより追加されます
独自のリポジトリも同様なので、設定レベルの dependencyResolutionManagement
Gradle のデフォルトの PREFER_PROJECT モードではブロックは無視されます。
すべてのプロジェクト {
リポジトリ {
maven { url " https://s3.amazonaws.com/synervoz-android-maven-repository " }
}
}
2.プレハブ。アプリの android/app/build.gradle で有効にします。
アンドロイド {
buildFeature {
プレハブの真
}
}
3. NDK 29. アプリの buildscript { ext { … } } ブロックで ndkVersion を上げます。
root android/build.gradle — React Native テンプレートは 27.x を固定しますが、27.x は固定しません。
仕事:
ビルドスクリプト {
内線 {
ndkVersion = " 29.0.14206865 " // テンプレートの 27.x ではありません — 以下を参照してください
// …
}
}
お持ちでない場合は、一度インストールしてください。
" $ANDROID_HOME /cmdline-tools/latest/bin/sdkmanager " " ndk;29.0.14206865 "
スイッチボードのネイティブ ライブラリには、NDK 29 の libc++_shared.so が必要です。 27.x ではアプリがビルドされます
インストールは正常に完了しますが、起動時に UnsatisfiedLinkError: シンボル "__cxa_init_primary_Exception" が見つかりませんで失敗します。
4. 32 ビット x86 を削除します。 Switchboard ライブラリには arm64 が同梱されています

-v8a 、 armeabi-v7a および
x86_64 — 32 ビットの x86 スライスはありませんが、React Native テンプレートの
反応NativeArchitecturesはそれを要求します。アプリ内で削除してください
android/gradle.properties :
反応NativeArchitectures =armeabi-v7a、arm64-v8a、x86_64
run-android と expo run:android は接続されたデバイスの ABI のみを構築するため、
いずれにしても、リリース ビルド、CI または EAS がそれが表面化する場所です。
タスク「:app:configureCMakeDebug[x86]」の実行に失敗しました。
> [CXX1210] … デバッグ|x86 : 互換性のあるライブラリが見つかりません [//SwitchboardSmartTurn/SwitchboardSmartTurn]
次に、以下をビルドします。
npx 反応ネイティブ実行アンドロイド
NDK の確認 (どちらかのパス): すべての Android ビルドは、使用した NDK を出力します。
正しいものが着地したことを確認できます。
[ExpoRootProject] - ndk: 29.0.14206865
openai-realtime-toolkit は Expo アプリで動作しますが、ネイティブ コード (C++
TurboModule + Switchboard フレームワーク）Expo Go では実行できません - あなた
開発ビルドが必要です。
新しいアーキテクチャが必要です。Expo SDK ≥ 52 ではデフォルトでオンになります。
SDK バージョン — React Native ≥ 0.76 (Expo SDK ≥ 52) が必要です。開発され、
RN 0.86、つまり Expo SDK 57 に対してテスト済み: RN が完全に一致するため、一致するものはありません。
TurboModule コード生成の不一致。 52 までの古い SDK も動作するはずです。
npx expo インストール @synervoz/openai-realtime-toolkit
2. 構成プラグインを app.json に追加します。これは、プリビルドの前に行ってください。
それを適用するのが prebuild です。 Switchboard Maven リポジトリを宣言し、Prefab を有効にし、
ndkVersion を 29 に上げ、サポートされていない 32 ビット x86 アーキテクチャを削除します。
アプリ独自の Android ビルド ファイル内の ReactNativeArchitectures — 配線 Expo ではできません
独自に実行します ( expo-build-properties には ndkVersion オプションがありません)。
ネイティブ CLI 自動リンク パスは、プリビルドでは十分に早く着陸できません。マイクの紐と
権限

は組み込み (下記) によって処理されるため、プラグインにはオプションがありません。必須のみです
Android の場合 — 書き込まれるすべての設定は Gradle のものであるため、iOS のみのプロジェクトでは何も行いません。
{
"エキスポ" : {
"プラグイン" : [ " @synervoz/openai-realtime-toolkit " ],
"ios" : {
"infoPlist" : {
"NSMicrophoneUsageDescription" : " 音声アシスタントに使用されます。"
}
}
}
}
iOS マイク プロンプトには NSMicrophoneUsageDescription が必要です。アンドロイド
権限 ( RECORD_AUDIO 、 INTERNET 、 MODIFY_AUDIO_SETTINGS ) は、
ライブラリのマニフェストが自動的にマージされ、追加するものは何もありません。
3. 事前ビルドしてから、開発ビルドをビルドして起動します。
npx expo prebuild # ios/ + android/ を生成します。 iOS はポッドのインストールを実行します (フレームワークを取得します)。
npx expo run:ios # または: npx expo run:android
-p ios / -p android を追加して、単一プラットフォーム用にプリビルドします。以前に事前にビルドした場合
プラグインを追加し、npx expo prebuild を再度実行して、Android 配線を接続します。それ以外の場合は、
生成された android/ には Maven リポジトリ、プレハブ、NDK 27 がなく、アプリは次の時点でクラッシュします。
起動します (Android を参照)。
openai-realtime-toolkit には独自の iOS プライバシー マニフェストがバンドルされており、
そのため、使用する Apple フラグ付き API (モデル ファイルのロードからの FileTimestamp ) は次のとおりです。
すでに宣言されていますので、何もする必要はありません。
アプリレベルでは、引き続き以下を処理します。
Info.plist の NSMicrophoneUsageDescription (iOS インストールを参照) — マイク
プロンプト文字列。 (マイクはプライバシーマニフェスト API ではありません。必要なのはこの文字列だけです。)
App Store のプライバシー ラベル - 音声は OpenAI Realtime API にストリーミングされるため、
それを開示します（例：「音声データ」）。 openai-realtime-toolkit 自体は何も保存しません。
プロバイダーが取得するすべての資格情報はオプションです。これはすぐに機能し、それぞれの資格情報はオプションです。
渡したものは、組み込みのデフォルトをオーバーライドします。
Switchboard — appId / appSecret 、テストと開発ではオプション:
リー

brary には、フォールバック先となる共有のデフォルト資格情報が同梱されています。制作の場合はサインアップしてください
console.switchboard.audio (無料) および
アプリを作成して、その APP_ID と APP_SECRET を取得します。
OpenAI — のリアルタイム対応キー
platform.openai.com 。 openAIApiKey と
セッションは代わりに共有テストキーで実行されます。
OpenAI テスト キーは評価専用です。共有され、レートが制限され、ローテーションされます
予告なしに、それに依存しているアプリは、それが裏返った瞬間に動作を停止し、あなたは
割り当て、請求、使用量の制御は受けられません。独自の openAIApiKey を渡します
あなたが発送するものは何でも。キーが設定されていない場合、ツールキットの console.warn は初期化時に発生します。バンドルされたもの
スイッチボードの認証情報は構築およびテストには問題ありませんが、独自の appId / を使用してください。
運用環境では appSecret を使用するため、アプリは自分の Switchboard アカウントで実行されます。
スイッチボードの APP_ID と APP_SECRET をアプリケーションに安全にバンドルできます。彼らは
公開キーのように機能し、アプリと一緒に配布されることを目的としています。あなたの
OpenAI キーはそうではありません — ソースから外しておいてください (react-native-dotenv など)。
出荷されたアプリは、常駐キーを埋め込むのではなく、サーバー側で一時的なキーを作成します。
アプリを OpenAIRealtimeToolkitProvider でラップし、から駆動します
useOpenAIRealtimeToolkit() フックを持つコンポーネント。 start() はマイクをリクエストし、
音声グラフを構築します — マイク → OpenAI.Realtime → スピーカー (オプションのオンデバイスあり)
方向転換検出とバージイン、およびハードウェア エコー キャンセル (VPIO):
'react' から React をインポートします。
import { Text , TouchableOpacity , View } from 'react-native' ;
インポート {
OpenAIRealtimeToolkitProvider 、
OpenAIRealtimeToolkit を使用し、
useTool 、
'@synervoz/openai-realtime-toolkit' から;
デフォルト関数のエクスポート App ( ) {
戻る (
// appId / appSecret はオプションです。ライブラリのデフォルトを使用する場合は省略してください。
< 開く

AIリアルタイムツールキットプロバイダー
appId = "YOUR_SWITCHBOARD_APP_ID"
appSecret = "YOUR_SWITCHBOARD_APP_SECRET"
openAIApiKey = "YOUR_OPENAI_API_KEY"
指示 = 「あなたは簡潔でフレンドリーな音声アシスタントです。」 >
< 画面 />
</ OpenAIRealtimeToolkitProvider >
) ;
}
関数画面 ( ) {
const {isRunning , connectionStatus , error , start , stop } = useOpenAIRealtimeToolkit();
// モデルが呼び出すことができるツール。その戻り値は自動的に送り返されます。
useTool ( {
名前: 'get_time' 、
description : '現在の時刻を取得します。' 、
パラメータ: { タイプ: 'オブジェクト' 、プロパティ: { } }、
ハンドラー: async() => ({ time: new Date() .toLocaleTimeString() }) 、
} ) ;
戻る (
< ビュー スタイル = { { フレックス : 1 、パディング : 24 、ギャップ : 16 } } >
< TouchableOpacity onPress = { isRunning ?停止 : 開始 } >
< テキスト > { isRunning ? 'やめて' : '話し始めてください' } </ Text >
</ TouchableOpacity >
< テキスト > 接続: { connectionStatus } </ テキスト >
{ ! ！エラー && < テキスト > エラー: { エラー .メッセージ } </ テキスト > }
</ ビュー >
) ;
}
これがアプリ全体です。start() はマイクの許可を処理して接続し、モデルは次のことを行うことができます。
get_time ツールを呼び出します (時間を尋ねてみてください)。 3 つの認証情報行をすべて削除すると、
ライブラリのデフォルトのスイッチボード資格情報と共有 OpenAI テスト キーで引き続き実行されます。
(上記) — あなた自身のものを入れてください

[切り捨てられた]

## Original Extract

openai-realtime-toolkit is a React Native SDK that provides easy integration with OpenAI's voice-to-voice agents, along with on-device VAD, local turn handling and barge-in detection as first-class primitives for building voice-first agentic applications. - switchboard-sdk/openai-realtime-toolkit

GitHub - switchboard-sdk/openai-realtime-toolkit: openai-realtime-toolkit is a React Native SDK that provides easy integration with OpenAI's voice-to-voice agents, along with on-device VAD, local turn handling and barge-in detection as first-class primitives for building voice-first agentic applications. · GitHub
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
switchboard-sdk
/
openai-realtime-toolkit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits .github/ workflows .github/ workflows android android cpp cpp example example ios ios plugin plugin scripts scripts src src .clang-format .clang-format .editorconfig .editorconfig .gitignore .gitignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.js .prettierrc.js CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE.txt LICENSE.txt OpenAIRealtimeToolkit.podspec OpenAIRealtimeToolkit.podspec README.md README.md app.plugin.js app.plugin.js babel.config.js babel.config.js eslint.config.js eslint.config.js jest.config.js jest.config.js jest.setup.js jest.setup.js package-lock.json package-lock.json package.json package.json react-native.config.js react-native.config.js tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json View all files Repository files navigation
openai-realtime-toolkit is a React Native SDK that provides easy integration with OpenAI's voice-to-voice agents, along with on-device VAD, local turn detection, barge-in handling, and tool calling as first-class primitives for building voice-first agentic applications. Powered by the Switchboard SDK .
Platform
Status
iOS
Supported
Android
Supported
Install
npm install @synervoz/openai-realtime-toolkit
Requirements
Requirement
Minimum
React Native
0.76+
New Architecture
Required (enabled)
iOS
13.4+
Android NDK
29 — see Android
Node.js
22+
openai-realtime-toolkit is a C++ TurboModule and requires the New Architecture . It works in both React Native CLI and Expo apps — but not in Expo Go, since it ships native code.
From your app's ios/ directory:
cd ios && pod install
Add a microphone usage string to your app's Info.plist (required — without it the
app crashes when the mic is requested):
< key >NSMicrophoneUsageDescription</ key >
< string >Used for the voice assistant.</ string >
Android
There are two paths here — pick the one that matches your app:
Expo → skip this section entirely. The config plugin applies all four
steps below for you during prebuild . Go to Expo .
React Native CLI (bare React Native) → apply the four steps below to your app by
hand.
openai-realtime-toolkit's C++ TurboModule is compiled in your app's native build, so your
app needs to (a) know the Switchboard Maven repo, (b) enable Prefab, and (c) build with
NDK 29 . All three are required.
1. Maven repo. Add it to your app's root android/build.gradle (public, no
credentials). Declare it at the project level — React Native's Gradle plugin adds
its own repos the same way, so a settings-level dependencyResolutionManagement
block would be ignored under Gradle's default PREFER_PROJECT mode:
allprojects {
repositories {
maven { url " https://s3.amazonaws.com/synervoz-android-maven-repository " }
}
}
2. Prefab. Enable it in your app's android/app/build.gradle :
android {
buildFeatures {
prefab true
}
}
3. NDK 29. Raise ndkVersion in the buildscript { ext { … } } block of your app's
root android/build.gradle — the React Native template pins 27.x, which does not
work:
buildscript {
ext {
ndkVersion = " 29.0.14206865 " // not the template's 27.x — see below
// …
}
}
Install it once if you don't have it:
" $ANDROID_HOME /cmdline-tools/latest/bin/sdkmanager " " ndk;29.0.14206865 "
The Switchboard native libraries need NDK 29's libc++_shared.so . On 27.x the app builds
and installs fine but fails at launch with UnsatisfiedLinkError: cannot locate symbol "__cxa_init_primary_exception" .
4. Drop 32-bit x86. The Switchboard libraries ship arm64-v8a , armeabi-v7a and
x86_64 — there is no 32-bit x86 slice, but the React Native template's
reactNativeArchitectures asks for one. Remove it in your app's
android/gradle.properties :
reactNativeArchitectures =armeabi-v7a,arm64-v8a,x86_64
run-android and expo run:android build only the connected device's ABI, so they pass
either way — a release build, CI or EAS is where it surfaces:
Execution failed for task ':app:configureCMakeDebug[x86]'.
> [CXX1210] … debug|x86 : No compatible library found [//SwitchboardSmartTurn/SwitchboardSmartTurn]
Then build:
npx react-native run-android
Confirming the NDK (either path): every Android build prints the NDK it used, so
you can check the right one landed:
[ExpoRootProject] - ndk: 29.0.14206865
openai-realtime-toolkit works in an Expo app, but because it ships native code (a C++
TurboModule + the Switchboard frameworks) it can not run in Expo Go — you
need a development build .
New architecture is required — on by default in Expo SDK ≥ 52.
SDK version — needs React Native ≥ 0.76 (Expo SDK ≥ 52). Developed and
tested against RN 0.86, i.e. Expo SDK 57 : an exact RN match, so there's no
TurboModule codegen mismatch. Older SDKs down to 52 should also work.
npx expo install @synervoz/openai-realtime-toolkit
2. Add the config plugin to app.json — do this before prebuilding, because
prebuild is what applies it. It declares the Switchboard Maven repo, enables Prefab,
raises ndkVersion to 29, and drops the unsupported 32-bit x86 architecture from
reactNativeArchitectures in your app's own Android build files — the wiring Expo can't
do on its own ( expo-build-properties has no ndkVersion option), and which the React
Native CLI autolinking path can't land early enough under prebuild. The microphone string and
permissions are handled by built-ins (below), so the plugin takes no options. It is only required
for Android — every setting it writes is a Gradle one, so on an iOS-only project it does nothing:
{
"expo" : {
"plugins" : [ " @synervoz/openai-realtime-toolkit " ],
"ios" : {
"infoPlist" : {
"NSMicrophoneUsageDescription" : " Used for the voice assistant. "
}
}
}
}
NSMicrophoneUsageDescription is required for the iOS mic prompt. Android
permissions ( RECORD_AUDIO , INTERNET , MODIFY_AUDIO_SETTINGS ) ship in the
library's manifest and merge in automatically — nothing to add.
3. Prebuild, then build and launch the dev build.
npx expo prebuild # generates ios/ + android/; iOS runs pod install (fetches the frameworks)
npx expo run:ios # or: npx expo run:android
Add -p ios / -p android to prebuild for a single platform. If you prebuilt before
adding the plugin, run npx expo prebuild again so the Android wiring lands — otherwise
the generated android/ has no Maven repo, no Prefab and NDK 27, and the app crashes at
launch (see Android ).
openai-realtime-toolkit bundles its own iOS privacy manifest ,
so the one Apple-flagged API it uses ( FileTimestamp , from loading model files) is
already declared for you — nothing to do.
At the app level you still handle:
NSMicrophoneUsageDescription in Info.plist (see iOS install ) — the mic
prompt string. (The microphone isn't a privacy-manifest API; this string is all it needs.)
App Store privacy labels — audio is streamed to the OpenAI Realtime API , so
disclose that (e.g. "Audio Data"). openai-realtime-toolkit itself stores nothing.
Every credential the provider takes is optional — it works out of the box, and each
one you pass overrides a built-in default.
Switchboard — appId / appSecret , optional for testing and development: the
library ships with shared default credentials it falls back to. For production, sign up
at console.switchboard.audio (free) and
create an app to get its APP_ID and APP_SECRET .
OpenAI — a Realtime-capable key from
platform.openai.com . Omit openAIApiKey and the
session runs on a shared test key instead.
The OpenAI test key is for evaluation only. It's shared, rate-limited, and rotated
without notice — an app relying on it stops working the moment it turns over, and you
get no quota, billing, or usage control over it. Pass your own openAIApiKey for
anything you ship; the toolkit console.warn s at init when no key is set. The bundled
Switchboard credentials are fine to build and test against, but use your own appId /
appSecret in production so the app runs under your own Switchboard account.
Your Switchboard APP_ID and APP_SECRET are safe to bundle in your application . They
function like a publishing key and are intended to be distributed with your app. Your
OpenAI key is not — keep it out of source (e.g. react-native-dotenv ), and for a
shipped app mint an ephemeral key server-side rather than embedding a standing one.
Wrap your app in OpenAIRealtimeToolkitProvider , then drive it from
any component with the useOpenAIRealtimeToolkit() hook. start() requests the mic and
builds the voice graph — microphone → OpenAI.Realtime → speaker, with optional on-device
turn detection and barge-in, and hardware echo cancellation (VPIO):
import React from 'react' ;
import { Text , TouchableOpacity , View } from 'react-native' ;
import {
OpenAIRealtimeToolkitProvider ,
useOpenAIRealtimeToolkit ,
useTool ,
} from '@synervoz/openai-realtime-toolkit' ;
export default function App ( ) {
return (
// appId / appSecret are optional — omit them to use the library's defaults.
< OpenAIRealtimeToolkitProvider
appId = "YOUR_SWITCHBOARD_APP_ID"
appSecret = "YOUR_SWITCHBOARD_APP_SECRET"
openAIApiKey = "YOUR_OPENAI_API_KEY"
instructions = "You are a terse, friendly voice assistant." >
< Screen />
</ OpenAIRealtimeToolkitProvider >
) ;
}
function Screen ( ) {
const { isRunning , connectionStatus , error , start , stop } = useOpenAIRealtimeToolkit ( ) ;
// A tool the model can call; its return value is sent back automatically.
useTool ( {
name : 'get_time' ,
description : 'Get the current time.' ,
parameters : { type : 'object' , properties : { } } ,
handler : async ( ) => ( { time : new Date ( ) . toLocaleTimeString ( ) } ) ,
} ) ;
return (
< View style = { { flex : 1 , padding : 24 , gap : 16 } } >
< TouchableOpacity onPress = { isRunning ? stop : start } >
< Text > { isRunning ? 'Stop' : 'Start talking' } </ Text >
</ TouchableOpacity >
< Text > Connection: { connectionStatus } </ Text >
{ ! ! error && < Text > Error: { error . message } </ Text > }
</ View >
) ;
}
That's the whole app — start() handles mic permission and connects, and the model can
call the get_time tool (try asking it the time). Drop all three credential lines and it
still runs, on the library's default Switchboard credentials and the shared OpenAI test key
( above ) — put your own

[truncated]
