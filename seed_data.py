import requests
import random
from datetime import datetime, timedelta

# ===== CONFIGURATION =====

API_BASE = "http://localhost:8080/api"
AUTH_TOKEN = "SjrQgqqNSjW4Ty09gswNf0EPqjwFHtrO"  # Replace this before running

HEADERS = {
    "Content-Type": "application/json",
    "Authorization": f"Bearer {AUTH_TOKEN}"
}

# Existing references
existing_team_ids = [1, 2]
existing_game_ids = [1]
existing_set_ids = [1, 2]
existing_player_ids = [1]
round_id = 1
user_id = 1

# ===== STEP 1: Create 2 new teams =====

team_names = ["Кронверские Барсы", "Северные Волки"]
new_team_ids = []

for name in team_names:
    team_payload = {
        "name": name,
        "user_id": user_id
    }
    res = requests.post(f"{API_BASE}/teams", json=team_payload, headers=HEADERS)
    res.raise_for_status()
    team = res.json()
    new_team_ids.append(team["team_id"])

all_team_ids = existing_team_ids + new_team_ids

# ===== STEP 2: Create 4 new games =====

new_game_ids = []

for i in range(4):
    team1, team2 = random.sample(all_team_ids, 2)
    date = (datetime(2025, 7, 1) + timedelta(days=i)).isoformat() + "Z"

    game_payload = {
        "date": date,
        "round_id": round_id,
        "team_id": team1,
        "opp_team_id": team2
    }

    res = requests.post(f"{API_BASE}/games", json=game_payload, headers=HEADERS)
    res.raise_for_status()
    game = res.json()
    new_game_ids.append(game["game_id"])

# ===== STEP 3: Create 3–5 sets per game =====

new_set_ids = []

for game_id in new_game_ids:
    for serial_number in range(1, random.randint(4, 6)):
        set_payload = {
            "serial_number": serial_number,
            "game_id": game_id
        }
        res = requests.post(f"{API_BASE}/sets", json=set_payload, headers=HEADERS)
        res.raise_for_status()
        new_set = res.json()
        new_set_ids.append(new_set["set_id"])

all_set_ids = existing_set_ids + new_set_ids

# ===== STEP 4: Create 6–10 players per team =====

player_number = 2
amplua_ids = [1, 2, 3, 4, 5]
team_to_players = {}

for team_id in all_team_ids:
    team_to_players[team_id] = []
    for _ in range(random.randint(6, 10)):
        player_payload = {
            "first_name": f"Имя{player_number}",
            "last_name": f"Фамилия{player_number}",
            "birthdate": "2006-01-01T00:00:00Z",
            "gender": "male",
            "height": round(random.uniform(1.75, 2.05), 2),
            "number": player_number,
            "team_id": team_id,
            "amplua_id": random.choice(amplua_ids)
        }
        res = requests.post(f"{API_BASE}/players", json=player_payload, headers=HEADERS)
        res.raise_for_status()
        player = res.json()
        team_to_players[team_id].append((player["player_id"], player_number))
        player_number += 1

# ===== STEP 5: Create random set actions =====

action_map = {
    "A": "Атака",
    "B": "Блок",
    "C": "Защита",
    "S": "Подача",
    "R": "Прием"
}

rate_map = {
    "Подача": ["++", "_", "-"],
    "Блок": ["++", "-"],
    "Атака": ["++", "+-", "_", "-"],
    "Защита": ["++", "-"],
    "Прием": ["++", "+", "_", "-"]
}

for set_id in new_set_ids:
    # Get game ID from set
    res = requests.get(f"{API_BASE}/sets/{set_id}", headers=HEADERS)
    res.raise_for_status()
    set_info = res.json()
    game_id = set_info["game_id"]

    # Get teams from game
    res = requests.get(f"{API_BASE}/games/{game_id}", headers=HEADERS)
    res.raise_for_status()
    game_info = res.json()
    team_id = game_info["team_id"]

    players = team_to_players.get(team_id, [])
    if not players:
        continue

    for player_id, number in random.sample(players, min(5, len(players))):
        for _ in range(random.randint(2, 4)):
            abbr = random.choice(list(action_map.keys()))
            action_name = action_map[abbr]
            signature = random.choice(rate_map[action_name])
            record = f"{number}{abbr}{signature}"

            action_payload = {
                "set_id": set_id,
                "record": record
            }

            res = requests.post(f"{API_BASE}/record/action", json=action_payload, headers=HEADERS)
            res.raise_for_status()
