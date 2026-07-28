---
source: "https://giodicanio.com/2026/07/28/why-we-should-double-check-the-ai-output-a-bug-which-wasnt/"
hn_url: "https://news.ycombinator.com/item?id=49086556"
title: "We Cannot 100% Trust AI-Generated Output: A Concrete Example"
article_title: "Why We Should Double-Check the AI Output: A Bug Which Wasn’t – Giovanni Dicanio's Blog"
author: "movd128"
captured_at: "2026-07-28T17:19:13Z"
capture_tool: "hn-digest"
hn_id: 49086556
score: 2
comments: 1
posted_at: "2026-07-28T16:42:57Z"
tags:
  - hacker-news
  - translated
---

# We Cannot 100% Trust AI-Generated Output: A Concrete Example

- HN: [49086556](https://news.ycombinator.com/item?id=49086556)
- Source: [giodicanio.com](https://giodicanio.com/2026/07/28/why-we-should-double-check-the-ai-output-a-bug-which-wasnt/)
- Score: 2
- Comments: 1
- Posted: 2026-07-28T16:42:57Z

## Translation

タイトル: AI によって生成された出力を 100% 信頼することはできません: 具体的な例
記事のタイトル: AI 出力をダブルチェックする必要がある理由: 存在しなかったバグ – Giovanni Dicanio のブログ
説明: これは、AI によって生成された結果が 100% 完全に信頼されるべきではないことを示す具体的な実例です。

記事本文:
AI 出力を再確認する必要がある理由: 存在しなかったバグ – Giovanni Dicanio のブログ
コンテンツにスキップ
ジョバンニ・ディカーニオのブログ
インターネット上のジョバンニ ディカーニオのプログラミング コーナー
AI 出力を再確認する必要がある理由: 存在しなかったバグ
これは、AI が生成した結果を 100% 完全に信頼すべきではないことを示す具体的な実例です。
最近、私はクロードに、WinReg C++ ライブラリ (低レベル C インターフェイス Windows レジストリ API の高レベル C++ ラッパー) のコードをレビューするよう依頼しました。
分析の結果、クロードは私のコードに長さゼロのバグがあると報告しました。特に、長さゼロの REG_SZ/REG_EXPAND_SZ 値が RegKey クラスの GetStringValue、GetExpandStringValue、TryGetStringValue、および TryGetExpandStringValue メソッドをクラッシュさせるとクロードは述べました。
特に、Claude は、バイナリを返すゲッター (RegKey::GetBinaryValue など) では dataSize == 0 を正しく保護していましたが、文字列ゲッターにはそのような保護がないことを指摘しました。彼らは無条件に次のことを行います。
result.resize((dataSize / sizeof(wchar_t)) - 1);
Claude は、REG_SZ 値と REG_EXPAND_SZ 値が cbData == 0 (ゼロバイト、NUL がまったくない、単一の NUL '\0' で作成された空の文字列とは異なる) で正当に格納できることを指定しました。
上記のステートメントで dataSize 変数を 0 に置き換えると、次のようになります。
result.resize(SIZE_MAX);
これは std::length_error 例外をスローします。
Claude は、長さがゼロの場合を防ぐためにバイナリ ゲッターに既に実装したものと同じエッジケース チェック ロジックを使用して、上記のコードを修正することを提案しました。
if (データサイズ == 0)
{
result.clear();
}
それ以外の場合
{
result.resize((dataSize / sizeof(wchar_t)) - 1);
}
私の WinReg ライブラリはかなり厳しいテストを受けており、いくつかのエッジ ケースに関連するバグがあり、すでに修正されていました。

d、それで私は興味を持って、レジストリに長さ 0 の文字列値を書き込んで、既存のコードでそれを読み取ってみました。そして、(少なくともコードをテストした Windows 11 では) dataSize == 0 条件が実行時にヒットしなかったことに注目しました。
これは、C++ コードで RegGetValue(W) API を呼び出しているためです。この API は、レジストリに格納されている文字列に NUL 終端文字がない場合でも、契約により NUL 終端文字列を返すことが保証されています。 (これは RegQueryValueEx のような古い API には当てはまりません。)
私はその点を指摘したクロードに返信しましたが、クロードは次のように修正しました。
あなたは正しいです。反発してくれてありがとう。バグとしてフラグを立てる前に、RegGetValueW の null 終了保証を考慮すべきでした。
これらの AI ツールは非常に強力である可能性がありますが、ここで重要な点は、それらは単なるツールであることを忘れてはならず、AI によって生成された出力やコードを 100% 信頼してはいけないということです。その AI コードにはバグや間違った仮定が隠れている可能性があるためです。
X で共有 (新しいウィンドウで開きます)
×
Facebook で共有 (新しいウィンドウで開きます)
フェイスブック
AI 出力を再確認する必要がある理由: 存在しなかったバグ
メモリのブロブを受け取る C++ 関数を宣言するにはどうすればよいですか?
Windows API ネイティブ C/C++ プログラミングにおける char-TCHAR-wchar_t 振り子
記事を提案するための IsoCpp.org プロセスが壊れているため修正する必要がある
文字列内の次の Unicode コード ポイントの検索: UTF-8 と UTF-16
アルゴリズムとデータ構造
組み立て
ATL
ベストプラクティス
BSTR
バグ
C
C#
C++
チャットGPT
コードレビュー
コンソール
C文字列
DLL
整数オーバーフロー
LPCWSTR
LPWSTR
MFC
破棄しない
OOP
最適化
PCWSTR
パフォーマンス
複数の視力
プログラミング
PWSTR
レジストリ
リソース
さび
SafeInt
SSO
STL
文字列
文字列ビュー
ユニコード
Unicode 変換
UNICODE_STRING
符号なし整数
UTF-8
UTF-16
ベクトル
VSコード
窓
Windows カーネルモード
WinReg

購読する
購読済み
ジョバンニ・ディカーニオのブログ
すでに WordPress.com アカウントをお持ちですか?今すぐログインしてください。

## Original Extract

This is a concrete real-world example showing how AI-generated results should not be 100% completely trusted.

Why We Should Double-Check the AI Output: A Bug Which Wasn’t – Giovanni Dicanio's Blog
Skip to content
Giovanni Dicanio's Blog
Giovanni Dicanio's Programming Corner on the Internet
Why We Should Double-Check the AI Output: A Bug Which Wasn’t
This is a concrete real-world example showing how AI-generated results should not be 100% completely trusted.
Recently I asked Claude to review the code of my WinReg C++ library (which is a C++ high-level wrapper around the low-level C-interface Windows Registry API).
As a result of its analysis, Claude reported that there were zero-length bugs in my code, in particular Claude stated that zero-length REG_SZ/REG_EXPAND_SZ values crash the GetStringValue, GetExpandStringValue, TryGetStringValue and TryGetExpandStringValue methods of the RegKey class.
In particular, Claude noted that I correctly guarded against dataSize == 0 in the binary-returning getters (like RegKey::GetBinaryValue), but the string getters do not have such guard; they unconditionally do:
result.resize((dataSize / sizeof(wchar_t)) - 1);
Claude specified that REG_SZ and REG_EXPAND_SZ values can be legitimately stored with cbData == 0 (zero bytes, no NUL at all; different from an empty string made by a single NUL ‘\0’).
If you substitute zero for the dataSize variable in the above statement, you end up with:
result.resize(SIZE_MAX);
which would throw a std::length_error exception.
Claude proposed to fix the above code using the same edge-case check logic I had already implemented in the binary getters to guard against the zero-length case:
if (dataSize == 0)
{
result.clear();
}
else
{
result.resize((dataSize / sizeof(wchar_t)) - 1);
}
My WinReg library is quite battle-tested, and there were bugs related to some edge cases that I had already fixed, so I was curious, and tried writing a zero-length string value in the registry, and read it back with my existing code. And I noted that (at least in Windows 11 where I tested my code) the dataSize == 0 condition was not hit at run-time.
That is because in my C++ code I invoke the RegGetValue(W) API, which by contract guarantees to return a NUL-terminated string, even if the string stored in the registry doesn’t have a NUL-terminator. (This is not the case for older APIs like RegQueryValueEx .)
I replied to Claude pointing that out, and Claude corrected itself:
You’re right, and thanks for the pushback — I should have accounted for RegGetValueW ‘s null-termination guarantee before flagging that as a bug.
Those AI tools can be very powerful, but the key takeway here is that we should not forget that they are just tools , and we should not trust AI-generated output and code 100%, because bugs and wrong assumptions can be hidden in that AI code, too.
Share on X (Opens in new window)
X
Share on Facebook (Opens in new window)
Facebook
Why We Should Double-Check the AI Output: A Bug Which Wasn’t
How to Declare a C++ Function that Takes a Blob of Memory?
The char-TCHAR-wchar_t Pendulum in Windows API Native C/C++ Programming
The IsoCpp.org Process for Suggesting Articles Is Broken and Should Be Fixed
Finding the Next Unicode Code Point in Strings: UTF-8 vs. UTF-16
Algorithms and Data Structures
Assembly
ATL
Best practices
BSTR
Bugs
C
C#
C++
ChatGPT
Code Review
Console
CString
DLL
Integer overflow
LPCWSTR
LPWSTR
MFC
nodiscard
OOP
Optimization
PCWSTR
Performance
Pluralsight
Programming
PWSTR
Registry
Resources
Rust
SafeInt
SSO
STL
Strings
string_view
Unicode
Unicode Conversions
UNICODE_STRING
unsigned int
UTF-8
UTF-16
vector
VSCode
Windows
Windows Kernel Mode
WinReg
Subscribe
Subscribed
Giovanni Dicanio's Blog
Already have a WordPress.com account? Log in now.
