def generate_all_possible_games(max_red, max_green, max_blue):
    games = []
    game_id = 1

    for red in range(max_red + 1):
        for green in range(max_green + 1):
            for blue in range(max_blue + 1):
                game_str = f"Game {game_id}: {red} red, {green} green, {blue} blue"
                games.append(game_str)
                game_id += 1

    return games

def calculate_sum_of_ids(max_red, max_green, max_blue):
    total_games = (max_red + 1) * (max_green + 1) * (max_blue + 1)
    sum_of_ids = total_games * (total_games + 1) // 2  # Sum of the first N natural numbers
    return sum_of_ids

# Generate all possible games within the constraints
max_red, max_green, max_blue = 12, 13, 14
sum_ids = calculate_sum_of_ids(max_red, max_green, max_blue)

print("Sum of IDs of all possible games:", sum_ids)

