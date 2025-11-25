#!/bin/bash

# Test script for the player search feature

echo "Starting server..."
PORT=8080 ./target/bin/server &
SERVER_PID=$!

# Wait for server to start
sleep 3

echo "Testing player search page..."
curl -s http://localhost:8080/player-search > /tmp/player_search_page.html
if grep -q "Player Career Search" /tmp/player_search_page.html; then
    echo "✓ Player search page loads correctly"
else
    echo "✗ Player search page did not load correctly"
fi

echo ""
echo "Testing player search API..."
curl -s "http://localhost:8080/player-search/search?query=Crosby" > /tmp/player_search_results.html
if grep -q "Sidney Crosby" /tmp/player_search_results.html; then
    echo "✓ Player search API works - found Sidney Crosby"
else
    echo "✗ Player search API did not return expected results"
fi

echo ""
echo "Testing player career API..."
curl -s "http://localhost:8080/player-search/career?playerId=8471675" > /tmp/player_career.html
if grep -q "Pittsburgh Penguins" /tmp/player_career.html; then
    echo "✓ Player career API works - shows Pittsburgh Penguins"
else
    echo "✗ Player career API did not return expected results"
fi

echo ""
echo "Testing team roster page (existing feature)..."
curl -s http://localhost:8080/roster > /tmp/roster_page.html
if grep -q "Team Roster" /tmp/roster_page.html; then
    echo "✓ Team roster page still works"
else
    echo "✗ Team roster page broken"
fi

# Cleanup
echo ""
echo "Stopping server..."
kill $SERVER_PID

echo ""
echo "All tests completed!"

