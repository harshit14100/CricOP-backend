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

ALTER TABLE deliveries
    ADD CONSTRAINT deliveries_striker_id_fkey
        FOREIGN KEY (striker_id)
            REFERENCES users(id);

ALTER TABLE deliveries
    ADD CONSTRAINT deliveries_non_striker_id_fkey
        FOREIGN KEY (non_striker_id)
            REFERENCES users(id);

ALTER TABLE deliveries
    ADD CONSTRAINT deliveries_bowler_id_fkey
        FOREIGN KEY (bowler_id)
            REFERENCES users(id);

ALTER TABLE deliveries
    ADD CONSTRAINT deliveries_fielder_id_fkey
        FOREIGN KEY (fielder_id)
            REFERENCES users(id);

ALTER TABLE deliveries
    ADD CONSTRAINT deliveries_player_out_id_fkey
        FOREIGN KEY (player_out_id)
            REFERENCES users(id);
