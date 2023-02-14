package main

// This example shows how to set up test suite runner with Go subtests and godog command line parameters.
// Sample commands:
// * run all scenarios from default directory (features): go test -test.run "^TestFeatures/"
// * run all scenarios and list subtest names: go test -test.v -test.run "^TestFeatures/"
// * run all scenarios from one feature file: go test -test.v -godog.paths features/nodogs.feature -test.run "^TestFeatures/"
// * run all scenarios from multiple feature files: go test -test.v -godog.paths features/nodogs.feature,features/godogs.feature -test.run "^TestFeatures/"
// * run single scenario as a subtest: go test -test.v -test.run "^TestFeatures/Eat_5_out_of_12$"
// * show usage help: go test -godog.help
// * show usage help if there were other test files in directory: go test -godog.help godogs_test.go
// * run scenarios with multiple formatters: go test -test.v -godog.format cucumber:cuc.json,pretty -test.run "^TestFeatures/"

import (
	"context"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestFeatures(t *testing.T) {
	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		Name:                 "Mangoes",
		Options:              &o,
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}

func thereAreMangoes(available int) error {
	Mangoes = available
	return nil
}

func iEat(num int) error {
	if Mangoes < num {
		return fmt.Errorf("you cannot eat %d mangoes, there are %d available", num, Mangoes)
	}
	Mangoes -= num
	return nil
}

func thereShouldBeRemaining(remaining int) error {
	if Mangoes != remaining {
		return fmt.Errorf("expected %d mangoes to be remaining, but there is %d", remaining, Mangoes)
	}
	return nil
}

func thereShouldBeNoneRemaining() error {
	return thereShouldBeRemaining(0)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() { Mangoes = 0 })
}

func shouldBuyGoDogs(buy int) {
	Mangoes += buy
}

func removeBadMangoes(badMangoes int) {
	Mangoes -= badMangoes
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		Mangoes = 0 // clean the state before every scenario
		return ctx, nil
	})

	ctx.Step(`^there are (\d+) mangoes$`, thereAreMangoes)
	ctx.Step(`^I eat (\d+)$`, iEat)
	ctx.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
	ctx.Step(`^there should be none remaining$`, thereShouldBeNoneRemaining)
	ctx.Step(`^I Bought (\d+) mangoes`, shouldBuyGoDogs)
	ctx.Step(`^Found Bad (\d+) mangoes`, removeBadMangoes)

}
