---
source: "https://github.com/microsoft/terminal/issues/20819"
hn_url: "https://news.ycombinator.com/item?id=50166355"
title: "Claude Code abuses WSL2 interop to just send clipboard"
article_title: "OSC 52 clipboard write freezes the UI thread for ~10s when another process holds the clipboard · Issue #20819 · microsoft/terminal · GitHub"
image: "https://opengraph.githubassets.com/eae33f43d462dbafc741a58b65e324689f31a06870aa0e0084abb0f989b453aa/microsoft/terminal/issues/20819"
author: "lazymio"
captured_at: "2026-10-20T14:44:28Z"
capture_tool: "hn-digest"
hn_id: 50166355
score: 1
comments: 0
posted_at: "2026-10-20T13:45:10Z"
tags:
  - hacker-news
---

# Claude Code abuses WSL2 interop to just send clipboard

- HN: [50166355](https://news.ycombinator.com/item?id=50166355)
- Source: [github.com](https://github.com/microsoft/terminal/issues/20819)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T13:45:10Z

## Translation

Title: Claude Code abuses WSL2 interop to just send clipboard
Article title: OSC 52 clipboard write freezes the UI thread for ~10s when another process holds the clipboard · Issue #20819 · microsoft/terminal · GitHub
Description: Windows Terminal version 1.26.2734.0 Windows build number 10.0.26300.0 Other Software No response Steps to reproduce Follow up of #20210 and #20728 Environment: WSL2 Script (assisted by Claude Code): #!/bin/bash # Repro: Windows Terminal...

Article text:
OSC 52 clipboard write freezes the UI thread for ~10s when another process holds the clipboard · Issue #20819 · microsoft/terminal · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
microsoft
OSC 52 clipboard write freezes the UI thread for ~10s when another process holds the clipboard #20819
Copy link New issue Copy link Open Bug Open OSC 52 clipboard write freezes the UI thread for ~10s when another process holds the clipboard #20819 Bug Copy link Labels Issue-Bug It either shouldn't be doing this or needs an investigation. It either shouldn't be doing this or needs an investigation. Needs-Triage It's a new issue that the core contributor team needs to triage at the next triage meeting It's a new issue that the core contributor team needs to triage at the next triage meeting Description
Issue body actions Windows Terminal version
Follow up of #20210 and #20728
Script (assisted by Claude Code):
#! /bin/bash
# Repro: Windows Terminal UI freezes ~10s when an OSC 52 arrives while another
# process has the clipboard open and then calls EmptyClipboard().
# Run from WSL in a *focused* Windows Terminal pane (no tmux needed).
osc52 () { printf ' \e]52;c;%s\a ' " $( printf %s " $1 " | base64 -w0 ) " > /dev/tty ; }
# Helper (separate Windows process): OpenClipboard, say HELD, wait 800ms,
# then EmptyClipboard -> sends WM_DESTROYCLIPBOARD to the owner (= WT).
ps= ' Add-Type -Namespace W -Name G -MemberDefinition @"
[DllImport("user32.dll")] public static extern bool OpenClipboard(IntPtr h);
[DllImport("user32.dll")] public static extern bool EmptyClipboard();
[DllImport("user32.dll")] public static extern bool CloseClipboard();
"@
while (-not [W.G]::OpenClipboard([IntPtr]::Zero)) { Start-Sleep -Milliseconds 5 }
[Console]::Out.WriteLine("HELD"); [Console]::Out.Flush()
Start-Sleep -Milliseconds 800
$sw = [Diagnostics.Stopwatch]::StartNew(); [void][W.G]::EmptyClipboard()
[Console]::Out.WriteLine("EmptyClipboard blocked for " + $sw.ElapsedMilliseconds + " ms")
[void][W.G]::CloseClipboard() '
enc= $( printf %s " $ps " | iconv -t UTF-16LE | base64 -w0 )
osc52 " first " ; sleep 0.5 # 1. WT becomes the clipboard owner
out= $( mktemp )
powershell.exe -NoProfile -NonInteractive -EncodedCommand " $enc " > " $out " 2> /dev/null &
until grep -q HELD " $out " ; do sleep 0.01 ; done # 2. helper now holds the clipboard
t0= $( date +%s.%N )
osc52 " second " # 3. WT: OpenClipboard fails -> Sleep() backoff on the UI thread
wait # 4. helper's EmptyClipboard waits on WT's UI thread
tr -d ' \r ' < " $out " | grep -v HELD
printf ' helper finished after %.1fs; WT should have been "Not Responding" for ~10s\n ' " $( echo " $( date +%s.%N ) - $t0 " | bc ) "
rm -f " $out "
Run this script and this should trigger a freeze. In real world, Claude Code spawns a powershell to send clipboard when tmux exists and thus cause such freezes.
I'm not a Windows expert and much analysis was assisted by Claude Code so feel free to correct me.
Reactions are currently unavailable Activity
lazymio ( wtdcode ) added the Bug issue type on Oct 19, 2026
lazymio ( wtdcode ) added Issue-Bug It either shouldn't be doing this or needs an investigation. It either shouldn't be doing this or needs an investigation. Needs-Triage It's a new issue that the core contributor team needs to triage at the next triage meeting It's a new issue that the core contributor team needs to triage at the next triage meeting on Oct 19, 2026
lhecker commented on Oct 19, 2026
More actions I can imagine that this may happen from time to time, e.g. once a year, but is it realistic for this to be a commonly encountered issue? I don't think so.
wtdcode commented on Oct 20, 2026
More actions
4. I can imagine that this may happen from time to time, e.g. once a year, but is it realistic for this to be a commonly encountered issue? I don't think so.
Hi, thanks for the response. I can reproduce this roughly once per day when copying from Clauce Code in tmux.
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
Author show comment More actions
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
Author show comment More actions
wtdcode commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
lhecker commented on Oct 20, 2026
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
Author More actions
21. Can your AI think of ways to investigate this issue on your machine? For instance, WPR traces could possibly help, but maybe there are some other approaches I can't think of right now.
It unfortunately really isn't reproducible on my machine.
The report from Claude Code points out the powershell spawned by itself, exactly like my reproduction. Does my reproduction script work for you?
Below is the full response from Claude Code (AI Slop warning!):
Every opener is the Set-Clipboard that Clauderop on each copy:
powershell.exe -NoProfile -NonInteractive -Command "[Console]::InputEncoding = [Text.Encoding]::UTF8; Set-Clipboard -Value ([Console]::In.ReadToEnd())"
No other process touched the clipboard while it was held. The clipboard sequence number did not change during the hold, so powershell never managed to write. I believe it is blocked in EmptyClipboard() → WM_DESTROYCLIPBOARD to the owner, which is WT's UI thread. That thread is sleepiackoff because powershell holds the clipboard.Each freeze lasts about 2.55 s, which equals 10+20+…+1280 ms of backoff. So it's not an app that forgets to call CloseClipboard() ; the two sides wait on each other.
Why rapid copying is needed: on each copy, the OSC 52 reaches WT about 20 ms later, while powershell starts writing about 200 ms later. When copies are about 200 ms apart, the next copy's OSC 52 arrives exactly while the previous copy's powershell holds the clipboard. The standaloneroduces the same condition deterministically.
In case you need any further information or investigation, let me know. I can also build Windows Terminal locally if that helps.
wtdcode commented on Oct 20, 2026
Author More actions WPR analysis from Claude Code below. The stack during the freeze is:
KernelBase!SleepEx
TerminalApp!clipboard::open(HWND) TerminalPage.cpp:118 <- Sleep(sleep) in the OpenClipboard retry loop
TerminalApp!TerminalPage::_copyToClipboard$_ResumeCoro$1 TerminalPage.cpp:3641 <- after co_await wil::resume_foreground(Dispatcher())
TerminalApp!delegate<DispatcherQueueHandler, wil::details::dispatcher_handler>::Invoke
Windows.UI.dll / CoreMessaging.dll dispatch
win32kfull!xxxRealInternalGetMessage / NtUserGetMessage
Details
I recorded a WPR trace (GeneralProfile + CPU) while reproducing and looked at the wait chains. The trace has 3 freezes, and all of them look the same.
WT UI thread (WindowsTerminal.exe 1.26.2734.0, tid 76100): in DelayExecution for 640 ms, then for 1280 ms. Symbolized with TerminalApp.pdb from the public symbol server:
KernelBase!SleepEx
TerminalApp!clipboard::open(HWND) TerminalPage.cpp:118 <- Sleep(sleep) in the OpenClipboard retry loop
TerminalApp!TerminalPage::_copyToClipboard$_ResumeCoro$1 TerminalPage.cpp:3641 <- after co_await wil::resume_foreground(Dispatcher())
TerminalApp!delegate<DispatcherQueueHandler, wil::details::dispatcher_handler>::Invoke
Windows.UI.dll / CoreMessaging.dll dispatch
win32kfull!xxxRealInternalGetMessage / NtUserGetMessage
The other side : the powershell.exe -NoProfile -NonInteractive -Command "[Console]::InputEncoding = [Text.Encoding]::UTF8; Set-Clipboard -Value ([Console]::In.ReadToEnd())" that Claude Code spawns through WSL interop on every copy. Blocked in WrUserRequest for ~1.7 s:
win32kfull!xxxSleepThread
win32kfull!xxxInterSendMsgEx
win32kfull!xxxSendTransformableMessageTimeout
win32kfull!xxxSendClipboardMessage
win32kfull!xxxEmptyClipboard
win32kfull!NtUserEmptyClipboard
So powershell has the clipboard open and is sending WM_DESTROYCLIPBOARD to the current owner. The owner is WT's UI thread, which is sleeping in the OpenClipboard retry loop (powershell holds the clipboard) and doesn't pump messages. Neither side can make progress.
The wait never completes normally. In all 3 cases powershell is readied by wsl.exe -> NtTerminateProcess , about 1.99 s after it was spawned (Claude Code's timeout kills it). Only then is the clipboard released, and WT's next retry succeeds. That explains the ~2.6 s freeze here: the WT UI thread stays in the 10+20+...+1280 ms backoff until the retry after the kill. In the standalone repro script nobody kills the other process, so the freeze lasts the full ~10 s backoff.
No other process (clipboard history, Parsec, etc.) appears in the wait chain. My clipboard monitor (GetOpenClipboardWindow / GetClipboardOwner / GetClipboardSequenceNumber sampled every 1 ms) shows the same thing: during every freeze, the clipboard is held by that powershell, WT is the owner, and the sequence number doesn't change, so the powershell write is lost too.
Why it takes rapid copying: on each copy, the OSC 52 reaches WT ~20 ms later, while the native Set-Clipboard opens the clipboard ~200 ms later. When copies come ~200 ms apart, the next copy's OSC 52 arrives just as the previous copy's powershell holds the clipboard. The multiplexer only makes this more likely, because tmux/zellij forward more than one OSC 52 per copy.
So I don't think anyone is "acquiring but not releasing" the clipboard. The cause is sleeping on the UI thread while holding clipboard ownership: whoever opens the clipboard next has to wait for the UI thread to process WM_DESTROYCLIPBOARD. Either pumping messages during the retry (e.g. MsgWaitForMultipleObjectsEx instead of Sleep ) or doing the retry loop off the UI thread should break the cycle.
lhecker commented on Oct 20, 2026
More actions
The other side: the powershell.exe -NoProfile -NonInteractive -Command "[Console]::InputEncoding = [Text.Encoding]::UTF8; Set-Clipboard -Value ([hat Claude Code spawns through WSL interop onevery copy.
wtdcode commented on Oct 20, 2026
More actions
24. > The other side: the powershell.exe -NoProfile -NonInteractive -Command "[Console]::InputEncoding = [Text.Encoding]::UTF8; Set-Clipboard -Value ([hat Claude Code spawns through WSL interop onevery copy.
IT DOES WHAT?
Excuse me...
# IT DOES WHAT?!?!?!?!?!?!?!?!
## WHAT THE FUCK.
Lol. So you now understand why my PoC script looks weird?
wtdcode commented on Oct 20, 2026
More actions Quote the sentence from my original post:
In real world, Claude Code spawns a powershell to send clipboard...
wtdcode commented on Oct 20, 2026
More actions *Sorry I notice that the format of Claude Code output is

[truncated]

## Original Extract

Windows Terminal version 1.26.2734.0 Windows build number 10.0.26300.0 Other Software No response Steps to reproduce Follow up of #20210 and #20728 Environment: WSL2 Script (assisted by Claude Code): #!/bin/bash # Repro: Windows Terminal...

OSC 52 clipboard write freezes the UI thread for ~10s when another process holds the clipboard · Issue #20819 · microsoft/terminal · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
microsoft
OSC 52 clipboard write freezes the UI thread for ~10s when another process holds the clipboard #20819
Copy link New issue Copy link Open Bug Open OSC 52 clipboard write freezes the UI thread for ~10s when another process holds the clipboard #20819 Bug Copy link Labels Issue-Bug It either shouldn't be doing this or needs an investigation. It either shouldn't be doing this or needs an investigation. Needs-Triage It's a new issue that the core contributor team needs to triage at the next triage meeting It's a new issue that the core contributor team needs to triage at the next triage meeting Description
Issue body actions Windows Terminal version
Follow up of #20210 and #20728
Script (assisted by Claude Code):
#! /bin/bash
# Repro: Windows Terminal UI freezes ~10s when an OSC 52 arrives while another
# process has the clipboard open and then calls EmptyClipboard().
# Run from WSL in a *focused* Windows Terminal pane (no tmux needed).
osc52 () { printf ' \e]52;c;%s\a ' " $( printf %s " $1 " | base64 -w0 ) " > /dev/tty ; }
# Helper (separate Windows process): OpenClipboard, say HELD, wait 800ms,
# then EmptyClipboard -> sends WM_DESTROYCLIPBOARD to the owner (= WT).
ps= ' Add-Type -Namespace W -Name G -MemberDefinition @"
[DllImport("user32.dll")] public static extern bool OpenClipboard(IntPtr h);
[DllImport("user32.dll")] public static extern bool EmptyClipboard();
[DllImport("user32.dll")] public static extern bool CloseClipboard();
"@
while (-not [W.G]::OpenClipboard([IntPtr]::Zero)) { Start-Sleep -Milliseconds 5 }
[Console]::Out.WriteLine("HELD"); [Console]::Out.Flush()
Start-Sleep -Milliseconds 800
$sw = [Diagnostics.Stopwatch]::StartNew(); [void][W.G]::EmptyClipboard()
[Console]::Out.WriteLine("EmptyClipboard blocked for " + $sw.ElapsedMilliseconds + " ms")
[void][W.G]::CloseClipboard() '
enc= $( printf %s " $ps " | iconv -t UTF-16LE | base64 -w0 )
osc52 " first " ; sleep 0.5 # 1. WT becomes the clipboard owner
out= $( mktemp )
powershell.exe -NoProfile -NonInteractive -EncodedCommand " $enc " > " $out " 2> /dev/null &
until grep -q HELD " $out " ; do sleep 0.01 ; done # 2. helper now holds the clipboard
t0= $( date +%s.%N )
osc52 " second " # 3. WT: OpenClipboard fails -> Sleep() backoff on the UI thread
wait # 4. helper's EmptyClipboard waits on WT's UI thread
tr -d ' \r ' < " $out " | grep -v HELD
printf ' helper finished after %.1fs; WT should have been "Not Responding" for ~10s\n ' " $( echo " $( date +%s.%N ) - $t0 " | bc ) "
rm -f " $out "
Run this script and this should trigger a freeze. In real world, Claude Code spawns a powershell to send clipboard when tmux exists and thus cause such freezes.
I'm not a Windows expert and much analysis was assisted by Claude Code so feel free to correct me.
Reactions are currently unavailable Activity
lazymio ( wtdcode ) added the Bug issue type on Oct 19, 2026
lazymio ( wtdcode ) added Issue-Bug It either shouldn't be doing this or needs an investigation. It either shouldn't be doing this or needs an investigation. Needs-Triage It's a new issue that the core contributor team needs to triage at the next triage meeting It's a new issue that the core contributor team needs to triage at the next triage meeting on Oct 19, 2026
lhecker commented on Oct 19, 2026
More actions I can imagine that this may happen from time to time, e.g. once a year, but is it realistic for this to be a commonly encountered issue? I don't think so.
wtdcode commented on Oct 20, 2026
More actions
4. I can imagine that this may happen from time to time, e.g. once a year, but is it realistic for this to be a commonly encountered issue? I don't think so.
Hi, thanks for the response. I can reproduce this roughly once per day when copying from Clauce Code in tmux.
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
Author show comment More actions
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
Author show comment More actions
wtdcode commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
lhecker commented on Oct 20, 2026
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
lhecker commented on Oct 20, 2026
wtdcode commented on Oct 20, 2026
Author More actions
21. Can your AI think of ways to investigate this issue on your machine? For instance, WPR traces could possibly help, but maybe there are some other approaches I can't think of right now.
It unfortunately really isn't reproducible on my machine.
The report from Claude Code points out the powershell spawned by itself, exactly like my reproduction. Does my reproduction script work for you?
Below is the full response from Claude Code (AI Slop warning!):
Every opener is the Set-Clipboard that Clauderop on each copy:
powershell.exe -NoProfile -NonInteractive -Command "[Console]::InputEncoding = [Text.Encoding]::UTF8; Set-Clipboard -Value ([Console]::In.ReadToEnd())"
No other process touched the clipboard while it was held. The clipboard sequence number did not change during the hold, so powershell never managed to write. I believe it is blocked in EmptyClipboard() → WM_DESTROYCLIPBOARD to the owner, which is WT's UI thread. That thread is sleepiackoff because powershell holds the clipboard.Each freeze lasts about 2.55 s, which equals 10+20+…+1280 ms of backoff. So it's not an app that forgets to call CloseClipboard() ; the two sides wait on each other.
Why rapid copying is needed: on each copy, the OSC 52 reaches WT about 20 ms later, while powershell starts writing about 200 ms later. When copies are about 200 ms apart, the next copy's OSC 52 arrives exactly while the previous copy's powershell holds the clipboard. The standaloneroduces the same condition deterministically.
In case you need any further information or investigation, let me know. I can also build Windows Terminal locally if that helps.
wtdcode commented on Oct 20, 2026
Author More actions WPR analysis from Claude Code below. The stack during the freeze is:
KernelBase!SleepEx
TerminalApp!clipboard::open(HWND) TerminalPage.cpp:118 <- Sleep(sleep) in the OpenClipboard retry loop
TerminalApp!TerminalPage::_copyToClipboard$_ResumeCoro$1 TerminalPage.cpp:3641 <- after co_await wil::resume_foreground(Dispatcher())
TerminalApp!delegate<DispatcherQueueHandler, wil::details::dispatcher_handler>::Invoke
Windows.UI.dll / CoreMessaging.dll dispatch
win32kfull!xxxRealInternalGetMessage / NtUserGetMessage
Details
I recorded a WPR trace (GeneralProfile + CPU) while reproducing and looked at the wait chains. The trace has 3 freezes, and all of them look the same.
WT UI thread (WindowsTerminal.exe 1.26.2734.0, tid 76100): in DelayExecution for 640 ms, then for 1280 ms. Symbolized with TerminalApp.pdb from the public symbol server:
KernelBase!SleepEx
TerminalApp!clipboard::open(HWND) TerminalPage.cpp:118 <- Sleep(sleep) in the OpenClipboard retry loop
TerminalApp!TerminalPage::_copyToClipboard$_ResumeCoro$1 TerminalPage.cpp:3641 <- after co_await wil::resume_foreground(Dispatcher())
TerminalApp!delegate<DispatcherQueueHandler, wil::details::dispatcher_handler>::Invoke
Windows.UI.dll / CoreMessaging.dll dispatch
win32kfull!xxxRealInternalGetMessage / NtUserGetMessage
The other side : the powershell.exe -NoProfile -NonInteractive -Command "[Console]::InputEncoding = [Text.Encoding]::UTF8; Set-Clipboard -Value ([Console]::In.ReadToEnd())" that Claude Code spawns through WSL interop on every copy. Blocked in WrUserRequest for ~1.7 s:
win32kfull!xxxSleepThread
win32kfull!xxxInterSendMsgEx
win32kfull!xxxSendTransformableMessageTimeout
win32kfull!xxxSendClipboardMessage
win32kfull!xxxEmptyClipboard
win32kfull!NtUserEmptyClipboard
So powershell has the clipboard open and is sending WM_DESTROYCLIPBOARD to the current owner. The owner is WT's UI thread, which is sleeping in the OpenClipboard retry loop (powershell holds the clipboard) and doesn't pump messages. Neither side can make progress.
The wait never completes normally. In all 3 cases powershell is readied by wsl.exe -> NtTerminateProcess , about 1.99 s after it was spawned (Claude Code's timeout kills it). Only then is the clipboard released, and WT's next retry succeeds. That explains the ~2.6 s freeze here: the WT UI thread stays in the 10+20+...+1280 ms backoff until the retry after the kill. In the standalone repro script nobody kills the other process, so the freeze lasts the full ~10 s backoff.
No other process (clipboard history, Parsec, etc.) appears in the wait chain. My clipboard monitor (GetOpenClipboardWindow / GetClipboardOwner / GetClipboardSequenceNumber sampled every 1 ms) shows the same thing: during every freeze, the clipboard is held by that powershell, WT is the owner, and the sequence number doesn't change, so the powershell write is lost too.
Why it takes rapid copying: on each copy, the OSC 52 reaches WT ~20 ms later, while the native Set-Clipboard opens the clipboard ~200 ms later. When copies come ~200 ms apart, the next copy's OSC 52 arrives just as the previous copy's powershell holds the clipboard. The multiplexer only makes this more likely, because tmux/zellij forward more than one OSC 52 per copy.
So I don't think anyone is "acquiring but not releasing" the clipboard. The cause is sleeping on the UI thread while holding clipboard ownership: whoever opens the clipboard next has to wait for the UI thread to process WM_DESTROYCLIPBOARD. Either pumping messages during the retry (e.g. MsgWaitForMultipleObjectsEx instead of Sleep ) or doing the retry loop off the UI thread should break the cycle.
lhecker commented on Oct 20, 2026
More actions
The other side: the powershell.exe -NoProfile -NonInteractive -Command "[Console]::InputEncoding = [Text.Encoding]::UTF8; Set-Clipboard -Value ([hat Claude Code spawns through WSL interop onevery copy.
wtdcode commented on Oct 20, 2026
More actions
24. > The other side: the powershell.exe -NoProfile -NonInteractive -Command "[Console]::InputEncoding = [Text.Encoding]::UTF8; Set-Clipboard -Value ([hat Claude Code spawns through WSL interop onevery copy.
IT DOES WHAT?
Excuse me...
# IT DOES WHAT?!?!?!?!?!?!?!?!
## WHAT THE FUCK.
Lol. So you now understand why my PoC script looks weird?
wtdcode commented on Oct 20, 2026
More actions Quote the sentence from my original post:
In real world, Claude Code spawns a powershell to send clipboard...
wtdcode commented on Oct 20, 2026
More actions *Sorry I notice that the format of Claude Code output is

[truncated]
