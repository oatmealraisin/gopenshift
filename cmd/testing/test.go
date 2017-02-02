package main

import "github.com/oatmealraisin/gopenshift/pkg/gopenshift"

func main() {
	thing := gopenshift.New()

	thing.GetPods()

}
