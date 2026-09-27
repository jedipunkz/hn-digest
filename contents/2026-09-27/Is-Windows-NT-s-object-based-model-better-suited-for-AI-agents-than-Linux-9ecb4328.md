---
source: "https://www.windowslatest.com/2026/09/26/google-researcher-explains-why-windows-nt-puts-linux-to-shame-and-imagines-an-alternate-history-where-it-won/"
hn_url: "https://news.ycombinator.com/item?id=49866112"
title: "Is Windows NT's object-based model better suited for AI agents than Linux?"
article_title: "Google researcher explains why Windows NT \"puts Linux to shame,\" and imagines an alternate history where it won"
image: "https://www.windowslatest.com/wp-content/uploads/2026/09/Windows-NT.jpg"
author: "indigodaddy"
captured_at: "2026-09-27T12:49:01Z"
capture_tool: "hn-digest"
hn_id: 49866112
score: 1
comments: 0
posted_at: "2026-09-27T12:30:13Z"
tags:
  - hacker-news
---

# Is Windows NT's object-based model better suited for AI agents than Linux?

- HN: [49866112](https://news.ycombinator.com/item?id=49866112)
- Source: [www.windowslatest.com](https://www.windowslatest.com/2026/09/26/google-researcher-explains-why-windows-nt-puts-linux-to-shame-and-imagines-an-alternate-history-where-it-won/)
- Score: 1
- Comments: 0
- Posted: 2026-09-27T12:30:13Z

## Translation

Title: Is Windows NT's object-based model better suited for AI agents than Linux?
Article title: Google researcher explains why Windows NT "puts Linux to shame," and imagines an alternate history where it won
Description: Ex-Microsoft reverse engineer Laurie Kirk says Windows NT's kernel beats Linux on security design, and AI agents are her strongest case.

Article text:
Google researcher explains why Windows NT "puts Linux to shame," and imagines an alternate history where it won
Facebook
Mail
RSS
Twitter
Youtube
Windows 11
Windows 10
Windows 10 PC Apps & Games
Select Theme:
System
Light
Dim
Dark
Windows 11
Electron is one of Windows 11’s biggest RAM offenders, and Discord’s…
Microsoft was wrong about Windows 11, and Windows 10 users watched…
Microsoft tried to ban “Microslop,” and six months later it has…
Microsoft confirms Windows 11 quietly stopped backing up your files due…
I tested Microsoft’s secret tool that debloats Windows 11 and makes…
Windows 10
All Windows 10 PC Apps & Games
Microsoft extended Windows 10, but Nvidia just made the extension meaningless…
Microsoft is limiting Microsoft 365 & OneDrive on Windows 10 to…
Windows 10 KB5120249 keeps the OS protected, direct download links for…
5 years later, Windows 10 refuses to die, and Microsoft just…
Select Theme:
System
Light
Dim
Dark
Google researcher explains why Windows NT “puts Linux to shame,” and imagines an alternate history where it won
A Google researcher who spent four years as a reverse engineer at Microsoft has reopened one of computing’s oldest arguments, and she thinks Linux is on the losing side.
Laurie Kirk , who joined Google in 2024 and runs the LaurieWired YouTube channel (with more than half a million subscribers), wrote on X and LinkedIn that “the NT kernel really is an engineering marvel that still puts Linux to shame in many ways.”
Kirk’s acclaim is for the architecture under Windows, meaning how NT represents resources and decides who can access them. We are not discussing the Start menu, or Copilot or the ads in Windows. I say that because Kirk takes the argument into 2026, asking whether Linux suits AI agents that get more permissions every month, and imagines an alternate history where Microsoft shipped an “Open NT” that Amazon could fork.
As with every unrestrained social media post about Windows and Linux, veteran engineers in the replies pushed back hard, arguing that Linux’s modular design is intentional and that forks would sink any Open NT.
Remember, Laurie Kirk is praising the NT kernel under Windows, not Windows 11
Windows NT has been the base of every major Windows release since 1993 , Windows 11 included. Microsoft hired Dave Cutler in October 1988 after he led VMS development at DEC. According to Microsoft , his small team of ex-DEC engineers spent six months on a specification before writing code, with portability, multiprocessor support, and C2 security certification as goals.
Dave Plummer , the retired Microsoft engineer behind Task Manager, replied that “NT wasn’t Dave Cutler’s first rodeo, and it shows in how forward-looking it was and is.” He added, “After RSX-11M and VMS, Windows NT was his THIRD time crafting the kernel from scratch.”
Still, NT isn’t VMS with a Windows logo. It began as NT OS/2 , a portable OS first written for Intel’s i860 chip before moving to MIPS. Microsoft swapped its main personality from OS/2 to Win32 after Windows 3.1 sold 16 million copies in six months!
As for Windows 11, Franco Minucci, an SoC engineer at Nokia, said it better: “the problems with Windows 11 are more shell issues than kernel issues.” Most of the hate towards Windows 11 was about AI, from the agentic OS backlash to Microsoft telling power users “we care deeply about you” and its AI chief not understanding why AI felt underwhelming .
However, that Microsoft is long gone, as Copilot just got a massive update , with impressive features that CEO Satya Nadella felt compelled to call an OS for work.
Anyway, when a user asked how Kirk could judge a closed-source kernel, she simply replied, “maybe because I used to work there…” reminding everyone that she worked 4 years at Microsoft.
Windows NT is object-based, which is what Kirk means by object-oriented
Kirk’s short version is that NT “was more like an object-oriented language, with a strong security model from day one, whereas Linux is very…not.”
Of course, “object-oriented” is an analogy. Microsoft said Windows isn’t object-oriented in the classical C++ sense , and calls it object-based.
In NT, processes, threads, files, devices, registry keys, mutexes, jobs and access tokens are all objects, across more than 25 object types. One component, the Object Manager, creates and destroys them, keeps their namespace, tracks what each process holds, and tracks access rights. Other components can only use an object through routines exported by its owner, which Microsoft says lets it rewrite a component’s internals without breaking anything else.
Developers see this through handles. When an app opens a file, Windows returns a handle and stores the granted rights with it, and every later request is checked against those rights. If you duplicate a handle, rights can be removed from the copy, but not added.
One consistent model for every resource feels like a luxury, so I can see why Kirk finds it elegant.
NT’s security model was designed around resources, identities, and permissions
Microsoft’s documentation described a security model based primarily on per-object rights with a few system-wide privileges on top.
When you sign in, Windows creates an access token carrying your security identifier (SID), your groups, and your privileges, and every process you start gets a copy. Each securable object has a security descriptor with access control lists (ACLs) saying which SIDs are allowed or denied which rights. Ask for an object, and Windows checks your token against its ACL, then hands you a handle with the granted rights.
The descriptor also holds a system ACL that controls auditing , so admins can log successful and failed access to objects.
However, a strong model doesn’t make a secure OS. NT 3.5 earned its C2 rating in 1995, but only as a standalone PC with no networking , and only after tightening default file and registry permissions. Ken C., an infrastructure engineer formerly at Bloomberg and Reuters, asked Kirk to name “any other product requiring monthly updates… or else you risk being hacked.” Drivers, defaults and decades of compatibility work shape Windows security as much as the kernel does.
Linux’s security model is more complicated, but it isn’t necessarily a flaw
Kirk’s other target is Linux. “Do you use UIDs? GIDs? ACLs? CGROUPs? Policies? SELinux? Filesystem modes?” she asks, arguing there’s no “singular coherent graph of capabilities.”
Each piece does a different job. UIDs and GIDs identify users and groups, while file modes and ACLs guard files. Capabilities split root’s powers, so a web server can bind to port 80 without full root. Namespaces give a process its own view of mounts, processes, or networks (which is how containers work), cgroups cap CPU and memory; and SELinux and other Linux Security Modules (LSMs) enforce extra policy.
Kirk is right that most of it came later. The kernel’s LSM documentation says the NSA presented SELinux in 2001 as a separate patch. Linus Torvalds responded by proposing a framework of hooks any security model could plug into, which was merged in December 2003. The kernel now ships SELinux, AppArmor, Smack, TOMOYO, Landlock and more .
Ken C. argued the add-on model is the unifying philosophy of Unix and Linux and “the whole point.” Kevin Faulkner added that optional SELinux lets some HPC projects skip overhead they can’t afford.
Neither camp gets to define clean architecture. NT centralizes around one object model, while Linux grew composable mechanisms on a Unix base.
AI agents are why Kirk thinks Linux is a poor fit in 2026
The Google researcher brings the AI problem down to one line: “What exactly is this AI agent allowed to do?”
An admin still clicks one thing at a time. An AI agent can run thousands of actions in minutes, execute code it just wrote, read files, and chain permissions together without asking you. The OS has to define, limit, and prove what that software could do.
Kirk says NT “can go the route of explicitly typed resources” and give “really strong centralized audit trails when an agent goes haywire.” It’s a thesis, though.
Linux has a strong answer. Landlock , added in Linux 5.13, lets even unprivileged processes limit their own file system and network access. The rules pass to child processes and can only get stricter, and denied requests go to the audit log by default . Add seccomp, namespaces, cgroups, and LSMs, and Linux can contain an agent. The Landlock docs do admit namespaces “are not designed for access-control,” which is similar to Kirk’s complaint.
So, the debate is about whether an OS built around agents would make these controls easier to reason about, and Microsoft is working on it. At Build 2026, it announced Microsoft Execution Containers (MXC) , where developers declare what an agent can access and Windows enforces it at runtime. MXC session isolation runs agents under their own user accounts, so every action is attributed to the agent. Windows 11’s Agent Workspace already uses ACLs to keep agent accounts within the user’s permissions .
Agents on Windows are getting SIDs, tokens, and ACLs, which is the NT model Kirk is praising. Of course, Microsoft’s agentic OS push, from an AI agent on the taskbar to a background agent with access to personal folders , came with their own malware warnings . Its work doesn’t prove Microsoft agrees with Kirk about Linux, but Redmond is asking her question.
“Open NT” is the alternate history Laurie Kirk wishes Microsoft had written
Kirk says Microsoft “should have made an ‘Open NT’” in the early 2000s, not fully GPL-style, “but one where a large org could say…swap out a memory allocator for their own.” Her example is an early Amazon building “AmazonNT” for EC2 with its own scheduler or network stack, while “the security+compatibility contract keeps a stable baseline on the Microsoft side.”
She admits Microsoft “sorta did this.” Its Shared Source Initiative offered Windows source code to about 1,600 enterprise customers, universities and governments. In 2001, Austria’s Interior Ministry became the first European government to get the Windows XP source. In 2006, the Windows Research Kernel let academics modify NT’s scheduler and memory manager, but only for teaching and research.
Shared Source let you study the code. It didn’t let Amazon ship a commercial NT fork.
But why didn’t Microsoft go further? In my opinion, it has to do with licensing revenue, intellectual property, support costs, and the fear of breaking the compatibility promise that kept Win32 apps running for three decades.
Linux’s biggest counterargument to Open NT is the cost of forks
The strongest pushback came from David Airlie , a longtime Linux kernel developer who maintains its graphics subsystem. “Forks don’t work like that. Too expensive over time to maintain,” he wrote, asking Kirk to imagine still maintaining a fork of NT’s memory manager 20 years later.
Kirk pointed to Java, where Corretto, GraalVM, Azul and OpenJ9 all target one spec with different engines. Airlie replied that those have paying customers, and an internal NT fork without revenue would constantly face pressure to disappear. “Lots of people fork Linux, last a year or five, then realize they needed to work upstream due to cost pressure,” he added.
If Amazon had rewritten NT’s scheduler, every Microsoft security fix would need to be merged and retested every month, until the fork became a second kernel team. With Linux, companies send useful changes upstream, where the community maintains them, so its architecture and development model can’t be separated.
Kirk assumes a stable NT contract would keep forks sustainable. Airlie says a divergent kernel is expensive even when the source is fully open, which is a fair argument to say the least.
NT has its own baggage, and critics are right to bring it up
Ken C, who has years of experience in IT and Systems Reliability, said VMS’s consistency only partly carried over to NT. He asked how to get deterministic pe

[truncated]

## Original Extract

Ex-Microsoft reverse engineer Laurie Kirk says Windows NT's kernel beats Linux on security design, and AI agents are her strongest case.

Google researcher explains why Windows NT "puts Linux to shame," and imagines an alternate history where it won
Facebook
Mail
RSS
Twitter
Youtube
Windows 11
Windows 10
Windows 10 PC Apps & Games
Select Theme:
System
Light
Dim
Dark
Windows 11
Electron is one of Windows 11’s biggest RAM offenders, and Discord’s…
Microsoft was wrong about Windows 11, and Windows 10 users watched…
Microsoft tried to ban “Microslop,” and six months later it has…
Microsoft confirms Windows 11 quietly stopped backing up your files due…
I tested Microsoft’s secret tool that debloats Windows 11 and makes…
Windows 10
All Windows 10 PC Apps & Games
Microsoft extended Windows 10, but Nvidia just made the extension meaningless…
Microsoft is limiting Microsoft 365 & OneDrive on Windows 10 to…
Windows 10 KB5120249 keeps the OS protected, direct download links for…
5 years later, Windows 10 refuses to die, and Microsoft just…
Select Theme:
System
Light
Dim
Dark
Google researcher explains why Windows NT “puts Linux to shame,” and imagines an alternate history where it won
A Google researcher who spent four years as a reverse engineer at Microsoft has reopened one of computing’s oldest arguments, and she thinks Linux is on the losing side.
Laurie Kirk , who joined Google in 2024 and runs the LaurieWired YouTube channel (with more than half a million subscribers), wrote on X and LinkedIn that “the NT kernel really is an engineering marvel that still puts Linux to shame in many ways.”
Kirk’s acclaim is for the architecture under Windows, meaning how NT represents resources and decides who can access them. We are not discussing the Start menu, or Copilot or the ads in Windows. I say that because Kirk takes the argument into 2026, asking whether Linux suits AI agents that get more permissions every month, and imagines an alternate history where Microsoft shipped an “Open NT” that Amazon could fork.
As with every unrestrained social media post about Windows and Linux, veteran engineers in the replies pushed back hard, arguing that Linux’s modular design is intentional and that forks would sink any Open NT.
Remember, Laurie Kirk is praising the NT kernel under Windows, not Windows 11
Windows NT has been the base of every major Windows release since 1993 , Windows 11 included. Microsoft hired Dave Cutler in October 1988 after he led VMS development at DEC. According to Microsoft , his small team of ex-DEC engineers spent six months on a specification before writing code, with portability, multiprocessor support, and C2 security certification as goals.
Dave Plummer , the retired Microsoft engineer behind Task Manager, replied that “NT wasn’t Dave Cutler’s first rodeo, and it shows in how forward-looking it was and is.” He added, “After RSX-11M and VMS, Windows NT was his THIRD time crafting the kernel from scratch.”
Still, NT isn’t VMS with a Windows logo. It began as NT OS/2 , a portable OS first written for Intel’s i860 chip before moving to MIPS. Microsoft swapped its main personality from OS/2 to Win32 after Windows 3.1 sold 16 million copies in six months!
As for Windows 11, Franco Minucci, an SoC engineer at Nokia, said it better: “the problems with Windows 11 are more shell issues than kernel issues.” Most of the hate towards Windows 11 was about AI, from the agentic OS backlash to Microsoft telling power users “we care deeply about you” and its AI chief not understanding why AI felt underwhelming .
However, that Microsoft is long gone, as Copilot just got a massive update , with impressive features that CEO Satya Nadella felt compelled to call an OS for work.
Anyway, when a user asked how Kirk could judge a closed-source kernel, she simply replied, “maybe because I used to work there…” reminding everyone that she worked 4 years at Microsoft.
Windows NT is object-based, which is what Kirk means by object-oriented
Kirk’s short version is that NT “was more like an object-oriented language, with a strong security model from day one, whereas Linux is very…not.”
Of course, “object-oriented” is an analogy. Microsoft said Windows isn’t object-oriented in the classical C++ sense , and calls it object-based.
In NT, processes, threads, files, devices, registry keys, mutexes, jobs and access tokens are all objects, across more than 25 object types. One component, the Object Manager, creates and destroys them, keeps their namespace, tracks what each process holds, and tracks access rights. Other components can only use an object through routines exported by its owner, which Microsoft says lets it rewrite a component’s internals without breaking anything else.
Developers see this through handles. When an app opens a file, Windows returns a handle and stores the granted rights with it, and every later request is checked against those rights. If you duplicate a handle, rights can be removed from the copy, but not added.
One consistent model for every resource feels like a luxury, so I can see why Kirk finds it elegant.
NT’s security model was designed around resources, identities, and permissions
Microsoft’s documentation described a security model based primarily on per-object rights with a few system-wide privileges on top.
When you sign in, Windows creates an access token carrying your security identifier (SID), your groups, and your privileges, and every process you start gets a copy. Each securable object has a security descriptor with access control lists (ACLs) saying which SIDs are allowed or denied which rights. Ask for an object, and Windows checks your token against its ACL, then hands you a handle with the granted rights.
The descriptor also holds a system ACL that controls auditing , so admins can log successful and failed access to objects.
However, a strong model doesn’t make a secure OS. NT 3.5 earned its C2 rating in 1995, but only as a standalone PC with no networking , and only after tightening default file and registry permissions. Ken C., an infrastructure engineer formerly at Bloomberg and Reuters, asked Kirk to name “any other product requiring monthly updates… or else you risk being hacked.” Drivers, defaults and decades of compatibility work shape Windows security as much as the kernel does.
Linux’s security model is more complicated, but it isn’t necessarily a flaw
Kirk’s other target is Linux. “Do you use UIDs? GIDs? ACLs? CGROUPs? Policies? SELinux? Filesystem modes?” she asks, arguing there’s no “singular coherent graph of capabilities.”
Each piece does a different job. UIDs and GIDs identify users and groups, while file modes and ACLs guard files. Capabilities split root’s powers, so a web server can bind to port 80 without full root. Namespaces give a process its own view of mounts, processes, or networks (which is how containers work), cgroups cap CPU and memory; and SELinux and other Linux Security Modules (LSMs) enforce extra policy.
Kirk is right that most of it came later. The kernel’s LSM documentation says the NSA presented SELinux in 2001 as a separate patch. Linus Torvalds responded by proposing a framework of hooks any security model could plug into, which was merged in December 2003. The kernel now ships SELinux, AppArmor, Smack, TOMOYO, Landlock and more .
Ken C. argued the add-on model is the unifying philosophy of Unix and Linux and “the whole point.” Kevin Faulkner added that optional SELinux lets some HPC projects skip overhead they can’t afford.
Neither camp gets to define clean architecture. NT centralizes around one object model, while Linux grew composable mechanisms on a Unix base.
AI agents are why Kirk thinks Linux is a poor fit in 2026
The Google researcher brings the AI problem down to one line: “What exactly is this AI agent allowed to do?”
An admin still clicks one thing at a time. An AI agent can run thousands of actions in minutes, execute code it just wrote, read files, and chain permissions together without asking you. The OS has to define, limit, and prove what that software could do.
Kirk says NT “can go the route of explicitly typed resources” and give “really strong centralized audit trails when an agent goes haywire.” It’s a thesis, though.
Linux has a strong answer. Landlock , added in Linux 5.13, lets even unprivileged processes limit their own file system and network access. The rules pass to child processes and can only get stricter, and denied requests go to the audit log by default . Add seccomp, namespaces, cgroups, and LSMs, and Linux can contain an agent. The Landlock docs do admit namespaces “are not designed for access-control,” which is similar to Kirk’s complaint.
So, the debate is about whether an OS built around agents would make these controls easier to reason about, and Microsoft is working on it. At Build 2026, it announced Microsoft Execution Containers (MXC) , where developers declare what an agent can access and Windows enforces it at runtime. MXC session isolation runs agents under their own user accounts, so every action is attributed to the agent. Windows 11’s Agent Workspace already uses ACLs to keep agent accounts within the user’s permissions .
Agents on Windows are getting SIDs, tokens, and ACLs, which is the NT model Kirk is praising. Of course, Microsoft’s agentic OS push, from an AI agent on the taskbar to a background agent with access to personal folders , came with their own malware warnings . Its work doesn’t prove Microsoft agrees with Kirk about Linux, but Redmond is asking her question.
“Open NT” is the alternate history Laurie Kirk wishes Microsoft had written
Kirk says Microsoft “should have made an ‘Open NT’” in the early 2000s, not fully GPL-style, “but one where a large org could say…swap out a memory allocator for their own.” Her example is an early Amazon building “AmazonNT” for EC2 with its own scheduler or network stack, while “the security+compatibility contract keeps a stable baseline on the Microsoft side.”
She admits Microsoft “sorta did this.” Its Shared Source Initiative offered Windows source code to about 1,600 enterprise customers, universities and governments. In 2001, Austria’s Interior Ministry became the first European government to get the Windows XP source. In 2006, the Windows Research Kernel let academics modify NT’s scheduler and memory manager, but only for teaching and research.
Shared Source let you study the code. It didn’t let Amazon ship a commercial NT fork.
But why didn’t Microsoft go further? In my opinion, it has to do with licensing revenue, intellectual property, support costs, and the fear of breaking the compatibility promise that kept Win32 apps running for three decades.
Linux’s biggest counterargument to Open NT is the cost of forks
The strongest pushback came from David Airlie , a longtime Linux kernel developer who maintains its graphics subsystem. “Forks don’t work like that. Too expensive over time to maintain,” he wrote, asking Kirk to imagine still maintaining a fork of NT’s memory manager 20 years later.
Kirk pointed to Java, where Corretto, GraalVM, Azul and OpenJ9 all target one spec with different engines. Airlie replied that those have paying customers, and an internal NT fork without revenue would constantly face pressure to disappear. “Lots of people fork Linux, last a year or five, then realize they needed to work upstream due to cost pressure,” he added.
If Amazon had rewritten NT’s scheduler, every Microsoft security fix would need to be merged and retested every month, until the fork became a second kernel team. With Linux, companies send useful changes upstream, where the community maintains them, so its architecture and development model can’t be separated.
Kirk assumes a stable NT contract would keep forks sustainable. Airlie says a divergent kernel is expensive even when the source is fully open, which is a fair argument to say the least.
NT has its own baggage, and critics are right to bring it up
Ken C, who has years of experience in IT and Systems Reliability, said VMS’s consistency only partly carried over to NT. He asked how to get deterministic pe

[truncated]
