---
source: "https://github.com/AdrianCX/picohive-robotarm"
hn_url: "https://news.ycombinator.com/item?id=50235958"
title: "Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai"
article_title: "GitHub - AdrianCX/picohive-robotarm: Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai · GitHub"
image: "https://opengraph.githubassets.com/8ff4e3137d03028f4464d2b58a1cb83c0b73469f42c15a7459ae48ecc1b1149c/AdrianCX/picohive-robotarm"
author: "AdrianCX"
captured_at: "2026-10-25T21:28:58Z"
capture_tool: "hn-digest"
hn_id: 50235958
score: 1
comments: 1
posted_at: "2026-10-25T20:52:34Z"
tags:
  - hacker-news
---

# Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai

- HN: [50235958](https://news.ycombinator.com/item?id=50235958)
- Source: [github.com](https://github.com/AdrianCX/picohive-robotarm)
- Score: 1
- Comments: 1
- Posted: 2026-10-25T20:52:34Z

## Translation

Title: Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai
Article title: GitHub - AdrianCX/picohive-robotarm: Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai · GitHub
Description: Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai - AdrianCX/picohive-robotarm

Article text:
GitHub - AdrianCX/picohive-robotarm: Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai · GitHub
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
AdrianCX
picohive-robotarm Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
docs/ pictures docs/ pictures os @ d6b0d77 os @ d6b0d77 src src .gitignore .gitignore .gitmodules .gitmodules CMakeLists.txt CMakeLists.txt README.md README.md View all files Repository files navigation
Standalone Raspberry Pi Pico W / Pico 2 W controller for an SO-101 / SO-100 arm
Can control over WiFi via an internal Web Page, MQTT or openAI response API compatible remote\
Video: https://www.youtube.com/shorts/e_Y3U83MK2A
I had an RC car and a LeRobot SO-ARM100 and wanted to see if I can combine them and use an LLM to control...
The arm is drawn in 3D from what the servos report; click a dot to put a target down, Go (or Enter) sends the jaws there.
Calibration : set up limits for joints so arm doesn't self destruct... (don't ask)
3D mapping : tell robot which way is left or up
Poses : various poses described for openAI LLM knowledge base
AI : chat box with status responses. The microphone needs https
Drive : Virtual joystick to control the wheels, or you can connect a gamepad
Items to do before
Picohive (optional) - for Debug logs and OTA updates
Add Device on picohive (generate the UUID, MQTT password and log key, pick a name)
Open the device, go to Device Config , press Export and copy the line for step 4 below
OpenAI (optional) - for AI chat box
An API key and a model (ex: gpt-6.1-sol ) on a server that has responses API with storage
Flash a pico, it starts as access point picow_config / randompassword (...ultra secure password...)
Connect to the access point, the page opens by itself, if not, use http://192.168.4.1
Configure wifi and a user/pass for the web ui
Optional :
4.1. Remote server enter line from picohive for logging/OTA/crash reporting
4.2. OpenAI enter the URL, key, model and prompt
Reboot - joins your wifi, you can access via its new IP now (find it in picohive logs or on router)
Reconfigure any time from its new ip address
If you need to wipe, just use the universal flash_nuke.uf2 to wipe the flash then re-flash the firmware
This follows LeRobot, in short:
Add servos one at a time to pico to set their IDs to not overlap (and baud rate to 115200, 1M was too much for my setup with 1K resistor)
Build up the arm and connect things
Calibrate from web page (just press the calibrate and move the arm by hand so it learns limits)
Go to 3D model and map some positions so it knows left from right and up from down...
git clone --recursive https://github.com/AdrianCX/picohive-robotarm.git
cd picohive-robotarm
os/build_via_podman.sh pico2_w # or pico_w
Flash build/<board>/robotarm.uf2
Later updates for OTA need the encrypted file robotarm.uf2.encrypted
Before you build your own, create your own OTA keys and certificate, as described in os/docs/ota.md . The HTTPS certificate is util/make_cert.sh <address> .
The same app builds on Linux with simulated servos
Useful to work on web page with valgrind and gdb
os/build_via_podman.sh linux # build/linux/robotarm
os/run_via_podman.sh # http://localhost:8080, https://localhost:8443
os/run_via_podman.sh --valgrind # or --gdb
Memory usage
[2026-10-25 13:52:12] 192.168.100.69 657 app[build/pico_w/robotarm] version[1.0] os[b697060-dirty] buildid[db40f1ac39e2b3ee3d07d2199aa628118cf588b3] Heap[12308/223652] Free[211344] MinFree[151552] LargestFree[192458] Stack[4236/16384] lastRestartReason[reboot_command]
There's quite a lot of free memory on pico even in worst case (150KB), stack usage is reasonable
Wiring (defaults in src/config/general_config.h , change them there if you wire it differently):
Servo bus: GPIO 0 (uart0 tx) through 1k, and GPIO 1 (rx), both onto the servos' data wire, 115200 baud ( ARM_BAUDRATE )
Servos: 12V and ground straight from the supply, the Pico's ground joins them
Front DRV8833: left front on GPIO 17/16, right front on GPIO 18/19
Rear DRV8833: left rear on GPIO 15/14, right rear on GPIO 12/13 (forward/back of each)
GPIO 20 to both drivers' sleep (EEP), held high
Pico: 5V from a buck converter into VSYS
The bus is half duplex - one data wire shared by every servo. The pico's pins are 3.3V: check a servo's data wire idles at or below that, or put a level shifter in between. The Waveshare bus servo driver board does all of this already, and plugs into the same two pins.
Two pins of one motor share a PWM slice (gpio/2 % 8, even A, odd B), keep that if you move them. MOTOR_PWM 1 drives the pins as PWM, 0 is on/off.
The diagram is generated: python3 docs/pictures/wiring.py docs/pictures/wiring.svg
SO-101 printed parts, or a kit
Printed parts for this build in stl/
[Raspberry Pi Pico 1 W] - you should use Pico2 to add a camera later
5V buck converter for Pico (mini360 is fine)
AGPL-3.0-or-later for this app. The os/ submodule is MPL-2.0, see its LICENSE.
Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai - AdrianCX/picohive-robotarm

GitHub - AdrianCX/picohive-robotarm: Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai · GitHub
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
AdrianCX
picohive-robotarm Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
docs/ pictures docs/ pictures os @ d6b0d77 os @ d6b0d77 src src .gitignore .gitignore .gitmodules .gitmodules CMakeLists.txt CMakeLists.txt README.md README.md View all files Repository files navigation
Standalone Raspberry Pi Pico W / Pico 2 W controller for an SO-101 / SO-100 arm
Can control over WiFi via an internal Web Page, MQTT or openAI response API compatible remote\
Video: https://www.youtube.com/shorts/e_Y3U83MK2A
I had an RC car and a LeRobot SO-ARM100 and wanted to see if I can combine them and use an LLM to control...
The arm is drawn in 3D from what the servos report; click a dot to put a target down, Go (or Enter) sends the jaws there.
Calibration : set up limits for joints so arm doesn't self destruct... (don't ask)
3D mapping : tell robot which way is left or up
Poses : various poses described for openAI LLM knowledge base
AI : chat box with status responses. The microphone needs https
Drive : Virtual joystick to control the wheels, or you can connect a gamepad
Items to do before
Picohive (optional) - for Debug logs and OTA updates
Add Device on picohive (generate the UUID, MQTT password and log key, pick a name)
Open the device, go to Device Config , press Export and copy the line for step 4 below
OpenAI (optional) - for AI chat box
An API key and a model (ex: gpt-6.1-sol ) on a server that has responses API with storage
Flash a pico, it starts as access point picow_config / randompassword (...ultra secure password...)
Connect to the access point, the page opens by itself, if not, use http://192.168.4.1
Configure wifi and a user/pass for the web ui
Optional :
4.1. Remote server enter line from picohive for logging/OTA/crash reporting
4.2. OpenAI enter the URL, key, model and prompt
Reboot - joins your wifi, you can access via its new IP now (find it in picohive logs or on router)
Reconfigure any time from its new ip address
If you need to wipe, just use the universal flash_nuke.uf2 to wipe the flash then re-flash the firmware
This follows LeRobot, in short:
Add servos one at a time to pico to set their IDs to not overlap (and baud rate to 115200, 1M was too much for my setup with 1K resistor)
Build up the arm and connect things
Calibrate from web page (just press the calibrate and move the arm by hand so it learns limits)
Go to 3D model and map some positions so it knows left from right and up from down...
git clone --recursive https://github.com/AdrianCX/picohive-robotarm.git
cd picohive-robotarm
os/build_via_podman.sh pico2_w # or pico_w
Flash build/<board>/robotarm.uf2
Later updates for OTA need the encrypted file robotarm.uf2.encrypted
Before you build your own, create your own OTA keys and certificate, as described in os/docs/ota.md . The HTTPS certificate is util/make_cert.sh <address> .
The same app builds on Linux with simulated servos
Useful to work on web page with valgrind and gdb
os/build_via_podman.sh linux # build/linux/robotarm
os/run_via_podman.sh # http://localhost:8080, https://localhost:8443
os/run_via_podman.sh --valgrind # or --gdb
Memory usage
[2026-10-25 13:52:12] 192.168.100.69 657 app[build/pico_w/robotarm] version[1.0] os[b697060-dirty] buildid[db40f1ac39e2b3ee3d07d2199aa628118cf588b3] Heap[12308/223652] Free[211344] MinFree[151552] LargestFree[192458] Stack[4236/16384] lastRestartReason[reboot_command]
There's quite a lot of free memory on pico even in worst case (150KB), stack usage is reasonable
Wiring (defaults in src/config/general_config.h , change them there if you wire it differently):
Servo bus: GPIO 0 (uart0 tx) through 1k, and GPIO 1 (rx), both onto the servos' data wire, 115200 baud ( ARM_BAUDRATE )
Servos: 12V and ground straight from the supply, the Pico's ground joins them
Front DRV8833: left front on GPIO 17/16, right front on GPIO 18/19
Rear DRV8833: left rear on GPIO 15/14, right rear on GPIO 12/13 (forward/back of each)
GPIO 20 to both drivers' sleep (EEP), held high
Pico: 5V from a buck converter into VSYS
The bus is half duplex - one data wire shared by every servo. The pico's pins are 3.3V: check a servo's data wire idles at or below that, or put a level shifter in between. The Waveshare bus servo driver board does all of this already, and plugs into the same two pins.
Two pins of one motor share a PWM slice (gpio/2 % 8, even A, odd B), keep that if you move them. MOTOR_PWM 1 drives the pins as PWM, 0 is on/off.
The diagram is generated: python3 docs/pictures/wiring.py docs/pictures/wiring.svg
SO-101 printed parts, or a kit
Printed parts for this build in stl/
[Raspberry Pi Pico 1 W] - you should use Pico2 to add a camera later
5V buck converter for Pico (mini360 is fine)
AGPL-3.0-or-later for this app. The os/ submodule is MPL-2.0, see its LICENSE.
Raspberry Pi PicoW + LeRobot SO-ARM100 controlled via a webpage/mqtt/openai
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
