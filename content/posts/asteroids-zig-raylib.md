---
cover: /static/images/.webp
date: 2024-08-15
summary: How to build a game with raylib+zig+emscripten.
tags:
  - "gamedev"
  - "raylib"
  - "zig"
title: a weekend love story - raylib/zig
---

![raylib-zig-emscripten](/static/images/zig-raylib/steroids.zig.webp)

### Zig Psyop

Before this weekend, I was a plebeian who used JavaScript for 80% of my tasks. Now I am an esoteric plebeian who has used zig once. Anyway, I decided to give zig a shot and try to build a game with it. There was really only two libraries I wanted to learn (sokol and raylib), I went with raylib.
<br/>

### Initial References

I went through the official raylib [docs](https://www.raylib.com) to get a feel for the API, then I searched for zig bindings and found two, [Not-Nik/raylib-zig](https://github.com/Not-Nik/raylib-zig) and [ryupold/raylib-zig](https://github.com/ryupold/raylib.zig)

The second one hasn't been updated in 6 months, and it isn't using the zig package manager. Ergo, leaving me no choice but to use the first one.

### Setup

Fairly easy to set things up, I just needed to install zig and raylib-zig. 

> *NOTE* : We will be using zig v0.12.0 instead of v0.13.0, we will see why later.


I used [zvm](https://github.com/tristanisham/zvm) to make life easier mangaging various zig versions, but you can install zig however you want.

Follow this [part](https://github.com/Not-Nik/raylib-zig#installation) of the docs to install raylib-zig. Do this inside your game directory. This adds raylib-zig as a dependency to our project. You can check the `build.zig.zon` file to verify.

<br/>

Now the most important part if you want to compile your game for the web with emscripten. *IF YOU'RE NOT BUILDING FOR THE WEB, PLEASE SKIP*.

<br/>

Install `emsdk` as mentioned here: [emscripten installation guide](https://emscripten.org/docs/getting_started/downloads.html) and make sure to do the `source ./emsdk_env.sh` command.

> *NOTE* : Just make sure not to clone the emsdk into $HOME/.emscripten, because emscripten uses that as the default cache directory. It will fuck your build up. (I didn't do this at all)

If you followed the raylib-zig installation guide, your build.zig should look something like this:

<br/>

```zig
const std = @import("std");

pub fn build(b: *std.Build) !void {
    const target = b.standardTargetOptions(.{});

    const optimize = b.standardOptimizeOption();

    const exe = b.addExecutable(.{
        .name = "steroids.zig",
        .root_source_file = b.path("src/main.zig"),
        .target = target,
        .optimize = optimize,
    });

    const raylib_dep = b.dependency("raylib-zig", .{
        .target = target,
        .optimize = optimize,
    });

    const raylib = raylib_dep.module("raylib");
    const raylib_artifact = raylib_dep.artifact("raylib");

    exe.linkLibrary(raylib_artifact);
    exe.root_module.addImport("raylib", raylib);

    b.installArtifact(exe);

    const run_cmd = b.addRunArtifact(exe);

    run_cmd.step.dependOn(b.getInstallStep());

    if (b.args) |args| {
        run_cmd.addArgs(args);
    }

    const run_step = b.step("run", "Run the app");
    run_step.dependOn(&run_cmd.step);
}
```

<br/>

If it isn't, it should be now! Now we need to add a block specific to the emscripten target.

<br/>

```zig
// add these two imports
const fs = std.fs;
const rlz = @import("raylib-zig");

pub fn build(b: *std.Build) !void {
    // ... other code

    const raylib = raylib_dep.module("raylib");
    const raylib_artifact = raylib_dep.artifact("raylib");

    if (target.query.os_tag == .emscripten) {
        const exe_lib = rlz.emcc.compileForEmscripten(b, "steroids.zig", "src/main.zig", target, optimize);

        exe_lib.linkLibrary(raylib_artifact);
        exe_lib.root_module.addImport("raylib", raylib);

        const include_path = try fs.path.join(b.allocator, &.{ b.sysroot.?, "cache", "sysroot", "include" });
        defer b.allocator.free(include_path);
        exe_lib.addIncludePath(.{ .path = include_path });
        exe_lib.linkLibC();

        // linking raylib to the exe_lib output file.
        const link_step = try rlz.emcc.linkWithEmscripten(b, &[_]*std.Build.Step.Compile{ exe_lib, raylib_artifact });

        // Use the custom HTML template
        // This will be the index.html where the game is rendered.
        // You can find an example in my repository.
        link_step.addArg("--shell-file");
        link_step.addArg("shell.html");

        // Embed the assets directory
        // This generates an index.data file which is neede for the game to run.
        link_step.addArg("--preload-file");
        link_step.addArg("assets");

        link_step.addArg("-sALLOW_MEMORY_GROWTH");
        link_step.addArg("-sWASM_MEM_MAX=16MB");
        link_step.addArg("-sTOTAL_MEMORY=16MB");
        link_step.addArg("-sERROR_ON_UNDEFINED_SYMBOLS=0");
        // Add any other flags you need

        b.getInstallStep().dependOn(&link_step.step);
        const run_step = try rlz.emcc.emscriptenRunStep(b);
        run_step.step.dependOn(&link_step.step);
        const run_option = b.step("run", "Run the game");
        run_option.dependOn(&run_step.step);
        return;
    }

    // rest of the build script
}
```
<br/>

I missed this code block and was stuck trying to figure out how to get the emscripten run step to work. :(

### Write some code!

I wrote an asteroids clone following along with this video: [zig space rocks - jdh](https://www.youtube.com/live/ajbYYgbDXGk?si=J4GzjKmHucSjhkaT). Go write whatever game you want!

<br/>

Here is a screenshot of the game in action:
<br/>

![space rocks](/static/images/zig-raylib/ingame.webp)

<br/>

Here are some other cool games made in zig: [tetris](https://github.com/jlafayette/zig-tetris/blob/master/src/main.zig), [terrain-zigger](https://github.com/JosefAlbers/TerrainZigger)

<br/>

Below is some cool zig code I wrote that I quite like:

<br/>

```zig
rl.initAudioDevice();
defer rl.closeAudioDevice();
sound = Sound{
    .asteroid = rl.loadSound("assets/asteroid.wav"),
    .bloop_lo = rl.loadSound("assets/bloop_lo.wav"),
    .bloop_hi = rl.loadSound("assets/bloop_hi.wav"),
    .explode = rl.loadSound("assets/explode.wav"),
    .shoot = rl.loadSound("assets/shoot.wav"),
    .thrust = rl.loadSound("assets/thrust.wav"),
};

rl.setSoundVolume(sound.explode, 0.5);

// this unwraps the loop at compile time for each field in the Sound struct
// and frees the memory allocated for each sound
defer inline for (std.meta.fields(Sound)) |f| {
    rl.unloadSound(@field(sound, f.name));
};

defer state.asteroids.deinit();
defer state.asteroids_q.deinit();
defer state.particles.deinit();
defer state.projectiles.deinit();
defer state.aliens.deinit();
```

<br/>

This `defer` syntax is so fucking cool, I will never forget to `free` again.

<br/>

```zig
const bloop_intensity: usize = @min(@as(usize, @intFromFloat(state.now - state.stage_start)) / 16, 4);
const bloop_mod: usize = 60;
const adjusted_bloop_mod: usize = @max(1, bloop_mod >> @intCast(bloop_intensity));

if (state.frame % adjusted_bloop_mod == 0) {
    state.bloop += 1;
}

if (!state.ship.isDead() and state.bloop != state.last_bloop) {
    rl.playSound(if (state.bloop % 2 == 0) sound.bloop_hi else sound.bloop_lo);
}
state.last_bloop = state.bloop;
```

<br/>

This code is for deciding when to play the low/high bloop sound. It increases in intensity the longer a player is alive. Makes it feel like the game is getting harder. Pretty cool!
Also avoided for loops with bit shifting. absolutely unnecessary. But I did it anyway.

### Memory management

One thing I think is important is for some reason I wasn't able to just use a `std.heap.GeneralPurposeAllocator` for the heap allocations. So i reused some code from the [ryupold/examples-raylib.zig](https://github.com/ryupold/examples-raylib.zig/blob/main/src/allocator.zig) repo and made a custom emscripten allocator.

<br/>

```zig
// i modified only one line
// line 50
const c = @cImport({
    @cDefine("__EMSCRIPTEN__", "1");
    @cInclude("emscripten.h"); // <- I changed it to just emscripten.h
    @cInclude("stdlib.h");
});
```
<br/>

Without this, I was getting `Out of Memory` issues. I also increased the memory limit to 16MB. Idk which of those fixed it, but it works now! After that, I was sitting at a solid `1.23 MB` of memory usage, so I think I was good.

<br/>

![memory usage](/static/images/zig-raylib/heaptrace.webp)

<br/>

### Build the game

You can build the game for desktop with `zig build run`. For the web, you need to run the following command:

<br/>

```bash
zig build -Dtarget=wasm32-emscripten --sysroot "/absolute/path/to/emsdk/upstream/emscripten"
# the absolute path is crucial, otherwise it will not work.
# now to run the game, you can use the following command:
emrun ./zig-out/htmlout/index.html
```

<br/>

Now, as to why we used `v0.12.0`. For some reason the build command above fails for the emscripten target with `v0.13.0` as mentioned in this open [issue](https://github.com/Not-Nik/raylib-zig/issues/108). If someone can figure out why, please mention it in the issue!

<br/>

You can drop everything in the `zig-out/htmlout` folder and host it wherever.

### Fin!

Was a fun start to my zig arc! I hope you guys don't waste time debugging shit like me lol.

May we zig harder every day.

<br/>

[Source Code](https://github.com/seatedro/steroids.zig)
