import random

def monty_hall_one_round():
    """
    Simulates one round of the Monty Hall problem with the switching strategy.
    
    Returns:
        bool: True if the player wins (guesses correctly after switching), False otherwise
    """
    # Randomly place the prize behind one of the three doors (1, 2, or 3)
    prize_door = random.randint(1, 3)
    
    # Player makes their first choice randomly
    player_first_choice = random.randint(1, 3)
    
    # Monty reveals a door that is not the prize and not the player's choice
    # Find the door to reveal (the one that's neither the prize nor the player's choice)
    doors = [1, 2, 3]
    doors.remove(prize_door)
    if player_first_choice in doors:
        doors.remove(player_first_choice)
    revealed_door = random.choice(doors)
    
    # Player switches to the remaining door
    remaining_doors = [1, 2, 3]
    remaining_doors.remove(player_first_choice)
    remaining_doors.remove(revealed_door)
    player_final_choice = remaining_doors[0]
    
    # Check if the player wins
    return player_final_choice == prize_door

def run_multiple_rounds(num_rounds=10):
    """
    Runs multiple rounds of the Monty Hall problem and counts wins.
    
    Args:
        num_rounds (int): Number of rounds to simulate (default: 10)
    
    Returns:
        tuple: (number of wins, total rounds, win percentage)
    """
    wins = 0
    for _ in range(num_rounds):
        if monty_hall_one_round():
            wins += 1
    
    win_percentage = (wins / num_rounds) * 100
    return wins, num_rounds, win_percentage

# Test the function
if __name__ == "__main__":
    # Test single round
    result = monty_hall_one_round()
    print(f"Single round - Player wins: {result}")
    
    # Test multiple rounds
    wins, total, percentage = run_multiple_rounds(1000)
    print(f"\nMultiple rounds - Wins: {wins}/{total} ({percentage:.1f}%)")
