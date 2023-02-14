Feature: eat mangoes
  In order to be happy
  As a hungry gopher
  I need to be able to eat mango

  Scenario: Eat 5 out of 12
    Given there are 12 mango
    When I eat 5
    Then there should be 7 remaining

  Scenario: Eat 12 out of 12
    Given there are 12 mango
    When I eat 12
    Then there should be none remaining

  Scenario: Add after eating 
    Given there are 12 mango
    When I eat 12
    And I Bought 1 mango
    Then there should be 1 remaining  

  Scenario: Just Main
    Given there are 0 mango
    And I Bought 1 mango
    Then there should be 1 remaining    