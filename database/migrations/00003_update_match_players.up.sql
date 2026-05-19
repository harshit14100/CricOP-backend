ALTER TABLE match_players
    DROP CONSTRAINT match_players_match_id_fkey;

ALTER TABLE match_players
    DROP CONSTRAINT match_players_team_id_fkey;

ALTER TABLE match_players
    ADD CONSTRAINT match_players_player_id_fkey
        FOREIGN KEY (player_id)
            REFERENCES players(id);

ALTER TABLE match_players
    ADD CONSTRAINT match_players_match_id_fkey
        FOREIGN KEY (match_id)
            REFERENCES matches(id);

ALTER TABLE match_players
    ADD CONSTRAINT match_players_team_id_fkey
        FOREIGN KEY (team_id)
            REFERENCES teams(team_id);
