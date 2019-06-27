Feature: Subtraction happens with addition

  Scenario:
    Given calculator is cleared
    When i press 2
    And i add 5
    And i subtract 3
    And i press 10
    And i add 2
    And i multiply by 2
    Then the result should be 24

