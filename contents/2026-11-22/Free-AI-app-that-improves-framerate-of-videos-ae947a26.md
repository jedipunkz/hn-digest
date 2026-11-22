---
source: "https://flowframes.app/"
hn_url: "https://news.ycombinator.com/item?id=50604005"
title: "Free AI app that improves framerate of videos"
article_title: "Flowframes - AI Video Frame Interpolation for Windows"
image: ""
author: "illidan1000"
captured_at: "2026-11-22T21:39:48Z"
capture_tool: "hn-digest"
hn_id: 50604005
score: 1
comments: 0
posted_at: "2026-11-22T21:19:32Z"
tags:
  - hacker-news
---

# Free AI app that improves framerate of videos

- HN: [50604005](https://news.ycombinator.com/item?id=50604005)
- Source: [flowframes.app](https://flowframes.app/)
- Score: 1
- Comments: 0
- Posted: 2026-11-22T21:19:32Z

## Translation

Title: Free AI app that improves framerate of videos
Article title: Flowframes - AI Video Frame Interpolation for Windows
Description: Flowframes uses AI (RIFE) to generate new frames between your video’s existing ones - turn 24 or 30 FPS footage into smooth 60, 120 or higher FPS. Free, open-source, GPU-accelerated frame interpolation for Windows 10 and 11.

Article text:
Flowframes - AI Video Frame Interpolation for Windows Flowframes Features How it works FAQ Download Flowframes
Create smoother, higher-frame-rate video with AI-generated frames. Flowframes turns 24 or 30 FPS footage into fluid 60, 120 or custom-frame-rate video using RIFE and other advanced interpolation models - all processed locally on your own GPU.
Download for Windows How it works Free download · Flowframes 1.36.0 · Windows 10 & 11 · build options
Windows 10 & 11 Open-source donationware Nvidia, AMD & Intel GPUs Enlarge Smoother motion, generated between the original frames
Flowframes analyses the motion between existing frames and generates new ones that follow that movement - for a higher frame rate, smoother motion and cleaner slow motion. Unlike simple frame blending, AI interpolation reconstructs where objects should be at each new moment in time, preserving clarity that conventional resampling tends to lose.
Built for real video workflows
Made for video editors, animators, visual artists, gamers and restoration workflows - anyone who needs better motion without rebuilding a project frame by frame.
Convert 24 or 30 FPS to 60 FPS
Increase the frame rate of films, edits, gameplay captures, screen recordings and generated video while keeping the original duration.
Generate extra frames before slowing footage down in your editor, for cleaner slow-motion that does not rely on duplicated or blended frames.
Detect and remove repeated frames before interpolation to get smoother motion from anime, motion graphics and content drawn on twos or threes.
Render alternate frames from a 3D scene, then generate the in-between frames with Flowframes to cut rendering time on suitable shots.
Smooth clips produced at low or inconsistent frame rates. Scene detection and de-duplication make Flowframes a natural fit for generative-video workflows.
Create seamless animated loops
Loop interpolation generates frames between the final frame and the first frame so an animated loop returns smoothly to its starting point.
Import PNG, JPEG and other supported image sequences directly, then export the result as a video, animated image or new frame sequence.
Three steps from source file to finished, higher-frame-rate video. No separate editor required.
Drop in a supported video, GIF or sequence of images. Flowframes reads common formats including MP4, MKV, MOV, WEBM, GIF, PNG and JPEG.
Choose the interpolation settings
Pick an interpolation factor or target frame rate, choose the AI model, and adjust optional controls. The defaults work well for many videos.
Encode to MP4, MKV, WEBM, MOV, GIF or an image sequence. Audio, subtitles and compatible metadata can be carried over from the source.
RIFE, DAIN and more - built around the best models
Flowframes is designed around RIFE, a fast AI model for frame interpolation, with additional and experimental models available depending on the build. For most people, RIFE is the recommended start.
Recommended for compatible Nvidia GPUs when speed is the priority. CUDA processing can use half-precision acceleration to reduce VRAM use and run faster on supported hardware.
Runs through the NCNN framework using Vulkan, for broad compatibility across supported Nvidia, AMD and Intel graphics - no Nvidia CUDA environment required.
An alternative interpolation model available through Vulkan-compatible NCNN processing. Useful for testing different motion characteristics and comparing against RIFE.
Experimental model options in supported builds. They provide additional approaches for specific footage, testing and advanced interpolation workflows.
Made for Nvidia, AMD and Intel graphics
Flowframes does not require an Nvidia GPU for standard NCNN processing - Vulkan-based models run on supported graphics hardware from multiple vendors.
Use RIFE CUDA for the fastest processing on supported Nvidia hardware, or run the Vulkan-based NCNN implementation.
Use RIFE NCNN or another available NCNN model through Vulkan. An Nvidia GPU is not required to use Flowframes.
Modern Vulkan-capable Intel graphics can use the NCNN path. Performance and compatibility depend on the specific GPU and driver.
Actual speed depends on the model, source resolution, interpolation factor, GPU performance, available VRAM and encoding settings.
Smart processing for clean results
Frame interpolation is only as good as how it handles duplicates, cuts and audio. Flowframes covers all three so you spend less time fixing artifacts.
Intelligent frame de-duplication
Many 2D animations hold the same drawing across several frames. Flowframes can detect and remove those duplicates before interpolation, so the model works with real movement instead of identical images - and it stays off for camera footage and gameplay.
A hard cut has no continuous motion across it. Flowframes detects scene changes and avoids generating frames across recognised cuts, preventing the morphing artifacts that ruin automated interpolation. Sensitivity is adjustable.
Keep audio, subtitles & metadata
Raising a video’s frame rate should not mean rebuilding its audio. Flowframes can transfer audio tracks, subtitle tracks and compatible metadata from the source, copying streams without re-encoding where the container allows.
Import common production and distribution formats, then export to the container and codec you need.
Export codecs include H.264, H.265 / HEVC, VP9, AV1 , subject to the selected encoder and hardware configuration.
Advanced control, without the busywork
Sensible defaults let you start immediately. These controls are there when a project needs them.
Encode the output while interpolation is still running, and optionally delete processed frames to save temporary storage.
RIFE UHD Mode adjusts model scaling for high-resolution footage; a maximum processing resolution can also be set for speed.
Generate a multiplied frame rate or cap the export to a specific maximum, such as 24 FPS in to exactly 60 FPS out.
Preserve and interpolate alpha transparency when the input, intermediate and output formats all support it.
Generate motion from the end of an animation back to its first frame to create a smoother repeating loop.
Use JPEG intermediate frames to reduce storage use and speed up extraction when lossless PNG is unnecessary.
Run extraction, interpolation and encoding as separate steps when frames need to be inspected, edited or removed.
Assign one or more supported GPUs by device ID for compatible processing modes and hardware setups.
Configure codec, quality, bitrate, container and other export parameters with advanced FFmpeg options.
Define output-name patterns from project and source variables for more consistent file organisation.
Download Flowframes for Windows
Free, open-source AI frame interpolation that runs on your own GPU - no uploads, no per-video fees, no cloud queues. Built for Windows 10 and 11.
Download Flowframes for Windows Installer (.exe) · Flowframes 1.36.0 · Windows 10 & 11
Flowframes is open-source donationware: the newest early-access builds reach supporters first, and the source is always public. Want the very latest version? Support development and get early-access builds on Patreon .
Flowframes runs interpolation locally on the Windows PC where it is installed. Older graphics hardware may lack the Vulkan features NCNN needs.
A Vulkan-capable graphics card
A supported Nvidia, AMD or Intel graphics driver
Storage for the source, temporary frames and exported result
A modern CUDA-capable Nvidia GPU
A modern multi-core Intel Core or AMD Ryzen CPU
Solid-state storage for frame extraction and encoding
Not supported: GPUs such as the Nvidia GT 710 and Intel HD 3000/4000 series, and Windows 7.
When frame interpolation works best
AI interpolation performs best when movement is visible and reasonably predictable between adjacent frames. It creates plausible new frames - it cannot recover detail that was never in the source.
Camera movement with consistent motion
Anime and hand-drawn animation (after correct de-duplication)
AI-generated clips with low frame rates
• Large objects appearing or disappearing between frames
• Fast cuts that are not detected correctly
• Flashing lights or rapid exposure changes
• Reflections, particles, smoke, water or transparent objects
• Severe compression artifacts or corrupted source timing
Flowframes is a Windows application for AI-powered video frame interpolation. It analyses the motion between existing frames and generates new intermediate frames to create smoother playback, higher frame rates and more natural slow motion. It is built around advanced models such as RIFE and supports both CUDA and Vulkan-based processing.
Flowframes can increase the frame rate of video, prepare footage for slow motion, smooth 2D animation, improve AI-generated clips, interpolate gameplay recordings, create seamless animated loops and generate missing frames for suitable 3D rendering workflows. It works with complete videos, animated GIFs and image sequences.
No. Flowframes increases temporal resolution by generating extra frames between the original frames. It does not increase the width, height or spatial detail of the source. A 1080p video stays 1080p - the difference is that motion looks smoother because more frames are shown each second.
Yes. Flowframes can generate additional frames and export at a higher frame rate, including common targets such as 48, 60, 120 or 240 FPS. You can also set a maximum output frame rate - for example, generating frames from a 24 FPS source and downsampling the result to exactly 60 FPS.
RIFE is the recommended starting point for most videos, with a strong balance of quality, speed and hardware compatibility. Nvidia systems can use RIFE CUDA for faster processing; RIFE NCNN uses Vulkan and supports a wider range of GPUs. DAIN, FLAVR and XVFI are available for comparison and specific footage.
RIFE CUDA is designed for compatible Nvidia graphics cards and is generally the faster option on supported hardware. RIFE NCNN runs through Vulkan and offers broader GPU compatibility, making it the right choice for AMD systems and for Nvidia systems that do not use the CUDA path. The intended result is similar; performance, memory use and model availability may differ.
Yes. Compatible AMD GPUs can use Vulkan-based NCNN models such as RIFE NCNN. CUDA modes are exclusive to Nvidia hardware, but an Nvidia GPU is not required to use Flowframes - the graphics card just needs to support the Vulkan features the selected NCNN model requires.
Standard builds include the components required for normal use, and the installer can download necessary dependencies automatically. Most people do not need to install Python, PyTorch, FFmpeg or AI models by hand. Advanced users can configure a system Python environment or compile the open-source project themselves.
Flowframes reads widely used formats such as MP4, MKV, MOV, WEBM, GIF and BIK, and can import image sequences in PNG and JPEG. Results export as MP4, MKV, MOV, WEBM, GIF or a sequence of frames, with codecs including H.264, H.265/HEVC, VP9 and AV1 depending on the encoding configuration.
Flowframes can transfer supported audio tracks, subtitle tracks and compatible metadata from the source to the exported video. Streams can often be copied without quality loss when the chosen output container supports them; compatibility depends on the source format, output format and encoding settings.
Frame de-duplication detects and removes consecutive duplicate frames before interpolation. It is mainly for 2D animation, anime and motion graphics that intentionally hold the same image across several frames. Removing those repeats lets the model work with real movement instead of identical images. It should usually be disabled for camera footage, gameplay and 3D-rendered video.
Choppy results are often caused by incorrect frame de-duplication. Dark, low-contrast or slowly changing frames can be wrongly detected as duplicates and removed. For

[truncated]

## Original Extract

Flowframes uses AI (RIFE) to generate new frames between your video’s existing ones - turn 24 or 30 FPS footage into smooth 60, 120 or higher FPS. Free, open-source, GPU-accelerated frame interpolation for Windows 10 and 11.

Flowframes - AI Video Frame Interpolation for Windows Flowframes Features How it works FAQ Download Flowframes
Create smoother, higher-frame-rate video with AI-generated frames. Flowframes turns 24 or 30 FPS footage into fluid 60, 120 or custom-frame-rate video using RIFE and other advanced interpolation models - all processed locally on your own GPU.
Download for Windows How it works Free download · Flowframes 1.36.0 · Windows 10 & 11 · build options
Windows 10 & 11 Open-source donationware Nvidia, AMD & Intel GPUs Enlarge Smoother motion, generated between the original frames
Flowframes analyses the motion between existing frames and generates new ones that follow that movement - for a higher frame rate, smoother motion and cleaner slow motion. Unlike simple frame blending, AI interpolation reconstructs where objects should be at each new moment in time, preserving clarity that conventional resampling tends to lose.
Built for real video workflows
Made for video editors, animators, visual artists, gamers and restoration workflows - anyone who needs better motion without rebuilding a project frame by frame.
Convert 24 or 30 FPS to 60 FPS
Increase the frame rate of films, edits, gameplay captures, screen recordings and generated video while keeping the original duration.
Generate extra frames before slowing footage down in your editor, for cleaner slow-motion that does not rely on duplicated or blended frames.
Detect and remove repeated frames before interpolation to get smoother motion from anime, motion graphics and content drawn on twos or threes.
Render alternate frames from a 3D scene, then generate the in-between frames with Flowframes to cut rendering time on suitable shots.
Smooth clips produced at low or inconsistent frame rates. Scene detection and de-duplication make Flowframes a natural fit for generative-video workflows.
Create seamless animated loops
Loop interpolation generates frames between the final frame and the first frame so an animated loop returns smoothly to its starting point.
Import PNG, JPEG and other supported image sequences directly, then export the result as a video, animated image or new frame sequence.
Three steps from source file to finished, higher-frame-rate video. No separate editor required.
Drop in a supported video, GIF or sequence of images. Flowframes reads common formats including MP4, MKV, MOV, WEBM, GIF, PNG and JPEG.
Choose the interpolation settings
Pick an interpolation factor or target frame rate, choose the AI model, and adjust optional controls. The defaults work well for many videos.
Encode to MP4, MKV, WEBM, MOV, GIF or an image sequence. Audio, subtitles and compatible metadata can be carried over from the source.
RIFE, DAIN and more - built around the best models
Flowframes is designed around RIFE, a fast AI model for frame interpolation, with additional and experimental models available depending on the build. For most people, RIFE is the recommended start.
Recommended for compatible Nvidia GPUs when speed is the priority. CUDA processing can use half-precision acceleration to reduce VRAM use and run faster on supported hardware.
Runs through the NCNN framework using Vulkan, for broad compatibility across supported Nvidia, AMD and Intel graphics - no Nvidia CUDA environment required.
An alternative interpolation model available through Vulkan-compatible NCNN processing. Useful for testing different motion characteristics and comparing against RIFE.
Experimental model options in supported builds. They provide additional approaches for specific footage, testing and advanced interpolation workflows.
Made for Nvidia, AMD and Intel graphics
Flowframes does not require an Nvidia GPU for standard NCNN processing - Vulkan-based models run on supported graphics hardware from multiple vendors.
Use RIFE CUDA for the fastest processing on supported Nvidia hardware, or run the Vulkan-based NCNN implementation.
Use RIFE NCNN or another available NCNN model through Vulkan. An Nvidia GPU is not required to use Flowframes.
Modern Vulkan-capable Intel graphics can use the NCNN path. Performance and compatibility depend on the specific GPU and driver.
Actual speed depends on the model, source resolution, interpolation factor, GPU performance, available VRAM and encoding settings.
Smart processing for clean results
Frame interpolation is only as good as how it handles duplicates, cuts and audio. Flowframes covers all three so you spend less time fixing artifacts.
Intelligent frame de-duplication
Many 2D animations hold the same drawing across several frames. Flowframes can detect and remove those duplicates before interpolation, so the model works with real movement instead of identical images - and it stays off for camera footage and gameplay.
A hard cut has no continuous motion across it. Flowframes detects scene changes and avoids generating frames across recognised cuts, preventing the morphing artifacts that ruin automated interpolation. Sensitivity is adjustable.
Keep audio, subtitles & metadata
Raising a video’s frame rate should not mean rebuilding its audio. Flowframes can transfer audio tracks, subtitle tracks and compatible metadata from the source, copying streams without re-encoding where the container allows.
Import common production and distribution formats, then export to the container and codec you need.
Export codecs include H.264, H.265 / HEVC, VP9, AV1 , subject to the selected encoder and hardware configuration.
Advanced control, without the busywork
Sensible defaults let you start immediately. These controls are there when a project needs them.
Encode the output while interpolation is still running, and optionally delete processed frames to save temporary storage.
RIFE UHD Mode adjusts model scaling for high-resolution footage; a maximum processing resolution can also be set for speed.
Generate a multiplied frame rate or cap the export to a specific maximum, such as 24 FPS in to exactly 60 FPS out.
Preserve and interpolate alpha transparency when the input, intermediate and output formats all support it.
Generate motion from the end of an animation back to its first frame to create a smoother repeating loop.
Use JPEG intermediate frames to reduce storage use and speed up extraction when lossless PNG is unnecessary.
Run extraction, interpolation and encoding as separate steps when frames need to be inspected, edited or removed.
Assign one or more supported GPUs by device ID for compatible processing modes and hardware setups.
Configure codec, quality, bitrate, container and other export parameters with advanced FFmpeg options.
Define output-name patterns from project and source variables for more consistent file organisation.
Download Flowframes for Windows
Free, open-source AI frame interpolation that runs on your own GPU - no uploads, no per-video fees, no cloud queues. Built for Windows 10 and 11.
Download Flowframes for Windows Installer (.exe) · Flowframes 1.36.0 · Windows 10 & 11
Flowframes is open-source donationware: the newest early-access builds reach supporters first, and the source is always public. Want the very latest version? Support development and get early-access builds on Patreon .
Flowframes runs interpolation locally on the Windows PC where it is installed. Older graphics hardware may lack the Vulkan features NCNN needs.
A Vulkan-capable graphics card
A supported Nvidia, AMD or Intel graphics driver
Storage for the source, temporary frames and exported result
A modern CUDA-capable Nvidia GPU
A modern multi-core Intel Core or AMD Ryzen CPU
Solid-state storage for frame extraction and encoding
Not supported: GPUs such as the Nvidia GT 710 and Intel HD 3000/4000 series, and Windows 7.
When frame interpolation works best
AI interpolation performs best when movement is visible and reasonably predictable between adjacent frames. It creates plausible new frames - it cannot recover detail that was never in the source.
Camera movement with consistent motion
Anime and hand-drawn animation (after correct de-duplication)
AI-generated clips with low frame rates
• Large objects appearing or disappearing between frames
• Fast cuts that are not detected correctly
• Flashing lights or rapid exposure changes
• Reflections, particles, smoke, water or transparent objects
• Severe compression artifacts or corrupted source timing
Flowframes is a Windows application for AI-powered video frame interpolation. It analyses the motion between existing frames and generates new intermediate frames to create smoother playback, higher frame rates and more natural slow motion. It is built around advanced models such as RIFE and supports both CUDA and Vulkan-based processing.
Flowframes can increase the frame rate of video, prepare footage for slow motion, smooth 2D animation, improve AI-generated clips, interpolate gameplay recordings, create seamless animated loops and generate missing frames for suitable 3D rendering workflows. It works with complete videos, animated GIFs and image sequences.
No. Flowframes increases temporal resolution by generating extra frames between the original frames. It does not increase the width, height or spatial detail of the source. A 1080p video stays 1080p - the difference is that motion looks smoother because more frames are shown each second.
Yes. Flowframes can generate additional frames and export at a higher frame rate, including common targets such as 48, 60, 120 or 240 FPS. You can also set a maximum output frame rate - for example, generating frames from a 24 FPS source and downsampling the result to exactly 60 FPS.
RIFE is the recommended starting point for most videos, with a strong balance of quality, speed and hardware compatibility. Nvidia systems can use RIFE CUDA for faster processing; RIFE NCNN uses Vulkan and supports a wider range of GPUs. DAIN, FLAVR and XVFI are available for comparison and specific footage.
RIFE CUDA is designed for compatible Nvidia graphics cards and is generally the faster option on supported hardware. RIFE NCNN runs through Vulkan and offers broader GPU compatibility, making it the right choice for AMD systems and for Nvidia systems that do not use the CUDA path. The intended result is similar; performance, memory use and model availability may differ.
Yes. Compatible AMD GPUs can use Vulkan-based NCNN models such as RIFE NCNN. CUDA modes are exclusive to Nvidia hardware, but an Nvidia GPU is not required to use Flowframes - the graphics card just needs to support the Vulkan features the selected NCNN model requires.
Standard builds include the components required for normal use, and the installer can download necessary dependencies automatically. Most people do not need to install Python, PyTorch, FFmpeg or AI models by hand. Advanced users can configure a system Python environment or compile the open-source project themselves.
Flowframes reads widely used formats such as MP4, MKV, MOV, WEBM, GIF and BIK, and can import image sequences in PNG and JPEG. Results export as MP4, MKV, MOV, WEBM, GIF or a sequence of frames, with codecs including H.264, H.265/HEVC, VP9 and AV1 depending on the encoding configuration.
Flowframes can transfer supported audio tracks, subtitle tracks and compatible metadata from the source to the exported video. Streams can often be copied without quality loss when the chosen output container supports them; compatibility depends on the source format, output format and encoding settings.
Frame de-duplication detects and removes consecutive duplicate frames before interpolation. It is mainly for 2D animation, anime and motion graphics that intentionally hold the same image across several frames. Removing those repeats lets the model work with real movement instead of identical images. It should usually be disabled for camera footage, gameplay and 3D-rendered video.
Choppy results are often caused by incorrect frame de-duplication. Dark, low-contrast or slowly changing frames can be wrongly detected as duplicates and removed. For

[truncated]
