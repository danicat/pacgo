# Step 10: Ghosts with power ups!

In this lesson you will learn how to:

- Use a timer 
- When and how to use a Mutex lock

## Overview

In this lesson we will be adding support for the power up pill to the application. We will update the configuration with the new setups and add code to draw the pill in the maze. We will also manage the process after pacman swallows a pill and collides with a ghost. Finally, we will manipulate cases where pacman tries to swallow a pill while the previous one is active and how to tackle this.

## Task 01: Drawing the Pills
Before we even start we should update the configuration to support power up pills! So for both `config_noemoji.json` and `config.json` we have to add the `ghost_blue` (string) and the `pill_duration_secs` (int) configurations.
Accordingly we update our `Config` struct:
```go
type config struct {
    ...
	GhostBlue        string        `json:"ghost_blue"`
	PillDurationSecs time.Duration `json:"pill_duration_secs"`
}
```

## Task 02: Enable Pill swallowing
To enable the pill swallowing by pacman we should add another case in the `movePlayer` func for the pill case
```go
case 'X':
	score += 10
	removeDot(player.row, player.col)
	go processPill()
```
Where `X` is the pill config character. 

Now, before moving to the `processPill` func, we should add some more code for the ghosts to support the 'Blue Ghosts'! We should add a new `GhostStatus` of string type which will hold the status of a ghost. The two statuses we have to support are the `Normal` and the `Blue`.


```go
type GhostStatus string

const (
	GhostStatusNormal GhostStatus = "Normal"
	GhostStatusBlue   GhostStatus = "Blue"
)
```

Now, each ghost should hold alongside with it's current position, the `initialPosition` where it will be spawned after it's been eaten by the pacman and it's current status.

```go
type ghost struct {
	position sprite
	status   GhostStatus
}
```
So, the `loadMaze` func will initially draw the ghosts with the `Normal` status and store it's initial position.

```go
ghosts = append(ghosts, &ghost{sprite{row, col, row, col}, GhostStatusNormal})
```

The `printScreen` func should be updated as well to support printing ghost of both types - Normal and Blue ghosts!

```go
for _, g := range ghosts {
		moveCursor(g.position.row, g.position.col)
		if g.status == Normal {
			fmt.Printf(cfg.Ghost)
		} else if g.status == Blue {
			fmt.Printf(cfg.GhostBlue)
		}
	}
```

The last thing that has left is the `processPill` func we added just before. This func should change all Ghosts' status to `Blue` for the defined period by the `PillDurationSecs` config.
For the pill processing  we are going to use a `Timer` from the ['time' package](https://golang.org/pkg/time/). We will use the `NewTimer` func which creates a new Timer that will send the current time on its channel after at least the specified duration.

The processPill code changes all ghosts' statuses to `GhostStatusBlue`, then it blocks for `PillDurationSecs` and then changes back all ghosts' statuses back to `GhostStatusNormal`.

```go
var pillTimer *time.Timer

func processPill() {
	for _, g := range ghosts {
		g.status = GhostStatusBlue
	}
	pillTimer = time.NewTimer(time.Second * cfg.PillDurationSecs)
	<-pillTimer.C
    for _, g := range ghosts {
		g.status = GhostStatusNormal
    }
}
```

## Task 03: Support simultaneous pill swallowing
The `processPill` function we discussed just before has a simple issue. Imagine what can happen if pacman tries to swallow a power-up pill while another pill is still active! Currently, with the proposed `processPill` function, when a second pill is being swallowed by the pacman, while the first on is still active, when the 1st pill's effect ends (after PillDurationSecs) all ghosts will turn back to Normal. In order to overcome this, we should check if a pill is already active by checking the timer and then stopping it and re-initializing it if it's already active.

```go
var pillTimer *time.Timer

func processPill() {
	updateGhosts(ghosts, GhostStatusBlue)
	if pillTimer != nil {
		pillTimer.Stop()
	}
	pillTimer = time.NewTimer(time.Second * cfg.PillDurationSecs)
	<-pillTimer.C
	pillTimer.Stop()
	updateGhosts(ghosts, GhostStatusNormal)
}
```

## Task 04: Avoiding Race Conditions
In our scenarios there are two possible race conditions. The first one is about the pill timer we mentioned just before. The `processPill` function is called asynchronously. So in the case that the first `processPill` function is just after the `pillTimer.Stop()` while the second one is inside the `if pillTimer != nil {` block. In this rare case it seems that while one pill is active, consuming a next one while code is at this point we might loose the second pill as Ghosts will come back to normal. 

For this reason we are introducing a pillMx Mutex lock which we are going to acquire at the beginning of the `processPill` function and release just before starting to wait on the timer channel. Also we are going to acquire it just after the blocking function and release it at the end of the function.

```go
var pillTimer *time.Timer
var pillMx sync.Mutex

func processPill() {
	pillMx.Lock()
	updateGhosts(ghosts, GhostStatusBlue)
	if pillTimer != nil {
		pillTimer.Stop()
	}
	pillTimer = time.NewTimer(time.Second * cfg.PillDurationSecs)
	pillMx.Unlock()
	<-pillTimer.C
	pillMx.Lock()
	pillTimer.Stop()
	updateGhosts(ghosts, GhostStatusNormal)
	pillMx.Unlock()
}
```

Another possible race condition that might arise during execution is when we update the ghosts' status. For this purpose we are going to use a RWMutex lock. We have to acquire the lock whenever we read or update a ghost's status. RWMutex supports locking even for read or write access. So we are introducing the `var ghostsStatusMx sync.RWMutex` and a `updateGhosts` function that updates one or more ghost's status.

```go 
var ghostsStatusMx sync.RWMutex

func updateGhosts(ghosts []*Ghost, ghostStatus GhostStatus) {
	ghostsStatusMx.Lock()
	defer ghostsStatusMx.Unlock()
	for _, g := range ghosts {
		g.status = ghostStatus
	}
}
```

Also we have to acquire a RLock whenever we read a ghost's status. Multiple read locks can be acquire simultaneously but only one write lock can be acquired. We are going to use the `ghostsStatusMx.RLock()` and `ghostsStatusMx.RUnlock()` while reading the ghosts' status. We have to always unlock the RLock before updating a ghost's status otherwise a deadlock will occur.

Now we have a more challenging pacman! Happy gaming/coding! :) 

## That's All Folks!

Congratulations! You've completed all the steps of the tutorial.

But your journey must not end here. If you are interested in contributing with a new step, have a look at the [TODO list](../TODO.md) or any open issues and submit a PR!
