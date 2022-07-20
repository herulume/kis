package pipeline

import (
	"errors"

	"github.com/herulume/kis/pkg/nonempty"
)

type Build struct {
	pipeline Pipeline
	state    BuildState
}

func NewBuild(pipeline Pipeline) Build {
	return Build{
		pipeline: pipeline,
		state:    BuildReady,
	}
}

type BuildState int64

const (
	BuildReady BuildState = iota
	BuildRunning
	BuildSucceeded
	BuildFailed
)

func (bs BuildState) String() string {
	switch bs {
	case BuildReady:
		return "Build Ready"
	case BuildFailed:
		return "Build Failed"
	case BuildRunning:
		return "Build Running"
	case BuildSucceeded:
		return "Build Succeeded"
	}
	return "unknown state"
}

type Pipeline struct {
	steps nonempty.NonEmpty[Step]
}

func NewPipeline(steps nonempty.NonEmpty[Step]) Pipeline {
	return Pipeline{
		steps: steps,
	}
}

type Step struct {
	name     string
	image    string
	commands nonempty.NonEmpty[string]
}

func NewStep(name, image string, commands nonempty.NonEmpty[string]) (Step, error) {
	if image == "" || name == "" {
		return Step{}, errors.New("empty argument")
	}

	return Step{
		name:     name,
		image:    image,
		commands: commands,
	}, nil
}
