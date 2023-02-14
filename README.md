# BDD-goDog


## Initialize
  ### execute for go setup 
  go mod init 

<br/>

## Install GoDog

  go install github.com/cucumber/godog/cmd/godog@latest

## create Feature File
<pre>
Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining

</pre>

## execute godog run for code suggestions
<pre>
 go run
</pre>


## create test file and copy generated code in that


## Our module should now look like this:
<pre>
MainFolder
- features
  - godogs.feature
- go.mod
- go.sum
- godogs_test.go
</pre>

## Add Main code
<pre>
package main

// Godogs available to eat
var Godogs int

func main() { /* usual main func */ }
</pre>


## Our module should now look like this:
<pre>
MainFolder
- features
  - godogs.feature
- go.mod
- go.sum
- godogs_test.go
- godog.go
</pre>

## Run Test

<pre>
Run code from main folder
    godog run 

<pre>