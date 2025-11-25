# Small App to Show Hockey Rosters

## Features

### Team Roster
View current and historical rosters for any NHL team. Select a team, season, and sort players by various attributes like number, name, position, height, weight, or age.

### Player Career Search
Search for any NHL player by name (or partial name) to view their complete career history:
- Search for players across all teams
- View which teams a player has played for throughout their career
- See season-by-season statistics including:
  - Teams and leagues played in
  - Games played, goals, assists, points (for skaters)
  - Wins, losses, GAA, save percentage (for goalies)

## Usage

### Command Line

1. Build the application:
   ```bash
   make build
   ```

2. Set the PORT environment variable and run:
   ```bash
   PORT=8080 ./target/bin/server
   ```

3. Navigate to:
   - `http://localhost:8080/` or `http://localhost:8080/roster` - Team roster viewer
   - `http://localhost:8080/player-search` - Player career search

### IDE (GoLand/IntelliJ IDEA)

Pre-configured run configurations are available:
- **Run Server** - Start the server on port 8080
- **Run Server (Port 3000)** - Alternative port configuration
- **Test All** - Run all tests
- **Test Player Search** - Run player search feature tests

Simply select a configuration from the dropdown and click the Run button (▶).

See `.idea/runConfigurations/README.md` for details.

## API Endpoints

- `GET /roster` - Team roster page
- `GET /roster/players-for-team` - Fetch players for selected team/season
- `GET /player-search` - Player search page
- `GET /player-search/search?query=<name>` - Search for players by name
- `GET /player-search/career?playerId=<id>` - View player career details

