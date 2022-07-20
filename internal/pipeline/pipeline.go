package pipeline

import (
	"errors"

	"github.com/herulume/kis/pkg/nonempty"
)

type Build struct {
	pipeline Pipeline
}

func NewBuild(pipeline Pipeline) Build {
	return Build{
		pipeline: pipeline,
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

func NewPipeline(steps []Step) (Pipeline, error) {
	s, err := nonempty.NewNonEmpty(steps)
	if err != nil {
		return Pipeline{}, errors.New("no steps provided")
	}
	return Pipeline{
		steps: s,
	}, nil
}

type Step struct {
	name     string
	image    string
	commands nonempty.NonEmpty[string]
}

func NewStep(name, image string, commands []string) (Step, error) {
	if image == "" || name == "" {
		return Step{}, errors.New("empty argument")
	}

	c, err := nonempty.NewNonEmpty(commands)
	if err != nil {
		return Step{}, errors.New("no commands provided")
	}
	return Step{
		name:     name,
		image:    image,
		commands: c,
	}, nil
}
