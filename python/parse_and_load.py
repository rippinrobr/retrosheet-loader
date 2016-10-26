import hall_of_fame
import managers
import players
import season
import parks
import home_games

print "### PROCESSING MANAGER DATA ###"
raw_managers = managers.parse()

print "### PROCESSING PLAYER DATA ###"
raw_players = players.parse()

print "### PROCESSING SEASON  DATA ###"
raw_teams, raw_seasons, raw_franchises = season.parse()

print "### PROCESSING HOF DATA ###"
raw_hof = hall_of_fame.parse()

print "### PROCESSING PARKS DATA ###"
raw_parks = parks.parse()

print "### PROCESSING HOME GAMES DATA ###"
raw_home_games = home_games.parse()
