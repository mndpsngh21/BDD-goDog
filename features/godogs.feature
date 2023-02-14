Feature: eat mangoes
  In order to be happy
  As a hungry Person
  I need to be able to eat mango

  Scenario: Eat 5 out of 12
    Given there are 12 mangoes
    When I eat 5
    Then there should be 7 remaining

  Scenario: Eat 7 out of 12
    Given there are 12 mangoes
    When I eat 7
    Then there should be 5 remaining

  Scenario: Eat 12 out of 12
    Given there are 12 mangoes
    When I eat 12
    Then there should be none remaining

  Scenario: Add after eating 
    Given there are 12 mangoes
    When I eat 12
    And I Bought 1 mangoes
    Then there should be 1 remaining  

  Scenario: New Purchase
    Given there are 0 mangoes
    And I Bought 1 mangoes
    Then there should be 1 remaining    

  Scenario: First Buy then eat
    Given there are 0 mangoes
    And I Bought 5 mangoes
    When I eat 1
    Then there should be 4 remaining    


  Scenario: Eat, Buy and Throw
    Given there are 0 mangoes
    And I Bought 5 mangoes
    When I eat 1
    And Found Bad 3 mangoes
    Then there should be 1 remaining   
