import random

def generate_random_good_games(num_games, max_id):
    def random_subset(max_red, max_green, max_blue):
        red = random.randint(0, max_red)
        green = random.randint(0, max_green)
        blue = random.randint(0, max_blue)
        return f"{red} red, {green} green, {blue} blue"

    games = []
    for _ in range(num_games):
        game_id = random.randint(1, max_id)
        subsets = [random_subset(12, 13, 14) for _ in range(random.randint(1, 5))]
        game_str = f"Game {game_id}: " + "; ".join(subsets)
        games.append(game_str)

    return games

def sum_of_possible_game_ids(games):
    possible_games = []

    for game in games:
        parts = game.split(": ")
        game_id = int(parts[0].split(" ")[1])  # Correct parsing of the game ID
        game_data = parts[1]
        max_red, max_green, max_blue = 0, 0, 0

        for subset in game_data.split("; "):
            cubes = subset.split(", ")
            red, green, blue = 0, 0, 0

            for cube in cubes:
                count, color = cube.split(" ")
                count = int(count)

                if "red" in color:
                    red = max(red, count)
                elif "green" in color:
                    green = max(green, count)
                elif "blue" in color:
                    blue = max(blue, count)

            max_red, max_green, max_blue = max(max_red, red), max(max_green, green), max(max_blue, blue)

        if max_red <= 12 and max_green <= 13 and max_blue <= 14:
            possible_games.append(game_id)

    return sum(possible_games)

# Generate a set of random games
random_games = generate_random_good_games(10, 100)

# Calculate the sum of IDs of the generated games
sum_ids = sum_of_possible_game_ids(random_games)

print("Generated Games:")
for game in random_games:
    print(game)
print("\nSum of IDs of possible games:", sum_ids)

