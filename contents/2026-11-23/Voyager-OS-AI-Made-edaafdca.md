---
source: "https://github.com/LetMePlaySong/VOYAGER-OS-AI-MADE-"
hn_url: "https://news.ycombinator.com/item?id=50620710"
title: "Voyager OS – AI Made"
article_title: "GitHub - LetMePlaySong/VOYAGER-OS-AI-MADE-: Voyager OS is a from-scratch, bootable 64-bit operating system written entirely in C and C++ — no host OS at runtime. It boots via Limine on BIOS/UEFI and ships a full GUI desktop: file manager, web browser, media player, editors, networking, and sound. Bu\n[truncated]"
image: "https://opengraph.githubassets.com/a924d1e782f42eb84d5ff0125f70df44d9760c4563c604971d314e5ea660a2d9/LetMePlaySong/VOYAGER-OS-AI-MADE-"
author: "retro95"
captured_at: "2026-11-23T23:27:58Z"
capture_tool: "hn-digest"
hn_id: 50620710
score: 1
comments: 0
posted_at: "2026-11-23T23:02:31Z"
tags:
  - hacker-news
---

# Voyager OS – AI Made

- HN: [50620710](https://news.ycombinator.com/item?id=50620710)
- Source: [github.com](https://github.com/LetMePlaySong/VOYAGER-OS-AI-MADE-)
- Score: 1
- Comments: 0
- Posted: 2026-11-23T23:02:31Z

## Translation

Title: Voyager OS – AI Made
Article title: GitHub - LetMePlaySong/VOYAGER-OS-AI-MADE-: Voyager OS is a from-scratch, bootable 64-bit operating system written entirely in C and C++ — no host OS at runtime. It boots via Limine on BIOS/UEFI and ships a full GUI desktop: file manager, web browser, media player, editors, networking, and sound. Bu
[truncated]
Description: Voyager OS is a from-scratch, bootable 64-bit operating system written entirely in C and C++ — no host OS at runtime. It boots via Limine on BIOS/UEFI and ships a full GUI desktop: file manager, web browser, media player, editors, networking, and sound. Built completely by AI. QEMU recommended; Virt
[truncated]

Article text:
GitHub - LetMePlaySong/VOYAGER-OS-AI-MADE-: Voyager OS is a from-scratch, bootable 64-bit operating system written entirely in C and C++ — no host OS at runtime. It boots via Limine on BIOS/UEFI and ships a full GUI desktop: file manager, web browser, media player, editors, networking, and sound. Built completely by AI. QEMU recommended; VirtualBox is unstable. Experimental — it has bugs. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
LetMePlaySong
VOYAGER-OS-AI-MADE- Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Voyager OS is a from-scratch, bootable 64-bit operating system written entirely in C and C++ — no host OS at runtime. It boots via Limine on BIOS/UEFI and ships a full GUI desktop: file manager, web browser, media player, editors, networking, and sound. Built completely by AI. QEMU recommended; VirtualBox is unstable. Experimental — it has bugs.
Readme Unlicense Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
Your ISO is here Your ISO is here .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md RUN VOYAGER.bat RUN VOYAGER.bat RUN VOYAGER.command RUN VOYAGER.command run-voyager.sh run-voyager.sh View all files Repository files navigation
Voyager OS is a from-scratch, bootable 64-bit operating system written
entirely in C and C++ — no host OS at runtime. It boots via Limine on
BIOS/UEFI and ships a full GUI desktop: file manager, web browser,
media player, editors, networking, and sound. Built completely by AI.
Experimental — it has bugs. Not a daily driver.
Voyager OS is a from-scratch, bootable 64-bit operating system with a
full graphical desktop — built entirely in C and C++ as a freestanding
kernel with no host OS dependencies at runtime.
It boots on both BIOS and UEFI through the Limine bootloader, runs in
QEMU or on real hardware (USB), and ships with a complete desktop
environment: file manager, web browser, media player, editors, games,
persistent settings, networking, and sound.
Not a daily driver. Voyager OS is a hobby/experimental project.
Do not use it as your main operating system — expect crashes,
data loss, and rough edges. It's for trying out, learning, and fun.
VM note: Use QEMU — it's the best and most stable option, and
every run script uses it automatically. VirtualBox is unstable
with this OS: expect display glitches, input problems, and audio
issues. Booting the ISO on real hardware via USB works too.
Made entirely by AI. Every line of code in Voyager OS —
kernel, drivers, GUI, apps, web engine — was written by an AI coding
agent (Devin, by Cognition) working interactively with a human who
supplied the idea and feedback. No human wrote the source.
Quick start — no building required
This repository ships a ready-to-boot ISO. Just pick your OS and
double-click (or run) the matching launcher — it boots Voyager OS in
QEMU with networking + sound enabled automatically. No compilers,
no dependencies to install beyond QEMU itself.
The script finds QEMU (or tells you how to install it), creates the
persistent settings disk ( voyager-hd.vdi ) on first run, and boots
the ISO. Close the QEMU window to exit.
Double-click RUN VOYAGER.command
If macOS Gatekeeper blocks the double-click, open Terminal ,
cd into this folder, and run bash run-voyager.sh — the .sh
script is the universal fallback on every platform.
The script auto-installs QEMU via Homebrew if needed, creates the
settings disk, and boots the ISO.
Open a terminal in this folder and run:
./run-voyager.sh # or: bash run-voyager.sh
It auto-installs QEMU via apt/dnf/pacman if needed, creates the
settings disk, and boots the ISO.
Flash Your ISO is here/voyager.iso to a USB stick (4 GB+) with
balenaEtcher (etcher.balena.io) or Rufus (rufus.ie, Windows):
select the ISO, select the drive, flash. Boot the PC from the USB.
Boot splash, animated window spin-in / spin-out transitions
Draggable, resizable windows with traffic-light controls
Zoom button: boing resize to fullscreen and back, live relayout
Animated dock with magnification, scroll arrows, window lists,
auto-hide in fullscreen, drop shadows on all chrome
Desktop icons with drag-and-drop, marquee selection, context menus
Menu bar: Voyager menu, network popup, volume slider, calendar,
live clock
Full cursor system: arrow, hand, I-beam, busy, resize grips
9 free-licensed photographic wallpapers (CC0 / public domain)
Persistent settings across reboots (wallpaper, username, files)
App
Description
Files
Finder-style file manager — System HD, user files, detected disks, drag & drop, thumbnails, rename/delete, disk eject/refresh/properties
Voyage Web
Web browser — dynamic tabs, back/forward/reload, URL bar with Frogfind search, HTML/CSS/JS rendering engine, images, animated GIFs, downloads
Voyage Player
Music player — MP3/WAV/OGG, playlist, scrolling song title, elapsed/remaining LCD, keeps playing while minimized
Quill
Plain-text editor with Save As picker
Memo
Sticky notes
Photos
Image viewer for PNG/JPEG/GIF
Calc
Calculator
Console
Command-line shell (disabled — doesn't work yet)
Settings
Wallpaper picker, Date & Time (timezone, auto-DST, SNTP sync), Task Manager with End Task
System
Kernel : freestanding x86-64, Limine boot protocol, GDT/IDT, heap,
framebuffer abstraction, PS/2 keyboard + mouse
Storage : custom Voyager volume format (VFS) user disk, System HD
catalog from ISO modules, ATA driver with hot-plug media polling
Networking : e1000 driver, ARP, ICMP, UDP, DNS, TCP, HTTP client,
SNTP time sync
Audio : AC97 driver with PCM ring buffer, MP3 via minimp3,
WAV/OGG support, software volume
Web engine : custom HTML parser + layout + JS interpreter subset,
CSS-ish styling, forms, links, images, scrolling
Reliability : OOM watchdog, disk-space guard on saves/downloads,
styled modal dialogs
Your ISO is here/
voyager.iso Ready-to-boot Voyager OS image (54 MB)
RUN VOYAGER.bat Windows launcher — boots the ISO in QEMU
RUN VOYAGER.command macOS launcher (double-click)
run-voyager.sh Linux/macOS launcher — boots the ISO in QEMU
LICENSE Project license
All three launchers do the same thing: locate/install QEMU, create the
16 MB voyager-hd.vdi settings disk on first run, then boot
voyager.iso with e1000 networking, AC97 sound, and a serial log
( qemu-serial.log ).
Bootloader : Limine (BSD-2-Clause)
MP3 decoding : minimp3 (CC0 / public domain)
JPEG decoding : picojpeg (public domain)
Wallpapers : free-licensed photography via Openverse
Test music : "Straight Fuse" by French Fuse
Voyager OS is an independent project — not affiliated with any
commercial OS vendor.
Voyager OS is a from-scratch, bootable 64-bit operating system written entirely in C and C++ — no host OS at runtime. It boots via Limine on BIOS/UEFI and ships a full GUI desktop: file manager, web browser, media player, editors, networking, and sound. Built completely by AI. QEMU recommended; VirtualBox is unstable. Experimental — it has bugs.
Readme Unlicense Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Voyager OS is a from-scratch, bootable 64-bit operating system written entirely in C and C++ — no host OS at runtime. It boots via Limine on BIOS/UEFI and ships a full GUI desktop: file manager, web browser, media player, editors, networking, and sound. Built completely by AI. QEMU recommended; Virt
[truncated]

GitHub - LetMePlaySong/VOYAGER-OS-AI-MADE-: Voyager OS is a from-scratch, bootable 64-bit operating system written entirely in C and C++ — no host OS at runtime. It boots via Limine on BIOS/UEFI and ships a full GUI desktop: file manager, web browser, media player, editors, networking, and sound. Built completely by AI. QEMU recommended; VirtualBox is unstable. Experimental — it has bugs. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
LetMePlaySong
VOYAGER-OS-AI-MADE- Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Voyager OS is a from-scratch, bootable 64-bit operating system written entirely in C and C++ — no host OS at runtime. It boots via Limine on BIOS/UEFI and ships a full GUI desktop: file manager, web browser, media player, editors, networking, and sound. Built completely by AI. QEMU recommended; VirtualBox is unstable. Experimental — it has bugs.
Readme Unlicense Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
Your ISO is here Your ISO is here .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md RUN VOYAGER.bat RUN VOYAGER.bat RUN VOYAGER.command RUN VOYAGER.command run-voyager.sh run-voyager.sh View all files Repository files navigation
Voyager OS is a from-scratch, bootable 64-bit operating system written
entirely in C and C++ — no host OS at runtime. It boots via Limine on
BIOS/UEFI and ships a full GUI desktop: file manager, web browser,
media player, editors, networking, and sound. Built completely by AI.
Experimental — it has bugs. Not a daily driver.
Voyager OS is a from-scratch, bootable 64-bit operating system with a
full graphical desktop — built entirely in C and C++ as a freestanding
kernel with no host OS dependencies at runtime.
It boots on both BIOS and UEFI through the Limine bootloader, runs in
QEMU or on real hardware (USB), and ships with a complete desktop
environment: file manager, web browser, media player, editors, games,
persistent settings, networking, and sound.
Not a daily driver. Voyager OS is a hobby/experimental project.
Do not use it as your main operating system — expect crashes,
data loss, and rough edges. It's for trying out, learning, and fun.
VM note: Use QEMU — it's the best and most stable option, and
every run script uses it automatically. VirtualBox is unstable
with this OS: expect display glitches, input problems, and audio
issues. Booting the ISO on real hardware via USB works too.
Made entirely by AI. Every line of code in Voyager OS —
kernel, drivers, GUI, apps, web engine — was written by an AI coding
agent (Devin, by Cognition) working interactively with a human who
supplied the idea and feedback. No human wrote the source.
Quick start — no building required
This repository ships a ready-to-boot ISO. Just pick your OS and
double-click (or run) the matching launcher — it boots Voyager OS in
QEMU with networking + sound enabled automatically. No compilers,
no dependencies to install beyond QEMU itself.
The script finds QEMU (or tells you how to install it), creates the
persistent settings disk ( voyager-hd.vdi ) on first run, and boots
the ISO. Close the QEMU window to exit.
Double-click RUN VOYAGER.command
If macOS Gatekeeper blocks the double-click, open Terminal ,
cd into this folder, and run bash run-voyager.sh — the .sh
script is the universal fallback on every platform.
The script auto-installs QEMU via Homebrew if needed, creates the
settings disk, and boots the ISO.
Open a terminal in this folder and run:
./run-voyager.sh # or: bash run-voyager.sh
It auto-installs QEMU via apt/dnf/pacman if needed, creates the
settings disk, and boots the ISO.
Flash Your ISO is here/voyager.iso to a USB stick (4 GB+) with
balenaEtcher (etcher.balena.io) or Rufus (rufus.ie, Windows):
select the ISO, select the drive, flash. Boot the PC from the USB.
Boot splash, animated window spin-in / spin-out transitions
Draggable, resizable windows with traffic-light controls
Zoom button: boing resize to fullscreen and back, live relayout
Animated dock with magnification, scroll arrows, window lists,
auto-hide in fullscreen, drop shadows on all chrome
Desktop icons with drag-and-drop, marquee selection, context menus
Menu bar: Voyager menu, network popup, volume slider, calendar,
live clock
Full cursor system: arrow, hand, I-beam, busy, resize grips
9 free-licensed photographic wallpapers (CC0 / public domain)
Persistent settings across reboots (wallpaper, username, files)
App
Description
Files
Finder-style file manager — System HD, user files, detected disks, drag & drop, thumbnails, rename/delete, disk eject/refresh/properties
Voyage Web
Web browser — dynamic tabs, back/forward/reload, URL bar with Frogfind search, HTML/CSS/JS rendering engine, images, animated GIFs, downloads
Voyage Player
Music player — MP3/WAV/OGG, playlist, scrolling song title, elapsed/remaining LCD, keeps playing while minimized
Quill
Plain-text editor with Save As picker
Memo
Sticky notes
Photos
Image viewer for PNG/JPEG/GIF
Calc
Calculator
Console
Command-line shell (disabled — doesn't work yet)
Settings
Wallpaper picker, Date & Time (timezone, auto-DST, SNTP sync), Task Manager with End Task
System
Kernel : freestanding x86-64, Limine boot protocol, GDT/IDT, heap,
framebuffer abstraction, PS/2 keyboard + mouse
Storage : custom Voyager volume format (VFS) user disk, System HD
catalog from ISO modules, ATA driver with hot-plug media polling
Networking : e1000 driver, ARP, ICMP, UDP, DNS, TCP, HTTP client,
SNTP time sync
Audio : AC97 driver with PCM ring buffer, MP3 via minimp3,
WAV/OGG support, software volume
Web engine : custom HTML parser + layout + JS interpreter subset,
CSS-ish styling, forms, links, images, scrolling
Reliability : OOM watchdog, disk-space guard on saves/downloads,
styled modal dialogs
Your ISO is here/
voyager.iso Ready-to-boot Voyager OS image (54 MB)
RUN VOYAGER.bat Windows launcher — boots the ISO in QEMU
RUN VOYAGER.command macOS launcher (double-click)
run-voyager.sh Linux/macOS launcher — boots the ISO in QEMU
LICENSE Project license
All three launchers do the same thing: locate/install QEMU, create the
16 MB voyager-hd.vdi settings disk on first run, then boot
voyager.iso with e1000 networking, AC97 sound, and a serial log
( qemu-serial.log ).
Bootloader : Limine (BSD-2-Clause)
MP3 decoding : minimp3 (CC0 / public domain)
JPEG decoding : picojpeg (public domain)
Wallpapers : free-licensed photography via Openverse
Test music : "Straight Fuse" by French Fuse
Voyager OS is an independent project — not affiliated with any
commercial OS vendor.
Voyager OS is a from-scratch, bootable 64-bit operating system written entirely in C and C++ — no host OS at runtime. It boots via Limine on BIOS/UEFI and ships a full GUI desktop: file manager, web browser, media player, editors, networking, and sound. Built completely by AI. QEMU recommended; VirtualBox is unstable. Experimental — it has bugs.
Readme Unlicense Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
